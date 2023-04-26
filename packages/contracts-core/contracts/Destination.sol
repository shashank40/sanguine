// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "./libs/Attestation.sol";
import {AttestationReport} from "./libs/AttestationReport.sol";
import {ByteString} from "./libs/ByteString.sol";
import {AGENT_ROOT_OPTIMISTIC_PERIOD} from "./libs/Constants.sol";
import {AgentStatus, DestinationStatus, DisputeFlag} from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "./base/AgentSecured.sol";
import {DestinationEvents} from "./events/DestinationEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceDestination} from "./interfaces/InterfaceDestination.sol";
import {InterfaceLightManager} from "./interfaces/InterfaceLightManager.sol";
import {ExecutionHub} from "./hubs/ExecutionHub.sol";

contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    using AttestationLib for bytes;
    using ByteString for bytes;

    // TODO: this could be further optimized in terms of storage
    struct StoredAttData {
        bytes32 agentRoot;
        bytes32 r;
        bytes32 s;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    /// @dev Invariant: this is either current LightManager root,
    /// or the pending root to be passed to LightManager once its optimistic period is over.
    bytes32 public nextAgentRoot;

    /// @inheritdoc InterfaceDestination
    DestinationStatus public destStatus;

    /// @dev Stored lookup data for all accepted Notary Attestations
    StoredAttData[] internal _storedAttestations;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 domain, address agentManager_) AgentSecured("0.0.3", domain, agentManager_) {}

    /// @notice Initializes Destination contract:
    /// - msg.sender is set as contract owner
    function initialize(bytes32 agentRoot) external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        // Set Agent Merkle Root in Light Manager
        nextAgentRoot = agentRoot;
        InterfaceLightManager(address(agentManager)).setAgentRoot(agentRoot);
        destStatus.agentRootTime = uint40(block.timestamp);
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function acceptAttestation(
        address notary,
        AgentStatus memory status,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {
        // First, try passing current agent merkle root
        (bool rootPassed, bool rootPending) = passAgentRoot();
        // Don't accept attestation, if the agent root was updated in LightManager,
        // as the following agent check will fail.
        if (rootPassed) return false;
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        (bytes32 r, bytes32 s, uint8 v) = attSignature.castToSignature().toRSV();
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, status.index, v);
        bytes32 agentRoot = att.agentRoot();
        _storedAttestations.push(StoredAttData(agentRoot, r, s));
        // Save Agent Root if required, and update the Destination's Status
        // TODO: rework when signatures are stored in AgentManager
        destStatus = _saveAgentRoot(rootPending, agentRoot, _agentStatus(notary).index);
        emit AttestationAccepted(status.domain, notary, attPayload, attSignature);
        return true;
    }

    // ═══════════════════════════════════════════ AGENT ROOT QUARANTINE ═══════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function passAgentRoot() public returns (bool rootPassed, bool rootPending) {
        bytes32 oldRoot = IAgentManager(agentManager).agentRoot();
        bytes32 newRoot = nextAgentRoot;
        // Check if agent root differs from the current one in LightManager
        if (oldRoot == newRoot) return (false, false);
        DestinationStatus memory status = destStatus;
        // Invariant: Notary who supplied `newRoot` was registered as active against `oldRoot`
        // So we just need to check the Dispute status of the Notary
        if (_disputes[status.notaryIndex] != DisputeFlag.None) {
            // Remove the pending agent merkle root, as its signer is in dispute
            nextAgentRoot = oldRoot;
            return (false, false);
        }
        // Check if agent root optimistic period is over
        if (status.agentRootTime + AGENT_ROOT_OPTIMISTIC_PERIOD > block.timestamp) {
            // We didn't pass anything, but there is a pending root
            return (false, true);
        }
        // `newRoot` signer was not disputed, and the root optimistic period is over.
        // Finally, pass the Agent Merkle Root to LightManager
        InterfaceLightManager(address(agentManager)).setAgentRoot(newRoot);
        return (true, false);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    // solhint-disable-next-line ordering
    function attestationsAmount() external view returns (uint256) {
        return _roots.length;
    }

    /// @inheritdoc InterfaceDestination
    function getSignedAttestation(uint256 index)
        external
        view
        returns (bytes memory attPayload, bytes memory attSignature)
    {
        require(index < _roots.length, "Index out of range");
        bytes32 snapRoot = _roots[index];
        SnapRootData memory rootData = _rootData[snapRoot];
        StoredAttData memory storedAtt = _storedAttestations[index];
        attPayload = AttestationLib.formatAttestation({
            snapRoot_: snapRoot,
            agentRoot_: storedAtt.agentRoot,
            nonce_: rootData.attNonce,
            blockNumber_: rootData.attBN,
            timestamp_: rootData.attTS
        });
        attSignature = ByteString.formatSignature({r: storedAtt.r, s: storedAtt.s, v: rootData.notaryV});
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Saves Agent Merkle Root from the accepted attestation, if there is
    /// no pending root to be passed to LightManager.
    /// Returns the updated "last snapshot root / last agent root" status struct.
    function _saveAgentRoot(bool rootPending, bytes32 agentRoot, uint32 notaryIndex)
        internal
        returns (DestinationStatus memory status)
    {
        status = destStatus;
        // Update the timestamp for the latest snapshot root
        status.snapRootTime = uint40(block.timestamp);
        // Don't update agent root, if there is already a pending one
        // Update the data for latest agent root only if it differs from the saved one
        if (!rootPending && nextAgentRoot != agentRoot) {
            status.agentRootTime = uint40(block.timestamp);
            status.notaryIndex = notaryIndex;
            nextAgentRoot = agentRoot;
            emit AgentRootAccepted(agentRoot);
        }
    }
}

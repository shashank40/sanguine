// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainBatch, InterchainBatchLib} from "../../contracts/libs/InterchainBatch.sol";

contract InterchainBatchLibHarness {
    function constructLocalBatch(uint256 dbNonce, bytes32 batchRoot) external view returns (InterchainBatch memory) {
        return InterchainBatchLib.constructLocalBatch(dbNonce, batchRoot);
    }

    function batchKey(InterchainBatch memory batch) external pure returns (bytes32) {
        return InterchainBatchLib.batchKey(batch);
    }
}
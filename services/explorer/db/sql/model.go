package sql

import (
	"database/sql"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"math/big"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	TxHashFieldName = namer.GetConsistentName("TxHash")
	ChainIDFieldName = namer.GetConsistentName("ChainID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	ContractAddressFieldName = namer.GetConsistentName("ContractAddress")
	BlockIndexFieldName = namer.GetConsistentName("BlockIndex")
}

var (
	// TxHashFieldName is the field name of the tx hash.
	TxHashFieldName string
	// ChainIDFieldName gets the chain id field name.
	ChainIDFieldName string
	// BlockNumberFieldName is the name of the block number field.
	BlockNumberFieldName string
	// ContractAddressFieldName is the address of the contract.
	ContractAddressFieldName string
	// BlockIndexFieldName is the index field name.
	BlockIndexFieldName string
)

// SwapEvent stores data for emitted events from the Swap contract.
type SwapEvent struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// BlockNumber is the block number of the event
	BlockNumber uint64 `gorm:"column:block_number;primaryKey;auto_increment:false"`
	// TxHash is the transaction hash of the event
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type;primaryKey;auto_increment:false"`

	// TokenIndex is the index of the token in the pool
	TokenIndex *big.Int `gorm:"column:token_index;type:UInt256"`
	// Amount is the amount of tokens
	Amount *big.Int `gorm:"column:amount;type:UInt256"`
	// AmountFee is the amount of fees
	AmountFee *big.Int `gorm:"column:amount_fee;type:UInt256"`
	// ProtocolFee is the protocol fee
	ProtocolFee *big.Int `gorm:"column:protocol_fee;type:UInt256"`
	// Buyer is the address of the buyer
	Buyer sql.NullString `gorm:"column:buyer"`
	// TokensSold is the amount of tokens sold
	TokensSold *big.Int `gorm:"column:tokens_sold;type:UInt256"`
	// SoldID is the id of the token sold
	SoldID *big.Int `gorm:"column:sold_id;type:UInt256"`
	// TokensBought is the amount of tokens bought
	TokensBought *big.Int `gorm:"column:tokens_bought;type:UInt256"`
	// BoughtID is the id of the token bought
	BoughtID *big.Int `gorm:"column:bought_id;type:UInt256"`
	// Provider is the address of the provider
	Provider sql.NullString `gorm:"column:provider"`
	// TokenAmounts is the amounts of each token to transact
	TokenAmounts []*big.Int `gorm:"column:token_amounts;type:Array(UInt256)"`
	// Fees is the fees for each token
	Fees []*big.Int `gorm:"column:fees;type:Array(UInt256)"`
	// Invariant is the invariant of the pool
	Invariant *big.Int `gorm:"column:invariant;type:UInt256"`
	// LPTokenAmount is the amount of LP tokens
	LPTokenAmount *big.Int `gorm:"column:lp_token_amount;type:UInt256"`
	// LPTokenSupply is the supply of the LP token
	LPTokenSupply *big.Int `gorm:"column:lp_token_supply;type:UInt256"`
	// NewAdminFee is the new admin fee
	NewAdminFee *big.Int `gorm:"column:new_admin_fee;type:UInt256"`
	// NewSwapFee is the new swap fee
	NewSwapFee *big.Int `gorm:"column:new_swap_fee;type:UInt256"`
	// OldA is the old A value
	OldA *big.Int `gorm:"column:old_a;type:UInt256"`
	// NewA is the new A value
	NewA *big.Int `gorm:"column:new_a;type:UInt256"`
	// InitialTime is the initial time
	InitialTime *big.Int `gorm:"column:initial_time;type:UInt256"`
	// FutureTime is the future time
	FutureTime *big.Int `gorm:"column:future_time;type:UInt256"`
	// CurrentA is the current A value
	CurrentA *big.Int `gorm:"column:current_a;type:UInt256"`
	// Time is the time
	Time *big.Int `gorm:"column:time;type:UInt256"`
	// Receiver is the address of the receiver
	Receiver sql.NullString `gorm:"column:receiver"`
	// TokenID is the token's ID
	TokenID sql.NullString `gorm:"column:token_id"`
}

// BridgeEvent stores data for emitted events from the Bridge contract.
type BridgeEvent struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type;primaryKey;auto_increment:false"`
	// BlockNumber is the block number of the event
	BlockNumber uint64 `gorm:"column:block_number;primaryKey;auto_increment:false"`
	// TxHash is the transaction hash of the event
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// Token is the address of the token
	Token string `gorm:"column:token"`
	// Amount is the amount of tokens
	Amount *big.Int `gorm:"column:amount;type:UInt256"`

	// Recipient is the address to send the tokens to
	Recipient sql.NullString `gorm:"column:recipient"`
	// RecipientBytes is the recipient address in bytes
	RecipientBytes sql.NullString `gorm:"column:recipient_bytes"`
	// DestinationChainID is the chain id of the chain to send the tokens to
	DestinationChainID *big.Int `gorm:"column:destination_chain_id;type:UInt256"`
	// Fee is the fee
	Fee *big.Int `gorm:"column:fee;type:UInt256"`
	// Kappa is theFee keccak256 hash of the transaction
	Kappa sql.NullString `gorm:"column:kappa"`
	// TokenIndexFrom is the index of the from token in the pool
	TokenIndexFrom *big.Int `gorm:"column:token_index_from;type:UInt256"`
	// TokenIndexTo is the index of the to token in the pool
	TokenIndexTo *big.Int `gorm:"column:token_index_to;type:UInt256"`
	// MinDy is the minimum amount of tokens to receive
	MinDy *big.Int `gorm:"column:min_dy;type:UInt256"`
	// Deadline is the deadline of the transaction
	Deadline *big.Int `gorm:"column:deadline;type:UInt256"`
	// SwapSuccess is whether the swap was successful
	SwapSuccess *big.Int `gorm:"column:swap_success;type:UInt256"`
	// SwapTokenIndex is the index of the token in the pool
	SwapTokenIndex *big.Int `gorm:"column:swap_token_index;type:UInt256"`
	// SwapMinAmount is the minimum amount of tokens to receive
	SwapMinAmount *big.Int `gorm:"column:swap_min_amount;type:UInt256"`
	// SwapDeadline is the deadline of the swap transaction
	SwapDeadline *big.Int `gorm:"column:swap_deadline;type:UInt256"`
	// TokenID is the token's ID
	TokenID sql.NullString `gorm:"column:token_id"`
}

// LastLoggedBlockInfo stores the last block number that was logged for a given chain.
type LastLoggedBlockInfo struct {
	// ChainID is the chain id of the chain
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// BlockNumber is the last block number indexed
	BlockNumber uint64 `gorm:"column:block_number"`
}

// FailedLog stores the logs that were unsuccessfully parsed and stored.
type FailedLog struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// PrimaryTopic is the primary topic of the event. Topics[0]
	PrimaryTopic sql.NullString `gorm:"primary_topic"`
	// TopicA is the first topic. Topics[1]
	TopicA sql.NullString `gorm:"topic_a"`
	// TopicB is the second topic. Topics[2]
	TopicB sql.NullString `gorm:"topic_b"`
	// TopicC is the third topic. Topics[3]
	TopicC sql.NullString `gorm:"topic_c"`
	// Data is the data provided by the contract
	Data []byte `gorm:"data"`
	// BlockNumber is the block in which the transaction was included
	BlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// TxIndex is the index of the transaction in the block
	TxIndex uint64 `gorm:"tx_index"`
	// BlockHash is the hash of the block in which the transaction was included
	BlockHash string `gorm:"block_hash"`
	// Index is the index of the log in the block
	BlockIndex uint64 `gorm:"column:block_index;primaryKey;auto_increment:false"`
	// Removed is true if this log was reverted due to a chain re-organization
	Removed bool `gorm:"removed"`
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"confirmed"`
}
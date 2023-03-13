package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "loan"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_loan"
)

var (
	TypeCreateRequest = ModuleName + ":create"
)

var (
	LoanIDKey     = []byte{0x01}
	LoanKeyPrefix = []byte{0x02}
)

func LoanKey(id uint64) []byte {
	return append(LoanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

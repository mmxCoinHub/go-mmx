package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	// MmxBech32Prefix defines the Bech32 prefix of an account's address
	MmxBech32Prefix = "mmx"

	// MmxCoinType Mmx coin in https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	MmxCoinType = 118

	// MmxFullFundraiserPath BIP44Prefix is the parts of the BIP44 HD path that are fixed by what we used during the fundraiser.
	// use the registered cosmos stake token ATOM 118 as coin_type
	// m / purpose' / coin_type' / account' / change / address_index
	MmxFullFundraiserPath = "44'/118'/0'/0/0"

	// MmxBech32PrefixAccAddr defines the Bech32 prefix of an account's address
	MmxBech32PrefixAccAddr = MmxBech32Prefix
	// MmxBech32PrefixAccPub defines the Bech32 prefix of an account's public key
	MmxBech32PrefixAccPub = MmxBech32Prefix + sdk.PrefixPublic
	// MmxBech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	MmxBech32PrefixValAddr = MmxBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// MmxBech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	MmxBech32PrefixValPub = MmxBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// MmxBech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	MmxBech32PrefixConsAddr = MmxBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// MmxBech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	MmxBech32PrefixConsPub = MmxBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

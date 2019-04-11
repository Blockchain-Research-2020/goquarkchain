package types

import (
	"github.com/QuarkChain/goquarkchain/account"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type IHeader interface {
	Hash() common.Hash
	SealHash() common.Hash
	NumberU64() uint64
	GetParentHash() common.Hash
	GetPrevRootBlockHash() common.Hash
	GetCoinbase() account.Address
	GetTime() uint64
	GetCoinbaseAmount()*big.Int
	GetDifficulty() *big.Int
	GetNonce() uint64
	GetExtra() []byte
	GetGasLimit() *big.Int
	GetBranch() account.Branch
	GetMetaHash() common.Hash
	GetBloom() Bloom
	SetCoinbase(account.Address)
	SetExtra([]byte)
	SetDifficulty(*big.Int)
	SetNonce(uint64)
	GetMixDigest() common.Hash
	ValidateHeader() error
}

type IBlock interface {
	ValidateBlock() error
	Hash() common.Hash
	NumberU64() uint64
	IHeader() IHeader
	WithMingResult(nonce uint64, mixDigest common.Hash) IBlock
	Content() []IHashable
	GetMetaData() *MinorBlockMeta
	GetTrackingData() []byte
	GetTransactions() Transactions
	GetSize() common.StorageSize
	CoinbaseAmount() *big.Int
}

type IHashable interface {
	Hash() common.Hash
}

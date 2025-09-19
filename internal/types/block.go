package types

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockBody struct {
	Transactions TransactionSignedList `json:"transactions" gencodec:"required"`
	Ommers       Headers               `json:"ommers"       gencodec:"required"`
	Withdrawals  types.Withdrawals     `json:"withdrawals"`
}

func (bb *BlockBody) GethTransactions() []*types.Transaction {
	return bb.Transactions.GethType()
}

func (bb *BlockBody) GethOmmers() []*types.Header {
	return bb.Ommers.GethType()
}

type Block struct {
	Header *Header    `json:"header" gencodec:"required"`
	Body   *BlockBody `json:"body"   gencodec:"required"`
}

func (b *Block) GethType() *types.Block {
	return types.NewBlockWithHeader(b.Header.GethType()).WithBody(types.Body{
		Transactions: b.Body.GethTransactions(),
		Uncles:       b.Body.GethOmmers(),
		Withdrawals:  b.Body.Withdrawals,
	})
}

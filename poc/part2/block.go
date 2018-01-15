package main

import (
	"time"
)

// Block keeps block headers, 由区块头和交易两部分构成
type Block struct {
	Timestamp     int64  // 当前时间戳, 也就是区块创建的时间
	Data          []byte // 区块实际存储的信息, 比特币中也就是交易
	PrevBlockHash []byte // 前一个块的哈希
	Hash          []byte // 当前块的哈希
	Nonce         int
}

// NewBlock creates and returns Block, 用于生成新块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

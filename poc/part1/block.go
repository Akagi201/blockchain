package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps block headers, 由区块头和交易两部分构成
type Block struct {
	Timestamp     int64  // 当前时间戳, 也就是区块创建的时间
	Data          []byte // 区块实际存储的信息, 比特币中也就是交易
	PrevBlockHash []byte // 前一个块的哈希
	Hash          []byte // 当前块的哈希
}

// SetHash calculates and sets block hash, 设置当前块哈希
// Hash = sha256(PrevBlockHash + Data + Timestamp)
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns Block, 用于生成新块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

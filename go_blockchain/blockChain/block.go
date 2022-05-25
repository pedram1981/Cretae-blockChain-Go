package blockChain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) CreateHash() {
	var info = bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	var hash = sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	var block = &Block{[]byte{}, []byte(data), prevHash}
	block.CreateHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	var prevBlock = chain.Blocks[len(chain.Blocks)-1]
	var new = CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func MyCrypto(data string) *Block {
	return CreateBlock(data, []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{MyCrypto("Hello World,I am AI")}}
}

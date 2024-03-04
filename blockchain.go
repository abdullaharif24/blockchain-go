package blockchain

import (
    "fmt"
)

// Block represents a block in the blockchain
type Block struct {
    Index     int
    Timestamp string
    Data      string
    Hash      string
    PrevHash  string
}

// DisplayAllBlocks displays all blocks in the blockchain
func DisplayAllBlocks(blocks []Block) {
    for _, block := range blocks {
        fmt.Printf("Block %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Printf("Previous Hash: %s\n", block.PrevHash)
        fmt.Println("--------------------")
    }
}

// NewBlock creates a new block in the blockchain
func NewBlock(index int, timestamp, data, hash, prevHash string) Block {
    return Block{Index: index, Timestamp: timestamp, Data: data, Hash: hash, PrevHash: prevHash}
}


// ModifyBlock modifies an existing block in the blockchain
func ModifyBlock(block *Block, data string) {
    block.Data = data
}


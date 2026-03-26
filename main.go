package main

import (
	"fmt"
)

func main() {
	// 1. Create the Genesis Block
	genesis := NewGenesisBlock()
	fmt.Printf("--- Genesis Block ---\n")
	fmt.Printf("Hash: %x\n", genesis.Hash)
	fmt.Printf("Data: %s\n\n", genesis.Data)

	// 2. Create a Second Block (Linking it to Genesis)
	// We pass the hash of the genesis block as the 'Previous Hash'
	secondBlock := NewBlock("Identity: Sweety Raj (Person)", genesis.Hash)

	fmt.Printf("--- Second Block ---\n")
	fmt.Printf("Hash: %x\n", secondBlock.Hash)
	fmt.Printf("Prev Hash: %x\n", secondBlock.PrevBlockHash)
	fmt.Printf("Data: %s\n\n", secondBlock.Data)

	// 3. Test Serialization (What BadgerDB will see)
	serialized := secondBlock.Serialize()
	fmt.Printf("Serialized bytes (Ready for Database): %d bytes\n", len(serialized))

	// 4. Test Deserialization (Turning bytes back into a Block)
	reconstructed := DeserializeBlock(serialized)
	fmt.Printf("Reconstructed Data: %s\n", reconstructed.Data)
}

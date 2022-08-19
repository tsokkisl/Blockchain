package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	nonce    []byte
	hash     []byte
	data     [][]byte
	prevHash []byte
}

func randString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func serializeData(d [][]byte) []byte {
	result := []byte("")
	for _, val := range d {
		result = append(result, val...)
	}
	return result
}

func (b *Block) generateHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.nonce, serializeData(b.data), b.prevHash}, []byte{}))
	b.hash = hash[:]
}

func (bc *Blockchain) printBlockchain() {
	for i, block := range bc.blocks {
		fmt.Printf("\nBlock %d\n-----\nHash: %x\nData: %s\nPrev Hash: %x\nNonce: %s\n", i, block.hash, block.data, block.prevHash, block.nonce)
	}
}

func (bc *Blockchain) addBlock(b *Block) {
	bc.blocks = append(bc.blocks, b)
}

func (bc *Blockchain) generateBlock() *Block {
	var block = Block{[]byte(strconv.Itoa(100000000 + rand.Intn(100000000))), []byte{}, [][]byte{[]byte(randString(20)), []byte(randString(20))}, bc.blocks[len(bc.blocks)-1].hash}
	block.generateHash()
	return &block
}

func initBlockchain() *Blockchain {
	var bc = Blockchain{[]*Block{}}
	var genesisBlock = Block{[]byte(strconv.Itoa(100000000 + rand.Intn(100000000))), []byte{}, [][]byte{[]byte("Genesis"), []byte("Block")}, []byte{}}
	genesisBlock.generateHash()
	bc.addBlock(&genesisBlock)
	return &bc
}

func main() {
	blockChain := initBlockchain()
	i := 0
	for i != 10 {
		block := blockChain.generateBlock()
		blockChain.addBlock(block)
		i++
	}
	blockChain.printBlockchain()
}

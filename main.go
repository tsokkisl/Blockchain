package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type Blockchain struct {
	blocks     []*Block
	difficulty int
}

type Block struct {
	nonce    int
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
	hash := sha256.Sum256(bytes.Join([][]byte{[]byte(strconv.Itoa(b.nonce)), serializeData(b.data), b.prevHash}, []byte{}))
	b.hash = hash[:]
}

func (bc *Blockchain) printBlockchain() {
	for i, block := range bc.blocks {
		fmt.Printf("\nBlock %d\n-----\nHash: %x\nData: %s\nPrev Hash: %x\nNonce: %d\n", i, block.hash, block.data, block.prevHash, block.nonce)
	}
}

func (bc *Blockchain) addBlock(b *Block) {
	proofOfWork := false
	diff := fmt.Sprintf("%0*d", bc.difficulty, 0)
	hashPrefix := string(fmt.Sprintf("%x", b.hash))[:bc.difficulty]
	for proofOfWork != true {
		if hashPrefix == diff {
			proofOfWork = true
		} else {
			b.nonce += 1
			hash := sha256.Sum256(bytes.Join([][]byte{[]byte(strconv.Itoa(b.nonce)), serializeData(b.data), b.prevHash}, []byte{}))
			b.hash = hash[:]
			hashPrefix = string(fmt.Sprintf("%x", b.hash))[:bc.difficulty]
		}
	}
	bc.blocks = append(bc.blocks, b)
}

func (bc *Blockchain) generateBlock() *Block {
	var block = Block{100000000 + rand.Intn(100000000), []byte{}, [][]byte{[]byte(randString(20)), []byte(randString(20))}, bc.blocks[len(bc.blocks)-1].hash}
	block.generateHash()
	return &block
}

func initBlockchain(difficulty int) *Blockchain {
	var bc = Blockchain{[]*Block{}, difficulty}
	var genesisBlock = Block{100000000 + rand.Intn(100000000), []byte{}, [][]byte{[]byte("Genesis"), []byte("Block")}, []byte{}}
	genesisBlock.generateHash()
	bc.addBlock(&genesisBlock)
	return &bc
}

func main() {
	difficulty := 5
	blockChain := initBlockchain(difficulty)
	for i := 0; i != 5; {
		block := blockChain.generateBlock()
		blockChain.addBlock(block)
		i++
	}
	blockChain.printBlockchain()
}

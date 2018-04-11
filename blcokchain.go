package blockchain

import (
	"encoding/json"

	"crypto/sha256"
	"fmt"
	"strings"
	"time"
	"encoding/hex"
)

const Difficulty = 2

type Block struct {
	PreviousHash string
	Hash         string
	Data         string
	TimeStamp    int64
	Nonce        int64
}

//New 新建一个区块
func NewBlock(data, previousHash string) Block {
	block := Block{
		PreviousHash: previousHash,
		Data:         data,
		TimeStamp:    time.Now().Unix(),
		Nonce:        0,
	}
	block.Hash = block.calculateHash()
	return block
}

//calculateHash 计算sha256
func (b Block) calculateHash() string {
	hash := sha256.New()
	hashContent := fmt.Sprintf("%s%s%d%d", b.PreviousHash, b.Data, b.Nonce, b.TimeStamp)
	hash.Write([]byte(hashContent))
	md := hash.Sum(nil)
	return string(hex.EncodeToString(md))
}

// MineBlock 设置计算HASH的难度
func (b Block) MineBlock(difficulty int) {
	target := strings.Repeat("0", difficulty)
	for string([]rune(b.Hash)[:difficulty]) != target {
		b.Nonce++
		b.Hash = b.calculateHash()
	}
	fmt.Println("Block mined!!!")
}

type BlockChain struct {
	List []Block
}

func NewBlockChain() BlockChain {
	return BlockChain{
		List: make([]Block, 0),
	}
}

func (bc BlockChain) isChainValid() bool {
	var currentBlcok, previosBlock Block

	for i := 1; i < len(bc.List); i++ {
		currentBlcok = bc.List[i]
		previosBlock = bc.List[i-1]
		if currentBlcok.Hash != currentBlcok.calculateHash() {
			fmt.Println("Current hashs not equal")
			return false
		}

		if previosBlock.Hash != currentBlcok.PreviousHash {
			fmt.Println("Pervious hashs not equal")
			return false
		}
	}
	return true
}

func (bc *BlockChain) Make() {
	bc.List = append(bc.List, NewBlock("Data 1", "0"))
	fmt.Println("Trying to mine block 1")
	bc.List[0].MineBlock(Difficulty)

	bc.List = append(bc.List, NewBlock("Data 2", bc.List[len(bc.List)-1].Hash))
	fmt.Println("Trying to mine block 2")
	bc.List[1].MineBlock(Difficulty)

	bc.List = append(bc.List, NewBlock("Data 3", bc.List[len(bc.List)-1].Hash))
	fmt.Println("Trying to mine block 3")
	bc.List[2].MineBlock(Difficulty)

	fmt.Printf("BlockChain is Valid: %t\n", bc.isChainValid())

	if j, err := json.Marshal(bc); err == nil {
		fmt.Println(string(j))
	}

}

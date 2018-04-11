package blockchain_test

import (
	blockchain "ez-blockchain"
	"testing"
)

func TestBlockChain(t *testing.T) {
	bc := blockchain.NewBlockChain()
	bc.Make()
}

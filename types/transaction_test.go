package types

import (
	"testing"

	"github.com/thisguymartin/blbockchain-grpc-protobuf/crypto"
	"github.com/thisguymartin/blbockchain-grpc-protobuf/util"

	"github.com/thisguymartin/blbockchain-grpc-protobuf/proto"
)

// balance 100 coins
// want to send 5 coints to "AAAA"

func TestNewTransaction(t *testing.T) {
	fromPrivateKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivateKey.Public().Address().Bytes()

	toPriveKey := crypto.GeneratePrivateKey()
	toAdderss := toPriveKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PubKey:       fromPrivateKey.Public().Bytes(),
		PrevOutIndex: 0,
	}

	output := &proto.TxOutput{
		Amount:  5,
		Address: toAdderss,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output, output2},
	}

	sig := SignTransaction(fromPrivateKey, tx)

	input.Signature = sig.Bytes()

}

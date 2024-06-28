package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thisguymartin/blbockchain-grpc-protobuf/crypto"
	"github.com/thisguymartin/blbockchain-grpc-protobuf/util"
)

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}

func TestSignedBlock(t *testing.T) {
	block := util.RandomBlock()
	privKey := crypto.GeneratePrivateKey()
	pubKey := privKey.Public()

	sig := SignBlock(privKey, block)

	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))

}

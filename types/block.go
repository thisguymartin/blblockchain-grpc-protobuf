package types

import (
	"crypto/sha256"

	"github.com/thisguymartin/blbockchain-grpc-protobuf/crypto"
	"github.com/thisguymartin/blbockchain-grpc-protobuf/proto"
	pb "google.golang.org/protobuf/proto"
)

// Hashblock creates a SHA256 of header.
func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {

	return pk.Sign(HashBlock(b))
}

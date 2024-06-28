package types

import (
	"crypto/sha256"

	"github.com/thisguymartin/blbockchain-grpc-protobuf/crypto"
	"github.com/thisguymartin/blbockchain-grpc-protobuf/proto"

	pb "google.golang.org/protobuf/proto"
)

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

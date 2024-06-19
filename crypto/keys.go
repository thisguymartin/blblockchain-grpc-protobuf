package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

var (
	privateKeyLength = 64
	publicKeyLength  = 32
	seedLen          = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

type PublicKey struct {
	key ed25519.PublicKey
}

type Signature struct {
	value []byte
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)

	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, publicKeyLength)
	copy(b, p.key[32:])

	return &PublicKey{
		key: b,
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.key, msg, s.value)
}

func (s *Signature) Bytes() []byte {
	return s.value
}

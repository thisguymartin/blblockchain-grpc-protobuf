package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.Public()
	// fmt.Println("privateKey", privateKey)
	// fmt.Println("publicKey", publicKey)
	assert.Equal(t, len(privateKey.Bytes()), privateKeyLength)
	assert.Equal(t, len(publicKey.Bytes()), publicKeyLength)

}

func TestPrivateSignature(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.Public()
	msg := []byte("testing here")

	privateKeySignature := privateKey.Sign(msg)

	valid := privateKeySignature.Verify(publicKey, msg)
	assert.True(t, valid)

	new_verification := privateKeySignature.Verify(publicKey, []byte("This is some new test here for it to fail"))
	assert.False(t, new_verification)

	// Test invalid key
	invalidPrivateKey := GeneratePrivateKey()
	invalidPublicKey := invalidPrivateKey.Public()
	bad_pub_key := privateKeySignature.Verify(invalidPublicKey, msg)
	assert.False(t, bad_pub_key)
}

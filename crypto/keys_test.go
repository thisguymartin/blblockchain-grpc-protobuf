package crypto

import (
	"fmt"
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

func TestPublicToAddress(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.Public()
	address := publicKey.Address()
	assert.Equal(t, len(address.Bytes()), addressLength)
	fmt.Println(address)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	seed := "57b06d0522c43d39e960e853cddbfd8e7cba115f60aa7b9529b1edcb37062be4"
	addressStr := "0034f5b35248f827ea04189ced965bec1307100e"
	privKey := NewPrivateKeyFromString(seed)

	assert.Equal(t, privateKeyLength, len(privKey.Bytes()))

	address := privKey.Public().Address()

	assert.Equal(t, addressStr, address.String())
}

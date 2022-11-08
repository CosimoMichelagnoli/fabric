package sw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOqsSigner_Sign(t *testing.T) {
	signer := &oqsSigner{}
	privateKeyVerifier := &oqsPrivateKeyVerifier{}
	publicKeyVerifier := &oqsPublicKeyKeyVerifier{}
	//KeyVerifier := &oqsVerifier{}

	// Generate keypair, message digest
	kg := &oqsKeyGenerator{}
	priv, err := kg.KeyGen(nil)
	assert.NoError(t, err)
	pub, err := priv.PublicKey()
	//priv, err := kg.KeyGen(nil) //TODO need a way of passing an algorithm to test
	assert.NoError(t, err)
	digest := []byte("Hello world")

	// Sign and verify signature
	signature, err := signer.Sign(priv, digest, nil)
	assert.NoError(t, err)
	verify, err := publicKeyVerifier.Verify(pub, signature, digest, nil)
	assert.NoError(t, err)
	assert.True(t, verify)
	verify, err = privateKeyVerifier.Verify(priv, signature, digest, nil)
	/*
		verify, err := KeyVerifier.Verify(priv, signature, digest, nil)
		assert.NoError(t, err)
		assert.True(t, verify)
	*/
}

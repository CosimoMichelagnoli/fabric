package sw

import (
	dilithium5 "crypto/pqc/dilithium/dilithium5"
	"crypto/sha256"
	"errors"

	"github.com/hyperledger/fabric/bccsp"
)

const (
	sigName        = "dilithium5"
	PublicKeySize  = 2592
	PrivateKeySize = 4864
)

// dilithiumPrivateKey implements a bccsp.Key interface
type dilithiumPrivateKey struct {
	privKey *dilithium5.PrivateKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *dilithiumPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k *dilithiumPrivateKey) SKI() []byte {
	if k.privKey == nil {
		return nil
	}
	algBytes := []byte(sigName)

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.privKey.Sk, algBytes...))
	return hash.Sum(nil)
}

func (k *dilithiumPrivateKey) Symmetric() bool {
	return false
}

func (k *dilithiumPrivateKey) Private() bool {
	return true
}

func (k *dilithiumPrivateKey) PublicKey() (bccsp.Key, error) {
	return &dilithiumPublicKey{&k.privKey.PublicKey}, nil
}

// dilithiumPublicKey implements a bccsp.Key interface
type dilithiumPublicKey struct {
	pubKey *dilithium5.PublicKey
}

func (k *dilithiumPublicKey) Bytes() ([]byte, error) {
	return k.pubKey.Pk, nil
}

// SKI returns the subject key identifier of this key.
func (k *dilithiumPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}
	algBytes := []byte(sigName)

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.pubKey.Pk, algBytes...))
	return hash.Sum(nil)
} // TODO not used

func (k *dilithiumPublicKey) Symmetric() bool {
	return false
}

func (k *dilithiumPublicKey) Private() bool {
	return false
}

func (k *dilithiumPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

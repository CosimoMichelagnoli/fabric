package sw

import (
	dilithium2 "crypto/pqc/dilithium/dilithium2"
	"crypto/sha256"
	"errors"

	"github.com/hyperledger/fabric/bccsp"
)

const (
	sigName        = "Dilithium2"
	PublicKeySize  = 1312
	PrivateKeySize = 2528
)

// dilithiumPrivateKey implements a bccsp.Key interface
type dilithiumPrivateKey struct {
	privKey *dilithium2.PrivateKey
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
	pubKey *dilithium2.PublicKey
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

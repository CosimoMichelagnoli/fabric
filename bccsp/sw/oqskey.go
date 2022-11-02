package sw

import (
	"crypto/sha256"
	"errors"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/open-quantum-safe/liboqs-go/oqs"
)

// oqsPrivateKey implements a bccsp.Key interface
type oqsPrivateKey struct {
	sig       *oqs.Signature
	publicKey []byte
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *oqsPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k *oqsPrivateKey) SKI() []byte {
	if k.sig.ExportSecretKey() == nil {
		return nil
	}
	algBytes := []byte(k.sig.Details().Name)

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.publicKey, algBytes...))
	return hash.Sum(nil)
}

func (k *oqsPrivateKey) Symmetric() bool {
	return false
}

func (k *oqsPrivateKey) Private() bool {
	return true
}

func (k *oqsPrivateKey) PublicKey() (bccsp.Key, error) {
	return &oqsPublicKey{k.publicKey}, nil
}

// oqsPublicKey implements a bccsp.Key interface
type oqsPublicKey struct {
	pubKey []byte
}

func (k *oqsPublicKey) Bytes() ([]byte, error) {
	return k.pubKey, nil
}

// SKI returns the subject key identifier of this key.
func (k *oqsPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}
	algBytes := []byte(k.pubKey)

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.pubKey, algBytes...))
	return hash.Sum(nil)
} //TODO not used

func (k *oqsPublicKey) Symmetric() bool {
	return false
}

func (k *oqsPublicKey) Private() bool {
	return false
}

func (k *oqsPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

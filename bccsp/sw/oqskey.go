package sw

import (
	"crypto/liboqs-go/oqs"
	"crypto/sha256"
	"errors"

	"github.com/hyperledger/fabric/bccsp"
)

// oqsPrivateKey implements a bccsp.Key interface
type oqsPrivateKey struct {
	privKey *oqs.Signature
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *oqsPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k *oqsPrivateKey) SKI() []byte {
	if k.privKey == nil {
		return nil
	}
	algBytes := []byte(k.privKey.Details().Name)

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.privKey.ExportSecretKey(), algBytes...))
	return hash.Sum(nil)
}

func (k *oqsPrivateKey) Symmetric() bool {
	return false
}

func (k *oqsPrivateKey) Private() bool {
	return true
}

func (k *oqsPrivateKey) PublicKey() (bccsp.Key, error) {
	return &oqsPublicKey{k.privKey}, nil
}

// oqsPublicKey implements a bccsp.Key interface
type oqsPublicKey struct {
	pubKey *oqs.Signature
}

func (k *oqsPublicKey) Bytes() ([]byte, error) {
	return k.pubKey.ExportPublicKey(), nil
}

// SKI returns the subject key identifier of this key.
func (k *oqsPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}
	algBytes := []byte(k.pubKey.ExportPublicKey())

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.pubKey.ExportPublicKey(), algBytes...))
	return hash.Sum(nil)
} // TODO not used

func (k *oqsPublicKey) Symmetric() bool {
	return false
}

func (k *oqsPublicKey) Private() bool {
	return false
}

func (k *oqsPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

/*
type oqsSignatureKey struct {
	sig *oqs.Signature
}

// oqsPrivateKey implements a bccsp.Key interface
type oqsPrivateKey struct {
	oqsSignatureKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *oqsPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k *oqsPrivateKey) SKI() []byte {
	if k.sig == nil {
		return nil
	}
	algBytes := []byte(k.sig.Details().Name)

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.oqsSignatureKey.sig.ExportSecretKey(), algBytes...))
	return hash.Sum(nil)
}

func (k *oqsPrivateKey) Symmetric() bool {
	return false
}

func (k *oqsPrivateKey) Private() bool {
	return true
}

func (k *oqsPrivateKey) PublicKey() (bccsp.Key, error) {
	return &oqsPublicKey{k.oqsSignatureKey}, nil
}

// oqsPublicKey implements a bccsp.Key interface
type oqsPublicKey struct {
	oqsSignatureKey
}

func (k *oqsPublicKey) Bytes() ([]byte, error) {
	return k.oqsSignatureKey.sig.ExportPublicKey(), nil
}

// SKI returns the subject key identifier of this key.
func (k *oqsPublicKey) SKI() []byte {
	if k.oqsSignatureKey.sig == nil {
		return nil
	}
	algBytes := []byte(k.oqsSignatureKey.sig.ExportPublicKey())

	// Hash public key with algorithm
	hash := sha256.New()
	hash.Write(append(k.oqsSignatureKey.sig.ExportPublicKey(), algBytes...))
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
*/

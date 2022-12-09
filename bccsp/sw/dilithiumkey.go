/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	dilithium5 "crypto/pqc/dilithium/dilithium5"
	"crypto/sha256"
	"crypto/x509"
	"errors"
	"fmt"

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
	hash := sha256.New()
	hash.Write(k.privKey.Sk)
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

func (k *dilithiumPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *dilithiumPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}
	hash := sha256.New()
	hash.Write(k.pubKey.Pk)
	return hash.Sum(nil)
}

func (k *dilithiumPublicKey) Symmetric() bool {
	return false
}

func (k *dilithiumPublicKey) Private() bool {
	return false
}

func (k *dilithiumPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

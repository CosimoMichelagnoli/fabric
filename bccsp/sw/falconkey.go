/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	falcon1024 "crypto/pqc/falcon/falcon1024"
	"crypto/sha256"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
)

/*var (
	sigName        = "Falcon-1024"
	PublicKeySize  = 1793
	PrivateKeySize = 2305
)*/

// falconPrivateKey implements a bccsp.Key interface
type falconPrivateKey struct {
	privKey *falcon1024.PrivateKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *falconPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k *falconPrivateKey) SKI() []byte {
	if k.privKey == nil {
		return nil
	}
	hash := sha256.New()
	hash.Write(k.privKey.Sk)
	return hash.Sum(nil)
}

func (k *falconPrivateKey) Symmetric() bool {
	return false
}

func (k *falconPrivateKey) Private() bool {
	return true
}

func (k *falconPrivateKey) PublicKey() (bccsp.Key, error) {
	return &falconPublicKey{&k.privKey.PublicKey}, nil
}

// dilithiumPublicKey implements a bccsp.Key interface
type falconPublicKey struct {
	pubKey *falcon1024.PublicKey
}

func (k *falconPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *falconPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}
	hash := sha256.New()
	hash.Write(k.pubKey.Pk)
	return hash.Sum(nil)
}

func (k *falconPublicKey) Symmetric() bool {
	return false
}

func (k *falconPublicKey) Private() bool {
	return false
}

func (k *falconPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

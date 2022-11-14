/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	"crypto/ecdsa"
	"crypto/liboqs-go/oqs"
	dilithium2 "crypto/pqc/dilithium/dilithium2"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"errors"
	"fmt"
	"reflect"

	"github.com/hyperledger/fabric/bccsp"
)

var (
	oidAltPublicKeyInfo    = asn1.ObjectIdentifier{2, 5, 29, 71}
	oidSignatureDilithium2 = asn1.ObjectIdentifier{1, 2, 1, 0}
)

type aes256ImportKeyOptsKeyImporter struct{}

func (*aes256ImportKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	aesRaw, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if aesRaw == nil {
		return nil, errors.New("Invalid raw material. It must not be nil.")
	}

	if len(aesRaw) != 32 {
		return nil, fmt.Errorf("Invalid Key Length [%d]. Must be 32 bytes", len(aesRaw))
	}

	return &aesPrivateKey{aesRaw, false}, nil
}

type hmacImportKeyOptsKeyImporter struct{}

func (*hmacImportKeyOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	aesRaw, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if len(aesRaw) == 0 {
		return nil, errors.New("Invalid raw material. It must not be nil.")
	}

	return &aesPrivateKey{aesRaw, false}, nil
}

type ecdsaPKIXPublicKeyImportOptsKeyImporter struct{}

func (*ecdsaPKIXPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if len(der) == 0 {
		return nil, errors.New("Invalid raw. It must not be nil.")
	}

	lowLevelKey, err := derToPublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("Failed converting PKIX to ECDSA public key [%s]", err)
	}

	ecdsaPK, ok := lowLevelKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Failed casting to ECDSA public key. Invalid raw material.")
	}

	return &ecdsaPublicKey{ecdsaPK}, nil
}

type ecdsaPrivateKeyImportOptsKeyImporter struct{}

func (*ecdsaPrivateKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("[ECDSADERPrivateKeyImportOpts] Invalid raw material. Expected byte array.")
	}

	if len(der) == 0 {
		return nil, errors.New("[ECDSADERPrivateKeyImportOpts] Invalid raw. It must not be nil.")
	}

	lowLevelKey, err := derToPrivateKey(der)
	if err != nil {
		return nil, fmt.Errorf("Failed converting PKIX to ECDSA public key [%s]", err)
	}

	ecdsaSK, ok := lowLevelKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("Failed casting to ECDSA private key. Invalid raw material.")
	}

	return &ecdsaPrivateKey{ecdsaSK}, nil
}

type ecdsaGoPublicKeyImportOptsKeyImporter struct{}

func (*ecdsaGoPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	lowLevelKey, ok := raw.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *ecdsa.PublicKey.")
	}

	return &ecdsaPublicKey{lowLevelKey}, nil
}

type x509PublicKeyImportOptsKeyImporter struct {
	bccsp *CSP
}

func (ki *x509PublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	x509Cert, ok := raw.(*x509.Certificate)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *x509.Certificate.")
	}

	pk := x509Cert.PublicKey

	switch pk := pk.(type) {
	case *ecdsa.PublicKey:
		return ki.bccsp.KeyImporters[reflect.TypeOf(&bccsp.ECDSAGoPublicKeyImportOpts{})].KeyImport(
			pk,
			&bccsp.ECDSAGoPublicKeyImportOpts{Temporary: opts.Ephemeral()})
	case *rsa.PublicKey:
		// This path only exists to support environments that use RSA certificate
		// authorities to issue ECDSA certificates.
		return &rsaPublicKey{pubKey: pk}, nil
	default:
		return nil, errors.New("Certificate's public key type not recognized. Supported keys: [ECDSA, RSA]")
	}
}

type x509AltPublicKeyImportOptsKeyImporter struct {
	bccsp *CSP
}

func (ki *x509AltPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	x509Cert, ok := raw.(*x509.Certificate)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *x509.Certificate.")
	}

	// TODO: check if it's better Extensions or Extraextension. CheckId to be defined
	for _, ext := range x509Cert.ExtraExtensions {
		if reflect.DeepEqual(ext.Id, oidAltPublicKeyInfo) {
			return ki.bccsp.KeyImporters[reflect.TypeOf(&bccsp.OQSGoPublicKeyImportOpts{})].KeyImport(
				ext.Value,
				&bccsp.OQSGoPublicKeyImportOpts{Temporary: opts.Ephemeral()})
		}

		if reflect.DeepEqual(ext.Id, oidSignatureDilithium2) {
			return ki.bccsp.KeyImporters[reflect.TypeOf(&bccsp.DILITHIUMGoPublicKeyImportOpts{})].KeyImport(
				ext.Value,
				&bccsp.DILITHIUMGoPublicKeyImportOpts{Temporary: opts.Ephemeral()})
		}
	}

	opts.(*bccsp.X509AltPublicKeyImportOpts).Temporary = true
	return nil, nil
}

type oqsGoPublicKeyImportOptsKeyImporter struct{}

func (*oqsGoPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	lowLevelKey, ok := raw.(*oqs.Signature)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *oqs.PublicKey.")
	}

	return &oqsPublicKey{lowLevelKey}, nil
}

type dilithiumGoPublicKeyImportOptsKeyImporter struct{}

func (*dilithiumGoPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	lowLevelKey, ok := raw.(*dilithium2.PublicKey)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *oqs.PublicKey.")
	}

	return &dilithiumPublicKey{lowLevelKey}, nil
}

/*
type oqsPKIXPublicKeyImportOptsKeyImporter struct{}

func (*oqsPKIXPublicKeyImportOptsKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (bccsp.Key, error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected byte array.")
	}

	if len(der) == 0 {
		return nil, errors.New("Invalid raw. It must not be nil.")
	}

	lowLevelKey, err := oqs.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("Failed converting PKIX to OQS public key [%s]", err)
	}

	oqsPK, ok := lowLevelKey.(*oqs.PublicKey)
	if !ok {
		return nil, errors.New("Failed casting to OQS public key. Invalid raw material.")
	}

	return &oqsPublicKey{oqsPK}, nil
}*/

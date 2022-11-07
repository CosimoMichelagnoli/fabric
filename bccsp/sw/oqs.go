package sw

import (
	"github.com/hyperledger/fabric/bccsp"
)

type oqsSigner struct{}

func (s *oqsSigner) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return k.(*oqsSignatureKey).sig.Sign(digest)
}

type oqsVerifier struct{}

func (v *oqsVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return k.(*oqsSignatureKey).sig.Verify(digest, signature, k.(*oqsSignatureKey).pubKey)
}

/*
type oqsPrivateKeyVerifier struct{}

func (v *oqsPrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return k.(*oqsPrivateKey).sig.Verify(digest, signature, k.(*oqsPrivateKey).sig.ExportSecretKey())
}

type oqsPublicKeyKeyVerifier struct{}

func (v *oqsPublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return k.(*oqsPrivateKey).sig.Verify(digest, signature, k.(*oqsPrivateKey).publicKey)
}
*/

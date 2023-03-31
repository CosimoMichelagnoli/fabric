/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sw

import (
	"crypto/rand"

	"github.com/hyperledger/fabric/bccsp"
)

type falconSigner struct{}

func (s *falconSigner) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return k.(*falconPrivateKey).privKey.Sign(rand.Reader, digest, opts)
}

type falconPrivateKeyVerifier struct{}

func (v *falconPrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return false, nil
}

type falconPublicKeyKeyVerifier struct{}

func (v *falconPublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return k.(*falconPublicKey).pubKey.Verify(digest, signature), nil
}

/*
Copyright IBM Corp. 2017 All Rights Reserved.

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
	"crypto/ecdsa"
	"crypto/elliptic"
	dilithium5 "crypto/pqc/dilithium/dilithium5"
	falcon1024 "crypto/pqc/falcon/falcon1024"
	"crypto/rand"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
)

type ecdsaKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *ecdsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	privKey, err := ecdsa.GenerateKey(kg.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", kg.curve, err)
	}

	return &ecdsaPrivateKey{privKey}, nil
}

type aesKeyGenerator struct {
	length int
}

func (kg *aesKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	lowLevelKey, err := GetRandomBytes(int(kg.length))
	if err != nil {
		return nil, fmt.Errorf("Failed generating AES %d key [%s]", kg.length, err)
	}

	return &aesPrivateKey{lowLevelKey, false}, nil
}

type dilithiumKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *dilithiumKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	privKey, err := dilithium5.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("Failed generating DILITHIUM key: [%s]", err)
	}

	return &dilithiumPrivateKey{privKey}, nil
}

type falconKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *falconKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	privKey, err := falcon1024.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("Failed generating FALCON key: [%s]", err)
	}

	return &falconPrivateKey{privKey}, nil
}

// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keys

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/google/trillian/crypto/keyspb"
	"github.com/letsencrypt/pkcs11key"
)

// The size of an RSA key generated by this library, in bits, if not overridden.
const defaultRsaKeySizeInBits = 2048

// MinRsaKeySizeInBits is the smallest RSA key that this package will generate.
const MinRsaKeySizeInBits = 2048

// SignerFactory gets and creates cryptographic signers.
// This may be done by loading a key from a file, interfacing with a HSM, or
// sending requests to a remote key management service, to give a few examples.
type SignerFactory interface {
	// NewSigner uses the information in the provided protobuf message to obtain and return a crypto.Signer.
	NewSigner(context.Context, proto.Message) (crypto.Signer, error)

	// Generate creates a new private key based on a key specification.
	// It returns a proto that describes how to access that key.
	// This proto can be passed to NewSigner() to get a crypto.Signer.
	Generate(context.Context, *keyspb.Specification) (proto.Message, error)
}

// NewFromPrivatePEMFile reads a PEM-encoded private key from a file.
// The key must be protected by a password.
func NewFromPrivatePEMFile(keyFile, keyPassword string) (crypto.Signer, error) {
	if keyPassword == "" {
		return nil, fmt.Errorf("empty password for PEM key file %q", keyFile)
	}
	pemData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key from file %q: %v", keyFile, err)
	}

	k, err := NewFromPrivatePEM(string(pemData), keyPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key from file %q: %v", keyFile, err)
	}

	return k, nil
}

// NewFromPrivatePEM reads a PEM-encoded private key from a string.
// The key may be protected by a password.
func NewFromPrivatePEM(pemEncodedKey, password string) (crypto.Signer, error) {
	block, rest := pem.Decode([]byte(pemEncodedKey))
	if len(rest) > 0 {
		return nil, errors.New("extra data found after PEM decoding")
	}

	der := block.Bytes
	if password != "" {
		pwdDer, err := x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			return nil, err
		}
		der = pwdDer
	}

	return NewFromPrivateDER(der)
}

// NewFromPrivateDER reads a DER-encoded private key.
func NewFromPrivateDER(der []byte) (crypto.Signer, error) {
	key1, err1 := x509.ParsePKCS1PrivateKey(der)
	if err1 == nil {
		return key1, nil
	}

	key2, err2 := x509.ParsePKCS8PrivateKey(der)
	if err2 == nil {
		switch key2 := key2.(type) {
		case *ecdsa.PrivateKey:
			return key2, nil
		case *rsa.PrivateKey:
			return key2, nil
		}
		return nil, fmt.Errorf("unsupported private key type: %T", key2)
	}

	key3, err3 := x509.ParseECPrivateKey(der)
	if err3 == nil {
		return key3, nil
	}

	return nil, fmt.Errorf("could not parse DER private key as PKCS1 (%v), PKCS8 (%v), or SEC1 (%v)", err1, err2, err3)
}

// NewFromSpec generates a new private key based on a key specification.
// If an RSA key is specified, the key size must be at least MinRsaKeySizeInBits.
func NewFromSpec(spec *keyspb.Specification) (crypto.Signer, error) {
	switch params := spec.GetParams().(type) {
	case *keyspb.Specification_EcdsaParams:
		curve := curveFromParams(params.EcdsaParams)
		if curve == nil {
			return nil, fmt.Errorf("unsupported ECDSA curve: %s", params.EcdsaParams.GetCurve())
		}

		return ecdsa.GenerateKey(curve, rand.Reader)
	case *keyspb.Specification_RsaParams:
		bits := int(params.RsaParams.GetBits())
		if bits == 0 {
			bits = defaultRsaKeySizeInBits
		}
		if bits < MinRsaKeySizeInBits {
			return nil, fmt.Errorf("minimum RSA key size is %v bits, got %v bits", MinRsaKeySizeInBits, bits)
		}

		return rsa.GenerateKey(rand.Reader, bits)
	default:
		return nil, fmt.Errorf("unsupported keygen params type: %T", params)
	}
}

// curveFromParams returns the curve specified by the given parameters.
// Returns nil if the curve is not supported.
func curveFromParams(params *keyspb.Specification_ECDSA) elliptic.Curve {
	switch params.GetCurve() {
	case keyspb.Specification_ECDSA_DEFAULT_CURVE:
		return elliptic.P256()
	case keyspb.Specification_ECDSA_P256:
		return elliptic.P256()
	case keyspb.Specification_ECDSA_P384:
		return elliptic.P384()
	case keyspb.Specification_ECDSA_P521:
		return elliptic.P521()
	}
	return nil
}

// MarshalPrivateKey serializes an RSA or ECDSA private key as DER.
func MarshalPrivateKey(key crypto.Signer) ([]byte, error) {
	switch key := key.(type) {
	case *ecdsa.PrivateKey:
		return x509.MarshalECPrivateKey(key)
	case *rsa.PrivateKey:
		return x509.MarshalPKCS1PrivateKey(key), nil
	}

	return nil, fmt.Errorf("unsupported key type: %T", key)
}

// NewFromPKCS11Config returns a crypto.Signer that uses a PKCS#11 interface.
func NewFromPKCS11Config(modulePath string, config *keyspb.PKCS11Config) (crypto.Signer, error) {
	if modulePath == "" {
		return nil, errors.New("No PKCS#11 module path set, cannot create signer")
	}
	pubKey, err := NewFromPublicPEM(config.GetPublicKey())
	if err != nil {
		return nil, fmt.Errorf("Failed to load public key from %q: %s", config.GetPublicKey(), err)
	}
	return pkcs11key.New(modulePath, config.GetTokenLabel(), config.GetPin(), pubKey)
}

/*
 * Copyright (c) 2025-present the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rsas

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"

	perrors "github.com/pkg/errors"
)

type RsaKey string

var (
	PRIVATE RsaKey = "PRIVATE KEY"
	PUBLIC  RsaKey = "PUBLIC KEY"
)

var (
	ErrPrivateKeyDecode = perrors.New("decode the private key to block failed")
	ErrPublicKeyDecode  = perrors.New("decode the public key to block failed")

	ErrPublicKeyParse    = perrors.New("parse the public key failed")
	ErrInvalidPublicKey  = perrors.New("invalid public key")
	ErrInvalidPrivateKey = perrors.New("invalid private key")

	ErrInvalidPublicKeyType  = perrors.New("invalid public key type")
	ErrInvalidPrivateKeyType = perrors.New("invalid private key type")
)

// ----------------------------------------------------------------

func LoadPublicKeyPem(publicKeyBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, ErrPublicKeyDecode
	}
	if !strings.Contains(block.Type, string(PUBLIC)) {
		return nil, ErrInvalidPublicKeyType
	}

	return LoadPublicKey(block.Bytes)
}

func LoadPrivateKeyPem(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, ErrPrivateKeyDecode
	}
	if !strings.Contains(block.Type, string(PRIVATE)) {
		return nil, ErrInvalidPrivateKeyType
	}

	return LoadPrivateKey(block.Bytes)
}

// ----------------------------------------------------------------

func LoadPublicKey(blockByte []byte) (*rsa.PublicKey, error) {
	pubKey, err := x509.ParsePKIXPublicKey(blockByte)
	if err != nil {
		return nil, ErrPublicKeyParse
	}
	publicKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, ErrInvalidPublicKey
	}

	return publicKey, nil
}

func LoadPrivateKey(blockByte []byte) (*rsa.PrivateKey, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(blockByte)
	if err != nil {
		pkcs8PrivateKey, err := x509.ParsePKCS8PrivateKey(blockByte)
		if err != nil {
			return nil, err
		}
		privateKey, ok := pkcs8PrivateKey.(*rsa.PrivateKey)
		if !ok {
			return nil, ErrInvalidPrivateKey
		}
		return privateKey, nil
	}

	return privateKey, err
}

// ----------------------------------------------------------------

func split(data []byte, blockSize int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(data)/blockSize+1)
	for len(data) >= blockSize {
		chunk, data = data[:blockSize], data[blockSize:]
		chunks = append(chunks, chunk)
	}
	if len(data) > 0 {
		chunks = append(chunks, data)
	}

	return chunks
}

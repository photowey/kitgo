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
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type Sha256 struct{}

type ISha256 interface {
	Signer
	OAEP
	PKCS1
}

func (algo Sha256) Sign(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return sign(privateKey, crypto.SHA256, hashed)
}

func (algo Sha256) Verify(publicKey *rsa.PublicKey, data, sign []byte) (bool, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return verify(publicKey, crypto.SHA256, hashed, sign)
}

// EncryptOAEP RSA 公钥加密 OAEP
func (algo Sha256) EncryptOAEP(publicKey *rsa.PublicKey, data []byte) (string, error) {
	if publicKey == nil {
		return EmptyString, ErrPrivateKeyNULL
	}
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
	if err != nil {
		return EmptyString, fmt.Errorf("rsas: encrypt data err: %s", err.Error())
	}
	encrypted := encryptBase64(encryptedBytes)

	return encrypted, nil
}

func (algo Sha256) EncryptBlockOAEP(publicKey *rsa.PublicKey, data []byte) (string, error) {
	if publicKey == nil {
		return EmptyString, ErrPrivateKeyNULL
	}

	encryptedBytes, err := blockEncrypt(publicKey, sha256.New(), data, OAEPMode)
	if err != nil {
		return EmptyString, err
	}

	encrypted := encryptBase64(encryptedBytes)

	return encrypted, nil
}

// DecryptOAEP RSA 私钥块解密 OAEP
func (algo Sha256) DecryptOAEP(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error) {
	if privateKey == nil {
		return EmptyString, ErrPublicKeyNULL
	}

	encryptedBytes, err := decryptBase64(encryptedBase64)
	if err != nil {
		return EmptyString, err
	}

	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedBytes, nil)
	if err != nil {
		return EmptyString, fmt.Errorf("rsas: decrypt encrypted err:%s", err)
	}

	return string(decryptedBytes), err
}

func (algo Sha256) DecryptBlockOAEP(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error) {
	if privateKey == nil {
		return EmptyString, ErrPublicKeyNULL
	}

	encryptedBytes, err := decryptBase64(encryptedBase64)
	if err != nil {
		return EmptyString, err
	}

	// 分块处理
	decryptedBytes, err := blockDecrypt(privateKey, sha256.New(), encryptedBytes, OAEPMode)
	if err != nil {
		return EmptyString, err
	}

	return string(decryptedBytes.Bytes()), nil
}

// EncryptBlockPKCS1 RSA 私钥块解密 PKCS1
func (algo Sha256) EncryptBlockPKCS1(publicKey *rsa.PublicKey, data []byte) (string, error) {
	if publicKey == nil {
		return EmptyString, ErrPrivateKeyNULL
	}

	encryptedBytes, err := blockEncrypt(publicKey, sha256.New(), data, PKCS1Mode)
	if err != nil {
		return EmptyString, err
	}

	encrypted := encryptBase64(encryptedBytes)

	return encrypted, nil
}

// DecryptBlockPKCS1 RSA 私钥块解密 PKCS1
func (algo Sha256) DecryptBlockPKCS1(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error) {
	if privateKey == nil {
		return EmptyString, ErrPublicKeyNULL
	}

	encryptedBytes, err := decryptBase64(encryptedBase64)
	if err != nil {
		return EmptyString, err
	}
	// 分块处理
	decryptedBytes, err := blockDecrypt(privateKey, sha256.New(), encryptedBytes, PKCS1Mode)
	if err != nil {
		return EmptyString, err
	}

	return string(decryptedBytes.Bytes()), err
}

func Sha256Algo() ISha256 {
	return _sha256
}

func sha256Algo() ISha256 {
	return &Sha256{}
}

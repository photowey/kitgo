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
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type PKCS1 interface {
	EncryptBlockPKCS1(publicKey *rsa.PublicKey, data []byte) (string, error)
	DecryptBlockPKCS1(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error)
}

func EncryptPKCS1(publicKey *rsa.PublicKey, data []byte) (string, error) {
	if publicKey == nil {
		return "", ErrPrivateKeyNULL
	}
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return "", fmt.Errorf("rsas: encrypt data err: %s", err.Error())
	}
	encrypted := encryptBase64(encryptedBytes)

	return encrypted, nil
}

func DecryptPKCS1(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error) {
	if privateKey == nil {
		return "", ErrPublicKeyNULL
	}

	encryptedBytes, err := decryptBase64(encryptedBase64)
	if err != nil {
		return "", err
	}

	encryptedByte, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
	if err != nil {
		return "", fmt.Errorf("rsas: decrypt encrypted err:%s", err)
	}

	return string(encryptedByte), nil
}

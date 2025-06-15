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
	"crypto/rsa"
	"strings"
)

type Signer interface {
	Sign(privateKey *rsa.PrivateKey, data []byte) (string, error)
	Verify(publicKey *rsa.PublicKey, data, sign []byte) (bool, error)
}

func sign(privateKey *rsa.PrivateKey, hash crypto.Hash, hashed []byte) (string, error) {
	got, err := rsa.SignPKCS1v15(nil, privateKey, hash, hashed)
	if err != nil {
		return "", err
	}
	signed := strings.TrimSpace(encryptBase64(got))

	return signed, nil
}

func verify(publicKey *rsa.PublicKey, hash crypto.Hash, hashed []byte, sign []byte) (bool, error) {
	err := rsa.VerifyPKCS1v15(publicKey, hash, hashed, sign)
	if err != nil {
		return false, err
	}

	return true, nil
}

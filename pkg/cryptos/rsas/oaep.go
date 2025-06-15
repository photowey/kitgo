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
)

type OAEP interface {
	EncryptOAEP(publicKey *rsa.PublicKey, data []byte) (string, error)
	EncryptBlockOAEP(publicKey *rsa.PublicKey, data []byte) (string, error)
	DecryptOAEP(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error)
	DecryptBlockOAEP(privateKey *rsa.PrivateKey, encryptedBase64 string) (string, error)
}

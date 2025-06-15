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
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	base64s "github.com/photowey/kitgo/pkg/base64s"
	"hash"
)

type CodecMode uint

const (
	EmptyString = ""
)

var (
	ErrPrivateKeyNULL = errors.New("the *rsa.PublicKey can't be nil")
	ErrPublicKeyNULL  = errors.New("the *rsa.PrivateKey can't be nil")
)

var (
	PKCS1Mode CodecMode = 1 >> 1
	OAEPMode  CodecMode = 1 >> 2
)

var (
	_sha1   ISha1
	_sha256 ISha256
)

func init() {
	_sha1 = sha1Algo()
	_sha256 = sha256Algo()
}

// ----------------------------------------------------------------

func blockEncrypt(publicKey *rsa.PublicKey, hash hash.Hash, data []byte, codecMode CodecMode) ([]byte, error) {
	blockSize := publicKey.N.BitLen()/8 - 11
	// OAEP 42
	// blockSize = publicKey.N.BitLen()/8 - 42

	if OAEPMode == codecMode {
		blockSize = publicKey.Size() - 2*hash.Size() - 2
	}

	blocks := split(data, blockSize)
	encryptedBytes := bytes.NewBufferString("")
	for _, block := range blocks {
		var encryptedByte []byte
		var err error
		if OAEPMode == codecMode {
			encryptedByte, err = rsa.EncryptOAEP(hash, rand.Reader, publicKey, block, nil)
		} else {
			encryptedByte, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, block)
		}
		if err != nil {
			return nil, err
		}

		encryptedBytes.Write(encryptedByte)
	}

	return encryptedBytes.Bytes(), nil
}

func blockDecrypt(
	privateKey *rsa.PrivateKey,
	hash hash.Hash,
	encryptedBytes []byte,
	codecMode CodecMode) (*bytes.Buffer, error) {
	blockSize := privateKey.PublicKey.N.BitLen() / 8
	if OAEPMode == codecMode {
		blockSize = privateKey.PublicKey.Size()
	}

	blocks := split(encryptedBytes, blockSize)
	decryptedBytes := bytes.NewBufferString("")
	for _, block := range blocks {

		var decryptedByte []byte
		var err error
		if OAEPMode == codecMode {
			decryptedByte, err = rsa.DecryptOAEP(hash, rand.Reader, privateKey, block, nil)
		} else {
			decryptedByte, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, block)
		}
		if err != nil {
			return nil, fmt.Errorf("rsas:decrypt data err:%s", err)
		}

		decryptedBytes.Write(decryptedByte)
	}

	return decryptedBytes, nil
}

// ----------------------------------------------------------------

func encryptBase64(encryptedBytes []byte) string {
	return base64s.EncodeByteToString(encryptedBytes)
}

func decryptBase64(encryptedBase64 string) ([]byte, error) {
	return base64s.DecodeToByte(encryptedBase64)
}

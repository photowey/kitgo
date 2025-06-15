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

package charsets

import (
	"bytes"
	"io"
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func ToGBK(src string) (string, error) {
	buf := []byte(src)
	gbk, err := ByteToGBK(buf)

	return string(gbk), err
}

func ToUTF8(src string) (string, error) {
	buf := []byte(src)
	utf8z, err := ByteToUTF8(buf)

	return string(utf8z), err
}

func ByteToUTF8(buf []byte) ([]byte, error) {
	if IsUtf8(buf) {
		return buf, nil
	}
	decodeBuf, err := decode(buf, simplifiedchinese.GBK)
	if err != nil {
		return buf, err
	}
	if IsUtf8(decodeBuf) {
		return decodeBuf, nil
	}
	decodeBuf, err = decode(buf, simplifiedchinese.GB18030)
	if err != nil {
		return buf, err
	}

	return decodeBuf, nil
}

func ByteToGBK(buf []byte) ([]byte, error) {
	if IsUtf8(buf) {
		encodeBuf, err := encode(buf, simplifiedchinese.GBK)
		if err != nil {
			return buf, err
		}
		if IsNotUtf8(encodeBuf) {
			return encodeBuf, nil
		}
		encodeBuf, err = encode(buf, simplifiedchinese.GB18030)
		if err != nil {
			return buf, err
		}

		return encodeBuf, nil
	}

	return buf, nil
}

// ----------------------------------------------------------------

func IsUtf8(buf []byte) bool {
	return utf8.Valid(buf)
}

func IsNotUtf8(buf []byte) bool {
	return !IsUtf8(buf)
}

// ----------------------------------------------------------------

func encode(buf []byte, encoding encoding.Encoding) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(buf), encoding.NewEncoder())
	dataBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

func decode(buf []byte, encoding encoding.Encoding) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(buf), encoding.NewDecoder())
	dataBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

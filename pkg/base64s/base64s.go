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

package base64s

import (
	"encoding/base64"
	"github.com/photowey/kitgo/pkg/constants"
)

func EncodeToString(source string) string {
	return EncodeByteToString([]byte(source))
}

func EncodeByteToString(source []byte) string {
	base64String := base64.StdEncoding.EncodeToString(source)

	return base64String
}

func DecodeToString(base64String string) (string, error) {
	bytes, err := DecodeToByte(base64String)
	if err != nil {
		return constants.EmptyString, err
	}

	return string(bytes), nil
}

func DecodeToByte(base64String string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

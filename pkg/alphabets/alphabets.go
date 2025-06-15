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

package alphabets

import (
	"github.com/photowey/kitgo/pkg/constants"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Snake2Pascal(snakeCase string) string {
	snakeCase = strings.Replace(snakeCase, "_", " ", -1)
	snakeCase = cases.Title(language.Dutch).String(snakeCase)
	return strings.Replace(snakeCase, " ", "", -1)
}

func Snake2Camel(snakeCase string) string {
	return Camel(Snake2Pascal(snakeCase))
}

// ----------------------------------------------------------------

func Camel(word string) string {
	if word == constants.EmptyString {
		return constants.EmptyString
	}

	return strings.ToLower(word[:1]) + word[1:]
}

func Pascal(word string) string {
	if word == constants.EmptyString {
		return constants.EmptyString
	}

	return strings.ToUpper(strings.ToLower(word[:1])) + word[1:]
}

func Snake(word string) string {
	if word == constants.EmptyString {
		return word
	}
	srcLen := len(word)
	result := make([]byte, 0, srcLen*2)
	caseSymbol := false
	for i := 0; i < srcLen; i++ {
		char := word[i]
		if i > 0 && char >= 'A' && char <= 'Z' && caseSymbol { // _xxx || yyy__zzz
			result = append(result, '_')
		}
		caseSymbol = char != '_'

		result = append(result, char)
	}

	snakeCase := strings.ToLower(string(result))

	return snakeCase
}

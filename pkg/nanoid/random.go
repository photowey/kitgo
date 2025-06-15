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

package nanoid

import (
	"math/rand"
	"time"
)

const (
	AlphabetLower          = "abcdefghijklmnopqrstuvwxyz"
	AlphabetUpper          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetNumber         = "0123456789"
	AlphabetLowerAndNumber = AlphabetLower + AlphabetNumber
	Alphabet               = AlphabetLower + AlphabetNumber + AlphabetUpper
)

func Random(length int) string {
	alphabetLength := len(Alphabet)
	haystacks := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(alphabetLength - 1)
		haystacks[i] = Alphabet[idx]
	}

	return string(haystacks)
}

func RandomAlphabet(alphabet string, length int) string {
	alphabetLength := len(alphabet)
	haystack := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(alphabetLength - 1)
		haystack[i] = alphabet[idx]
	}

	return string(haystack)
}

func RandomNumber(length int) string {
	alphabetLength := len(AlphabetNumber)
	haystacks := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(alphabetLength - 1)
		haystacks[i] = AlphabetNumber[idx]
	}

	return string(haystacks)
}

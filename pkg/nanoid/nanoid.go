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
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	single      = 1
	DefaultSize = 21
)

func New(sizes ...int) (string, error) {
	size := determineSize(sizes...)
	id, err := gonanoid.New(size)

	return id, err
}

func MustNew(sizes ...int) string {
	id, err := New(sizes...)
	if err != nil {
		size := determineSize(sizes...)
		return Random(size)
	}

	return id
}

func Generate(alphabet string, sizes ...int) (string, error) {
	size := determineSize(sizes...)
	id, err := gonanoid.Generate(alphabet, size)

	return id, err
}

func MustGenerate(alphabet string, sizes ...int) string {
	id, err := Generate(alphabet, sizes...)
	if err != nil {
		size := determineSize(sizes...)
		return RandomAlphabet(alphabet, size)
	}

	return id
}

func determineSize(sizes ...int) int {
	size := DefaultSize
	switch len(sizes) {
	case single:
		size = sizes[0]
	}

	return size
}

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

package collections

func NewInt64Slice(cap ...int) Int64Slice {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(Int64Slice, cap[0])
		}
	}
	return make(Int64Slice, 0)
}

func ToInt64Slice(haystack ...int64) Int64Slice {
	array := make(Int64Slice, len(haystack))
	for i, v := range haystack {
		array[i] = v
	}

	return array
}

func ArrayInt64Contains(haystack []int64, needle int64) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

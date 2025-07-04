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

func NewStringSlice(cap ...int) StringSlice {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(StringSlice, cap[0])
		}
	}

	return make(StringSlice, 0)
}

func ToStringSlice(haystack ...string) StringSlice {
	array := make(StringSlice, len(haystack))
	for i, v := range haystack {
		array[i] = v
	}

	return array
}

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

package strings

import (
	"github.com/photowey/kitgo/pkg/constants"
	"strings"
)

func SliceContains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}

	return false
}

func SliceNotContains(haystack []string, needle string) bool {
	return !SliceContains(haystack, needle)
}

// ----------------------------------------------------------------

func IsEmptyStringSlice(haystack []string) bool {
	return len(haystack) == 0
}

func IsNotEmptyStringSlice(haystack []string) bool {
	return !IsEmptyStringSlice(haystack)
}

// ----------------------------------------------------------------

func Implode(haystack []string, separator string) string {
	if len(haystack) == 0 {
		return constants.EmptyString
	}

	var buf strings.Builder
	for _, str := range haystack {
		if constants.EmptyString == str {
			continue
		}
		buf.WriteString(str)
		buf.WriteString(separator)
	}

	return strings.TrimRight(buf.String(), separator)
}

// ----------------------------------------------------------------

func Explode(data string, separator string) []string {
	if constants.EmptyString == data {
		return MakeStringSlice(0)
	}

	return strings.Split(data, separator)
}

// ----------------------------------------------------------------

func StringSliceIntersect(haystack ...[]string) []string {
	var inter []string
	mp := make(map[string]int)
	length := len(haystack)

	if length == 0 {
		return make([]string, 0)
	}
	if length == 1 {
		for _, needle := range haystack[0] {
			if _, ok := mp[needle]; !ok {
				mp[needle] = 1
				inter = append(inter, needle)
			}
		}
		return inter
	}

	for _, needle := range haystack[0] {
		if _, ok := mp[needle]; !ok {
			mp[needle] = 1
		}
	}

	for _, haystack := range haystack[1 : length-1] {
		for _, needle := range haystack {
			if _, ok := mp[needle]; ok {
				mp[needle]++
			}
		}
	}

	for _, needle := range haystack[length-1] {
		if _, ok := mp[needle]; ok {
			if mp[needle] == length-1 {
				inter = append(inter, needle)
			}
		}
	}

	return inter
}

// ----------------------------------------------------------------

func Clone(haystack []string) []string {
	dst := make([]string, len(haystack))
	copy(dst, haystack)

	return dst
}

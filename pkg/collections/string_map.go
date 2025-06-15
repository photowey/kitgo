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

func NewStringMap(cap ...int) StringMap {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(StringMap, cap[0])
		}
	}

	return make(StringMap)
}

func InitStringMap(ctx map[string]string) StringMap {
	sm := make(StringMap)
	if ctx != nil {
		for k, v := range ctx {
			sm[k] = v
		}
	}

	return sm
}

func (sm StringMap) Put(key, value string) StringMap {
	sm[key] = value

	return sm
}

func (sm StringMap) Get(key string) (string, bool) {
	value, ok := sm[key]

	return value, ok
}

func (sm StringMap) Remove(key string) {
	if sm.Has(key) {
		delete(sm, key)
	}
}

func (sm StringMap) Has(key string) bool {
	_, ok := sm[key]

	return ok
}

func (sm StringMap) Size() int {
	return len(sm)
}

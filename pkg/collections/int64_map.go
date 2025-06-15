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

func NewInt64Map(cap ...int) Int64Map {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(Int64Map, cap[0])
		}
	}

	return make(Int64Map)
}

func InitInt64Map(ctx map[int64]int64) Int64Map {
	sm := make(Int64Map)
	if ctx != nil {
		for k, v := range ctx {
			sm[k] = v
		}
	}

	return sm
}

func (im Int64Map) Put(key, value int64) Int64Map {
	im[key] = value

	return im
}

func (im Int64Map) Get(key int64) (int64, bool) {
	value, ok := im[key]

	return value, ok
}

func (im Int64Map) Remove(key int64) {
	if im.Has(key) {
		delete(im, key)
	}
}

func (im Int64Map) Has(key int64) bool {
	_, ok := im[key]

	return ok
}

func (im Int64Map) Size() int {
	return len(im)
}

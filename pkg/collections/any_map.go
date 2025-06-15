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

import "github.com/photowey/kitgo/pkg/casts"

func NewAnyMap() AnyMap {
	return make(AnyMap)
}

func InitAnyMap(ctx map[any]any) AnyMap {
	am := make(AnyMap)
	if ctx != nil {
		for k, v := range ctx {
			am[k] = v
		}
	}

	return am
}

func (am AnyMap) Put(key, value any) AnyMap {
	am[key] = value

	return am
}

func (am AnyMap) Get(key any) (any, bool) {
	value, ok := am[key]

	return value, ok
}

func (am AnyMap) Remove(key any) {
	if am.Has(key) {
		delete(am, key)
	}
}

func (am AnyMap) Has(key any) bool {
	_, ok := am[key]

	return ok
}

func (am AnyMap) Size() int {
	return len(am)
}

// ---------------------------------------------------------------- type value

func (am AnyMap) GetString(key string) (string, bool) {
	if value, ok := am[key]; ok {
		return casts.ToStringB(value)
	}

	return "", false
}

// ---------------------------------------------------------------- int

func (am AnyMap) GetInt(key string) (int, bool) {
	if value, ok := am[key]; ok {
		return casts.ToIntB(value)
	}

	return 0, false
}

func (am AnyMap) GetInt64(key string) (int64, bool) {
	if value, ok := am[key]; ok {
		return casts.ToInt64B(value)
	}

	return 0, false
}

func (am AnyMap) GetInt32(key string) (int32, bool) {
	if value, ok := am[key]; ok {
		return casts.ToInt32B(value)
	}

	return 0, false
}

func (am AnyMap) GetInt16(key string) (int16, bool) {
	if value, ok := am[key]; ok {
		return casts.ToInt16B(value)
	}

	return 0, false
}

func (am AnyMap) GetInt8(key string) (int8, bool) {
	if value, ok := am[key]; ok {
		return casts.ToInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- uint

func (am AnyMap) GetUInt(key string) (uint, bool) {
	if value, ok := am[key]; ok {
		return casts.ToUIntB(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt64(key string) (uint64, bool) {
	if value, ok := am[key]; ok {
		return casts.ToUInt64B(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt32(key string) (uint32, bool) {
	if value, ok := am[key]; ok {
		return casts.ToUInt32B(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt16(key string) (uint16, bool) {
	if value, ok := am[key]; ok {
		return casts.ToUInt16B(value)
	}

	return 0, false
}

func (am AnyMap) GetUInt8(key string) (uint8, bool) {
	if value, ok := am[key]; ok {
		return casts.ToUInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- float

func (am AnyMap) GetFloat64(key string) (float64, bool) {
	if value, ok := am[key]; ok {
		return casts.ToFloat64B(value)
	}

	return 0, true
}

func (am AnyMap) GetFloat32(key string) (float32, bool) {
	if value, ok := am[key]; ok {
		return casts.ToFloat32B(value)
	}

	return 0, true
}

// ---------------------------------------------------------------- bool

func (am AnyMap) GetBool(key string) (bool, bool) {
	if value, ok := am[key]; ok {
		return casts.ToBoolB(value)
	}

	return false, false
}

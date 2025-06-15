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

import (
	"github.com/photowey/kitgo/pkg/casts"
)

func NewMixedMap(cap ...int) MixedMap {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(MixedMap, cap[0])
		}
	}

	return make(MixedMap)
}

func InitMixedMap(ctx map[string]any) MixedMap {
	sm := make(MixedMap)
	if ctx != nil {
		for k, v := range ctx {
			sm[k] = v
		}
	}

	return sm
}

func (mm MixedMap) Put(key string, value any) MixedMap {
	mm[key] = value

	return mm
}

func (mm MixedMap) Get(key string) (any, bool) {
	value, ok := mm[key]

	return value, ok
}

func (mm MixedMap) Remove(key string) {
	if mm.Has(key) {
		delete(mm, key)
	}
}

func (mm MixedMap) Has(key string) bool {
	_, ok := mm[key]

	return ok
}

func (mm MixedMap) Size() int {
	return len(mm)
}

// ---------------------------------------------------------------- type value

func (mm MixedMap) GetString(key string) (string, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToStringB(value)
	}

	return "", false
}

// ---------------------------------------------------------------- int

func (mm MixedMap) GetInt(key string) (int, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToIntB(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt64(key string) (int64, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToInt64B(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt32(key string) (int32, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToInt32B(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt16(key string) (int16, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToInt16B(value)
	}

	return 0, false
}

func (mm MixedMap) GetInt8(key string) (int8, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- uint

func (mm MixedMap) GetUInt(key string) (uint, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToUIntB(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt64(key string) (uint64, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToUInt64B(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt32(key string) (uint32, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToUInt32B(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt16(key string) (uint16, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToUInt16B(value)
	}

	return 0, false
}

func (mm MixedMap) GetUInt8(key string) (uint8, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToUInt8B(value)
	}

	return 0, false
}

// ---------------------------------------------------------------- float

func (mm MixedMap) GetFloat64(key string) (float64, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToFloat64B(value)
	}

	return 0, true
}

func (mm MixedMap) GetFloat32(key string) (float32, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToFloat32B(value)
	}

	return 0, true
}

// ---------------------------------------------------------------- bool

func (mm MixedMap) GetBool(key string) (bool, bool) {
	if value, ok := mm[key]; ok {
		return casts.ToBoolB(value)
	}

	return false, false
}

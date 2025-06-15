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

package casts

import (
	"time"

	"github.com/spf13/cast"
)

// ---------------------------------------------------------------- bool

func ToBoolB(src any) (bool, bool) {
	v, err := cast.ToBoolE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- int

func ToInt64B(src any) (int64, bool) {
	v, err := cast.ToInt64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToInt32B(src any) (int32, bool) {
	v, err := cast.ToInt32E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToInt16B(src any) (int16, bool) {
	v, err := cast.ToInt16E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToInt8B(src any) (int8, bool) {
	v, err := cast.ToInt8E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToIntB(src any) (int, bool) {
	v, err := cast.ToIntE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToUInt64B(src any) (uint64, bool) {
	v, err := cast.ToUint64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToUInt32B(src any) (uint32, bool) {
	v, err := cast.ToUint32E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToUInt16B(src any) (uint16, bool) {
	v, err := cast.ToUint16E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToUInt8B(src any) (uint8, bool) {
	v, err := cast.ToUint8E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToUIntB(src any) (uint, bool) {
	v, err := cast.ToUintE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- float

func ToFloat64B(src any) (float64, bool) {
	v, err := cast.ToFloat64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToFloat32B(src any) (float32, bool) {
	v, err := cast.ToFloat32E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- string

func ToStringB(src any) (string, bool) {
	v, err := cast.ToStringE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- map<string,x>

func ToStringMapB(src any) (map[string]any, bool) {
	v, err := cast.ToStringMapE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToStringMapStringB(src any) (map[string]string, bool) {
	v, err := cast.ToStringMapStringE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToStringMapStringSliceB(src any) (map[string][]string, bool) {
	v, err := cast.ToStringMapStringSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToStringMapInt64B(src any) (map[string]int64, bool) {
	v, err := cast.ToStringMapInt64E(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- slice

func ToSliceB(src any) ([]any, bool) {
	v, err := cast.ToSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToStringSliceB(src any) ([]string, bool) {
	v, err := cast.ToStringSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

func ToIntSliceB(src any) ([]int, bool) {
	v, err := cast.ToIntSliceE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

// ---------------------------------------------------------------- time

func ToTimeB(src any) (time.Time, bool) {
	v, err := cast.ToTimeE(src)
	if err != nil {
		return v, false
	}

	return v, true
}

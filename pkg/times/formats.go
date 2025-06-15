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

package times

import (
	"time"
)

// ---------------------------------------------------------------- format time

// d == default == default pattern

func Nowd() string {
	return NowPattern(YyyyMMddHHmmss)
}

func NowPattern(pattern TimePattern) string {
	now := time.Now()
	return FormatNow(now, pattern)
}

// ----------------------------------------------------------------

func FormatTime(now time.Time) string {
	return FormatTimePattern(now, YyyyMMddHHmmss)
}

func FormatTimePattern(now time.Time, pattern TimePattern) string {
	return now.Format(String(pattern))
}

// ----------------------------------------------------------------

func FormatNow(now time.Time, pattern TimePattern) string {
	return now.Format(String(pattern))
}

func FormatNowd(now time.Time) string {
	return FormatNow(now, YyyyMMddHHmmss)
}

// ----------------------------------------------------------------

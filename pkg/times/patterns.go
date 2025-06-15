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

type TimePattern string

const (
	YyyyMMddHHmmss        TimePattern = "2006-01-02 15:04:05"
	YyyyMMdd              TimePattern = "2006-01-02"
	HHmmss                TimePattern = "15:04:05"
	YyyyMMddShort         TimePattern = "20060102"
	YyyyMMddHHmmssShort   TimePattern = "20060102150405"
	YyyyMMddHHmmssChinese TimePattern = "2006年01月02日 15:04:05"
)

func ToTimePattern(pattern string) TimePattern {
	return TimePattern(pattern)
}

func String(pattern TimePattern) string {
	return string(pattern)
}

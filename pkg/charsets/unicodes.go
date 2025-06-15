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

package charsets

import (
	"strconv"
	"strings"
)

const (
	oldUniChar = `\\u`
	newUniChar = `\u`
)

// ToZhCN 将 unicode 转化为简体中文
func ToZhCN(text string) string {
	zhCN, err := strconv.Unquote(strings.Replace(strconv.Quote(text), oldUniChar, newUniChar, -1))
	if err != nil {
		return text
	}

	return zhCN
}

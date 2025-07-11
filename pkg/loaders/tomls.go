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

package loaders

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func LoadToml(tomlPath string, tomlBean any) (err error) {
	if NotHasSuffix(tomlPath, SuffixToml) {
		return fmt.Errorf("TOML file must have .toml extension: %v", tomlPath)
	}

	_, err = toml.DecodeFile(tomlPath, tomlBean)

	return
}

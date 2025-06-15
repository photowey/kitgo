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
	"gopkg.in/yaml.v3"
	"os"
)

func LoadYaml(yamlPath string, yamlBean any) (err error) {
	if NotHasSuffix(yamlPath, SuffixYml) && NotHasSuffix(yamlPath, SuffixYaml) {
		return fmt.Errorf("YAML file must have .yml or .yaml extension: %v", yamlPath)
	}
	yamlStream, err := os.ReadFile(yamlPath)
	if err != nil {
		return fmt.Errorf("error reading YAML file [%v]: %v", yamlPath, err)
	}
	err = yaml.Unmarshal(yamlStream, yamlBean)

	return
}

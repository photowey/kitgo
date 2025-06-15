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
	"fmt"
	"time"
)

func ParseTime(datetime string) (time.Time, error) {
	return ParseTimed(datetime, YyyyMMddHHmmss)
}

func ParseTimed(datetime string, pattern TimePattern) (time.Time, error) {
	return ParseTimez(datetime, TimezoneAsiaShanghai, pattern)
}

func ParseTimez(datetime, timezone string, pattern TimePattern) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load timezone '%s': %w", timezone, err)
	}

	layout := String(pattern)
	parsedTime, err := time.ParseInLocation(layout, datetime, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf(
			"failed to parse datetime '%s' using layout '%s' in timezone '%s': %w",
			datetime,
			layout,
			timezone,
			err,
		)
	}

	return parsedTime, nil
}

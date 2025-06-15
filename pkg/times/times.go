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

var epoch = time.Now().AddDate(-1, -1, -1)

func CurrentTimeMicros() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CurrentTimeSeconds() int64 {
	return time.Now().Unix()
}

func Now() time.Duration {
	return time.Since(epoch)
}

func Since(past time.Duration) time.Duration {
	return time.Since(epoch) - past
}

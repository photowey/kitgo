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

package stopwatch

import (
	"fmt"
	"github.com/photowey/kitgo/pkg/times"
	"time"
)

var epoch = time.Now().AddDate(-1, -1, -1)

type StopWatch struct {
	start time.Duration
}

func NewDefaultStopWatch() *StopWatch {
	return &StopWatch{
		start: times.Now(),
	}
}

func NewStopWatch(start time.Duration) *StopWatch {
	return &StopWatch{
		start: start,
	}
}

func (sw *StopWatch) Duration() time.Duration {
	return times.Since(sw.start)
}

func (sw *StopWatch) Elapse() string {
	return times.Since(sw.start).String()
}

func (sw *StopWatch) ElapseMillis() string {
	return fmt.Sprintf("%.2f ms", float32(times.Since(sw.start))/float32(time.Millisecond))
}

func (sw *StopWatch) ElapseDurationMs(past time.Duration) string {
	return fmt.Sprintf("%.2f ms", float32(times.Since(past))/float32(time.Millisecond))
}

func (sw *StopWatch) ElapseMicro() string {
	return fmt.Sprintf("%.2f macros", float32(times.Since(sw.start))/float32(time.Microsecond))
}

func (sw *StopWatch) ElapseDurationMicro(past time.Duration) string {
	return fmt.Sprintf("%.2f macros", float32(times.Since(past))/float32(time.Microsecond))
}

func (sw *StopWatch) ElapseAuto() string {
	cost := float32(times.Since(sw.start)) / float32(time.Millisecond)
	if cost > 0 {
		return fmt.Sprintf("%.2f ms", cost)
	}

	return sw.ElapseMicro()
}

func (sw *StopWatch) ElapseDurationAuto(past time.Duration) string {
	cost := float32(times.Since(past)) / float32(time.Millisecond)
	if cost > 0 {
		return fmt.Sprintf("%.2f ms", cost)
	}

	return sw.ElapseDurationMicro(past)
}

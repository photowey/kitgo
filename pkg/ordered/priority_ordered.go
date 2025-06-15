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

package ordered

import (
	"sort"
)

type PriorityOrdered interface {
	Ordered
}

type PrioritySorter []PriorityOrdered

func (sorter PrioritySorter) Len() int {
	return len(sorter)
}

func (sorter PrioritySorter) Less(i, j int) bool {
	return sorter[i].Order() < sorter[j].Order()
}

func (sorter PrioritySorter) Swap(i, j int) {
	sorter[i], sorter[j] = sorter[j], sorter[i]
}

func NewPrioritySorter(notifiers ...Ordered) PrioritySorter {
	sorter := make(PrioritySorter, len(notifiers))
	for i, notifier := range notifiers {
		sorter[i] = notifier
	}

	return sorter
}

func PrioritySort(sorter PrioritySorter) {
	sort.Slice(sorter, func(i, j int) bool {
		return sorter[i].Order() < sorter[j].Order()
	})
}

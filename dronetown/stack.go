/*
 * Copyright 2021 Banco Bilbao Vizcaya Argentaria, S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dronetown

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (string, bool) {
	if len(s.items) > 0 {
		last := len(s.items) - 1
		item := s.items[last]
		s.items = s.items[:last]
		return item, true
	}
	return "", false
}

func (s *Stack) Fork() Stack {
	return List2Stack(s.items)
}

func (s *Stack) String() string {
	maxL := 20
	for _, item := range s.items {
		size := len(item)
		if size > maxL {
			maxL = size
		}
	}
	if maxL%2 == 0 {
		maxL += 1
	}

	out := "\n--- Stack ---\n"

	for i, item := range s.items {
		if i == 0 {
			out += "Top\n"
			out += fmt.Sprintf("\t%s\n", item)
			out += "Rest\n"
		} else {
			out += fmt.Sprintf("\t%s\n", item)
		}
	}
	out += "--- Stack ---\n"
	return out
}

func List2Stack(list []string) Stack {
	s := Stack{}
	for _, item := range list {
		s.Push(item)
	}
	return s
}

type Brain = Stack
type Action = func(Brain, Stack) error

func RunShop(m map[string]Action, brain Brain) error {
	stack := Stack{}
	b := brain.Fork()
	for {
		currentItem, ok := b.Pop()
		if !ok {
			break
		}
		f, found := m[currentItem]
		if found {
			e := f(b, stack)
			return e
		} else {
			stack.Push(currentItem)
		}
	}
	return &EOS{stack}
}

type EOS struct { // by Canon
	S Stack
}

func (e *EOS) Error() string {
	return fmt.Sprintf("End Of Stack Reached:\n%s", e.S)
}

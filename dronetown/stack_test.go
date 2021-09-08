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

import (
	"testing"
)

func TestTryToRetrieveFromEmtpyStack(t *testing.T) {
	stack := Stack{}

	_, ok := stack.Pop()

	if ok {
		t.Log("Unexpected success!")
		t.FailNow()
	}
}

func TestAppendAndRecoverAndItem(t *testing.T) {
	item := "cheese"

	stack := Stack{}

	stack.Push(item)
	res, _ := stack.Pop()

	if res != item {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", res, item)
		t.FailNow()
	}
}

func TestAppendServeralAndRecoverInCorrectOrder(t *testing.T) {
	stack := Stack{}

	stack.Push("Last")
	stack.Push("Middle")
	stack.Push("First")

	f, ok := stack.Pop()
	if !ok {
		t.Log("unexpected empty stack")
		t.Fail()
	}
	if f != "First" {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", "First", f)
		t.Fail()
	}
	f, ok = stack.Pop()
	if !ok {
		t.Log("unexpected empty stack")
		t.Fail()
	}
	if f != "Middle" {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", "Middle", f)
		t.Fail()
	}
	f, ok = stack.Pop()
	if !ok {
		t.Log("unexpected empty stack")
		t.Fail()
	}
	if f != "Last" {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", "Last", f)
		t.Fail()
	}
}

func TestCreateFromList(t *testing.T) {
	inp := []string{
		"Last",
		"Middle",
		"First",
	}

	stack := List2Stack(inp)
	t.Log(&stack)

	f, ok := stack.Pop()
	if !ok {
		t.Log("unexpected empty stack")
		t.Fail()
	}
	if f != "First" {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", "First", f)
		t.Fail()
	}
	f, ok = stack.Pop()
	if !ok {
		t.Log("unexpected empty stack")
		t.Fail()
	}
	if f != "Middle" {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", "Middle", f)
		t.Fail()
	}
	f, ok = stack.Pop()
	if !ok {
		t.Log("unexpected empty stack")
		t.Fail()
	}
	if f != "Last" {
		t.Logf("Unexpected item recovered:\n\t%s\nI expect:\n\t%s", "Last", f)
		t.Fail()
	}
}

func TestForkStack(t *testing.T) {
	stack := Stack{}

	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	fork := stack.Fork()

	same := true
	for {
		a, aok := stack.Pop()
		b, bok := fork.Pop()

		if !aok {
			break
		}

		if aok != bok {
			same = false
			break
		}

		if a != b {
			same = false
			break
		}
	}
	if !same {
		t.Log("expect the same, but...")
		t.Log("Stack")
		t.Log(stack)
		t.Log("Fork")
		t.Log(fork)
		t.Fail()
	}
}

func TestRunShop(t *testing.T) {
	flag := false
	sel := map[string]func(Brain, Stack) error{
		"buy": func(b Brain, s Stack) error {
			id, _ := s.Pop()
			item, _ := s.Pop()

			t.Logf("id: %s item: %s", id, item)
			if id != "soul" {
				t.Logf("soul id expected but: %s", id)
				t.Fail()
			}

			if item != "happiness" {
				t.Logf("happiness item expected but: %s", item)
				t.Fail()
			}

			flag = true
			return nil
		},
	}

	brain := List2Stack([]string{
		"go",
		"home",
		"buy",
		"soul",
		"happiness",
	})

	RunShop(sel, brain)

	if !flag {
		t.Log("Expect to run buy selector but not")
		t.Fail()
	}
}

func TestRunShopWithContext(t *testing.T) {
	flag := false

	buyWContext := func(drone Drone) Action {
		return func(b Brain, s Stack) error {
			id, _ := s.Pop()
			item, _ := s.Pop()

			t.Logf("id: %s item: %s", id, item)
			t.Logf("drone to Manipulate: %v", drone)

			flag = true
			return nil
		}
	}

	brain := List2Stack([]string{
		"go",
		"home",
		"buy",
		"happiness",
	})

	drone := Drone{
		Brain: brain,
	}

	sel := map[string]func(Brain, Stack) error{
		"buy": buyWContext(drone),
	}

	RunShop(sel, brain)

	if !flag {
		t.Log("Expect to run buy selector but not")
		t.Fail()
	}
}

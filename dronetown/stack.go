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

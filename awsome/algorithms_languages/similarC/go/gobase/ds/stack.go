// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ds

// Stack is a FILO stack
type Stack struct {
	v []interface{}
}

// NewStack returns a new stack
func NewStack() *Stack {
	return &Stack{v: make([]interface{}, 0)}
}

// Push pushes a value to the stack
func (s *Stack) Push(v interface{}) {
	s.v = append(s.v, v)
}

// Pop pops the top value out of the stack
func (s *Stack) Pop() interface{} {
	v := s.v[len(s.v)]
	s.v = s.v[:len(s.v)-1]
	return v
}

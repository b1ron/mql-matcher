package ast

import (
	"reflect"
)

// leaf types
const (
	EQ = iota
	LTE
	LT
	GT
	GTE
)

// restrict nesting to 2 levels for trivial cases for now
const maxDepth = 2

type Expr interface {
	eval() any
}

type tree struct {
	root *node
}

type node struct {
	key  string // can also have a key which could be an operator or ident
	expr Expr
}

type leaf struct {
	key   any
	value any
}

// returns a leaf value.
func (l *leaf) eval() any {
	switch l.value.(type) {
	// needs the same precedence as the switch statement in isSubset
	case int:
		return l.value.(int)
	case string:
		return l.value.(string)
	case []any:
		return l.value.([]any)
	}
	return nil
}

func (l *leaf) field() string {
	return l.key.(string)
}

// evaluates a node expression
func (n *node) eval() any {
	switch n.expr.(type) {
	case *node:
		return n.expr.eval()
	case *leaf:
		return n.expr.eval()
	}
	return nil
}

// evaluates an expression tree
func (t *tree) eval() any {
	switch t.root.expr.(type) {
	case *node:
		return t.root.expr.eval()
	case *leaf:
		return t.root.expr.eval()
	}
	return nil
}

func (t *tree) isNil() bool {
	return t.root == nil || t.root.expr == nil
}

type stack struct {
	frames   []any
	i        int
	children int
}

func (s *stack) push(v any) {
	s.frames = append(s.frames, v)
	s.i++
	s.children = s.i
}

func (s *stack) pop() any {
	v := s.frames[s.i-1]
	s.i--
	s.frames = s.frames[:s.i]
	return v
}

func (s *stack) size() int {
	return s.children
}

func (s *stack) empty() bool {
	return s.frames == nil || len(s.frames) == 0
}

func (t *tree) depth() int {
	l := t.root.expr.(*leaf)

	s := stack{}
	s.push(l.value)

	level := 0
	for !s.empty() {
		item := s.pop()
		switch item := item.(type) {
		case []any:
			level++
			for _, v := range item {
				s.push(v)
			}
		}
	}

	return level
}

func isSubset(a, b *tree) bool {
	if a.isNil() || b.isNil() {
		return false
	}

	if a.root.expr.(*leaf).field() != b.root.expr.(*leaf).field() {
		return false
	}

	// this switch statement is fragile
	// it uses precedence rules to determine the order of evaluation
	lhs := a.eval()
	switch lhs.(type) {
	case int:
		switch b.eval().(type) {
		case int:
			return lhs.(int) == b.eval().(int)
		case string:
			return false
		case []any:
			return contains(lhs, b.eval().([]any))
		}
	case string:
		switch b.eval().(type) {
		case int:
			return false
		case string:
			return lhs.(string) == b.eval().(string)
		case []any:
			return contains(lhs, b.eval().([]any))
		}
	case []any:
		switch b.eval().(type) {
		case int:
			return false
		case string:
			return false
		case []any:
			return containsAll(lhs.([]any), b.eval().([]any))
		}
	}
	return false
}

func contains[E comparable](v E, s []E) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}
	return false
}

func containsAll[E comparable](s1, s2 []E) bool {
	return reflect.DeepEqual(s1, s2)
}

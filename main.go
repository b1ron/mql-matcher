package main

import "fmt"

// leaf types
const (
	EQ = iota
	LTE
	LT
	GT
	GTE
)

type Expr interface {
	eval() int
}

type tree struct {
	node *node
}

type node struct {
	expr Expr
}

type leaf struct {
	key   any
	value any
}

// returns a leaf value.
func (l *leaf) eval() int {
	switch l.value.(type) {
	case *node:
		return l.value.(*node).expr.eval()
	case []any:
		return len(l.value.([]any))
	}
	return -1
}

func main() {
	// build a tree of simple nested expressions
	t := &tree{
		node: &node{
			expr: &leaf{
				key: "foo",
				value: &node{
					expr: &leaf{
						key:   "bar",
						value: []any{1, 2, 3},
					},
				},
			},
		},
	}
	fmt.Println(t.node.expr.eval())
}

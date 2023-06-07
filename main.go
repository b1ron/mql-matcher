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
	// ...
	left  *node
	right *node
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
	case []any:
		// TODO iterate the slice
		return len(l.value.([]any))
	case int:
		return l.value.(int)
	}
	return -1
}

// evaluates a node expression
func (n *node) eval() int {
	switch n.expr.(type) {
	case *node:
		return n.expr.eval()
	case *leaf:
		return n.expr.eval()
	}
	return -1
}

// evaluates an expression tree
func (t *tree) eval() int {
	switch t.node.expr.(type) {
	case *node:
		return t.node.expr.eval()
	case *leaf:
		return t.node.expr.eval()
	}
	return -1
}

func traverse(t *tree) {
	// TODO
}

func main() {
	// build a tree of simple nested expressions
	t1 := &tree{
		node: &node{
			expr: &leaf{
				key:   "foo",
				value: 1,
			},
		},
	}

	t2 := &tree{
		node: &node{
			expr: &leaf{
				key: "foo",
				value: []any{
					[]any{1, 2, 3},
				},
			},
		},
	}
	// traverses the expression tree post-order and compares each leaf node
	// if A is a subset of B, then A EQ B.
	tree := &tree{left: t1.node, right: t2.node}
	fmt.Println(tree.eval())
}

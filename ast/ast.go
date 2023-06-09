package ast

// leaf types
const (
	EQ = iota
	LTE
	LT
	GT
	GTE
)

type Expr interface {
	eval() any
}

type tree struct {
	root *node
	next *node // make it simple for now
}

type node struct {
	expr Expr
}

type leaf struct {
	key   any
	value any
}

// returns a leaf value.
func (l *leaf) eval() any {
	switch l.value.(type) {
	case []any:
		return l.value.([]any)
	case int:
		return l.value
	}
	return nil
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

func isSubset(a, b *tree) bool {
	lhs := a.eval()
	switch lhs.(type) {
	case int:
		switch b.eval().(type) {
		case int:
			return lhs.(int) == b.eval().(int)
		case []any:
			return contains(b.eval().([]any), lhs.(int))
		}
	}
	return false
}

type number interface {
	int | float64 | float32 | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func contains[K comparable, V number]([]any, any) bool {
	return false
}

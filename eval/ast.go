package eval

// clause
//   = leaf_clause
//   / expression_tree_clause
//   / expression_clause
//   / where_clause
//   / text_clause

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

type ident map[string]any

var idents = ident{
	"leaf-clause":            expr{},
	"expression-tree-clause": nil,
	"expression-clause":      expr{},
	"where-clause":           nil,
	"text-clause":            nil,
	"leaf-value":             nil,
}

type expr struct {
	id  ident
	key string
	e   []Expr
}

type leaf struct {
	id    ident
	key   string
	value any
}

type value struct {
	id ident
	e  Expr
}

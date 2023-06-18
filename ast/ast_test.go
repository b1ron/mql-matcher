package ast

import (
	"testing"
)

func TestSubset(t *testing.T) {
	tests := []struct {
		a    *tree
		b    *tree
		want bool
	}{
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: 0}}},
			want: false,
		},
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			want: true,
		},
		// slightly more complex
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, 2}}}},
			want: true,
		},
		// mismatched keys
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			b:    &tree{root: &node{expr: &leaf{key: "b", value: 1}}},
			want: false,
		},
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			b:    &tree{root: nil},
			want: false,
		},
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 1}}},
			b:    &tree{root: &node{key: "a", expr: nil}},
			want: false,
		},
	}
	for _, test := range tests {
		got := isSubset(test.a, test.b)
		if got != test.want {
			t.Errorf("isSubset(%v, %v) = %v, want %v", test.a.eval(), test.b.eval(), got, test.want)
		}
	}
}

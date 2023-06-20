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
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: "x"}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: "x"}}},
			want: true,
		},
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: "x"}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: "z"}}},
			want: false,
		},
		// strict array comparison
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1}}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, 2}}}},
			want: false,
		},
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, 2}}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, 2}}}},
			want: true,
		},
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, 2}}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1}}}},
			want: false,
		},
		// deeply nested array that exceeds max depth
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1}}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, []any{2, []any{3, []any{4, nil}}}}}}},
			want: false,
		},
		// FIXME: this should be true, elements should be compared at the second-level
		{
			a:    &tree{root: &node{expr: &leaf{key: "a", value: 3}}},
			b:    &tree{root: &node{expr: &leaf{key: "a", value: []any{1, []any{2, 3}}}}},
			want: true,
		},
	}
	for _, test := range tests {
		got := isSubset(test.a, test.b)
		if got != test.want {
			t.Errorf("isSubset(%v, %v) = %v, want %v", test.a.eval(), test.b.eval(), got, test.want)
		}
	}
}

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
			&tree{
				node: &node{
					expr: &leaf{
						key:   "a",
						value: 1,
					},
				},
			},
			&tree{
				node: &node{
					expr: &leaf{
						key:   "a",
						value: 0,
					},
				},
			},
			false,
		},
		{
			&tree{
				node: &node{
					expr: &leaf{
						key:   "a",
						value: 1,
					},
				},
			},
			&tree{
				node: &node{
					expr: &node{
						expr: &leaf{
							key:   "a",
							value: []int{1, 2, 3},
						},
					},
				},
			},
			true,
		},
	}
	for i, test := range tests {
		t.Log(test.a.eval(), test.b.eval())
		if got := isSubset(test.a, test.b); got != test.want {
			t.Errorf("#%d isSubset(%v, %v) = %v", i, test.a, test.b, got)
		}
	}
}

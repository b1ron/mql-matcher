package eval

import "testing"

func TestEval(t *testing.T) {
	tests := []struct {
		e    Expr
		want any
	}{
		{
			e:    &expr{},
			want: nil,
		},
		{
			e: &expr{
				id: idents["expression-clause"].(ident),
				e: []Expr{
					&expr{
						// id: idents["leaf-clause"].(ident),
						e: []Expr{
							&leaf{
								// id: idents["leaf-value"].(ident),
								key: "a",
								value: value{
									literal: 1,
								},
							},
						},
					},
					&expr{},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		if got := tt.e.Eval(); got != tt.want {
			t.Errorf("Eval() = %v, want %v", got, tt.want)
		}
	}
}

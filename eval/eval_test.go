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
					&leaf{
						id: idents["leaf-clause"].(ident),
						value: value{
							id:      idents["leaf-value"].(ident),
							literal: "hello",
						},
					},
					&leaf{
						id: idents["leaf-clause"].(ident),
						value: value{
							id:      idents["leaf-value"].(ident),
							literal: "world",
						},
					},
				},
			},
			want: "helloworld",
		},
	}
	for _, tt := range tests {
		if got := tt.e.Eval(); got != tt.want {
			t.Errorf("Eval() = %v, want %v", got, tt.want)
		}
	}
}

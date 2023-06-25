package eval

func (l *leaf) Eval() any {
	switch {
	case l.id["leaf-value"] != nil:
		return l.value.literal
	case l.id["operator-expression"] != nil:
		// return operators
	case l.id["value-operator"] != nil:
		return l.value.e.Eval()
	}
	return nil
}

func (e *expr) Eval() any {
	switch {
	case e.id["expression-clause"] != nil:
		// has many leaf-clauses
		for _, lc := range e.e {
			lc.Eval()
		}
	case e.id["value-operator"] != nil:
		// TODO implement
	}
	return nil
}

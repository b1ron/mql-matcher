package eval

func (l *leaf) Eval() any {
	switch {
	case l.id["leaf-value"] != nil:
		// return value
	case l.id["operator-expression"] != nil:
		// return operators
	case l.id["value-operator"] != nil:
		// return leaf-value
	}

	return nil
}

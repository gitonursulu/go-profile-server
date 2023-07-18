// pkg/predicate/predicate.go
package predicate

type PredicateNode struct {
	Operator string
	Left     Expression
	Right    Expression
}

type ValueNode struct {
	Value interface{}
}

type Expression interface {
	Evaluate() bool
}

func (p *PredicateNode) Evaluate() bool {
	switch p.Operator {
	case "AND":
		return p.Left.Evaluate() && p.Right.Evaluate()
	case "OR":
		return p.Left.Evaluate() || p.Right.Evaluate()
	case "NOT":
		return !p.Right.Evaluate()
	default:
		panic("Geçersiz koşul operatörü: " + p.Operator)
	}
}

func (v *ValueNode) Evaluate() bool {
	return v.Value.(bool)
}

package predicate

import (
	"fmt"
)

type PredicateNode struct {
	Operator string
	Left     Expression
	Right    Expression
}

type ValueSpec int

const (
	Other  ValueSpec = iota // 0
	Column                  // 1
)

type ValueNode struct {
	Value     interface{}
	ValueSpec ValueSpec
}

type Expression interface {
	Evaluate() string
}

func (p *PredicateNode) Evaluate() string {
	switch p.Operator {
	case "AND":
		left := p.Left.Evaluate()
		right := p.Right.Evaluate()
		return fmt.Sprintf("(%s AND %s)", left, right)
	case "OR":
		left := p.Left.Evaluate()
		right := p.Right.Evaluate()
		return fmt.Sprintf("(%s OR %s)", left, right)
	case "NOT":
		if p.Left != nil {
			panic("aallaaaaah")
		}
		right := p.Right.Evaluate()
		return fmt.Sprintf("(NOT %s)", right)
	case ">":
		left := p.Left.Evaluate()
		right := p.Right.Evaluate()
		return fmt.Sprintf("(%s > %s)", left, right)
	case "<":
		left := p.Left.Evaluate()
		right := p.Right.Evaluate()
		return fmt.Sprintf("(%s < %s)", left, right)
	case "=":
		left := p.Left.Evaluate()
		right := p.Right.Evaluate()
		return fmt.Sprintf("(%s = %s)", left, right)
	default:
		panic("Geçersiz koşul operatörü: " + p.Operator)
	}
}

func (v *ValueNode) Evaluate() string {
	valueSpec := v.ValueSpec
	switch value := v.Value.(type) {
	case string:
		if valueSpec == Column {
			return fmt.Sprintf("%s", value)
		}
		return fmt.Sprintf("'%s'", value)
	case int, int64, float32, float64, bool:
		return fmt.Sprintf("%v", value)
	default:
		panic("Geçersiz değer türü: " + fmt.Sprintf("%T", value))
	}
}

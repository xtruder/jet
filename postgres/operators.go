package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// NOT returns negation of bool expression result
var NOT = jet.NOT

// BIT_NOT inverts every bit in integer expression result
var BIT_NOT = jet.BIT_NOT

func ARRAY(expressions ...Expression) ArrayExpression {
	return ArrayExp(jet.NewSequenceOperatorExpression(expressions, "ARRAY"))
}

func OVERLAPS(lhs Expression, rhs Expression) BoolExpression {
	return jet.NewBinaryBoolOperatorExpression(lhs, rhs, "&&")
}

func CONTAINS(lhs Expression, rhs Expression) BoolExpression {
	return jet.NewBinaryBoolOperatorExpression(lhs, rhs, "@>")
}

func IS_CONTAINED_BY(lhs Expression, rhs Expression) BoolExpression {
	return jet.NewBinaryBoolOperatorExpression(lhs, rhs, "<@")
}

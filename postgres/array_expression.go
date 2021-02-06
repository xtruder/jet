package postgres

import "github.com/go-jet/jet/v2/internal/jet"

type ArrayExpression interface {
	Expression

	EQ(rhs ArrayExpression) BoolExpression
	NOT_EQ(rhs ArrayExpression) BoolExpression

	LT(rhs ArrayExpression) BoolExpression
	LT_EQ(rhs ArrayExpression) BoolExpression
	GT(rhs ArrayExpression) BoolExpression
	GT_EQ(rhs ArrayExpression) BoolExpression

	CONTAINS(rhs ArrayExpression) BoolExpression
	IS_CONTAINED_BY(rhs ArrayExpression) BoolExpression
	OVERLAPS(rhs ArrayExpression) BoolExpression

	CONCAT(rhs Expression) ArrayExpression
}

type arrayInterfaceImpl struct {
	parent ArrayExpression
}

func (s *arrayInterfaceImpl) EQ(rhs ArrayExpression) BoolExpression {
	return jet.Eq(s.parent, rhs)
}

func (s *arrayInterfaceImpl) NOT_EQ(rhs ArrayExpression) BoolExpression {
	return jet.NotEq(s.parent, rhs)
}

func (s *arrayInterfaceImpl) GT(rhs ArrayExpression) BoolExpression {
	return jet.Gt(s.parent, rhs)
}

func (s *arrayInterfaceImpl) GT_EQ(rhs ArrayExpression) BoolExpression {
	return jet.GtEq(s.parent, rhs)
}

func (s *arrayInterfaceImpl) LT(rhs ArrayExpression) BoolExpression {
	return jet.Lt(s.parent, rhs)
}

func (s *arrayInterfaceImpl) LT_EQ(rhs ArrayExpression) BoolExpression {
	return jet.LtEq(s.parent, rhs)
}

func (s *arrayInterfaceImpl) CONTAINS(rhs ArrayExpression) BoolExpression {
	return CONTAINS(s.parent, rhs)
}

func (s *arrayInterfaceImpl) IS_CONTAINED_BY(rhs ArrayExpression) BoolExpression {
	return IS_CONTAINED_BY(s.parent, rhs)
}

func (s *arrayInterfaceImpl) OVERLAPS(rhs ArrayExpression) BoolExpression {
	return OVERLAPS(s.parent, rhs)
}

func (s *arrayInterfaceImpl) CONCAT(rhs Expression) ArrayExpression {
	return newBinaryArrayOperatorExpression(s.parent, rhs, jet.StringConcatOperator)
}

//---------------------------------------------------//
func newBinaryArrayOperatorExpression(lhs, rhs Expression, operator string) ArrayExpression {
	return ArrayExp(jet.NewBinaryOperatorExpression(lhs, rhs, operator))
}

//---------------------------------------------------//

type arrayExpressionWrapper struct {
	arrayInterfaceImpl
	Expression
}

func newArrayExpressionWrap(expression Expression) ArrayExpression {
	arrayExpressionWrap := arrayExpressionWrapper{Expression: expression}
	arrayExpressionWrap.arrayInterfaceImpl.parent = &arrayExpressionWrap
	return &arrayExpressionWrap
}

// ArrayExp is array expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as array expression.
// Does not add sql cast to generated sql builder output.
func ArrayExp(expression Expression) ArrayExpression {
	return newArrayExpressionWrap(expression)
}

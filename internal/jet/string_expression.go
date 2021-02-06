package jet

// StringExpression interface
type StringExpression interface {
	Expression

	EQ(rhs StringExpression) BoolExpression
	NOT_EQ(rhs StringExpression) BoolExpression
	IS_DISTINCT_FROM(rhs StringExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs StringExpression) BoolExpression

	LT(rhs StringExpression) BoolExpression
	LT_EQ(rhs StringExpression) BoolExpression
	GT(rhs StringExpression) BoolExpression
	GT_EQ(rhs StringExpression) BoolExpression

	CONCAT(rhs Expression) StringExpression

	LIKE(pattern StringExpression) BoolExpression
	NOT_LIKE(pattern StringExpression) BoolExpression

	REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression
	NOT_REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression
}

type stringInterfaceImpl struct {
	parent StringExpression
}

func (s *stringInterfaceImpl) EQ(rhs StringExpression) BoolExpression {
	return Eq(s.parent, rhs)
}

func (s *stringInterfaceImpl) NOT_EQ(rhs StringExpression) BoolExpression {
	return NotEq(s.parent, rhs)
}

func (s *stringInterfaceImpl) IS_DISTINCT_FROM(rhs StringExpression) BoolExpression {
	return IsDistinctFrom(s.parent, rhs)
}

func (s *stringInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs StringExpression) BoolExpression {
	return IsNotDistinctFrom(s.parent, rhs)
}

func (s *stringInterfaceImpl) GT(rhs StringExpression) BoolExpression {
	return Gt(s.parent, rhs)
}

func (s *stringInterfaceImpl) GT_EQ(rhs StringExpression) BoolExpression {
	return GtEq(s.parent, rhs)
}

func (s *stringInterfaceImpl) LT(rhs StringExpression) BoolExpression {
	return Lt(s.parent, rhs)
}

func (s *stringInterfaceImpl) LT_EQ(rhs StringExpression) BoolExpression {
	return LtEq(s.parent, rhs)
}

func (s *stringInterfaceImpl) CONCAT(rhs Expression) StringExpression {
	return NewBinaryStringOperatorExpression(s.parent, rhs, StringConcatOperator)
}

func (s *stringInterfaceImpl) LIKE(pattern StringExpression) BoolExpression {
	return IsLike(s.parent, pattern)
}

func (s *stringInterfaceImpl) NOT_LIKE(pattern StringExpression) BoolExpression {
	return IsNotLike(s.parent, pattern)
}

func (s *stringInterfaceImpl) REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression {
	return RegExpLike(s.parent, pattern, caseSensitive...)
}

func (s *stringInterfaceImpl) NOT_REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression {
	return NotRegExpLike(s.parent, pattern, caseSensitive...)
}

//---------------------------------------------------//
func NewBinaryStringOperatorExpression(lhs, rhs Expression, operator string) StringExpression {
	return StringExp(NewBinaryOperatorExpression(lhs, rhs, operator))
}

//---------------------------------------------------//

type stringExpressionWrapper struct {
	stringInterfaceImpl
	Expression
}

func newStringExpressionWrap(expression Expression) StringExpression {
	stringExpressionWrap := stringExpressionWrapper{Expression: expression}
	stringExpressionWrap.stringInterfaceImpl.parent = &stringExpressionWrap
	return &stringExpressionWrap
}

// StringExp is string expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as string expression.
// Does not add sql cast to generated sql builder output.
func StringExp(expression Expression) StringExpression {
	return newStringExpressionWrap(expression)
}

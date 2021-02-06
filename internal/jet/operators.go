package jet

// Operators
const (
	StringConcatOperator        = "||"
	StringRegexpLikeOperator    = "REGEXP"
	StringNotRegexpLikeOperator = "NOT REGEXP"
)

//----------- Logical operators ---------------//

// NOT returns negation of bool expression result
func NOT(exp BoolExpression) BoolExpression {
	return NewPrefixBoolOperatorExpression(exp, "NOT")
}

// BIT_NOT inverts every bit in integer expression result
func BIT_NOT(expr IntegerExpression) IntegerExpression {
	if literalExp, ok := expr.(LiteralExpression); ok {
		literalExp.SetConstant(true)
	}
	return newPrefixIntegerOperatorExpression(expr, "~")
}

//----------- Comparison operators ---------------//

// EXISTS checks for existence of the rows in subQuery
func EXISTS(subQuery Expression) BoolExpression {
	return NewPrefixBoolOperatorExpression(subQuery, "EXISTS")
}

// Eq returns a representation of "a=b"
func Eq(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "=")
}

// NotEq returns a representation of "a!=b"
func NotEq(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "!=")
}

// IsDistinctFrom returns a representation of "a IS DISTINCT FROM b"
func IsDistinctFrom(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "IS DISTINCT FROM")
}

// IsNotDistinctFrom returns a representation of "a IS NOT DISTINCT FROM b"
func IsNotDistinctFrom(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "IS NOT DISTINCT FROM")
}

// IsLike returns a representation of "a LIKE b"
func IsLike(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "LIKE")
}

// IsNotLike returns a representation of "a NOT LIKE b"
func IsNotLike(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "NOT LIKE")
}

func RegExpLike(lhs, rhs Expression, caseSensitive ...bool) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, StringRegexpLikeOperator,
		Bool(len(caseSensitive) > 0 && caseSensitive[0]))
}

func NotRegExpLike(lhs, rhs Expression, caseSensitive ...bool) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, StringNotRegexpLikeOperator,
		Bool(len(caseSensitive) > 0 && caseSensitive[0]))
}

// Lt returns a representation of "a<b"
func Lt(lhs Expression, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "<")
}

// LtEq returns a representation of "a<=b"
func LtEq(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, "<=")
}

// Gt returns a representation of "a>b"
func Gt(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, ">")
}

// GtEq returns a representation of "a>=b"
func GtEq(lhs, rhs Expression) BoolExpression {
	return NewBinaryBoolOperatorExpression(lhs, rhs, ">=")
}

// Add notEq returns a representation of "a + b"
func Add(lhs, rhs Serializer) Expression {
	return NewBinaryOperatorExpression(lhs, rhs, "+")
}

// Sub notEq returns a representation of "a - b"
func Sub(lhs, rhs Serializer) Expression {
	return NewBinaryOperatorExpression(lhs, rhs, "-")
}

// Mul returns a representation of "a * b"
func Mul(lhs, rhs Serializer) Expression {
	return NewBinaryOperatorExpression(lhs, rhs, "*")
}

// Div returns a representation of "a / b"
func Div(lhs, rhs Serializer) Expression {
	return NewBinaryOperatorExpression(lhs, rhs, "/")
}

// Mod returns a representation of "a % b"
func Mod(lhs, rhs Serializer) Expression {
	return NewBinaryOperatorExpression(lhs, rhs, "%")
}

// --------------- CASE operator -------------------//

// CaseOperator is interface for SQL case operator
type CaseOperator interface {
	Expression

	WHEN(condition Expression) CaseOperator
	THEN(then Expression) CaseOperator
	ELSE(els Expression) CaseOperator
}

type caseOperatorImpl struct {
	expressionInterfaceImpl

	expression Expression
	when       []Expression
	then       []Expression
	els        Expression
}

// CASE create CASE operator with optional list of expressions
func CASE(expression ...Expression) CaseOperator {
	caseExp := &caseOperatorImpl{}

	if len(expression) > 0 {
		caseExp.expression = expression[0]
	}

	caseExp.expressionInterfaceImpl.parent = caseExp

	return caseExp
}

func (c *caseOperatorImpl) WHEN(when Expression) CaseOperator {
	c.when = append(c.when, when)
	return c
}

func (c *caseOperatorImpl) THEN(then Expression) CaseOperator {
	c.then = append(c.then, then)
	return c
}

func (c *caseOperatorImpl) ELSE(els Expression) CaseOperator {
	c.els = els

	return c
}

func (c *caseOperatorImpl) Serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	out.WriteString("(CASE")

	if c.expression != nil {
		c.expression.Serialize(statement, out, FallTrough(options)...)
	}

	if len(c.when) == 0 || len(c.then) == 0 {
		panic("jet: invalid case Statement. There should be at least one WHEN/THEN pair. ")
	}

	if len(c.when) != len(c.then) {
		panic("jet: WHEN and THEN expression count mismatch. ")
	}

	for i, when := range c.when {
		out.WriteString("WHEN")
		when.Serialize(statement, out, NoWrap)

		out.WriteString("THEN")
		c.then[i].Serialize(statement, out, NoWrap)
	}

	if c.els != nil {
		out.WriteString("ELSE")
		c.els.Serialize(statement, out, NoWrap)
	}

	out.WriteString("END)")
}

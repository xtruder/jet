package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// JSONExpression interface
type JSONExpression interface {
	Expression

	//----------------- Operators -----------------//

	// EXTRACT method extracts json value by key or by index
	EXTRACT(rhs Expression) JSONExpression

	// EXTRACT method extracts json value as text by key or by index
	EXTRACT_TEXT(rhs Expression) StringExpression

	// EXTRACT_PATH method extract json value by path
	EXTRACT_PATH(path ArrayExpression) JSONExpression

	// EXTRACT_PATH_TEXT method extracts text value by path
	EXTRACT_PATH_TEXT(path ArrayExpression) StringExpression

	//----------------- Helper functions -----------//

	// ARRAY_LEN gets len of json array
	ARRAY_LEN() IntegerExpression

	// TYPEOF returns the type of the outermost JSON value as a text string
	TYPEOF() StringExpression

	// STRIP_NULLS returns json with all object fields that have null values omitted
	STRIP_NULLS() JSONExpression
}

type jsonInterfaceImpl struct {
	parent JSONExpression
}

func (j *jsonInterfaceImpl) EXTRACT(rhs Expression) JSONExpression {
	return newBinaryJSONOperatorExpression(j.parent, rhs, "->")
}

func (j *jsonInterfaceImpl) EXTRACT_TEXT(rhs Expression) StringExpression {
	return jet.NewBinaryStringOperatorExpression(j.parent, rhs, "->>")
}

func (j *jsonInterfaceImpl) EXTRACT_PATH(path ArrayExpression) JSONExpression {
	return newBinaryJSONOperatorExpression(j.parent, path, "#>")
}

func (j *jsonInterfaceImpl) EXTRACT_PATH_TEXT(path ArrayExpression) StringExpression {
	return jet.NewBinaryStringOperatorExpression(j.parent, path, "#>>")
}

func (j *jsonInterfaceImpl) ARRAY_LEN() IntegerExpression {
	return JSON_ARRAY_LENGTH(j.parent)
}

func (j *jsonInterfaceImpl) TYPEOF() StringExpression {
	return JSON_TYPEOF(j.parent)
}

func (j *jsonInterfaceImpl) STRIP_NULLS() JSONExpression {
	return JSON_STRIP_NULLS(j.parent)
}

//---------------------------------------------------//
func newBinaryJSONOperatorExpression(lhs, rhs Expression, operator string) JSONExpression {
	return JSONExp(jet.NewBinaryOperatorExpression(lhs, rhs, operator))
}

//---------------------------------------------------//

type jsonExpressionWrapper struct {
	jsonInterfaceImpl
	Expression
}

func newJSONExpressionWrap(expression Expression) JSONExpression {
	jsonExpressionWrap := jsonExpressionWrapper{Expression: expression}
	jsonExpressionWrap.jsonInterfaceImpl.parent = &jsonExpressionWrap
	return &jsonExpressionWrap
}

// JSONExp is json expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as json expression.
// Does not add sql cast to generated sql builder output.
func JSONExp(expression Expression) JSONExpression {
	return newJSONExpressionWrap(expression)
}

package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// JSONExpression interface
type JSONBExpression interface {
	Expression

	//----------------- Operators -----------------//

	// EXTRACT method extracts json value by key or by index
	EXTRACT(rhs Expression) JSONBExpression

	// EXTRACT method extracts json value as text by key or by index
	EXTRACT_TEXT(rhs Expression) StringExpression

	// EXTRACT_PATH method extract json value by path
	EXTRACT_PATH(path ArrayExpression) JSONBExpression

	// EXTRACT_PATH_TEXT method extracts text value by path
	EXTRACT_PATH_TEXT(path ArrayExpression) StringExpression

	// CONTAINS method checks if json on righ hand is is contained in parent
	CONTAINS(rhs JSONBExpression) BoolExpression

	// IS_CONTAINED method checks if json is contained in json on righ hand side
	IS_CONTAINED_BY(rhs JSONBExpression) BoolExpression

	// HAS_KEY method checks if the string exists as top-level keys
	HAS_KEY(key StringExpression) BoolExpression

	// HAS_KEYS method checks whether any of these array strings exist as top-level keys
	HAS_KEYS(keys ArrayExpression) BoolExpression

	// HAS_ANY_KEY method checks if any of the keys exist on top level
	HAS_ANY_KEY(key ArrayExpression) BoolExpression

	// CONCAT concatenates two jsonb values into a new jsonb value
	CONCAT(json JSONBExpression) JSONBExpression

	// DELETE_KEY method deletes key/value pair or string element from json
	DELETE_KEY(key StringExpression) JSONBExpression

	// DELETE_INDEX method deletes the array element with specified index
	DELETE_INDEX(key IntegerExpression) JSONBExpression

	// DELETE method deletes the field or element with specified path
	DELETE(keys ArrayExpression) JSONBExpression

	//----------------- Helper functions -----------//

	// ARRAY_LEN gets len of jsonb array
	ARRAY_LEN() IntegerExpression

	// TYPEOF returns the type of the outermost jsonb value as a text string
	TYPEOF() StringExpression

	// STRIP_NULLS returns jsonb with all object fields that have null values omitted
	STRIP_NULLS() JSONBExpression
}

type jsonbInterfaceImpl struct {
	parent JSONBExpression
}

func (j *jsonbInterfaceImpl) EXTRACT(rhs Expression) JSONBExpression {
	return newBinaryJSONBOperatorExpression(j.parent, rhs, "->")
}

func (j *jsonbInterfaceImpl) EXTRACT_TEXT(rhs Expression) StringExpression {
	return jet.NewBinaryStringOperatorExpression(j.parent, rhs, "->>")
}

func (j *jsonbInterfaceImpl) EXTRACT_PATH(path ArrayExpression) JSONBExpression {
	return newBinaryJSONBOperatorExpression(j.parent, path, "#>")
}

func (j *jsonbInterfaceImpl) EXTRACT_PATH_TEXT(path ArrayExpression) StringExpression {
	return jet.NewBinaryStringOperatorExpression(j.parent, path, "#>>")
}

func (j *jsonbInterfaceImpl) CONTAINS(rhs JSONBExpression) BoolExpression {
	return CONTAINS(j.parent, rhs)
}

func (j *jsonbInterfaceImpl) IS_CONTAINED_BY(rhs JSONBExpression) BoolExpression {
	return IS_CONTAINED_BY(j.parent, rhs)
}

func (j *jsonbInterfaceImpl) HAS_KEY(key StringExpression) BoolExpression {
	return jet.NewBinaryBoolOperatorExpression(j.parent, key, "?")
}

func (j *jsonbInterfaceImpl) HAS_KEYS(keys ArrayExpression) BoolExpression {
	return jet.NewBinaryBoolOperatorExpression(j.parent, keys, "?&")
}

func (j *jsonbInterfaceImpl) HAS_ANY_KEY(key ArrayExpression) BoolExpression {
	return jet.NewBinaryBoolOperatorExpression(j.parent, key, "?|")
}

func (j *jsonbInterfaceImpl) CONCAT(json JSONBExpression) JSONBExpression {
	return newBinaryJSONBOperatorExpression(j.parent, json, "||")
}

func (j *jsonbInterfaceImpl) DELETE_KEY(key StringExpression) JSONBExpression {
	return newBinaryJSONBOperatorExpression(j.parent, key, "-")
}

func (j *jsonbInterfaceImpl) DELETE_INDEX(key IntegerExpression) JSONBExpression {
	return newBinaryJSONBOperatorExpression(j.parent, key, "-")
}

func (j *jsonbInterfaceImpl) DELETE(keys ArrayExpression) JSONBExpression {
	return newBinaryJSONBOperatorExpression(j.parent, keys, "#-")
}

func (j *jsonbInterfaceImpl) ARRAY_LEN() IntegerExpression {
	return JSONB_ARRAY_LENGTH(j.parent)
}

func (j *jsonbInterfaceImpl) TYPEOF() StringExpression {
	return JSONB_TYPEOF(j.parent)
}

func (j *jsonbInterfaceImpl) STRIP_NULLS() JSONBExpression {
	return JSONB_STRIP_NULLS(j.parent)
}

//---------------------------------------------------//
func newBinaryJSONBOperatorExpression(lhs, rhs Expression, operator string) JSONBExpression {
	return JSONBExp(jet.NewBinaryOperatorExpression(lhs, rhs, operator))
}

//---------------------------------------------------//

type jsonbExpressionWrapper struct {
	jsonbInterfaceImpl
	Expression
}

func newJSONBExpressionWrap(expression Expression) JSONBExpression {
	jsonbExpressionWrap := jsonbExpressionWrapper{Expression: expression}
	jsonbExpressionWrap.jsonbInterfaceImpl.parent = &jsonbExpressionWrap
	return &jsonbExpressionWrap
}

// JSONBExp is jsonb expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as json expression.
// Does not add sql cast to generated sql builder output.
func JSONBExp(expression Expression) JSONBExpression {
	return newJSONBExpressionWrap(expression)
}

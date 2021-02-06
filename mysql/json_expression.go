package mysql

import "github.com/go-jet/jet/v2/internal/jet"

// JSONExpression interface
type JSONExpression interface {
	StringExpression

	//----------------- Operators -----------------//

	// EXTRACT method returns value from JSON column after evaluating path
	EXTRACT(path StringExpression, paths ...StringExpression) JSONExpression

	// EXTRACT_UNQUOTE method returns value from JSON column after evaluating path and unquoting the result
	EXTRACT_UNQUOTE(path StringExpression) Expression

	// UNQUOUTE method unquotes JSON value
	UNQUOTE() Expression

	// ARRAY_APPEND method appends values to the end of the indicated arrays within a JSON document
	ARRAY_APPEND(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression

	// ARRAY_INSERT method updates a JSON document, inserting into an array within the document
	ARRAY_INSERT(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression

	// CONTAINS indicates by whether a given candidate JSON document is contained within a target JSON document
	CONTAINS(rhs JSONExpression, path StringExpression) BoolExpression

	// CONTAINS_PATH indicate whether a JSON document contains data at a given path or paths
	CONTAINS_PATH(oneOrAll bool, path StringExpression, paths ...StringExpression) BoolExpression

	// DEPTH returns the maximum depth of a JSON document
	DEPTH() IntegerExpression

	// KEYS returns the keys from the top-level value of a JSON object as a JSON array, or,
	// if a path argument is given, the top-level keys from the selected path
	KEYS(path ...StringExpression) JSONExpression

	// LENGTH returns the length of a JSON document, or, if a path argument is given,
	// the length of the value within the document identified by the path
	LENGTH(path ...StringExpression) IntegerExpression

	// MERGE_PATCH performs an RFC 7396 compliant merge of two or more JSON documents and returns the merged result
	MERGE_PATCH(rhs JSONExpression) JSONExpression

	// MERGE_PRESERVE merges two or more JSON documents and returns the merged result.
	MERGE_PRESERVE(rhs JSONExpression) JSONExpression

	// OVERLAPS compares two JSON documents.
	// Returns true (1) if the two document have any key-value pairs or array elements in common
	OVERLAPS(rhs JSONExpression) BoolExpression

	// REMOVE method removes data from a JSON document and returns the result
	REMOVE(path StringExpression, paths ...StringExpression) JSONExpression

	// REPLACE method replaces existing values in a JSON document and returns the result
	REPLACE(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression

	// SCHEMA_VALID method validates a JSON document against a JSON schema
	SCHEMA_VALID(schema JSONExpression) BoolExpression

	// SCHEMA_VALIDATION_REPORT method validates a JSON document against a JSON schema and returns
	// validation report as json document
	SCHEMA_VALIDATION_REPORT(schema JSONExpression) JSONExpression

	// SEARCH method returns the path to the given string within a JSON document
	SEARCH(oneOrAll bool, searchStr StringExpression, path ...StringExpression) JSONExpression

	// SET method inserts or updates data in a JSON document and returns the result.
	SET_VALUE(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression

	// TYPE method returns string indicating the type of a JSON value. This can be an object, an array, or a scalar type
	TYPE() StringExpression
}

//---------------------------------------------------//
func newBinaryJSONOperatorExpression(lhs, rhs Expression, operator string) JSONExpression {
	return JSONExp(jet.NewBinaryOperatorExpression(lhs, rhs, operator))
}

//---------------------------------------------------//

type jsonInterfaceImpl struct {
	parent JSONExpression
}

func (j *jsonInterfaceImpl) EQ(rhs StringExpression) BoolExpression {
	return jet.Eq(j.parent, optionalCastToJSON(rhs))
}

func (j *jsonInterfaceImpl) NOT_EQ(rhs StringExpression) BoolExpression {
	return jet.NotEq(j.parent, optionalCastToJSON(rhs))
}

func (j *jsonInterfaceImpl) IS_DISTINCT_FROM(rhs StringExpression) BoolExpression {
	return jet.IsDistinctFrom(j.parent, rhs)
}

func (j *jsonInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs StringExpression) BoolExpression {
	return jet.IsNotDistinctFrom(j.parent, rhs)
}

func (j *jsonInterfaceImpl) LT(rhs StringExpression) BoolExpression {
	return jet.Lt(j.parent, optionalCastToJSON(rhs))
}

func (j *jsonInterfaceImpl) LT_EQ(rhs StringExpression) BoolExpression {
	return jet.LtEq(j.parent, optionalCastToJSON(rhs))
}

func (j *jsonInterfaceImpl) GT(rhs StringExpression) BoolExpression {
	return jet.Gt(j.parent, optionalCastToJSON(rhs))
}

func (j *jsonInterfaceImpl) GT_EQ(rhs StringExpression) BoolExpression {
	return jet.GtEq(j.parent, optionalCastToJSON(rhs))
}

func (j *jsonInterfaceImpl) CONCAT(rhs Expression) StringExpression {
	return jet.NewBinaryStringOperatorExpression(j.parent, rhs, jet.StringConcatOperator)
}

func (j *jsonInterfaceImpl) LIKE(pattern StringExpression) BoolExpression {
	return jet.IsLike(j.parent, pattern)
}

func (j *jsonInterfaceImpl) NOT_LIKE(pattern StringExpression) BoolExpression {
	return jet.IsNotLike(j.parent, pattern)
}

func (j *jsonInterfaceImpl) REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression {
	return jet.RegExpLike(j.parent, pattern, caseSensitive...)
}

func (j *jsonInterfaceImpl) NOT_REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression {
	return jet.NotRegExpLike(j.parent, pattern, caseSensitive...)
}

func (j *jsonInterfaceImpl) EXTRACT(path StringExpression, paths ...StringExpression) JSONExpression {
	if len(paths) == 0 {
		return newBinaryJSONOperatorExpression(j.parent, path, "->")
	}

	allPaths := []StringExpression{path}
	allPaths = append(allPaths, paths...)

	return JSON_EXTRACT(j.parent, allPaths...)
}

func (j *jsonInterfaceImpl) EXTRACT_UNQUOTE(path StringExpression) Expression {
	return newBinaryJSONOperatorExpression(j.parent, path, "->>")
}

func (j *jsonInterfaceImpl) UNQUOTE() Expression {
	return JSON_UNQUOTE(j.parent)
}

func (j *jsonInterfaceImpl) ARRAY_APPEND(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression {
	return JSON_ARRAY_APPEND(j.parent, path, value, pathOrValue...)
}

func (j *jsonInterfaceImpl) ARRAY_INSERT(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression {
	return JSON_ARRAY_INSERT(j.parent, path, value, pathOrValue...)
}

func (j *jsonInterfaceImpl) CONTAINS(rhs JSONExpression, path StringExpression) BoolExpression {
	return JSON_CONTAINS(j.parent, rhs, path)
}

func (j *jsonInterfaceImpl) CONTAINS_PATH(oneOrAll bool, path StringExpression, paths ...StringExpression) BoolExpression {
	return JSON_CONTAINS_PATH(j.parent, oneOrAll, path, paths...)
}

func (j *jsonInterfaceImpl) DEPTH() IntegerExpression {
	return JSON_DEPTH(j.parent)
}

func (j *jsonInterfaceImpl) KEYS(paths ...StringExpression) JSONExpression {
	return JSON_KEYS(j.parent, paths...)
}

func (j *jsonInterfaceImpl) LENGTH(paths ...StringExpression) IntegerExpression {
	return JSON_LENGTH(j.parent, paths...)
}

func (j *jsonInterfaceImpl) MERGE_PATCH(rhs JSONExpression) JSONExpression {
	return JSON_MERGE_PATCH(j.parent, rhs)
}

func (j *jsonInterfaceImpl) MERGE_PRESERVE(rhs JSONExpression) JSONExpression {
	return JSON_MERGE_PRESERVE(j.parent, rhs)
}

func (j *jsonInterfaceImpl) OVERLAPS(rhs JSONExpression) BoolExpression {
	return JSON_OVERLAPS(j.parent, rhs)
}

func (j *jsonInterfaceImpl) REMOVE(path StringExpression, paths ...StringExpression) JSONExpression {
	return JSON_REMOVE(j.parent, path, paths...)
}

func (j *jsonInterfaceImpl) REPLACE(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression {
	return JSON_REPLACE(j.parent, path, value, pathOrValue...)
}

func (j *jsonInterfaceImpl) SCHEMA_VALID(schema JSONExpression) BoolExpression {
	return JSON_SCHEMA_VALID(j.parent, schema)
}

func (j *jsonInterfaceImpl) SCHEMA_VALIDATION_REPORT(schema JSONExpression) JSONExpression {
	return JSON_SCHEMA_VALIDATION_REPORT(j.parent, schema)
}

func (j *jsonInterfaceImpl) SEARCH(oneOrAll bool, searchStr StringExpression, paths ...StringExpression) JSONExpression {
	return JSON_SEARCH(j.parent, oneOrAll, searchStr, jet.String(""), paths...)
}

func (j *jsonInterfaceImpl) SET_VALUE(path StringExpression, value Expression, pathOrValue ...Expression) JSONExpression {
	return JSON_SET(j.parent, path, value, pathOrValue...)
}

func (j *jsonInterfaceImpl) TYPE() StringExpression {
	return JSON_TYPE(j.parent)
}

//---------------------------------------------------//

func optionalCastToJSON(expression Expression) Expression {
	if _, ok := expression.(JSONExpression); ok {
		return expression
	}

	return CAST(expression).AS_JSON()
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

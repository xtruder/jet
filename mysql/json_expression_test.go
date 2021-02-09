package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestJSONExpression_EXTRACT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single path",
			Test: JSONExp(table2ColJSON).EXTRACT(String("$.id"))},
		{Name: "multiple path",
			Test: JSONExp(table2ColJSON).EXTRACT(String("$[0]"), String("$[1]"))},
	}.Run(t, Dialect)
}

func TestJSONExpression_EXTRACT_UNQUOTE(t *testing.T) {
	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).EXTRACT_UNQUOTE(String("$.id")),
	}.Assert(t, Dialect)
}

func TestJSONExpression_UNQUOTE(t *testing.T) {
	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).UNQUOTE(),
	}.Assert(t, Dialect)
}

func TestJSONExpression_ARRAY_APPEND(t *testing.T) {
	val := []string{"val1", "val2"}

	testutils.SerializerTests{
		{Name: "single",
			Test: JSONExp(table2ColJSON).ARRAY_APPEND(String("$.id"), JSON(val))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).ARRAY_APPEND(String("$.id"), JSON(val), String("$.key"), JSON(val))},
	}.Run(t, Dialect)
}

func TestJSONExpression_ARRAY_INSERT(t *testing.T) {
	val := []string{"val1", "val2"}

	testutils.SerializerTests{
		{Name: "single",
			Test: JSONExp(table2ColJSON).ARRAY_INSERT(String("$.id"), JSON(val))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).ARRAY_INSERT(String("$.id"), JSON(val), String("$.key"), JSON(val))},
	}.Run(t, Dialect)
}

func TestJSONExpression_CONTAINS(t *testing.T) {
	val := []string{"val1", "val2"}

	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).CONTAINS(JSON(val), String("$.id")),
	}.Assert(t, Dialect)
}

func TestJSONExpression_CONTAINS_PATH(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single",
			Test: JSONExp(table2ColJSON).CONTAINS_PATH(false, String("$.id"))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).CONTAINS_PATH(true, String("$.key1"), String("$.key2"))},
	}.Run(t, Dialect)
}

func TestJSONExpression_DEPTH(t *testing.T) {
	testutils.SerializerTest{Test: JSONExp(table2ColJSON).DEPTH()}.Assert(t, Dialect)
}

func TestJSONExpression_KEYS(t *testing.T) {
	testutils.SerializerTests{
		{Name: "no params", Test: JSONExp(table2ColJSON).KEYS()},
		{Name: "with path", Test: JSONExp(table2ColJSON).KEYS(String("$.id"))},
	}.Run(t, Dialect)
}

func TestJSONExpression_LENGTH(t *testing.T) {
	testutils.SerializerTests{
		{Name: "no params", Test: JSONExp(table2ColJSON).LENGTH()},
		{Name: "with path", Test: JSONExp(table2ColJSON).LENGTH(String("$.id"))},
	}.Run(t, Dialect)
}

func TestJSONExpression_MERGE_PATCH(t *testing.T) {
	val := []string{"val1", "val2"}

	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).MERGE_PATCH(JSON(val)),
	}.Assert(t, Dialect)
}

func TestJSONExpression_MERGE_PRESERVE(t *testing.T) {
	val := []string{"val1", "val2"}

	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).MERGE_PRESERVE(JSON(val)),
	}.Assert(t, Dialect)
}

func TestJSONExpression_OVERLAPS(t *testing.T) {
	val := []string{"val1", "val2"}

	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).OVERLAPS(JSON(val)),
	}.Assert(t, Dialect)
}

func TestJSONExpression_REMOVE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single path", Test: JSONExp(table2ColJSON).REMOVE(String("$.id"))},
		{Name: "multiple path", Test: JSONExp(table2ColJSON).REMOVE(String("$.key1"), String("$.key2"))},
	}.Run(t, Dialect)
}

func TestJSONExpression_REPLACE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: JSONExp(table2ColJSON).REPLACE(String("$.id"), String("value"))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).REPLACE(String("$.key1"), String("value"), String("$.key2"), Int(2))},
	}.Run(t, Dialect)
}

func TestJSONExpression_SCHEMA_VALID(t *testing.T) {
	schema := map[string]string{"type": "string"}

	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).SCHEMA_VALID(JSON(schema)),
	}.Assert(t, Dialect)
}

func TestJSONExpression_SCHEMA_VALIDATION_REPORT(t *testing.T) {
	schema := map[string]string{"type": "string"}

	testutils.SerializerTest{
		Test: JSONExp(table2ColJSON).SCHEMA_VALIDATION_REPORT(JSON(schema)),
	}.Assert(t, Dialect)
}

func TestJSONExpression_SEARCH(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: JSONExp(table2ColJSON).SEARCH(false, String("abc"))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).SEARCH(true, String("10"), String("$[1]"), String("$[2]"))},
	}.Run(t, Dialect)
}

func TestJSONExpression_SET_VALUE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single",
			Test: JSONExp(table2ColJSON).SET_VALUE(String("$.id"), String("abc"))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).SET_VALUE(String("$.key1"), String("value"), String("$.key2"), Int(2))},
	}.Run(t, Dialect)
}

func TestJSONExpression_TYPE(t *testing.T) {
	testutils.SerializerTest{Test: JSONExp(table2ColJSON).TYPE()}.Assert(t, Dialect)
}

func TestJSONExpression_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: JSONExp(table2ColJSON).EQ(String("value"))},
		{Name: "multiple",
			Test: JSONExp(table2ColJSON).EQ(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_NOT_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).NOT_EQ(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).NOT_EQ(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_LT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).LT(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).LT(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_LT_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).LT_EQ(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).LT_EQ(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_GT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).GT(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).GT(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_GT_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).GT_EQ(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).GT_EQ(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_IS_DISTINCT_FROM(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).IS_DISTINCT_FROM(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).IS_DISTINCT_FROM(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_IS_NOT_DISTINCT_FROM(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).IS_NOT_DISTINCT_FROM(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).IS_NOT_DISTINCT_FROM(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_CONCAT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).CONCAT(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).CONCAT(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_LIKE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).LIKE(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).LIKE(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_NOT_LIKE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).NOT_LIKE(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).NOT_LIKE(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_REGEXP_LIKE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).REGEXP_LIKE(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).REGEXP_LIKE(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

func TestJSONExpression_NOT_REGEXP_LIKE(t *testing.T) {
	testutils.SerializerTests{
		{Name: "string", Test: JSONExp(table2ColJSON).NOT_REGEXP_LIKE(String("value"))},
		{Name: "object",
			Test: JSONExp(table2ColJSON).NOT_REGEXP_LIKE(JSON(map[string]string{"key": "value"}))},
	}.Run(t, Dialect)
}

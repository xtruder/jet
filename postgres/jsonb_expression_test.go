package postgres

import (
	"testing"

	"github.com/lib/pq"
)

func TestJSONBExpression_EXTRACT(t *testing.T) {
	expr := JSONBExp(table2ColJSONB).EXTRACT(String("p1"))
	assertSerialize(t, expr, "(table2.col_jsonb -> $1)", "p1")
}

func TestJSONBExpression_EXTRACT_TEXT(t *testing.T) {
	expr := JSONBExp(table2ColJSONB).EXTRACT_TEXT(String("p1"))
	assertSerialize(t, expr, "(table2.col_jsonb ->> $1)", "p1")
}

func TestJSONBExpression_EXTRACT_PATH(t *testing.T) {
	path := []string{"p1", "p2"}
	expr := JSONBExp(table2ColJSONB).EXTRACT_PATH(newArrayLiteral(path))
	assertSerialize(t, expr, "(table2.col_jsonb #> $1)", pq.Array(path))
}

func TestJSONBExpression_EXTRACT_PATH_TEXT(t *testing.T) {
	path := []string{"p1", "p2"}
	expr := JSONBExp(table2ColJSONB).EXTRACT_PATH_TEXT(newArrayLiteral(path))
	assertSerialize(t, expr, "(table2.col_jsonb #>> $1)", pq.Array(path))
}

func TestJSONBExpression_CONTAINS(t *testing.T) {
	value := map[string]string{"key": "value"}
	expr := JSONBExp(table2ColJSONB).CONTAINS(newJsonbLiteral(value))
	assertSerialize(t, expr, "(table2.col_jsonb @> $1)", `{"key":"value"}`)
}

func TestJSONBExpression_IS_CONTAINED_BY(t *testing.T) {
	value := map[string]string{"key": "value"}
	expr := JSONBExp(table2ColJSONB).IS_CONTAINED_BY(newJsonbLiteral(value))
	assertSerialize(t, expr, "(table2.col_jsonb <@ $1)", `{"key":"value"}`)
}

func TestJSONBExpression_DELETE_KEY(t *testing.T) {
	expr := JSONBExp(table2ColJSONB).DELETE_KEY(String("key"))
	assertSerialize(t, expr, "(table2.col_jsonb - $1)", "key")
}

func TestJSONBExpression_DELETE_INDEX(t *testing.T) {
	var val int64 = 1
	expr := JSONBExp(table2ColJSONB).DELETE_INDEX(Int(val))
	assertSerialize(t, expr, "(table2.col_jsonb - $1)", val)
}

func TestJSONBExpression_DELETE(t *testing.T) {
	expr := JSONBExp(table2ColJSONB).DELETE(newArrayLiteral([]string{"key"}))
	assertSerialize(t, expr, "(table2.col_jsonb #- $1)", pq.Array([]string{"key"}))
}

func TestJSONBExpression_HAS_KEY(t *testing.T) {
	expr := JSONBExp(table2ColJSONB).HAS_KEY(String("key"))
	assertSerialize(t, expr, "(table2.col_jsonb ? $1)", "key")
}

func TestJSONBExpression_HAS_ANY_KEY(t *testing.T) {
	keys := []string{"key1", "key2"}
	expr := JSONBExp(table2ColJSONB).HAS_ANY_KEY(newArrayLiteral(keys))
	assertSerialize(t, expr, "(table2.col_jsonb ?| $1)", pq.Array(keys))
}

func TestJSONBExpression_HAS_KEYS(t *testing.T) {
	keys := []string{"key1", "key2"}
	expr := JSONBExp(table2ColJSONB).HAS_KEYS(newArrayLiteral(keys))
	assertSerialize(t, expr, "(table2.col_jsonb ?& $1)", pq.Array(keys))
}

func TestJSONBExpression_CONCAT(t *testing.T) {
	expr := JSONBExp(table1ColJSONB).CONCAT(table2ColJSONB)
	assertSerialize(t, expr, "(table1.col_jsonb || table2.col_jsonb)")
}

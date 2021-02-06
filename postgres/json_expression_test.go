package postgres

import (
	"testing"

	"github.com/lib/pq"
)

func TestJSONExpression_EXTRACT(t *testing.T) {
	expr := JSONExp(table2ColJSON).EXTRACT(String("p1"))
	assertSerialize(t, expr, "(table2.col_json -> $1)", "p1")
}

func TestJSONExpression_EXTRACT_TEXT(t *testing.T) {
	expr := JSONExp(table2ColJSON).EXTRACT_TEXT(String("p1"))
	assertSerialize(t, expr, "(table2.col_json ->> $1)", "p1")
}

func TestJSONExpression_EXTRACT_PATH(t *testing.T) {
	path := []string{"p1", "p2"}
	expr := JSONExp(table2ColJSON).EXTRACT_PATH(newArrayLiteral(path))
	assertSerialize(t, expr, "(table2.col_json #> $1)", pq.Array(path))
}

func TestJSONExpression_EXTRACT_PATH_TEXT(t *testing.T) {
	path := []string{"p1", "p2"}
	expr := JSONExp(table2ColJSON).EXTRACT_PATH_TEXT(newArrayLiteral(path))
	assertSerialize(t, expr, "(table2.col_json #>> $1)", pq.Array(path))
}

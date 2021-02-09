package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestString_REGEXP_LIKE_operator(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table3StrCol.REGEXP_LIKE(table2ColStr)},
		{Name: "literal", Test: table3StrCol.REGEXP_LIKE(String("JOHN"))},
		{Name: "literal not case sensitive", Test: table3StrCol.REGEXP_LIKE(String("JOHN"), false)},
		{Name: "literal case sensitive", Test: table3StrCol.REGEXP_LIKE(String("JOHN"), true)},
	}.Run(t, Dialect)
}

func TestString_NOT_REGEXP_LIKE_operator(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table3StrCol.NOT_REGEXP_LIKE(table2ColStr)},
		{Name: "literal", Test: table3StrCol.NOT_REGEXP_LIKE(String("JOHN"))},
		{Name: "literal not case sensitive", Test: table3StrCol.NOT_REGEXP_LIKE(String("JOHN"), false)},
		{Name: "literal case sensitive", Test: table3StrCol.NOT_REGEXP_LIKE(String("JOHN"), true)},
	}.Run(t, Dialect)
}

func TestExists(t *testing.T) {
	testutils.SerializerTests{
		{Name: "select", Test: EXISTS(
			table2.
				SELECT(Int(1)).
				WHERE(table1Col1.EQ(table2Col3)),
		)},
		{Name: "condition", Test: EXISTS(
			SELECT(Int(1)),
		).EQ(Bool(true))},
	}.Run(t, Dialect)

	testutils.ProjectionTests{
		{Name: "projection", Test: EXISTS(SELECT(Int(1))).AS("exists")},
	}.Run(t, Dialect)
}

func TestIN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single col", Test: Float(1.11).IN(table1.SELECT(table1Col1))},
		{Name: "multiple col", Test: ROW(Int(12), table1Col1).IN(table2.SELECT(table2Col3, table3Col1))},
	}.Run(t, Dialect)
}

func TestNOT_IN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single col", Test: Float(1.11).NOT_IN(table1.SELECT(table1Col1))},
		{Name: "multiple col", Test: ROW(Int(12), table1Col1).NOT_IN(table2.SELECT(table2Col3, table3Col1))},
	}.Run(t, Dialect)
}

func TestReservedWordEscaped(t *testing.T) {
	table1ColUser := IntervalColumn("user")
	table1ColVariadic := IntervalColumn("VARIADIC")
	table1ColProcedure := IntervalColumn("procedure")

	_ = NewTable(
		"db",
		"table1",
		table1ColUser,
		table1ColVariadic,
		table1ColProcedure,
	)

	testutils.SerializerTests{
		{Name: "word user", Test: table1ColUser},
		{Name: "word variadic", Test: table1ColVariadic},
		{Name: "word procedure", Test: table1ColProcedure},
	}.Run(t, Dialect)
}

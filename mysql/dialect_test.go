package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestDialectBoolExpressionIS_DISTINCT_FROM(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColBool.IS_DISTINCT_FROM(table2ColBool)},
		{Name: "literal", Test: table1ColBool.IS_DISTINCT_FROM(Bool(false))},
	}.Run(t, Dialect)
}

func TestDialectBoolExpressionIS_NOT_DISTINCT_FROM(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColBool.IS_NOT_DISTINCT_FROM(table2ColBool)},
		{Name: "literal", Test: table1ColBool.IS_NOT_DISTINCT_FROM(Bool(false))},
	}.Run(t, Dialect)
}

func TestDialectBoolLiteral(t *testing.T) {
	testutils.SerializerTests{
		{Name: "true", Test: Bool(true)},
		{Name: "false", Test: Bool(false)},
	}.Run(t, Dialect)
}

func TestDialectIntegerExpressionDIV(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColInt.DIV(table2ColInt)},
		{Name: "literal", Test: table1ColInt.DIV(Int(11))},
	}.Run(t, Dialect)
}

func TestDialectIntExpressionPOW(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColInt.POW(table2ColInt)},
		{Name: "literal", Test: table1ColInt.POW(Int(11))},
	}.Run(t, Dialect)
}

func TestDialectIntExpressionBIT_XOR(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColInt.BIT_XOR(table2ColInt)},
		{Name: "literal", Test: table1ColInt.BIT_XOR(Int(11))},
	}.Run(t, Dialect)
}

func TestDialectSelectExists(t *testing.T) {
	testutils.SerializerTest{Test: table2.
		SELECT(Int(1)).
		WHERE(table1Col1.EQ(table2Col3)),
	}.Assert(t, Dialect)
}

func TestDialectString_REGEXP_LIKE_operator(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table3StrCol.REGEXP_LIKE(table2ColStr)},
		{Name: "literal", Test: table3StrCol.REGEXP_LIKE(String("JOHN"))},
		{Name: "literal case sensitive", Test: table3StrCol.REGEXP_LIKE(String("JOHN"), true)},
		{Name: "literal not sensitive", Test: table3StrCol.REGEXP_LIKE(String("JOHN"), false)},
	}.Run(t, Dialect)
}

func TestDialectString_NOT_REGEXP_LIKE_operator(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table3StrCol.NOT_REGEXP_LIKE(table2ColStr)},
		{Name: "literal", Test: table3StrCol.NOT_REGEXP_LIKE(String("JOHN"))},
		{Name: "literal case sensitive", Test: table3StrCol.NOT_REGEXP_LIKE(String("JOHN"), true)},
		{Name: "literal not case sensitive", Test: table3StrCol.NOT_REGEXP_LIKE(String("JOHN"), false)},
	}.Run(t, Dialect)
}

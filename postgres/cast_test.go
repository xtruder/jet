package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestExpressionCAST_AS(t *testing.T) {
	testutils.SerializerTest{Test: CAST(String("test")).AS("text")}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_BOOL(t *testing.T) {
	testutils.SerializerTests{
		{Name: "literal", Test: CAST(Int(1)).AS_BOOL()},
		{Name: "column", Test: CAST(table2Col3).AS_BOOL()},
		{Name: "expression", Test: CAST(table2Col3.ADD(table2Col3)).AS_BOOL()},
	}.Run(t, Dialect)
}

func TestExpressionCAST_AS_SMALLINT(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_SMALLINT()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_INTEGER(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_INTEGER()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_BIGINT(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_BIGINT()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_NUMERIC(t *testing.T) {
	testutils.SerializerTests{
		{Name: "precision", Test: CAST(table2Col3).AS_NUMERIC(11)},
		{Name: "precision and scale", Test: CAST(table2Col3).AS_NUMERIC(11, 11)},
	}.Run(t, Dialect)
}

func TestExpressionCAST_AS_REAL(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_REAL()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_DOUBLE(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_DOUBLE()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_TEXT(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_TEXT()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_DATE(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_DATE()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_TIME(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_TIME()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_TIMEZ(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_TIMEZ()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_TIMESTAMP(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_TIMESTAMP()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_TIMESTAMPZ(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2Col3).AS_TIMESTAMPZ()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_INTERVAL(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: CAST(table2ColTimez).AS_INTERVAL()},
		{Name: "literal", Test: CAST(Time(20, 11, 10)).AS_INTERVAL()},
		{Name: "expression", Test: table2ColDate.SUB(CAST(Time(20, 11, 10)).AS_INTERVAL())},
	}.Run(t, Dialect)
}

func TestExpressionCAST_AS_ARRAY(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2ColStr).AS_ARRAY("text")}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_JSON(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2ColStr).AS_JSON()}.Assert(t, Dialect)
}

func TestExpressionCAST_AS_JSONB(t *testing.T) {
	testutils.SerializerTest{Test: CAST(table2ColStr).AS_JSONB()}.Assert(t, Dialect)
}

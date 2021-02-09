package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestArrayEQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.EQ(table2ColArray)},
		{Name: "literal", Test: table2ColArray.EQ(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayNOT_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.NOT_EQ(table2ColArray)},
		{Name: "literal", Test: table2ColArray.NOT_EQ(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayGT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.GT(table2ColArray)},
		{Name: "literal", Test: table2ColArray.GT(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayGT_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.GT_EQ(table2ColArray)},
		{Name: "literal", Test: table2ColArray.GT_EQ(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayLT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.LT(table2ColArray)},
		{Name: "literal", Test: table2ColArray.LT(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayLT_EQ(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.LT_EQ(table2ColArray)},
		{Name: "literal", Test: table2ColArray.LT_EQ(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayCONTAINS(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.CONTAINS(table2ColArray)},
		{Name: "literal", Test: table2ColArray.CONTAINS(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayIS_CONTAINED_BY(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.IS_CONTAINED_BY(table2ColArray)},
		{Name: "literal", Test: table2ColArray.IS_CONTAINED_BY(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayOVERLAPS(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.OVERLAPS(table2ColArray)},
		{Name: "literal", Test: table2ColArray.OVERLAPS(newArrayLiteral([]string{"JOHN"}))},
	}.Run(t, Dialect)
}

func TestArrayCONCAT(t *testing.T) {
	testutils.SerializerTests{
		{Name: "column", Test: table1ColArray.CONCAT(table2ColArray)},
		{Name: "literal", Test: table2ColArray.CONCAT(String("v"))},
	}.Run(t, Dialect)
}

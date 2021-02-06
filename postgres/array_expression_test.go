package postgres

import (
	"testing"

	"github.com/lib/pq"
)

func TestArrayEQ(t *testing.T) {
	exp := table1ColArray.EQ(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array = table2.col_array)")
	exp = table2ColArray.EQ(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array = $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayNOT_EQ(t *testing.T) {
	exp := table1ColArray.NOT_EQ(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array != table2.col_array)")
	exp = table2ColArray.NOT_EQ(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array != $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayGT(t *testing.T) {
	exp := table1ColArray.GT(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array > table2.col_array)")
	exp = table2ColArray.GT(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array > $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayGT_EQ(t *testing.T) {
	exp := table1ColArray.GT_EQ(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array >= table2.col_array)")
	exp = table2ColArray.GT_EQ(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array >= $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayLT(t *testing.T) {
	exp := table1ColArray.LT(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array < table2.col_array)")
	exp = table2ColArray.LT(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array < $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayLT_EQ(t *testing.T) {
	exp := table1ColArray.LT_EQ(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array <= table2.col_array)")
	exp = table2ColArray.LT_EQ(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array <= $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayCONTAINS(t *testing.T) {
	exp := table1ColArray.CONTAINS(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array @> table2.col_array)")
	exp = table2ColArray.CONTAINS(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array @> $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayIS_CONTAINED_BY(t *testing.T) {
	exp := table1ColArray.IS_CONTAINED_BY(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array <@ table2.col_array)")
	exp = table2ColArray.IS_CONTAINED_BY(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array <@ $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayOVERLAPS(t *testing.T) {
	exp := table1ColArray.OVERLAPS(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array && table2.col_array)")
	exp = table2ColArray.OVERLAPS(newArrayLiteral([]string{"JOHN"}))
	assertSerialize(t, exp, "(table2.col_array && $1)", pq.Array([]string{"JOHN"}))
}

func TestArrayCONCAT(t *testing.T) {
	exp := table1ColArray.CONCAT(table2ColArray)
	assertSerialize(t, exp, "(table1.col_array || table2.col_array)")
	exp = table2ColArray.CONCAT(String("v"))
	assertSerialize(t, exp, "(table2.col_array || $1)", "v")
}

package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestJoinNilInputs(t *testing.T) {
	testutils.SerializerTests{
		{
			Name:   "panics right side nil",
			Test:   table2.INNER_JOIN(nil, table1ColBool.EQ(table2ColBool)),
			Panics: "jet: right hand side of join operation is nil table",
		},
		{
			Name:   "panics nil join condition",
			Test:   table2.INNER_JOIN(table1, nil),
			Panics: "jet: join condition is nil",
		},
	}.Run(t, Dialect)
}

func TestINNER_JOIN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: table1.
			INNER_JOIN(table2, table1ColInt.EQ(table2ColInt))},
		{Name: "multiple col", Test: table1.
			INNER_JOIN(table2, table1ColInt.EQ(table2ColInt)).
			INNER_JOIN(table3, table1ColInt.EQ(table3ColInt))},
		{Name: "multiple lit", Test: table1.
			INNER_JOIN(table2, table1ColInt.EQ(Int(1))).
			INNER_JOIN(table3, table1ColInt.EQ(Int(2)))},
	}.Run(t, Dialect)
}

func TestLEFT_JOIN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: table1.
			LEFT_JOIN(table2, table1ColInt.EQ(table2ColInt))},
		{Name: "multiple col", Test: table1.
			LEFT_JOIN(table2, table1ColInt.EQ(table2ColInt)).
			LEFT_JOIN(table3, table1ColInt.EQ(table3ColInt))},
		{Name: "multiple lit", Test: table1.
			LEFT_JOIN(table2, table1ColInt.EQ(Int(1))).
			LEFT_JOIN(table3, table1ColInt.EQ(Int(2)))},
	}.Run(t, Dialect)
}

func TestRIGHT_JOIN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: table1.
			RIGHT_JOIN(table2, table1ColInt.EQ(table2ColInt))},
		{Name: "multiple col", Test: table1.
			RIGHT_JOIN(table2, table1ColInt.EQ(table2ColInt)).
			RIGHT_JOIN(table3, table1ColInt.EQ(table3ColInt))},
		{Name: "multiple lit", Test: table1.
			RIGHT_JOIN(table2, table1ColInt.EQ(Int(1))).
			RIGHT_JOIN(table3, table1ColInt.EQ(Int(2)))},
	}.Run(t, Dialect)
}

func TestFULL_JOIN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: table1.
			FULL_JOIN(table2, table1ColInt.EQ(table2ColInt))},
		{Name: "multiple col", Test: table1.
			FULL_JOIN(table2, table1ColInt.EQ(table2ColInt)).
			FULL_JOIN(table3, table1ColInt.EQ(table3ColInt))},
		{Name: "multiple lit", Test: table1.
			FULL_JOIN(table2, table1ColInt.EQ(Int(1))).
			FULL_JOIN(table3, table1ColInt.EQ(Int(2)))},
	}.Run(t, Dialect)
}

func TestCROSS_JOIN(t *testing.T) {
	testutils.SerializerTests{
		{Name: "single", Test: table1.CROSS_JOIN(table2)},
		{Name: "multiple tbls", Test: table1.
			CROSS_JOIN(table2).
			CROSS_JOIN(table3)},
	}.Run(t, Dialect)
}

package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestDeleteUnconditionally(t *testing.T) {
	testutils.StatementTests{
		{
			Name:   "panics without where",
			Test:   table1.DELETE(),
			Panics: `jet: WHERE clause not set`,
		},
		{
			Name:   "panics nil where",
			Test:   table1.DELETE().WHERE(nil),
			Panics: `jet: WHERE clause not set`,
		},
	}.Run(t)
}

func TestDeleteWithWhere(t *testing.T) {
	testutils.StatementTest{
		Test: table1.DELETE().WHERE(table1Col1.EQ(Int(1))),
	}.Assert(t)
}

func TestDeleteWithWhereOrderByLimit(t *testing.T) {
	testutils.StatementTest{
		Test: table1.DELETE().WHERE(table1Col1.EQ(Int(1))).ORDER_BY(table1Col1).LIMIT(1),
	}.Assert(t)
}

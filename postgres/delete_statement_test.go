package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestDeleteUnconditionally(t *testing.T) {
	testutils.StatementTests{
		{Name: "panics if where not set", Test: table1.DELETE(),
			Panics: `jet: WHERE clause not set`},
		{Name: "panics if where is nil", Test: table1.DELETE().WHERE(nil),
			Panics: `jet: WHERE clause not set`},
	}.Run(t)
}

func TestDeleteWithWhere(t *testing.T) {
	stmt := table1.DELETE().WHERE(table1Col1.EQ(Int(1)))
	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestDeleteWithWhereAndReturning(t *testing.T) {
	stmt := table1.DELETE().WHERE(table1Col1.EQ(Int(1))).RETURNING(table1Col1)
	testutils.StatementTest{Test: stmt}.Assert(t)
}

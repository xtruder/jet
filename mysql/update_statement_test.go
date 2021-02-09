package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestUpdateWithOneValue(t *testing.T) {
	stmt := table1.UPDATE(table1ColInt).
		SET(1).
		WHERE(table1ColInt.GT_EQ(Int(33)))

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestUpdateWithValues(t *testing.T) {
	stmt := table1.UPDATE(table1ColInt, table1ColFloat).
		SET(1, 22.2).
		WHERE(table1ColInt.GT_EQ(Int(33)))

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestUpdateOneColumnWithSelect(t *testing.T) {
	stmt := table1.
		UPDATE(table1ColFloat).
		SET(
			table1.SELECT(table1ColFloat),
		).
		WHERE(table1Col1.EQ(Int(2)))

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInvalidInputs(t *testing.T) {
	testutils.StatementTests{
		{
			Name:   "panics with no value",
			Test:   table1.UPDATE(table1ColInt).SET(1),
			Panics: "jet: WHERE clause not set",
		},
		{
			Name:   "panics with nil columns",
			Test:   table1.UPDATE(nil).SET(1),
			Panics: "jet: nil column in columns list for SET clause",
		},
	}.Run(t)
}

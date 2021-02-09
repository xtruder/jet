package postgres

import (
	"testing"
)

func TestUpdateWithOneValue(t *testing.T) {
	stmt := table1.UPDATE(table1ColInt).
		SET(1).
		WHERE(table1ColInt.GT_EQ(Int(33)))

	assertStatementRecordSQL(t, stmt)
}

func TestUpdateWithValues(t *testing.T) {
	stmt := table1.UPDATE(table1ColInt, table1ColFloat).
		SET(1, 22.2).
		WHERE(table1ColInt.GT_EQ(Int(33)))

	assertStatementRecordSQL(t, stmt)
}

func TestUpdateOneColumnWithSelect(t *testing.T) {
	stmt := table1.
		UPDATE(table1ColFloat).
		SET(
			table1.SELECT(table1ColFloat),
		).
		WHERE(table1Col1.EQ(Int(2))).
		RETURNING(table1Col1)

	assertStatementRecordSQL(t, stmt)
}

func TestInvalidInputs(t *testing.T) {
	assertStatementSqlErr(t, table1.UPDATE(table1ColInt).SET(1),
		"jet: WHERE clause not set")
	assertStatementSqlErr(t, table1.UPDATE(nil).SET(1),
		"jet: nil column in columns list")
}

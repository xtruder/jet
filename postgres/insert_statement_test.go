package postgres

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestInvalidInsert(t *testing.T) {
	testutils.StatementTests{
		{Name: "panics if values not set", Test: table1.INSERT(table1Col1),
			Panics: "jet: VALUES or QUERY has to be specified for INSERT statement"},
		{Name: "panics if nil columns", Test: table1.INSERT(nil).VALUES(1),
			Panics: "jet: nil column in columns list"},
	}.Run(t)
}

func TestInsertNilValue(t *testing.T) {
	testutils.StatementTest{Test: table1.INSERT(table1Col1).VALUES(nil)}.Assert(t)
}

func TestInsertSingleValue(t *testing.T) {
	testutils.StatementTest{Test: table1.INSERT(table1Col1).VALUES(1)}.Assert(t)
}

func TestInsertWithColumnList(t *testing.T) {
	columnList := ColumnList{table3ColInt, table3StrCol}

	testutils.StatementTest{Test: table3.INSERT(columnList).VALUES(1, 3)}.Assert(t)

}

func TestInsertDate(t *testing.T) {
	date := time.Date(1999, 1, 2, 3, 4, 5, 0, time.UTC)

	testutils.StatementTest{Test: table1.INSERT(table1ColTime).VALUES(date)}.Assert(t)
}

func TestInsertMultipleValues(t *testing.T) {
	stmt := table1.INSERT(table1Col1, table1ColFloat, table1ColBool).VALUES(1, 2, 3)

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsertMultipleRows(t *testing.T) {
	stmt := table1.INSERT(table1Col1, table1ColFloat).
		VALUES(1, 2).
		VALUES(11, 22).
		VALUES(111, 222)

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsertValuesFromModel(t *testing.T) {
	type Table1Model struct {
		Col1     *int
		ColFloat float64
	}

	one := 1

	toInsert := Table1Model{
		Col1:     &one,
		ColFloat: 1.11,
	}

	stmt := table1.INSERT(table1Col1, table1ColFloat).
		MODEL(toInsert).
		MODEL(&toInsert)

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsertValuesFromModelColumnMismatch(t *testing.T) {
	type Table1Model struct {
		Col1Prim int
		Col2     string
	}

	newData := Table1Model{
		Col1Prim: 1,
		Col2:     "one",
	}

	assert.PanicsWithValue(t, "missing struct field for column : col1", func() {
		table1.
			INSERT(table1Col1, table1ColFloat).
			MODEL(newData)
	})
}

func TestInsertFromNonStructModel(t *testing.T) {
	assert.PanicsWithValue(t, "jet: data has to be a struct", func() {
		table2.INSERT(table2ColInt).MODEL([]int{})
	})
}

func TestInsertQuery(t *testing.T) {
	stmt := table1.INSERT(table1Col1).
		QUERY(table1.SELECT(table1Col1))

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsertDefaultValue(t *testing.T) {
	stmt := table1.INSERT(table1Col1, table1ColFloat).
		VALUES(DEFAULT, "two")

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsert_ON_CONFLICT(t *testing.T) {
	stmt := table1.INSERT(table1Col1, table1ColBool).
		VALUES("one", "two").
		VALUES("1", "2").
		VALUES("theta", "beta").
		ON_CONFLICT(table1ColBool).WHERE(table1ColBool.IS_NOT_FALSE()).DO_UPDATE(
		SET(table1ColBool.SET(Bool(true)),
			table2ColInt.SET(Int(1)),
			ColumnList{table1Col1, table1ColBool}.SET(jet.ROW(Int(2), String("two"))),
		).WHERE(table1Col1.GT(Int(2))),
	).
		RETURNING(table1Col1, table1ColBool)

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsert_ON_CONFLICT_ON_CONSTRAINT(t *testing.T) {
	stmt := table1.INSERT(table1Col1, table1ColBool).
		VALUES("one", "two").
		VALUES("1", "2").
		ON_CONFLICT().ON_CONSTRAINT("idk_primary_key").DO_UPDATE(
		SET(table1ColBool.SET(Bool(false)),
			table2ColInt.SET(Int(1)),
			ColumnList{table1Col1, table1ColBool}.SET(jet.ROW(Int(2), String("two"))),
		).WHERE(table1Col1.GT(Int(2))),
	).
		RETURNING(table1Col1, table1ColBool)

	testutils.StatementTest{Test: stmt}.Assert(t)
}

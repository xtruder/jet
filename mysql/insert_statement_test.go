package mysql

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestInvalidInsert(t *testing.T) {
	testutils.StatementTests{
		{
			Name:   "panics without values or query",
			Test:   table1.INSERT(table1Col1),
			Panics: "jet: VALUES or QUERY has to be specified for INSERT statement",
		},
		{
			Name:   "panics with nil columns",
			Test:   table1.INSERT(nil).VALUES(1),
			Panics: "jet: nil column in columns list",
		},
	}.Run(t)
}

func TestInsertNilValue(t *testing.T) {
	testutils.StatementTest{Test: table1.INSERT(table1Col1).VALUES(nil)}.Assert(t)
}

func TestInsertSingleValue(t *testing.T) {
	testutils.StatementTest{Test: table1.INSERT(table1Col1).VALUES(1)}.Assert(t)
}

func TestInsertWithColumnList(t *testing.T) {
	columnList := ColumnList{table3ColInt}

	columnList = append(columnList, table3StrCol)

	testutils.StatementTest{Test: table3.INSERT(columnList).VALUES(1, 3)}.Assert(t)
}

func TestInsertDate(t *testing.T) {
	date := time.Date(1999, 1, 2, 3, 4, 5, 0, time.UTC)

	testutils.StatementTest{Test: table1.INSERT(table1ColTimestamp).VALUES(date)}.Assert(t)
}

func TestInsertMultipleValues(t *testing.T) {
	testutils.StatementTest{
		Test: table1.INSERT(table1Col1, table1ColFloat, table1Col3).VALUES(1, 2, 3),
	}.Assert(t)
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

func TestInsertDefaultValue(t *testing.T) {
	stmt := table1.INSERT(table1Col1, table1ColFloat).
		VALUES(DEFAULT, "two")

	testutils.StatementTest{Test: stmt}.Assert(t)
}

func TestInsertOnDuplicateKeyUpdate(t *testing.T) {
	stmt := func() InsertStatement {
		return table1.INSERT(table1Col1, table1ColFloat).
			VALUES(DEFAULT, "two")
	}

	t.Run("empty list", func(t *testing.T) {
		stmt := stmt().ON_DUPLICATE_KEY_UPDATE()
		testutils.StatementTest{Test: stmt}.Assert(t)
	})

	t.Run("one set", func(t *testing.T) {
		stmt := stmt().ON_DUPLICATE_KEY_UPDATE(table1ColFloat.SET(Float(11.1)))
		testutils.StatementTest{Test: stmt}.Assert(t)
	})

	t.Run("all types set", func(t *testing.T) {
		stmt := stmt().ON_DUPLICATE_KEY_UPDATE(
			table1ColBool.SET(Bool(true)),
			table1ColInt.SET(Int(11)),
			table1ColFloat.SET(Float(11.1)),
			table1ColString.SET(String("str")),
			table1ColTime.SET(Time(11, 23, 11)),
			table1ColTimestamp.SET(Timestamp(2020, 1, 22, 3, 4, 5)),
			table1ColDate.SET(Date(2020, 12, 1)),
		)
		testutils.StatementTest{Test: stmt}.Assert(t)
	})
}

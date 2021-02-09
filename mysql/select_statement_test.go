package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestInvalidSelect(t *testing.T) {
	testutils.StatementTest{
		Test:   SELECT(nil),
		Panics: "jet: Projection is nil",
	}.Assert(t)
}

func TestSelectColumnList(t *testing.T) {
	columnList := ColumnList{table2ColInt, table2ColFloat, table3ColInt}

	testutils.StatementTest{Test: SELECT(columnList).FROM(table2)}.Assert(t)
}

func TestSelectLiterals(t *testing.T) {
	testutils.StatementTest{Test: SELECT(Int(1), Float(2.2), Bool(false)).FROM(table1)}.Assert(t)
}

func TestSelectDistinct(t *testing.T) {
	testutils.StatementTest{Test: SELECT(table1ColBool).DISTINCT().FROM(table1)}.Assert(t)
}

func TestSelectFrom(t *testing.T) {
	testutils.StatementTests{
		{Name: "two cols",
			Test: SELECT(table1ColInt, table2ColFloat).FROM(table1)},
		{Name: "two cols and join",
			Test: SELECT(table1ColInt, table2ColFloat).FROM(table1.INNER_JOIN(table2, table1ColInt.EQ(table2ColInt)))},
		{Name: "inner join",
			Test: table1.INNER_JOIN(table2, table1ColInt.EQ(table2ColInt)).SELECT(table1ColInt, table2ColFloat)},
	}.Run(t)
}

func TestSelectWhere(t *testing.T) {
	testutils.StatementTests{
		{Name: "simple condition",
			Test: SELECT(table1ColInt).FROM(table1).WHERE(Bool(true))},
		{Name: "complex condition",
			Test: SELECT(table1ColInt).FROM(table1).WHERE(table1ColInt.GT_EQ(Int(10)))},
	}.Run(t)
}

func TestSelectGroupBy(t *testing.T) {
	testutils.StatementTest{
		Test: SELECT(table2ColInt).FROM(table2).GROUP_BY(table2ColFloat),
	}.Assert(t)
}

func TestSelectHaving(t *testing.T) {
	testutils.StatementTest{
		Test: SELECT(table3ColInt).FROM(table3).HAVING(table1ColBool.EQ(Bool(true))),
	}.Assert(t)
}

func TestSelectOrderBy(t *testing.T) {
	testutils.StatementTests{
		{Name: "single field",
			Test: SELECT(table2ColFloat).FROM(table2).ORDER_BY(table2ColInt.DESC())},
		{Name: "multiple field",
			Test: SELECT(table2ColFloat).FROM(table2).ORDER_BY(table2ColInt.DESC(), table2ColInt.ASC())},
	}.Run(t)
}

func TestSelectLimitOffset(t *testing.T) {
	testutils.StatementTests{
		{Name: "only limit",
			Test: SELECT(table2ColInt).FROM(table2).LIMIT(10)},
		{Name: "limit and offset",
			Test: SELECT(table2ColInt).FROM(table2).LIMIT(10).OFFSET(2)},
	}.Run(t)
}

func TestSelectLock(t *testing.T) {
	testutils.StatementTests{
		{Name: "for update",
			Test: SELECT(table1ColBool).FROM(table1).FOR(UPDATE())},
		{Name: "for update no wait",
			Test: SELECT(table1ColBool).FROM(table1).FOR(SHARE().NOWAIT())},
	}.Run(t)
}

func TestSelect_LOCK_IN_SHARE_MODE(t *testing.T) {
	testutils.StatementTest{Test: SELECT(table1ColBool).FROM(table1).LOCK_IN_SHARE_MODE()}.Assert(t)
}

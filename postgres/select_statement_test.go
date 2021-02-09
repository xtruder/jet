package postgres

import (
	"testing"
)

func TestInvalidSelect(t *testing.T) {
	assertStatementSqlErr(t, SELECT(nil), "jet: Projection is nil")
}

func TestSelectColumnList(t *testing.T) {
	columnList := ColumnList{table2ColInt, table2ColFloat, table3ColInt}

	assertStatementRecordSQL(t, SELECT(columnList).FROM(table2))
}

func TestSelectLiterals(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(Int(1), Float(2.2), Bool(false)).FROM(table1))
}

func TestSelectDistinct(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table1ColBool).DISTINCT().FROM(table1))
}

func TestSelectFrom(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table1ColInt, table2ColFloat).FROM(table1))

	assertStatementRecordSQL(t, SELECT(table1ColInt, table2ColFloat).FROM(table1.INNER_JOIN(table2, table1ColInt.EQ(table2ColInt))))

	assertStatementRecordSQL(t, table1.INNER_JOIN(table2, table1ColInt.EQ(table2ColInt)).SELECT(table1ColInt, table2ColFloat))
}

func TestSelectWhere(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table1ColInt).FROM(table1).WHERE(Bool(true)))

	assertStatementRecordSQL(t, SELECT(table1ColInt).FROM(table1).WHERE(table1ColInt.GT_EQ(Int(10))))
}

func TestSelectGroupBy(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table2ColInt).FROM(table2).GROUP_BY(table2ColFloat))
}

func TestSelectHaving(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table3ColInt).FROM(table3).HAVING(table1ColBool.EQ(Bool(true))))
}

func TestSelectOrderBy(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table2ColFloat).FROM(table2).ORDER_BY(table2ColInt.DESC()))
	assertStatementRecordSQL(t, SELECT(table2ColFloat).FROM(table2).ORDER_BY(table2ColInt.DESC(), table2ColInt.ASC()))
}

func TestSelectLimitOffset(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table2ColInt).FROM(table2).LIMIT(10))
	assertStatementRecordSQL(t, SELECT(table2ColInt).FROM(table2).LIMIT(10).OFFSET(2))
}

func TestSelectLock(t *testing.T) {
	assertStatementRecordSQL(t, SELECT(table1ColBool).FROM(table1).FOR(UPDATE()))

	assertStatementRecordSQL(t, SELECT(table1ColBool).FROM(table1).FOR(SHARE().NOWAIT()))

	assertStatementRecordSQL(t, SELECT(table1ColBool).FROM(table1).FOR(KEY_SHARE().NOWAIT()))

	assertStatementRecordSQL(t, SELECT(table1ColBool).FROM(table1).FOR(NO_KEY_UPDATE().SKIP_LOCKED()))
}

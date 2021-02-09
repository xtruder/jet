package postgres

import (
	"testing"
)

func TestSelectSets(t *testing.T) {
	select1 := SELECT(table1ColBool).FROM(table1)
	select2 := SELECT(table2ColBool).FROM(table2)

	assertStatementRecordSQL(t, select1.UNION(select2))

	assertStatementRecordSQL(t, select1.UNION_ALL(select2))

	assertStatementRecordSQL(t, select1.INTERSECT(select2))

	assertStatementRecordSQL(t, select1.INTERSECT_ALL(select2))

	assertStatementRecordSQL(t, select1.EXCEPT(select2))

	assertStatementRecordSQL(t, select1.EXCEPT_ALL(select2))
}

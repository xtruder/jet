package postgres

import "testing"

func TestWithStatement(t *testing.T) {
	cte1 := CTE("cte1")
	cte2 := CTE("cte2")

	stmt := WITH(
		cte1.AS(table1.SELECT(table1ColInt)),
		cte2.AS(
			table2.UPDATE().
				SET(table2ColInt.SET(Int(1))).
				WHERE(table2Col3.IN(cte1.SELECT(table1ColInt.From(cte1)))),
		),
	)(SELECT(cte1.AllColumns()).FROM(cte1))

	assertStatementRecordSQL(t, stmt)
}

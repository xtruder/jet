package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestOnConflict(t *testing.T) {
	modify := func(value *onConflictClause, modify func(v *onConflictClause)) *onConflictClause {
		modify(value)
		return value
	}

	doNothing := func(v *onConflictClause) { v.DO_NOTHING() }

	testutils.SerializerTests{
		{Name: "empty", Test: &onConflictClause{}},

		{Name: "empty do nothing", Test: modify(&onConflictClause{}, doNothing)},

		{Name: "do nothing",
			Test: modify(&onConflictClause{indexExpressions: ColumnList{table1ColBool}}, doNothing)},

		{Name: "on constraint do nothing",
			Test: modify(&onConflictClause{indexExpressions: ColumnList{table1ColBool}},
				func(v *onConflictClause) {
					v.ON_CONSTRAINT("table_pkey").DO_NOTHING()
				})},

		{Name: "where do update",
			Test: modify(&onConflictClause{indexExpressions: ColumnList{table1ColBool, table2ColFloat}},
				func(v *onConflictClause) {
					v.WHERE(table2ColFloat.ADD(table1ColInt).GT(table1ColFloat)).
						DO_UPDATE(
							SET(table1ColBool.SET(Bool(true)),
								table1ColInt.SET(Int(11))).
								WHERE(table2ColFloat.GT(Float(11.1))),
						)
				})},
	}.Run(t, Dialect)
}

package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestSelectSets(t *testing.T) {
	select1 := SELECT(table1ColBool).FROM(table1)
	select2 := SELECT(table2ColBool).FROM(table2)

	testutils.StatementTests{
		{Name: "union", Test: select1.UNION(select2)},
		{Name: "union all", Test: select1.UNION_ALL(select2)},
	}.Run(t)
}

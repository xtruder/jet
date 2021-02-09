package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestNewIntervalColumn(t *testing.T) {
	subQuery := SELECT(Int(1)).AsTable("sub_query")

	subQueryIntervalColumn := IntervalColumn("col_interval").From(subQuery)
	subQueryIntervalColumn2 := table1ColInterval.From(subQuery)

	testutils.SerializerTests{
		{Name: "interval col", Test: subQueryIntervalColumn},
		{Name: "interval col eq",
			Test: subQueryIntervalColumn.EQ(INTERVAL(2, HOUR, 10, MINUTE))},
		{Name: "interval col 2", Test: subQueryIntervalColumn2},
		{Name: "interval col 2 eq", Test: subQueryIntervalColumn2.EQ(INTERVAL(1, DAY))},
	}.Run(t, Dialect)

	testutils.ProjectionTests{
		{Name: "interval col projection", Test: subQueryIntervalColumn},
		{Name: "interval col 2 projection", Test: subQueryIntervalColumn2},
	}.Run(t, Dialect)
}

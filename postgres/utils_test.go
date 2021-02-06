package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/testutils"
)

var (
	table1Col1          = IntegerColumn("col1")
	table1ColInt        = IntegerColumn("col_int")
	table1ColFloat      = FloatColumn("col_float")
	table1ColTime       = TimeColumn("col_time")
	table1ColTimez      = TimezColumn("col_timez")
	table1ColTimestamp  = TimestampColumn("col_timestamp")
	table1ColTimestampz = TimestampzColumn("col_timestampz")
	table1ColBool       = BoolColumn("col_bool")
	table1ColDate       = DateColumn("col_date")
	table1ColInterval   = IntervalColumn("col_interval")
	table1ColArray      = ArrayColumn("col_array")
)

var table1 = NewTable(
	"db",
	"table1",
	table1Col1,
	table1ColInt,
	table1ColFloat,
	table1ColTime,
	table1ColTimez,
	table1ColBool,
	table1ColDate,
	table1ColTimestamp,
	table1ColTimestampz,
	table1ColInterval,
	table1ColArray,
)

var (
	table2Col3          = IntegerColumn("col3")
	table2Col4          = IntegerColumn("col4")
	table2ColInt        = IntegerColumn("col_int")
	table2ColFloat      = FloatColumn("col_float")
	table2ColStr        = StringColumn("col_str")
	table2ColBool       = BoolColumn("col_bool")
	table2ColTime       = TimeColumn("col_time")
	table2ColTimez      = TimezColumn("col_timez")
	table2ColTimestamp  = TimestampColumn("col_timestamp")
	table2ColTimestampz = TimestampzColumn("col_timestampz")
	table2ColDate       = DateColumn("col_date")
	table2ColInterval   = IntervalColumn("col_interval")
	table2ColArray      = ArrayColumn("col_array")
)

var table2 = NewTable(
	"db",
	"table2",
	table2Col3,
	table2Col4,
	table2ColInt,
	table2ColFloat,
	table2ColStr,
	table2ColBool,
	table2ColTime,
	table2ColTimez,
	table2ColDate,
	table2ColTimestamp,
	table2ColTimestampz,
	table2ColInterval,
	table2ColArray,
)

var (
	table3Col1   = IntegerColumn("col1")
	table3ColInt = IntegerColumn("col_int")
	table3StrCol = StringColumn("col2")
)

var table3 = NewTable(
	"db",
	"table3",
	table3Col1,
	table3ColInt,
	table3StrCol)

func assertSerialize(t *testing.T, serializer jet.Serializer, query string, args ...interface{}) {
	testutils.AssertSerialize(t, Dialect, serializer, query, args...)
}

func assertClauseSerialize(t *testing.T, clause jet.Clause, query string, args ...interface{}) {
	testutils.AssertClauseSerialize(t, Dialect, clause, query, args...)
}

func assertSerializeErr(t *testing.T, serializer jet.Serializer, errString string) {
	testutils.AssertSerializeErr(t, Dialect, serializer, errString)
}

func assertProjectionSerialize(t *testing.T, projection jet.Projection, query string, args ...interface{}) {
	testutils.AssertProjectionSerialize(t, Dialect, projection, query, args...)
}

var assertStatementSql = testutils.AssertStatementSql
var assertDebugStatementSql = testutils.AssertDebugStatementSql
var assertStatementSqlErr = testutils.AssertStatementSqlErr
var assertPanicErr = testutils.AssertPanicErr

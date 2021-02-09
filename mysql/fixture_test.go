package mysql

var table1Col1 = IntegerColumn("col1")
var table1ColBool = BoolColumn("col_bool")
var table1ColInt = IntegerColumn("col_int")
var table1ColFloat = FloatColumn("col_float")
var table1ColString = StringColumn("col_string")
var table1Col3 = IntegerColumn("col3")
var table1ColTimestamp = TimestampColumn("col_timestamp")
var table1ColDate = DateColumn("col_date")
var table1ColTime = TimeColumn("col_time")
var table1ColJSON = JSONColumn("col_json")

var table1 = NewTable(
	"db",
	"table1",
	table1Col1,
	table1ColInt,
	table1ColFloat,
	table1ColString,
	table1Col3,
	table1ColBool,
	table1ColDate,
	table1ColTimestamp,
	table1ColTime,
	table1ColJSON,
)

var table2Col3 = IntegerColumn("col3")
var table2Col4 = IntegerColumn("col4")
var table2ColInt = IntegerColumn("col_int")
var table2ColFloat = FloatColumn("col_float")
var table2ColStr = StringColumn("col_str")
var table2ColBool = BoolColumn("col_bool")
var table2ColTimestamp = TimestampColumn("col_timestamp")
var table2ColDate = DateColumn("col_date")
var table2ColJSON = JSONColumn("col_json")

var table2 = NewTable(
	"db",
	"table2",
	table2Col3,
	table2Col4,
	table2ColInt,
	table2ColFloat,
	table2ColStr,
	table2ColBool,
	table2ColDate,
	table2ColTimestamp,
	table2ColJSON,
)

var table3Col1 = IntegerColumn("col1")
var table3ColInt = IntegerColumn("col_int")
var table3StrCol = StringColumn("col2")
var table3 = NewTable(
	"db",
	"table3",
	table3Col1,
	table3ColInt,
	table3StrCol)

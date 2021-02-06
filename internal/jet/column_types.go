package jet

// ColumnBool is interface for SQL boolean columns.
type ColumnBool interface {
	BoolExpression
	Column

	From(subQuery SelectTable) ColumnBool
	SET(boolExp BoolExpression) ColumnAssigment
}

type boolColumnImpl struct {
	boolInterfaceImpl
	ColumnExpression
}

func (i *boolColumnImpl) From(subQuery SelectTable) ColumnBool {
	newBoolColumn := BoolColumn(i.Name())
	newBoolColumn.SetTableName(i.TableName())
	newBoolColumn.SetSubQuery(subQuery)

	return newBoolColumn
}

func (i *boolColumnImpl) SET(boolExp BoolExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: boolExp,
	}
}

// BoolColumn creates named bool column.
func BoolColumn(name string) ColumnBool {
	boolColumn := &boolColumnImpl{}
	boolColumn.ColumnExpression = NewColumnExpression(name, "", boolColumn)
	boolColumn.boolInterfaceImpl.parent = boolColumn

	return boolColumn
}

//------------------------------------------------------//

// ColumnFloat is interface for SQL real, numeric, decimal or double precision column.
type ColumnFloat interface {
	FloatExpression
	Column

	From(subQuery SelectTable) ColumnFloat
	SET(floatExp FloatExpression) ColumnAssigment
}

type floatColumnImpl struct {
	floatInterfaceImpl
	ColumnExpression
}

func (i *floatColumnImpl) From(subQuery SelectTable) ColumnFloat {
	newFloatColumn := FloatColumn(i.Name())
	newFloatColumn.SetTableName(i.TableName())
	newFloatColumn.SetSubQuery(subQuery)

	return newFloatColumn
}

func (i *floatColumnImpl) SET(floatExp FloatExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: floatExp,
	}
}

// FloatColumn creates named float column.
func FloatColumn(name string) ColumnFloat {
	floatColumn := &floatColumnImpl{}
	floatColumn.floatInterfaceImpl.parent = floatColumn
	floatColumn.ColumnExpression = NewColumnExpression(name, "", floatColumn)

	return floatColumn
}

//------------------------------------------------------//

// ColumnInteger is interface for SQL smallint, integer, bigint columns.
type ColumnInteger interface {
	IntegerExpression
	Column

	From(subQuery SelectTable) ColumnInteger
	SET(intExp IntegerExpression) ColumnAssigment
}

type integerColumnImpl struct {
	integerInterfaceImpl
	ColumnExpression
}

func (i *integerColumnImpl) From(subQuery SelectTable) ColumnInteger {
	newIntColumn := IntegerColumn(i.Name())
	newIntColumn.SetTableName(i.TableName())
	newIntColumn.SetSubQuery(subQuery)

	return newIntColumn
}

func (i *integerColumnImpl) SET(intExp IntegerExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: intExp,
	}
}

// IntegerColumn creates named integer column.
func IntegerColumn(name string) ColumnInteger {
	integerColumn := &integerColumnImpl{}
	integerColumn.integerInterfaceImpl.parent = integerColumn
	integerColumn.ColumnExpression = NewColumnExpression(name, "", integerColumn)

	return integerColumn
}

//------------------------------------------------------//

// ColumnString is interface for SQL text, character, character varying
// bytea, uuid columns and enums types.
type ColumnString interface {
	StringExpression
	Column

	From(subQuery SelectTable) ColumnString
	SET(stringExp StringExpression) ColumnAssigment
}

type stringColumnImpl struct {
	stringInterfaceImpl
	ColumnExpression
}

func (i *stringColumnImpl) From(subQuery SelectTable) ColumnString {
	newStrColumn := StringColumn(i.Name())
	newStrColumn.SetTableName(i.TableName())
	newStrColumn.SetSubQuery(subQuery)

	return newStrColumn
}

func (i *stringColumnImpl) SET(stringExp StringExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: stringExp,
	}
}

// StringColumn creates named string column.
func StringColumn(name string) ColumnString {
	stringColumn := &stringColumnImpl{}
	stringColumn.stringInterfaceImpl.parent = stringColumn
	stringColumn.ColumnExpression = NewColumnExpression(name, "", stringColumn)

	return stringColumn
}

//------------------------------------------------------//

// ColumnTime is interface for SQL time column.
type ColumnTime interface {
	TimeExpression
	Column

	From(subQuery SelectTable) ColumnTime
	SET(timeExp TimeExpression) ColumnAssigment
}

type timeColumnImpl struct {
	timeInterfaceImpl
	ColumnExpression
}

func (i *timeColumnImpl) From(subQuery SelectTable) ColumnTime {
	newTimeColumn := TimeColumn(i.Name())
	newTimeColumn.SetTableName(i.TableName())
	newTimeColumn.SetSubQuery(subQuery)

	return newTimeColumn
}

func (i *timeColumnImpl) SET(timeExp TimeExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: timeExp,
	}
}

// TimeColumn creates named time column
func TimeColumn(name string) ColumnTime {
	timeColumn := &timeColumnImpl{}
	timeColumn.timeInterfaceImpl.parent = timeColumn
	timeColumn.ColumnExpression = NewColumnExpression(name, "", timeColumn)
	return timeColumn
}

//------------------------------------------------------//

// ColumnTimez is interface of SQL time with time zone columns.
type ColumnTimez interface {
	TimezExpression
	Column

	From(subQuery SelectTable) ColumnTimez
}

type timezColumnImpl struct {
	timezInterfaceImpl
	ColumnExpression
}

func (i *timezColumnImpl) From(subQuery SelectTable) ColumnTimez {
	newTimezColumn := TimezColumn(i.Name())
	newTimezColumn.SetTableName(i.TableName())
	newTimezColumn.SetSubQuery(subQuery)

	return newTimezColumn
}

func (i *timezColumnImpl) SET(timezExp TimezExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: timezExp,
	}
}

// TimezColumn creates named time with time zone column.
func TimezColumn(name string) ColumnTimez {
	timezColumn := &timezColumnImpl{}
	timezColumn.timezInterfaceImpl.parent = timezColumn
	timezColumn.ColumnExpression = NewColumnExpression(name, "", timezColumn)

	return timezColumn
}

//------------------------------------------------------//

// ColumnTimestamp is interface of SQL timestamp columns.
type ColumnTimestamp interface {
	TimestampExpression
	Column

	From(subQuery SelectTable) ColumnTimestamp
	SET(timestampExp TimestampExpression) ColumnAssigment
}

type timestampColumnImpl struct {
	timestampInterfaceImpl
	ColumnExpression
}

func (i *timestampColumnImpl) From(subQuery SelectTable) ColumnTimestamp {
	newTimestampColumn := TimestampColumn(i.Name())
	newTimestampColumn.SetTableName(i.TableName())
	newTimestampColumn.SetSubQuery(subQuery)

	return newTimestampColumn
}

func (i *timestampColumnImpl) SET(timestampExp TimestampExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: timestampExp,
	}
}

// TimestampColumn creates named timestamp column
func TimestampColumn(name string) ColumnTimestamp {
	timestampColumn := &timestampColumnImpl{}
	timestampColumn.timestampInterfaceImpl.parent = timestampColumn
	timestampColumn.ColumnExpression = NewColumnExpression(name, "", timestampColumn)

	return timestampColumn
}

//------------------------------------------------------//

// ColumnTimestampz is interface of SQL timestamp with timezone columns.
type ColumnTimestampz interface {
	TimestampzExpression
	Column

	From(subQuery SelectTable) ColumnTimestampz
	SET(timestampzExp TimestampzExpression) ColumnAssigment
}

type timestampzColumnImpl struct {
	timestampzInterfaceImpl
	ColumnExpression
}

func (i *timestampzColumnImpl) From(subQuery SelectTable) ColumnTimestampz {
	newTimestampzColumn := TimestampzColumn(i.Name())
	newTimestampzColumn.SetTableName(i.TableName())
	newTimestampzColumn.SetSubQuery(subQuery)

	return newTimestampzColumn
}

func (i *timestampzColumnImpl) SET(timestampzExp TimestampzExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: timestampzExp,
	}
}

// TimestampzColumn creates named timestamp with time zone column.
func TimestampzColumn(name string) ColumnTimestampz {
	timestampzColumn := &timestampzColumnImpl{}
	timestampzColumn.timestampzInterfaceImpl.parent = timestampzColumn
	timestampzColumn.ColumnExpression = NewColumnExpression(name, "", timestampzColumn)

	return timestampzColumn
}

//------------------------------------------------------//

// ColumnDate is interface of SQL date columns.
type ColumnDate interface {
	DateExpression
	Column

	From(subQuery SelectTable) ColumnDate
	SET(dateExp DateExpression) ColumnAssigment
}

type dateColumnImpl struct {
	dateInterfaceImpl
	ColumnExpression
}

func (i *dateColumnImpl) From(subQuery SelectTable) ColumnDate {
	newDateColumn := DateColumn(i.Name())
	newDateColumn.SetTableName(i.TableName())
	newDateColumn.SetSubQuery(subQuery)

	return newDateColumn
}

func (i *dateColumnImpl) SET(dateExp DateExpression) ColumnAssigment {
	return columnAssigmentImpl{
		column:     i,
		expression: dateExp,
	}
}

// DateColumn creates named date column.
func DateColumn(name string) ColumnDate {
	dateColumn := &dateColumnImpl{}
	dateColumn.dateInterfaceImpl.parent = dateColumn
	dateColumn.ColumnExpression = NewColumnExpression(name, "", dateColumn)
	return dateColumn
}

package postgres

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

// Column is common column interface for all types of columns.
type Column = jet.ColumnExpression

// ColumnList function returns list of columns that be used as projection or column list for UPDATE and INSERT statement.
type ColumnList = jet.ColumnList

// ColumnBool is interface for SQL boolean columns.
type ColumnBool = jet.ColumnBool

// BoolColumn creates named bool column.
var BoolColumn = jet.BoolColumn

// ColumnString is interface for SQL text, character, character varying
// bytea, uuid columns and enums types.
type ColumnString = jet.ColumnString

// StringColumn creates named string column.
var StringColumn = jet.StringColumn

// ColumnInteger is interface for SQL smallint, integer, bigint columns.
type ColumnInteger = jet.ColumnInteger

// IntegerColumn creates named integer column.
var IntegerColumn = jet.IntegerColumn

// ColumnFloat is interface for SQL real, numeric, decimal or double precision column.
type ColumnFloat = jet.ColumnFloat

// FloatColumn creates named float column.
var FloatColumn = jet.FloatColumn

// ColumnDate is interface of SQL date columns.
type ColumnDate = jet.ColumnDate

// DateColumn creates named date column.
var DateColumn = jet.DateColumn

// ColumnTime is interface for SQL time column.
type ColumnTime = jet.ColumnTime

// TimeColumn creates named time column
var TimeColumn = jet.TimeColumn

// ColumnTimez is interface of SQL time with time zone columns.
type ColumnTimez = jet.ColumnTimez

// TimezColumn creates named time with time zone column.
var TimezColumn = jet.TimezColumn

// ColumnTimestamp is interface of SQL timestamp columns.
type ColumnTimestamp = jet.ColumnTimestamp

// TimestampColumn creates named timestamp column
var TimestampColumn = jet.TimestampColumn

// ColumnTimestampz is interface of SQL timestamp with timezone columns.
type ColumnTimestampz = jet.ColumnTimestampz

// TimestampzColumn creates named timestamp with time zone column.
var TimestampzColumn = jet.TimestampzColumn

//------------------------------------------------------//

// ColumnInterval is interface of PostgreSQL interval columns.
type ColumnInterval interface {
	IntervalExpression
	jet.Column

	From(subQuery SelectTable) ColumnInterval
}

type intervalColumnImpl struct {
	intervalInterfaceImpl
	jet.ColumnExpression
}

func (i *intervalColumnImpl) From(subQuery SelectTable) ColumnInterval {
	newIntervalColumn := IntervalColumn(i.Name())
	jet.SetTableName(newIntervalColumn, i.TableName())
	jet.SetSubQuery(newIntervalColumn, subQuery)

	return newIntervalColumn
}

// IntervalColumn creates named interval column.
func IntervalColumn(name string) ColumnInterval {
	intervalColumn := &intervalColumnImpl{}
	intervalColumn.intervalInterfaceImpl.parent = intervalColumn
	intervalColumn.ColumnExpression = jet.NewColumnExpression(name, "", intervalColumn)
	return intervalColumn
}

//------------------------------------------------------//

// ColumnArray is interface for SQL array types.
type ColumnArray interface {
	ArrayExpression
	jet.Column

	From(subQuery SelectTable) ColumnArray
	SET(arrayExp ArrayExpression) ColumnAssigment
}

type arrayColumnImpl struct {
	arrayInterfaceImpl
	jet.ColumnExpression
}

func (i *arrayColumnImpl) From(subQuery SelectTable) ColumnArray {
	newArrayColumn := ArrayColumn(i.Name())
	newArrayColumn.SetTableName(i.TableName())
	newArrayColumn.SetSubQuery(subQuery)

	return newArrayColumn
}

func (i *arrayColumnImpl) SET(arrayExp ArrayExpression) ColumnAssigment {
	return jet.NewColumnAssigment(i, arrayExp)
}

// ArrayColumn creates named array column.
func ArrayColumn(name string) ColumnArray {
	arrayColumn := &arrayColumnImpl{}
	arrayColumn.arrayInterfaceImpl.parent = arrayColumn
	arrayColumn.ColumnExpression = jet.NewColumnExpression(name, "", arrayColumn)

	return arrayColumn
}

//------------------------------------------------------//

// ColumnJSON is interface for SQL json types.
type ColumnJSON interface {
	JSONExpression
	jet.Column

	From(subQuery SelectTable) ColumnJSON
	SET(jsonExp JSONExpression) ColumnAssigment
}

type jsonColumnImpl struct {
	jsonInterfaceImpl
	jet.ColumnExpression
}

func (i *jsonColumnImpl) From(subQuery SelectTable) ColumnJSON {
	newStrColumn := JSONColumn(i.Name())
	newStrColumn.SetTableName(i.TableName())
	newStrColumn.SetSubQuery(subQuery)

	return newStrColumn
}

func (i *jsonColumnImpl) SET(jsonExp JSONExpression) ColumnAssigment {
	return jet.NewColumnAssigment(i, jsonExp)
}

// JSONColumn creates named json column.
func JSONColumn(name string) ColumnJSON {
	jsonColumn := &jsonColumnImpl{}
	jsonColumn.jsonInterfaceImpl.parent = jsonColumn
	jsonColumn.ColumnExpression = jet.NewColumnExpression(name, "", jsonColumn)

	return jsonColumn
}

//------------------------------------------------------//

// ColumnJSON is interface for SQL json types.
type ColumnJSONB interface {
	JSONBExpression
	jet.Column

	From(subQuery SelectTable) ColumnJSONB
	SET(jsonbExp JSONBExpression) ColumnAssigment
}

type jsonbColumnImpl struct {
	jsonbInterfaceImpl
	jet.ColumnExpression
}

func (i *jsonbColumnImpl) From(subQuery SelectTable) ColumnJSONB {
	newStrColumn := JSONBColumn(i.Name())
	newStrColumn.SetTableName(i.TableName())
	newStrColumn.SetSubQuery(subQuery)

	return newStrColumn
}

func (i *jsonbColumnImpl) SET(jsonbExp JSONBExpression) ColumnAssigment {
	return jet.NewColumnAssigment(i, jsonbExp)
}

// JSONBColumn creates named jsonb column.
func JSONBColumn(name string) ColumnJSONB {
	jsonbColumn := &jsonbColumnImpl{}
	jsonbColumn.jsonbInterfaceImpl.parent = jsonbColumn
	jsonbColumn.ColumnExpression = jet.NewColumnExpression(name, "", jsonbColumn)

	return jsonbColumn
}

package postgres

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestINTERVAL(t *testing.T) {
	testutils.SerializerTests{
		{Name: "year", Test: INTERVAL(1, YEAR)},
		{Name: "month", Test: INTERVAL(1, MONTH)},
		{Name: "week", Test: INTERVAL(1, WEEK)},
		{Name: "day", Test: INTERVAL(1, DAY)},
		{Name: "hour", Test: INTERVAL(1, HOUR)},
		{Name: "minute", Test: INTERVAL(1, MINUTE)},
		{Name: "second", Test: INTERVAL(1, SECOND)},
		{Name: "millisecond", Test: INTERVAL(1, MILLISECOND)},
		{Name: "microsecond", Test: INTERVAL(1, MICROSECOND)},
		{Name: "decade", Test: INTERVAL(1, DECADE)},
		{Name: "century", Test: INTERVAL(1, CENTURY)},
		{Name: "millenium", Test: INTERVAL(1, MILLENNIUM)},
		{Name: "year month", Test: INTERVAL(1, YEAR, 10, MONTH)},
		{Name: "year month day", Test: INTERVAL(1, YEAR, 10, MONTH, 20, DAY)},
		{Name: "year month day hour", Test: INTERVAL(1, YEAR, 10, MONTH, 20, DAY, 3, HOUR)},
		{Name: "year not null", Test: INTERVAL(1, YEAR).IS_NOT_NULL()},
		{Name: "float year", Test: INTERVAL(5.2, YEAR)},
	}.Run(t, Dialect)

	testutils.ProjectionTests{{Name: "as one year",
		Test: INTERVAL(1, YEAR).AS("one year")}}.Run(t, Dialect)
}

func TestINTERVALd(t *testing.T) {
	testutils.SerializerTests{
		{Name: "zero", Test: INTERVALd(0)},
		{Name: "micro", Test: INTERVALd(1 * time.Microsecond)},
		{Name: "milli", Test: INTERVALd(1 * time.Millisecond)},
		{Name: "second", Test: INTERVALd(1 * time.Second)},
		{Name: "minute", Test: INTERVALd(1 * time.Minute)},
		{Name: "hour", Test: INTERVALd(1 * time.Hour)},
		{Name: "hours", Test: INTERVALd(24 * time.Hour)},
		{Name: "mixed",
			Test: INTERVALd(24*time.Hour + 2*time.Hour + 3*time.Minute + 4*time.Second + 5*time.Microsecond)},
	}.Run(t, Dialect)
}

func TestINTERVAL_InvalidParams(t *testing.T) {
	assert.PanicsWithValue(t, "jet: invalid number of quantity and unit fields", func() { INTERVAL() })
	assert.PanicsWithValue(t, "jet: invalid number of quantity and unit fields", func() { INTERVAL(1) })
	assert.PanicsWithValue(t, "jet: invalid INTERVAL unit type", func() { INTERVAL(1, 2) })
}

func TestDateTimeIntervalArithmetic(t *testing.T) {
	testutils.SerializerTests{
		{Name: "date add", Test: table2ColDate.ADD(INTERVAL(1, HOUR))},
		{Name: "date sub", Test: table2ColDate.SUB(INTERVAL(1, HOUR))},
		{Name: "time add", Test: table2ColTime.ADD(INTERVAL(1, HOUR))},
		{Name: "time sub", Test: table2ColTime.SUB(INTERVAL(1, HOUR))},
		{Name: "timez add", Test: table2ColTimez.ADD(INTERVAL(1, HOUR))},
		{Name: "timez sub", Test: table2ColTimez.SUB(INTERVAL(1, HOUR))},
		{Name: "timestamp add", Test: table2ColTimestamp.ADD(INTERVAL(1, HOUR))},
		{Name: "timestamp sub", Test: table2ColTimestamp.SUB(INTERVAL(1, HOUR))},
		{Name: "timestampz add", Test: table2ColTimestampz.ADD(INTERVAL(1, HOUR))},
		{Name: "timestampz sub", Test: table2ColTimestampz.SUB(INTERVAL(1, HOUR))},
	}.Run(t, Dialect)
}

func TestIntervalExpressionMethods(t *testing.T) {
	testutils.SerializerTests{
		{Name: "eq col", Test: table1ColInterval.EQ(table2ColInterval)},
		{Name: "eq interval literal", Test: table1ColInterval.EQ(INTERVAL(10, SECOND))},
		{Name: "eq intervald literal", Test: table1ColInterval.EQ(INTERVALd(11 * time.Minute))},
		{Name: "eq intervald and eq bool", Test: table1ColInterval.EQ(INTERVALd(11 * time.Minute)).EQ(Bool(false))},
		{Name: "not eq", Test: table1ColInterval.NOT_EQ(table2ColInterval)},
		{Name: "is distinct from", Test: table1ColInterval.IS_DISTINCT_FROM(table2ColInterval)},
		{Name: "is not distinct from", Test: table1ColInterval.IS_NOT_DISTINCT_FROM(table2ColInterval)},
		{Name: "lt", Test: table1ColInterval.LT(table2ColInterval)},
		{Name: "lt eq", Test: table1ColInterval.LT_EQ(table2ColInterval)},
		{Name: "gt", Test: table1ColInterval.GT(table2ColInterval)},
		{Name: "gt eq", Test: table1ColInterval.GT_EQ(table2ColInterval)},
		{Name: "add", Test: table1ColInterval.ADD(table2ColInterval)},
		{Name: "sub", Test: table1ColInterval.SUB(table2ColInterval)},
		{Name: "mul int", Test: table1ColInterval.MUL(table2ColInt)},
		{Name: "mul float", Test: table1ColInterval.MUL(table2ColFloat)},
		{Name: "div int", Test: table1ColInterval.DIV(table2ColInt)},
		{Name: "div float", Test: table1ColInterval.DIV(table2ColFloat)},
	}.Run(t, Dialect)
}

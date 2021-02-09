package mysql

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestINTERVAL(t *testing.T) {
	testutils.SerializerTests{
		{Name: "year month", Test: INTERVAL("3-2", YEAR_MONTH)},
		{Name: "year month neg", Test: INTERVAL("-3-2", YEAR_MONTH)},
		{Name: "day hour", Test: INTERVAL("10 25", DAY_HOUR)},
		{Name: "day hour neg", Test: INTERVAL("-10 25", DAY_HOUR)},
		{Name: "day minute", Test: INTERVAL("10 25:15", DAY_MINUTE)},
		{Name: "day minute neg", Test: INTERVAL("-10 25:15", DAY_MINUTE)},
		{Name: "day second", Test: INTERVAL("10 25:15:08", DAY_SECOND)},
		{Name: "day second neg", Test: INTERVAL("-10 25:15:08", DAY_SECOND)},
		{Name: "day micros", Test: INTERVAL("10 25:15:08.000100", DAY_MICROSECOND)},
		{Name: "day micros neg", Test: INTERVAL("-10 25:15:08.000100", DAY_MICROSECOND)},
		{Name: "hour minute", Test: INTERVAL("15:08", HOUR_MINUTE)},
		{Name: "hour minute neg", Test: INTERVAL("-15:08", HOUR_MINUTE)},
		{Name: "hour second", Test: INTERVAL("15:08:03", HOUR_SECOND)},
		{Name: "hour second neg", Test: INTERVAL("-15:08:03", HOUR_SECOND)},
		{Name: "hour micros neg", Test: INTERVAL("25:15:08.000100", HOUR_MICROSECOND)},
		{Name: "minute second", Test: INTERVAL("08:03", MINUTE_SECOND)},
		{Name: "minute second neg", Test: INTERVAL("-08:03", MINUTE_SECOND)},
		{Name: "minute micros", Test: INTERVAL("15:08.000100", MINUTE_MICROSECOND)},
		{Name: "minute micros neg", Test: INTERVAL("-15:08.000100", MINUTE_MICROSECOND)},
		{Name: "second micros", Test: INTERVAL("08.000100", SECOND_MICROSECOND)},
		{Name: "second micros neg", Test: INTERVAL("-08.000100", SECOND_MICROSECOND)},

		{Name: "int second", Test: INTERVAL(15, SECOND)},
		{Name: "int micros", Test: INTERVAL(1, MICROSECOND)},
		{Name: "int minute", Test: INTERVAL(2, MINUTE)},
		{Name: "int hour", Test: INTERVAL(3, HOUR)},
		{Name: "int day", Test: INTERVAL(4, DAY)},
		{Name: "int month", Test: INTERVAL(5, MONTH)},
		{Name: "int year", Test: INTERVAL(6, YEAR)},
		{Name: "negint year", Test: INTERVAL(-6, YEAR)},

		{Name: "uint year", Test: INTERVAL(uint(6), YEAR)},
		{Name: "uint16 year", Test: INTERVAL(int16(7), YEAR)},
		{Name: "float year", Test: INTERVAL(3.5, YEAR)},
	}.Run(t, Dialect)
}

func TestINTERVAL_InvalidUnitType(t *testing.T) {
	assert.PanicsWithValue(t, "jet: INTERVAL invalid value type. Numeric type expected", func() { INTERVAL("11", HOUR) })
	assert.PanicsWithValue(t, "jet: INTERVAL invalid format", func() { INTERVAL("11", YEAR_MONTH) })
	assert.PanicsWithValue(t, "jet: INTERVAL invalid format", func() { INTERVAL("11+11", YEAR_MONTH) })
	assert.PanicsWithValue(t, "jet: INTERNAL invalid value type. String type expected", func() { INTERVAL(156.11, YEAR_MONTH) })
}

func TestINTERVALd(t *testing.T) {
	testutils.SerializerTests{
		{Name: "micros", Test: INTERVALd(3 * time.Microsecond)},
		{Name: "micros neg", Test: INTERVALd(-1 * time.Microsecond)},
		{Name: "seconds", Test: INTERVALd(3 * time.Second)},
		{Name: "seconds micros", Test: INTERVALd(3*time.Second + 4*time.Microsecond)},
		{Name: "seconds neg", Test: INTERVALd(-1 * time.Second)},
		{Name: "minutes", Test: INTERVALd(3 * time.Minute)},
		{Name: "minutes seconds", Test: INTERVALd(3*time.Minute + 4*time.Second)},
		{Name: "minutes seconds and micros", Test: INTERVALd(3*time.Minute + 4*time.Second + 5*time.Microsecond)},
		{Name: "minutes neg", Test: INTERVALd(-11 * time.Minute)},
		{Name: "minutes seconds neg", Test: INTERVALd(-11*time.Minute - 22*time.Second)},
		{Name: "hours", Test: INTERVALd(3 * time.Hour)},
		{Name: "hours and minutes", Test: INTERVALd(3*time.Hour + 4*time.Minute)},
		{Name: "hours minutes seconds", Test: INTERVALd(3*time.Hour + 4*time.Minute + 5*time.Second)},
		{Name: "hours minuts seconds millis", Test: INTERVALd(3*time.Hour + 4*time.Minute + 5*time.Second + 6*time.Millisecond)},
		{Name: "hours", Test: INTERVALd(-11 * time.Hour)},
		{Name: "hours and minutes neg", Test: INTERVALd(-11*time.Hour - 22*time.Minute)},
		{Name: "hours", Test: INTERVALd(3 * 24 * time.Hour)},
		{Name: "hours", Test: INTERVALd(3*24*time.Hour + 4*time.Hour)},
		{Name: "hours", Test: INTERVALd(3*24*time.Hour + 4*time.Hour + 5*time.Minute)},
		{Name: "hours minutes seconds", Test: INTERVALd(3*24*time.Hour + 4*time.Hour + 5*time.Minute + 6*time.Second)},
		{Name: "hours minutes seconds micros", Test: INTERVALd(3*24*time.Hour + 4*time.Hour + 5*time.Minute + 6*time.Second + 7*time.Microsecond)},
		{Name: "hours neg", Test: INTERVALd(-11 * 24 * time.Hour)},
		{Name: "all", Test: INTERVALd(1*time.Hour + 2*time.Minute + 3*time.Second + 345*time.Microsecond)},
		{Name: "all neg", Test: INTERVALd(-1 * (1*time.Hour + 2*time.Minute + 3*time.Second + 345*time.Microsecond))},
	}.Run(t, Dialect)
}

func TestINTERVALe(t *testing.T) {
	testutils.SerializerTests{
		{Name: "micros", Test: INTERVALe(table1ColFloat, MICROSECOND)},
		{Name: "seconds", Test: INTERVALe(table1ColFloat, SECOND)},
		{Name: "minutes", Test: INTERVALe(table1ColFloat, MINUTE)},
		{Name: "hours", Test: INTERVALe(table1ColFloat, HOUR)},
		{Name: "days", Test: INTERVALe(table1ColFloat, DAY)},
		{Name: "weeks", Test: INTERVALe(table1ColFloat, WEEK)},
		{Name: "months", Test: INTERVALe(table1ColFloat, MONTH)},
		{Name: "quater", Test: INTERVALe(table1ColFloat, QUARTER)},
		{Name: "year", Test: INTERVALe(table1ColFloat, YEAR)},
	}.Run(t, Dialect)
}

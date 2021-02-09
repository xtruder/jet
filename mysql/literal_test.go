package mysql

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestBool(t *testing.T) {
	testutils.SerializerTest{Test: Bool(false)}.Assert(t, Dialect)
}

func TestInt(t *testing.T) {
	testutils.SerializerTest{Test: Int(1)}.Assert(t, Dialect)
}

func TestFloat(t *testing.T) {
	testutils.SerializerTest{Test: Float(12.34)}.Assert(t, Dialect)
}

func TestString(t *testing.T) {
	testutils.SerializerTest{Test: String("Some text")}.Assert(t, Dialect)
}

func TestDate(t *testing.T) {
	testutils.SerializerTests{
		{Name: "Date", Test: Date(2014, time.January, 2)},
		{Name: "DateT", Test: DateT(testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0))},
	}.Run(t, Dialect)
}

func TestTime(t *testing.T) {
	testutils.SerializerTests{
		{Name: "Time", Test: Time(10, 15, 30)},
		{Name: "TimeT", Test: TimeT(testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0))},
	}.Run(t, Dialect)
}

func TestDateTime(t *testing.T) {
	testutils.SerializerTests{
		{Name: "DateTime", Test: DateTime(2010, time.March, 30, 10, 15, 30)},
		{Name: "DateTimeT", Test: DateTimeT(testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0))},
	}.Run(t, Dialect)
}

func TestTimestamp(t *testing.T) {
	testutils.SerializerTests{
		{Name: "Timestamp", Test: Timestamp(2010, time.March, 30, 10, 15, 30)},
		{Name: "TimestampT", Test: TimestampT(testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0))},
	}.Run(t, Dialect)
}

func TestJSON(t *testing.T) {
	testutils.SerializerTests{
		{Name: "object", Test: JSON(map[string]string{"key": "value"})},
		{Name: "string", Test: JSON("value")},
	}.Run(t, Dialect)
}

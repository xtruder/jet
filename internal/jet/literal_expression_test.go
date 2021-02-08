package jet

import (
	"testing"
	"time"
)

func TestRawExpression(t *testing.T) {
	assertSerialize(t, Raw("current_database()"), "current_database()")

	var timeT = time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	assertSerialize(t, DateT(timeT), "$1", timeT)
}

func TestTimeLiteral(t *testing.T) {
	assertDebugSerialize(t, Time(11, 5, 30), "'11:05:30'")
	assertDebugSerialize(t, Time(11, 5, 30, 0), "'11:05:30'")
	assertDebugSerialize(t, Time(11, 5, 30, 3*time.Millisecond), "'11:05:30.003'")
	assertDebugSerialize(t, Time(11, 5, 30, 30*time.Millisecond), "'11:05:30.030'")
	assertDebugSerialize(t, Time(11, 5, 30, 300*time.Millisecond), "'11:05:30.300'")
	assertDebugSerialize(t, Time(11, 5, 30, 300*time.Microsecond), "'11:05:30.0003'")
	assertDebugSerialize(t, Time(11, 5, 30, 4*time.Nanosecond), "'11:05:30.000000004'")
}

func TestTimeT(t *testing.T) {
	timeT := time.Date(2000, 1, 1, 11, 40, 20, 124, time.UTC)
	assertDebugSerialize(t, TimeT(timeT), `'2000-01-01 11:40:20.000000124Z'`)
}

func TestTimezLiteral(t *testing.T) {
	assertDebugSerialize(t, Timez(11, 5, 30, 10*time.Nanosecond, "UTC"), "'11:05:30.00000001 UTC'")
	assertDebugSerialize(t, Timez(11, 5, 30, 0, "+1"), "'11:05:30 +1'")
	assertDebugSerialize(t, Timez(11, 5, 30, 3*time.Microsecond, "-7"), "'11:05:30.000003 -7'")
	assertDebugSerialize(t, Timez(11, 5, 30, 30*time.Millisecond, "+8:00"), "'11:05:30.030 +8:00'")
	assertDebugSerialize(t, Timez(11, 5, 30, 300*time.Nanosecond, "America/New_Yor"), "'11:05:30.0000003 America/New_Yor'")
	assertDebugSerialize(t, Timez(11, 5, 30, 3000*time.Nanosecond, "zulu"), "'11:05:30.000003 zulu'")
}

func TestTimestampLiteral(t *testing.T) {
	assertDebugSerialize(t, Timestamp(2011, 1, 8, 11, 5, 30), "'2011-01-08 11:05:30'")
	assertDebugSerialize(t, Timestamp(2011, 2, 7, 11, 5, 30, 0), "'2011-02-07 11:05:30'")
	assertDebugSerialize(t, Timestamp(2011, 3, 6, 11, 5, 30, 3*time.Millisecond), "'2011-03-06 11:05:30.003'")
	assertDebugSerialize(t, Timestamp(2011, 4, 5, 11, 5, 30, 30*time.Millisecond), "'2011-04-05 11:05:30.030'")
	assertDebugSerialize(t, Timestamp(2011, 5, 4, 11, 5, 30, 300*time.Millisecond), "'2011-05-04 11:05:30.300'")
	assertDebugSerialize(t, Timestamp(2011, 6, 3, 11, 5, 30, 3000*time.Microsecond), "'2011-06-03 11:05:30.003'")
}

func TestTimestampzLiteral(t *testing.T) {
	assertDebugSerialize(t, Timestampz(2011, 1, 8, 11, 5, 30, 0, "UTC"), "'2011-01-08 11:05:30 UTC'")
	assertDebugSerialize(t, Timestampz(2011, 2, 7, 11, 5, 30, 0, "PST"), "'2011-02-07 11:05:30 PST'")
	assertDebugSerialize(t, Timestampz(2011, 3, 6, 11, 5, 30, 3, "+4:00"), "'2011-03-06 11:05:30.000000003 +4:00'")
	assertDebugSerialize(t, Timestampz(2011, 4, 5, 11, 5, 30, 30, "-8:00"), "'2011-04-05 11:05:30.00000003 -8:00'")
	assertDebugSerialize(t, Timestampz(2011, 5, 4, 11, 5, 30, 300, "400"), "'2011-05-04 11:05:30.0000003 400'")
	assertDebugSerialize(t, Timestampz(2011, 6, 3, 11, 5, 30, 3000, "zulu"), "'2011-06-03 11:05:30.000003 zulu'")
}

func TestDate(t *testing.T) {
	assertDebugSerialize(t, Date(2019, 8, 8), `'2019-08-08'`)
}

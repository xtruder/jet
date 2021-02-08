package mysql

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/utils"
)

type unitType string

// List of interval unit types for MySQL
const (
	MICROSECOND        unitType = "MICROSECOND"
	SECOND                      = "SECOND"
	MINUTE                      = "MINUTE"
	HOUR                        = "HOUR"
	DAY                         = "DAY"
	WEEK                        = "WEEK"
	MONTH                       = "MONTH"
	QUARTER                     = "QUARTER"
	YEAR                        = "YEAR"
	SECOND_MICROSECOND          = "SECOND_MICROSECOND"
	MINUTE_MICROSECOND          = "MINUTE_MICROSECOND"
	MINUTE_SECOND               = "MINUTE_SECOND"
	HOUR_MICROSECOND            = "HOUR_MICROSECOND"
	HOUR_SECOND                 = "HOUR_SECOND"
	HOUR_MINUTE                 = "HOUR_MINUTE"
	DAY_MICROSECOND             = "DAY_MICROSECOND"
	DAY_SECOND                  = "DAY_SECOND"
	DAY_MINUTE                  = "DAY_MINUTE"
	DAY_HOUR                    = "DAY_HOUR"
	YEAR_MONTH                  = "YEAR_MONTH"
)

// Interval is representation of MySQL interval
type Interval = jet.Interval

// INTERVAL creates new temporal interval.
// In a case of MICROSECOND, SECOND, MINUTE, HOUR, DAY, WEEK, MONTH, QUARTER, YEAR unit type
// value parameter should be number. For example: INTERVAL(1, DAY)
// In a case of other unit types, value should be string with appropriate format.
// For example: INTERVAL("10:08:50", HOUR_SECOND)
func INTERVAL(value interface{}, unitType unitType) Interval {
	switch unitType {
	case MICROSECOND, SECOND, MINUTE, HOUR, DAY, WEEK, MONTH, QUARTER, YEAR:
		if !isNumericType(value) {
			panic("jet: INTERVAL invalid value type. Numeric type expected")
		}
		return INTERVALe(jet.FixedLiteral(value), unitType)
	default:
		strValue, ok := value.(string)

		if !ok {
			panic("jet: INTERNAL invalid value type. String type expected")
		}

		var regexp *regexp.Regexp

		switch unitType {
		case SECOND_MICROSECOND:
			regexp = regexSecondMicrosecond
		case MINUTE_MICROSECOND:
			regexp = regexMinuteMicrosecond
		case MINUTE_SECOND:
			regexp = regexMinuteSecond
		case HOUR_MICROSECOND:
			regexp = regexHourMicrosecond
		case HOUR_SECOND:
			regexp = regexHourSecond
		case HOUR_MINUTE:
			regexp = regexHourMinute
		case DAY_MICROSECOND:
			regexp = regexDayMicrosecond
		case DAY_SECOND:
			regexp = regexDaySecond
		case DAY_MINUTE:
			regexp = regexDayMinute
		case DAY_HOUR:
			regexp = regexDayHour
		case YEAR_MONTH:
			regexp = regexYearMonth
		default:
			panic("jet: INTERVAL invalid unit type")
		}

		if !regexp.MatchString(strValue) {
			panic("jet: INTERVAL invalid format")
		}

		return INTERVALe(jet.Literal(value), unitType)
	}
}

// INTERVALe creates new temporal interval from expresion and unit type.
func INTERVALe(expr Expression, unitType unitType) Interval {
	return jet.NewInterval(jet.ListSerializer{
		Serializers: []jet.Serializer{expr, jet.Raw(string(unitType))},
		Separator:   " ",
	})
}

// INTERVALd temoral interval from time.Duration
func INTERVALd(duration time.Duration) Interval {
	var sign int64 = 1
	if duration < 0 {
		sign = -1
		duration = -duration
	}

	days, hours, minutes, sec, microsec := utils.ExtractDateTimeComponents(duration)

	if days != 0 {
		switch {
		case microsec > 0:
			intervalStr := fmt.Sprintf("%d %02d:%02d:%02d.%06d", sign*days, hours, minutes, sec, microsec)
			return INTERVAL(intervalStr, DAY_MICROSECOND)
		case sec > 0:
			intervalStr := fmt.Sprintf("%d %02d:%02d:%02d", sign*days, hours, minutes, sec)
			return INTERVAL(intervalStr, DAY_SECOND)
		case minutes > 0:
			intervalStr := fmt.Sprintf("%d %02d:%02d", sign*days, hours, minutes)
			return INTERVAL(intervalStr, DAY_MINUTE)
		case hours > 0:
			intervalStr := fmt.Sprintf("%d %02d", sign*days, hours)
			return INTERVAL(intervalStr, DAY_HOUR)
		default:
			return INTERVAL(sign*days, DAY)
		}
	}

	if hours != 0 {
		switch {
		case microsec > 0:
			intervalStr := fmt.Sprintf("%02d:%02d:%02d.%06d", sign*hours, minutes, sec, microsec)
			return INTERVAL(intervalStr, HOUR_MICROSECOND)
		case sec > 0:
			intervalStr := fmt.Sprintf("%02d:%02d:%02d", sign*hours, minutes, sec)
			return INTERVAL(intervalStr, HOUR_SECOND)
		case minutes > 0:
			intervalStr := fmt.Sprintf("%02d:%02d", sign*hours, minutes)
			return INTERVAL(intervalStr, HOUR_MINUTE)
		default:
			return INTERVAL(sign*hours, HOUR)
		}
	}

	if minutes != 0 {
		switch {
		case microsec > 0:
			intervalStr := fmt.Sprintf("%02d:%02d.%06d", sign*minutes, sec, microsec)
			return INTERVAL(intervalStr, MINUTE_MICROSECOND)
		case sec > 0:
			intervalStr := fmt.Sprintf("%02d:%02d", sign*minutes, sec)
			return INTERVAL(intervalStr, MINUTE_SECOND)
		default:
			return INTERVAL(sign*minutes, MINUTE)
		}
	}

	if sec != 0 {
		if microsec > 0 {
			intervalStr := fmt.Sprintf("%02d.%06d", sign*sec, microsec)
			return INTERVAL(intervalStr, SECOND_MICROSECOND)
		}
		return INTERVAL(sign*sec, SECOND)
	}

	return INTERVAL(sign*microsec, MICROSECOND)
}

var (
	regexSecondMicrosecond = regexp.MustCompile(`^-?\d{1,2}\.\d+$`)                //'SECONDS.MICROSECONDS'
	regexMinuteMicrosecond = regexp.MustCompile(`^-?\d{1,2}:\d{2}\.\d+$`)          //'MINUTE:SECONDS.MICROSECONDS'
	regexMinuteSecond      = regexp.MustCompile(`^-?\d{1,2}:\d{2}$`)               //'MINUTE:SECONDS'
	regexHourMicrosecond   = regexp.MustCompile(`^-?\d{1,2}:\d{2}:\d{2}\.\d+$`)    //'HOUR:MINUTE:SECONDS.MICROSECONDS'
	regexHourSecond        = regexp.MustCompile(`^-?\d{1,2}:\d{2}:\d{2}$`)         //'HOUR:MINUTE:SECONDS'
	regexHourMinute        = regexp.MustCompile(`^-?\d{1,2}:\d{2}$`)               //'HOUR:MINUTE'
	regexDayMicrosecond    = regexp.MustCompile(`^-?\d+ \d{1,2}:\d{2}:\d{2}.\d+$`) //'DAY HOUR:MINUTE:SECONDS'
	regexDaySecond         = regexp.MustCompile(`^-?\d+ \d{1,2}:\d{2}:\d{2}$`)     //'DAY HOUR:MINUTE:SECONDS'
	regexDayMinute         = regexp.MustCompile(`^-?\d+ \d{1,2}:\d{2}$`)           //'DAY HOUR:MINUTE'
	regexDayHour           = regexp.MustCompile(`^-?\d+ \d{1,2}$`)                 //'DAY HOUR:MINUTE'
	regexYearMonth         = regexp.MustCompile(`^-?\d+-\d{1,2}$`)                 //'YEAR-MONTH'
)

func isNumericType(value interface{}) bool {
	switch value.(type) {
	case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

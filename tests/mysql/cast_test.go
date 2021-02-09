package mysql

import (
	"testing"
	"time"

	. "github.com/go-jet/jet/v2/mysql"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/table"
)

func TestCast(t *testing.T) {
	query := SELECT(
		CAST(String("test")).AS("CHAR CHARACTER SET utf8").AS("AS1"),
		CAST(String("2011-02-02")).AS_DATE().AS("date1"),
		CAST(String("14:06:10")).AS_TIME().AS("time"),
		CAST(String("2011-02-02 14:06:10")).AS_DATETIME().AS("datetime"),

		CAST(Int(150)).AS_CHAR().AS("char1"),
		CAST(Int(150)).AS_CHAR(30).AS("char2"),

		CAST(Int(5).SUB(Int(10))).AS_SIGNED().AS("signed"),
		CAST(Int(5).ADD(Int(10))).AS_UNSIGNED().AS("unsigned"),
		CAST(String("Some text")).AS_BINARY().AS("binary"),
	).FROM(AllTypes)

	dest := struct {
		As1      string
		Date1    time.Time
		Time     time.Time
		DateTime time.Time
		Char1    string
		Char2    string
		Signed   int
		Unsigned int
		Binary   string
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

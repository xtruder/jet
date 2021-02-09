package mariadb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"

	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-jet/jet/v2/tests/common"
	"github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/table"
	"github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/view"

	. "github.com/go-jet/jet/v2/mysql"
)

func TestAllTypes(t *testing.T) {
	dest := []model.AllTypes{}

	query := AllTypes.
		SELECT(AllTypes.AllColumns).
		LIMIT(2)

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

type AllTypesView model.AllTypes

func TestAllTypesViewSelect(t *testing.T) {
	query := view.AllTypesView.SELECT(view.AllTypesView.AllColumns)
	dest := []AllTypesView{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestUUID(t *testing.T) {
	query := AllTypes.
		SELECT(
			String("dc8daae3-b83b-11e9-8eb4-98ded00c39c6").AS("uuid"),
			String("dc8daae3-b83b-11e9-8eb4-98ded00c39c6").AS("ptr_uuid"),
			Raw("unhex(replace('dc8daae3-b83b-11e9-8eb4-98ded00c39c6','-',''))").AS("bin_uuid"),
		).LIMIT(1)

	dest := struct {
		UUID    uuid.UUID
		PtrUUID *uuid.UUID
		BinUUID uuid.UUID
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestExpressionOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Integer.IS_NULL().AS("result.is_null"),
		AllTypes.DatePtr.IS_NOT_NULL().AS("result.is_not_null"),
		AllTypes.SmallIntPtr.IN(Int(11), Int(22)).AS("result.in"),
		AllTypes.SmallIntPtr.IN(AllTypes.SELECT(AllTypes.Integer)).AS("result.in_select"),
		AllTypes.SmallIntPtr.NOT_IN(Int(11), Int(22), NULL).AS("result.not_in"),
		AllTypes.SmallIntPtr.NOT_IN(AllTypes.SELECT(AllTypes.Integer)).AS("result.not_in_select"),

		Raw("DATABASE()"),
	).LIMIT(2)

	dest := []struct {
		common.ExpressionTestResult `alias:"result.*"`
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestBoolOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Boolean.EQ(AllTypes.BooleanPtr).AS("EQ1"),
		AllTypes.Boolean.EQ(Bool(true)).AS("EQ2"),
		AllTypes.Boolean.NOT_EQ(AllTypes.BooleanPtr).AS("NEq1"),
		AllTypes.Boolean.NOT_EQ(Bool(false)).AS("NEq2"),
		AllTypes.Boolean.IS_DISTINCT_FROM(AllTypes.BooleanPtr).AS("distinct1"),
		AllTypes.Boolean.IS_DISTINCT_FROM(Bool(true)).AS("distinct2"),
		AllTypes.Boolean.IS_NOT_DISTINCT_FROM(AllTypes.BooleanPtr).AS("not_distinct_1"),
		AllTypes.Boolean.IS_NOT_DISTINCT_FROM(Bool(true)).AS("NOTDISTINCT2"),
		AllTypes.Boolean.IS_TRUE().AS("ISTRUE"),
		AllTypes.Boolean.IS_NOT_TRUE().AS("isnottrue"),
		AllTypes.Boolean.IS_FALSE().AS("is_False"),
		AllTypes.Boolean.IS_NOT_FALSE().AS("is not false"),
		AllTypes.Boolean.IS_UNKNOWN().AS("is unknown"),
		AllTypes.Boolean.IS_NOT_UNKNOWN().AS("is_not_unknown"),

		AllTypes.Boolean.AND(AllTypes.Boolean).EQ(AllTypes.Boolean.AND(AllTypes.Boolean)).AS("complex1"),
		AllTypes.Boolean.OR(AllTypes.Boolean).EQ(AllTypes.Boolean.AND(AllTypes.Boolean)).AS("complex2"),
	)

	dest := []struct {
		Eq1          *bool
		Eq2          *bool
		NEq1         *bool
		NEq2         *bool
		Distinct1    *bool
		Distinct2    *bool
		NotDistinct1 *bool
		NotDistinct2 *bool
		IsTrue       *bool
		IsNotTrue    *bool
		IsFalse      *bool
		IsNotFalse   *bool
		IsUnknown    *bool
		IsNotUnknown *bool

		Complex1 *bool
		Complex2 *bool
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestFloatOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Numeric.EQ(AllTypes.Numeric).AS("eq1"),
		AllTypes.Decimal.EQ(Float(12.22)).AS("eq2"),
		AllTypes.Real.EQ(Float(12.12)).AS("eq3"),
		AllTypes.Numeric.IS_DISTINCT_FROM(AllTypes.Numeric).AS("distinct1"),
		AllTypes.Decimal.IS_DISTINCT_FROM(Float(12)).AS("distinct2"),
		AllTypes.Real.IS_DISTINCT_FROM(Float(12.12)).AS("distinct3"),
		AllTypes.Numeric.IS_NOT_DISTINCT_FROM(AllTypes.Numeric).AS("not_distinct1"),
		AllTypes.Decimal.IS_NOT_DISTINCT_FROM(Float(12)).AS("not_distinct2"),
		AllTypes.Real.IS_NOT_DISTINCT_FROM(Float(12.12)).AS("not_distinct3"),
		AllTypes.Numeric.LT(Float(124)).AS("lt1"),
		AllTypes.Numeric.LT(Float(34.56)).AS("lt2"),
		AllTypes.Numeric.GT(Float(124)).AS("gt1"),
		AllTypes.Numeric.GT(Float(34.56)).AS("gt2"),

		TRUNC(AllTypes.Decimal.ADD(AllTypes.Decimal), Int(2)).AS("add1"),
		TRUNC(AllTypes.Decimal.ADD(Float(11.22)), Int(2)).AS("add2"),
		TRUNC(AllTypes.Decimal.SUB(AllTypes.DecimalPtr), Int(2)).AS("sub1"),
		TRUNC(AllTypes.Decimal.SUB(Float(11.22)), Int(2)).AS("sub2"),
		TRUNC(AllTypes.Decimal.MUL(AllTypes.DecimalPtr), Int(2)).AS("mul1"),
		TRUNC(AllTypes.Decimal.MUL(Float(11.22)), Int(2)).AS("mul2"),
		TRUNC(AllTypes.Decimal.DIV(AllTypes.DecimalPtr), Int(2)).AS("div1"),
		TRUNC(AllTypes.Decimal.DIV(Float(11.22)), Int(2)).AS("div2"),
		TRUNC(AllTypes.Decimal.MOD(AllTypes.DecimalPtr), Int(2)).AS("mod1"),
		TRUNC(AllTypes.Decimal.MOD(Float(11.22)), Int(2)).AS("mod2"),
		TRUNC(AllTypes.Decimal.POW(AllTypes.DecimalPtr), Int(2)).AS("pow1"),
		TRUNC(AllTypes.Decimal.POW(Float(2.1)), Int(2)).AS("pow2"),

		TRUNC(ABSf(AllTypes.Decimal), Int(2)).AS("abs"),
		TRUNC(POWER(AllTypes.Decimal, Float(2.1)), Int(2)).AS("power"),
		TRUNC(SQRT(AllTypes.Decimal), Int(2)).AS("sqrt"),
		TRUNC(CBRT(AllTypes.Decimal), Int(2)).AS("cbrt"),

		CEIL(AllTypes.Real).AS("ceil"),
		FLOOR(AllTypes.Real).AS("floor"),
		ROUND(AllTypes.Decimal).AS("round1"),
		ROUND(AllTypes.Decimal, Int(2)).AS("round2"),

		SIGN(AllTypes.Real).AS("sign"),
		TRUNC(AllTypes.Decimal, Int(1)).AS("trunc"),
	).LIMIT(2)

	dest := []struct {
		common.FloatExpressionTestResult `alias:"."`
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestIntegerOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.BigInt,
		AllTypes.BigIntPtr,
		AllTypes.SmallInt,
		AllTypes.SmallIntPtr,

		AllTypes.BigInt.EQ(AllTypes.BigInt).AS("eq1"),
		AllTypes.BigInt.EQ(Int(12)).AS("eq2"),

		AllTypes.BigInt.NOT_EQ(AllTypes.BigIntPtr).AS("neq1"),
		AllTypes.BigInt.NOT_EQ(Int(12)).AS("neq2"),

		AllTypes.BigInt.IS_DISTINCT_FROM(AllTypes.BigInt).AS("distinct1"),
		AllTypes.BigInt.IS_DISTINCT_FROM(Int(12)).AS("distinct2"),

		AllTypes.BigInt.IS_NOT_DISTINCT_FROM(AllTypes.BigInt).AS("not distinct1"),
		AllTypes.BigInt.IS_NOT_DISTINCT_FROM(Int(12)).AS("not distinct2"),

		AllTypes.BigInt.LT(AllTypes.BigIntPtr).AS("lt1"),
		AllTypes.BigInt.LT(Int(65)).AS("lt2"),

		AllTypes.BigInt.LT_EQ(AllTypes.BigIntPtr).AS("lte1"),
		AllTypes.BigInt.LT_EQ(Int(65)).AS("lte2"),

		AllTypes.BigInt.GT(AllTypes.BigIntPtr).AS("gt1"),
		AllTypes.BigInt.GT(Int(65)).AS("gt2"),

		AllTypes.BigInt.GT_EQ(AllTypes.BigIntPtr).AS("gte1"),
		AllTypes.BigInt.GT_EQ(Int(65)).AS("gte2"),

		AllTypes.BigInt.ADD(AllTypes.BigInt).AS("add1"),
		AllTypes.BigInt.ADD(Int(11)).AS("add2"),

		AllTypes.BigInt.SUB(AllTypes.BigInt).AS("sub1"),
		AllTypes.BigInt.SUB(Int(11)).AS("sub2"),

		AllTypes.BigInt.MUL(AllTypes.BigInt).AS("mul1"),
		AllTypes.BigInt.MUL(Int(11)).AS("mul2"),

		AllTypes.BigInt.DIV(AllTypes.BigInt).AS("div1"),
		AllTypes.BigInt.DIV(Int(11)).AS("div2"),

		AllTypes.BigInt.MOD(AllTypes.BigInt).AS("mod1"),
		AllTypes.BigInt.MOD(Int(11)).AS("mod2"),

		AllTypes.SmallInt.POW(AllTypes.SmallInt.DIV(Int(3))).AS("pow1"),
		AllTypes.SmallInt.POW(Int(6)).AS("pow2"),

		AllTypes.SmallInt.BIT_AND(AllTypes.SmallInt).AS("bit_and1"),
		AllTypes.SmallInt.BIT_AND(AllTypes.SmallInt).AS("bit_and2"),

		AllTypes.SmallInt.BIT_OR(AllTypes.SmallInt).AS("bit or 1"),
		AllTypes.SmallInt.BIT_OR(Int(22)).AS("bit or 2"),

		AllTypes.SmallInt.BIT_XOR(AllTypes.SmallInt).AS("bit xor 1"),
		AllTypes.SmallInt.BIT_XOR(Int(11)).AS("bit xor 2"),

		BIT_NOT(Int(-1).MUL(AllTypes.SmallInt)).AS("bit_not_1"),
		BIT_NOT(Int(-1).MUL(Int(11))).AS("bit_not_2"),

		AllTypes.SmallInt.BIT_SHIFT_LEFT(AllTypes.SmallInt.DIV(Int(2))).AS("bit shift left 1"),
		AllTypes.SmallInt.BIT_SHIFT_LEFT(Int(4)).AS("bit shift left 2"),

		AllTypes.SmallInt.BIT_SHIFT_RIGHT(AllTypes.SmallInt.DIV(Int(5))).AS("bit shift right 1"),
		AllTypes.SmallInt.BIT_SHIFT_RIGHT(Int(1)).AS("bit shift right 2"),

		ABSi(AllTypes.BigInt).AS("abs"),
		SQRT(ABSi(AllTypes.BigInt)).AS("sqrt"),
		CBRT(ABSi(AllTypes.BigInt)).AS("cbrt"),
	).LIMIT(2)

	dest := []struct {
		common.AllTypesIntegerExpResult `alias:"."`
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestStringOperators(t *testing.T) {
	projectionList := []Projection{
		AllTypes.Text.EQ(AllTypes.Char),
		AllTypes.Text.EQ(String("Text")),
		AllTypes.Text.NOT_EQ(AllTypes.VarCharPtr),
		AllTypes.Text.NOT_EQ(String("Text")),
		AllTypes.Text.GT(AllTypes.Text),
		AllTypes.Text.GT(String("Text")),
		AllTypes.Text.GT_EQ(AllTypes.TextPtr),
		AllTypes.Text.GT_EQ(String("Text")),
		AllTypes.Text.LT(AllTypes.Char),
		AllTypes.Text.LT(String("Text")),
		AllTypes.Text.LT_EQ(AllTypes.VarCharPtr),
		AllTypes.Text.LT_EQ(String("Text")),
		AllTypes.Text.CONCAT(String("text2")),
		AllTypes.Text.CONCAT(Int(11)),
		AllTypes.Text.LIKE(String("abc")),
		AllTypes.Text.NOT_LIKE(String("_b_")),
		AllTypes.Text.REGEXP_LIKE(String("aba")),
		AllTypes.Text.REGEXP_LIKE(String("aba"), false),
		String("ABA").REGEXP_LIKE(String("aba"), true),
		AllTypes.Text.NOT_REGEXP_LIKE(String("aba")),
		AllTypes.Text.NOT_REGEXP_LIKE(String("aba"), false),
		String("ABA").NOT_REGEXP_LIKE(String("aba"), true),

		BIT_LENGTH(AllTypes.Text),
		CHAR_LENGTH(AllTypes.Char),
		OCTET_LENGTH(AllTypes.Text),
		LOWER(AllTypes.VarCharPtr),
		UPPER(AllTypes.Char),
		LTRIM(AllTypes.VarCharPtr),
		RTRIM(AllTypes.VarCharPtr),
		CONCAT(String("string1"), Int(1), Float(11.12)),
		CONCAT_WS(String("string1"), Int(1), Float(11.12)),
		FORMAT(String("Hello %s, %1$s"), String("World")),
		LEFT(String("abcde"), Int(2)),
		RIGHT(String("abcde"), Int(2)),
		LENGTH(String("jose")),
		LPAD(String("Hi"), Int(5), String("xy")),
		RPAD(String("Hi"), Int(5), String("xy")),
		MD5(AllTypes.VarCharPtr),
		REPEAT(AllTypes.Text, Int(33)),
		REPLACE(AllTypes.Char, String("BA"), String("AB")),
		REVERSE(AllTypes.VarCharPtr),
		SUBSTR(AllTypes.CharPtr, Int(3)),
		SUBSTR(AllTypes.CharPtr, Int(3), Int(2)),
	}

	query := SELECT(projectionList[0], projectionList[1:]...).FROM(AllTypes)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestTimeExpressions(t *testing.T) {
	query := AllTypes.SELECT(
		Time(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC).Clock()),

		AllTypes.Time.EQ(AllTypes.Time),
		AllTypes.Time.EQ(Time(23, 6, 6)),
		AllTypes.Time.EQ(Time(22, 6, 6, 11*time.Millisecond)),
		AllTypes.Time.EQ(Time(21, 6, 6, 11111*time.Microsecond)),

		AllTypes.TimePtr.NOT_EQ(AllTypes.Time),
		AllTypes.TimePtr.NOT_EQ(Time(20, 16, 6)),

		AllTypes.Time.IS_DISTINCT_FROM(AllTypes.Time),
		AllTypes.Time.IS_DISTINCT_FROM(Time(19, 26, 6)),

		AllTypes.Time.IS_NOT_DISTINCT_FROM(AllTypes.Time),
		AllTypes.Time.IS_NOT_DISTINCT_FROM(Time(18, 36, 6)),

		AllTypes.Time.LT(AllTypes.Time),
		AllTypes.Time.LT(Time(17, 46, 6)),

		AllTypes.Time.LT_EQ(AllTypes.Time),
		AllTypes.Time.LT_EQ(Time(16, 56, 56)),

		AllTypes.Time.GT(AllTypes.Time),
		AllTypes.Time.GT(Time(15, 16, 46)),

		AllTypes.Time.GT_EQ(AllTypes.Time),
		AllTypes.Time.GT_EQ(Time(14, 26, 36)),

		AllTypes.Time.ADD(INTERVAL(10, MINUTE)),
		AllTypes.Time.ADD(INTERVALe(AllTypes.Integer, MINUTE)),
		AllTypes.Time.ADD(INTERVALd(3*time.Hour)),

		AllTypes.Time.SUB(INTERVAL(20, MINUTE)),
		AllTypes.Time.SUB(INTERVALe(AllTypes.SmallInt, MINUTE)),
		AllTypes.Time.SUB(INTERVALd(3*time.Minute)),

		AllTypes.Time.ADD(INTERVAL(20, MINUTE)).SUB(INTERVAL(11, HOUR)),

		CURRENT_TIME(),
		CURRENT_TIME(3),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestDateExpressions(t *testing.T) {
	query := AllTypes.SELECT(
		Date(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC).Date()),

		AllTypes.Date.EQ(AllTypes.Date),
		AllTypes.Date.EQ(Date(2019, 6, 6)),

		AllTypes.DatePtr.NOT_EQ(AllTypes.Date),
		AllTypes.DatePtr.NOT_EQ(Date(2019, 1, 6)),

		AllTypes.Date.IS_DISTINCT_FROM(AllTypes.Date),
		AllTypes.Date.IS_DISTINCT_FROM(Date(2019, 2, 6)),

		AllTypes.Date.IS_NOT_DISTINCT_FROM(AllTypes.Date),
		AllTypes.Date.IS_NOT_DISTINCT_FROM(Date(2019, 3, 6)),

		AllTypes.Date.LT(AllTypes.Date),
		AllTypes.Date.LT(Date(2019, 4, 6)),

		AllTypes.Date.LT_EQ(AllTypes.Date),
		AllTypes.Date.LT_EQ(Date(2019, 5, 5)),

		AllTypes.Date.GT(AllTypes.Date),
		AllTypes.Date.GT(Date(2019, 1, 4)),

		AllTypes.Date.GT_EQ(AllTypes.Date),
		AllTypes.Date.GT_EQ(Date(2019, 2, 3)),

		AllTypes.Date.ADD(INTERVAL("10:20.000100", MINUTE_MICROSECOND)),
		AllTypes.Date.ADD(INTERVALe(AllTypes.BigInt, MINUTE)),
		AllTypes.Date.ADD(INTERVALd(15*time.Hour)),

		AllTypes.Date.SUB(INTERVAL(20, MINUTE)),
		AllTypes.Date.SUB(INTERVALe(AllTypes.SmallInt, MINUTE)),
		AllTypes.Date.SUB(INTERVALd(3*time.Minute)),

		CURRENT_DATE(),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestDateTimeExpressions(t *testing.T) {
	dateTime := DateTime(2019, 6, 6, 10, 2, 46)

	query := AllTypes.SELECT(
		AllTypes.DateTime.EQ(AllTypes.DateTime),
		AllTypes.DateTime.EQ(dateTime),

		AllTypes.DateTimePtr.NOT_EQ(AllTypes.DateTime),
		AllTypes.DateTimePtr.NOT_EQ(DateTime(2019, 6, 6, 10, 2, 46, 100*time.Millisecond)),

		AllTypes.DateTime.IS_DISTINCT_FROM(AllTypes.DateTime),
		AllTypes.DateTime.IS_DISTINCT_FROM(dateTime),

		AllTypes.DateTime.IS_NOT_DISTINCT_FROM(AllTypes.DateTime),
		AllTypes.DateTime.IS_NOT_DISTINCT_FROM(dateTime),

		AllTypes.DateTime.LT(AllTypes.DateTime),
		AllTypes.DateTime.LT(dateTime),

		AllTypes.DateTime.LT_EQ(AllTypes.DateTime),
		AllTypes.DateTime.LT_EQ(dateTime),

		AllTypes.DateTime.GT(AllTypes.DateTime),
		AllTypes.DateTime.GT(dateTime),

		AllTypes.DateTime.GT_EQ(AllTypes.DateTime),
		AllTypes.DateTime.GT_EQ(dateTime),

		AllTypes.DateTime.ADD(INTERVAL("05:10:20.000100", HOUR_MICROSECOND)),
		AllTypes.DateTime.ADD(INTERVALe(AllTypes.BigInt, HOUR)),
		AllTypes.DateTime.ADD(INTERVALd(2*time.Hour)),

		AllTypes.DateTime.SUB(INTERVAL("05:10:20.000100", HOUR_MICROSECOND)),
		AllTypes.DateTime.SUB(INTERVALe(AllTypes.IntegerPtr, HOUR)),
		AllTypes.DateTime.SUB(INTERVALd(3*time.Hour)),

		NOW(),
		NOW(1),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestTimestampExpressions(t *testing.T) {
	timestamp := Timestamp(2019, 6, 6, 10, 2, 46)

	query := AllTypes.SELECT(
		AllTypes.Timestamp.EQ(AllTypes.Timestamp),
		AllTypes.Timestamp.EQ(timestamp),

		AllTypes.TimestampPtr.NOT_EQ(AllTypes.Timestamp),
		AllTypes.TimestampPtr.NOT_EQ(Timestamp(2019, 6, 6, 10, 2, 46, 100*time.Millisecond)),

		AllTypes.Timestamp.IS_DISTINCT_FROM(AllTypes.Timestamp),
		AllTypes.Timestamp.IS_DISTINCT_FROM(timestamp),

		AllTypes.Timestamp.IS_NOT_DISTINCT_FROM(AllTypes.Timestamp),
		AllTypes.Timestamp.IS_NOT_DISTINCT_FROM(timestamp),

		AllTypes.Timestamp.LT(AllTypes.Timestamp),
		AllTypes.Timestamp.LT(timestamp),

		AllTypes.Timestamp.LT_EQ(AllTypes.Timestamp),
		AllTypes.Timestamp.LT_EQ(timestamp),

		AllTypes.Timestamp.GT(AllTypes.Timestamp),
		AllTypes.Timestamp.GT(timestamp),

		AllTypes.Timestamp.GT_EQ(AllTypes.Timestamp),
		AllTypes.Timestamp.GT_EQ(timestamp),

		AllTypes.Timestamp.ADD(INTERVAL("05:10:20.000100", HOUR_MICROSECOND)),
		AllTypes.Timestamp.ADD(INTERVALe(AllTypes.BigInt, HOUR)),
		AllTypes.Timestamp.ADD(INTERVALd(2*time.Hour)),

		AllTypes.Timestamp.SUB(INTERVAL("05:10:20.000100", HOUR_MICROSECOND)),
		AllTypes.Timestamp.SUB(INTERVALe(AllTypes.IntegerPtr, HOUR)),
		AllTypes.Timestamp.SUB(INTERVALd(3*time.Hour)),

		CURRENT_TIMESTAMP(),
		CURRENT_TIMESTAMP(2),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestTimeLiterals(t *testing.T) {
	loc := time.FixedZone("GMT+1", 60*60)

	timeT := time.Date(2009, 11, 17, 20, 34, 58, 351387237, loc)

	query := SELECT(
		Date(timeT.Date()).AS("date"),
		DateT(timeT).AS("dateT"),
		Time(timeT.Clock()).AS("time"),
		TimeT(timeT).AS("timeT"),
		DateTimeT(timeT).AS("datetime"),
		Timestamp(2019, 8, 6, 10, 10, 30, 123456*time.Millisecond).AS("timestamp"),
		TimestampT(timeT).AS("timestampT"),
	).FROM(AllTypes).LIMIT(1)

	dest := struct {
		Date       time.Time
		DateT      time.Time
		Time       time.Time
		TimeT      time.Time
		DateTime   time.Time
		Timestamp  time.Time
		TimestampT time.Time
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestINTERVAL(t *testing.T) {
	query := SELECT(
		Date(2000, 2, 10).ADD(INTERVAL(1, MICROSECOND)).
			EQ(Timestamp(2000, 2, 10, 0, 0, 0, 1*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVAL(2, SECOND)),
		Date(2000, 2, 10).ADD(INTERVAL(3, MINUTE)),
		Date(2000, 2, 10).SUB(INTERVAL(4, HOUR)),
		Date(2000, 2, 10).ADD(INTERVAL(5, DAY)),
		Date(2000, 2, 10).SUB(INTERVAL(6, MONTH)),
		Date(2000, 2, 10).ADD(INTERVAL(7, YEAR)),
		Date(2000, 2, 10).ADD(INTERVAL(-7, YEAR)),
		Date(2000, 2, 10).ADD(INTERVAL("20.0000100", SECOND_MICROSECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("02:20.0000100", MINUTE_MICROSECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("11:02:20.0000100", HOUR_MICROSECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("100 11:02:20.0000100", DAY_MICROSECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("11:02", MINUTE_SECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("11:02:20", HOUR_SECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("11:02", HOUR_MINUTE)),
		Date(2000, 2, 10).SUB(INTERVAL("11 02:03:04", DAY_SECOND)),
		Date(2000, 2, 10).SUB(INTERVAL("11 02:03", DAY_MINUTE)),
		Date(2000, 2, 10).SUB(INTERVAL("11 2", DAY_HOUR)),
		Date(2000, 2, 10).SUB(INTERVAL("2000-2", YEAR_MONTH)),

		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, MICROSECOND)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, SECOND)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, MINUTE)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, HOUR)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, DAY)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, WEEK)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, MONTH)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, QUARTER)),
		Date(2000, 2, 10).SUB(INTERVALe(AllTypes.IntegerPtr, YEAR)),

		Date(2000, 2, 10).SUB(INTERVALd(3*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVALd(-3*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVALd(3*time.Second)),
		Date(2000, 2, 10).SUB(INTERVALd(3*time.Second+4*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVALd(3*time.Minute+4*time.Second+5*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVALd(3*time.Hour+4*time.Minute+5*time.Second+6*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVALd(2*24*time.Hour+3*time.Hour+4*time.Minute+5*time.Second+6*time.Microsecond)),
		Date(2000, 2, 10).SUB(INTERVALd(2*24*time.Hour+3*time.Hour+4*time.Minute+5*time.Second)),
		Date(2000, 2, 10).SUB(INTERVALd(2*24*time.Hour+3*time.Hour+4*time.Minute)),
		Date(2000, 2, 10).SUB(INTERVALd(2*24*time.Hour+3*time.Hour)),
		Date(2000, 2, 10).SUB(INTERVALd(2*24*time.Hour)),
		Date(2000, 2, 10).SUB(INTERVALd(3*time.Hour)),
		Date(2000, 2, 10).SUB(INTERVALd(1*time.Hour+2*time.Minute+3*time.Second+345*time.Microsecond)),
	).FROM(AllTypes)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestAllTypesInsert(t *testing.T) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	stmt := AllTypes.INSERT(AllTypes.AllColumns).
		MODEL(toInsert)

	assertExec(t, stmt, tx, 1)

	dest := model.AllTypes{}
	err = AllTypes.SELECT(AllTypes.AllColumns).
		WHERE(AllTypes.BigInt.EQ(Int(toInsert.BigInt))).
		Query(tx, &dest)

	require.NoError(t, err)
	require.Equal(t, toInsert.TinyInt, dest.TinyInt)

	err = tx.Rollback()
	require.NoError(t, err)
}

func TestAllTypesInsertOnDuplicateKeyUpdate(t *testing.T) {
	tx, err := db.Begin()
	require.NoError(t, err)

	toInsert := model.AllTypes{
		Boolean:   true,
		Integer:   124,
		Float:     45.67,
		Blob:      []byte("blob"),
		Text:      "text",
		JSON:      qrm.JSON("{}"),
		Time:      time.Now(),
		Timestamp: time.Now(),
		Date:      time.Now(),
	}

	stmt := AllTypes.INSERT(
		AllTypes.Boolean,
		AllTypes.Integer,
		AllTypes.Float,
		AllTypes.Blob,
		AllTypes.Text,
		AllTypes.JSON,
		AllTypes.Time,
		AllTypes.Timestamp,
		AllTypes.Date,
	).
		MODEL(toInsert).
		ON_DUPLICATE_KEY_UPDATE(
			AllTypes.Boolean.SET(Bool(false)),
			AllTypes.Integer.SET(Int(4)),
			AllTypes.Float.SET(Float(0.67)),
			AllTypes.Text.SET(String("new text")),
			AllTypes.Time.SET(TimeT(time.Now())),
			AllTypes.Timestamp.SET(TimestampT(time.Now())),
			AllTypes.Date.SET(DateT(time.Now())),
		)

	_, err = stmt.Exec(tx)
	require.NoError(t, err)

	err = tx.Rollback()
	require.NoError(t, err)
}

var toInsert = model.AllTypes{
	Boolean:       false,
	BooleanPtr:    testutils.Ptr(true).(*bool),
	TinyInt:       1,
	UTinyInt:      2,
	SmallInt:      3,
	USmallInt:     4,
	MediumInt:     5,
	UMediumInt:    6,
	Integer:       7,
	UInteger:      8,
	BigInt:        9,
	UBigInt:       1122334455,
	TinyIntPtr:    testutils.Ptr(int8(11)).(*int8),
	UTinyIntPtr:   testutils.Ptr(uint8(22)).(*uint8),
	SmallIntPtr:   testutils.Ptr(int16(33)).(*int16),
	USmallIntPtr:  testutils.Ptr(uint16(44)).(*uint16),
	MediumIntPtr:  testutils.Ptr(int32(55)).(*int32),
	UMediumIntPtr: testutils.Ptr(uint32(66)).(*uint32),
	IntegerPtr:    testutils.Ptr(int32(77)).(*int32),
	UIntegerPtr:   testutils.Ptr(uint32(88)).(*uint32),
	BigIntPtr:     testutils.Ptr(int64(99)).(*int64),
	UBigIntPtr:    testutils.Ptr(uint64(111)).(*uint64),
	Decimal:       11.22,
	DecimalPtr:    testutils.Ptr(33.44).(*float64),
	Numeric:       55.66,
	NumericPtr:    testutils.Ptr(77.88).(*float64),
	Float:         99.00,
	FloatPtr:      testutils.Ptr(11.22).(*float64),
	Double:        33.44,
	DoublePtr:     testutils.Ptr(55.66).(*float64),
	Real:          77.88,
	RealPtr:       testutils.Ptr(99.00).(*float64),
	Bit:           "1",
	BitPtr:        testutils.Ptr("0").(*string),
	Time:          time.Now(),
	TimePtr:       testutils.Ptr(time.Now()).(*time.Time),
	Date:          time.Now(),
	DatePtr:       testutils.Ptr(time.Now()).(*time.Time),
	DateTime:      time.Now(),
	DateTimePtr:   testutils.Ptr(time.Now()).(*time.Time),
	Timestamp:     time.Now(),
	//TimestampPtr:  testutils.TimePtr(time.Now()), // TODO: build fails for MariaDB
	Year:         2000,
	YearPtr:      testutils.Ptr(int16(001)).(*int16),
	Char:         "abcd",
	CharPtr:      testutils.Ptr("absd").(*string),
	VarChar:      "abcd",
	VarCharPtr:   testutils.Ptr("absd").(*string),
	Binary:       []byte("1010"),
	BinaryPtr:    testutils.Ptr([]byte("100001")).(*[]byte),
	VarBinary:    []byte("1010"),
	VarBinaryPtr: testutils.Ptr([]byte("100001")).(*[]byte),
	Blob:         []byte("large file"),
	BlobPtr:      testutils.Ptr([]byte("very large file")).(*[]byte),
	Text:         "some text",
	TextPtr:      testutils.Ptr("text").(*string),
	Enum:         model.AllTypesEnum_Value1,
	JSON:         qrm.JSON("{}"),
	JSONPtr:      testutils.Ptr(qrm.JSON(`{"a": 1}`)).(*qrm.JSON),
}

func TestReservedWord(t *testing.T) {
	stmt := SELECT(User.AllColumns).
		FROM(User)

	dest := []model.User{}

	assertStatementRecordSQL(t, stmt)
	assertQueryRecordValues(t, stmt, &dest)
}

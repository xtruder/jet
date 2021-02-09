package postgres

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"

	"github.com/go-jet/jet/v2/internal/testutils"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-jet/jet/v2/tests/common"
	"github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/table"
	"github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/view"
	"github.com/google/uuid"
)

func TestAllTypesSelect(t *testing.T) {
	dest := []model.AllTypes{}

	err := AllTypes.SELECT(AllTypes.AllColumns).LIMIT(2).Query(db, &dest)

	require.NoError(t, err)
	require.EqualValues(t, testparrot.RecordNext(t, dest[:2]), dest[:2])
}

type AllTypesView model.AllTypes

func TestAllTypesViewSelect(t *testing.T) {
	dest := []AllTypesView{}

	err := view.AllTypesView.SELECT(view.AllTypesView.AllColumns).Query(db, &dest)
	require.NoError(t, err)

	require.EqualValues(t, testparrot.RecordNext(t, dest[:2]), dest[:2])
}

func TestAllTypesInsertModel(t *testing.T) {
	query := AllTypes.INSERT(AllTypes.AllColumns).
		MODEL(allTypesRow0).
		MODEL(&allTypesRow1).
		RETURNING(AllTypes.AllColumns)

	dest := []model.AllTypes{}

	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, 2)
	require.EqualValues(t, dest[0], allTypesRow0)
	require.EqualValues(t, dest[1], allTypesRow1)
}

func TestAllTypesInsertQuery(t *testing.T) {
	query := AllTypes.INSERT(AllTypes.AllColumns).
		QUERY(
			AllTypes.
				SELECT(AllTypes.AllColumns).
				LIMIT(2),
		).
		RETURNING(AllTypes.AllColumns)

	dest := []model.AllTypes{}
	err := query.Query(db, &dest)

	require.NoError(t, err)
	require.Equal(t, len(dest), 2)
	require.EqualValues(t, dest[0], allTypesRow0)
	require.EqualValues(t, dest[1], allTypesRow1)
}

func TestAllTypesFromSubQuery(t *testing.T) {

	subQuery := SELECT(AllTypes.AllColumns).
		FROM(AllTypes).
		AsTable("allTypesSubQuery")

	mainQuery := SELECT(subQuery.AllColumns()).
		FROM(subQuery).
		LIMIT(2)

	sql := mainQuery.String()
	require.Equal(t, testparrot.RecordNext(t, sql), sql)

	dest := []model.AllTypes{}
	err := mainQuery.Query(db, &dest)

	require.NoError(t, err)
	require.Equal(t, len(dest), 2)
}

type expressionOperatorsResult struct {
	common.ExpressionTestResult `alias:"result.*"`
}

func TestExpressionOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Integer.IS_NULL().AS("result.is_null"),
		AllTypes.DatePtr.IS_NOT_NULL().AS("result.is_not_null"),
		AllTypes.SmallIntPtr.IN(Int(11), Int(22)).AS("result.in"),
		AllTypes.SmallIntPtr.IN(AllTypes.SELECT(AllTypes.Integer)).AS("result.in_select"),
		AllTypes.SmallIntPtr.NOT_IN(Int(11), Int(22), NULL).AS("result.not_in"),
		AllTypes.SmallIntPtr.NOT_IN(AllTypes.SELECT(AllTypes.Integer)).AS("result.not_in_select"),
	).LIMIT(2)

	dest := []expressionOperatorsResult{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestExpressionCast(t *testing.T) {
	query := AllTypes.SELECT(
		CAST(Int(150)).AS_CHAR(12).AS("char12"),
		CAST(String("TRUE")).AS_BOOL(),
		CAST(String("111")).AS_SMALLINT(),
		CAST(String("111")).AS_INTEGER(),
		CAST(String("111")).AS_BIGINT(),
		CAST(String("11.23")).AS_NUMERIC(30, 10),
		CAST(String("11.23")).AS_NUMERIC(30),
		CAST(String("11.23")).AS_NUMERIC(),
		CAST(String("11.23")).AS_REAL(),
		CAST(String("11.23")).AS_DOUBLE(),
		CAST(Int(234)).AS_TEXT(),
		CAST(String("1/8/1999")).AS_DATE(),
		CAST(String("04:05:06.789")).AS_TIME(),
		CAST(String("04:05:06 PST")).AS_TIMEZ(),
		CAST(String("1999-01-08 04:05:06")).AS_TIMESTAMP(),
		CAST(String("January 8 04:05:06 1999 PST")).AS_TIMESTAMPZ(),
		CAST(String("04:05:06")).AS_INTERVAL(),
		CAST(JSON(map[string]string{"key": "value"})).AS_TEXT(),

		TO_CHAR(AllTypes.Timestamp, String("HH12:MI:SS")),
		TO_CHAR(AllTypes.Integer, String("999")),
		TO_CHAR(AllTypes.DoublePrecision, String("999D9")),
		TO_CHAR(AllTypes.Numeric, String("999D99S")),

		TO_DATE(String("05 Dec 2000"), String("DD Mon YYYY")),
		TO_NUMBER(String("12,454"), String("99G999D9S")),
		TO_TIMESTAMP(String("05 Dec 2000"), String("DD Mon YYYY")),

		COALESCE(AllTypes.IntegerPtr, AllTypes.SmallIntPtr, NULL, Int(11)),
		NULLIF(AllTypes.Text, String("(none)")),
		GREATEST(AllTypes.Numeric, AllTypes.NumericPtr),
		LEAST(AllTypes.Numeric, AllTypes.NumericPtr),
		Raw("current_database()"),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestStringOperators(t *testing.T) {
	query := AllTypes.SELECT(
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
		AllTypes.Text.LT_EQ(AllTypes.VarChar),
		AllTypes.Text.LT_EQ(String("Text")),
		AllTypes.Text.CONCAT(String("text2")),
		AllTypes.Text.CONCAT(Int(11)),
		AllTypes.Text.LIKE(String("abc")),
		AllTypes.Text.NOT_LIKE(String("_b_")),
		AllTypes.Text.REGEXP_LIKE(String("^t")),
		AllTypes.Text.REGEXP_LIKE(String("^t"), true),
		AllTypes.Text.NOT_REGEXP_LIKE(String("^t")),
		AllTypes.Text.NOT_REGEXP_LIKE(String("^t"), true),

		BIT_LENGTH(String("length")),
		CHAR_LENGTH(AllTypes.Char),
		CHAR_LENGTH(String("length")),
		OCTET_LENGTH(AllTypes.Text),
		OCTET_LENGTH(String("length")),
		LOWER(AllTypes.VarCharPtr),
		LOWER(String("length")),
		UPPER(AllTypes.Char),
		UPPER(String("upper")),
		BTRIM(AllTypes.VarChar),
		BTRIM(String("btrim")),
		BTRIM(AllTypes.VarChar, String("AA")),
		BTRIM(String("btrim"), String("AA")),
		LTRIM(AllTypes.VarChar),
		LTRIM(String("ltrim")),
		LTRIM(AllTypes.VarChar, String("A")),
		LTRIM(String("Ltrim"), String("A")),
		RTRIM(String("rtrim")),
		RTRIM(AllTypes.VarChar, String("B")),
		CHR(Int(65)),
		CONCAT(AllTypes.VarCharPtr, AllTypes.VarCharPtr, String("aaa"), Int(1)),
		CONCAT(Bool(false), Int(1), Float(22.2), String("test test")),
		CONCAT_WS(String("string1"), Int(1), Float(11.22), String("bytea"), Bool(false)), //Float(11.12)),
		CONVERT(String("bytea"), String("UTF8"), String("LATIN1")),
		CONVERT(AllTypes.Bytea, String("UTF8"), String("LATIN1")),
		CONVERT_FROM(String("text_in_utf8"), String("UTF8")),
		CONVERT_TO(String("text_in_utf8"), String("UTF8")),
		//ENCODE(String("123\000\001"), String("base64")),
		DECODE(String("MTIzAAE="), String("base64")),
		FORMAT(String("Hello %s, %1$s"), String("World")),
		INITCAP(String("hi THOMAS")),
		LEFT(String("abcde"), Int(2)),
		RIGHT(String("abcde"), Int(2)),
		LENGTH(String("jose")),
		LENGTH(String("jose"), String("UTF8")),
		LPAD(String("Hi"), Int(5)),
		LPAD(String("Hi"), Int(5), String("xy")),
		RPAD(String("Hi"), Int(5)),
		RPAD(String("Hi"), Int(5), String("xy")),
		MD5(AllTypes.VarChar),
		REPEAT(AllTypes.Text, Int(33)),
		REPLACE(AllTypes.Char, String("BA"), String("AB")),
		REVERSE(AllTypes.VarChar),
		STRPOS(AllTypes.Text, String("A")),
		SUBSTR(AllTypes.Char, Int(3)),
		SUBSTR(AllTypes.CharPtr, Int(3), Int(2)),
		TO_HEX(AllTypes.IntegerPtr),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

type boolOperatorsResult struct {
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
	).LIMIT(2)

	dest := []boolOperatorsResult{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

type floatOperatorsResult struct {
	common.FloatExpressionTestResult `alias:"."`
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
		TRUNC(CAST(CBRT(AllTypes.Decimal)).AS_DECIMAL(), Int(2)).AS("cbrt"),

		CEIL(AllTypes.Real).AS("ceil"),
		FLOOR(AllTypes.Real).AS("floor"),
		ROUND(AllTypes.Decimal).AS("round1"),
		ROUND(AllTypes.Decimal, AllTypes.Integer).AS("round2"),

		SIGN(AllTypes.Real).AS("sign"),
		TRUNC(AllTypes.Decimal, Int(1)).AS("trunc"),
	).LIMIT(2)

	dest := []floatOperatorsResult{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

type integerOperatorsResult struct {
	common.AllTypesIntegerExpResult `alias:"."`
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
		BIT_NOT(Int(-11)).AS("bit_not_2"),

		AllTypes.SmallInt.BIT_SHIFT_LEFT(AllTypes.SmallInt.DIV(Int(2))).AS("bit shift left 1"),
		AllTypes.SmallInt.BIT_SHIFT_LEFT(Int(4)).AS("bit shift left 2"),

		AllTypes.SmallInt.BIT_SHIFT_RIGHT(AllTypes.SmallInt.DIV(Int(5))).AS("bit shift right 1"),
		AllTypes.SmallInt.BIT_SHIFT_RIGHT(Int(1)).AS("bit shift right 2"),

		ABSi(AllTypes.BigInt).AS("abs"),
		SQRT(ABSi(AllTypes.BigInt)).AS("sqrt"),
		CBRT(ABSi(AllTypes.BigInt)).AS("cbrt"),
	).LIMIT(2)

	dest := []integerOperatorsResult{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestTimeExpression(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Time.EQ(AllTypes.Time),
		AllTypes.Time.EQ(Time(23, 6, 6, 1)),
		AllTypes.Timez.EQ(AllTypes.TimezPtr),
		AllTypes.Timez.EQ(Timez(23, 6, 6, 222, "+200")),
		AllTypes.Timestamp.EQ(AllTypes.TimestampPtr),
		AllTypes.Timestamp.EQ(Timestamp(2010, 10, 21, 15, 30, 12, 333)),
		AllTypes.Timestampz.EQ(AllTypes.TimestampzPtr),
		AllTypes.Timestampz.EQ(Timestampz(2010, 10, 21, 15, 30, 12, 444, "PST")),
		AllTypes.Date.EQ(AllTypes.DatePtr),
		AllTypes.Date.EQ(Date(2010, 12, 3)),

		AllTypes.Time.NOT_EQ(AllTypes.Time),
		AllTypes.Time.NOT_EQ(Time(23, 6, 6, 10)),
		AllTypes.Timez.NOT_EQ(AllTypes.TimezPtr),
		AllTypes.Timez.NOT_EQ(Timez(23, 6, 6, 555, "+4:00")),
		AllTypes.Timestamp.NOT_EQ(AllTypes.TimestampPtr),
		AllTypes.Timestamp.NOT_EQ(Timestamp(2010, 10, 21, 15, 30, 12, 666)),
		AllTypes.Timestampz.NOT_EQ(AllTypes.TimestampzPtr),
		AllTypes.Timestampz.NOT_EQ(Timestampz(2010, 10, 21, 15, 30, 12, 777, "UTC")),
		AllTypes.Date.NOT_EQ(AllTypes.DatePtr),
		AllTypes.Date.NOT_EQ(Date(2010, 12, 3)),

		AllTypes.Time.IS_DISTINCT_FROM(AllTypes.Time),
		AllTypes.Time.IS_DISTINCT_FROM(Time(23, 6, 6, 100)),

		AllTypes.Time.IS_NOT_DISTINCT_FROM(AllTypes.Time),
		AllTypes.Time.IS_NOT_DISTINCT_FROM(Time(23, 6, 6, 200)),

		AllTypes.Time.LT(AllTypes.Time),
		AllTypes.Time.LT(Time(23, 6, 6, 22)),

		AllTypes.Time.LT_EQ(AllTypes.Time),
		AllTypes.Time.LT_EQ(Time(23, 6, 6, 33)),

		AllTypes.Time.GT(AllTypes.Time),
		AllTypes.Time.GT(Time(23, 6, 6, 0)),

		AllTypes.Time.GT_EQ(AllTypes.Time),
		AllTypes.Time.GT_EQ(Time(23, 6, 6, 1)),

		AllTypes.Date.ADD(INTERVAL(1, HOUR)),
		AllTypes.Date.SUB(INTERVAL(1, MINUTE)),
		AllTypes.Time.ADD(INTERVAL(1, HOUR)),
		AllTypes.Time.SUB(INTERVAL(1, MINUTE)),
		AllTypes.Timez.ADD(INTERVAL(1, HOUR)),
		AllTypes.Timez.SUB(INTERVAL(1, MINUTE)),
		AllTypes.Timestamp.ADD(INTERVAL(1, HOUR)),
		AllTypes.Timestamp.SUB(INTERVAL(1, MINUTE)),
		AllTypes.Timestampz.ADD(INTERVAL(1, HOUR)),
		AllTypes.Timestampz.SUB(INTERVAL(1, MINUTE)),

		AllTypes.Date.SUB(CAST(String("04:05:06")).AS_INTERVAL()),

		CURRENT_DATE(),
		CURRENT_TIME(),
		CURRENT_TIME(2),
		CURRENT_TIMESTAMP(),
		CURRENT_TIMESTAMP(1),
		LOCALTIME(),
		LOCALTIME(11),
		LOCALTIMESTAMP(),
		LOCALTIMESTAMP(4),
		NOW(),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestInterval(t *testing.T) {
	query := SELECT(
		INTERVAL(1, YEAR),
		INTERVAL(1, MONTH),
		INTERVAL(1, WEEK),
		INTERVAL(1, DAY),
		INTERVAL(1, HOUR),
		INTERVAL(1, MINUTE),
		INTERVAL(1, SECOND),
		INTERVAL(1, MILLISECOND),
		INTERVAL(1, MICROSECOND),
		INTERVAL(1, DECADE),
		INTERVAL(1, CENTURY),
		INTERVAL(1, MILLENNIUM),

		INTERVAL(1, YEAR, 10, MONTH),
		INTERVAL(1, YEAR, 10, MONTH, 20, DAY),
		INTERVAL(1, YEAR, 10, MONTH, 20, DAY, 3, HOUR),

		INTERVAL(1, YEAR).IS_NOT_NULL(),
		INTERVAL(1, YEAR).AS("one year"),

		INTERVALd(0),
		INTERVALd(1*time.Microsecond),
		INTERVALd(1*time.Millisecond),
		INTERVALd(1*time.Second),
		INTERVALd(1*time.Minute),
		INTERVALd(1*time.Hour),
		INTERVALd(24*time.Hour),
		INTERVALd(24*time.Hour+2*time.Hour+3*time.Minute+4*time.Second+5*time.Microsecond),

		AllTypes.Interval.EQ(INTERVAL(2, HOUR, 20, MINUTE)).EQ(Bool(true)),
		AllTypes.IntervalPtr.NOT_EQ(INTERVAL(2, HOUR, 20, MINUTE)).EQ(Bool(false)),
		AllTypes.Interval.IS_DISTINCT_FROM(INTERVAL(2, HOUR, 20, MINUTE)).EQ(AllTypes.Boolean),
		AllTypes.IntervalPtr.IS_NOT_DISTINCT_FROM(INTERVALd(10*time.Microsecond)).EQ(AllTypes.Boolean),
		AllTypes.Interval.LT(AllTypes.IntervalPtr).EQ(AllTypes.BooleanPtr),
		AllTypes.Interval.LT_EQ(AllTypes.IntervalPtr).EQ(AllTypes.BooleanPtr),
		AllTypes.Interval.GT(AllTypes.IntervalPtr).EQ(AllTypes.BooleanPtr),
		AllTypes.Interval.GT_EQ(AllTypes.IntervalPtr).EQ(AllTypes.BooleanPtr),
		AllTypes.Interval.ADD(AllTypes.IntervalPtr).EQ(INTERVALd(17*time.Second)),
		AllTypes.Interval.SUB(AllTypes.IntervalPtr).EQ(INTERVAL(100, MICROSECOND)),
		AllTypes.IntervalPtr.MUL(Int(11)).EQ(AllTypes.Interval),
		AllTypes.IntervalPtr.DIV(Float(22.222)).EQ(AllTypes.IntervalPtr),
	).FROM(AllTypes)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestSubQueryColumnReference(t *testing.T) {
	subQueries := []SelectTable{
		AllTypes.SELECT(
			AllTypes.Boolean,
			AllTypes.Integer,
			AllTypes.Real,
			AllTypes.Text,
			AllTypes.Time,
			AllTypes.Timez,
			AllTypes.Timestamp,
			AllTypes.Timestampz,
			AllTypes.Date,
			AllTypes.Bytea.AS("aliasedColumn"),
		).LIMIT(2).AsTable("subQuery"),

		UNION_ALL(
			AllTypes.SELECT(
				AllTypes.Boolean,
				AllTypes.Integer,
				AllTypes.Real,
				AllTypes.Text,
				AllTypes.Time,
				AllTypes.Timez,
				AllTypes.Timestamp,
				AllTypes.Timestampz,
				AllTypes.Date,
				AllTypes.Bytea.AS("aliasedColumn"),
			).LIMIT(1),

			AllTypes.SELECT(
				AllTypes.Boolean,
				AllTypes.Integer,
				AllTypes.Real,
				AllTypes.Text,
				AllTypes.Time,
				AllTypes.Timez,
				AllTypes.Timestamp,
				AllTypes.Timestampz,
				AllTypes.Date,
				AllTypes.Bytea.AS("aliasedColumn"),
			).LIMIT(1).OFFSET(1),
		).AsTable("subQuery"),
	}

	for _, subQuery := range subQueries {
		boolColumn := AllTypes.Boolean.From(subQuery)
		intColumn := AllTypes.Integer.From(subQuery)
		floatColumn := AllTypes.Real.From(subQuery)
		stringColumn := AllTypes.Text.From(subQuery)
		timeColumn := AllTypes.Time.From(subQuery)
		timezColumn := AllTypes.Timez.From(subQuery)
		timestampColumn := AllTypes.Timestamp.From(subQuery)
		timestampzColumn := AllTypes.Timestampz.From(subQuery)
		dateColumn := AllTypes.Date.From(subQuery)
		aliasedColumn := StringColumn("aliasedColumn").From(subQuery)

		query := SELECT(
			boolColumn,
			intColumn,
			floatColumn,
			stringColumn,
			timeColumn,
			timezColumn,
			timestampColumn,
			timestampzColumn,
			dateColumn,
			aliasedColumn,
		).FROM(subQuery)

		require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

		dest1 := []model.AllTypes{}
		err := query.Query(db, &dest1)
		require.NoError(t, err)
		require.EqualValues(t, testparrot.RecordNext(t, dest1), dest1)

		query = SELECT(
			subQuery.AllColumns(),
		).FROM(subQuery)

		require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

		dest2 := []model.AllTypes{}
		err = query.Query(db, &dest2)

		require.NoError(t, err)
		require.EqualValues(t, dest1, dest2)
	}
}

type timeLiteralsResult struct {
	Date      time.Time
	Time      time.Time
	Timez     time.Time
	Timestamp time.Time
}

func TestTimeLiterals(t *testing.T) {

	loc, err := time.LoadLocation("Europe/Berlin")
	require.NoError(t, err)

	timeT := time.Date(2009, 11, 17, 20, 34, 58, 651387237, loc)

	query := SELECT(
		DateT(timeT).AS("date"),
		TimeT(timeT).AS("time"),
		TimezT(timeT).AS("timez"),
		TimestampT(timeT).AS("timestamp"),
		TimestampzT(timeT).AS("timestampz"),
	).FROM(AllTypes).LIMIT(1)

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &timeLiteralsResult{})
}

type jsonbResult struct {
	JSONB            qrm.JSON
	JSONBArray       qrm.JSON
	Extract          qrm.JSON
	ExtractText      string
	ExtractPath      qrm.JSON
	ExtractPathText  string
	HasKey           bool
	HasAnyKey        bool
	HasKeys          bool
	Contains         bool
	Concat           qrm.JSON
	DeleteKey        qrm.JSON
	DeleteIndex      qrm.JSON
	Delete           qrm.JSON
	ArrayLen         int
	TypeOf           string
	ToJSONB          qrm.JSON
	JSONBBuildArray  qrm.JSON
	JSONBBuildObject qrm.JSON
}

func TestJSONBOperators(t *testing.T) {
	mapVal := map[string]interface{}{"key": "value", "key2": []int{0, 1, 2}}
	arrayVal := []string{"a", "b"}
	query := SELECT(
		JSONB(mapVal).AS("jsonb"),
		JSONB(arrayVal).AS("jsonb_array"),
		JSONB(mapVal).EXTRACT(String("key")).AS("extract"),
		JSONB(mapVal).EXTRACT_TEXT(String("key")).AS("extract_text"),
		JSONB(mapVal).EXTRACT_PATH(StringArray("key")).AS("extract_path"),
		JSONB(mapVal).EXTRACT_PATH_TEXT(StringArray("key2", "0")).AS("extract_path_text"),
		JSONB(mapVal).CONTAINS(JSONB(mapVal)).AS("contains"),
		JSONB(mapVal).IS_CONTAINED_BY(JSONB(mapVal)).AS("contains"),
		JSONB(mapVal).HAS_KEY(String("key")).AS("has_key"),
		JSONB(mapVal).HAS_ANY_KEY(StringArray("key", "key3")).AS("has_any_key"),
		JSONB(mapVal).HAS_KEYS(StringArray("key", "key3")).AS("has_keys"),
		JSONB(mapVal).CONCAT(JSONB(map[string]string{"key3": "value3"})).AS("concat"),
		JSONB(mapVal).DELETE_KEY(String("key2")).AS("delete_key"),
		JSONB(arrayVal).DELETE_INDEX(Int(1)).AS("delete_index"),
		JSONB(mapVal).DELETE(StringArray("key2", "0")).AS("delete"),
		JSONB(arrayVal).ARRAY_LEN().AS("array_len"),
		JSONB(arrayVal).TYPEOF().AS("typeof"),
		TO_JSONB(StringArray("elem1", "elem2")).AS("to_jsonb"),
		JSONB_BUILD_ARRAY(
			String("2"),
			Int(2),
			Float(2.2),
			JSONB(map[string]string{"key": "value"}),
		).AS("jsonb_build_array"),
		JSONB_BUILD_OBJECT(
			String("text"), String("2"),
			String("int"), Int(2),
			String("real"), Float(2.2),
			String("jsonb"), JSONB(map[string]string{"key": "value"}),
			String("subobject"), JSONB_BUILD_OBJECT(
				String("key1"), String("value1"),
				String("key2"), Int(2),
			),
		).AS("jsonb_build_object"),
	).
		FROM(AllTypes).
		LIMIT(1)

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &jsonbResult{})
}

var allTypesRow0 = model.AllTypes{
	SmallIntPtr:        testutils.Ptr(int16(14)).(*int16),
	SmallInt:           14,
	IntegerPtr:         testutils.Ptr(int32(300)).(*int32),
	Integer:            300,
	BigIntPtr:          testutils.Ptr(int64(50000)).(*int64),
	BigInt:             5000,
	DecimalPtr:         testutils.Ptr(float64(1.11)).(*float64),
	Decimal:            1.11,
	NumericPtr:         testutils.Ptr(float64(2.22)).(*float64),
	Numeric:            2.22,
	RealPtr:            testutils.Ptr(float32(5.55)).(*float32),
	Real:               5.55,
	DoublePrecisionPtr: testutils.Ptr(float64(11111111.22)).(*float64),
	DoublePrecision:    11111111.22,
	Smallserial:        1,
	Serial:             1,
	Bigserial:          1,
	//MoneyPtr: nil,
	//Money:
	VarCharPtr:           testutils.Ptr("ABBA").(*string),
	VarChar:              "ABBA",
	CharPtr:              testutils.Ptr("JOHN                                                                            ").(*string),
	Char:                 "JOHN                                                                            ",
	TextPtr:              testutils.Ptr("Some text").(*string),
	Text:                 "Some text",
	ByteaPtr:             testutils.Ptr([]byte("bytea")).(*[]byte),
	Bytea:                []byte("bytea"),
	TimestampzPtr:        testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0),
	Timestampz:           *testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0),
	TimestampPtr:         testutils.TimestampWithoutTimeZone("1999-01-08 04:05:06", 0),
	Timestamp:            *testutils.TimestampWithoutTimeZone("1999-01-08 04:05:06", 0),
	DatePtr:              testutils.TimestampWithoutTimeZone("1999-01-08 00:00:00", 0),
	Date:                 *testutils.TimestampWithoutTimeZone("1999-01-08 00:00:00", 0),
	TimezPtr:             testutils.TimeWithTimeZone("04:05:06 -0800"),
	Timez:                *testutils.TimeWithTimeZone("04:05:06 -0800"),
	TimePtr:              testutils.TimeWithoutTimeZone("04:05:06"),
	Time:                 *testutils.TimeWithoutTimeZone("04:05:06"),
	IntervalPtr:          testutils.Ptr("3 days 04:05:06").(*string),
	Interval:             "3 days 04:05:06",
	BooleanPtr:           testutils.Ptr(true).(*bool),
	Boolean:              false,
	PointPtr:             testutils.Ptr("(2,3)").(*string),
	BitPtr:               testutils.Ptr("101").(*string),
	Bit:                  "101",
	BitVaryingPtr:        testutils.Ptr("101111").(*string),
	BitVarying:           "101111",
	TsvectorPtr:          testutils.Ptr("'supernova':1").(*string),
	Tsvector:             "'supernova':1",
	UUIDPtr:              testutils.Ptr(uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")).(*uuid.UUID),
	UUID:                 uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
	XMLPtr:               testutils.Ptr("<Sub>abc</Sub>").(*string),
	XML:                  "<Sub>abc</Sub>",
	JSONPtr:              testutils.Ptr(qrm.JSON(`{"a": 1, "b": 3}`)).(*qrm.JSON),
	JSON:                 qrm.JSON(`{"a": 1, "b": 3}`),
	JsonbPtr:             testutils.Ptr(qrm.JSON(`{"a": 1, "b": 3}`)).(*qrm.JSON),
	Jsonb:                qrm.JSON(`{"a": 1, "b": 3}`),
	IntegerArrayPtr:      testutils.Ptr("{1,2,3}").(*string),
	IntegerArray:         "{1,2,3}",
	TextArrayPtr:         testutils.Ptr("{breakfast,consulting}").(*string),
	TextArray:            "{breakfast,consulting}",
	JsonbArray:           `{"{\"a\": 1, \"b\": 2}","{\"a\": 3, \"b\": 4}"}`,
	TextMultiDimArrayPtr: testutils.Ptr("{{meeting,lunch},{training,presentation}}").(*string),
	TextMultiDimArray:    "{{meeting,lunch},{training,presentation}}",
}

var allTypesRow1 = model.AllTypes{
	SmallIntPtr:        nil,
	SmallInt:           14,
	IntegerPtr:         nil,
	Integer:            300,
	BigIntPtr:          nil,
	BigInt:             5000,
	DecimalPtr:         nil,
	Decimal:            1.11,
	NumericPtr:         nil,
	Numeric:            2.22,
	RealPtr:            nil,
	Real:               5.55,
	DoublePrecisionPtr: nil,
	DoublePrecision:    11111111.22,
	Smallserial:        2,
	Serial:             2,
	Bigserial:          2,
	//MoneyPtr: nil,
	//Money:
	VarCharPtr:           nil,
	VarChar:              "ABBA",
	CharPtr:              nil,
	Char:                 "JOHN                                                                            ",
	TextPtr:              nil,
	Text:                 "Some text",
	ByteaPtr:             nil,
	Bytea:                []byte("bytea"),
	TimestampzPtr:        nil,
	Timestampz:           *testutils.TimestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0),
	TimestampPtr:         nil,
	Timestamp:            *testutils.TimestampWithoutTimeZone("1999-01-08 04:05:06", 0),
	DatePtr:              nil,
	Date:                 *testutils.TimestampWithoutTimeZone("1999-01-08 00:00:00", 0),
	TimezPtr:             nil,
	Timez:                *testutils.TimeWithTimeZone("04:05:06 -0800"),
	TimePtr:              nil,
	Time:                 *testutils.TimeWithoutTimeZone("04:05:06"),
	IntervalPtr:          nil,
	Interval:             "3 days 04:05:06",
	BooleanPtr:           nil,
	Boolean:              false,
	PointPtr:             nil,
	BitPtr:               nil,
	Bit:                  "101",
	BitVaryingPtr:        nil,
	BitVarying:           "101111",
	TsvectorPtr:          nil,
	Tsvector:             "'supernova':1",
	UUIDPtr:              nil,
	UUID:                 uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
	XMLPtr:               nil,
	XML:                  "<Sub>abc</Sub>",
	JSONPtr:              nil,
	JSON:                 qrm.JSON(`{"a": 1, "b": 3}`),
	JsonbPtr:             nil,
	Jsonb:                qrm.JSON(`{"a": 1, "b": 3}`),
	IntegerArrayPtr:      nil,
	IntegerArray:         "{1,2,3}",
	TextArrayPtr:         nil,
	TextArray:            "{breakfast,consulting}",
	JsonbArray:           `{"{\"a\": 1, \"b\": 2}","{\"a\": 3, \"b\": 4}"}`,
	TextMultiDimArrayPtr: nil,
	TextMultiDimArray:    "{{meeting,lunch},{training,presentation}}",
}

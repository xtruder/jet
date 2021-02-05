package jet

import (
	"testing"
)

func TestFuncAVG(t *testing.T) {
	assertSerialize(t, AVG(table1ColFloat), "AVG(table1.col_float)")
	assertSerialize(t, AVG(table1ColInt), "AVG(table1.col_int)")
}

func TestFuncBIT_AND(t *testing.T) {
	assertSerialize(t, BIT_AND(table1ColInt), "BIT_AND(table1.col_int)")
}

func TestFuncBIT_OR(t *testing.T) {
	assertSerialize(t, BIT_OR(table1ColInt), "BIT_OR(table1.col_int)")
}

func TestFuncBOOL_AND(t *testing.T) {
	assertSerialize(t, BOOL_AND(table1ColBool), "BOOL_AND(table1.col_bool)")
}

func TestFuncBOOL_OR(t *testing.T) {
	assertSerialize(t, BOOL_OR(table1ColBool), "BOOL_OR(table1.col_bool)")
}

func TestFuncEVERY(t *testing.T) {
	assertSerialize(t, EVERY(table1ColBool), "EVERY(table1.col_bool)")
}

func TestFuncMIN(t *testing.T) {
	t.Run("expression", func(t *testing.T) {
		assertSerialize(t, MIN(table1ColDate), "MIN(table1.col_date)")
		assertSerialize(t, MIN(Date(2001, 1, 1)), "MIN($1)", "2001-01-01")
		assertSerialize(t, MIN(Time(12, 10, 10)), "MIN($1)", "12:10:10")
		assertSerialize(t, MIN(Timestamp(2001, 1, 1, 12, 10, 10)), "MIN($1)", "2001-01-01 12:10:10")
		assertSerialize(t, MIN(Timestampz(2001, 1, 1, 12, 10, 10, 1, "UTC")), "MIN($1)", "2001-01-01 12:10:10.000000001 UTC")
	})

	t.Run("float", func(t *testing.T) {
		assertSerialize(t, MINf(table1ColFloat), "MIN(table1.col_float)")
	})

	t.Run("integer", func(t *testing.T) {
		assertSerialize(t, MINi(table1ColInt), "MIN(table1.col_int)")
	})
}

func TestFuncMAX(t *testing.T) {
	t.Run("expression", func(t *testing.T) {
		assertSerialize(t, MAX(table1ColDate), "MAX(table1.col_date)")
		assertSerialize(t, MAX(Date(2001, 1, 1)), "MAX($1)", "2001-01-01")
		assertSerialize(t, MAX(Time(12, 10, 10)), "MAX($1)", "12:10:10")
		assertSerialize(t, MAX(Timestamp(2001, 1, 1, 12, 10, 10)), "MAX($1)", "2001-01-01 12:10:10")
		assertSerialize(t, MAX(Timestampz(2001, 1, 1, 12, 10, 10, 1, "UTC")), "MAX($1)", "2001-01-01 12:10:10.000000001 UTC")
	})

	t.Run("float", func(t *testing.T) {
		assertSerialize(t, MAXf(table1ColFloat), "MAX(table1.col_float)")
		assertSerialize(t, MAXf(Float(11.2222)), "MAX($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertSerialize(t, MAXi(table1ColInt), "MAX(table1.col_int)")
		assertSerialize(t, MAXi(Int(11)), "MAX($1)", int64(11))
	})
}

func TestFuncSUM(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertSerialize(t, SUMf(table1ColFloat), "SUM(table1.col_float)")
		assertSerialize(t, SUMf(Float(11.2222)), "SUM($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertSerialize(t, SUMi(table1ColInt), "SUM(table1.col_int)")
		assertSerialize(t, SUMi(Int(11)), "SUM($1)", int64(11))
	})
}

func TestFuncCOUNT(t *testing.T) {
	assertSerialize(t, COUNT(STAR), "COUNT(*)")
	assertSerialize(t, COUNT(table1ColFloat), "COUNT(table1.col_float)")
	assertSerialize(t, COUNT(Float(11.2222)), "COUNT($1)", float64(11.2222))
}

func TestFuncABS(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		assertSerialize(t, ABSf(table1ColFloat), "ABS(table1.col_float)")
		assertSerialize(t, ABSf(Float(11.2222)), "ABS($1)", float64(11.2222))
	})

	t.Run("integer", func(t *testing.T) {
		assertSerialize(t, ABSi(table1ColInt), "ABS(table1.col_int)")
		assertSerialize(t, ABSi(Int(11)), "ABS($1)", int64(11))
	})
}

func TestFuncSQRT(t *testing.T) {
	assertSerialize(t, SQRT(table1ColFloat), "SQRT(table1.col_float)")
	assertSerialize(t, SQRT(Float(11.2222)), "SQRT($1)", float64(11.2222))
	assertSerialize(t, SQRT(table1ColInt), "SQRT(table1.col_int)")
	assertSerialize(t, SQRT(Int(11)), "SQRT($1)", int64(11))
}

func TestFuncCBRT(t *testing.T) {
	assertSerialize(t, CBRT(table1ColFloat), "CBRT(table1.col_float)")
	assertSerialize(t, CBRT(Float(11.2222)), "CBRT($1)", float64(11.2222))
	assertSerialize(t, CBRT(table1ColInt), "CBRT(table1.col_int)")
	assertSerialize(t, CBRT(Int(11)), "CBRT($1)", int64(11))
}

func TestFuncCEIL(t *testing.T) {
	assertSerialize(t, CEIL(table1ColFloat), "CEIL(table1.col_float)")
	assertSerialize(t, CEIL(Float(11.2222)), "CEIL($1)", float64(11.2222))
}

func TestFuncFLOOR(t *testing.T) {
	assertSerialize(t, FLOOR(table1ColFloat), "FLOOR(table1.col_float)")
	assertSerialize(t, FLOOR(Float(11.2222)), "FLOOR($1)", float64(11.2222))
}

func TestFuncROUND(t *testing.T) {
	assertSerialize(t, ROUND(table1ColFloat), "ROUND(table1.col_float)")
	assertSerialize(t, ROUND(Float(11.2222)), "ROUND($1)", float64(11.2222))

	assertSerialize(t, ROUND(table1ColFloat, Int(2)), "ROUND(table1.col_float, $1)", int64(2))
	assertSerialize(t, ROUND(Float(11.2222), Int(1)), "ROUND($1, $2)", float64(11.2222), int64(1))
}

func TestFuncSIGN(t *testing.T) {
	assertSerialize(t, SIGN(table1ColFloat), "SIGN(table1.col_float)")
	assertSerialize(t, SIGN(Float(11.2222)), "SIGN($1)", float64(11.2222))
}

func TestFuncTRUNC(t *testing.T) {
	assertSerialize(t, TRUNC(table1ColFloat), "TRUNC(table1.col_float)")
	assertSerialize(t, TRUNC(Float(11.2222)), "TRUNC($1)", float64(11.2222))

	assertSerialize(t, TRUNC(table1ColFloat, Int(2)), "TRUNC(table1.col_float, $1)", int64(2))
	assertSerialize(t, TRUNC(Float(11.2222), Int(1)), "TRUNC($1, $2)", float64(11.2222), int64(1))
}

func TestFuncLN(t *testing.T) {
	assertSerialize(t, LN(table1ColFloat), "LN(table1.col_float)")
	assertSerialize(t, LN(Float(11.2222)), "LN($1)", float64(11.2222))
}

func TestFuncLOG(t *testing.T) {
	assertSerialize(t, LOG(table1ColFloat), "LOG(table1.col_float)")
	assertSerialize(t, LOG(Float(11.2222)), "LOG($1)", float64(11.2222))
}

func TestFuncCOALESCE(t *testing.T) {
	assertSerialize(t, COALESCE(table1ColFloat), "COALESCE(table1.col_float)")
	assertSerialize(t, COALESCE(Float(11.2222), NULL, String("str")), "COALESCE($1, NULL, $2)", float64(11.2222), "str")
}

func TestFuncNULLIF(t *testing.T) {
	assertSerialize(t, NULLIF(table1ColFloat, table2ColInt), "NULLIF(table1.col_float, table2.col_int)")
	assertSerialize(t, NULLIF(Float(11.2222), NULL), "NULLIF($1, NULL)", float64(11.2222))
}

func TestFuncGREATEST(t *testing.T) {
	assertSerialize(t, GREATEST(table1ColFloat), "GREATEST(table1.col_float)")
	assertSerialize(t, GREATEST(Float(11.2222), NULL, String("str")), "GREATEST($1, NULL, $2)", float64(11.2222), "str")
}

func TestFuncLEAST(t *testing.T) {
	assertSerialize(t, LEAST(table1ColFloat), "LEAST(table1.col_float)")
	assertSerialize(t, LEAST(Float(11.2222), NULL, String("str")), "LEAST($1, NULL, $2)", float64(11.2222), "str")
}

func TestTO_ASCII(t *testing.T) {
	assertSerialize(t, TO_ASCII(String("Karel")), `TO_ASCII($1)`, "Karel")
	assertSerialize(t, TO_ASCII(String("Karel")), `TO_ASCII($1)`, "Karel")
}

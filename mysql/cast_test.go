package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestCAST(t *testing.T) {
	testutils.SerializerTests{
		{Name: "as custom", Test: CAST(Float(11.22)).AS("bigint")},
		{Name: "as char", Test: CAST(Int(22)).AS_CHAR()},
		{Name: "as decimal", Test: CAST(Int(22)).AS_DECIMAL()},
		{Name: "as date", Test: CAST(Int(22)).AS_DATE()},
		{Name: "as time", Test: CAST(Int(22)).AS_TIME()},
		{Name: "as datetime", Test: CAST(Int(22)).AS_DATETIME()},
		{Name: "as signed", Test: CAST(Int(22)).AS_SIGNED()},
		{Name: "as unsigned", Test: CAST(Int(22)).AS_UNSIGNED()},
		{Name: "as binary", Test: CAST(Int(22)).AS_BINARY()},
	}.Run(t, Dialect)
}

package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestOperatorARRAY(t *testing.T) {
	testutils.SerializerTest{Test: ARRAY(table2ColStr, String("test"))}.Assert(t, Dialect)
}

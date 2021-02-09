package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/testutils"
)

var assertStatementRecordSQL = testutils.AssertStatementRecordSQL

func assertExec(t *testing.T, stmt jet.Statement, rowsAffected ...int64) {
	t.Helper()
	testutils.AssertExec(t, stmt, db, rowsAffected...)
}

func assertQuery(t *testing.T, stmt jet.Statement) {
	t.Helper()
	dest := []struct{}{}
	testutils.AssertQuery(t, db, stmt, &dest)
}

func assertQueryDest(t *testing.T, stmt jet.Statement, dest interface{}) {
	t.Helper()
	testutils.AssertQuery(t, db, stmt, dest)
}

func assertQueryRecordValues(t *testing.T, stmt jet.Statement, dest interface{}) {
	t.Helper()
	testutils.AssertQueryRecordValues(t, db, stmt, dest)
}

func assertQueryPanicErr(t *testing.T, stmt jet.Statement, dest interface{}, errString string) {
	t.Helper()
	testutils.AssertQueryPanicErr(t, stmt, db, dest, errString)
}

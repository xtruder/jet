package mariadb

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/go-jet/jet/v2/qrm"
)

var assertStatementRecordSQL = testutils.AssertStatementRecordSQL
var assertStatementSql = testutils.AssertStatementSql

func assertExec(t *testing.T, stmt jet.Statement, db qrm.DB, rowsAffected ...int64) {
	t.Helper()
	testutils.AssertExec(t, stmt, db, rowsAffected...)
}

func assertQueryRecordValues(t *testing.T, stmt jet.Statement, dest interface{}) {
	t.Helper()
	testutils.AssertQueryRecordValues(t, db, stmt, dest)
}

func assertQuery(t *testing.T, query jet.Statement) {
	t.Helper()
	dest := []struct{}{}
	testutils.AssertQuery(t, db, query, &dest)
}

func assertStatementSqlErr(t *testing.T, stmt jet.Statement, errorStr string) {
	t.Helper()
	testutils.AssertStatementSqlErr(t, stmt, errorStr)
}

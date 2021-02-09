package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestLockRead(t *testing.T) {
	testutils.StatementTest{Test: table2.LOCK().READ()}.Assert(t)
}

func TestLockWrite(t *testing.T) {
	testutils.StatementTest{Test: table2.LOCK().WRITE()}.Assert(t)
}

func TestUNLOCK_TABLES(t *testing.T) {
	testutils.StatementTest{Test: UNLOCK_TABLES()}.Assert(t)
}

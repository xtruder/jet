package mariadb

import (
	"testing"

	. "github.com/go-jet/jet/v2/mysql"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/dvds/table"
)

func TestLockRead(t *testing.T) {
	query := Customer.LOCK().READ()

	defer db.Exec("UNLOCK TABLES;")

	assertStatementSql(t, query, `LOCK TABLES dvds.customer READ;`)
	assertExec(t, query, db)
}

func TestLockWrite(t *testing.T) {
	query := Customer.LOCK().WRITE()

	defer db.Exec("UNLOCK TABLES;")

	assertStatementSql(t, query, `LOCK TABLES dvds.customer WRITE;`)
	assertExec(t, query, db)
}

func TestUnlockTables(t *testing.T) {
	query := UNLOCK_TABLES()

	assertStatementSql(t, query, `UNLOCK TABLES;`)
	assertExec(t, query, db)
}

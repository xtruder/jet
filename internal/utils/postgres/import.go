package postgresutils

import (
	"database/sql"
)

func ImportSQL(db *sql.DB, sql string) (err error) {
	_, err = db.Exec(sql)
	return
}

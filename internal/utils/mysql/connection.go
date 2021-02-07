package mysqlutils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ConnOptions struct {
	ConnStr  string
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Params   string
}

func Connect(opts ConnOptions) (*sql.DB, error) {
	connStr := opts.ConnStr

	if connStr == "" {
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", opts.User, opts.Password, opts.Host, opts.Port, opts.DBName)
		if opts.Params != "" {
			connStr += "?" + opts.Params
		}
	}

	fmt.Println("Connecting to MySQL database: " + connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err

}

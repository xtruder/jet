package postgresutils

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
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

func Connect(dbOpts ConnOptions) (*sql.DB, error) {
	connStr := dbOpts.ConnStr

	if connStr == "" {
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
			dbOpts.Host, strconv.Itoa(dbOpts.Port), dbOpts.User, dbOpts.Password, dbOpts.DBName, dbOpts.Params)
	}

	fmt.Println("Connecting to postgres database: " + connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

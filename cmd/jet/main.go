package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	mysqlgen "github.com/go-jet/jet/v2/generator/mysql"
	postgresgen "github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/internal/utils"
	mysqlutils "github.com/go-jet/jet/v2/internal/utils/mysql"
	postgresutils "github.com/go-jet/jet/v2/internal/utils/postgres"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/postgres"
)

var (
	source  string
	destDir string

	connString string
	host       string
	port       int
	user       string
	password   string
	sslmode    string
	params     string
	dbName     string
	schemaName string

	logger = log.New(os.Stderr, "", 0)
)

func init() {
	flag.StringVar(&source, "source", "", "Database system name (PostgreSQL, MySQL or MariaDB)")
	flag.StringVar(&destDir, "path", "", "Destination dir for generated files")

	flag.StringVar(&connString, "connstr", "", "Database connection string")
	flag.StringVar(&host, "host", "", "Database host path (Example: localhost)")
	flag.IntVar(&port, "port", 0, "Database port")
	flag.StringVar(&user, "user", "", "Database user")
	flag.StringVar(&password, "password", "", "The user’s password")
	flag.StringVar(&params, "params", "", "Additional connection string parameters(optional)")
	flag.StringVar(&dbName, "dbname", "", "Database name")
	flag.StringVar(&schemaName, "schema", "public", `Database schema name. (default "public") (ignored for MySQL and MariaDB)`)
	flag.StringVar(&sslmode, "sslmode", "disable", `Whether or not to use SSL(optional)(default "disable") (ignored for MySQL and MariaDB)`)

	// parse flags from env variables
	utils.ParseEnv("jet")
}

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprint(os.Stdout, `
Jet generator 2.3.0

Usage:
  -source string
    	Database system name (PostgreSQL, MySQL or MariaDB)
  -host string
        Database host path (Example: localhost)
  -port int
        Database port
  -user string
        Database user
  -password string
        The user’s password
  -dbname string
        Database name
  -params string
        Additional connection string parameters(optional)
  -schema string
        Database schema name. (default "public") (ignored for MySQL and MariaDB)
  -sslmode string
        Whether or not to use SSL(optional) (default "disable") (ignored for MySQL and MariaDB)
  -path string
        Destination dir for files generated.
`)
	}

	flag.Parse()

	if source == "" || host == "" || port == 0 || user == "" || dbName == "" {
		flag.Usage()
		logger.Fatalf("\nERROR: required flag(s) missing")
	}

	var err error
	var db *sql.DB

	switch strings.ToLower(strings.TrimSpace(source)) {
	case strings.ToLower(postgres.Dialect.Name()),
		strings.ToLower(postgres.Dialect.PackageName()):
		if sslmode != "" {
			params += fmt.Sprintf(" sslmode=%s ", sslmode)
		}

		db, err = postgresutils.Connect(postgresutils.ConnOptions{
			User:     user,
			Password: password,
			Host:     host,
			Port:     port,
			DBName:   dbName,
			Params:   params,
		})
		if err != nil {
			break
		}

		err = postgresgen.Generate(db, schemaName, destDir)

	case strings.ToLower(mysql.Dialect.Name()), "mariadb":
		db, err = mysqlutils.Connect(mysqlutils.ConnOptions{
			User:     user,
			Password: password,
			Host:     host,
			Port:     port,
			DBName:   dbName,
			Params:   params,
		})
		if err != nil {
			break
		}

		err = mysqlgen.Generate(db, dbName, destDir)

	default:
		logger.Fatalf("ERROR: unsupported source %s. %s and %s are currently supported.\n",
			postgres.Dialect.Name(), source, mysql.Dialect.Name())
	}

	if err != nil {
		logger.Fatalln(err.Error())
	}
}

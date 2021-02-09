package main

import (
	"bytes"
	"flag"
	"fmt"
	"path"

	testdata "github.com/go-jet/jet-test-data"
	mysqlgen "github.com/go-jet/jet/v2/generator/mysql"
	"github.com/go-jet/jet/v2/internal/utils"
	mysqlutils "github.com/go-jet/jet/v2/internal/utils/mysql"
)

var (
	generate  bool
	importDbs bool

	host     string
	port     int
	user     string
	password string
	genPath  string
)

func init() {
	flag.StringVar(&host, "mysql-host", "localhost", "mysql host")
	flag.IntVar(&port, "mysql-port", 3306, "mysql port")
	flag.StringVar(&user, "mysql-user", "jet", "mysql user")
	flag.StringVar(&password, "mysql-password", "jet", "mysql password")
	flag.BoolVar(&importDbs, "import", true, "whether to import databases")
	flag.BoolVar(&generate, "generate", true, "whether to generate models")
	flag.StringVar(&genPath, "gen-path", path.Join(utils.PkgPath(0), "../gen"), "path where to generate files")

	// parse including environment variables
	utils.ParseEnv("")
}

var initDBs = []string{
	"dvds",
	"test_sample",
}

func main() {
	mysqlConnOpts := mysqlutils.ConnOptions{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}

	for _, dbName := range initDBs {
		fmt.Printf("Importing mysql database: %s\n", dbName)

		if importDbs {
			data, err := testdata.Asset("mysql/" + dbName + ".sql")
			utils.PanicOnError(err)

			mysqlConnOpts.DBName = dbName

			utils.PanicOnError(mysqlutils.ImportSQL(mysqlConnOpts, bytes.NewReader(data)))
		}

		if generate {
			fmt.Printf("Generating jet models: %s\n", dbName)

			db, err := mysqlutils.Connect(mysqlConnOpts)
			utils.PanicOnError(err)

			// get current package path
			genPath := path.Join(genPath, dbName)

			utils.PanicOnError(mysqlgen.Generate(db, dbName, genPath))
		}
	}
}

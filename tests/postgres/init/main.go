package main

import (
	"flag"
	"fmt"
	"path"

	testdata "github.com/go-jet/jet-test-data"
	postgresgen "github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/internal/utils"
	postgresutils "github.com/go-jet/jet/v2/internal/utils/postgres"
)

var (
	generate   bool
	importDbs  bool
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string
	genPath    string
)

func init() {
	flag.StringVar(&dbHost, "postgres-host", "localhost", "postgres host")
	flag.IntVar(&dbPort, "postgres-port", 5432, "postgres port")
	flag.StringVar(&dbUser, "postgres-user", "jet", "postgres user")
	flag.StringVar(&dbPassword, "postgres-password", "jet", "postgres password")
	flag.StringVar(&dbName, "postgres-db", "jet", "postgres database")
	flag.BoolVar(&importDbs, "import", true, "whether to import databases")
	flag.BoolVar(&generate, "generate", true, "whether to generate models")
	flag.StringVar(&genPath, "gen-path", path.Join(utils.PkgPath(0), "../gen"), "path where to generate files")

	// parses flags from env
	utils.ParseEnv("")
}

var initDBs = []string{
	"dvds",
	"test_sample",
	"chinook",
	"northwind",
}

func main() {
	postgresConnOpts := postgresutils.ConnOptions{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}

	for _, schemaName := range initDBs {
		postgresConnOpts.Params = fmt.Sprintf("search_path=%s sslmode=disable", schemaName)

		db, err := postgresutils.Connect(postgresConnOpts)
		utils.PanicOnError(err)

		if importDbs {
			fmt.Printf("Importing postgres database: %s\n", schemaName)

			data, err := testdata.Asset("postgres/" + schemaName + ".sql")
			utils.PanicOnError(err)

			fmt.Printf("Importing SQL: %s\n", schemaName)

			utils.PanicOnError(postgresutils.ImportSQL(db, string(data)))
		}

		if generate {
			fmt.Printf("Generating jet models: %s\n", schemaName)

			// get current package path
			genPath := path.Join(genPath, schemaName)

			utils.PanicOnError(postgresgen.Generate(db, schemaName, genPath))
		}
	}
}

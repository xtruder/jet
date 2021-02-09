//go:generate go run ./init --import=false
//go:generate go test ./ -testparrot.record -testparrot.splitfiles
package postgres

import (
	"context"
	"database/sql"
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/utils"
	postgresutils "github.com/go-jet/jet/v2/internal/utils/postgres"
	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"github.com/pkg/profile"
	"github.com/xtruder/go-testparrot"
)

var (
	db *sql.DB

	loggedSQL      string
	loggedSQLArgs  []interface{}
	loggedDebugSQL string

	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string
)

func init() {
	postgres.SetLogger(func(ctx context.Context, statement postgres.PrintableStatement) {
		loggedSQL, loggedSQLArgs = statement.Sql()
		loggedDebugSQL = statement.String()
	})
}

func TestMain(m *testing.M) {
	flag.StringVar(&dbHost, "postgres-host", "localhost", "postgres host")
	flag.IntVar(&dbPort, "postgres-port", 5432, "postgres port")
	flag.StringVar(&dbUser, "postgres-user", "jet", "postgres user")
	flag.StringVar(&dbPassword, "postgres-password", "jet", "postgres password")
	flag.StringVar(&dbName, "postgres-db", "jet", "postgres database")

	testparrot.BeforeTests(testparrot.R)

	utils.ParseEnv("")

	rand.Seed(time.Now().Unix())
	defer profile.Start().Stop()

	postgresConnOpts := postgresutils.ConnOptions{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
		Params:   "sslmode=disable",
	}

	var err error
	db, err = postgresutils.Connect(postgresConnOpts)
	utils.PanicOnError(err)
	defer db.Close()

	code := m.Run()

	if code != 0 {
		os.Exit(code)
		return
	}

	testparrot.AfterTests(testparrot.R, "")
}

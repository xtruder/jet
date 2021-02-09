//go:generate go run ./init --import=false
//go:generate go test -v ./ -testparrot.record -testparrot.splitfiles
package mariadb

import (
	"context"
	"database/sql"
	"flag"
	"math/rand"
	"time"

	"github.com/go-jet/jet/v2/internal/utils"
	mysqlutils "github.com/go-jet/jet/v2/internal/utils/mysql"
	jetmysql "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"

	_ "github.com/go-sql-driver/mysql"

	"os"
	"testing"

	"github.com/pkg/profile"
)

var (
	db     *sql.DB
	source string

	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string

	loggedSQL      string
	loggedSQLArgs  []interface{}
	loggedDebugSQL string
)

func init() {
	jetmysql.SetLogger(func(ctx context.Context, statement jetmysql.PrintableStatement) {
		loggedSQL, loggedSQLArgs = statement.Sql()
		loggedDebugSQL = statement.String()
		//fmt.Println(statement.String())
	})
}

func TestMain(m *testing.M) {
	flag.StringVar(&dbHost, "mariadb-host", "localhost", "mysql host")
	flag.IntVar(&dbPort, "mariadb-port", 3306, "mysql port")
	flag.StringVar(&dbUser, "mariadb-user", "jet", "mysql user")
	flag.StringVar(&dbPassword, "mariadb-password", "jet", "mysql password")
	flag.StringVar(&dbName, "mariadb-db", "dvds", "mysql database")

	testparrot.BeforeTests(testparrot.R)

	utils.ParseEnv("")

	rand.Seed(time.Now().Unix())
	defer profile.Start().Stop()

	mysqlConnOpts := mysqlutils.ConnOptions{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}

	var err error
	db, err = mysqlutils.Connect(mysqlConnOpts)
	if err != nil {
		panic("Failed to connect to test db" + err.Error())
	}
	defer db.Close()

	code := m.Run()

	if code != 0 {
		os.Exit(code)
		return
	}

	testparrot.AfterTests(testparrot.R, "")
}

func requireLogged(t *testing.T, statement postgres.Statement) {
	query, args := statement.Sql()
	require.Equal(t, loggedSQL, query)
	require.Equal(t, loggedSQLArgs, args)
	require.Equal(t, loggedDebugSQL, statement.String())
}

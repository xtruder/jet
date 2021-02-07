package mysql

import (
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/generator/internal/metadata"
	"github.com/go-jet/jet/v2/generator/internal/template"
	"github.com/go-jet/jet/v2/internal/utils"
	"github.com/go-jet/jet/v2/mysql"
)

func Generate(db *sql.DB, dbName string, destDir string) (err error) {
	defer utils.ErrorCatch(&err)

	fmt.Println("Retrieving database information...")
	// No schemas in MySQL
	dbInfo := metadata.GetSchemaMetaData(db, dbName, &mySqlQuerySet{})

	template.GenerateFiles(destDir, dbInfo, mysql.Dialect)

	return
}

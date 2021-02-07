package postgres

import (
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/generator/internal/metadata"
	"github.com/go-jet/jet/v2/generator/internal/template"
	"github.com/go-jet/jet/v2/internal/utils"
	"github.com/go-jet/jet/v2/postgres"
)

func Generate(db *sql.DB, schemaName string, destDir string) (err error) {
	defer utils.ErrorCatch(&err)

	fmt.Println("Retrieving schema information...")
	schemaInfo := metadata.GetSchemaMetaData(db, schemaName, &postgresQuerySet{})

	template.GenerateFiles(destDir, schemaInfo, postgres.Dialect)

	return
}

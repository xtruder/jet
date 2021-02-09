package mariadb

import (
	"testing"

	"github.com/go-jet/jet/v2/generator/mysql"
	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	genDir := t.TempDir()

	require.NoError(t,
		mysql.Generate(db, "dvds", genDir))

	testutils.AssertDirRecordContent(t, genDir)
}

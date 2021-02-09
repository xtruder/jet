package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"

	. "github.com/go-jet/jet/v2/postgres"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/dvds/table"
)

func TestLockTable(t *testing.T) {
	var testData = []TableLockMode{
		LOCK_ACCESS_SHARE,
		LOCK_ROW_SHARE,
		LOCK_ROW_EXCLUSIVE,
		LOCK_SHARE_UPDATE_EXCLUSIVE,
		LOCK_SHARE,
		LOCK_SHARE_ROW_EXCLUSIVE,
		LOCK_EXCLUSIVE,
		LOCK_ACCESS_EXCLUSIVE,
	}

	for _, lockMode := range testData {
		stmt := Address.LOCK().IN(lockMode)

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		tx, _ := db.Begin()

		_, err := stmt.Exec(tx)

		require.NoError(t, err)

		err = tx.Rollback()

		require.NoError(t, err)
	}

	for _, lockMode := range testData {
		stmt := Address.LOCK().IN(lockMode).NOWAIT()

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		tx, _ := db.Begin()

		_, err := stmt.Exec(tx)

		require.NoError(t, err)

		err = tx.Rollback()

		require.NoError(t, err)
	}
}

func TestLockExecContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	tx, _ := db.Begin()
	defer tx.Rollback()

	_, err := Address.LOCK().IN(LOCK_ACCESS_SHARE).ExecContext(ctx, tx)

	require.Error(t, err, "context deadline exceeded")
}

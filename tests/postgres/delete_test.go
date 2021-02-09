package postgres

import (
	"context"
	"testing"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/table"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestDeleteWithWhere(t *testing.T) {
	initForDeleteTest(t)

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(String("Gmail"), String("Outlook")))

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	assertExec(t, stmt, 2)
}

func TestDeleteWithWhereAndReturning(t *testing.T) {
	initForDeleteTest(t)

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(String("Gmail"), String("Outlook"))).
		RETURNING(Link.AllColumns)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	dest := []model.Link{}
	require.NoError(t, stmt.Query(db, &dest))

	// ID can change, so zero it out
	for i := range dest {
		dest[i].ID = 0
	}

	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func initForDeleteTest(t *testing.T) {
	cleanUpLinkTable(t)

	stmt := Link.INSERT(Link.URL, Link.Name, Link.Description).
		VALUES("www.gmail.com", "Gmail", "Email service developed by Google").
		VALUES("www.outlook.live.com", "Outlook", "Email service developed by Microsoft")

	assertExec(t, stmt, 2)
}

func TestDeleteQueryContext(t *testing.T) {
	initForDeleteTest(t)

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(String("Gmail"), String("Outlook")))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	dest := []model.Link{}
	err := stmt.QueryContext(ctx, db, &dest)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	require.Error(t, err, "context deadline exceeded")
}

func TestDeleteExecContext(t *testing.T) {
	initForDeleteTest(t)

	list := []Expression{String("Gmail"), String("Outlook")}

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(list...))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	_, err := stmt.ExecContext(ctx, db)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	require.Error(t, err, "context deadline exceeded")
}

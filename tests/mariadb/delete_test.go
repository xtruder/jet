package mariadb

import (
	"context"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/utils"
	. "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/table"
	"github.com/stretchr/testify/assert"
)

func TestDeleteWithWhere(t *testing.T) {
	initForDeleteTest(t)

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(String("Gmail"), String("Outlook")))

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestDeleteWithWhereOrderByLimit(t *testing.T) {
	initForDeleteTest(t)

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(String("Gmail"), String("Outlook"))).
		ORDER_BY(Link.Name).
		LIMIT(1)

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
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

	assertStatementRecordSQL(t, stmt)
	assert.Error(t, err, "context deadline exceeded")
}

func TestDeleteExecContext(t *testing.T) {
	initForDeleteTest(t)

	stmt := Link.
		DELETE().
		WHERE(Link.Name.IN(String("Gmail"), String("Outlook")))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	_, err := stmt.ExecContext(ctx, db)

	assertStatementRecordSQL(t, stmt)
	assert.Error(t, err, "context deadline exceeded")
}

func initForDeleteTest(t *testing.T) {
	_, err := Link.DELETE().WHERE(Link.ID.GT(Int(1))).Exec(db)
	utils.PanicOnError(err)

	stmt := Link.INSERT(Link.URL, Link.Name, Link.Description).
		VALUES("www.gmail.com", "Gmail", "Email service developed by Google").
		VALUES("www.outlook.live.com", "Outlook", "Email service developed by Microsoft")

	_, err = stmt.Exec(db)
	utils.PanicOnError(err)
}

package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/utils"
	. "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/table"
	"github.com/stretchr/testify/require"
)

func TestInsertValues(t *testing.T) {
	initForInsertTest(t)

	stmt := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		VALUES(101, "http://www.google.com", "Google", DEFAULT).
		VALUES(102, "http://www.yahoo.com", "Yahoo", nil)

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)

	query := Link.SELECT(Link.AllColumns).
		WHERE(Link.ID.GT_EQ(Int(100))).
		ORDER_BY(Link.ID)

	dest := []model.Link{}

	assertQueryRecordValues(t, query, &dest)
}

func TestInsertEmptyColumnList(t *testing.T) {
	initForInsertTest(t)

	stmt := Link.INSERT().
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT)

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)

	query := Link.SELECT(Link.AllColumns).
		WHERE(Link.ID.GT_EQ(Int(100))).
		ORDER_BY(Link.ID)

	dest := []model.Link{}

	assertQueryRecordValues(t, query, &dest)
}

func TestInsertModelObject(t *testing.T) {
	initForInsertTest(t)

	linkData := model.Link{
		URL:  "http://www.duckduckgo.com",
		Name: "Duck Duck go",
	}

	stmt := Link.
		INSERT(Link.URL, Link.Name).
		MODEL(linkData)

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestInsertModelObjectEmptyColumnList(t *testing.T) {
	initForInsertTest(t)

	linkData := model.Link{
		ID:   1000,
		URL:  "http://www.duckduckgo.com",
		Name: "Duck Duck go",
	}

	stmt := Link.
		INSERT().
		MODEL(linkData)

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestInsertModelsObject(t *testing.T) {
	initForInsertTest(t)

	tutorial := model.Link{
		URL:  "http://www.postgresqltutorial.com",
		Name: "PostgreSQL Tutorial",
	}

	google := model.Link{
		URL:  "http://www.google.com",
		Name: "Google",
	}

	yahoo := model.Link{
		URL:  "http://www.yahoo.com",
		Name: "Yahoo",
	}

	stmt := Link.
		INSERT(Link.URL, Link.Name).
		MODELS([]model.Link{tutorial, google, yahoo})

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestInsertUsingMutableColumns(t *testing.T) {
	initForInsertTest(t)

	google := model.Link{
		URL:  "http://www.google.com",
		Name: "Google",
	}

	yahoo := model.Link{
		URL:  "http://www.yahoo.com",
		Name: "Yahoo",
	}

	stmt := Link.
		INSERT(Link.MutableColumns).
		VALUES("http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		MODEL(google).
		MODELS([]model.Link{google, yahoo})

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestInsertQuery(t *testing.T) {
	initForInsertTest(t)

	stmt := Link.
		INSERT(Link.URL, Link.Name, Link.ID).
		QUERY(
			SELECT(Link.URL, Link.Name, Int(1000)).
				FROM(Link).
				WHERE(Link.ID.EQ(Int(1))),
		)

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)

	dest := []model.Link{}
	query := Link.
		SELECT(Link.AllColumns).
		WHERE(Link.Name.EQ(String("Youtube")))

	assertQueryRecordValues(t, query, &dest)
}

func TestInsertOnDuplicateKey(t *testing.T) {
	initForInsertTest(t)

	randId := 1000

	stmt := Link.INSERT().
		VALUES(randId, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		VALUES(randId, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		ON_DUPLICATE_KEY_UPDATE(
			Link.ID.SET(Link.ID.ADD(Int(11))),
			Link.Name.SET(String("PostgreSQL Tutorial 2")),
		)
	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)

	dest := []model.Link{}

	query := SELECT(Link.AllColumns).
		FROM(Link).
		WHERE(Link.ID.EQ(Int(int64(randId)).ADD(Int(11))))

	assertQueryRecordValues(t, query, &dest)
}

func TestInsertWithQueryContext(t *testing.T) {
	initForInsertTest(t)

	stmt := Link.INSERT().
		VALUES(1100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	dest := []model.Link{}
	err := stmt.QueryContext(ctx, db, &dest)

	require.Error(t, err, "context deadline exceeded")
}

func TestInsertWithExecContext(t *testing.T) {
	initForInsertTest(t)

	stmt := Link.INSERT().
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	_, err := stmt.ExecContext(ctx, db)

	require.Error(t, err, "context deadline exceeded")
}

func initForInsertTest(t *testing.T) {
	_, err := Link.DELETE().WHERE(Link.ID.GT(Int(1))).Exec(db)
	utils.PanicOnError(err)
}

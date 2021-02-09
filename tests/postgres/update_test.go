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

func TestUpdateValues(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	t.Run("deprecated version", func(t *testing.T) {
		stmt := Link.
			UPDATE(Link.Name, Link.URL).
			SET("Bong", "http://bong.com").
			WHERE(Link.Name.EQ(String("Bing")))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 1)

		links := []model.Link{}

		selQuery := Link.
			SELECT(Link.AllColumns).
			WHERE(Link.Name.IN(String("Bong")))

		err := selQuery.Query(db, &links)

		require.NoError(t, err)
		require.EqualValues(t, testparrot.RecordNext(t, links), links)
	})

	t.Run("new version", func(t *testing.T) {
		stmt := Link.UPDATE().
			SET(
				Link.Name.SET(String("DuckDuckGo")),
				Link.URL.SET(String("www.duckduckgo.com")),
			).
			WHERE(Link.Name.EQ(String("Yahoo")))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 1)
	})
}

func TestUpdateWithSubQueries(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	t.Run("deprecated version", func(t *testing.T) {
		stmt := Link.
			UPDATE(Link.Name, Link.URL).
			SET(
				SELECT(String("Bong")),
				SELECT(Link.URL).
					FROM(Link).
					WHERE(Link.Name.EQ(String("Bing"))),
			).
			WHERE(Link.Name.EQ(String("Bing")))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 1)
	})

	t.Run("new version", func(t *testing.T) {
		stmt := Link.UPDATE().
			SET(
				Link.Name.SET(String("Bong")),
				Link.URL.SET(StringExp(
					SELECT(Link.URL).
						FROM(Link).
						WHERE(Link.Name.EQ(String("Bing")))),
				),
			).
			WHERE(Link.Name.EQ(String("Bing")))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		_, err := stmt.Exec(db)
		require.NoError(t, err)
	})
}

func TestUpdateAndReturning(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	stmt := Link.
		UPDATE(Link.Name, Link.URL).
		SET("DuckDuckGo", "http://www.duckduckgo.com").
		WHERE(Link.Name.EQ(String("Ask"))).
		RETURNING(Link.AllColumns)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	links := []model.Link{}

	require.NoError(t, stmt.Query(db, &links))
	require.EqualValues(t, testparrot.RecordNext(t, links), links)
}

func TestUpdateWithSelect(t *testing.T) {
	t.Run("deprecated version", func(t *testing.T) {
		stmt := Link.UPDATE(Link.AllColumns).
			SET(
				Link.
					SELECT(Link.AllColumns).
					WHERE(Link.ID.EQ(Int(0))),
			).
			WHERE(Link.ID.EQ(Int(0)))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 1)
	})

	t.Run("new version", func(t *testing.T) {
		stmt := Link.UPDATE().
			SET(
				Link.MutableColumns.SET(
					SELECT(Link.MutableColumns).
						FROM(Link).
						WHERE(Link.ID.EQ(Int(0))),
				),
			).
			WHERE(Link.ID.EQ(Int(0)))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 1)
	})
}

func TestUpdateWithInvalidSelect(t *testing.T) {
	t.Run("deprecated version", func(t *testing.T) {
		stmt := Link.UPDATE(Link.AllColumns).
			SET(
				Link.
					SELECT(Link.ID, Link.Name).
					WHERE(Link.ID.EQ(Int(0))),
			).
			WHERE(Link.ID.EQ(Int(0)))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		_, err := stmt.Exec(db)
		require.Error(t, err, "pq: number of columns does not match number of values")
	})

	t.Run("new version", func(t *testing.T) {
		stmt := Link.UPDATE().
			SET(Link.AllColumns.SET(Link.SELECT(Link.MutableColumns))).
			WHERE(Link.ID.EQ(Int(0)))

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		_, err := stmt.Exec(db)
		require.Error(t, err, "pq: number of columns does not match number of values")
	})
}

func TestUpdateWithModelData(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	link := model.Link{
		ID:   201,
		URL:  "http://www.duckduckgo.com",
		Name: "DuckDuckGo",
	}

	stmt := Link.
		UPDATE(Link.AllColumns).
		MODEL(link).
		WHERE(Link.ID.EQ(Int(int64(link.ID))))

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 1)
}

func TestUpdateWithModelDataAndPredefinedColumnList(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	link := model.Link{
		ID:   201,
		URL:  "http://www.duckduckgo.com",
		Name: "DuckDuckGo",
	}

	updateColumnList := ColumnList{Link.Description, Link.Name, Link.URL}

	stmt := Link.
		UPDATE(updateColumnList).
		MODEL(link).
		WHERE(Link.ID.EQ(Int(int64(link.ID))))

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 1)
}

func TestUpdateWithInvalidModelData(t *testing.T) {
	defer func() {
		r := recover()

		require.Equal(t, r, "missing struct field for column : id")
	}()

	setupLinkTableForUpdateTest(t)

	link := struct {
		Ident       int
		URL         string
		Name        string
		Description *string
		Rel         *string
	}{
		Ident: 201,
		URL:   "http://www.duckduckgo.com",
		Name:  "DuckDuckGo",
	}

	stmt := Link.
		UPDATE(Link.AllColumns).
		MODEL(link).
		WHERE(Link.ID.EQ(Int(int64(link.Ident))))

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	_, err := stmt.Exec(db)
	require.Error(t, err, "pq: number of columns does not match number of values")
}

func TestUpdateQueryContext(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	updateStmt := Link.
		UPDATE(Link.Name, Link.URL).
		SET("Bong", "http://bong.com").
		WHERE(Link.Name.EQ(String("Bing")))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	dest := []model.Link{}
	err := updateStmt.QueryContext(ctx, db, &dest)

	require.Error(t, err, "context deadline exceeded")
}

func TestUpdateExecContext(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	updateStmt := Link.
		UPDATE(Link.Name, Link.URL).
		SET("Bong", "http://bong.com").
		WHERE(Link.Name.EQ(String("Bing")))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	_, err := updateStmt.ExecContext(ctx, db)

	require.Error(t, err, "context deadline exceeded")
}

func setupLinkTableForUpdateTest(t *testing.T) {
	cleanUpLinkTable(t)

	_, err := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
		VALUES(200, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		VALUES(201, "http://www.ask.com", "Ask", DEFAULT).
		VALUES(202, "http://www.ask.com", "Ask", DEFAULT).
		VALUES(203, "http://www.yahoo.com", "Yahoo", DEFAULT).
		VALUES(204, "http://www.bing.com", "Bing", DEFAULT).
		Exec(db)

	require.NoError(t, err)
}

func cleanUpLinkTable(t *testing.T) {
	_, err := Link.DELETE().WHERE(Link.ID.GT(Int(0))).Exec(db)
	require.NoError(t, err)
}

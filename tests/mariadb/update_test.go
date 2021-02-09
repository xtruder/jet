package mariadb

import (
	"context"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/utils"
	. "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/tests/mysql/gen/dvds/table"
	"github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/test_sample/table"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateValues(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	t.Run("old version", func(t *testing.T) {
		stmt := Link.
			UPDATE(Link.Name, Link.URL).
			SET("Bong", "http://bong.com").
			WHERE(Link.Name.EQ(String("Bing")))

		assertStatementRecordSQL(t, stmt)
		assertExec(t, stmt, db)
	})

	t.Run("new version", func(t *testing.T) {
		stmt := Link.UPDATE().
			SET(
				Link.Name.SET(String("Bong")),
				Link.URL.SET(String("http://bong.com")),
			).
			WHERE(Link.Name.EQ(String("Bing")))

		assertStatementRecordSQL(t, stmt)
		assertExec(t, stmt, db)
	})

	links := []model.Link{}

	err := Link.
		SELECT(Link.AllColumns).
		WHERE(Link.Name.EQ(String("Bong"))).
		Query(db, &links)
	utils.PanicOnError(err)

	require.Len(t, links, 1)
	assert.EqualValues(t, links[0], model.Link{
		ID:   204,
		URL:  "http://bong.com",
		Name: "Bong",
	})
}

func TestUpdateWithSubQueries(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	t.Run("old version", func(t *testing.T) {
		stmt := Link.
			UPDATE(Link.Name, Link.URL).
			SET(
				SELECT(String("Bong")),
				SELECT(Link2.URL).
					FROM(Link2).
					WHERE(Link2.Name.EQ(String("Youtube"))),
			).
			WHERE(Link.Name.EQ(String("Bing")))

		assertStatementRecordSQL(t, stmt)
		//assertExec(t, stmt, db)
	})

	t.Run("new version", func(t *testing.T) {
		stmt := Link.
			UPDATE().
			SET(
				Link.Name.SET(StringExp(SELECT(String("Bong")))),
				Link.URL.SET(StringExp(
					SELECT(Link2.URL).
						FROM(Link2).
						WHERE(Link2.Name.EQ(String("Youtube"))),
				)),
			).
			WHERE(Link.Name.EQ(String("Bing")))

		assertStatementRecordSQL(t, stmt)
		//assertExec(t, stmt, db)
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

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
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

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestUpdateWithModelDataAndMutableColumns(t *testing.T) {
	setupLinkTableForUpdateTest(t)

	link := model.Link{
		ID:   201,
		URL:  "http://www.duckduckgo.com",
		Name: "DuckDuckGo",
	}

	stmt := Link.
		UPDATE(Link.MutableColumns).
		MODEL(link).
		WHERE(Link.ID.EQ(Int(int64(link.ID))))

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, db)
}

func TestUpdateWithInvalidModelData(t *testing.T) {
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

	assert.PanicsWithValue(t, "missing struct field for column : id", func() {
		Link.
			UPDATE(Link.AllColumns).
			MODEL(link).
			WHERE(Link.ID.EQ(Int(int64(link.Ident))))
	})
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

	assert.Error(t, err, "context deadline exceeded")
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

	assert.Error(t, err, "context deadline exceeded")
}

func TestUpdateWithJoin(t *testing.T) {
	query := table.Staff.
		INNER_JOIN(table.Address, table.Address.AddressID.EQ(table.Staff.AddressID)).
		UPDATE(table.Staff.LastName).
		SET(String("New name")).
		WHERE(table.Staff.StaffID.EQ(Int(1)))

	_, err := query.Exec(db)
	assert.NoError(t, err)
}

func setupLinkTableForUpdateTest(t *testing.T) {
	_, err := Link.DELETE().WHERE(Link.ID.GT(Int(1))).Exec(db)
	utils.PanicOnError(err)

	_, err = Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
		VALUES(200, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		VALUES(201, "http://www.ask.com", "Ask", DEFAULT).
		VALUES(202, "http://www.ask.com", "Ask", DEFAULT).
		VALUES(203, "http://www.yahoo.com", "Yahoo", DEFAULT).
		VALUES(204, "http://www.bing.com", "Bing", DEFAULT).
		Exec(db)
	utils.PanicOnError(err)
}

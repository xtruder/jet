package postgres

import (
	"context"
	"math/rand"
	"testing"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/table"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestInsertValues(t *testing.T) {
	cleanUpLinkTable(t)

	stmt := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		VALUES(101, "http://www.google.com", "Google", DEFAULT).
		VALUES(102, "http://www.yahoo.com", "Yahoo", nil).
		RETURNING(Link.AllColumns)

	inserted := []model.Link{}

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	require.NoError(t, stmt.Query(db, &inserted))
	require.Len(t, inserted, 3)
	require.EqualValues(t, testparrot.RecordNext(t, inserted), inserted)

	retrieved := []model.Link{}

	require.NoError(t, Link.SELECT(Link.AllColumns).
		WHERE(Link.ID.GT_EQ(Int(100))).
		ORDER_BY(Link.ID).
		Query(db, &retrieved))

	require.EqualValues(t, inserted, retrieved)
}

func TestInsertEmptyColumnList(t *testing.T) {
	cleanUpLinkTable(t)

	stmt := Link.INSERT().
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 1)
}

func TestInsertOnConflict(t *testing.T) {
	t.Run("do nothing", func(t *testing.T) {
		cleanUpLinkTable(t)

		employee := model.Employee{EmployeeID: rand.Int31()}

		stmt := Employee.INSERT(Employee.AllColumns).
			MODEL(employee).
			MODEL(employee).
			ON_CONFLICT(Employee.EmployeeID).DO_NOTHING()

		sql, _ := stmt.Sql()
		require.Equal(t, testparrot.RecordNext(t, sql), sql)

		assertExec(t, stmt, 1)
	})

	t.Run("on constraint do nothing", func(t *testing.T) {
		employee := model.Employee{EmployeeID: rand.Int31()}

		stmt := Employee.INSERT(Employee.AllColumns).
			MODEL(employee).
			MODEL(employee).
			ON_CONFLICT().ON_CONSTRAINT("employee_pkey").DO_NOTHING()

		sql, _ := stmt.Sql()
		require.Equal(t, testparrot.RecordNext(t, sql), sql)

		assertExec(t, stmt, 1)
	})

	t.Run("do update", func(t *testing.T) {
		cleanUpLinkTable(t)

		stmt := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
			VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
			VALUES(200, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
			ON_CONFLICT(Link.ID).DO_UPDATE(
			SET(
				Link.ID.SET(Link.EXCLUDED.ID),
				Link.URL.SET(String("http://www.postgresqltutorial2.com")),
			),
		).RETURNING(Link.AllColumns)

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 2)
	})

	t.Run("on constraint do update", func(t *testing.T) {
		cleanUpLinkTable(t)

		stmt := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
			VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
			VALUES(200, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
			ON_CONFLICT().ON_CONSTRAINT("link_pkey").DO_UPDATE(
			SET(
				Link.ID.SET(Link.EXCLUDED.ID),
				Link.URL.SET(String("http://www.postgresqltutorial2.com")),
			),
		).RETURNING(Link.AllColumns)

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 2)
	})

	t.Run("do update complex", func(t *testing.T) {
		cleanUpLinkTable(t)

		stmt := Link.INSERT(Link.ID, Link.URL, Link.Name, Link.Description).
			VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
			ON_CONFLICT(Link.ID).WHERE(Link.ID.MUL(Int(2)).GT(Int(10))).DO_UPDATE(
			SET(
				Link.ID.SET(
					IntExp(SELECT(MAXi(Link.ID).ADD(Int(1))).
						FROM(Link)),
				),
				ColumnList{Link.Name, Link.Description}.SET(ROW(Link.EXCLUDED.Name, String("new description"))),
			).WHERE(Link.Description.IS_NOT_NULL()),
		)

		require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

		assertExec(t, stmt, 1)
	})
}

func TestInsertModelObject(t *testing.T) {
	cleanUpLinkTable(t)

	linkData := model.Link{
		URL:  "http://www.duckduckgo.com",
		Name: "Duck Duck go",
	}

	stmt := Link.
		INSERT(Link.URL, Link.Name).
		MODEL(linkData)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 1)
}

func TestInsertModelObjectEmptyColumnList(t *testing.T) {
	cleanUpLinkTable(t)

	linkData := model.Link{
		ID:   1000,
		URL:  "http://www.duckduckgo.com",
		Name: "Duck Duck go",
	}

	stmt := Link.
		INSERT().
		MODEL(linkData)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 1)
}

func TestInsertModelsObject(t *testing.T) {
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

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 3)
}

func TestInsertUsingMutableColumns(t *testing.T) {
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

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	assertExec(t, stmt, 4)
}

func TestInsertQuery(t *testing.T) {
	cleanUpLinkTable(t)

	stmt := Link.
		INSERT(Link.URL, Link.Name).
		QUERY(
			SELECT(Link.URL, Link.Name).
				FROM(Link).
				WHERE(Link.ID.EQ(Int(0))),
		).
		RETURNING(Link.AllColumns)

	dest := []model.Link{}

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	require.NoError(t, stmt.Query(db, &dest))

	youtubeLinks := []model.Link{}

	require.NoError(t, Link.
		SELECT(Link.AllColumns).
		WHERE(Link.Name.EQ(String("Youtube"))).
		Query(db, &youtubeLinks))
	require.Len(t, youtubeLinks, 2)
}

func TestInsertWithQueryContext(t *testing.T) {
	cleanUpLinkTable(t)

	stmt := Link.INSERT().
		VALUES(1100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT).
		RETURNING(Link.AllColumns)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	dest := []model.Link{}
	err := stmt.QueryContext(ctx, db, &dest)

	require.Error(t, err, "context deadline exceeded")
}

func TestInsertWithExecContext(t *testing.T) {
	cleanUpLinkTable(t)

	stmt := Link.INSERT().
		VALUES(100, "http://www.postgresqltutorial.com", "PostgreSQL Tutorial", DEFAULT)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	_, err := stmt.ExecContext(ctx, db)

	require.Error(t, err, "context deadline exceeded")
}

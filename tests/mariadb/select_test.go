package mariadb

import (
	"testing"

	. "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/tests/mysql/gen/dvds/enum"
	"github.com/go-jet/jet/v2/tests/mysql/gen/dvds/model"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/dvds/table"
	"github.com/go-jet/jet/v2/tests/mysql/gen/dvds/view"
)

func TestSelect_ScanToStruct(t *testing.T) {
	query := Actor.
		SELECT(Actor.AllColumns).
		DISTINCT().
		WHERE(Actor.ActorID.EQ(Int(2))).
		LIMIT(2)

	dest := model.Actor{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestSelect_ScanToSlice(t *testing.T) {
	query := Actor.
		SELECT(Actor.AllColumns).
		ORDER_BY(Actor.ActorID).
		LIMIT(10)

	dest := []model.Actor{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestSelectGroupByHaving(t *testing.T) {
	query := Payment.
		INNER_JOIN(Customer, Customer.CustomerID.EQ(Payment.CustomerID)).
		SELECT(
			Customer.AllColumns,

			SUMf(Payment.Amount).AS("amount.sum"),
			AVG(Payment.Amount).AS("amount.avg"),
			MAX(Payment.PaymentDate).AS("amount.max_date"),
			MAXf(Payment.Amount).AS("amount.max"),
			MIN(Payment.PaymentDate).AS("amount.min_date"),
			MINf(Payment.Amount).AS("amount.min"),
			COUNT(Payment.Amount).AS("amount.count"),
		).
		GROUP_BY(Payment.CustomerID).
		HAVING(
			SUMf(Payment.Amount).GT(Float(125.6)),
		).
		ORDER_BY(
			Payment.CustomerID, SUMf(Payment.Amount).ASC(),
		).
		LIMIT(10)

	dest := []struct {
		model.Customer

		Amount struct {
			Sum   float64
			Avg   float64
			Max   float64
			Min   float64
			Count int64
		} `alias:"amount"`
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestSubQuery(t *testing.T) {
	rRatingFilms := Film.
		SELECT(
			Film.FilmID,
			Film.Title,
			Film.Rating,
		).
		WHERE(Film.Rating.EQ(enum.FilmRating.R)).
		AsTable("rFilms")

	rFilmID := Film.FilmID.From(rRatingFilms)

	query := rRatingFilms.
		INNER_JOIN(FilmActor, FilmActor.FilmID.EQ(rFilmID)).
		INNER_JOIN(Actor, Actor.ActorID.EQ(FilmActor.ActorID)).
		SELECT(
			Actor.AllColumns,
			FilmActor.AllColumns,
			rRatingFilms.AllColumns(),
		).
		ORDER_BY(rFilmID, Actor.ActorID).
		LIMIT(10)

	dest := []struct {
		model.Film

		Actors []model.Actor
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestSelectAndUnionInProjection(t *testing.T) {
	query := Payment.
		SELECT(
			Payment.PaymentID,
			Customer.SELECT(Customer.CustomerID).LIMIT(1),
			UNION(
				Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(10),
				Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(2),
			).LIMIT(1),
		).
		LIMIT(12)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestSelectUNION(t *testing.T) {
	query := UNION(
		Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(10),
		Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(2),
	).LIMIT(1)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestSelectUNION_ALL(t *testing.T) {
	query := UNION_ALL(
		Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(10),
		Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(2),
	).ORDER_BY(Payment.PaymentID).
		LIMIT(4).
		OFFSET(3)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)

	query = Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(10).
		UNION_ALL(Payment.SELECT(Payment.PaymentID).LIMIT(1).OFFSET(2)).
		ORDER_BY(Payment.PaymentID).
		LIMIT(4).
		OFFSET(3)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestJoinQueryStruct(t *testing.T) {
	query := Language.
		INNER_JOIN(Film, Film.LanguageID.EQ(Language.LanguageID)).
		INNER_JOIN(FilmActor, FilmActor.FilmID.EQ(Film.FilmID)).
		INNER_JOIN(Actor, Actor.ActorID.EQ(FilmActor.ActorID)).
		LEFT_JOIN(Inventory, Inventory.FilmID.EQ(Film.FilmID)).
		LEFT_JOIN(Rental, Rental.InventoryID.EQ(Inventory.InventoryID)).
		SELECT(
			FilmActor.AllColumns,
			Film.AllColumns,
			Language.AllColumns,
			Actor.AllColumns,
			Inventory.AllColumns,
			Rental.AllColumns,
		).
		ORDER_BY(
			Language.LanguageID.ASC(),
			Film.FilmID.ASC(),
			Actor.ActorID.ASC(),
			Inventory.InventoryID.ASC(),
			Rental.RentalID.ASC(),
		).
		LIMIT(10)

	dest := []struct {
		model.Language

		Films []struct {
			model.Film

			Actors []struct {
				model.Actor
			}

			Inventories []struct {
				model.Inventory

				Rentals []model.Rental
			}
		}
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestRowLock(t *testing.T) {
	getRowLockTestData := func() []RowLock {
		return []RowLock{
			UPDATE(),
		}
	}

	query := Address.
		SELECT(STAR).
		LIMIT(3).
		OFFSET(1)

	for _, lockType := range getRowLockTestData() {
		query.FOR(lockType)

		assertExec(t, query, db)

		db.Exec("UNLOCK TABLES;")
	}

	for _, lockType := range getRowLockTestData() {
		query.FOR(lockType.NOWAIT())

		assertExec(t, query, db)

		db.Exec("UNLOCK TABLES;")
	}
}

func TestExpressionWrappers(t *testing.T) {
	query := SELECT(
		BoolExp(Raw("true")),
		IntExp(Raw("11")),
		FloatExp(Raw("11.22")),
		StringExp(Raw("'stringer'")),
		TimeExp(Raw("'raw'")),
		TimestampExp(Raw("'raw'")),
		DateTimeExp(Raw("'raw'")),
		DateExp(Raw("'date'")),
	)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestLockInShareMode(t *testing.T) {
	query := Address.
		SELECT(STAR).
		LIMIT(3).
		OFFSET(1).
		LOCK_IN_SHARE_MODE()

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

func TestWindowClause(t *testing.T) {
	query := Payment.SELECT(
		AVG(Payment.Amount).OVER(),
		AVG(Payment.Amount).OVER(Window("w1")),
		AVG(Payment.Amount).OVER(
			Window("w2").
				ORDER_BY(Payment.CustomerID).
				RANGE(PRECEDING(UNBOUNDED), FOLLOWING(UNBOUNDED)),
		),
		AVG(Payment.Amount).OVER(Window("w3").RANGE(PRECEDING(UNBOUNDED), FOLLOWING(UNBOUNDED))),
	).
		WHERE(Payment.PaymentID.LT(Int(10))).
		WINDOW("w1").AS(PARTITION_BY(Payment.PaymentDate)).
		WINDOW("w2").AS(Window("w1")).
		WINDOW("w3").AS(Window("w2").ORDER_BY(Payment.CustomerID)).
		ORDER_BY(Payment.CustomerID)

	assertStatementRecordSQL(t, query)
	assertQuery(t, query)
}

type ActorInfo struct {
	ActorID   int
	FirstName string
	LastName  string
	FilmInfo  string
}

func TestSimpleView(t *testing.T) {
	query := SELECT(
		view.ActorInfo.AllColumns,
	).
		FROM(view.ActorInfo).
		ORDER_BY(view.ActorInfo.ActorID).
		LIMIT(10)

	dest := []ActorInfo{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestJoinViewWithTable(t *testing.T) {
	query := SELECT(
		view.CustomerList.AllColumns,
		Rental.AllColumns,
	).
		FROM(view.CustomerList.
			INNER_JOIN(Rental, view.CustomerList.ID.EQ(Rental.CustomerID)),
		).
		ORDER_BY(view.CustomerList.ID).
		WHERE(view.CustomerList.ID.LT_EQ(Int(2))).
		LIMIT(2)

	dest := []struct {
		model.CustomerList `sql:"primary_key=ID"`
		Rentals            []model.Rental
	}{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

func TestConditionalProjectionList(t *testing.T) {
	projectionList := ProjectionList{}

	columnsToSelect := []string{"customer_id", "create_date"}

	for _, columnName := range columnsToSelect {
		switch columnName {
		case Customer.CustomerID.Name():
			projectionList = append(projectionList, Customer.CustomerID)
		case Customer.Email.Name():
			projectionList = append(projectionList, Customer.Email)
		case Customer.CreateDate.Name():
			projectionList = append(projectionList, Customer.CreateDate)
		}
	}

	query := SELECT(projectionList).
		FROM(Customer).
		LIMIT(3)

	dest := []model.Customer{}

	assertStatementRecordSQL(t, query)
	assertQueryRecordValues(t, query, &dest)
}

package postgres

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/testutils"
	. "github.com/go-jet/jet/v2/postgres"
	qrm "github.com/go-jet/jet/v2/qrm"
	"github.com/go-jet/jet/v2/tests/postgres/gen/dvds/enum"
	"github.com/go-jet/jet/v2/tests/postgres/gen/dvds/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/dvds/table"
	"github.com/go-jet/jet/v2/tests/postgres/gen/dvds/view"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestSelect_ScanToStruct(t *testing.T) {
	query := Actor.
		SELECT(Actor.AllColumns).
		DISTINCT().
		WHERE(Actor.ActorID.EQ(Int(2)))

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	actor := model.Actor{}
	require.NoError(t, query.Query(db, &actor))
	require.EqualValues(t, testparrot.RecordNext(t, actor), actor)
}

func TestClassicSelect(t *testing.T) {
	query := SELECT(
		Payment.AllColumns,
		Customer.AllColumns,
	).
		FROM(Payment.
			INNER_JOIN(Customer, Payment.CustomerID.EQ(Customer.CustomerID))).
		ORDER_BY(Payment.PaymentID.ASC()).
		LIMIT(2)

	dest := []model.Payment{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestSelect_ScanToSlice(t *testing.T) {
	customers := []model.Customer{}
	query := Customer.SELECT(Customer.AllColumns).ORDER_BY(Customer.CustomerID.ASC())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &customers))
	require.Len(t, customers, 599)
	require.EqualValues(t, testparrot.RecordNext(t, customers[0]), customers[0])
	require.EqualValues(t, testparrot.RecordNext(t, customers[1]), customers[1])
	require.EqualValues(t, testparrot.RecordNext(t, customers[598]), customers[598])
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

	dest := []struct{}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
}

func TestJoinQueryStruct(t *testing.T) {
	query := FilmActor.
		INNER_JOIN(Actor, FilmActor.ActorID.EQ(Actor.ActorID)).
		INNER_JOIN(Film, FilmActor.FilmID.EQ(Film.FilmID)).
		INNER_JOIN(Language, Film.LanguageID.EQ(Language.LanguageID)).
		INNER_JOIN(Inventory, Inventory.FilmID.EQ(Film.FilmID)).
		INNER_JOIN(Rental, Rental.InventoryID.EQ(Inventory.InventoryID)).
		SELECT(
			FilmActor.AllColumns,
			Film.AllColumns,
			Language.AllColumns,
			Actor.AllColumns,
			Inventory.AllColumns,
			Rental.AllColumns,
		).
		ORDER_BY(Film.FilmID.ASC()).
		LIMIT(1000)

	languageActorFilm := []struct {
		model.Language

		Films []struct {
			model.Film
			Actors []struct {
				model.Actor
			}

			Inventory []struct {
				model.Inventory

				Rental []model.Rental
			}
		}
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &languageActorFilm))
	require.Len(t, languageActorFilm, 1)
	require.Len(t, languageActorFilm[0].Films, 10)
	require.Len(t, languageActorFilm[0].Films[0].Actors, 10)
}

func TestJoinQuerySlice(t *testing.T) {
	result := []struct {
		Language *model.Language
		Film     []model.Film
	}{}

	query := Film.
		INNER_JOIN(Language, Film.LanguageID.EQ(Language.LanguageID)).
		SELECT(Language.AllColumns, Film.AllColumns).
		WHERE(Film.Rating.EQ(enum.MpaaRating.Nc17)).
		LIMIT(15)

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &result))
	require.EqualValues(t, testparrot.RecordNext(t, result), result)

	resultPtr := []*struct {
		Language *model.Language
		Film     []model.Film
	}{}
	require.NoError(t, query.Query(db, &resultPtr))
	require.Equal(t, len(result), 1)
	require.Equal(t, len(result[0].Film), 15)
}

func TestExecution1(t *testing.T) {
	query := City.
		INNER_JOIN(Address, Address.CityID.EQ(City.CityID)).
		INNER_JOIN(Customer, Customer.AddressID.EQ(Address.AddressID)).
		SELECT(
			City.CityID,
			City.City,
			Address.AddressID,
			Address.Address,
			Customer.CustomerID,
			Customer.LastName,
		).
		WHERE(City.City.EQ(String("London")).OR(City.City.EQ(String("York")))).
		ORDER_BY(City.CityID, Address.AddressID, Customer.CustomerID)

	dest := []struct {
		model.City

		Customers []struct {
			model.Customer
			Address model.Address
		}
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestExecution2(t *testing.T) {
	query := City.
		INNER_JOIN(Address, Address.CityID.EQ(City.CityID)).
		INNER_JOIN(Customer, Customer.AddressID.EQ(Address.AddressID)).
		SELECT(
			City.CityID.AS("my_city.id"),
			City.City.AS("myCity.Name"),
			Address.AddressID.AS("My_Address.id"),
			Address.Address.AS("my address.address line"),
			Customer.CustomerID.AS("my_customer.id"),
			Customer.LastName.AS("my_customer.last_name"),
		).
		WHERE(City.City.EQ(String("London")).OR(City.City.EQ(String("York")))).
		ORDER_BY(City.CityID, Address.AddressID, Customer.CustomerID)

	dest := []struct {
		ID   int32 `sql:"primary_key"`
		Name string

		Customers []struct {
			ID       int32 `sql:"primary_key"`
			LastName *string

			Address struct {
				ID          int32 `sql:"primary_key"`
				AddressLine string
			}
		}
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestExecution3(t *testing.T) {
	query := City.
		INNER_JOIN(Address, Address.CityID.EQ(City.CityID)).
		INNER_JOIN(Customer, Customer.AddressID.EQ(Address.AddressID)).
		SELECT(
			City.CityID.AS("city_id"),
			City.City.AS("city_name"),
			Customer.CustomerID.AS("customer_id"),
			Customer.LastName.AS("last_name"),
			Address.AddressID.AS("address_id"),
			Address.Address.AS("address_line"),
		).
		WHERE(City.City.EQ(String("London")).OR(City.City.EQ(String("York")))).
		ORDER_BY(City.CityID, Address.AddressID, Customer.CustomerID)

	dest := []struct {
		CityID   int32 `sql:"primary_key"`
		CityName string

		Customers []struct {
			CustomerID int32 `sql:"primary_key"`
			LastName   *string

			Address struct {
				AddressID   int32 `sql:"primary_key"`
				AddressLine string
			}
		}
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestExecution4(t *testing.T) {
	query := City.
		INNER_JOIN(Address, Address.CityID.EQ(City.CityID)).
		INNER_JOIN(Customer, Customer.AddressID.EQ(Address.AddressID)).
		SELECT(
			City.CityID,
			City.City,
			Customer.CustomerID,
			Customer.LastName,
			Address.AddressID,
			Address.Address,
		).
		WHERE(City.City.EQ(String("London")).OR(City.City.EQ(String("York")))).
		ORDER_BY(City.CityID, Address.AddressID, Customer.CustomerID)

	dest := []struct {
		CityID   int32  `sql:"primary_key" alias:"city.city_id"`
		CityName string `alias:"city.city"`

		Customers []struct {
			CustomerID int32   `sql:"primary_key" alias:"customer_id"`
			LastName   *string `alias:"last_name"`

			Address struct {
				AddressID   int32  `sql:"primary_key" alias:"AddressId"`
				AddressLine string `alias:"address.address"`
			} `alias:"address.*"`
		} `alias:"customer"`
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestJoinQuerySliceWithPtrs(t *testing.T) {
	limit := int64(3)

	query := Film.INNER_JOIN(Language, Film.LanguageID.EQ(Language.LanguageID)).
		SELECT(Language.AllColumns, Film.AllColumns).
		LIMIT(limit)

	dest := []*struct {
		Language model.Language
		Film     *[]*model.Film
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, 1)
	require.Len(t, *dest[0].Film, int(limit))
}

func TestSelect_WithoutUniqueColumnSelected(t *testing.T) {
	query := Customer.SELECT(Customer.FirstName, Customer.LastName, Customer.Email)
	dest := []model.Customer{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, 599)
}

func TestSelectOrderByAscDesc(t *testing.T) {
	asc := []model.Customer{}

	query := Customer.SELECT(Customer.CustomerID, Customer.FirstName, Customer.LastName).
		ORDER_BY(Customer.FirstName.ASC())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &asc))

	firstAsc := asc[0]
	lastAsc := asc[len(asc)-1]

	desc := []model.Customer{}
	query = Customer.SELECT(Customer.CustomerID, Customer.FirstName, Customer.LastName).
		ORDER_BY(Customer.FirstName.DESC())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &desc))

	firstDesc := desc[0]
	lastDesc := desc[len(desc)-1]

	require.EqualValues(t, firstAsc, lastDesc)
	require.EqualValues(t, lastAsc, firstDesc)

	ascDesc := []model.Customer{}
	query = Customer.SELECT(Customer.CustomerID, Customer.FirstName, Customer.LastName).
		ORDER_BY(Customer.FirstName.ASC(), Customer.LastName.DESC())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &ascDesc))

	require.EqualValues(t, testparrot.RecordNext(t, ascDesc[326:328]), ascDesc[326:328])
}

func TestSelectFullJoin(t *testing.T) {
	query := Customer.
		FULL_JOIN(Address, Customer.AddressID.EQ(Address.AddressID)).
		SELECT(Customer.AllColumns, Address.AllColumns).
		ORDER_BY(Customer.CustomerID.ASC())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	dest := []struct {
		Address  *model.Address
		Customer *model.Customer
	}{}

	require.NoError(t, query.Query(db, &dest))
	require.Equal(t, len(dest), 603)
	require.EqualValues(t, testparrot.RecordNext(t, dest[len(dest)-1]), dest[len(dest)-1])
}

func TestSelectFullCrossJoin(t *testing.T) {
	query := Customer.
		CROSS_JOIN(Address).
		SELECT(Customer.AllColumns, Address.AllColumns).
		ORDER_BY(Customer.CustomerID.ASC()).
		LIMIT(1000)

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	dest := []struct {
		model.Customer
		model.Address
	}{}

	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, 1000)
	require.EqualValues(t, testparrot.RecordNext(t, dest[len(dest)-1]), dest[len(dest)-1])
}

func TestSelectSelfJoin(t *testing.T) {
	f1 := Film.AS("f1")
	f2 := Film.AS("f2")
	query := f1.
		INNER_JOIN(f2, f1.FilmID.LT(f2.FilmID).AND(f1.Length.EQ(f2.Length))).
		SELECT(f1.AllColumns, f2.AllColumns).
		ORDER_BY(f1.FilmID.ASC())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	dest := []struct {
		F1 model.Film `alias:"f1.*"`
		F2 model.Film `alias:"f2.*"`
	}{}

	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest[len(dest)-1]), dest[len(dest)-1])
}

func TestSelectAliasColumn(t *testing.T) {
	f1 := Film.AS("f1")
	f2 := Film.AS("f2")

	f1.FilmID.EQ(Int(11))

	query := f1.
		INNER_JOIN(f2, f1.FilmID.NOT_EQ(f2.FilmID).AND(f1.Length.EQ(f2.Length))).
		SELECT(f1.Title.AS("title1"),
			f2.Title.AS("title2"),
			f1.Length.AS("length")).
		ORDER_BY(f1.Length.ASC(), f1.Title.ASC(), f2.Title.ASC()).
		LIMIT(1000)

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	films := []struct {
		Title1 string
		Title2 string
		Length int16
	}{}

	require.NoError(t, query.Query(db, &films))
	require.Len(t, films, 1000)
	require.EqualValues(t, testparrot.RecordNext(t, films[0]), films[0])
}

func TestSubQuery(t *testing.T) {
	rRatingFilms := Film.
		SELECT(
			Film.FilmID,
			Film.Title,
			Film.Rating,
		).
		WHERE(Film.Rating.EQ(enum.MpaaRating.R)).
		AsTable("rFilms")

	rFilmID := Film.FilmID.From(rRatingFilms)

	query := Actor.
		INNER_JOIN(FilmActor, Actor.ActorID.EQ(FilmActor.FilmID)).
		INNER_JOIN(rRatingFilms, FilmActor.FilmID.EQ(rFilmID)).
		SELECT(
			Actor.AllColumns,
			FilmActor.AllColumns,
			rRatingFilms.AllColumns(),
		)

	dest := []model.Actor{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
}

func TestSelectFunctions(t *testing.T) {
	query := Film.SELECT(
		MAXf(Film.RentalRate).AS("max_film_rate"),
	)

	result := struct{ MaxFilmRate float64 }{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &result))
	require.EqualValues(t, testparrot.RecordNext(t, result), result)
}

func TestSelectQueryScalar(t *testing.T) {
	maxFilmRentalRate := FloatExp(
		Film.
			SELECT(MAXf(Film.RentalRate)),
	)

	query := Film.
		SELECT(Film.AllColumns).
		WHERE(Film.RentalRate.EQ(maxFilmRentalRate)).
		ORDER_BY(Film.FilmID.ASC())

	dest := []model.Film{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
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
		GROUP_BY(Customer.CustomerID).
		HAVING(
			SUMf(Payment.Amount).GT(Float(125.6)),
		).
		ORDER_BY(
			Customer.CustomerID, SUMf(Payment.Amount).ASC(),
		)

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

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
}

func TestSelectGroupBy2(t *testing.T) {
	customersPayments := Payment.
		SELECT(
			Payment.CustomerID,
			SUMf(Payment.Amount).AS("amount_sum"),
		).
		GROUP_BY(Payment.CustomerID).
		AsTable("customer_payment_sum")

	customerID := Payment.CustomerID.From(customersPayments)
	amountSum := FloatColumn("amount_sum").From(customersPayments)

	query := Customer.
		INNER_JOIN(customersPayments, Customer.CustomerID.EQ(customerID)).
		SELECT(
			Customer.AllColumns,
			amountSum.AS("CustomerWithAmounts.AmountSum"),
		).
		ORDER_BY(amountSum.ASC())

	dest := []struct {
		Customer  *model.Customer
		AmountSum float64
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
}

func TestSelectStaff(t *testing.T) {
	query := Staff.SELECT(Staff.AllColumns)
	dest := []model.Staff{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

type JSONResult struct{ qrm.JSON }

func TestSelectJoinJSON(t *testing.T) {
	query := Film.
		SELECT(
			JSONB_BUILD_OBJECT(
				String("title"), Film.Title,
				String("description"), Film.Description,
				String("rating"), Film.Rating,
				String("length"), Film.Length,
				String("actors"), Actor.
					LEFT_JOIN(FilmActor, FilmActor.ActorID.EQ(Actor.ActorID)).
					SELECT(
						JSON_AGG(Actor.FirstName.CONCAT(String(" ")).CONCAT(Actor.LastName)),
					).WHERE(Film.FilmID.EQ(FilmActor.FilmID)),
				String("language"), Language.
					SELECT(
						CAST(
							ROW_TO_JSON(Language.TableName()),
						).AS_JSONB().DELETE_KEY(String("language_id")),
					).
					WHERE(Film.LanguageID.EQ(Language.LanguageID)).
					LIMIT(1),
			).AS("json"),
		).
		LIMIT(2)

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	result := []JSONResult{}

	err := query.Query(db, &result)
	require.NoError(t, err)

	require.Len(t, result, 2)
	require.EqualValues(t, testparrot.RecordNext(t, result), result)

	m := map[string]interface{}{}
	json.Unmarshal(result[0].JSON, &m)
	require.EqualValues(t, testparrot.RecordNext(t, m), m)
}

func TestSelectTimeColumns(t *testing.T) {
	query := Payment.SELECT(Payment.AllColumns).
		WHERE(Payment.PaymentDate.LT(Timestamp(2007, time.February, 14, 22, 16, 01, 0))).
		ORDER_BY(Payment.PaymentDate.ASC())

	dest := []model.Payment{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
}

func TestUnion(t *testing.T) {
	query := UNION_ALL(
		Payment.
			SELECT(Payment.PaymentID.AS("payment.payment_id"), Payment.Amount).
			WHERE(Payment.Amount.LT_EQ(Float(100))),
		Payment.
			SELECT(Payment.PaymentID, Payment.Amount).
			WHERE(Payment.Amount.GT_EQ(Float(200))),
	).
		ORDER_BY(IntegerColumn("payment.payment_id").ASC(), Payment.Amount.DESC()).
		LIMIT(10).
		OFFSET(20)

	dest := []model.Payment{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])
}

func TestAllSetOperators(t *testing.T) {
	var select1 = Payment.SELECT(Payment.AllColumns).
		WHERE(Payment.PaymentID.GT_EQ(Int(17600)).AND(Payment.PaymentID.LT(Int(17610))))
	var select2 = Payment.SELECT(Payment.AllColumns).
		WHERE(Payment.PaymentID.GT_EQ(Int(17620)).AND(Payment.PaymentID.LT(Int(17630))))

	t.Run("UNION", func(t *testing.T) {
		query := select1.UNION(select2)

		dest := []model.Payment{}
		err := query.Query(db, &dest)

		require.NoError(t, err)
		require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	})

	t.Run("UNION_ALL", func(t *testing.T) {
		query := select1.UNION_ALL(select2)

		dest := []model.Payment{}
		err := query.Query(db, &dest)

		require.NoError(t, err)
		require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	})

	t.Run("INTERSECT", func(t *testing.T) {
		query := select1.INTERSECT(select2)

		dest := []model.Payment{}
		err := query.Query(db, &dest)

		require.NoError(t, err)
		require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	})

	t.Run("INTERSECT_ALL", func(t *testing.T) {
		query := select1.INTERSECT_ALL(select2)

		dest := []model.Payment{}
		err := query.Query(db, &dest)

		require.NoError(t, err)
		require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	})

	t.Run("EXCEPT", func(t *testing.T) {
		query := select1.EXCEPT(select2)

		dest := []model.Payment{}
		err := query.Query(db, &dest)

		require.NoError(t, err)
		require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	})

	t.Run("EXCEPT_ALL", func(t *testing.T) {
		query := select1.EXCEPT_ALL(select2)

		dest := []model.Payment{}
		err := query.Query(db, &dest)

		require.NoError(t, err)
		require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	})
}

func TestSelectWithCase(t *testing.T) {
	query := Payment.SELECT(
		CASE(Payment.StaffID).
			WHEN(Int(1)).THEN(String("ONE")).
			WHEN(Int(2)).THEN(String("TWO")).
			WHEN(Int(3)).THEN(String("THREE")).
			ELSE(String("OTHER")).AS("staff_id_num"),
	).
		ORDER_BY(Payment.PaymentID.ASC()).
		LIMIT(20)

	dest := []struct{ StaffIDNum string }{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])
}

func TestRowLock(t *testing.T) {
	query := Address.
		SELECT(STAR).
		LIMIT(3)

	getLockTypes := func() []RowLock {
		return []RowLock{
			UPDATE(),
			NO_KEY_UPDATE(),
			SHARE(),
			KEY_SHARE(),
		}
	}

	for _, lockType := range getLockTypes() {
		query.FOR(lockType)

		require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

		tx, _ := db.Begin()

		res, err := query.Exec(tx)
		require.NoError(t, err)
		rowsAffected, _ := res.RowsAffected()
		require.Equal(t, rowsAffected, int64(3))

		err = tx.Rollback()
		require.NoError(t, err)
	}

	for _, lockType := range getLockTypes() {
		query.FOR(lockType.NOWAIT())

		require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

		tx, _ := db.Begin()

		res, err := query.Exec(tx)
		require.NoError(t, err)
		rowsAffected, _ := res.RowsAffected()
		require.Equal(t, rowsAffected, int64(3))

		err = tx.Rollback()
		require.NoError(t, err)
	}

	for _, lockType := range getLockTypes() {
		query.FOR(lockType.SKIP_LOCKED())

		require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

		tx, _ := db.Begin()

		res, err := query.Exec(tx)
		require.NoError(t, err)
		rowsAffected, _ := res.RowsAffected()
		require.Equal(t, rowsAffected, int64(3))

		err = tx.Rollback()
		require.NoError(t, err)
	}
}

func TestQuickStart(t *testing.T) {
	query := SELECT(
		Actor.ActorID, Actor.FirstName, Actor.LastName, Actor.LastUpdate, // list of all actor columns (equivalent to Actor.AllColumns)
		Film.AllColumns, // list of all film columns (equivalent to Film.FilmID, Film.Title, ...)
		Language.AllColumns,
		Category.AllColumns,
	).FROM(
		Actor.
			INNER_JOIN(FilmActor, Actor.ActorID.EQ(FilmActor.ActorID)). // INNER JOIN Actor with FilmActor on condition Actor.ActorID = FilmActor.ActorID
			INNER_JOIN(Film, Film.FilmID.EQ(FilmActor.FilmID)).         // then with Film, Language, FilmCategory and Category.
			INNER_JOIN(Language, Language.LanguageID.EQ(Film.LanguageID)).
			INNER_JOIN(FilmCategory, FilmCategory.FilmID.EQ(Film.FilmID)).
			INNER_JOIN(Category, Category.CategoryID.EQ(FilmCategory.CategoryID)),
	).WHERE(
		Language.Name.EQ(String("English")). // note that every column has type.
							AND(Category.Name.NOT_EQ(String("Action"))). // String column Language.Name and Category.Name can be compared only with string expression
							AND(Film.Length.GT(Int(180))),               // Film.Length is integer column and can be compared only with integer expression
	).ORDER_BY(
		Actor.ActorID.ASC(),
		Film.FilmID.ASC(),
	)

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	dest := []struct {
		model.Actor

		Films []struct {
			model.Film

			Language model.Language

			Categories []model.Category
		}
	}{}

	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])

	dest2 := []struct {
		model.Category

		Films  []model.Film
		Actors []model.Actor
	}{}

	require.NoError(t, query.Query(db, &dest2))
	require.EqualValues(t, testparrot.RecordNext(t, dest2[0:2]), dest2[0:2])
}

func TestQuickStartWithSubQueries(t *testing.T) {
	filmLogerThan180 := Film.
		SELECT(Film.AllColumns).
		WHERE(Film.Length.GT(Int(180))).
		AsTable("films")

	filmID := Film.FilmID.From(filmLogerThan180)
	filmLanguageID := Film.LanguageID.From(filmLogerThan180)

	categoriesNotAction := Category.
		SELECT(Category.AllColumns).
		WHERE(Category.Name.NOT_EQ(String("Action"))).
		AsTable("categories")

	categoryID := Category.CategoryID.From(categoriesNotAction)

	query := Actor.
		INNER_JOIN(FilmActor, Actor.ActorID.EQ(FilmActor.ActorID)).
		INNER_JOIN(filmLogerThan180, filmID.EQ(FilmActor.FilmID)).
		INNER_JOIN(Language, Language.LanguageID.EQ(filmLanguageID)).
		INNER_JOIN(FilmCategory, FilmCategory.FilmID.EQ(filmID)).
		INNER_JOIN(categoriesNotAction, categoryID.EQ(FilmCategory.CategoryID)).
		SELECT(
			Actor.AllColumns,
			filmLogerThan180.AllColumns(),
			Language.AllColumns,
			categoriesNotAction.AllColumns(),
		).ORDER_BY(
		Actor.ActorID.ASC(),
		filmID.ASC(),
	)

	dest := []struct {
		model.Actor

		Films []struct {
			model.Film

			Language model.Language

			Categories []model.Category
		}
	}{}

	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])

	dest2 := []struct {
		model.Category

		Films  []model.Film
		Actors []model.Actor
	}{}

	require.NoError(t, query.Query(db, &dest2))
	require.EqualValues(t, testparrot.RecordNext(t, dest2[0:2]), dest2[0:2])
}

func TestExpressionWrappers(t *testing.T) {
	query := SELECT(
		BoolExp(Raw("true")),
		IntExp(Raw("11")),
		FloatExp(Raw("11.22")),
		StringExp(Raw("'stringer'")),
		TimeExp(Raw("'raw'")),
		TimezExp(Raw("'raw'")),
		TimestampExp(Raw("'raw'")),
		TimestampzExp(Raw("'raw'")),
		DateExp(Raw("'date'")),
		JSONExp(Raw(`'{"key": "value"}'`)),
		JSONBExp(Raw(`'{"key": "value"}'`)),
	)

	dest := []struct{}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
}

func TestWindowFunction(t *testing.T) {
	query := Payment.
		SELECT(
			AVG(Payment.Amount).OVER(),
			AVG(Payment.Amount).OVER(PARTITION_BY(Payment.CustomerID)),
			MAXf(Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate.DESC())),
			MINf(Payment.Amount).OVER(PARTITION_BY(Payment.CustomerID).ORDER_BY(Payment.PaymentDate.DESC())),
			SUMf(Payment.Amount).OVER(PARTITION_BY(Payment.CustomerID).
				ORDER_BY(Payment.PaymentDate.DESC()).ROWS(PRECEDING(1), FOLLOWING(6))),
			SUMf(Payment.Amount).OVER(PARTITION_BY(Payment.CustomerID).
				ORDER_BY(Payment.PaymentDate.DESC()).RANGE(PRECEDING(UNBOUNDED), FOLLOWING(UNBOUNDED))),
			MAXi(Payment.CustomerID).OVER(ORDER_BY(Payment.PaymentDate.DESC()).ROWS(CURRENT_ROW, FOLLOWING(UNBOUNDED))),
			MINi(Payment.CustomerID).OVER(PARTITION_BY(Payment.CustomerID).ORDER_BY(Payment.PaymentDate.DESC())),
			SUMi(Payment.CustomerID).OVER(PARTITION_BY(Payment.CustomerID).ORDER_BY(Payment.PaymentDate.DESC())),
			ROW_NUMBER().OVER(ORDER_BY(Payment.PaymentDate)),
			RANK().OVER(ORDER_BY(Payment.PaymentDate)),
			DENSE_RANK().OVER(ORDER_BY(Payment.PaymentDate)),
			CUME_DIST().OVER(ORDER_BY(Payment.PaymentDate)),
			NTILE(11).OVER(ORDER_BY(Payment.PaymentDate)),
			LAG(Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate)),
			LAG(Payment.Amount, 2).OVER(ORDER_BY(Payment.PaymentDate)),
			LAG(Payment.Amount, 2, Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate)),
			LAG(Payment.Amount, 2, 100).OVER(ORDER_BY(Payment.PaymentDate)),
			LEAD(Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate)),
			LEAD(Payment.Amount, 2).OVER(ORDER_BY(Payment.PaymentDate)),
			LEAD(Payment.Amount, 2, Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate)),
			LEAD(Payment.Amount, 2, 100).OVER(ORDER_BY(Payment.PaymentDate)),
			FIRST_VALUE(Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate)),
			LAST_VALUE(Payment.Amount).OVER(ORDER_BY(Payment.PaymentDate)),
			NTH_VALUE(Payment.Amount, 3).OVER(ORDER_BY(Payment.PaymentDate)),
		).GROUP_BY(Payment.Amount, Payment.CustomerID, Payment.PaymentDate).
		WHERE(Payment.PaymentID.LT(Int(10)))

	dest := []struct{}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
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

	dest := []struct{}{}
	//fmt.Println(query.Sql())

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
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
	).FROM(view.ActorInfo).
		ORDER_BY(view.ActorInfo.ActorID).
		LIMIT(10)

	dest := []ActorInfo{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])
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
		WHERE(view.CustomerList.ID.LT_EQ(Int(2)))

	dest := []struct {
		model.CustomerList `sql:"primary_key=ID"`
		Rentals            []model.Rental
	}{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.Len(t, dest[0].Rentals, testparrot.RecordNext(t, len(dest[0].Rentals)).(int))
	require.Len(t, dest[1].Rentals, testparrot.RecordNext(t, len(dest[1].Rentals)).(int))
}

func TestDynamicProjectionList(t *testing.T) {
	request := struct {
		ColumnsToSelect []string
		ShowFullName    bool
	}{
		ColumnsToSelect: []string{"customer_id", "create_date"},
		ShowFullName:    true,
	}

	projectionList := ProjectionList{}

	for _, columnName := range request.ColumnsToSelect {
		switch columnName {
		case Customer.CustomerID.Name():
			projectionList = append(projectionList, Customer.CustomerID)
		case Customer.Email.Name():
			projectionList = append(projectionList, Customer.Email)
		case Customer.CreateDate.Name():
			projectionList = append(projectionList, Customer.CreateDate)
		}
	}

	var showFullName bool
	if showFullName {
		projectionList = append(projectionList, Customer.FirstName.CONCAT(Customer.LastName))
	}

	query := SELECT(projectionList).
		FROM(Customer).
		LIMIT(3)

	dest := []model.Customer{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
}

func TestDynamicCondition(t *testing.T) {
	request := struct {
		CustomerID *int64
		Email      *string
		Active     *bool
	}{
		CustomerID: testutils.Ptr(int64(1)).(*int64),
		Active:     testutils.Ptr(true).(*bool),
	}

	condition := Bool(true)

	if request.CustomerID != nil {
		condition = condition.AND(Customer.CustomerID.EQ(Int(*request.CustomerID)))
	}
	if request.Email != nil {
		condition = condition.AND(Customer.Email.EQ(String(*request.Email)))
	}
	if request.Active != nil {
		condition = condition.AND(Customer.Activebool.EQ(Bool(*request.Active)))
	}

	query := SELECT(Customer.AllColumns).
		FROM(Customer).
		WHERE(condition)

	dest := []model.Customer{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
}

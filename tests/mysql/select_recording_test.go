// Code generated by testparrot. DO NOT EDIT.

package mysql

import (
	model "github.com/go-jet/jet/v2/tests/mysql/gen/dvds/model"
	gotestparrot "github.com/xtruder/go-testparrot"
	"time"
)

func init() {
	gotestparrot.R.Load("TestConditionalProjectionList", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT customer.customer_id AS "customer.customer_id",
     customer.create_date AS "customer.create_date"
FROM dvds.customer
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(3)},
	}, {
		Key: 2,
		Value: []model.Customer{{
			CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
			CustomerID: uint16(0x1),
		}, {
			CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
			CustomerID: uint16(0x2),
		}, {
			CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
			CustomerID: uint16(0x3),
		}},
	}})
	gotestparrot.R.Load("TestExpressionWrappers", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT true,
     11,
     11.22,
     'stringer',
     'raw',
     'raw',
     'raw',
     'date';
`,
	}, {
		Key:   1,
		Value: []interface{}{},
	}})
	gotestparrot.R.Load("TestJoinQueryStruct", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT film_actor.actor_id AS "film_actor.actor_id",
     film_actor.film_id AS "film_actor.film_id",
     film_actor.last_update AS "film_actor.last_update",
     film.film_id AS "film.film_id",
     film.title AS "film.title",
     film.description AS "film.description",
     film.release_year AS "film.release_year",
     film.language_id AS "film.language_id",
     film.original_language_id AS "film.original_language_id",
     film.rental_duration AS "film.rental_duration",
     film.rental_rate AS "film.rental_rate",
     film.length AS "film.length",
     film.replacement_cost AS "film.replacement_cost",
     film.rating AS "film.rating",
     film.special_features AS "film.special_features",
     film.last_update AS "film.last_update",
     language.language_id AS "language.language_id",
     language.name AS "language.name",
     language.last_update AS "language.last_update",
     actor.actor_id AS "actor.actor_id",
     actor.first_name AS "actor.first_name",
     actor.last_name AS "actor.last_name",
     actor.last_update AS "actor.last_update",
     inventory.inventory_id AS "inventory.inventory_id",
     inventory.film_id AS "inventory.film_id",
     inventory.store_id AS "inventory.store_id",
     inventory.last_update AS "inventory.last_update",
     rental.rental_id AS "rental.rental_id",
     rental.rental_date AS "rental.rental_date",
     rental.inventory_id AS "rental.inventory_id",
     rental.customer_id AS "rental.customer_id",
     rental.return_date AS "rental.return_date",
     rental.staff_id AS "rental.staff_id",
     rental.last_update AS "rental.last_update"
FROM dvds.language
     INNER JOIN dvds.film ON (film.language_id = language.language_id)
     INNER JOIN dvds.film_actor ON (film_actor.film_id = film.film_id)
     INNER JOIN dvds.actor ON (actor.actor_id = film_actor.actor_id)
     LEFT JOIN dvds.inventory ON (inventory.film_id = film.film_id)
     LEFT JOIN dvds.rental ON (rental.inventory_id = inventory.inventory_id)
ORDER BY language.language_id ASC, film.film_id ASC, actor.actor_id ASC, inventory.inventory_id ASC, rental.rental_id ASC
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(10)},
	}, {
		Key: 2,
		Value: []struct {
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
		}{{
			Films: []struct {
				model.Film
				Actors []struct {
					model.Actor
				}
				Inventories []struct {
					model.Inventory
					Rentals []model.Rental
				}
			}{{
				Actors: []struct {
					model.Actor
				}{{Actor: model.Actor{
					ActorID:    uint16(0x1),
					FirstName:  "PENELOPE",
					LastName:   "GUINESS",
					LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
				}}},
				Film: model.Film{
					Description:     gotestparrot.Ptr("A Epic Drama of a Feminist And a Mad Scientist who must Battle a Teacher in The Canadian Rockies").(*string),
					FilmID:          uint16(0x1),
					LanguageID:      uint8(0x1),
					LastUpdate:      gotestparrot.Decode("2006-02-15T05:03:42Z", time.Time{}).(time.Time),
					Length:          gotestparrot.Ptr(uint16(0x56)).(*uint16),
					Rating:          gotestparrot.Ptr(model.FilmRating("PG")).(*model.FilmRating),
					ReleaseYear:     gotestparrot.Ptr(int16(2006)).(*int16),
					RentalDuration:  uint8(0x6),
					RentalRate:      0.99,
					ReplacementCost: 20.99,
					SpecialFeatures: gotestparrot.Ptr("Deleted Scenes,Behind the Scenes").(*string),
					Title:           "ACADEMY DINOSAUR",
				},
				Inventories: []struct {
					model.Inventory
					Rentals []model.Rental
				}{{
					Inventory: model.Inventory{
						FilmID:      uint16(0x1),
						InventoryID: uint32(0x1),
						LastUpdate:  gotestparrot.Decode("2006-02-15T05:09:17Z", time.Time{}).(time.Time),
						StoreID:     uint8(0x1),
					},
					Rentals: []model.Rental{{
						CustomerID:  uint16(0x1af),
						InventoryID: uint32(0x1),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-07-08T19:03:15Z", time.Time{}).(time.Time),
						RentalID:    int32(4863),
						ReturnDate:  gotestparrot.Decode("2005-07-11T21:29:15Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x2),
					}, {
						CustomerID:  uint16(0x206),
						InventoryID: uint32(0x1),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-08-02T20:13:10Z", time.Time{}).(time.Time),
						RentalID:    int32(11433),
						ReturnDate:  gotestparrot.Decode("2005-08-11T21:35:10Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}, {
						CustomerID:  uint16(0x117),
						InventoryID: uint32(0x1),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-08-21T21:27:43Z", time.Time{}).(time.Time),
						RentalID:    int32(14714),
						ReturnDate:  gotestparrot.Decode("2005-08-30T22:26:43Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}},
				}, {
					Inventory: model.Inventory{
						FilmID:      uint16(0x1),
						InventoryID: uint32(0x2),
						LastUpdate:  gotestparrot.Decode("2006-02-15T05:09:17Z", time.Time{}).(time.Time),
						StoreID:     uint8(0x1),
					},
					Rentals: []model.Rental{{
						CustomerID:  uint16(0x19b),
						InventoryID: uint32(0x2),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-05-30T20:21:07Z", time.Time{}).(time.Time),
						RentalID:    int32(972),
						ReturnDate:  gotestparrot.Decode("2005-06-06T00:36:07Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}, {
						CustomerID:  uint16(0xaa),
						InventoryID: uint32(0x2),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-06-17T20:24:00Z", time.Time{}).(time.Time),
						RentalID:    int32(2117),
						ReturnDate:  gotestparrot.Decode("2005-06-23T17:45:00Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x2),
					}, {
						CustomerID:  uint16(0xa1),
						InventoryID: uint32(0x2),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-07-07T10:41:31Z", time.Time{}).(time.Time),
						RentalID:    int32(4187),
						ReturnDate:  gotestparrot.Decode("2005-07-11T06:25:31Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}, {
						CustomerID:  uint16(0x245),
						InventoryID: uint32(0x2),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-07-30T22:02:34Z", time.Time{}).(time.Time),
						RentalID:    int32(9449),
						ReturnDate:  gotestparrot.Decode("2005-08-06T02:09:34Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}, {
						CustomerID:  uint16(0x167),
						InventoryID: uint32(0x2),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-08-23T01:01:01Z", time.Time{}).(time.Time),
						RentalID:    int32(15453),
						ReturnDate:  gotestparrot.Decode("2005-08-30T20:08:01Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}},
				}, {
					Inventory: model.Inventory{
						FilmID:      uint16(0x1),
						InventoryID: uint32(0x3),
						LastUpdate:  gotestparrot.Decode("2006-02-15T05:09:17Z", time.Time{}).(time.Time),
						StoreID:     uint8(0x1),
					},
					Rentals: []model.Rental{{
						CustomerID:  uint16(0x27),
						InventoryID: uint32(0x3),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-07-31T21:36:07Z", time.Time{}).(time.Time),
						RentalID:    int32(10126),
						ReturnDate:  gotestparrot.Decode("2005-08-03T23:59:07Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x1),
					}, {
						CustomerID:  uint16(0x21d),
						InventoryID: uint32(0x3),
						LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
						RentalDate:  gotestparrot.Decode("2005-08-22T23:56:37Z", time.Time{}).(time.Time),
						RentalID:    int32(15421),
						ReturnDate:  gotestparrot.Decode("2005-08-25T18:58:37Z", &time.Time{}).(*time.Time),
						StaffID:     uint8(0x2),
					}},
				}},
			}},
			Language: model.Language{
				LanguageID: uint8(0x1),
				LastUpdate: gotestparrot.Decode("2006-02-15T05:02:19Z", time.Time{}).(time.Time),
				Name:       "English",
			},
		}},
	}})
	gotestparrot.R.Load("TestJoinViewWithTable", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT customer_list.` + "`" + `ID` + "`" + ` AS "customer_list.ID",
     customer_list.name AS "customer_list.name",
     customer_list.address AS "customer_list.address",
     customer_list.` + "`" + `zip code` + "`" + ` AS "customer_list.zip code",
     customer_list.phone AS "customer_list.phone",
     customer_list.city AS "customer_list.city",
     customer_list.country AS "customer_list.country",
     customer_list.notes AS "customer_list.notes",
     customer_list.` + "`" + `SID` + "`" + ` AS "customer_list.SID",
     rental.rental_id AS "rental.rental_id",
     rental.rental_date AS "rental.rental_date",
     rental.inventory_id AS "rental.inventory_id",
     rental.customer_id AS "rental.customer_id",
     rental.return_date AS "rental.return_date",
     rental.staff_id AS "rental.staff_id",
     rental.last_update AS "rental.last_update"
FROM dvds.customer_list
     INNER JOIN dvds.rental ON (customer_list.` + "`" + `ID` + "`" + ` = rental.customer_id)
WHERE customer_list.` + "`" + `ID` + "`" + ` <= ?
ORDER BY customer_list.` + "`" + `ID` + "`" + `
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(2), int64(2)},
	}, {
		Key: 2,
		Value: []struct {
			model.CustomerList `sql:"primary_key=ID"`
			Rentals            []model.Rental
		}{{
			CustomerList: model.CustomerList{
				Address: "1913 Hanoi Way",
				City:    "Sasebo",
				Country: "Japan",
				ID:      uint16(0x1),
				Name:    gotestparrot.Ptr("MARY SMITH").(*string),
				Notes:   "active",
				Phone:   "28303384290",
				Sid:     uint8(0x1),
				ZipCode: gotestparrot.Ptr("35200").(*string),
			},
			Rentals: []model.Rental{{
				CustomerID:  uint16(0x1),
				InventoryID: uint32(0xbcd),
				LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
				RentalDate:  gotestparrot.Decode("2005-05-25T11:30:37Z", time.Time{}).(time.Time),
				RentalID:    int32(76),
				ReturnDate:  gotestparrot.Decode("2005-06-03T12:00:37Z", &time.Time{}).(*time.Time),
				StaffID:     uint8(0x2),
			}, {
				CustomerID:  uint16(0x1),
				InventoryID: uint32(0xfb4),
				LastUpdate:  gotestparrot.Decode("2006-02-15T21:30:53Z", time.Time{}).(time.Time),
				RentalDate:  gotestparrot.Decode("2005-05-28T10:35:23Z", time.Time{}).(time.Time),
				RentalID:    int32(573),
				ReturnDate:  gotestparrot.Decode("2005-06-03T06:32:23Z", &time.Time{}).(*time.Time),
				StaffID:     uint8(0x1),
			}},
		}},
	}})
	gotestparrot.R.Load("TestLockInShareMode", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT *
FROM dvds.address
LIMIT ?
OFFSET ?
LOCK IN SHARE MODE;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(3), int64(1)},
	}})
	gotestparrot.R.Load("TestSelectAndUnionInProjection", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT payment.payment_id AS "payment.payment_id",
     (
          SELECT customer.customer_id AS "customer.customer_id"
          FROM dvds.customer
          LIMIT ?
     ),
     (
          (
               SELECT payment.payment_id AS "payment.payment_id"
               FROM dvds.payment
               LIMIT ?
               OFFSET ?
          )
          UNION
          (
               SELECT payment.payment_id AS "payment.payment_id"
               FROM dvds.payment
               LIMIT ?
               OFFSET ?
          )
          LIMIT ?
     )
FROM dvds.payment
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(1), int64(1), int64(10), int64(1), int64(2), int64(1), int64(12)},
	}})
	gotestparrot.R.Load("TestSelectGroupByHaving", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT customer.customer_id AS "customer.customer_id",
     customer.store_id AS "customer.store_id",
     customer.first_name AS "customer.first_name",
     customer.last_name AS "customer.last_name",
     customer.email AS "customer.email",
     customer.address_id AS "customer.address_id",
     customer.active AS "customer.active",
     customer.create_date AS "customer.create_date",
     customer.last_update AS "customer.last_update",
     SUM(payment.amount) AS "amount.sum",
     AVG(payment.amount) AS "amount.avg",
     MAX(payment.payment_date) AS "amount.max_date",
     MAX(payment.amount) AS "amount.max",
     MIN(payment.payment_date) AS "amount.min_date",
     MIN(payment.amount) AS "amount.min",
     COUNT(payment.amount) AS "amount.count"
FROM dvds.payment
     INNER JOIN dvds.customer ON (customer.customer_id = payment.customer_id)
GROUP BY payment.customer_id
HAVING SUM(payment.amount) > ?
ORDER BY payment.customer_id, SUM(payment.amount) ASC
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{125.6, int64(10)},
	}, {
		Key: 2,
		Value: []struct {
			model.Customer
			Amount struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			} `alias:"amount"`
		}{{
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.767778,
				Count: int64(27),
				Max:   10.99,
				Min:   0.99,
				Sum:   128.73,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x6),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x2),
				Email:      gotestparrot.Ptr("PATRICIA.JOHNSON@dvdscustomer.org").(*string),
				FirstName:  "PATRICIA",
				LastName:   "JOHNSON",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   5.220769,
				Count: int64(26),
				Max:   10.99,
				Min:   0.99,
				Sum:   135.74,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x7),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x3),
				Email:      gotestparrot.Ptr("LINDA.WILLIAMS@dvdscustomer.org").(*string),
				FirstName:  "LINDA",
				LastName:   "WILLIAMS",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   3.805789,
				Count: int64(38),
				Max:   9.99,
				Min:   0.99,
				Sum:   144.62,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x9),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x5),
				Email:      gotestparrot.Ptr("ELIZABETH.BROWN@dvdscustomer.org").(*string),
				FirstName:  "ELIZABETH",
				LastName:   "BROWN",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.596061,
				Count: int64(33),
				Max:   8.99,
				Min:   0.99,
				Sum:   151.67,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0xb),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x7),
				Email:      gotestparrot.Ptr("MARIA.MILLER@dvdscustomer.org").(*string),
				FirstName:  "MARIA",
				LastName:   "MILLER",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.878889,
				Count: int64(27),
				Max:   11.99,
				Min:   0.99,
				Sum:   131.73,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x11),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0xd),
				Email:      gotestparrot.Ptr("KAREN.JACKSON@dvdscustomer.org").(*string),
				FirstName:  "KAREN",
				LastName:   "JACKSON",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x2),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.20875,
				Count: int64(32),
				Max:   8.99,
				Sum:   134.68,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x13),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0xf),
				Email:      gotestparrot.Ptr("HELEN.HARRIS@dvdscustomer.org").(*string),
				FirstName:  "HELEN",
				LastName:   "HARRIS",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   5.24,
				Count: int64(24),
				Max:   9.99,
				Min:   0.99,
				Sum:   125.76,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x17),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x13),
				Email:      gotestparrot.Ptr("RUTH.MARTINEZ@dvdscustomer.org").(*string),
				FirstName:  "RUTH",
				LastName:   "MARTINEZ",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.447143,
				Count: int64(35),
				Max:   10.99,
				Min:   0.99,
				Sum:   155.65,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x19),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x15),
				Email:      gotestparrot.Ptr("MICHELLE.CLARK@dvdscustomer.org").(*string),
				FirstName:  "MICHELLE",
				LastName:   "CLARK",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x1),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.49,
				Count: int64(34),
				Max:   9.99,
				Min:   0.99,
				Sum:   152.66,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x1e),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x1a),
				Email:      gotestparrot.Ptr("JESSICA.HALL@dvdscustomer.org").(*string),
				FirstName:  "JESSICA",
				LastName:   "HALL",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x2),
			},
		}, {
			Amount: struct {
				Sum   float64
				Avg   float64
				Max   float64
				Min   float64
				Count int64
			}{
				Avg:   4.086774,
				Count: int64(31),
				Max:   8.99,
				Min:   0.99,
				Sum:   126.69,
			},
			Customer: model.Customer{
				Active:     true,
				AddressID:  uint16(0x1f),
				CreateDate: gotestparrot.Decode("2006-02-14T22:04:36Z", time.Time{}).(time.Time),
				CustomerID: uint16(0x1b),
				Email:      gotestparrot.Ptr("SHIRLEY.ALLEN@dvdscustomer.org").(*string),
				FirstName:  "SHIRLEY",
				LastName:   "ALLEN",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:57:20Z", &time.Time{}).(*time.Time),
				StoreID:    uint8(0x2),
			},
		}},
	}})
	gotestparrot.R.Load("TestSelectUNION", []gotestparrot.Recording{{
		Key: 0,
		Value: `
(
     SELECT payment.payment_id AS "payment.payment_id"
     FROM dvds.payment
     LIMIT ?
     OFFSET ?
)
UNION
(
     SELECT payment.payment_id AS "payment.payment_id"
     FROM dvds.payment
     LIMIT ?
     OFFSET ?
)
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(1), int64(10), int64(1), int64(2), int64(1)},
	}})
	gotestparrot.R.Load("TestSelectUNION_ALL", []gotestparrot.Recording{{
		Key: 0,
		Value: `
(
     SELECT payment.payment_id AS "payment.payment_id"
     FROM dvds.payment
     LIMIT ?
     OFFSET ?
)
UNION ALL
(
     SELECT payment.payment_id AS "payment.payment_id"
     FROM dvds.payment
     LIMIT ?
     OFFSET ?
)
ORDER BY "payment.payment_id"
LIMIT ?
OFFSET ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(1), int64(10), int64(1), int64(2), int64(4), int64(3)},
	}, {
		Key: 2,
		Value: `
(
     SELECT payment.payment_id AS "payment.payment_id"
     FROM dvds.payment
     LIMIT ?
     OFFSET ?
)
UNION ALL
(
     SELECT payment.payment_id AS "payment.payment_id"
     FROM dvds.payment
     LIMIT ?
     OFFSET ?
)
ORDER BY "payment.payment_id"
LIMIT ?
OFFSET ?;
`,
	}, {
		Key:   3,
		Value: []interface{}{int64(1), int64(10), int64(1), int64(2), int64(4), int64(3)},
	}})
	gotestparrot.R.Load("TestSelect_ScanToSlice", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT actor.actor_id AS "actor.actor_id",
     actor.first_name AS "actor.first_name",
     actor.last_name AS "actor.last_name",
     actor.last_update AS "actor.last_update"
FROM dvds.actor
ORDER BY actor.actor_id
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(10)},
	}, {
		Key: 2,
		Value: []model.Actor{{
			ActorID:    uint16(0x1),
			FirstName:  "PENELOPE",
			LastName:   "GUINESS",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x2),
			FirstName:  "NICK",
			LastName:   "WAHLBERG",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x3),
			FirstName:  "ED",
			LastName:   "CHASE",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x4),
			FirstName:  "JENNIFER",
			LastName:   "DAVIS",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x5),
			FirstName:  "JOHNNY",
			LastName:   "LOLLOBRIGIDA",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x6),
			FirstName:  "BETTE",
			LastName:   "NICHOLSON",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x7),
			FirstName:  "GRACE",
			LastName:   "MOSTEL",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x8),
			FirstName:  "MATTHEW",
			LastName:   "JOHANSSON",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0x9),
			FirstName:  "JOE",
			LastName:   "SWANK",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}, {
			ActorID:    uint16(0xa),
			FirstName:  "CHRISTIAN",
			LastName:   "GABLE",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		}},
	}})
	gotestparrot.R.Load("TestSelect_ScanToStruct", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT DISTINCT actor.actor_id AS "actor.actor_id",
     actor.first_name AS "actor.first_name",
     actor.last_name AS "actor.last_name",
     actor.last_update AS "actor.last_update"
FROM dvds.actor
WHERE actor.actor_id = ?
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(2), int64(2)},
	}, {
		Key: 2,
		Value: model.Actor{
			ActorID:    uint16(0x2),
			FirstName:  "NICK",
			LastName:   "WAHLBERG",
			LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
		},
	}})
	gotestparrot.R.Load("TestSimpleView", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT actor_info.actor_id AS "actor_info.actor_id",
     actor_info.first_name AS "actor_info.first_name",
     actor_info.last_name AS "actor_info.last_name",
     actor_info.film_info AS "actor_info.film_info"
FROM dvds.actor_info
ORDER BY actor_info.actor_id
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(10)},
	}, {
		Key: 2,
		Value: []ActorInfo{{
			ActorID:   1,
			FilmInfo:  "Animation: ANACONDA CONFESSIONS; Children: LANGUAGE COWBOY; Classics: COLOR PHILADELPHIA, WESTWARD SEABISCUIT; Comedy: VERTIGO NORTHWEST; Documentary: ACADEMY DINOSAUR; Family: KING EVOLUTION, SPLASH GUMP; Foreign: MULHOLLAND BEAST; Games: BULWORTH COMMANDMENTS, HUMAN GRAFFITI; Horror: ELEPHANT TROJAN, LADY STAGE, RULES HUMAN; Music: WIZARD COLDBLOODED; New: ANGELS LIFE, OKLAHOMA JUMANJI; Sci-Fi: CHEAPER CLYDE; Sports: GLEAMING JAWBREAKER",
			FirstName: "PENELOPE",
			LastName:  "GUINESS",
		}, {
			ActorID:   2,
			FilmInfo:  "Action: BULL SHAWSHANK; Animation: FIGHT JAWBREAKER; Children: JERSEY SASSY; Classics: DRACULA CRYSTAL, GILBERT PELICAN; Comedy: MALLRATS UNITED, RUSHMORE MERMAID; Documentary: ADAPTATION HOLES; Drama: WARDROBE PHANTOM; Family: APACHE DIVINE, CHISUM BEHAVIOR, INDIAN LOVE, MAGUIRE APACHE; Foreign: BABY HALL, HAPPINESS UNITED; Games: ROOF CHAMPION; Music: LUCKY FLYING; New: DESTINY SATURDAY, FLASH WARS, JEKYLL FROGMEN, MASK PEACH; Sci-Fi: CHAINSAW UPTOWN, GOODFELLAS SALUTE; Travel: LIAISONS SWEET, SMILE EARRING",
			FirstName: "NICK",
			LastName:  "WAHLBERG",
		}, {
			ActorID:   3,
			FilmInfo:  "Action: CADDYSHACK JEDI, FORREST SONS; Classics: FROST HEAD, JEEPERS WEDDING; Documentary: ARMY FLINTSTONES, FRENCH HOLIDAY, HALLOWEEN NUTS, HUNTER ALTER, WEDDING APOLLO, YOUNG LANGUAGE; Drama: LUCK OPUS, NECKLACE OUTBREAK, SPICE SORORITY; Foreign: COWBOY DOOM, WHALE BIKINI; Music: ALONE TRIP; New: EVE RESURRECTION, PLATOON INSTINCT; Sci-Fi: WEEKEND PERSONAL; Sports: ARTIST COLDBLOODED, IMAGE PRINCESS; Travel: BOONDOCK BALLROOM",
			FirstName: "ED",
			LastName:  "CHASE",
		}, {
			ActorID:   4,
			FilmInfo:  "Action: BAREFOOT MANCHURIAN; Animation: ANACONDA CONFESSIONS, GHOSTBUSTERS ELF; Comedy: SUBMARINE BED; Documentary: BED HIGHBALL, NATIONAL STORY, RAIDERS ANTITRUST; Drama: BLADE POLISH, GREEDY ROOTS; Family: SPLASH GUMP; Horror: TREASURE COMMAND; Music: HANOVER GALAXY, REDS POCUS; New: ANGELS LIFE, JUMANJI BLADE, OKLAHOMA JUMANJI; Sci-Fi: RANDOM GO, SILVERADO GOLDFINGER, UNFORGIVEN ZOOLANDER; Sports: INSTINCT AIRPORT, POSEIDON FOREVER; Travel: BOONDOCK BALLROOM",
			FirstName: "JENNIFER",
			LastName:  "DAVIS",
		}, {
			ActorID:   5,
			FilmInfo:  "Action: AMADEUS HOLY, GRAIL FRANKENSTEIN, RINGS HEARTBREAKERS; Animation: SUNRISE LEAGUE; Children: HALL CASSIDY; Comedy: DADDY PITTSBURGH; Documentary: BONNIE HOLOCAUST, METAL ARMAGEDDON, PACIFIC AMISTAD, POCUS PULP; Drama: CHITTY LOCK, CONEHEADS SMOOCHY; Games: FIRE WOLVES; Horror: COMMANDMENTS EXPRESS, LOVE SUICIDES, PATTON INTERVIEW; Music: BANGER PINOCCHIO, HEAVENLY GUN; New: FRONTIER CABIN, RIDGEMONT SUBMARINE; Sci-Fi: DAISY MENAGERIE, GOODFELLAS SALUTE, SOLDIERS EVOLUTION; Sports: GROOVE FICTION, KRAMER CHOCOLATE, STAR OPERATION; Travel: ENOUGH RAGING, ESCAPE METROPOLIS, SMILE EARRING",
			FirstName: "JOHNNY",
			LastName:  "LOLLOBRIGIDA",
		}, {
			ActorID:   6,
			FilmInfo:  "Action: ANTITRUST TOMATOES; Animation: BIKINI BORROWERS, CROSSROADS CASUALTIES, POTLUCK MIXED, TITANIC BOONDOCK; Children: LANGUAGE COWBOY; Classics: BEAST HUNCHBACK; Documentary: COAST RAINBOW; Family: BANG KWAI; Foreign: CALENDAR GUNFIGHT, MULHOLLAND BEAST; New: WYOMING STORM; Sci-Fi: COLDBLOODED DARLING; Sports: DROP WATERFRONT, KRAMER CHOCOLATE, LESSON CLEOPATRA, LIBERTY MAGNIFICENT, TRADING PINOCCHIO; Travel: IGBY MAKER, SPEED SUIT",
			FirstName: "BETTE",
			LastName:  "NICHOLSON",
		}, {
			ActorID:   7,
			FilmInfo:  "Action: BERETS AGENT, EXCITEMENT EVE; Animation: SLEEPLESS MONSOON, TRACY CIDER; Children: WARLOCK WEREWOLF; Classics: MALKOVICH PET, OCTOBER SUBMARINE; Drama: CONFESSIONS MAGUIRE, DECEIVER BETRAYED, DESTINATION JERK, NECKLACE OUTBREAK, SAINTS BRIDE, SAVANNAH TOWN, TREATMENT JEKYLL; Foreign: COMMAND DARLING, HELLFIGHTERS SIERRA, SCISSORHANDS SLUMS, TOWN ARK, WAR NOTTING; Games: DAZED PUNK; Horror: ARACHNOPHOBIA ROLLERCOASTER, GASLIGHT CRUSADE; New: ANGELS LIFE, BREAKING HOME, SLEEPY JAPANESE, STING PERSONAL; Sci-Fi: OPEN AFRICAN; Sports: ANONYMOUS HUMAN, INSTINCT AIRPORT, POSEIDON FOREVER",
			FirstName: "GRACE",
			LastName:  "MOSTEL",
		}, {
			ActorID:   8,
			FilmInfo:  "Action: CAMPUS REMEMBER, DANCES NONE; Animation: SUGAR WONKA; Classics: LIGHTS DEER, MALKOVICH PET, TOMORROW HUSTLER; Drama: CONQUERER NUTS, HANGING DEEP, SCORPION APOLLO; Family: INDIAN LOVE; Foreign: BABY HALL, SCHOOL JACKET; Music: CLONES PINOCCHIO, DRIVING POLISH, RUNNER MADIGAN, VANISHING ROCKY; New: FLASH WARS; Sci-Fi: CROWDS TELEMARK; Sports: DURHAM PANKY, LOSER HUSTLER",
			FirstName: "MATTHEW",
			LastName:  "JOHANSSON",
		}, {
			ActorID:   9,
			FilmInfo:  "Action: PRIMARY GLASS, WATERFRONT DELIVERANCE; Animation: LAWLESS VISION, SUNRISE LEAGUE; Children: CROOKED FROGMEN, SWEETHEARTS SUSPECTS, TIES HUNGER; Classics: SNATCHERS MONTEZUMA; Documentary: MAJESTIC FLOATS, PACIFIC AMISTAD, UNTOUCHABLES SUNRISE; Drama: DALMATIONS SWEDEN, LEBOWSKI SOLDIERS; Family: CHOCOLAT HARRY; Foreign: CHOCOLATE DUCK; Games: CURTAIN VIDEOTAPE; Horror: ANYTHING SAVANNAH, REEF SALUTE; Music: BIRCH ANTITRUST, RUNNER MADIGAN; New: WILD APOLLO; Sports: PERDITION FARGO; Travel: HORROR REIGN, SMILE EARRING, TRAFFIC HOBBIT",
			FirstName: "JOE",
			LastName:  "SWANK",
		}, {
			ActorID:   10,
			FilmInfo:  "Action: LORD ARIZONA, WATERFRONT DELIVERANCE; Animation: PUNK DIVORCE; Children: CROOKED FROGMEN; Classics: JEEPERS WEDDING, PREJUDICE OLEANDER; Comedy: LIFE TWISTED; Documentary: ACADEMY DINOSAUR, MOD SECRETARY, WEDDING APOLLO; Drama: GOLDFINGER SENSIBILITY; Foreign: USUAL UNTOUCHABLES; Games: DIVINE RESURRECTION; Horror: ALABAMA DEVIL, REAP UNFAITHFUL; Music: JAWBREAKER BROOKLYN, WIZARD COLDBLOODED, WON DARES; New: DRAGONFLY STRANGERS; Sci-Fi: VACATION BOONDOCK; Sports: SHAKESPEARE SADDLE; Travel: TROUBLE DATE",
			FirstName: "CHRISTIAN",
			LastName:  "GABLE",
		}},
	}})
	gotestparrot.R.Load("TestSubQuery", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT actor.actor_id AS "actor.actor_id",
     actor.first_name AS "actor.first_name",
     actor.last_name AS "actor.last_name",
     actor.last_update AS "actor.last_update",
     film_actor.actor_id AS "film_actor.actor_id",
     film_actor.film_id AS "film_actor.film_id",
     film_actor.last_update AS "film_actor.last_update",
     ` + "`" + `rFilms` + "`" + `.` + "`" + `film.film_id` + "`" + ` AS "film.film_id",
     ` + "`" + `rFilms` + "`" + `.` + "`" + `film.title` + "`" + ` AS "film.title",
     ` + "`" + `rFilms` + "`" + `.` + "`" + `film.rating` + "`" + ` AS "film.rating"
FROM (
          SELECT film.film_id AS "film.film_id",
               film.title AS "film.title",
               film.rating AS "film.rating"
          FROM dvds.film
          WHERE film.rating = 'R'
     ) AS ` + "`" + `rFilms` + "`" + `
     INNER JOIN dvds.film_actor ON (film_actor.film_id = ` + "`" + `rFilms` + "`" + `.` + "`" + `film.film_id` + "`" + `)
     INNER JOIN dvds.actor ON (actor.actor_id = film_actor.actor_id)
ORDER BY ` + "`" + `rFilms` + "`" + `.` + "`" + `film.film_id` + "`" + `, actor.actor_id
LIMIT ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(10)},
	}, {
		Key: 2,
		Value: []struct {
			model.Film
			Actors []model.Actor
		}{{
			Actors: []model.Actor{{
				ActorID:    uint16(0x37),
				FirstName:  "FAY",
				LastName:   "KILMER",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0x60),
				FirstName:  "GENE",
				LastName:   "WILLIS",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0x6e),
				FirstName:  "SUSAN",
				LastName:   "DAVIS",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0x8a),
				FirstName:  "LUCILLE",
				LastName:   "DEE",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}},
			Film: model.Film{
				FilmID: uint16(0x8),
				Rating: gotestparrot.Ptr(model.FilmRating("R")).(*model.FilmRating),
				Title:  "AIRPORT POLLOCK",
			},
		}, {
			Actors: []model.Actor{{
				ActorID:    uint16(0x3),
				FirstName:  "ED",
				LastName:   "CHASE",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0xc),
				FirstName:  "KARL",
				LastName:   "BERRY",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0xd),
				FirstName:  "UMA",
				LastName:   "WOOD",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0x52),
				FirstName:  "WOODY",
				LastName:   "JOLIE",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0x64),
				FirstName:  "SPENCER",
				LastName:   "DEPP",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}, {
				ActorID:    uint16(0xa0),
				FirstName:  "CHRIS",
				LastName:   "DEPP",
				LastUpdate: gotestparrot.Decode("2006-02-15T04:34:33Z", time.Time{}).(time.Time),
			}},
			Film: model.Film{
				FilmID: uint16(0x11),
				Rating: gotestparrot.Ptr(model.FilmRating("R")).(*model.FilmRating),
				Title:  "ALONE TRIP",
			},
		}},
	}})
	gotestparrot.R.Load("TestWindowClause", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT AVG(payment.amount) OVER (),
     AVG(payment.amount) OVER (w1),
     AVG(payment.amount) OVER (w2 ORDER BY payment.customer_id RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING),
     AVG(payment.amount) OVER (w3 RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING)
FROM dvds.payment
WHERE payment.payment_id < ?
WINDOW w1 AS (PARTITION BY payment.payment_date), w2 AS (w1), w3 AS (w2 ORDER BY payment.customer_id)
ORDER BY payment.customer_id;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(10)},
	}})
	gotestparrot.R.Load("TestWindowFunction", []gotestparrot.Recording{{
		Key: 0,
		Value: `
SELECT AVG(payment.amount) OVER (),
     AVG(payment.amount) OVER (PARTITION BY payment.customer_id),
     MAX(payment.amount) OVER (ORDER BY payment.payment_date DESC),
     MIN(payment.amount) OVER (PARTITION BY payment.customer_id ORDER BY payment.payment_date DESC),
     SUM(payment.amount) OVER (PARTITION BY payment.customer_id ORDER BY payment.payment_date DESC ROWS BETWEEN 1 PRECEDING AND 6 FOLLOWING),
     SUM(payment.amount) OVER (PARTITION BY payment.customer_id ORDER BY payment.payment_date DESC RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING),
     MAX(payment.customer_id) OVER (ORDER BY payment.payment_date DESC ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING),
     MIN(payment.customer_id) OVER (PARTITION BY payment.customer_id ORDER BY payment.payment_date DESC),
     SUM(payment.customer_id) OVER (PARTITION BY payment.customer_id ORDER BY payment.payment_date DESC),
     ROW_NUMBER() OVER (ORDER BY payment.payment_date),
     RANK() OVER (ORDER BY payment.payment_date),
     DENSE_RANK() OVER (ORDER BY payment.payment_date),
     CUME_DIST() OVER (ORDER BY payment.payment_date),
     NTILE(11) OVER (ORDER BY payment.payment_date),
     LAG(payment.amount) OVER (ORDER BY payment.payment_date),
     LAG(payment.amount) OVER (ORDER BY payment.payment_date),
     LAG(payment.amount, 2, payment.amount) OVER (ORDER BY payment.payment_date),
     LAG(payment.amount, 2, ?) OVER (ORDER BY payment.payment_date),
     LEAD(payment.amount) OVER (ORDER BY payment.payment_date),
     LEAD(payment.amount) OVER (ORDER BY payment.payment_date),
     LEAD(payment.amount, 2, payment.amount) OVER (ORDER BY payment.payment_date),
     LEAD(payment.amount, 2, ?) OVER (ORDER BY payment.payment_date),
     FIRST_VALUE(payment.amount) OVER (ORDER BY payment.payment_date),
     LAST_VALUE(payment.amount) OVER (ORDER BY payment.payment_date),
     NTH_VALUE(payment.amount, 3) OVER (ORDER BY payment.payment_date)
FROM dvds.payment
WHERE payment.payment_id < ?
GROUP BY payment.amount, payment.customer_id, payment.payment_date;
`,
	}, {
		Key:   1,
		Value: []interface{}{100, 100, int64(10)},
	}})
}

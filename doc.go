/*
Package jet is a framework for writing type-safe SQL queries in Go, with ability to easily convert database query
result into desired arbitrary object structure.


Installation


Use the bellow command to add jet as a dependency into go.mod project:
	$ go get github.com/go-jet/jet/v2

Use the bellow command to add jet as a dependency into GOPATH project:
	$ go get -u github.com/go-jet/jet

Install jet generator to GOPATH bin folder. This will allow generating jet files from the command line.
	cd $GOPATH/src/ && GO111MODULE=off go get -u github.com/go-jet/jet/cmd/jet

Make sure GOPATH bin folder is added to the PATH environment variable.

Usage

Jet requires already defined database schema(with tables, enums etc), so that jet generator can generate SQL Builder
and Model files. File generation is very fast, and can be added as every pre-build step.
Sample command:
	jet -source=PostgreSQL -host=localhost -port=5432 -user=jet -password=pass -dbname=jetdb -schema=dvds -path=./gen

Then next step is to import generated SQL Builder and Model files and write SQL queries in Go:
	import . "some_path/gen/table"
	import "some_path/gen/model"

To write SQL queries for PostgreSQL import:
	. "github.com/go-jet/jet/v2/postgres"

To write SQL queries for MySQL and MariaDB import:
	. "github.com/go-jet/jet/v2/mysql"
*Dot import is used so that Go code resemble as much as native SQL. Dot import is not mandatory.

Write SQL:
	// sub-query
	rRatingFilms := SELECT(
				Film.FilmID,
				Film.Title,
				Film.Rating,
			).
			FROM(Film).
			WHERE(Film.Rating.EQ(enum.FilmRating.R)).
			AsTable("rFilms")

	// export column from sub-query
	rFilmID := Film.FilmID.From(rRatingFilms)

	// main-query
	query := SELECT(
			Actor.AllColumns,
			FilmActor.AllColumns,
			rRatingFilms.AllColumns(),
		).
		FROM(
			rRatingFilms.
			INNER_JOIN(FilmActor, FilmActor.FilmID.EQ(rFilmID)).
			INNER_JOIN(Actor, Actor.ActorID.EQ(FilmActor.ActorID)
		).
		ORDER_BY(rFilmID, Actor.ActorID)

Store result into desired destination:
	var dest []struct {
		model.Film

		Actors []model.Actor
	}

	err := query.Query(db, &dest)

Detail info about all features and use cases can be
found at project wiki page - https://github.com/go-jet/jet/wiki.
*/
package jet

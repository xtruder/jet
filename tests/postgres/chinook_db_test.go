package postgres

import (
	"context"
	"testing"
	"time"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/tests/postgres/gen/chinook/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/chinook/table"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestSelect(t *testing.T) {
	stmt := Album.
		SELECT(Album.AllColumns).
		ORDER_BY(Album.AlbumId.ASC())

	dest := []model.Album{}

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	require.NoError(t, stmt.Query(db, &dest))
	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
	require.EqualValues(t, testparrot.RecordNext(t, dest[len(dest)-1]), dest[len(dest)-1])
}

func TestJoinEverything(t *testing.T) {
	manager := Employee.AS("Manager")

	stmt := Artist.
		LEFT_JOIN(Album, Artist.ArtistId.EQ(Album.ArtistId)).
		LEFT_JOIN(Track, Track.AlbumId.EQ(Album.AlbumId)).
		LEFT_JOIN(Genre, Genre.GenreId.EQ(Track.GenreId)).
		LEFT_JOIN(MediaType, MediaType.MediaTypeId.EQ(Track.MediaTypeId)).
		LEFT_JOIN(PlaylistTrack, PlaylistTrack.TrackId.EQ(Track.TrackId)).
		LEFT_JOIN(Playlist, Playlist.PlaylistId.EQ(PlaylistTrack.PlaylistId)).
		LEFT_JOIN(InvoiceLine, InvoiceLine.TrackId.EQ(Track.TrackId)).
		LEFT_JOIN(Invoice, Invoice.InvoiceId.EQ(InvoiceLine.InvoiceId)).
		LEFT_JOIN(Customer, Customer.CustomerId.EQ(Invoice.CustomerId)).
		LEFT_JOIN(Employee, Employee.EmployeeId.EQ(Customer.SupportRepId)).
		LEFT_JOIN(manager, manager.EmployeeId.EQ(Employee.ReportsTo)).
		SELECT(
			Artist.AllColumns,
			Album.AllColumns,
			Track.AllColumns,
			Genre.AllColumns,
			MediaType.AllColumns,
			PlaylistTrack.AllColumns,
			Playlist.AllColumns,
			Invoice.AllColumns,
			Customer.AllColumns,
			Employee.AllColumns,
			manager.AllColumns,
		).
		ORDER_BY(Artist.ArtistId, Album.AlbumId, Track.TrackId,
			Genre.GenreId, MediaType.MediaTypeId, Playlist.PlaylistId,
			Invoice.InvoiceId, Customer.CustomerId)

	dest := []struct { //list of all artist
		model.Artist

		Albums []struct { // list of albums per artist
			model.Album

			Tracks []struct { // list of tracks per album
				model.Track

				Genre     model.Genre     // track genre
				MediaType model.MediaType // track media type

				Playlists []model.Playlist // list of playlist where track is used

				Invoices []struct { // list of invoices where track occurs
					model.Invoice

					Customer struct { // customer data for invoice
						model.Customer

						Employee *struct { // employee data for customer if exists
							model.Employee

							Manager *model.Employee `alias:"Manager"`
						}
					}
				}
			}
		}
	}{}

	assertStatementRecordSQL(t, stmt)
	assertQueryDest(t, stmt, &dest)

	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	assert.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
	assert.EqualValues(t, testparrot.RecordNext(t, dest[len(dest)-1]), dest[len(dest)-1])
}

func TestSelfJoin(t *testing.T) {
	manager := Employee.AS("Manager")

	stmt := Employee.
		LEFT_JOIN(manager, Employee.ReportsTo.EQ(manager.EmployeeId)).
		SELECT(
			Employee.EmployeeId,
			Employee.FirstName,
			Employee.LastName,
			manager.EmployeeId,
			manager.FirstName,
			manager.LastName,
		).
		ORDER_BY(Employee.EmployeeId)

	dest := []struct {
		model.Employee

		Manager *model.Employee `alias:"Manager.*"`
	}{}

	assertStatementRecordSQL(t, stmt)
	assertQueryDest(t, stmt, &dest)

	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	assert.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])
}

func TestUnionForQuotedNames(t *testing.T) {
	stmt := UNION_ALL(
		Album.SELECT(Album.AllColumns).WHERE(Album.AlbumId.EQ(Int(1))),
		Album.SELECT(Album.AllColumns).WHERE(Album.AlbumId.EQ(Int(2))),
	).ORDER_BY(Album.AlbumId)

	dest := []model.Album{}

	assertStatementRecordSQL(t, stmt)
	assertQueryDest(t, stmt, &dest)

	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])
}

func TestQueryWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	dest := []model.Album{}

	err := Album.
		CROSS_JOIN(Track).
		CROSS_JOIN(InvoiceLine).
		SELECT(Album.AllColumns, Track.AllColumns, InvoiceLine.AllColumns).
		QueryContext(ctx, db, &dest)

	require.Error(t, err, "context deadline exceeded")
}

func TestExecWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := Album.
		CROSS_JOIN(Track).
		CROSS_JOIN(InvoiceLine).
		SELECT(Album.AllColumns, Track.AllColumns, InvoiceLine.AllColumns).
		ExecContext(ctx, db)

	require.Error(t, err, "pq: canceling statement due to user request")
}

func TestSubQueriesForQuotedNames(t *testing.T) {
	first10Artist := Artist.
		SELECT(Artist.AllColumns).
		ORDER_BY(Artist.ArtistId).
		LIMIT(10).
		AsTable("first10Artist")

	artistID := Artist.ArtistId.From(first10Artist)

	first10Albums := Album.
		SELECT(Album.AllColumns).
		ORDER_BY(Album.AlbumId).
		LIMIT(10).
		AsTable("first10Albums")

	albumArtistID := Album.ArtistId.From(first10Albums)

	stmt := first10Artist.
		INNER_JOIN(first10Albums, artistID.EQ(albumArtistID)).
		SELECT(first10Artist.AllColumns(), first10Albums.AllColumns()).
		ORDER_BY(artistID)

	dest := []struct {
		model.Artist

		Album []model.Album
	}{}

	assertStatementRecordSQL(t, stmt)
	assertQueryDest(t, stmt, &dest)

	require.Len(t, dest, testparrot.RecordNext(t, len(dest)).(int))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0:2]), dest[0:2])
}

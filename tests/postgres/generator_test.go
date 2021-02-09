package postgres

import (
	"reflect"
	"testing"

	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/internal/testutils"
	"github.com/stretchr/testify/require"

	"github.com/go-jet/jet/v2/tests/postgres/gen/dvds/model"
)

func TestGeneratedModel(t *testing.T) {
	actor := model.Actor{}

	require.Equal(t, reflect.TypeOf(actor.ActorID).String(), "int32")
	actorIDField, ok := reflect.TypeOf(actor).FieldByName("ActorID")
	require.True(t, ok)
	require.Equal(t, actorIDField.Tag.Get("sql"), "primary_key")
	require.Equal(t, reflect.TypeOf(actor.FirstName).String(), "string")
	require.Equal(t, reflect.TypeOf(actor.LastName).String(), "string")
	require.Equal(t, reflect.TypeOf(actor.LastUpdate).String(), "time.Time")

	filmActor := model.FilmActor{}

	require.Equal(t, reflect.TypeOf(filmActor.FilmID).String(), "int16")
	filmIDField, ok := reflect.TypeOf(filmActor).FieldByName("FilmID")
	require.True(t, ok)
	require.Equal(t, filmIDField.Tag.Get("sql"), "primary_key")

	require.Equal(t, reflect.TypeOf(filmActor.ActorID).String(), "int16")
	actorIDField, ok = reflect.TypeOf(filmActor).FieldByName("ActorID")
	require.True(t, ok)
	require.Equal(t, filmIDField.Tag.Get("sql"), "primary_key")

	staff := model.Staff{}

	require.Equal(t, reflect.TypeOf(staff.Email).String(), "*string")
	require.Equal(t, reflect.TypeOf(staff.Picture).String(), "*[]uint8")
}

// func TestCmdGenerator(t *testing.T) {
// 	goInstallJet := exec.Command("sh", "-c", "cd $GOPATH/src/ && GO111MODULE=off go get github.com/go-jet/jet/cmd/jet")
// 	goInstallJet.Stderr = os.Stderr
// 	err := goInstallJet.Run()
// 	require.NoError(t, err)

// 	err = os.RemoveAll(genTestDir2)
// 	require.NoError(t, err)

// 	cmd := exec.Command("jet", "-source=PostgreSQL", "-dbname=jetdb", "-host=", "-port=5432",
// 		"-user=jet", "-password=jet", "-schema=dvds", "-path="+genTestDir2)
// 	cmd.Stderr = os.Stderr
// 	cmd.Stdout = os.Stdout

// 	err = cmd.Run()
// 	require.NoError(t, err)

// 	assertGeneratedFiles(t)

// 	err = os.RemoveAll(genTestDir2)
// 	require.NoError(t, err)
// }

func TestGenerator(t *testing.T) {
	genDir := t.TempDir()

	require.NoError(t,
		postgres.Generate(db, "dvds", genDir))

	testutils.AssertDirRecordContent(t, genDir)
}

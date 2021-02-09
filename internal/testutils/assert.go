package testutils

import (
	"hash/fnv"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

type ProjectionTest struct {
	Name string
	Test jet.Projection
}

func (test ProjectionTest) Assert(t *testing.T, dialect jet.Dialect) {
	t.Helper()

	out := jet.NewSQLBuilder(dialect, jet.SQLBuilderOptPretty)
	jet.SerializeForProjection(test.Test, jet.SelectStatementType, out)

	// snapshot serialized value
	cupaloy.SnapshotT(t, out.Buff.String(), out.Args)
}

type ProjectionTests []ProjectionTest

func (p ProjectionTests) Run(t *testing.T, dialect jet.Dialect) {
	t.Helper()

	for _, test := range p {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()
			test.Assert(t, dialect)
		})
	}
}

type SerializerTest struct {
	Name   string
	Test   jet.Serializer
	Panics string
}

func (test SerializerTest) Assert(t *testing.T, dialect jet.Dialect) {
	t.Helper()

	out := jet.NewSQLBuilder(dialect, jet.SQLBuilderOptPretty)

	if test.Panics != "" {
		assert.PanicsWithValue(t, test.Panics, func() {
			test.Test.Serialize(jet.SelectStatementType, out)
		})
		return
	}

	test.Test.Serialize(jet.SelectStatementType, out)

	// snapshot serialized value
	cupaloy.SnapshotT(t, out.Buff.String(), out.Args)
}

type SerializerTests []SerializerTest

func (s SerializerTests) Run(t *testing.T, dialect jet.Dialect) {
	t.Helper()

	for _, test := range s {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()
			test.Assert(t, dialect)
		})
	}
}

type StatementTest struct {
	// Name of the test
	Name string

	// Test to test
	Test jet.Statement

	Panics string
}

func (test StatementTest) Assert(t *testing.T) {
	t.Helper()

	if test.Panics != "" {
		assert.PanicsWithValue(t, test.Panics, func() {
			test.Test.Sql()
		})
		return
	}

	query, args := test.Test.Sql(jet.SQLBuilderOptPretty)

	cupaloy.SnapshotT(t, query, args)
}

type StatementTests []StatementTest

func (s StatementTests) Run(t *testing.T) {
	t.Helper()

	for _, test := range s {
		t.Run(test.Name, test.Assert)
	}
}

// AssertExec assert statement execution for successful execution and number of rows affected
func AssertExec(t *testing.T, stmt jet.Statement, db qrm.DB, rowsAffected ...int64) {
	t.Helper()

	res, err := stmt.Exec(db)
	require.NoError(t, err)

	rows, err := res.RowsAffected()
	require.NoError(t, err)

	if len(rowsAffected) > 0 {
		require.EqualValues(t, rowsAffected[0], rows)
	}
}

func AssertQuery(t *testing.T, db qrm.DB, stmt jet.Statement, dest interface{}) {
	t.Helper()

	err := stmt.Query(db, dest)
	require.NoError(t, err)
}

func AssertQueryRecordValues(t *testing.T, db qrm.DB, stmt jet.Statement, dest interface{}) {
	t.Helper()

	err := stmt.Query(db, dest)
	value := reflect.Indirect(reflect.ValueOf(dest)).Interface()
	require.NoError(t, err)
	require.EqualValues(t, testparrot.RecordNext(t, value), value)
}

// AssertExecErr assert statement execution for failed execution with error string errorStr
func AssertExecErr(t *testing.T, stmt jet.Statement, db qrm.DB, errorStr string) {
	t.Helper()

	_, err := stmt.Exec(db)
	require.Error(t, err, errorStr)
}

// AssertStatementSql check if statement Sql() is the same as expectedQuery and expectedArgs
func AssertStatementSql(t *testing.T, query jet.Statement, expectedQuery string, expectedArgs ...interface{}) {
	t.Helper()

	queryStr, args := query.Sql()
	assert.Equal(t, expectedQuery, queryStr)

	if len(expectedArgs) > 0 {
		assert.EqualValues(t, expectedArgs, args)
	}
}

func AssertStatementRecordSQL(t *testing.T, query jet.Statement) {
	t.Helper()

	queryStr, args := query.Sql(jet.SQLBuilderOptPretty)

	snapshotName := strings.Replace(t.Name(), "/", "-", -1)

	h := fnv.New32a()
	h.Write([]byte(queryStr))

	snapshotName += "-" + strconv.Itoa(int(h.Sum32()))

	cupaloy.SnapshotMulti(snapshotName, queryStr, args)
}

// AssertStatementSqlErr checks if statement Sql() panics with errorStr
func AssertStatementSqlErr(t *testing.T, stmt jet.Statement, errorStr string) {
	t.Helper()

	assert.PanicsWithValue(t, errorStr, func() { stmt.Sql() })
}

// AssertDebugStatementSql check if statement Sql() is the same as expectedQuery
func AssertDebugStatementSql(t *testing.T, query jet.Statement, expectedQuery string, expectedArgs ...interface{}) {
	t.Helper()

	_, args := query.Sql()

	if len(expectedArgs) > 0 {
		assert.EqualValues(t, expectedArgs, args, "arguments are not equal")
	}

	debuqSql := query.String()
	assert.Equal(t, debuqSql, expectedQuery)
}

// AssertSerialize checks if clause serialize produces expected query and args
func AssertSerialize(t *testing.T, dialect jet.Dialect, serializer jet.Serializer, query string, args ...interface{}) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect}
	serializer.Serialize(jet.SelectStatementType, &out)

	//fmt.Println(out.Buff.String())

	assert.Equal(t, query, out.Buff.String())

	if len(args) > 0 {
		assert.EqualValues(t, args, out.Args)
	}
}

// AssertSerialize checks if clause serialize produces expected query and args
func AssertRecordSerialize(t *testing.T, dialect jet.Dialect, serializer jet.Serializer) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect}
	serializer.Serialize(jet.SelectStatementType, &out)

	expectedQuery := testparrot.RecordNext(t, out.Buff.String())
	assert.EqualValues(t, out.Buff.String(), expectedQuery)

	if len(out.Args) > 0 {
		expectedArgs := testparrot.RecordNext(t, out.Args)
		assert.EqualValues(t, out.Args, expectedArgs)
	}
}

// AssertRecordClauseSerialize checks if clause serialize produces expected query and args
func AssertRecordClauseSerialize(t *testing.T, dialect jet.Dialect, clause jet.Clause) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect}
	clause.Serialize(jet.SelectStatementType, &out)

	expectedQuery := testparrot.RecordNext(t, out.Buff.String())
	assert.EqualValues(t, out.Buff.String(), expectedQuery)

	if len(out.Args) > 0 {
		expectedArgs := testparrot.RecordNext(t, out.Args)
		assert.EqualValues(t, out.Args, expectedArgs)
	}
}

// AssertClauseSerialize checks if clause serialize produces expected query and args
func AssertClauseSerialize(t *testing.T, dialect jet.Dialect, clause jet.Clause, query string, args ...interface{}) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect}
	clause.Serialize(jet.SelectStatementType, &out)

	assert.Equal(t, out.Buff.String(), query)

	if len(args) > 0 {
		assert.EqualValues(t, out.Args, args)
	}
}

// AssertDebugSerialize checks if clause serialize produces expected debug query and args
func AssertDebugSerialize(t *testing.T, dialect jet.Dialect, clause jet.Serializer, query string, args ...interface{}) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect, Debug: true}
	clause.Serialize(jet.SelectStatementType, &out)

	assert.EqualValues(t, out.Buff.String(), query)

	if len(args) > 0 {
		assert.EqualValues(t, out.Args, args)
	}
}

// AssertSerializeErr check if clause serialize panics with errString
func AssertSerializeErr(t *testing.T, dialect jet.Dialect, clause jet.Serializer, errString string) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect}
	assert.PanicsWithValue(t, errString, func() {
		clause.Serialize(jet.SelectStatementType, &out)
	})
}

// AssertProjectionSerialize check if projection serialize produces expected query and args
func AssertProjectionSerialize(t *testing.T, dialect jet.Dialect, projection jet.Projection, query string, args ...interface{}) {
	t.Helper()

	out := jet.SQLBuilder{Dialect: dialect}
	jet.SerializeForProjection(projection, jet.SelectStatementType, &out)

	assert.EqualValues(t, out.Buff.String(), query)
	assert.EqualValues(t, out.Args, args)
}

// AssertQueryPanicErr check if statement Query execution panics with error errString
func AssertQueryPanicErr(t *testing.T, stmt jet.Statement, db qrm.DB, dest interface{}, errString string) {
	t.Helper()

	assert.PanicsWithValue(t, errString, func() {
		stmt.Query(db, dest)
	})
}

func AssertDirRecordContent(t *testing.T, dir string) {
	fileNames := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			path = strings.TrimPrefix(path, dir)
			fileNames = append(fileNames, path)
			require.EqualValues(t, testparrot.Record(t, path, string(data)), string(data))
		}
		return nil
	})

	require.EqualValues(t, fileNames, testparrot.Record(t, "fileNames", fileNames))

	if err != nil {
		panic(err)
	}
}

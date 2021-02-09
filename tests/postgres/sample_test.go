package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/table"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestUUIDType(t *testing.T) {
	query := AllTypes.
		SELECT(AllTypes.UUID, AllTypes.UUIDPtr).
		WHERE(AllTypes.UUID.EQ(String("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")))

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

	result := model.AllTypes{}

	require.NoError(t, query.Query(db, &result))
	require.Equal(t, result.UUID, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"))
	require.EqualValues(t, result.UUIDPtr,
		testutils.Ptr(uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")))
}

func TestUUIDComplex(t *testing.T) {
	query := Person.INNER_JOIN(PersonPhone, PersonPhone.PersonID.EQ(Person.PersonID)).
		SELECT(Person.AllColumns, PersonPhone.AllColumns).
		ORDER_BY(Person.PersonID.ASC(), PersonPhone.PhoneID.ASC())

	t.Run("slice of structs", func(t *testing.T) {
		dest := []struct {
			model.Person
			Phones []struct {
				model.PersonPhone
			}
		}{}

		require.NoError(t, query.Query(db, &dest))
		require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
	})

	t.Run("single struct", func(t *testing.T) {
		singleQuery := query.WHERE(Person.PersonID.EQ(String("b68dbff6-a87d-11e9-a7f2-98ded00c39c8")))

		require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())

		dest := struct {
			model.Person
			Phones []struct {
				model.PersonPhone
			}
		}{}

		require.NoError(t, singleQuery.Query(db, &dest))
		require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
	})

	t.Run("slice of structs left join", func(t *testing.T) {
		leftQuery := Person.LEFT_JOIN(PersonPhone, PersonPhone.PersonID.EQ(Person.PersonID)).
			SELECT(Person.AllColumns, PersonPhone.AllColumns).
			ORDER_BY(Person.PersonID.ASC(), PersonPhone.PhoneID.ASC())

		dest := []struct {
			model.Person
			Phones []struct {
				model.PersonPhone
			}
		}{}

		require.NoError(t, leftQuery.Query(db, &dest))
		require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
	})
}
func TestEnumType(t *testing.T) {
	query := Person.
		SELECT(Person.AllColumns)

	dest := []model.Person{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestSelecSelfJoin1(t *testing.T) {

	// clean up
	_, err := Employee.DELETE().WHERE(Employee.EmployeeID.GT(Int(100))).Exec(db)
	require.NoError(t, err)

	var expectedSQL = `
SELECT employee.employee_id AS "employee.employee_id",
     employee.first_name AS "employee.first_name",
     employee.last_name AS "employee.last_name",
     employee.employment_date AS "employee.employment_date",
     employee.manager_id AS "employee.manager_id",
     manager.employee_id AS "manager.employee_id",
     manager.first_name AS "manager.first_name",
     manager.last_name AS "manager.last_name",
     manager.employment_date AS "manager.employment_date",
     manager.manager_id AS "manager.manager_id"
FROM test_sample.employee
     LEFT JOIN test_sample.employee AS manager ON (manager.employee_id = employee.manager_id)
ORDER BY employee.employee_id;
`

	manager := Employee.AS("manager")
	query := Employee.
		LEFT_JOIN(manager, manager.EmployeeID.EQ(Employee.ManagerID)).
		SELECT(
			Employee.AllColumns,
			manager.AllColumns,
		).
		ORDER_BY(Employee.EmployeeID)

	testutils.AssertDebugStatementSql(t, query, expectedSQL)

	type Manager model.Employee

	var dest []struct {
		model.Employee

		Manager *Manager
	}

	err = query.Query(db, &dest)

	require.NoError(t, err)
	require.Equal(t, len(dest), 8)
	require.EqualValues(t, dest[0].Employee, model.Employee{
		EmployeeID:     1,
		FirstName:      "Windy",
		LastName:       "Hays",
		EmploymentDate: testutils.TimestampWithTimeZone("1999-01-08 04:05:06.1 +0100 CET", 1),
		ManagerID:      nil,
	})

	require.True(t, dest[0].Manager == nil)

	require.EqualValues(t, dest[7].Employee, model.Employee{
		EmployeeID:     8,
		FirstName:      "Salley",
		LastName:       "Lester",
		EmploymentDate: testutils.TimestampWithTimeZone("1999-01-08 04:05:06 +0100 CET", 1),
		ManagerID:      testutils.Ptr(int32(3)).(*int32),
	})
}

func TestWierdNamesTable(t *testing.T) {
	query := WeirdNamesTable.SELECT(WeirdNamesTable.AllColumns)
	dest := []model.WeirdNamesTable{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

func TestReservedWordEscape(t *testing.T) {
	query := SELECT(User.AllColumns).FROM(User)
	dest := []model.User{}

	require.Equal(t, testparrot.RecordNext(t, query.String()), query.String())
	require.NoError(t, query.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest), dest)
}

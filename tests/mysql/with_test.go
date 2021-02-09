package mysql

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/utils"
	. "github.com/go-jet/jet/v2/mysql"
	. "github.com/go-jet/jet/v2/tests/mysql/gen/dvds/table"
)

func TestWITH_And_SELECT(t *testing.T) {
	salesRep := CTE("sales_rep")
	salesRepStaffID := Staff.StaffID.From(salesRep)
	salesRepFullName := StringColumn("sales_rep_full_name").From(salesRep)
	customerSalesRep := CTE("customer_sales_rep")

	stmt := WITH(
		salesRep.AS(
			SELECT(
				Staff.StaffID,
				Staff.FirstName.CONCAT(Staff.LastName).AS(salesRepFullName.Name()),
			).FROM(Staff),
		),
		customerSalesRep.AS(
			SELECT(
				Customer.FirstName.CONCAT(Customer.LastName).AS("customer_name"),
				salesRepFullName,
			).FROM(
				salesRep.
					INNER_JOIN(Store, Store.ManagerStaffID.EQ(salesRepStaffID)).
					INNER_JOIN(Customer, Customer.StoreID.EQ(Store.StoreID)),
			),
		),
	)(
		SELECT(customerSalesRep.AllColumns()).
			FROM(customerSalesRep).LIMIT(10),
	)

	dest := []struct {
		CustomerName     string
		SalesRepFullName string
	}{}

	assertStatementRecordSQL(t, stmt)
	assertQueryRecordValues(t, stmt, &dest)
}

//func TestWITH_And_INSERT(t *testing.T) {
//	paymentsToInsert := CTE("payments_to_insert")
//
//	stmt := WITH(
//		paymentsToInsert.AS(
//			SELECT(Payment.AllColumns).
//				FROM(Payment).
//				WHERE(Payment.Amount.LT(Float(0.5))),
//		),
//	)(
//		Payment.INSERT(Payment.AllColumns).
//			QUERY(
//				SELECT(paymentsToInsert.AllColumns()).
//					FROM(paymentsToInsert),
//			).ON_DUPLICATE_KEY_UPDATE(
//			Payment.PaymentID.SET(Payment.PaymentID.ADD(Int(100000))),
//		),
//	)
//
//	//fmt.Println(stmt.String())
//
//	tx, err := db.Begin()
//	require.NoError(t, err)
//	defer tx.Rollback()
//
//	testing.AssertExec(t, stmt, tx, 24)
//}

func TestWITH_And_UPDATE(t *testing.T) {
	paymentsToUpdate := CTE("payments_to_update")
	paymentsToDeleteID := Payment.PaymentID.From(paymentsToUpdate)

	stmt := WITH(
		paymentsToUpdate.AS(
			SELECT(Payment.AllColumns).
				FROM(Payment).
				WHERE(Payment.Amount.LT(Float(0.5))),
		),
	)(
		Payment.UPDATE().
			SET(Payment.Amount.SET(Float(0.0))).
			WHERE(Payment.PaymentID.IN(
				SELECT(paymentsToDeleteID).
					FROM(paymentsToUpdate),
			),
			),
	)

	tx, err := db.Begin()
	utils.PanicOnError(err)
	defer tx.Rollback()

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, tx)
}

func TestWITH_And_DELETE(t *testing.T) {
	paymentsToDelete := CTE("payments_to_delete")
	paymentsToDeleteID := Payment.PaymentID.From(paymentsToDelete)

	stmt := WITH(
		paymentsToDelete.AS(
			SELECT(Payment.AllColumns).
				FROM(Payment).
				WHERE(Payment.Amount.LT(Float(0.5))),
		),
	)(
		Payment.DELETE().
			WHERE(Payment.PaymentID.IN(
				SELECT(paymentsToDeleteID).
					FROM(paymentsToDelete),
			),
			),
	)

	tx, err := db.Begin()
	utils.PanicOnError(err)
	defer tx.Rollback()

	assertStatementRecordSQL(t, stmt)
	assertExec(t, stmt, tx)
}

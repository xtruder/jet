package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/tests/postgres/gen/northwind/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/northwind/table"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestNorthwindJoinEverything(t *testing.T) {
	stmt := Customers.
		LEFT_JOIN(CustomerCustomerDemo, Customers.CustomerID.EQ(CustomerCustomerDemo.CustomerID)).
		LEFT_JOIN(CustomerDemographics, CustomerCustomerDemo.CustomerTypeID.EQ(CustomerDemographics.CustomerTypeID)).
		LEFT_JOIN(Orders, Orders.CustomerID.EQ(Customers.CustomerID)).
		LEFT_JOIN(Shippers, Orders.ShipVia.EQ(Shippers.ShipperID)).
		LEFT_JOIN(OrderDetails, Orders.OrderID.EQ(OrderDetails.OrderID)).
		LEFT_JOIN(Products, OrderDetails.ProductID.EQ(Products.ProductID)).
		LEFT_JOIN(Categories, Products.CategoryID.EQ(Categories.CategoryID)).
		LEFT_JOIN(Suppliers, Products.SupplierID.EQ(Suppliers.SupplierID)).
		LEFT_JOIN(Employees, Orders.EmployeeID.EQ(Employees.EmployeeID)).
		LEFT_JOIN(EmployeeTerritories, EmployeeTerritories.EmployeeID.EQ(Employees.EmployeeID)).
		LEFT_JOIN(Territories, EmployeeTerritories.TerritoryID.EQ(Territories.TerritoryID)).
		LEFT_JOIN(Region, Territories.RegionID.EQ(Region.RegionID)).
		SELECT(
			Customers.AllColumns,
			CustomerDemographics.AllColumns,
			Orders.AllColumns,
			Shippers.AllColumns,
			OrderDetails.AllColumns,
			Products.AllColumns,
			Categories.AllColumns,
			Suppliers.AllColumns,
		).
		ORDER_BY(Customers.CustomerID, Orders.OrderID, Products.ProductID)

	dest := []struct {
		model.Customers

		Demographics model.CustomerDemographics

		Orders []struct {
			model.Orders

			Shipper model.Shippers

			Details struct {
				model.OrderDetails

				Products []struct {
					model.Products

					Category model.Categories
					Supplier model.Suppliers
				}
			}
		}
	}{}

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())
	require.NoError(t, stmt.Query(db, &dest))
	require.EqualValues(t, testparrot.RecordNext(t, dest[0]), dest[0])
	require.EqualValues(t, testparrot.RecordNext(t, dest[len(dest)-1]), dest[len(dest)-1])
}

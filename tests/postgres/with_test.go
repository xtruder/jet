package postgres

import (
	"testing"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/tests/postgres/gen/northwind/model"
	. "github.com/go-jet/jet/v2/tests/postgres/gen/northwind/table"
	"github.com/stretchr/testify/require"
	"github.com/xtruder/go-testparrot"
)

func TestWithRegionalSales(t *testing.T) {
	regionalSales := CTE("regional_sales")
	topRegion := CTE("top_region")

	regionalSalesTotalSales := IntegerColumn("total_sales").From(regionalSales)
	regionalSalesShipRegion := Orders.ShipRegion.From(regionalSales)
	topRegionShipRegion := regionalSalesShipRegion.From(topRegion)

	stmt := WITH(
		regionalSales.AS(
			SELECT(
				Orders.ShipRegion,
				SUM(OrderDetails.Quantity).AS(regionalSalesTotalSales.Name()),
			).
				FROM(Orders.INNER_JOIN(OrderDetails, OrderDetails.OrderID.EQ(Orders.OrderID))).
				GROUP_BY(Orders.ShipRegion),
		),
		topRegion.AS(
			SELECT(regionalSalesShipRegion).
				FROM(regionalSales).
				WHERE(regionalSalesTotalSales.GT(
					IntExp(
						SELECT(SUM(regionalSalesTotalSales)).
							FROM(regionalSales),
					).DIV(Int(50)),
				)),
		),
	)(
		SELECT(
			Orders.ShipRegion,
			OrderDetails.ProductID,
			COUNT(STAR).AS("product_units"),
			SUM(OrderDetails.Quantity).AS("product_sales"),
		).
			FROM(Orders.INNER_JOIN(OrderDetails, Orders.OrderID.EQ(OrderDetails.OrderID))).
			WHERE(Orders.ShipRegion.IN(
				topRegion.SELECT(topRegionShipRegion)),
			).
			GROUP_BY(Orders.ShipRegion, OrderDetails.ProductID).
			ORDER_BY(SUM(OrderDetails.Quantity).DESC()),
	)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	_, err := stmt.Exec(db)
	require.NoError(t, err)
}

func TestWithStatementDeleteAndInsert(t *testing.T) {
	removeDiscontinuedOrders := CTE("remove_discontinued_orders")
	updateDiscontinuedPrice := CTE("update_discontinued_price")
	logDiscontinuedProducts := CTE("log_discontinued")

	discontinuedProductID := OrderDetails.ProductID.From(removeDiscontinuedOrders)

	stmt := WITH(
		removeDiscontinuedOrders.AS(
			OrderDetails.DELETE().
				WHERE(OrderDetails.ProductID.IN(
					SELECT(Products.ProductID).
						FROM(Products).
						WHERE(Products.Discontinued.EQ(Int(1)))),
				).RETURNING(OrderDetails.ProductID),
		),
		updateDiscontinuedPrice.AS(
			Products.UPDATE().
				SET(
					Products.UnitPrice.SET(Float(0.0)),
				).
				WHERE(Products.ProductID.IN(removeDiscontinuedOrders.SELECT(discontinuedProductID))).
				RETURNING(Products.AllColumns),
		),
		logDiscontinuedProducts.AS(
			ProductLogs.INSERT(ProductLogs.AllColumns).
				QUERY(SELECT(updateDiscontinuedPrice.AllColumns()).FROM(updateDiscontinuedPrice)).
				RETURNING(
					ProductLogs.ProductID,
					ProductLogs.ProductName,
					ProductLogs.SupplierID,
					ProductLogs.CategoryID,
					ProductLogs.QuantityPerUnit,
					ProductLogs.UnitPrice,
					ProductLogs.UnitsInStock,
					ProductLogs.UnitsOnOrder,
					ProductLogs.ReorderLevel,
					ProductLogs.Discontinued,
				),
		),
	)(
		SELECT(logDiscontinuedProducts.AllColumns()).
			FROM(logDiscontinuedProducts),
	)

	require.Len(t, removeDiscontinuedOrders.AllColumns(), 1)
	require.Len(t, updateDiscontinuedPrice.AllColumns()[0].(ProjectionList), 10)
	require.Len(t, logDiscontinuedProducts.AllColumns(), 10)

	require.Equal(t, testparrot.RecordNext(t, stmt.String()), stmt.String())

	resp := []model.ProductLogs{}

	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback()

	require.NoError(t, stmt.Query(tx, &resp))
	require.EqualValues(t, testparrot.RecordNext(t, resp), resp)
}

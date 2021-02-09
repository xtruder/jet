// Code generated by testparrot. DO NOT EDIT.

package postgres

import (
	model "github.com/go-jet/jet/v2/tests/postgres/gen/northwind/model"
	gotestparrot "github.com/xtruder/go-testparrot"
)

func init() {
	gotestparrot.R.Load("TestWithRegionalSales", []gotestparrot.Recording{{
		Key: 0,
		Value: `
WITH regional_sales AS (
     SELECT orders.ship_region AS "orders.ship_region",
          SUM(order_details.quantity) AS "total_sales"
     FROM northwind.orders
          INNER JOIN northwind.order_details ON (order_details.order_id = orders.order_id)
     GROUP BY orders.ship_region
),top_region AS (
     SELECT regional_sales."orders.ship_region" AS "orders.ship_region"
     FROM regional_sales
     WHERE regional_sales.total_sales > ((
               SELECT SUM(regional_sales.total_sales)
               FROM regional_sales
          ) / 50)
)
SELECT orders.ship_region AS "orders.ship_region",
     order_details.product_id AS "order_details.product_id",
     COUNT(*) AS "product_units",
     SUM(order_details.quantity) AS "product_sales"
FROM northwind.orders
     INNER JOIN northwind.order_details ON (orders.order_id = order_details.order_id)
WHERE orders.ship_region IN ((
          SELECT top_region."orders.ship_region" AS "orders.ship_region"
          FROM top_region
     ))
GROUP BY orders.ship_region, order_details.product_id
ORDER BY SUM(order_details.quantity) DESC;
`,
	}})
	gotestparrot.R.Load("TestWithStatementDeleteAndInsert", []gotestparrot.Recording{{
		Key: 0,
		Value: `
WITH remove_discontinued_orders AS (
     DELETE FROM northwind.order_details
     WHERE order_details.product_id IN ((
               SELECT products.product_id AS "products.product_id"
               FROM northwind.products
               WHERE products.discontinued = 1
          ))
     RETURNING order_details.product_id AS "order_details.product_id"
),update_discontinued_price AS (
     UPDATE northwind.products
     SET unit_price = 0
     WHERE products.product_id IN ((
               SELECT remove_discontinued_orders."order_details.product_id" AS "order_details.product_id"
               FROM remove_discontinued_orders
          ))
     RETURNING products.product_id AS "products.product_id",
               products.product_name AS "products.product_name",
               products.supplier_id AS "products.supplier_id",
               products.category_id AS "products.category_id",
               products.quantity_per_unit AS "products.quantity_per_unit",
               products.unit_price AS "products.unit_price",
               products.units_in_stock AS "products.units_in_stock",
               products.units_on_order AS "products.units_on_order",
               products.reorder_level AS "products.reorder_level",
               products.discontinued AS "products.discontinued"
),log_discontinued AS (
     INSERT INTO northwind.product_logs (product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued) (
          SELECT update_discontinued_price."products.product_id" AS "products.product_id",
               update_discontinued_price."products.product_name" AS "products.product_name",
               update_discontinued_price."products.supplier_id" AS "products.supplier_id",
               update_discontinued_price."products.category_id" AS "products.category_id",
               update_discontinued_price."products.quantity_per_unit" AS "products.quantity_per_unit",
               update_discontinued_price."products.unit_price" AS "products.unit_price",
               update_discontinued_price."products.units_in_stock" AS "products.units_in_stock",
               update_discontinued_price."products.units_on_order" AS "products.units_on_order",
               update_discontinued_price."products.reorder_level" AS "products.reorder_level",
               update_discontinued_price."products.discontinued" AS "products.discontinued"
          FROM update_discontinued_price
     )
     RETURNING product_logs.product_id AS "product_logs.product_id",
               product_logs.product_name AS "product_logs.product_name",
               product_logs.supplier_id AS "product_logs.supplier_id",
               product_logs.category_id AS "product_logs.category_id",
               product_logs.quantity_per_unit AS "product_logs.quantity_per_unit",
               product_logs.unit_price AS "product_logs.unit_price",
               product_logs.units_in_stock AS "product_logs.units_in_stock",
               product_logs.units_on_order AS "product_logs.units_on_order",
               product_logs.reorder_level AS "product_logs.reorder_level",
               product_logs.discontinued AS "product_logs.discontinued"
)
SELECT log_discontinued."product_logs.product_id" AS "product_logs.product_id",
     log_discontinued."product_logs.product_name" AS "product_logs.product_name",
     log_discontinued."product_logs.supplier_id" AS "product_logs.supplier_id",
     log_discontinued."product_logs.category_id" AS "product_logs.category_id",
     log_discontinued."product_logs.quantity_per_unit" AS "product_logs.quantity_per_unit",
     log_discontinued."product_logs.unit_price" AS "product_logs.unit_price",
     log_discontinued."product_logs.units_in_stock" AS "product_logs.units_in_stock",
     log_discontinued."product_logs.units_on_order" AS "product_logs.units_on_order",
     log_discontinued."product_logs.reorder_level" AS "product_logs.reorder_level",
     log_discontinued."product_logs.discontinued" AS "product_logs.discontinued"
FROM log_discontinued;
`,
	}, {
		Key: 1,
		Value: []model.ProductLogs{{
			CategoryID:      gotestparrot.Ptr(int16(1)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(1),
			ProductName:     "Chai",
			QuantityPerUnit: gotestparrot.Ptr("10 boxes x 30 bags").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(10)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(8)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(39)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(1)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(2),
			ProductName:     "Chang",
			QuantityPerUnit: gotestparrot.Ptr("24 - 12 oz bottles").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(25)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(1)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(17)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(40)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(2)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(5),
			ProductName:     "Chef Anton's Gumbo Mix",
			QuantityPerUnit: gotestparrot.Ptr("36 boxes").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(2)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(0)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(6)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(9),
			ProductName:     "Mishi Kobe Niku",
			QuantityPerUnit: gotestparrot.Ptr("18 - 500 g pkgs.").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(4)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(29)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(6)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(17),
			ProductName:     "Alice Mutton",
			QuantityPerUnit: gotestparrot.Ptr("20 - 1 kg tins").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(7)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(0)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(1)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(24),
			ProductName:     "Guaraná Fantástica",
			QuantityPerUnit: gotestparrot.Ptr("12 - 355 ml cans").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(10)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(20)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(7)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(28),
			ProductName:     "Rössle Sauerkraut",
			QuantityPerUnit: gotestparrot.Ptr("25 - 825 g cans").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(12)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(26)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(6)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(29),
			ProductName:     "Thüringer Rostbratwurst",
			QuantityPerUnit: gotestparrot.Ptr("50 bags x 30 sausgs.").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(12)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(0)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(5)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(42),
			ProductName:     "Singaporean Hokkien Fried Mee",
			QuantityPerUnit: gotestparrot.Ptr("32 - 1 kg pkgs.").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(20)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(26)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}, {
			CategoryID:      gotestparrot.Ptr(int16(6)).(*int16),
			Discontinued:    int32(1),
			ProductID:       int16(53),
			ProductName:     "Perth Pasties",
			QuantityPerUnit: gotestparrot.Ptr("48 pieces").(*string),
			ReorderLevel:    gotestparrot.Ptr(int16(0)).(*int16),
			SupplierID:      gotestparrot.Ptr(int16(24)).(*int16),
			UnitPrice:       gotestparrot.Ptr(float32(0)).(*float32),
			UnitsInStock:    gotestparrot.Ptr(int16(0)).(*int16),
			UnitsOnOrder:    gotestparrot.Ptr(int16(0)).(*int16),
		}},
	}})
}

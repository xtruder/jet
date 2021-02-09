// Code generated by testparrot. DO NOT EDIT.

package postgres

import gotestparrot "github.com/xtruder/go-testparrot"

func init() {
	gotestparrot.R.Load("TestWithStatement", []gotestparrot.Recording{{
		Key: 0,
		Value: `
WITH cte1 AS (
     SELECT table1.col_int AS "table1.col_int"
     FROM db.table1
),cte2 AS (
     UPDATE db.table2
     SET col_int = $1
     WHERE table2.col3 IN ((
               SELECT cte1."table1.col_int" AS "table1.col_int"
               FROM cte1
          ))
)
SELECT cte1."table1.col_int" AS "table1.col_int"
FROM cte1;
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(1)},
	}})
}

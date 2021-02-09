// Code generated by testparrot. DO NOT EDIT.

package postgres

import gotestparrot "github.com/xtruder/go-testparrot"

func init() {
	gotestparrot.R.Load("TestUpdateOneColumnWithSelect", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE db.table1
SET col_float = (
     SELECT table1.col_float AS "table1.col_float"
     FROM db.table1
)
WHERE table1.col1 = $1
RETURNING table1.col1 AS "table1.col1";
`,
	}, {
		Key:   1,
		Value: []interface{}{int64(2)},
	}})
	gotestparrot.R.Load("TestUpdateWithOneValue", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE db.table1
SET col_int = $1
WHERE table1.col_int >= $2;
`,
	}, {
		Key:   1,
		Value: []interface{}{1, int64(33)},
	}})
	gotestparrot.R.Load("TestUpdateWithValues", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE db.table1
SET (col_int, col_float) = ($1, $2)
WHERE table1.col_int >= $3;
`,
	}, {
		Key:   1,
		Value: []interface{}{1, 22.2, int64(33)},
	}})
}

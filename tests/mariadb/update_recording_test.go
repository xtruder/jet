// Code generated by testparrot. DO NOT EDIT.

package mariadb

import gotestparrot "github.com/xtruder/go-testparrot"

func init() {
	gotestparrot.R.Load("TestUpdateValues/new_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET name = ?,
    url = ?
WHERE link.name = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{"Bong", "http://bong.com", "Bing"},
	}})
	gotestparrot.R.Load("TestUpdateValues/old_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET name = ?,
    url = ?
WHERE link.name = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{"Bong", "http://bong.com", "Bing"},
	}})
	gotestparrot.R.Load("TestUpdateWithModelData", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET id = ?,
    url = ?,
    name = ?,
    description = ?
WHERE link.id = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{int32(201), "http://www.duckduckgo.com", "DuckDuckGo", nil, int64(201)},
	}})
	gotestparrot.R.Load("TestUpdateWithModelDataAndMutableColumns", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET url = ?,
    name = ?,
    description = ?
WHERE link.id = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{"http://www.duckduckgo.com", "DuckDuckGo", nil, int64(201)},
	}})
	gotestparrot.R.Load("TestUpdateWithModelDataAndPredefinedColumnList", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET description = ?,
    name = ?,
    url = ?
WHERE link.id = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{nil, "DuckDuckGo", "http://www.duckduckgo.com", int64(201)},
	}})
	gotestparrot.R.Load("TestUpdateWithSubQueries/new_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET name = (
         SELECT ?
    ),
    url = (
         SELECT link2.url AS "link2.url"
         FROM test_sample.link2
         WHERE link2.name = ?
    )
WHERE link.name = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{"Bong", "Youtube", "Bing"},
	}})
	gotestparrot.R.Load("TestUpdateWithSubQueries/old_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET name = (
         SELECT ?
    ),
    url = (
         SELECT link2.url AS "link2.url"
         FROM test_sample.link2
         WHERE link2.name = ?
    )
WHERE link.name = ?;
`,
	}, {
		Key:   1,
		Value: []interface{}{"Bong", "Youtube", "Bing"},
	}})
}

// Code generated by testparrot. DO NOT EDIT.

package postgres

import (
	model "github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/model"
	gotestparrot "github.com/xtruder/go-testparrot"
)

func init() {
	gotestparrot.R.Load("TestUpdateAndReturning", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (name, url) = ('DuckDuckGo', 'http://www.duckduckgo.com')
WHERE link.name = 'Ask'
RETURNING link.id AS "link.id",
          link.url AS "link.url",
          link.name AS "link.name",
          link.description AS "link.description";
`,
	}, {
		Key: 1,
		Value: []model.Link{{
			ID:   int32(201),
			Name: "DuckDuckGo",
			URL:  "http://www.duckduckgo.com",
		}, {
			ID:   int32(202),
			Name: "DuckDuckGo",
			URL:  "http://www.duckduckgo.com",
		}},
	}})
	gotestparrot.R.Load("TestUpdateValues/deprecated_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (name, url) = ('Bong', 'http://bong.com')
WHERE link.name = 'Bing';
`,
	}, {
		Key: 1,
		Value: []model.Link{{
			ID:   int32(204),
			Name: "Bong",
			URL:  "http://bong.com",
		}},
	}})
	gotestparrot.R.Load("TestUpdateValues/new_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET name = 'DuckDuckGo',
    url = 'www.duckduckgo.com'
WHERE link.name = 'Yahoo';
`,
	}})
	gotestparrot.R.Load("TestUpdateWithInvalidSelect/deprecated_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (id, url, name, description) = (
     SELECT link.id AS "link.id",
          link.name AS "link.name"
     FROM test_sample.link
     WHERE link.id = 0
)
WHERE link.id = 0;
`,
	}})
	gotestparrot.R.Load("TestUpdateWithInvalidSelect/new_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (id, url, name, description) = (
         SELECT link.url AS "link.url",
              link.name AS "link.name",
              link.description AS "link.description"
         FROM test_sample.link
    )
WHERE link.id = 0;
`,
	}})
	gotestparrot.R.Load("TestUpdateWithModelData", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (id, url, name, description) = (201, 'http://www.duckduckgo.com', 'DuckDuckGo', NULL)
WHERE link.id = 201;
`,
	}})
	gotestparrot.R.Load("TestUpdateWithModelDataAndPredefinedColumnList", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (description, name, url) = (NULL, 'DuckDuckGo', 'http://www.duckduckgo.com')
WHERE link.id = 201;
`,
	}})
	gotestparrot.R.Load("TestUpdateWithSelect/deprecated_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (id, url, name, description) = (
     SELECT link.id AS "link.id",
          link.url AS "link.url",
          link.name AS "link.name",
          link.description AS "link.description"
     FROM test_sample.link
     WHERE link.id = 0
)
WHERE link.id = 0;
`,
	}})
	gotestparrot.R.Load("TestUpdateWithSelect/new_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (url, name, description) = (
         SELECT link.url AS "link.url",
              link.name AS "link.name",
              link.description AS "link.description"
         FROM test_sample.link
         WHERE link.id = 0
    )
WHERE link.id = 0;
`,
	}})
	gotestparrot.R.Load("TestUpdateWithSubQueries/deprecated_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET (name, url) = ((
     SELECT 'Bong'
), (
     SELECT link.url AS "link.url"
     FROM test_sample.link
     WHERE link.name = 'Bing'
))
WHERE link.name = 'Bing';
`,
	}})
	gotestparrot.R.Load("TestUpdateWithSubQueries/new_version", []gotestparrot.Recording{{
		Key: 0,
		Value: `
UPDATE test_sample.link
SET name = 'Bong',
    url = (
         SELECT link.url AS "link.url"
         FROM test_sample.link
         WHERE link.name = 'Bing'
    )
WHERE link.name = 'Bing';
`,
	}})
}

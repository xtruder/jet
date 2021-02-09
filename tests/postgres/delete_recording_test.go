// Code generated by testparrot. DO NOT EDIT.

package postgres

import (
	model "github.com/go-jet/jet/v2/tests/postgres/gen/test_sample/model"
	gotestparrot "github.com/xtruder/go-testparrot"
)

func init() {
	gotestparrot.R.Load("TestDeleteExecContext", []gotestparrot.Recording{{
		Key: 0,
		Value: `
DELETE FROM test_sample.link
WHERE link.name IN ('Gmail', 'Outlook');
`,
	}})
	gotestparrot.R.Load("TestDeleteQueryContext", []gotestparrot.Recording{{
		Key: 0,
		Value: `
DELETE FROM test_sample.link
WHERE link.name IN ('Gmail', 'Outlook');
`,
	}})
	gotestparrot.R.Load("TestDeleteWithWhere", []gotestparrot.Recording{{
		Key: 0,
		Value: `
DELETE FROM test_sample.link
WHERE link.name IN ('Gmail', 'Outlook');
`,
	}})
	gotestparrot.R.Load("TestDeleteWithWhereAndReturning", []gotestparrot.Recording{{
		Key: 0,
		Value: `
DELETE FROM test_sample.link
WHERE link.name IN ('Gmail', 'Outlook')
RETURNING link.id AS "link.id",
          link.url AS "link.url",
          link.name AS "link.name",
          link.description AS "link.description";
`,
	}, {
		Key: 1,
		Value: []model.Link{{
			Description: gotestparrot.Ptr("Email service developed by Google").(*string),
			Name:        "Gmail",
			URL:         "www.gmail.com",
		}, {
			Description: gotestparrot.Ptr("Email service developed by Microsoft").(*string),
			Name:        "Outlook",
			URL:         "www.outlook.live.com",
		}},
	}})
}
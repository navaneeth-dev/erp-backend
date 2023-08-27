package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ezt92lv1megwhoh")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[
			"CREATE UNIQUE INDEX ` + "`" + `idx_eISOhNe` + "`" + ` ON ` + "`" + `hackathons` + "`" + ` (` + "`" + `original_url` + "`" + `)"
		]`), &collection.Indexes)

		// add
		new_original_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bsn8yerh",
			"name": "original_url",
			"type": "url",
			"required": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), new_original_url)
		collection.Schema.AddField(new_original_url)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ezt92lv1megwhoh")
		if err != nil {
			return err
		}

		json.Unmarshal([]byte(`[]`), &collection.Indexes)

		// remove
		collection.Schema.RemoveField("bsn8yerh")

		return dao.SaveCollection(collection)
	})
}

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

		// add
		new_source := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "hctspwow",
			"name": "source",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "2wm92beq199389v",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"name"
				]
			}
		}`), new_source)
		collection.Schema.AddField(new_source)

		// update
		edit_original_url := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bsn8yerh",
			"name": "original_url",
			"type": "url",
			"required": true,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), edit_original_url)
		collection.Schema.AddField(edit_original_url)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ezt92lv1megwhoh")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("hctspwow")

		// update
		edit_original_url := &schema.SchemaField{}
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
		}`), edit_original_url)
		collection.Schema.AddField(edit_original_url)

		return dao.SaveCollection(collection)
	})
}

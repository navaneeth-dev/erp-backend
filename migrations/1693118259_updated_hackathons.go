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

		// remove
		collection.Schema.RemoveField("0ponilzp")

		// add
		new_location := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "cnnwupqa",
			"name": "location",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 1,
				"max": 64,
				"pattern": ""
			}
		}`), new_location)
		collection.Schema.AddField(new_location)

		// add
		new_end := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "mdaqqjvs",
			"name": "end",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_end)
		collection.Schema.AddField(new_end)

		// update
		edit_start := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jhlhowkc",
			"name": "start",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_start)
		collection.Schema.AddField(edit_start)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ezt92lv1megwhoh")
		if err != nil {
			return err
		}

		// add
		del_location := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "0ponilzp",
			"name": "location",
			"type": "select",
			"required": true,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"online",
					"offline"
				]
			}
		}`), del_location)
		collection.Schema.AddField(del_location)

		// remove
		collection.Schema.RemoveField("cnnwupqa")

		// remove
		collection.Schema.RemoveField("mdaqqjvs")

		// update
		edit_start := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jhlhowkc",
			"name": "start",
			"type": "date",
			"required": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_start)
		collection.Schema.AddField(edit_start)

		return dao.SaveCollection(collection)
	})
}

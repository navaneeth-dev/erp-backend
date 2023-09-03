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

		collection, err := dao.FindCollectionByNameOrId("8nq055ts67h4xie")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("jsfpou5q")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8nq055ts67h4xie")
		if err != nil {
			return err
		}

		// add
		del_role := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jsfpou5q",
			"name": "role",
			"type": "select",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"student",
					"teacher"
				]
			}
		}`), del_role)
		collection.Schema.AddField(del_role)

		return dao.SaveCollection(collection)
	})
}

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

		collection, err := dao.FindCollectionByNameOrId("a8tcuksh1s4tfal")
		if err != nil {
			return err
		}

		// add
		new_users := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "deh7qdvo",
			"name": "users",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": []
			}
		}`), new_users)
		collection.Schema.AddField(new_users)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("a8tcuksh1s4tfal")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("deh7qdvo")

		return dao.SaveCollection(collection)
	})
}

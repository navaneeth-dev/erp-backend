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

		// add
		new_avatar := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ag7o4kbq",
			"name": "avatar",
			"type": "url",
			"required": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), new_avatar)
		collection.Schema.AddField(new_avatar)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8nq055ts67h4xie")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ag7o4kbq")

		return dao.SaveCollection(collection)
	})
}

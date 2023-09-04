package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "y3whxc3hhex2q1f",
			"created": "2023-09-04 05:37:32.463Z",
			"updated": "2023-09-04 05:37:32.463Z",
			"name": "attendance",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "kvnc5h0j",
					"name": "date",
					"type": "date",
					"required": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "z30vm73p",
					"name": "class",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "a8tcuksh1s4tfal",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "g06tjrgb",
					"name": "absent_users",
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
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("y3whxc3hhex2q1f")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

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
			"id": "a8tcuksh1s4tfal",
			"created": "2023-09-04 05:32:33.396Z",
			"updated": "2023-09-04 05:32:33.396Z",
			"name": "classes",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ivyvxm1c",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 1,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "xujsrj0z",
					"name": "teacher",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "8nq055ts67h4xie",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
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

		collection, err := dao.FindCollectionByNameOrId("a8tcuksh1s4tfal")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

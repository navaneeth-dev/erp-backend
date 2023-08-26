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
			"id": "ezt92lv1megwhoh",
			"created": "2023-08-26 14:10:34.484Z",
			"updated": "2023-08-26 14:10:34.484Z",
			"name": "hackathons",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "rwarseiw",
					"name": "title",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 3,
						"max": 64,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "6tgv5ilo",
					"name": "description",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 1,
						"max": 256,
						"pattern": ""
					}
				},
				{
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
				},
				{
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

		collection, err := dao.FindCollectionByNameOrId("ezt92lv1megwhoh")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

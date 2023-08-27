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
			"id": "2wm92beq199389v",
			"created": "2023-08-27 06:42:18.015Z",
			"updated": "2023-08-27 06:42:18.015Z",
			"name": "sources",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "csn0p9yi",
					"name": "url",
					"type": "url",
					"required": true,
					"unique": false,
					"options": {
						"exceptDomains": [],
						"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "bsvbazh1",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": 1,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_vrTin2a` + "`" + ` ON ` + "`" + `sources` + "`" + ` (` + "`" + `url` + "`" + `)"
			],
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

		collection, err := dao.FindCollectionByNameOrId("2wm92beq199389v")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

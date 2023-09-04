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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("w9rtbry146mk38y")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	}, func(db dbx.Builder) error {
		jsonData := `{
			"id": "w9rtbry146mk38y",
			"created": "2023-09-03 11:55:26.551Z",
			"updated": "2023-09-03 11:55:26.551Z",
			"name": "test",
			"type": "auth",
			"system": false,
			"schema": [],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": "",
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"allowEmailAuth": false,
				"allowOAuth2Auth": true,
				"allowUsernameAuth": false,
				"exceptEmailDomains": [],
				"manageRule": null,
				"minPasswordLength": 8,
				"onlyEmailDomains": [],
				"requireEmail": false
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	})
}

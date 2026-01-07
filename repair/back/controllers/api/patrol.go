package api

import (
	"encoding/json"
	"repair/controllers"
	"repair/models"
)

type PatrolController struct {
	controllers.Controller
}

// @POST()
func (c *PatrolController) Upload() {
	conn := c.NewConnection()

	manager := models.NewPatrolManager(conn)
	patrolimageManager := models.NewPatrolimageManager(conn)

	files := c.MultiUpload("patrol", "file")

	item := c.Get("item")

	var patrol models.Patrol
	json.Unmarshal([]byte(item), &patrol)

	manager.Insert(&patrol)

	patrol.Id = manager.GetIdentity()

	for _, file := range files {
		image := &models.Patrolimage{}

		image.Filename = file
		image.Patrol = patrol.Id
		image.Apt = patrol.Apt

		patrolimageManager.Insert(image)
	}

	c.Set("id", patrol.Id)
}

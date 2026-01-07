package api

import (
	"anb/controllers"
	"anb/models"
	"log"
)

type AptgroupController struct {
	controllers.Controller
}

func (c *AptgroupController) AjaxGet() {
	conn := c.NewConnection()

	id := c.Geti64("id")

	log.Println(id)

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(id)

	manager := models.NewAptgroupManager(conn)
	item := manager.Get(apt.Aptgroup)

	c.Set("item", item)
}

package api

import (
	"anb/controllers"
	"anb/models"
)

type StatusController struct {
	controllers.Controller
}

func (c *StatusController) AjaxList() {
	conn := c.NewConnection()

	company := c.Geti64("company")

	manager := models.NewStatusManager(conn)
	items := manager.GetListByCompany(company, 0, 0, "order, s_id")
	c.Set("items", items)
}

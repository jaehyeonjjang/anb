package api

import (
	"repair/controllers"
	"repair/models"
)

type DetailController struct {
	controllers.Controller
}

func (c *DetailController) Post_Insert(item *models.Detail) {
	conn := c.NewConnection()

	aptdetailManager := models.NewAptdetailManager(conn)

	apt := aptdetailManager.Get(item.Apt)

	if apt == nil || apt.Id == 0 {
		var aptdetail models.Aptdetail
		aptdetail.Id = item.Apt
		aptdetail.Record4 = "1000-01-01"
		aptdetail.Record5 = "1000-01-01"
		aptdetailManager.Insert(&aptdetail)
	}
}

// @Post()
func (c *DetailController) Duplication(id int64) {
	conn := c.NewConnection()

	detailManager := models.NewDetailManager(conn)
	detailtechnicianManager := models.NewDetailtechnicianManager(conn)

	detail := detailManager.Get(id)

	detail.Id = 0
	detail.Date = ""
	detailManager.Insert(detail)

	detail.Id = detailManager.GetIdentity()

	detailtechnicians := detailtechnicianManager.Find([]interface{}{models.Where{Column: "detail", Value: id, Compare: "="}, models.Ordering("dt_order,dt_id")})

	for _, item := range detailtechnicians {
		item.Id = 0
		item.Detail = detail.Id

		detailtechnicianManager.Insert(&item)
	}
}

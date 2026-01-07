package api

import (
	"repair/controllers"
	"repair/models"
)

type DongController struct {
	controllers.Controller
}

func (c *DongController) Post_Delete(item *models.Dong) {
	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)
	breakdownManager.DeleteWhere([]any{
		models.Where{Column: "dong", Value: item.Id, Compare: "="},
	})

}

func (c *DongController) Post_Deletebatch(item *[]models.Dong) {
	if len(*item) == 0 {
		return
	}

	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)

	for _, v := range *item {
		breakdownManager.DeleteWhere([]any{
			models.Where{Column: "dong", Value: v.Id, Compare: "="},
		})
	}
}

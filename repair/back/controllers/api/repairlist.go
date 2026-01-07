package api

import (
	"fmt"
	"repair/controllers"
	"repair/models"
	"repair/models/repairlist"
	"strings"
)

type RepairlistController struct {
	controllers.Controller
}

func (c *RepairlistController) Search() {
	page := c.Geti("page")
	pagesize := c.Geti("pagesize")
	search := c.Get("search")
	typeid := c.Geti("type")
	status := c.Geti("status")

	conn := c.NewConnection()

	manager := models.NewRepairlistManager(conn)

	var args []interface{}

	if search != "" {
		strs := strings.Split(search, " ")

		query := ""
		for i, v := range strs {
			if i > 0 {
				query += " and "
			}

			if i == 0 {
				query += fmt.Sprintf("a_name like '%%%v%%'", v)
			} else {
				query += fmt.Sprintf("(a_address like '%%%v%%' or a_address2 like '%%%v%%')", v, v)
			}
		}

		query = fmt.Sprintf("(%v)", query)
		args = append(args, models.Custom{Query: query})
	}

	if typeid != 0 {
		args = append(args, models.Where{Column: "type", Value: typeid, Compare: "="})
	}

	if status != 0 {
		args = append(args, models.Where{Column: "status", Value: status, Compare: "="})
	}

	if page != 0 && pagesize != 0 {
		args = append(args, models.Paging(page, pagesize))
	}

	orderby := "id desc"
	args = append(args, models.Ordering(orderby))

	types := []repairlist.Repairtype{repairlist.RepairtypeEstablishment, repairlist.RepairtypeReview}
	args = append(args, models.Where{Column: "repairtype", Value: types, Compare: "in"})
	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)
}

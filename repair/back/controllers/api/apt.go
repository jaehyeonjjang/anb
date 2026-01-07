package api

import (
	"fmt"
	"repair/controllers"
	"repair/models"
	"strings"
)

type AptController struct {
	controllers.Controller
}

func (c *AptController) Search() {
	page := c.Geti("page")
	pagesize := c.Geti("pagesize")
	search := c.Get("search")
	contracttype := c.Geti("contracttype")

	conn := c.NewConnection()

	manager := models.NewAptManager(conn)

	var args []interface{}

	if search != "" {
		strs := strings.Split(search, " ")

		query := ""
		for i, v := range strs {
			if i > 0 {
				query += " and "
			}

			query += fmt.Sprintf("(a_name like '%%%v%%' or a_address like '%%%v%%' or a_address2 like '%%%v%%' or a_tel like '%%%v%%' or a_fax like '%%%v%%' or a_personalemail like '%%%v%%' or a_email like '%%%v%%')", v, v, v, v, v, v, v)
		}

		query = fmt.Sprintf("(%v)", query)
		args = append(args, models.Custom{Query: query})
	}

	if contracttype != 0 {
		query := "(a_contracttype & 1 or a_contracttype & 2 or a_contracttype & 4 or a_contracttype & 128)"
		args = append(args, models.Custom{Query: query})
	}

	if page != 0 && pagesize != 0 {
		args = append(args, models.Paging(page, pagesize))
	}

	orderby := "id desc"
	args = append(args, models.Ordering(orderby))

	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)
}

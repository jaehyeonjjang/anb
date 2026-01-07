package api

import (
	"anb/controllers"
	"anb/models"
)

type PictureController struct {
	controllers.Controller
}

func (c *PictureController) AjaxList() {
	conn := c.NewConnection()

	apt := c.Geti64("apt")

	manager := models.NewPictureManager(conn)
	items := manager.GetListByApt(apt, 0, 0, "")

	c.Set("items", items)
}

func (c *PictureController) AjaxUpload() {
	conn := c.NewConnection()

	category := c.Get("category")
	content := c.Get("content")

	file := c.GetUpload("upload")

	manager := models.NewPictureManager(conn)

	item := models.Picture{Picturecategory: 0, Filename: file, Category: category, Content: content}
	manager.Insert(&item)
}

func (c *PictureController) AjaxInsert() {
	conn := c.NewConnection()

	apt := c.Geti64("apt")

	category := c.Get("category")
	content := c.Get("content")
	filename := c.Get("filename")

	manager := models.NewPictureManager(conn)

	item := models.Picture{Apt: apt, Picturecategory: 0, Filename: filename, Category: category, Content: content}
	manager.Insert(&item)
}

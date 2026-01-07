package api

import (
	"anb/controllers"
	"anb/models"
)

type ImagefloorController struct {
	controllers.Controller
}

func (c *ImagefloorController) AjaxGet() {
	conn := c.NewConnection()

	image := c.Geti64("image")
	name := c.Get("name")
	imagename := c.Get("imagename")

	manager := models.NewImagefloorManager(conn)

	item := manager.GetByImageNameImagename(image, name, imagename)

	var id int64
	if item == nil {
		imageManager := models.NewImageManager(conn)
		parent := imageManager.Get(image)

		imageItem := models.Image{Apt: parent.Apt, Name: name, Level: parent.Level + 1, Parent: parent.Id, Last: 1, Title: "", Type: 9}
		imageManager.Insert(&imageItem)

		imageId := imageManager.GetIdentity()

		var imagefloor = models.Imagefloor{Image: image, Name: name, Imagename: imagename, Target: imageId}
		manager.Insert(&imagefloor)

		id = manager.GetIdentity()

		item = &imagefloor
		item.Id = id
	} else {
		id = item.Id
	}

	c.Set("item", item)
}

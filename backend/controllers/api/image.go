package api

import (
	"anb/config"
	"anb/controllers"
	"anb/models"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

type ImageController struct {
	controllers.Controller
}

func (c *ImageController) AjaxList() {
	apt := c.Geti64("apt")
	image := c.Geti64("image")

	if config.LocalMode == "true" {
		items := make([]models.Image, 0)

		images := models.GetImages()

		for _, item := range images {
			if apt == item.Apt {
				items = append(items, item)
			}
		}

		c.Set("items", items)

		return
	}

	conn := c.NewConnection()

	manager := models.NewImageManager(conn)

	var items *[]models.Image
	items = manager.GetListByAptParent(apt, image, 0, 0, "order, i_id")

	c.Set("items", items)

	if image > 0 {
		item := manager.Get(image)
		c.Set("image", item)
	}
}

func (c *ImageController) AjaxInsert() {
	conn := c.NewConnection()

	apt := c.Geti64("apt")
	name := c.Get("name")
	parent := c.Geti64("image")
	standard := c.Geti("standard")

	manager := models.NewImageManager(conn)

	item := manager.Get(parent)

	order := 0
	if parent > 0 {
		parentImage := manager.Get(parent)
		if parentImage != nil {
			order = parentImage.Order
		}
	}
	image := models.Image{Apt: apt, Name: name, Level: item.Level + 1, Parent: parent, Last: 1, Title: "", Type: 8, Order: order, Standard: standard}
	manager.Insert(&image)
}

func (c *ImageController) AjaxDelete() {
	conn := c.NewConnection()

	id := c.Geti64("id")

	manager := models.NewImageManager(conn)
	manager.Delete(id)
}

func (c *ImageController) AjaxView() {
	conn := c.NewConnection()

	id := c.Geti64("id")

	if config.LocalMode == "true" {
		images := models.GetImages()

		for _, item := range images {
			if id == item.Id {
				c.Set("item", item)

				return
			}
		}

		return
	}

	manager := models.NewImageManager(conn)
	item := manager.Get(id)
	c.Set("item", item)
}

func (c *ImageController) AjaxUpload() {
	file := c.GetUpload("upload")

	c.Set("filename", file)
}

func (c *ImageController) AjaxUpload2() {
	file := c.GetUpload("upload")

	apt := c.Geti64("apt")
	image := c.Geti64("image")

	fn := fmt.Sprintf("webdata/%v-%v.png", apt, image)
	os.Rename(file, fn)
	c.Set("filename", file)
}

func (c *ImageController) AjaxUploadandroidbase64() {
	/*
		conn := c.NewConnection()

		file := c.GetUpload("upload")

		image := c.Geti64("image")
		filename := c.Get("filename")

		data, err := ioutil.ReadFile(file)
		if err != nil {
			return
		}

		img, _ := base64.StdEncoding.DecodeString(string(data[22:]))

		manager := models.NewImageManager(conn)
		item := manager.Get(image)

		width, height := global.GetImageSize(item.Filename)

		if width > 0 && height > 0 {
			fullFilename := fmt.Sprintf("webdata/_%v", filename)
			ioutil.WriteFile(fullFilename, img, 0644)

			targetFilename := fmt.Sprintf("webdata/%v", filename)
			global.MakeThumbnailFault(1600, 1193, width, height, fullFilename, targetFilename)
		} else {
			fullFilename := fmt.Sprintf("webdata/%v", filename)
			ioutil.WriteFile(fullFilename, img, 0644)
		}

		filename = fmt.Sprintf("webdata/%v", filename)

		os.Remove(file)
		c.Set("filename", filename)
	*/

	file := c.GetUpload("upload")

	filename := c.Get("filename")

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	img, _ := base64.StdEncoding.DecodeString(string(data[22:]))

	filename = fmt.Sprintf("webdata/%v", filename)
	ioutil.WriteFile(filename, img, 0644)

	os.Remove(file)
	c.Set("filename", filename)
}

func (c *ImageController) AjaxUploadandroid() {
	file := c.GetUpload("upload")

	filename := c.Get("filename")

	target := fmt.Sprintf("webdata/%v", filename)

	os.Rename(file, target)
	c.Set("filename", target)
}

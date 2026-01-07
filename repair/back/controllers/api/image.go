package api

import (
	"log"
	"path"
	"repair/controllers"
	"repair/global"
	"repair/global/config"
)

type ImageController struct {
	controllers.Controller
}

func (c *ImageController) Size(filename string) {
	fullFilename := path.Join(config.UploadPath, filename)
	log.Println(fullFilename)
	width, height := global.GetImageSize(fullFilename)

	c.Set("width", width)
	c.Set("height", height)
}

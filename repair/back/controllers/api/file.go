package api

import (
	"os"
	"path"
	"repair/global/config"
	"repair/controllers"
	"repair/models"
)

type FileController struct {
	controllers.Controller
}

func (c *FileController) Pre_Delete(item *models.File) {
	conn := c.NewConnection()

	manager := models.NewFileManager(conn)
	file := manager.Get(item.Id)
	fullFilename := path.Join(config.UploadPath, file.Filename)
	os.Remove(fullFilename)
}

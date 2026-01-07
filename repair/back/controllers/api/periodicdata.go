package api

import (
	"fmt"
	"os"
	"repair/controllers"
	"repair/global/config"
	"repair/models"
)

type PeriodicdataController struct {
	controllers.Controller
}

func (c *PeriodicdataController) Post_Index(items []models.Periodicdata) {
	for i, v := range items {
		filename := fmt.Sprintf("%v/periodicresult/%v/%v.jpg", config.UploadPath, v.Periodic, v.Extra["blueprint"].(models.Blueprint).Id)
		_, error := os.Stat(filename)

		if os.IsNotExist(error) {
			items[i].AddExtra("resultimage", "")
		} else {
			items[i].AddExtra("resultimage", fmt.Sprintf("periodicresult/%v/%v.jpg", v.Periodic, v.Extra["blueprint"].(models.Blueprint).Id))
		}
	}
}

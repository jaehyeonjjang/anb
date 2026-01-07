package api

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"repair/controllers"
	"repair/global"
	"repair/global/config"
	"time"

	"repair/models"

	"github.com/karmdip-mi/go-fitz"
)

type PeriodicimageController struct {
	controllers.Controller
}

func (c *PeriodicimageController) Pre_Delete(item *models.Periodicimage) {
	conn := c.NewConnection()

	periodicimageManager := models.NewPeriodicimageManager(conn)
	periodicimageItem := periodicimageManager.Get(item.Id)

	removePeriodicimageFile(periodicimageItem)
}

func (c *PeriodicimageController) Pre_Deletebatch(items *[]models.Periodicimage) {
	if len(*items) == 0 {
		return
	}

	for _, v := range *items {
		removePeriodicimageFile(&v)
	}
}

func removePeriodicimageFile(item *models.Periodicimage) {
	fullFilename := fmt.Sprintf("%v/periodic/%v", config.UploadPath, item.Filename)
	os.Remove(fullFilename)
}

// @Post()
func (c *PeriodicimageController) Process(item *models.Periodicimage) {
	conn := c.NewConnection()

	periodicimageManager := models.NewPeriodicimageManager(conn)
	items := periodicimageManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: item.Periodic, Compare: "="},
	})

	max := 0
	for _, v := range items {
		if v.Order > max {
			max = v.Order
		}
	}

	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, item.Filename)
	doc, err := fitz.New(fullFilename)
	if err != nil {
		log.Println(err)
		return
	}

	now := global.GetDatetime(time.Now())
	for i := 0; i < doc.NumPage(); i++ {
		img, err := doc.Image(i)
		if err != nil {
			log.Println(err)
			continue
		}

		id := global.UniqueId()
		filename := fmt.Sprintf("periodic/pastimage-%v.jpg", id)
		fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
		f, err := os.Create(fullFilename)
		if err != nil {
			log.Println(err)
			continue
		}

		jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		f.Close()

		item := models.Periodicimage{Type: item.Type, Filename: filename, Name: item.Name, Use: item.Use, Order: max + i + 1, Periodic: item.Periodic, Date: now}
		periodicimageManager.Insert(&item)

	}

}

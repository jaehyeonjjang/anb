package api

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"repair/controllers"
	"repair/global"
	"repair/global/config"
	"strings"
	"time"

	"repair/models"

	"github.com/karmdip-mi/go-fitz"
)

type ManagebookController struct {
	controllers.Controller
}

func (c *ManagebookController) Pre_Delete(item *models.Managebook) {
	conn := c.NewConnection()

	managebookManager := models.NewManagebookManager(conn)
	managebookItem := managebookManager.Get(item.Id)

	removeFile(managebookItem)
}

func (c *ManagebookController) Pre_Deletebatch(items *[]models.Managebook) {
	if len(*items) == 0 {
		return
	}

	for _, v := range *items {
		removeFile(&v)
	}
}

func removeFile(item *models.Managebook) {
	fullFilename := fmt.Sprintf("%v/periodicresult/%v/%v", config.UploadPath, item.Periodic, item.Filename)
	os.Remove(fullFilename)
}

// @Post()
func (c *ManagebookController) Process(id int64, name string, order int, filename string) {
	conn := c.NewConnection()

	managebookManager := models.NewManagebookManager(conn)
	managebookcategoryManager := models.NewManagebookcategoryManager(conn)

	categorys := managebookcategoryManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "name", Value: name, Compare: "="},
	})

	now := global.GetDatetime(time.Now())
	var categoryId int64
	if len(categorys) > 0 {
		category := categorys[0]
		managebookManager.DeleteByManagebookcategory(category.Id)

		categoryId = category.Id
	} else {
		category := models.Managebookcategory{Name: name, Order: order, Periodic: id, Date: now}
		managebookcategoryManager.Insert(&category)
		categoryId = managebookcategoryManager.GetIdentity()
	}

	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	doc, err := fitz.New(fullFilename)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < doc.NumPage(); i++ {
		img, err := doc.Image(i)
		if err != nil {
			log.Println(err)
			continue
		}

		os.Mkdir(fmt.Sprintf("%v/periodicresult/%v", config.UploadPath, id), 0755)

		filename := fmt.Sprintf("managebook-%v.jpg", global.UniqueId())
		fullFilename := fmt.Sprintf("%v/periodicresult/%v/%v", config.UploadPath, id, filename)
		f, err := os.Create(fullFilename)
		if err != nil {
			log.Println(err)
			continue
		}

		jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		f.Close()

		item := models.Managebook{Filename: filename, Order: i + 1, Periodic: id, Managebookcategory: categoryId, Date: now}
		managebookManager.Insert(&item)

	}

}

// @Post()
func (c *ManagebookController) Multiprocess(id int64, filename string, originalfilename string) {
	conn := c.NewConnection()

	managebookManager := models.NewManagebookManager(conn)
	managebookcategoryManager := models.NewManagebookcategoryManager(conn)

	filenames := strings.Split(filename, ",")
	originalfilenames := strings.Split(originalfilename, ",")

	for j, name := range originalfilenames {
		name = strings.ReplaceAll(name, ".pdf", "")
		filename := filenames[j]

		categorys := managebookcategoryManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Where{Column: "name", Value: name, Compare: "="},
		})

		now := global.GetDatetime(time.Now())
		var categoryId int64
		if len(categorys) > 0 {
			category := categorys[0]
			managebookManager.DeleteByManagebookcategory(category.Id)

			categoryId = category.Id
		} else {
			category := models.Managebookcategory{Name: name, Order: 0, Periodic: id, Date: now}
			managebookcategoryManager.Insert(&category)
			categoryId = managebookcategoryManager.GetIdentity()
		}

		fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
		doc, err := fitz.New(fullFilename)
		if err != nil {
			log.Println(err)
			return
		}

		for i := 0; i < doc.NumPage(); i++ {
			img, err := doc.Image(i)
			if err != nil {
				log.Println(err)
				continue
			}

			os.Mkdir(fmt.Sprintf("%v/periodicresult/%v", config.UploadPath, id), 0755)

			filename := fmt.Sprintf("managebook-%v.jpg", global.UniqueId())
			fullFilename := fmt.Sprintf("%v/periodicresult/%v/%v", config.UploadPath, id, filename)
			f, err := os.Create(fullFilename)
			if err != nil {
				log.Println(err)
				continue
			}

			jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			f.Close()

			item := models.Managebook{Filename: filename, Order: i + 1, Periodic: id, Managebookcategory: categoryId, Date: now}
			managebookManager.Insert(&item)

		}
	}
}

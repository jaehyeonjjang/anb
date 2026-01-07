package api

import (
	"anb/controllers"
	"anb/models"
	"log"
)

type JsonController struct {
	controllers.Controller
}

func (c *JsonController) AjaxList() {
	conn := c.NewConnection()

	apt := c.Geti64("apt")

	aptManager := models.NewAptManager(conn)
	aptItem := aptManager.Get(apt)

	log.Println("apt = ", apt)
	log.Println(aptItem)

	apts := aptManager.GetListByCompanyStatus(aptItem.Company, 1, 0, 0, "")
	c.Set("apts", apts)

	aptuserManager := models.NewAptuserManager(conn)
	aptusers := aptuserManager.GetListByApt(apt, 0, 0, "")
	c.Set("aptusers", aptusers)

	imageManager := models.NewImageManager(conn)
	images := imageManager.GetListByApt(apt, 0, 0, "order")
	c.Set("images", images)

	imagefloorManager := models.NewImagefloorManager(conn)
	imagefloors := imagefloorManager.GetList(0, 0, "")

	if imagefloors != nil {
		var items []models.Imagefloor

		for _, item := range *imagefloors {
			flag := false
			if images != nil {
				for _, image := range *images {
					if item.Image == image.Id {
						flag = true
						break
					}
				}
			} else {
				flag = true
			}

			if flag == false {
				continue
			}

			items = append(items, item)
		}
		c.Set("imagefloors", items)
	} else {
		c.Set("imagefloors", imagefloors)
	}

	statusManager := models.NewStatusManager(conn)
	statuss := statusManager.GetListByCompany(aptItem.Company, 0, 0, "order, s_name")
	c.Set("statuss", statuss)

	statuscategoryManager := models.NewStatuscategoryManager(conn)
	statuscategorys := statuscategoryManager.GetListByCompany(aptItem.Company, 0, 0, "order, sc_name")
	c.Set("statuscategorys", statuscategorys)

	userManager := models.NewUserManager(conn)
	users := userManager.GetListByCompanyStatus(aptItem.Company, 1, 0, 0, "")
	c.Set("users", users)
}

func (c *JsonController) AjaxData() {
	conn := c.NewConnection()

	apt := c.Geti64("apt")

	dataManager := models.NewDataManager(conn)
	datas := dataManager.GetListByApt(apt, 0, 0, "")
	c.Set("datas", datas)
}

func (c *JsonController) AjaxCheck() {
}

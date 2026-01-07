package services

import (
	"anb/models"
	"log"
)

func Summary(id int64) string {
	conn := models.NewConnection()
	defer conn.Close()

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(id)

	imageManager := models.NewImageManager(conn)
	images := imageManager.GetListByApt(id, 0, 0, "order, i_id")

	manager := models.NewDataManager(conn)

	if images == nil {
		return ""
	}

	for _, image := range *images {
		if image.Type == 1 || image.Type == 8 {
			items := manager.GetListByAptImage(apt.Id, image.Id, 0, 0, "id")

			if items == nil {
				continue
			}

			for _, item := range *items {
				if !isNumber(item.Type) {
					continue
				}

				if item.Name == "" {
					continue
				}

				if item.Count == "" || item.Count == "0" {
					item.Count = "1"
				}

				item.Group = image.Type

				log.Println(item)
			}
		}
	}

	return ""
}

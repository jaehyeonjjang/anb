package api

import (
	"anb/config"
	"anb/controllers"
	"anb/models"
	"encoding/json"
)

type SyncController struct {
	controllers.Controller
}

func (c *SyncController) AjaxInsert() {
	image := c.Geti64("image")

	var item models.Sync

	item.Image = image

	if config.LocalMode == "true" {
		models.InsertSync(item)
	} else {
		conn := c.NewConnection()
		manager := models.NewSyncManager(conn)
		manager.Insert(&item)
	}
}

func (c *SyncController) AjaxSync() {
	if config.LocalMode != "true" {
		return
	}

	models.DatabaseUploadSync()
}

func (c *SyncController) AjaxDownload() {
	models.DatabaseSync()
}

func (c *SyncController) AjaxUpload() {
	conn := c.NewConnection()

	dataManager := models.NewDataManager(conn)
	imageManager := models.NewImageManager(conn)
	imagefloorManager := models.NewImagefloorManager(conn)

	raw, _ := c.Context.GetRawData()

	var content models.Datas
	json.Unmarshal(raw, &content)

	imageMap := make(map[int64]int64)

	for _, item := range content.Images {
		oldId := item.Id

		item.Id = 0

		parentImage := imageManager.Get(item.Parent)

		if parentImage == nil {
			continue
		}

		item.Order = parentImage.Order
		imageManager.Insert(&item)
		id := imageManager.GetIdentity()

		imageMap[oldId] = id
	}

	for _, item := range content.Imagefloors {
		target := item.Target

		newTarget := imageMap[target]

		item.Id = 0
		item.Target = newTarget
		imagefloorManager.Insert(&item)
	}

	syncArray := make(map[int64]models.Sync)

	for _, sync := range content.Syncs {
		syncArray[sync.Image] = sync
	}

	for _, sync := range syncArray {
		dataManager.DeleteByImage(sync.Image)

		for _, item := range content.Datas {
			if item.Image < 0 {
				item.Image = imageMap[item.Image]
			} else {
				if item.Image != sync.Image {
					continue
				}
			}

			item.Id = 0

			dataManager.Insert(&item)
		}
	}
}

func (c *SyncController) AjaxComplete() {
	conn := c.NewConnection()

	image := c.Geti64("image")

	manager := models.NewImageManager(conn)
	item := manager.Get(image)

	reportManager := models.NewReportManager(conn)

	var report models.Report
	report.Apt = item.Apt
	report.Image = item.Id
	report.Status = 2

	reportManager.Insert(&report)
}

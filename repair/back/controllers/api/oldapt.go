package api

import (
	"fmt"
	"log"
	"repair/controllers"
	"repair/global"
	"repair/global/config"
	"repair/models"
	"strings"
)

type OldaptController struct {
	controllers.Controller
}

func convert(conn *models.Connection, imageId int64, periodicId int64, blueprintId int64) {
	oldimageManager := models.NewOldimageManager(conn)
	olddataManager := models.NewOlddataManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)

	image := oldimageManager.Get(imageId)

	url := fmt.Sprintf("http://anbweb.kr/%v", image.Filename)

	filename := strings.ReplaceAll(image.Filename, "webdata/", "")
	fullFilename := fmt.Sprintf("%v/blueprint/%v", config.UploadPath, filename)

	blueprintManager.UpdateFilenameById(fmt.Sprintf("blueprint/%v", filename), blueprintId)

	global.GetFile(url, fullFilename)

	olds := olddataManager.Find([]interface{}{
		models.Where{Column: "image", Value: image.Id, Compare: "="},
		models.Ordering("d_id"),
	})

	periodicdataManager.DeleteByPeriodicBlueprint(periodicId, blueprintId)
	periodicblueprintzoomManager.DeleteByPeriodicBlueprint(periodicId, blueprintId)
	order := 1

	zoom := 4.0
	for _, v := range olds {
		typeid := 0
		content := ""

		if v.Type == 100 {
			zoom = global.Atof(v.Content) * 6
			continue
		}

		content = fmt.Sprintf(`[{"dx":%v,"dy":%v}]`, v.X, v.Y)
		switch v.Type {
		case 20:
			typeid = 1
		case 21:
			typeid = 2
		case 1:
			typeid = 31
			content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
			content = strings.ReplaceAll(content, `"y"`, `"dy"`)
		case 15:
			typeid = 32
			content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
			content = strings.ReplaceAll(content, `"y"`, `"dy"`)
		case 53:
			typeid = 33
			content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
			content = strings.ReplaceAll(content, `"y"`, `"dy"`)
		case 2:
			typeid = 41
			content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
			content = strings.ReplaceAll(content, `"y"`, `"dy"`)
		case 16:
			typeid = 42
			content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
			content = strings.ReplaceAll(content, `"y"`, `"dy"`)
		case 54:
			typeid = 43
			content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
			content = strings.ReplaceAll(content, `"y"`, `"dy"`)
		case 660:
			typeid = 101
		case 670:
			typeid = 102
		case 620:
			typeid = 103
		case 630:
			typeid = 104
		case 611:
			typeid = 105
		case 680:
			typeid = 106
		case 640:
			typeid = 107
		case 641:
			typeid = 108
		case 642:
			typeid = 109
		case 643:
			typeid = 110
		case 591:
			typeid = 111
		default:
			continue
		}

		item := models.Periodicdata{}

		if v.Filename != "" {
			url := fmt.Sprintf("http://anbweb.kr/%v", v.Filename)
			filename := strings.ReplaceAll(v.Filename, "webdata/", "")
			fullFilename := fmt.Sprintf("%v/periodic/%v", config.UploadPath, filename)

			global.GetFile(url, fullFilename)
			item.Filename = fmt.Sprintf("periodic/%v", filename)
		}

		item.Group = v.Number
		item.Type = typeid
		item.Part = v.Fault
		item.Member = v.Name
		item.Shape = v.Content
		item.Width = fmt.Sprintf("%v", v.Width)
		item.Length = fmt.Sprintf("%v", v.Length)
		item.Count = global.Atoi(v.Count)
		item.Content = content

		if v.Progress == "X" {
			item.Progress = 2
		} else {
			item.Progress = 1
		}

		item.Remark = v.Remark
		item.Order = order
		item.Status = 1
		item.Blueprint = blueprintId
		item.Periodic = periodicId
		item.Date = v.Date
		order++

		periodicdataManager.Insert(&item)
	}

	periodicblueprintzoom := models.Periodicblueprintzoom{Iconzoom: models.Double(zoom), Zoom: models.Double(0.8), Status: 1, Blueprint: blueprintId, Periodic: periodicId}
	periodicblueprintzoomManager.Insert(&periodicblueprintzoom)
}

// @POST()
func (c *OldaptController) Convert(items *models.ConvertOldapt) {
	conn := c.NewConnection()

	conn.Begin()
	defer conn.Rollback()

	for _, v := range items.Items {
		log.Println(v.Id, "=>", v.Company)

		if v.Company == 0 {
			continue
		}
		convert(conn, v.Id, v.Master, v.Company)
	}

	conn.Commit()
}

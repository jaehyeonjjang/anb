package api

import (
	"anb/config"
	"anb/controllers"
	"anb/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type DataController struct {
	controllers.Controller
}

func (c *DataController) AjaxList() {
	apt := c.Geti64("apt")
	image := c.Geti64("image")

	if config.LocalMode == "true" {
		items := make([]models.Data, 0)

		datas := models.GetDatas()

		for _, item := range datas {
			if item.Apt == apt && item.Image == image {
				items = append(items, item)
			}
		}

		c.Set("items", items)

		return
	}

	conn := c.NewConnection()

	manager := models.NewDataManager(conn)
	items := manager.GetListByAptImage(apt, image, 0, 0, "")

	c.Set("items", items)
}

func (c *DataController) AjaxDelete() {
	image := c.Geti64("image")

	if config.LocalMode == "true" {
		models.DeleteDatas(image)
		return
	}

	conn := c.NewConnection()

	manager := models.NewDataManager(conn)

	//manager.DeleteByImageUser(image, c.Session.Id)
	manager.DeleteByImageUser(image, 0)
}

func (c *DataController) AjaxInsert() {
	apt := c.Geti64("apt")
	image := c.Geti64("image")

	typeid := c.Geti("type")
	x := c.Getf("x")
	y := c.Getf("y")
	point := c.Get("point")
	number := c.Geti("number")
	group := c.Geti("group")
	name := c.Get("name")
	fault := c.Get("fault")
	content := c.Get("content")
	report := c.Geti("report")
	usermemo := c.Get("usermemo")
	aptmemo := c.Get("aptmemo")
	width := c.Getf("width")
	length := c.Getf("length")
	count := c.Get("count")
	progress := c.Get("progress")
	remark := c.Get("remark")
	imagename := c.Get("imagename")
	filename := c.Get("filename")

	var item models.Data

	item.Apt = apt
	item.Image = image
	//item.User = c.Session.Id

	item.Type = typeid
	item.X = x
	item.Y = y
	item.Point = point
	item.Number = number
	item.Group = group
	item.Name = name
	item.Fault = fault
	item.Content = content
	item.Report = report
	item.Usermemo = usermemo
	item.Aptmemo = aptmemo
	item.Width = width
	item.Length = length
	item.Count = count
	item.Progress = progress
	item.Remark = remark
	item.Imagename = imagename
	item.Filename = filename

	if config.LocalMode == "true" {
		models.InsertData(item)
	} else {
		conn := c.NewConnection()

		manager := models.NewDataManager(conn)
		manager.Insert(&item)
	}
}

func (c *DataController) AjaxImage() {
	log.Println("AjaxImage Start")

	apt := c.Geti64("apt")
	image := c.Geti64("image")
	data := c.Get("data")

	//log.Println(len(image))
	//log.Println(data[22:])

	img, _ := base64.StdEncoding.DecodeString(data[22:])

	//log.Println(img)
	filename := fmt.Sprintf("webdata/%v-%v.png", apt, image)
	log.Println("filename = ", filename)
	ioutil.WriteFile(filename, img, 0644)

	var item models.Sync

	item.Image = image

	if config.LocalMode == "true" {
		models.InsertSync(item)
	} else {
		conn := c.NewConnection()

		manager := models.NewSyncManager(conn)
		manager.Insert(&item)
	}

	log.Println("AjaxImage End")
}

type Data struct {
	Apt       int64   `json:"apt"`
	Image     int64   `json:"image"`
	Imagetype int     `json:"imagetype"`
	User      int64   `json:"user"`
	Type      int     `json:"type"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Point     string  `json:"point"`
	Number    int     `json:"number"`
	Group     int     `json:"group"`
	Name      string  `json:"name"`
	Fault     string  `json:"fault"`
	Content   string  `json:"content"`
	Report    int     `json:"report"`
	Usermemo  string  `json:"usermemo"`
	Aptmemo   string  `json:"aptmemo"`
	Width     float64 `json:"width"`
	Length    float64 `json:"length"`
	Count     string     `json:"count"`
	Progress  string  `json:"progress"`
	Remark    string  `json:"remark"`
	Imagename string  `json:"imagename"`
	Filename  string  `json:"filename"`
	Memo      string  `json:"memo"`
	Date      string  `json:"date"`
}

type DataItem struct {
	Image int64  `json:"image"`
	Datas []Data `json:"data"`
}

func (c *DataController) AjaxMultiinsert() {
	raw, _ := c.Context.GetRawData()

	var content DataItem
	json.Unmarshal(raw, &content)

	var manager *models.DataManager

	if config.LocalMode != "true" {
		conn := c.NewConnection()
		manager = models.NewDataManager(conn)

		manager.DeleteByImage(content.Image)
	}

	items := make([]models.Data, 0)

	for _, data := range content.Datas {
		var item models.Data

		item.Apt = data.Apt
		item.Image = data.Image
		item.Imagetype = data.Imagetype
		item.User = data.User
		item.Type = data.Type
		item.X = data.X
		item.Y = data.Y
		item.Point = data.Point
		item.Number = data.Number
		item.Group = data.Group
		item.Name = data.Name
		item.Fault = data.Fault
		item.Content = data.Content
		item.Report = data.Report
		item.Usermemo = data.Usermemo
		item.Aptmemo = data.Aptmemo
		item.Width = data.Width
		item.Length = data.Length
		item.Count = data.Count
		item.Progress = data.Progress
		item.Remark = data.Remark
		item.Imagename = data.Imagename
		item.Filename = data.Filename
		item.Memo = data.Memo

		//item.User = c.Session.Id

		log.Printf("item.Apt : %v\n", data.Apt)
		log.Printf("item.Image : %v\n", data.Image)
		log.Printf("item.Imagetype : %v\n", data.Imagetype)
		log.Printf("item.User : %v\n", data.User)
		log.Printf("item.Type : %v\n", data.Type)
		log.Printf("item.X : %v\n", data.X)
		log.Printf("item.Y : %v\n", data.Y)
		log.Printf("item.Point : %v\n", data.Point)
		log.Printf("item.Number : %v\n", data.Number)
		log.Printf("item.Group : %v\n", data.Group)
		log.Printf("item.Name : %v\n", data.Name)
		log.Printf("item.Fault : %v\n", data.Fault)
		log.Printf("item.Content : %v\n", data.Content)
		log.Printf("item.Report : %v\n", data.Report)
		log.Printf("item.Usermemo : %v\n", data.Usermemo)
		log.Printf("item.Aptmemo : %v\n", data.Aptmemo)
		log.Printf("item.Width : %v\n", data.Width)
		log.Printf("item.Length : %v\n", data.Length)
		log.Printf("item.Count : %v\n", data.Count)
		log.Printf("item.Progress : %v\n", data.Progress)
		log.Printf("item.Remark : %v\n", data.Remark)
		log.Printf("item.Imagename : %v\n", data.Imagename)
		log.Printf("item.Filename : %v\n", data.Filename)
		log.Printf("item.Memo : %v\n", data.Memo)
		log.Printf("===========================================\n")

		items = append(items, item)

		if config.LocalMode != "true" {
			manager.Insert(&item)
		}
	}

	if config.LocalMode == "true" {
		models.InsertMultiDatas(content.Image, items)
	}
}

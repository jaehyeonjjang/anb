package models

import (

	//_ "github.com/denisenkom/go-mssqldb"
	"anb/global"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Lists struct {
	Apts     []Apt
	Aptusers []Aptuser
	Images   []Image
	Statuss  []Status
	Users    []User
}

type Datas struct {
	Datas []Data
	Syncs []Sync
	Images []Image
	Imagefloors []Imagefloor
}

func GetAll() Lists {
	content := global.ReadFile("webdata/json/list.json")

	var lists Lists
	json.Unmarshal([]byte(content), &lists)
	//log.Println(lists)

	return lists
}

func GetApts() []Apt {
	lists := GetAll()

	return lists.Apts
}

func GetAptusers() []Aptuser {
	lists := GetAll()

	return lists.Aptusers
}

func GetImages() []Image {
	lists := GetAll()

	return lists.Images
}

func GetImage(id int64) *Image {
	lists := GetAll()

	for _, item := range lists.Images {
		if item.Id == id {
			return &item
		}
	}

	return nil
}

func GetStatuss() []Status {
	lists := GetAll()

	return lists.Statuss
}

func GetUsers() []User {
	lists := GetAll()

	return lists.Users
}

func GetDatas() []Data {
	content := global.ReadFile("webdata/json/data.json")

	var lists Datas
	json.Unmarshal([]byte(content), &lists)
	//log.Println(lists)

	return lists.Datas
}

func GetSyncs() []Sync {
	content := global.ReadFile("webdata/json/data.json")

	var lists Datas
	json.Unmarshal([]byte(content), &lists)
	//log.Println(lists)

	return lists.Syncs
}

func SaveDatas(datas string) {
	if global.WriteFile("webdata/json/data.json", datas) != nil {
		log.Println("write file error")
		return
	}
}

func DeleteDatas(image int64) {
	datas := GetDatas()

	items := make([]Data, 0)

	for _, item := range datas {
		if item.Image != image {
			items = append(items, item)
		}
	}

	b, _ := json.Marshal(items)
	SaveDatas(string(b))
}

func TruncateSync() {
	content := global.ReadFile("webdata/json/data.json")

	var lists Datas
	json.Unmarshal([]byte(content), &lists)

	lists.Syncs = make([]Sync, 0)
	b, _ := json.Marshal(lists)
	SaveDatas(string(b))
}

func InsertData(item Data) {
	content := global.ReadFile("webdata/json/data.json")

	var lists Datas
	json.Unmarshal([]byte(content), &lists)

	datas := lists.Datas

	var max int64 = 0
	for _, item := range datas {
		if item.Id > max {
			max = item.Id
		}
	}

	item.Id = max + 1

	t := time.Now()
	item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	lists.Datas = append(lists.Datas, item)

	b, _ := json.Marshal(lists)
	SaveDatas(string(b))
}

func InsertSync(item Sync) {
	content := global.ReadFile("webdata/json/data.json")

	var lists Datas
	json.Unmarshal([]byte(content), &lists)

	syncs := lists.Syncs

	var max int64 = 0
	for _, item := range syncs {
		if item.Id > max {
			max = item.Id
		}
	}

	item.Id = max + 1

	t := time.Now()
	item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	lists.Syncs = append(lists.Syncs, item)

	b, _ := json.Marshal(lists)
	SaveDatas(string(b))
}

func InsertMultiDatas(image int64, items []Data) {
	content := global.ReadFile("webdata/json/data.json")

	var lists Datas
	json.Unmarshal([]byte(content), &lists)

	datas := lists.Datas

	log.Printf("old datas : %v\n", len(datas))

	newItems := make([]Data, 0)

	var max int64 = 0
	for _, item := range datas {
		if item.Id > max {
			max = item.Id
		}

		if item.Image != image {
			newItems = append(newItems, item)
		}
	}

	t := time.Now()

	for _, item := range items {
		max++

		item.Id = max
		item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

		newItems = append(newItems, item)
	}

	lists.Datas = newItems

	log.Printf("new datas : %v\n", len(lists.Datas))

	b, _ := json.Marshal(lists)
	SaveDatas(string(b))
}

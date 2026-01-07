package main

import (
	"fmt"
	"path"
	"repair/global"
	"repair/global/config"
	"repair/global/log"
	"repair/models"
	"testing"
)

/*
func TestCategory(t *testing.T) {

			conn := models.NewConnection()

			periodicManager := models.NewPeriodicManager(conn)
			periodicotherManager := models.NewPeriodicotherManager(conn)

			items := periodicManager.Find(nil)

			for _, v := range items {
				id := v.Id
				periodicotherManager.Insert(&models.Periodicother{Name: "해당사항 없음,해당사항 있음", Status: "해당사항 없음", Type: 1, Position: "강재구조 노후", Category: 15, Order: 150, Periodic: id})
				periodicotherManager.Insert(&models.Periodicother{Name: "부식,파손,이격", Type: 2, Position: "비구조형강", Category: 15, Order: 151, Periodic: id})
				periodicotherManager.Insert(&models.Periodicother{Name: "부식,파손,볼트풀림", Type: 2, Position: "철골구조물접합부위", Category: 15, Order: 152, Periodic: id})
				periodicotherManager.Insert(&models.Periodicother{Name: "들뜸,탈락", Type: 2, Position: "내화피복", Category: 15, Order: 153, Periodic: id})
			}
	}

	func TestConvert(t *testing.T) {
		conn := models.NewConnection()
		defer conn.Close()

		oldimageManager := models.NewOldimageManager(conn)
		olddataManager := models.NewOlddataManager(conn)
		blueprintManager := models.NewBlueprintManager(conn)
		periodicdataManager := models.NewPeriodicdataManager(conn)
		periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)

		image := oldimageManager.Get(70343)

		var periodicId int64 = 107
		var blueprintId int64 = 6153

		url := fmt.Sprintf("https://anbweb.kr/%v", image.Filename)
		log.Println(image)

		filename := strings.ReplaceAll(image.Filename, "webdata/", "")
		fullFilename := fmt.Sprintf("%v/blueprint/%v", config.UploadPath, filename)

		blueprintManager.UpdateFilename(fmt.Sprintf("blueprint/%v", filename), blueprintId)

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
			if v.Type == 20 {
				typeid = 1
			} else if v.Type == 21 {
				typeid = 2
			} else if v.Type == 1 {
				typeid = 31
				content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
				content = strings.ReplaceAll(content, `"y"`, `"dy"`)
			} else if v.Type == 15 {
				typeid = 32
				content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
				content = strings.ReplaceAll(content, `"y"`, `"dy"`)
			} else if v.Type == 53 {
				typeid = 33
				content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
				content = strings.ReplaceAll(content, `"y"`, `"dy"`)
			} else if v.Type == 2 {
				typeid = 41
				content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
				content = strings.ReplaceAll(content, `"y"`, `"dy"`)
			} else if v.Type == 16 {
				typeid = 42
				content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
				content = strings.ReplaceAll(content, `"y"`, `"dy"`)
			} else if v.Type == 54 {
				typeid = 43
				content = strings.ReplaceAll(v.Point, `"x"`, `"dx"`)
				content = strings.ReplaceAll(content, `"y"`, `"dy"`)
			} else if v.Type == 660 {
				typeid = 101
			} else if v.Type == 670 {
				typeid = 102
			} else if v.Type == 620 {
				typeid = 103
			} else if v.Type == 630 {
				typeid = 104
			} else if v.Type == 611 {
				typeid = 105
			} else if v.Type == 680 {
				typeid = 106
			} else if v.Type == 640 {
				typeid = 107
			} else if v.Type == 641 {
				typeid = 108
			} else if v.Type == 642 {
				typeid = 109
			} else if v.Type == 643 {
				typeid = 110
			} else if v.Type == 591 {
				typeid = 111
			} else {
				continue
			}

			item := models.Periodicdata{}

			if v.Filename != "" {
				url := fmt.Sprintf("https://anbweb.kr/%v", v.Filename)
				filename := strings.ReplaceAll(v.Filename, "webdata/", "")
				fullFilename := fmt.Sprintf("%v/periodic/%v", config.UploadPath, filename)

				global.GetFile(url, fullFilename)
				item.Filename = fmt.Sprintf("periodic/%v", filename)
			}

			item.Group = v.Number
			item.Type = typeid
			item.Part = v.Name
			item.Member = v.Fault
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

		periodicblueprintzoom := models.Periodicblueprintzoom{Iconzoom: zoom, Zoom: 0.8, Status: 1, Blueprint: blueprintId, Periodic: periodicId}
		periodicblueprintzoomManager.Insert(&periodicblueprintzoom)
	}
*/
/*
func TestFile(t *testing.T) {
	data, err := os.ReadFile("doc/detail/detail-00.jet")
	if err != nil {
		log.Println(err)
		return
	}

	str := string(data)

	lines := strings.Split(str, ">")

	indent := 0
	for _, line := range lines {
		flag := false
		if line[0:2] == "</" {
			if indent > 0 {
				indent--
			}

			flag = true
		}

		fmt.Println(strings.Repeat("\t", indent) + line + ">")

		if flag == false && line[0:1] == "<" {
			if strings.Contains(line, "</") {
			} else if line[len(line)-1:] == "/" {
			} else {
				indent++
			}
		} else {
			if strings.Contains(line, "</") {
				indent -= 1
				if indent < 0 {
					indent = 0
				}

			}
		}
	}
}
*/

/*
func TestDetail0(t *testing.T) {
		conn := models.NewConnection()

		blueprintManager := models.NewBlueprintManager(conn)
		periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
		periodicdataManager := models.NewPeriodicdataManager(conn)

		items := periodicblueprintzoomManager.Find([]interface{}{
			models.Where{Column: "periodic", Value: 113, Compare: "="},
			models.Ordering("pb_id"),
		})

		for _, v := range items {
			blueprint := blueprintManager.Get(v.Blueprint)
			periodicdatas := periodicdataManager.Find([]interface{}{
				models.Where{Column: "periodic", Value: v.Periodic, Compare: "="},
				models.Where{Column: "blueprint", Value: v.Blueprint, Compare: "="},
				models.Ordering("pd_order"),
			})

			api.MakeImage(v.Periodic, *blueprint, periodicdatas, float64(v.Iconzoom))

			inclination := false
			fiber := false
			meterial := false
			for _, v2 := range periodicdatas {
				if v2.Type >= 200 && v2.Type < 300 {
					inclination = true
				} else if v2.Type >= 300 && v2.Type < 400 {
					fiber = true
				} else if v2.Type >= 400 && v2.Type < 500 {
					meterial = true
				}
			}

			if inclination == true {
				image.MakeInclinationImage(v.Periodic, *blueprint, periodicdatas, float64(v.Iconzoom))
			}

			if fiber == true {
				image.MakeFiberImage(v.Periodic, *blueprint, periodicdatas, float64(v.Iconzoom))
			}

			if meterial == true {
				image.MakeMeterialImage(v.Periodic, *blueprint, periodicdatas, float64(v.Iconzoom))
			}
		}

		// conn := models.NewConnection()
		// var id int64 = 113
		// filename := detail.Detail0(id, conn)
		// log.Println(filename)
		// cmd := exec.Command("open", fmt.Sprintf("./webdata/detail"))
		// cmd.Run()
}
*/

// func TestMakeImage(t *testing.T) {
// 	conn := models.NewConnection()
// 	periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
// 	blueprintManager := models.NewBlueprintManager(conn)
// 	periodicdataManager := models.NewPeriodicdataManager(conn)

// 	var periodic int64 = 880
// 	var blueprintId int64 = 7863
// 	periodicblueprintzoom := periodicblueprintzoomManager.GetByPeriodicBlueprint(periodic, blueprintId)

// 	log.Println("zoom", periodicblueprintzoom)
// 	blueprint := blueprintManager.Get(blueprintId)
// 	periodicdatas := periodicdataManager.Find([]interface{}{
// 		models.Where{Column: "periodic", Value: periodic, Compare: "="},
// 		models.Where{Column: "blueprint", Value: blueprintId, Compare: "="},
// 		models.Ordering("pd_order"),
// 	})

// 	log.Println("blueprint nil?")

// 	if blueprint != nil {
// 		log.Println("Make image", periodic)
// 		api.MakeImage(periodic, *blueprint, periodicdatas, float64(periodicblueprintzoom.Iconzoom))
// 	}
// }

// func TestCategoryConvert(t *testing.T) {
// 	log.Init()

// 	// repair.Change(24372)
// 	conn := models.NewConnection()
// 	defer conn.Close()

// 	standardManager := models.NewStandardManager(conn)
// 	categoryManager := models.NewCategoryManager(conn)

// 	items := categoryManager.Find([]interface{}{
// 		models.Where{Column: "name", Value: "(2) 옥상 비상문 자동 개폐장치", Compare: "="},
// 	})

// 	for _, v := range items {
// 		cnt := standardManager.Count([]interface{}{
// 			models.Where{Column: "apt", Value: v.Apt, Compare: "="},
// 			models.Where{Column: "name", Value: "옥상 비상", Compare: "like"},
// 		})

// 		if cnt > 0 {
// 			continue
// 		}

// 		cnt = standardManager.Count([]interface{}{
// 			models.Where{Column: "apt", Value: v.Apt, Compare: "="},
// 			models.Where{Column: "name", Value: "옥상비상", Compare: "like"},
// 		})

// 		if cnt > 0 {
// 			continue
// 		}

// 		log.Println("추가해야함", v.Apt)
// 		item := models.Standard{
// 			Name:     "옥상 비상문 자동 개폐장치",
// 			Direct:   235394,
// 			Labor:    0,
// 			Cost:     0,
// 			Unit:     "ea",
// 			Order:    0,
// 			Category: v.Id,
// 			Apt:      v.Apt,
// 		}

// 		standardManager.Insert(&item)
// 	}
// }

// func TestJosa(t *testing.T) {
// 	log.Init()

// 	str := global.GetJosa("균열(층간)", hangul.I_GA)
// 	log.Println(str)
// 	str = global.GetJosa("균열", hangul.I_GA)
// 	log.Println(str)

// 	str = "균열(층간)"
// 	var re = regexp.MustCompile(`\([^)]+\)`)
// 	str = re.ReplaceAllString(str, "")
// 	log.Println(str)
// }

// func TestImage(t *testing.T) {
// 	log.Init()

// 	img, err := global.LoadFromPngFile("./20250403084753_8e3affbec17a42cea1f5a32d933f01ce.jpeg")
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		log.Println(img)
// 	}
// }

// func TestBlueprint(t *testing.T) {
// 	log.Init()

// 	conn := models.NewConnection()
// 	defer conn.Close()

// 	blueprintManager := models.NewBlueprintManager(conn)
// 	items := blueprintManager.Find([]interface{}{
// 		models.Where{Column: "name", Value: "지하", Compare: "like"},
// 	})

// 	for _, v := range items {
// 		fmt.Printf("update blueprint_tb set bp_floortype = %v where bp_id = %v;\n", v.Floortype, v.Id)
// 	}
// }

// func TestElectricccar(t *testing.T) {
// 	log.Init()
// 	conn := models.NewConnection()
// 	defer conn.Close()

// 	categoryManager := models.NewCategoryManager(conn)
// 	var item *models.Category

// 	var apt int64 = 24674

// 	item = categoryManager.GetByAptName(apt, "가. 옥외부대시설 및 옥외복리시설")
// 	if item == nil {
// 		item = categoryManager.GetByAptName(apt, "가. 옥외부대시설 및 옥외 복리시설")
// 		if item == nil {
// 			item = categoryManager.GetByAptName(apt, "가. 옥외 부대시설 및 옥외 복리시설")
// 			if item == nil {
// 				items := categoryManager.FindByAptOrder(apt, 601000000)
// 				for _, v := range items {
// 					item = &v
// 				}
// 			}
// 		}
// 	}

// 	if item != nil {
// 		item2 := &models.Category{
// 			Name:     "(14) 전기자동차의 고정형 충전기",
// 			Level:    3,
// 			Parent:   item.Id,
// 			Cycle:    0,
// 			Percent:  0,
// 			Unit:     "",
// 			Elevator: 2,
// 			Remark:   "",
// 			Order:    601140000,
// 			Apt:      apt,
// 		}
// 		categoryManager.Insert(item2)
// 		item2.Id = categoryManager.GetIdentity()

// 		item3 := &models.Category{
// 			Name:     "부분수선",
// 			Level:    4,
// 			Parent:   item2.Id,
// 			Cycle:    5,
// 			Percent:  10,
// 			Unit:     "",
// 			Elevator: 2,
// 			Remark:   "",
// 			Order:    601140100,
// 			Apt:      apt,
// 		}

// 		categoryManager.Insert(item3)

// 		item3 = &models.Category{
// 			Name:     "전면교체",
// 			Level:    4,
// 			Parent:   item2.Id,
// 			Cycle:    10,
// 			Percent:  100,
// 			Unit:     "",
// 			Elevator: 2,
// 			Remark:   "",
// 			Order:    601140200,
// 			Apt:      apt,
// 		}
// 		categoryManager.Insert(item3)
// 	}
// }

func TestDownloadImage(t *testing.T) {
	log.Init()

	serverUrl := "https://service.anbweb.kr"

	conn := models.NewConnection()

	periodicManager := models.NewPeriodicManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	periodicdataimageManager := models.NewPeriodicdataimageManager(conn)
	periodicotherManager := models.NewPeriodicotherManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)

	var id int64 = 1274
	periodicItem := periodicManager.Get(id)

	blueprints := blueprintManager.FindByApt(periodicItem.Apt)
	for _, v := range blueprints {
		if v.Filename == "" {
			continue
		}

		url := fmt.Sprintf("%v/webdata/%v", serverUrl, v.Filename)
		fullFilename := path.Join(config.UploadPath, v.Filename)
		global.DownloadImage(url, fullFilename)
	}

	{
		items := periodicdataManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
		})

		for _, v := range items {
			if v.Filename == "" {
				continue
			}

			url := fmt.Sprintf("%v/webdata/%v", serverUrl, v.Filename)
			fullFilename := path.Join(config.UploadPath, v.Filename)
			global.DownloadImage(url, fullFilename)
		}
	}

	{
		items := periodicdataimageManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
		})

		for _, v := range items {
			if v.Filename == "" {
				continue
			}

			url := fmt.Sprintf("%v/webdata/%v", serverUrl, v.Filename)
			fullFilename := path.Join(config.UploadPath, v.Filename)
			global.DownloadImage(url, fullFilename)
		}
	}

	{
		items := periodicotherManager.Find([]interface{}{
			models.Where{Column: "periodic", Value: id, Compare: "="},
		})

		for _, v := range items {
			if v.Filename == "" {
				continue
			}

			url := fmt.Sprintf("%v/webdata/%v", serverUrl, v.Filename)
			fullFilename := path.Join(config.UploadPath, v.Filename)
			global.DownloadImage(url, fullFilename)
		}
	}
}

// func TestDownloadImage(t *testing.T) {
// 	log.Init()

// 	conn := models.NewConnection()
// 	defer conn.Close()

// 	aptManager := models.NewAptManager(conn)
// 	estimateManager := models.NewEstimateManager(conn)
// 	contractManager := models.NewContractManager(conn)

// 	apts := aptManager.FindAll()

// 	for _, apt := range apts {
// 		contracts := contractManager.Find([]interface{}{
// 			models.Where{Column: "apt", Value: apt.Id, Compare: "="},
// 		})

// 		for _, contract := range contracts {
// 			estimates := estimateManager.Find([]interface{}{
// 				models.Where{Column: "apt", Value: contract.Apt, Compare: "="},
// 			})

// 			for _, estimate := range estimates {
// 				if contract.Apt != estimate.Apt {
// 					log.Println("NOT MATCH #################################################################")
// 					log.Println("ID", contract.Id)
// 				}
// 			}
// 		}
// 	}
// }

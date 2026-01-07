package services

import (
	"anb/global"
	"anb/models"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/CloudyKit/jet"
)

type Image struct {
	Filename       string
	Number         int64
	Position       string
	Content        string
	Name           string
	StatusFilename string
}

type ImagePage struct {
	Items []Image
}

func isNumber(num int) bool {
	if num == 20 || num == 21 {
		return true
	} else {
		return false
	}
}

func GetTitle(id int64) string {
	conn := models.NewConnection()
	defer conn.Close()

	title := ""

	manager := models.NewImageManager(conn)

	item := manager.Get(id)

	title = item.Name

	for {
		if item.Parent == 0 {
			break
		}

		item = manager.Get(item.Parent)

		if item == nil {
			break
		}

		title = item.Name + " " + title
	}

	return title
}

func GetTitleFloor(id int64) string {
	conn := models.NewConnection()
	defer conn.Close()

	title := ""

	manager := models.NewImageManager(conn)

	item := manager.Get(id)

	title = item.Name

	for {
		if item.Parent == 0 {
			break
		}

		item = manager.Get(item.Parent)

		if item == nil {
			break
		}

		title = item.Name + " " + title
	}

	return title
}

const (
	TypeTitle   = 1
	TypeContent = 2
	TypeFooter  = 3

	MaxTableRows = 18
)

type Item struct {
	Type  int
	Title string
	Data  models.Data
}

type Table struct {
	Items []Item
}

func GetHead(id int64, images []Image, total int) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	v := make(jet.VarMap)

	v.Set("total", total)
	v.Set("images", images)

	t, err := view.GetTemplate("head.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetTail(id int64) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	v := make(jet.VarMap)

	t, err := view.GetTemplate("tail.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetStrength(items []Image) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	v := make(jet.VarMap)

	v.Set("items", items)

	t, err := view.GetTemplate("strength.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetMaterial(items []Image) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	v := make(jet.VarMap)

	v.Set("items", items)

	t, err := view.GetTemplate("material.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetSlope(items []Image) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	v := make(jet.VarMap)

	v.Set("items", items)

	t, err := view.GetTemplate("slope.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetBody(id int64, faults []Image, strengths []Image, materials []Image, slopes []Image) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	v := make(jet.VarMap)

	v.Set("strength", GetStrength(strengths))
	v.Set("material", GetMaterial(materials))
	v.Set("slope", GetSlope(slopes))

	t, err := view.GetTemplate("body.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetDraw(id int64, allImages []Image, faults []Image, strengths []Image, materials []Image, slopes []Image) string {
	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	v := make(jet.VarMap)

	v.Set("total", len(allImages))
	v.Set("images", allImages)

	v.Set("strengths", strengths)
	v.Set("materials", materials)
	v.Set("slopes", slopes)

	t, err := view.GetTemplate("draw.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetFault(id int64, images []Image) string {
	conn := models.NewConnection()
	defer conn.Close()

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(id)

	aptgroupManager := models.NewAptgroupManager(conn)
	aptgroup := aptgroupManager.Get(apt.Aptgroup)

	if aptgroup == nil {
		return ""
	}

	title := aptgroup.Name + " " + apt.Name

	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	view.AddGlobal("imagename", func(id int64) string {
		return GetTitle(id)
	})

	v := make(jet.VarMap)

	v.Set("groupname", aptgroup.Name)
	v.Set("aptname", apt.Name)
	v.Set("name", title)

	v.Set("total", len(images)+6)
	v.Set("images", images)

	t, err := view.GetTemplate("fault.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetFloor(id int64, images []Image) string {
	items := make([]ImagePage, 0)

	total := len(images)

	remain := len(images) % 4

	pages := len(images) / 4

	index := 0
	for i := 0; i < pages; i++ {
		var p ImagePage
		p.Items = make([]Image, 0)

		for j := 0; j < 4; j++ {
			p.Items = append(p.Items, images[index])
			index++
		}

		items = append(items, p)
	}

	for i := 0; i < 1; i++ {
		var p ImagePage
		p.Items = make([]Image, 0)

		for j := 0; j < remain; j++ {
			p.Items = append(p.Items, images[index])
			index++
		}

		items = append(items, p)
	}

	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	v := make(jet.VarMap)

	v.Set("total", total)
	v.Set("images", images)

	v.Set("pages", items)

	t, err := view.GetTemplate("floor.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetTable(id int64) string {
	conn := models.NewConnection()
	defer conn.Close()

	hwps := make([]Table, 0)
	tables := make([]Item, 0)

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

			if items != nil {
				checkCnt := 0
				for _, item := range *items {
					if !isNumber(item.Type) {
						continue
					}

					checkCnt++
				}

				if checkCnt == 0 {
					continue
				}

				tables = append(tables, Item{Type: TypeTitle, Title: GetTitle(image.Id)})
				if len(tables) == MaxTableRows {
					tables = append(tables, Item{Type: TypeFooter})
					hwps = append(hwps, Table{Items: tables})
					tables = make([]Item, 0)
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

					tables = append(tables, Item{Type: TypeContent, Data: item})
					if len(tables) == MaxTableRows {
						tables = append(tables, Item{Type: TypeFooter})
						hwps = append(hwps, Table{Items: tables})
						tables = make([]Item, 0)
					}
				}
			}
		}

	}

	if len(tables) > 0 {
		tables = append(tables, Item{Type: TypeFooter})
		hwps = append(hwps, Table{Items: tables})
		tables = make([]Item, 0)
	}

	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("tableRows", func(row []Item) int {
		return len(row) + 2
	})

	view.AddGlobal("isMultiRows", func(str string) bool {
		return strings.Contains(str, "/")
	})

	view.AddGlobal("first", func(str string) string {
		strs := strings.Split(str, "/")
		return strs[0]
	})

	view.AddGlobal("second", func(str string) string {
		strs := strings.Split(str, "/")

		if len(strs) == 3 {
			return strs[1] + "/" + strs[2]
		} else {
			return strs[1]
		}
	})

	view.AddGlobal("floatFormat", func(val float64) string {
		return fmt.Sprintf("%.1f", val)
	})

	view.AddGlobal("zero", func(val float64) string {
		if val == 0.0 {
			return "-"
		} else {
			return fmt.Sprintf("%.1f", val)
		}
	})

	v := make(jet.VarMap)

	rows := 0
	rows = len(hwps) + 2

	v.Set("rows", rows)
	v.Set("tables", hwps)
	t, err := view.GetTemplate("table.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return b.String()
}

func GetImage(id int64) ([]Image, string) {
	conn := models.NewConnection()
	defer conn.Close()

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(id)

	imageManager := models.NewImageManager(conn)
	imageItems := imageManager.GetListByApt(id, 0, 0, "order, i_id")

	if imageItems == nil {
		return nil, ""
	}

	dataManager := models.NewDataManager(conn)

	images := make([]Image, 0)

	pos := 0
	for _, image := range *imageItems {
		datas := dataManager.GetListByAptImage(image.Apt, image.Id, 0, 0, "id")

		if datas == nil {
			continue
		}

		length := len(*datas)

		for i := 0; i < length; i++ {
			item := (*datas)[i]

			if item.Filename == "" {
				continue
			}

			if _, err := os.Stat(item.Filename); err == nil {
			} else if os.IsNotExist(err) {
				continue
			}

			fullFilename := item.Filename
			targetFilename := strings.Replace(item.Filename, "webdata/", "webdata/_", -1)
			log.Println(targetFilename)
			global.MakeThumbnailPicture(3624, 2448, fullFilename, targetFilename)

			item.Filename = targetFilename

			name := item.Name
			if item.Imagetype == 8 {
				name = item.Name + " " + strings.Replace(strings.Replace(item.Imagename, "좌 ", "", -1), "우 ", "", -1)
			} else {
				name = item.Fault + " " + item.Name
			}

			img := Image{Filename: item.Filename, Number: int64(pos) + 1, Position: fmt.Sprintf("%v %v", GetTitle(image.Id), name), Content: item.Content}
			images = append(images, img)

			pos++
		}

	}

	total := len(images)

	remain := len(images) % 6

	if remain > 0 {
		for i := 0; i < 6-remain; i++ {

			img := Image{Filename: "", Number: 0, Position: "", Content: ""}
			images = append(images, img)
		}
	}

	pages := len(images) / 6

	items := make([]ImagePage, 0)

	index := 0
	for i := 0; i < pages; i++ {
		var p ImagePage
		p.Items = make([]Image, 0)

		for j := 0; j < 6; j++ {
			p.Items = append(p.Items, images[index])
			index++
		}

		items = append(items, p)
	}

	var b bytes.Buffer

	var view = jet.NewHTMLSet("./template")
	view.SetDevelopmentMode(true)

	view.AddGlobal("basename", func(str string) string {
		return strings.Replace(str, "webdata/", "", 1)
	})

	view.AddGlobal("zero", func(value int) string {
		if value == 0 {
			return ""
		} else {
			return fmt.Sprintf("%v", value)
		}
	})

	v := make(jet.VarMap)

	v.Set("total", total)
	v.Set("images", images)
	v.Set("pages", items)

	title := ""
	if apt.Type == 1 {
		title = "정밀안전점검"
	} else if apt.Type == 2 {
		title = "정기안점점검"
	} else if apt.Type == 3 {
		title = "하자조사"
	} else if apt.Type == 4 {
		title = "장기수선"
	}

	v.Set("type", title)
	t, err := view.GetTemplate("image.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println(err)
	}

	return images, b.String()
}

func MakeHwp(id int64) {
	conn := models.NewConnection()
	defer conn.Close()

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(id)

	//aptgroupManager := models.NewAptgroupManager(conn)
	//aptgroup := aptgroupManager.Get(apt.Aptgroup)

	imageManager := models.NewImageManager(conn)
	images := imageManager.GetListByApt(id, 0, 0, "order, i_id")

	imagefloorManager := models.NewImagefloorManager(conn)
	manager := models.NewDataManager(conn)
	statusManager := models.NewStatusManager(conn)

	//zipFilename := fmt.Sprintf("%v-%04d%02d%02d.zip", apt.Id, t.Year(), t.Month(), t.Day())
	zipFilename := fmt.Sprintf("%v.zip", apt.Id)
	//filename := fmt.Sprintf("%v %v.zip", aptgroup.Name, apt.Name)

	fnames := make([]string, 0)
	filenames := make([]string, 0)

	faults := make([]Image, 0)
	strengths := make([]Image, 0)
	materials := make([]Image, 0)
	slopes := make([]Image, 0)

	allImages := make([]Image, 0)

	floors := make([]Image, 0)

	imageNo := 0

	if images != nil {
		statuss := statusManager.GetListByCompany(apt.Company, 0, 0, "")

		if statuss != nil {
			for _, item := range *statuss {
				if item.Type != 8 {
					continue
				}

				if item.Content == "" {
					continue
				}

				_, err := os.Stat(item.Content)

				if !os.IsNotExist(err) {
					fnames = append(fnames, item.Content)
					filenames = append(filenames, item.Content)
				}
			}
		}

		str2 := ""

		str2 += "		<table width=\"84%\" border=1 cellpadding=5 cellspacing=0>"
		str2 += "	<tr>"
		str2 += "		<td width=\"10%\" align=\"center\" valign=\"middle\">층별</td>"
		str2 += "		<td width=\"10%\" align=\"center\" valign=\"middle\">부위</td>"
		str2 += "		<td width=\"14%\" align=\"center\" valign=\"middle\">부재</td>"
		str2 += "		<td width=\"40%\" align=\"center\" valign=\"middle\">결함도</td>"
		str2 += "		<td width=\"26%\" align=\"center\" valign=\"middle\">결함 및 주변현황<br>(폭 mm x 길이 m)</td>"
		str2 += "	</tr>"

		floorFlag := false

		for _, image := range *images {
			if image.Type == 9 {

				filename := fmt.Sprintf("webdata/%v-%v.png", id, image.Id)
				_, err := os.Stat(filename)

				log.Println("filename = ", filename)
				if !os.IsNotExist(err) {
					target := imagefloorManager.GetByTarget(image.Id)

					if target != nil {
						//parent := imageManager.Get(target.Image)
						datas := manager.GetListByImageNameImagename(target.Image, target.Name, target.Imagename, 0, 0, "")

						status := statusManager.GetByCompanyTypeName(apt.Company, 8, target.Imagename)

						if datas != nil && status != nil {
							floorFlag = true

							content := ""
							name := ""

							for j, data := range *datas {
								name = data.Imagename
								if j > 0 {
									content += " "
								}

								content += data.Remark
								content += " "
								if data.Width > 0 && data.Length > 0 {
									content += fmt.Sprintf("%.1f %.1f", data.Width, data.Length)

									if data.Count != "" && data.Count != "0" {
										content += fmt.Sprintf(" X %v", data.Count)
									}
								}
							}

							fnames = append(fnames, status.Content)
							filenames = append(filenames, status.Content)

							fnames = append(fnames, filename)
							filenames = append(filenames, filename)

							str2 += fmt.Sprintf("<tr>\n")
							str2 += fmt.Sprintf("<td align=\"center\" valign=\"middle\">%v</td>\n", GetTitleFloor(image.Id))
							str2 += fmt.Sprintf("<td align=\"center\" valign=\"middle\">%v</td>\n", name)
							str2 += fmt.Sprintf("<td align=\"center\" valign=\"middle\"><img src=\"%v\" width=\"80%%\" /></td>\n", status.Content)
							str2 += fmt.Sprintf("<td align=\"center\" valign=\"middle\"><img src=\"%v\" width=\"80%%\" /></td>\n", filename)
							str2 += fmt.Sprintf("<td align=\"center\" valign=\"middle\">%v</td>\n", content)
							str2 += fmt.Sprintf("</tr>\n")

							log.Println("---------------------------------------------")
							img := Image{Name: name, Filename: filename, Number: 1, Position: GetTitleFloor(image.Id) + " " + "계단실", Content: content, StatusFilename: status.Content}
							floors = append(floors, img)
						}
					}

				}
			}

			if image.Type == 1 || image.Type == 2 || image.Type == 3 || image.Type == 4 {

				filename := fmt.Sprintf("webdata/%v-%v.png", id, image.Id)
				_, err := os.Stat(filename)

				if os.IsNotExist(err) {
					filename = image.Filename
				}

				if image.Type == 1 {
					fullFilename := filename
					targetFilename := strings.Replace(filename, "webdata/", "webdata/_", -1)
					log.Println(targetFilename)
					global.MakeThumbnail(1600, 1193, fullFilename, targetFilename)

					filename = targetFilename
				}

				fnames = append(fnames, filename)
				filenames = append(filenames, filename)

				if image.Type != 1 {
					imageNo++
				}

				if image.Type == 1 {
					img := Image{Filename: filename, Number: image.Id, Position: GetTitle(image.Id), Content: ""}
					faults = append(faults, img)
				} else if image.Type == 2 {
					img := Image{Filename: filename, Number: int64(imageNo), Position: GetTitle(image.Id), Content: ""}
					slopes = append(slopes, img)
					allImages = append(allImages, img)
				} else if image.Type == 3 {
					img := Image{Filename: filename, Number: int64(imageNo), Position: GetTitle(image.Id), Content: ""}
					strengths = append(strengths, img)
					allImages = append(allImages, img)
				} else if image.Type == 4 {
					img := Image{Filename: filename, Number: int64(imageNo), Position: GetTitle(image.Id), Content: ""}
					materials = append(materials, img)
					allImages = append(allImages, img)
				}
			}
		}

		if floorFlag == true {
			/*
				str2 += fmt.Sprintf("</table>")
				ioutil.WriteFile("webdata/floor.html", []byte(str2), 0644)

				fnames = append(fnames, "계단실.html")
				filenames = append(filenames, "webdata/floor.html")
			*/

			floor := GetFloor(id, floors)

			if floor != "" {
				ioutil.WriteFile(fmt.Sprintf("webdata/%v-floor.hml", id), []byte(floor), 0644)

				fnames = append(fnames, "계단실.hml")
				filenames = append(filenames, fmt.Sprintf("webdata/%v-floor.hml", id))
			}

		}

	}

	/*
		if flag1 == true {
			ioutil.WriteFile("webdata/1.html", []byte(html1), 0644)

			fnames = append(fnames, "결함도.html")
			filenames = append(filenames, "webdata/1.html")
		}

		if flag2 == true {
			ioutil.WriteFile("webdata/2.html", []byte(html2), 0644)

			fnames = append(fnames, "기울기 위치도.html")
			filenames = append(filenames, "webdata/2.html")
		}

		if flag3 == true {
			ioutil.WriteFile("webdata/3.html", []byte(html3), 0644)

			fnames = append(fnames, "강도 및 탄산화 위치도.html")
			filenames = append(filenames, "webdata/3.html")
		}

		if flag4 == true {
			ioutil.WriteFile("webdata/4.html", []byte(html4), 0644)

			fnames = append(fnames, "부재.html")
			filenames = append(filenames, "webdata/4.html")
		}
	*/

	if len(strengths)%2 == 1 {
		img := Image{Filename: "", Number: 0, Position: "", Content: ""}
		strengths = append(strengths, img)
	}

	if len(materials)%2 == 1 {
		img := Image{Filename: "", Number: 0, Position: "", Content: ""}
		materials = append(materials, img)
	}

	/*
		head := GetHead(id, allImages, imageNo)
		body := GetBody(id, faults, strengths, materials, slopes)
		tail := GetTail(id)

		ioutil.WriteFile(fmt.Sprintf("webdata/%v-report.hml", id), []byte(head+body+tail), 0644)

		fnames = append(fnames, "보고서 양식.hml")
		filenames = append(filenames, fmt.Sprintf("webdata/%v-report.hml", id))
	*/

	fault := GetFault(id, faults)

	ioutil.WriteFile(fmt.Sprintf("webdata/%v-fault.hml", id), []byte(fault), 0644)

	fnames = append(fnames, "결함도.hml")
	filenames = append(filenames, fmt.Sprintf("webdata/%v-fault.hml", id))

	draw := GetDraw(id, allImages, faults, strengths, materials, slopes)

	ioutil.WriteFile(fmt.Sprintf("webdata/%v-draw.hml", id), []byte(draw), 0644)

	fnames = append(fnames, "도면.hml")
	filenames = append(filenames, fmt.Sprintf("webdata/%v-draw.hml", id))

	table := GetTable(id)

	if table != "" {
		ioutil.WriteFile(fmt.Sprintf("webdata/%v-table.hml", id), []byte(table), 0644)

		fnames = append(fnames, "육안검사.hml")
		filenames = append(filenames, fmt.Sprintf("webdata/%v-table.hml", id))
	}

	imgs, str := GetImage(id)

	if imgs != nil {
		for _, item := range imgs {
			if item.Number > 0 {
				fnames = append(fnames, item.Filename)
				filenames = append(filenames, item.Filename)
			}
		}

		ioutil.WriteFile(fmt.Sprintf("webdata/%v-picture.hml", id), []byte(str), 0644)

		fnames = append(fnames, "사진자료.hml")
		filenames = append(filenames, fmt.Sprintf("webdata/%v-picture.hml", id))
	}

	if err := os.Remove("webdata/" + zipFilename); err != nil {
		fmt.Println(err)
	}

	global.MakeZipfile(zipFilename, fnames, filenames)
}

func Report() {
	for {
		time.Sleep(5 * time.Second)

		conn := models.NewConnection()

		aptManager := models.NewAptManager(conn)

		manager := models.NewReportManager(conn)
		items := manager.GetListByStatus(2, 0, 0, "")

		if items != nil {
			for _, item := range *items {
				item.Status = 3
				manager.Update(&item)

				aptManager.UpdateReportById(3, item.Apt)

				MakeHwp(item.Apt)

				item.Status = 4
				manager.Update(&item)

				aptManager.UpdateReportById(4, item.Apt)
			}
		}

		conn.Close()
	}
}

type TopsData struct {
	Id    int64
	Name  string
	Items []models.Image
}

type ResultItem struct {
	Name  string
	Fault string
	Width float64
}

type ResultData struct {
	Type    int
	Typestr string
	Items   map[string][]ResultItem
}

type ReportItem struct {
	Name   string
	Fault  []string
	Method []string
}

type ReportData struct {
	Name  string
	Items []ReportItem
}

func MakeSummary() {
	conn := models.NewConnection()
	defer conn.Close()

	//var id int64 = 833
	var id int64 = 906
	imageManager := models.NewImageManager(conn)
	images := imageManager.GetListByApt(id, 0, 0, "order, i_id")

	imageMap := make(map[int64]models.Image)
	tops := make([]TopsData, 0)
	levelType := 1

	pos := 0

	for _, v := range *images {
		//log.Println(v.Level, v.Name)
		if v.Level == 0 {
			var top TopsData
			top.Id = v.Id
			top.Name = v.Name
			top.Items = make([]models.Image, 0)
			tops = append(tops, top)
		} else {
			levelType = 2

			tops[pos].Items = append(tops[pos].Items, v)
		}

		imageMap[v.Id] = v

	}

	log.Println("====================")
	log.Println(levelType)
	log.Println("====================")

	manager := models.NewDataManager(conn)
	items := manager.GetListByApt(id, 0, 0, "")

	result := make([]ReportData, 0)
	for _, top := range tops {
		lists := make([]ResultData, 0)
		topFlag1 := false
		topFlag2 := false
		for _, image := range top.Items {
			baseTypestr := "지상층"
			if strings.Contains(image.Name, "옥탑") {
				baseTypestr = "지붕층"
				topFlag1 = true
			} else if strings.Contains(image.Name, "지붕") {
				baseTypestr = "지붕층"
				topFlag2 = true
			} else if strings.Contains(image.Name, "주차장") {
				baseTypestr = "주차장"
			} else if strings.Contains(image.Name, "지하") {
				baseTypestr = "지하층"
			} else {
			}

			for _, v := range *items {
				if image.Id != v.Image {
					continue
				}

				if v.Type != 20 && v.Type != 21 {
					continue
				}

				typestr := baseTypestr

				if v.Fault == "계단실" {
					typestr = "계단실"
				} else if v.Fault == "주차장" {
					typestr = "주차장"
				}

				key := typestr
				//log.Println(v.Content, v.Width)
				//key := fmt.Sprintf("%v %v %v", typestr, v.Fault, v.Name)
				//
				var target *ResultData

				flag := false
				for _, list := range lists {
					if list.Typestr == key {
						flag = true
						target = &list
						break
					}
				}

				if flag == false {
					var list ResultData
					list.Typestr = typestr
					list.Items = make(map[string][]ResultItem)
					lists = append(lists, list)

					target = &list
				}

				var ritem ResultItem
				ritem.Name = v.Content
				ritem.Width = v.Width
				ritem.Fault = v.Fault

				flag = false
				for k, _ := range target.Items {
					if k == v.Name {
						flag = true
						break
					}
				}

				if flag == false {
					target.Items[v.Name] = make([]ResultItem, 0)
				}

				target.Items[v.Name] = append(target.Items[v.Name], ritem)

				/*
					log.Println(v.Name)
					log.Println(v.Fault)
					log.Println(v.Content)
				*/
			}
		}

		if len(lists) == 0 {
			continue
		}

		log.Println(top.Name)
		var reportData ReportData
		reportData.Name = top.Name
		reportData.Items = make([]ReportItem, 0)

		for _, v := range lists {
			if v.Typestr == "지붕층" {
				if topFlag2 == false {
					v.Typestr = "옥탑"
				} else {
					if topFlag1 == true {
						v.Typestr = "지붕층 및 옥탑"
					}
				}
			}

			log.Println("********************************")
			//log.Println(v.Typestr)

			for k2, v2 := range v.Items {
				//log.Println(k2)

				width2 := false
				width3 := false

				crack1 := false
				crack2 := false

				leak := true
				leak1 := false
				leak2 := false
				leak3 := false

				items := make([]string, 0)
				faults := make([]string, 0)

				for _, v3 := range v2 {
					fflag := false
					for _, fault := range faults {
						if v3.Fault == fault {
							fflag = true
							break
						}
					}

					if fflag == false {
						faults = append(faults, v3.Fault)
					}

					strs := strings.Split(v3.Name, "/")

					if strings.Contains(v3.Name, "균열") {
						if v3.Width >= 0.3 {
							width3 = true
						} else {
							width2 = true
						}
					}

					if strings.Contains(v3.Name, "조적") {
						crack1 = true
						continue
					}

					if strings.Contains(v3.Name, "이질재") {
						crack2 = true
						continue
					}

					if strings.Contains(v3.Name, "균열/누수") {
						leak = true
					}

					items = append(items, strs...)
				}

				titles := make([]string, 0)
				for _, name := range items {
					if name == "누수" {
						leak1 = true
						continue
					} else if name == "누수흔적" {
						leak2 = true
						continue
					} else if name == "백태" {
						leak3 = true
						continue
					}

					flag := false
					for i, v4 := range titles {
						if v4 == name {
							flag = true
							break
						}

						if len(v4) > len(name) {
							if v4[:len(name)] == name {
								flag = true
								break
							}
						} else {
							if name[:len(v4)] == v4 {
								titles[i] = name
								flag = true
								break
							}
						}
					}

					if flag == false {
						titles = append(titles, name)
					}
				}

				var reportItem ReportItem
				reportItem.Fault = make([]string, 0)
				reportItem.Method = make([]string, 0)

				if v.Typestr == "주차장" || v.Typestr == "계단실" {
					reportItem.Name = fmt.Sprintf("%v %v", v.Typestr, k2)
				} else {
					if len(faults) == 1 {
						reportItem.Name = fmt.Sprintf("%v %v %v", v.Typestr, strings.Join(faults, " / "), k2)
					} else {
						reportItem.Name = fmt.Sprintf("%v (%v) %v", v.Typestr, strings.Join(faults, "/"), k2)
					}
				}

				rtitles := make([]string, 0)
				ctitles := make([]string, 0)
				ititles := make([]string, 0)

				crack := false
				for _, title := range titles {
					if title == "균열" {
						crack = true
						continue
					}

					if strings.Contains(title, "균열") {
						rtitles = append(rtitles, strings.ReplaceAll(title, "균열", ""))
					}

					if strings.Contains(title, "콘크리트") {
						ctitles = append(ctitles, title)
					}

					if strings.Contains(title, "철근") {
						ititles = append(ititles, title)
					}
				}

				if len(rtitles) > 0 {
					reportItem.Fault = append(reportItem.Fault, fmt.Sprintf("%v균열", strings.Join(rtitles, "/")))
				}

				if leak1 == false && leak2 == false && leak3 == false && crack == true {
					reportItem.Fault = append(reportItem.Fault, "균열")
				}

				if width2 == true {
					reportItem.Method = append(reportItem.Method, "건식주입 공법(0.3mm 이상 균열)")
				}

				if width3 == true {
					reportItem.Method = append(reportItem.Method, "표면처리 공법(0.3mm 미만 균열)")
				}

				if crack1 == true && crack2 == true {
					reportItem.Fault = append(reportItem.Fault, "조적/이질재균열")
					reportItem.Method = append(reportItem.Method, "충전식 공법")
				} else if crack1 == true {
					reportItem.Fault = append(reportItem.Fault, "조적균열")
					reportItem.Method = append(reportItem.Method, "충전식 공법")
				} else if crack2 == true {
					reportItem.Fault = append(reportItem.Fault, "이질재균열")
					reportItem.Method = append(reportItem.Method, "충전식 공법")
				}

				leakTitles := make([]string, 0)
				if leak == true {
					leakTitles = append(leakTitles, "균열")
				}

				if leak1 == true {
					leakTitles = append(leakTitles, "누수")
				} else if leak2 == true {
					leakTitles = append(leakTitles, "누수흔적")
				}

				if leak3 == true {
					leakTitles = append(leakTitles, "백태")
				}

				if len(leakTitles) > 0 {
					reportItem.Fault = append(reportItem.Fault, strings.Join(leakTitles, "/"))
					reportItem.Method = append(reportItem.Method, "습식주입 공법")
				}

				if len(ctitles) > 0 {
					reportItem.Fault = append(reportItem.Fault, strings.Join(ctitles, "/"))
					reportItem.Method = append(reportItem.Method, "단면복구 공법")
				}

				if len(ititles) > 0 {
					reportItem.Fault = append(reportItem.Fault, strings.Join(ititles, "/"))
					reportItem.Method = append(reportItem.Method, "철근노출 보수공법")
				}

				reportData.Items = append(reportData.Items, reportItem)
			}

			log.Println("====================")
		}

		result = append(result, reportData)
	}

	log.Println(result)
}

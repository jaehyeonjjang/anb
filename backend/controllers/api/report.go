package api

import (
	"anb/controllers"
	"anb/global"
	"anb/models"
	"fmt"
	"log"
	"strings"
)

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

type ReportController struct {
	controllers.Controller
}

func (c *ReportController) Index() {
	conn := c.NewConnection()

	search := c.Get("search")

	page := c.DefaultGeti("page", 1)
	pagesize := c.DefaultGeti("pagesize", 20)

	manager := models.NewAptManager(conn)
	items := manager.GetListByName(search, page, pagesize, "")
	count := manager.GetCountByName(search)

	c.Set("items", items)

	c.Paging(page, count, pagesize)

	c.Set("statuss", global.AptStatuss)
}

func (c *ReportController) DownloadHwp() {
}

func (c *ReportController) AjaxSummary() {
	conn := c.NewConnection()

	id := c.Geti64("id")

	//var id int64 = 833
	//var id int64 = 906

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(id)

	imageManager := models.NewImageManager(conn)
	images := imageManager.GetListByApt(id, 0, 0, "order, i_id")

	tops := make([]TopsData, 0)
	levelType := 1

	pos := -1

	for _, v := range *images {
		//log.Println(v.Level, v.Name)
		if v.Level == 0 {
			var top TopsData
			top.Id = v.Id
			top.Name = v.Name
			top.Items = make([]models.Image, 0)
			tops = append(tops, top)

			pos++
		} else {
			levelType = 2

			tops[pos].Items = append(tops[pos].Items, v)
		}
	}

	log.Println("type", levelType)

	if levelType == 1 || apt.Summarytype == 2 {
		pos = 0
		tops = make([]TopsData, 0)
		var top TopsData
		top.Id = 0
		top.Name = "요약표"
		top.Items = make([]models.Image, 0)
		tops = append(tops, top)

		for _, v := range *images {
			tops[pos].Items = append(tops[pos].Items, v)
		}
	}

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

				log.Println(v.Name)

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
			}
		}

		if len(lists) == 0 {
			continue
		}

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

				v.Typestr = "지붕(옥탑)층"
			}

			for k2, v2 := range v.Items {
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
						log.Println(v3.Name)
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

				if leak1 == false && leak2 == false && leak3 == false && leak == false && crack == true {
					log.Println("insert 1")
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
		}

		result = append(result, reportData)
	}

	log.Println(result)
	c.Set("items", result)
}

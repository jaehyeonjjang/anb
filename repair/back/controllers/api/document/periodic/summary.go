package periodic

import (
	"fmt"
	"log"
	"repair/global"
	"repair/models"
	"strings"
)

type SummaryResultItem struct {
	Part     string
	Result   string
	Methods  []string
	Causes   []string
	MaxWidth float64
	MinWidth float64
	Height   int
	Type     int
}

type SummaryResultType struct {
	Type  int
	Count int
	Items []SummaryResultItem
}

type SummaryResultList struct {
	Type         int
	Originaltype int
	Count        int
	Height       int
	Item         SummaryResultItem
	Last         int
}

type CheckResult struct {
	Crack         []string
	Leak          []string
	Concret       []string
	Material      []string
	CrackCount    int
	LeakCount     int
	ConcretCount  int
	MaterialCount int
}

type SummaryResult struct {
	Name         string
	Id           int64
	Count        int
	Height       int
	Types        []SummaryResultType
	Items        []SummaryResultList
	Allcount     int
	Check        CheckResult
	Result       []string
	ResultBottom string
}

type SummaryDong struct {
	Name     string
	Id       int64
	Private  int
	Topcount int
}

type SummaryCheck struct {
	Dong   int64
	Crack  []string
	Leak   []string
	Rebar  []string
	Steel  []string
	Finish []string
}

type DataSummary struct {
	Data []models.Periodicdata
}

func GetSummary(id int64, periodic models.Periodic, dongs []models.Aptdong, items []models.Periodicdata, blueprints []models.Blueprint, periodicchanges []models.Periodicchange, others map[int]OtherResult, otherMap map[int][]models.Periodicother) []SummaryResult {
	var privateDong models.Aptdong
	for _, v := range dongs {
		if v.Private == 3 {
			privateDong = v
		}
	}

	summaryResults := make([]SummaryResult, 0)

	summaryDongs := make([]SummaryDong, 0)

	publicTopcount := 0

	publicFlag := false
	commonFlag := false
	var commonDong models.Aptdong

	for _, dong := range dongs {
		if dong.Private == 2 {
			publicFlag = true
			publicTopcount = dong.Topcount
			continue
		}

		if dong.Private == 4 {
			commonFlag = true
			commonDong = dong
			continue
		}

		if privateDong.Id == 0 {
			summaryDongs = append(summaryDongs, SummaryDong{Name: dong.Dong, Id: dong.Id, Private: 1, Topcount: dong.Topcount})
		}
	}

	if privateDong.Id > 0 {
		summaryDongs = append(summaryDongs, SummaryDong{Name: "", Id: privateDong.Id, Private: 1, Topcount: privateDong.Topcount})
	}

	if publicFlag == true {
		summaryDongs = append(summaryDongs, SummaryDong{Name: "공용부위", Id: 0, Private: 2, Topcount: publicTopcount})
	}

	floorCount := 7
	for _, dong := range summaryDongs {
		topcount := dong.Topcount
		summarys := make([]DataSummary, floorCount)
		for i := 0; i < floorCount; i++ {
			summarys[i].Data = make([]models.Periodicdata, 0)
		}

		summaryResult := SummaryResult{Name: dong.Name, Id: dong.Id, Types: make([]SummaryResultType, 0), Items: make([]SummaryResultList, 0)}

		for _, blueprint := range blueprints {
			if dong.Private == 1 {
				if privateDong.Id == 0 {
					if blueprint.Aptdong != dong.Id {
						continue
					}
				}
			} else if dong.Private == 3 {
				if blueprint.Aptdong != dong.Id {
					continue
				}
			} else {
				flag := false
				for i := 0; i < len(summaryDongs)-1; i++ {
					if blueprint.Aptdong == summaryDongs[i].Id {
						flag = true
					}
				}

				if flag == true {
					continue
				}
			}

			for _, v := range items {
				if blueprint.Id != v.Blueprint {
					continue
				}

				if v.Type != 1 && v.Type != 2 {
					continue
				}

				position := 3
				if blueprint.Floortype == 4 || blueprint.Floortype == 5 {
					position = 1
				} else if blueprint.Floortype == 1 || blueprint.Floortype == 2 {
					position = 5
				}

				if v.Type == 2 {
					position++
				}

				if v.Part == "외부" {
					position = 7
				}

				summarys[position-1].Data = append(summarys[position-1].Data, v)
			}

			if dong.Private == 1 && commonFlag == true {
				for _, v := range items {
					blueprint := v.Extra["blueprint"].(models.Blueprint)

					findBlueprint := false
					for _, v2 := range blueprints {
						if v2.Aptdong == commonDong.Id {
							if v2.Id == blueprint.Id {
								findBlueprint = true
								break
							}
						}
					}

					if findBlueprint == false {
						continue
					}

					if v.Type != 1 && v.Type != 2 {
						continue
					}

					position := 3
					if blueprint.Floortype == 4 || blueprint.Floortype == 5 {
						position = 1
					} else if blueprint.Floortype == 1 || blueprint.Floortype == 2 {
						position = 5
					}

					if v.Type == 2 {
						position++
					}

					if v.Part == "외부" {
						position = 7
					}

					summarys[position-1].Data = append(summarys[position-1].Data, v)
				}
			}
		}

		for i := 0; i < floorCount; i++ {
			width2 := false
			width3 := false
			maxWidth := 0.0
			minWidth := 9999.0

			crack1 := false
			crack2 := false
			crack3 := false
			crack4 := false
			crack5 := false

			othercrack1 := false
			othercrack2 := false

			leak1 := false
			leak2 := false
			leak3 := false

			concrete1 := false
			concrete2 := false
			concrete3 := false

			rebar1 := false
			rebar2 := false

			material1 := false
			material2 := false

			parts := make([][]string, floorCount)
			causes := make([][]string, floorCount)
			resultItems := make([]SummaryResultItem, 0)

			for j := 0; j < floorCount; j++ {
				parts[j] = make([]string, 0)
				causes[j] = make([]string, 0)
			}

			for _, data := range summarys[i].Data {
				check := []bool{false, false, false, false, false, false}

				if strings.Contains(data.Shape, "조적") {
					othercrack1 = true
					check[1] = true
				}

				if strings.Contains(data.Shape, "이질재") {
					othercrack2 = true
					check[1] = true
				}

				if strings.Contains(data.Shape, "마감재") {
					if strings.Contains(data.Shape, "오염") {
						material1 = true
					}

					if strings.Contains(data.Shape, "파손") {
						material2 = true
					}

					check[5] = true
				}

				if strings.Contains(data.Shape, "콘크리트") {
					if strings.Contains(data.Shape, "박리") {
						concrete1 = true
					}

					if strings.Contains(data.Shape, "박락") {
						concrete2 = true
					}

					if strings.Contains(data.Shape, "건조수축") {
						concrete3 = true
					}

					check[3] = true
				}

				if strings.Contains(data.Shape, "철근") {
					if strings.Contains(data.Shape, "노출") {
						rebar1 = true
					}

					if strings.Contains(data.Shape, "부식") {
						rebar2 = true
					}

					check[4] = true
				}

				if strings.Contains(data.Shape, "균열") {
					width := global.Atof(data.Width)
					if width >= 0.3 {
						width2 = true
					} else {
						width3 = true
					}

					if width > maxWidth {
						maxWidth = width
					}

					if width < minWidth {
						minWidth = width
					}

					if strings.Contains(data.Shape, "수직") {
						crack1 = true
					}

					if strings.Contains(data.Shape, "수평") {
						crack2 = true
					}

					if strings.Contains(data.Shape, "경사") {
						crack3 = true
					}

					if strings.Contains(data.Shape, "층간") {
						crack4 = true
					}

					if strings.Contains(data.Shape, "망상") {
						crack5 = true
					}

					check[0] = true
				}

				if strings.Contains(data.Shape, "누수") {
					strs := strings.Split(data.Shape, "/")
					for _, name := range strs {
						if name == "누수" {
							leak1 = true
						} else if name == "누수흔적" {
							leak2 = true
						} else if name == "백태" {
							leak3 = true
						}
					}

					check[2] = true
				}

				for j := 0; j < 6; j++ {
					if check[j] == false {
						continue
					}

					find := false
					for _, v := range parts[j] {
						if v == data.Part {
							find = true
							break
						}
					}

					if find == false {
						parts[j] = append(parts[j], data.Part)
					}

					find = false
					for _, v := range causes[j] {
						if v == data.Remark {
							find = true
							break
						}
					}

					if find == false {
						causes[j] = append(causes[j], data.Remark)
					}
				}
			}

			for j := 0; j < floorCount; j++ {
				if len(parts[j]) == 0 {
					continue
				}

				title := ""
				titles := make([]string, 0)
				methods := make([]string, 0)

				if j == 0 {
					if crack1 == true {
						titles = append(titles, "수직")
					}

					if crack2 == true {
						titles = append(titles, "수평")
					}

					if crack3 == true {
						titles = append(titles, "경사")
					}

					if crack4 == true {
						titles = append(titles, "층간")
					}

					if crack5 == true {
						titles = append(titles, "망상")
					}

					title = fmt.Sprintf("%v균열", strings.Join(titles, "/"))

					if width2 == true {
						methods = append(methods, "건식주입 공법(0.3mm 이상 균열)")
					}

					if width3 == true {
						methods = append(methods, "표면처리 공법(0.3mm 미만 균열)")
					}
				} else if j == 1 {
					if othercrack1 == true {
						titles = append(titles, "조적")
					}

					if othercrack2 == true {
						titles = append(titles, "이질재")
					}

					title = fmt.Sprintf("%v균열", strings.Join(titles, "/"))
					methods = append(methods, "충전식 공법")
				} else if j == 2 {
					titles = append(titles, "균열")
					if leak1 == true {
						titles = append(titles, "누수")
					} else if leak2 == true {
						titles = append(titles, "누수흔적")
					}

					if leak3 == true {
						titles = append(titles, "백태")
					}

					title = strings.Join(titles, "/")
					methods = append(methods, "습식주입 공법")
				} else if j == 3 {
					if concrete1 == true {
						titles = append(titles, "박리")
					}

					if concrete2 == true {
						titles = append(titles, "박락")
					}

					if concrete3 == true {
						titles = append(titles, "건조수축")
					}

					title = fmt.Sprintf("콘크리트 %v", strings.Join(titles, "/"))
					methods = append(methods, "단면복구 공법")
				} else if j == 4 {
					if rebar1 == true {
						titles = append(titles, "노출")
					}

					if rebar2 == true {
						titles = append(titles, "부식")
					}

					title = fmt.Sprintf("철근 %v", strings.Join(titles, "/"))
					methods = append(methods, "철근노출 보수공법")
				} else if j == 5 {
					if material1 == true {
						titles = append(titles, "오염")
					}

					if material2 == true {
						titles = append(titles, "파손")
					}

					title = fmt.Sprintf("마감재 %v", strings.Join(titles, "/"))
					methods = append(methods, "외부 우수유입 및 충격")
				}

				height := len(methods)
				if height < 3 {
					height = 3
				}
				summaryItem := SummaryResultItem{Part: strings.Join(parts[j], ", "), Type: j + 1, Result: title, Causes: causes[j], Methods: methods, Height: height, MaxWidth: maxWidth, MinWidth: minWidth}
				resultItems = append(resultItems, summaryItem)
			}

			if len(resultItems) > 0 {
				summaryResult.Types = append(summaryResult.Types, SummaryResultType{Type: i + 1, Count: len(resultItems), Items: resultItems})
				summaryResult.Count++

				for k, v := range resultItems {
					typeid := i + 1
					if k > 0 {
						typeid = 0
					}

					height := 0

					for _, v2 := range resultItems {
						height = len(v2.Methods)
					}

					summaryResult.Items = append(summaryResult.Items, SummaryResultList{Originaltype: i + 1, Type: typeid, Count: len(resultItems), Height: height, Item: v, Last: 0})
					summaryResult.Allcount++
					summaryResult.Height += height
				}
			}

		}

		for j := len(summaryResult.Items) - 1; j >= 0; j-- {
			if summaryResult.Items[j].Type > 0 {
				summaryResult.Items[j].Last = 1
				break
			}
		}

		cracks := make([]SummaryResultList, 0)
		leaks := make([]SummaryResultList, 0)
		concrets := make([]SummaryResultList, 0)
		materials := make([]SummaryResultList, 0)

		for _, v := range summaryResult.Items {
			if v.Item.Type == 1 || v.Item.Type == 2 {
				cracks = append(cracks, v)
			} else if v.Item.Type == 3 {
				leaks = append(leaks, v)
			} else if v.Item.Type == 4 || v.Item.Type == 5 {
				concrets = append(concrets, v)
			} else if v.Item.Type == 6 {
				materials = append(materials, v)
			}
		}

		item := CheckResult{Crack: make([]string, 0), Leak: make([]string, 0), Concret: make([]string, 0), Material: make([]string, 0)}

		positionTitles := []string{"지붕, 옥탑층", "지상층", "지하층", "외부벽체"}

		if topcount == 0 {
			positionTitles[0] = "지붕층"
		}

		{
			check := []bool{false, false, false, false, false, false, false, false}

			items := make([][]SummaryResultList, 4)

			for i := 0; i < 4; i++ {
				items[i] = make([]SummaryResultList, 0)
			}

			for _, v := range cracks {
				check[v.Originaltype-1] = true

				pos := (v.Originaltype - 1) / 2
				items[pos] = append(items[pos], v)
			}

			for i := 0; i <= 3; i++ {

				if check[i*2] == true || check[i*2+1] == true {
					results := make([]string, 0)
					for _, v := range items[i] {
						result := strings.ReplaceAll(v.Item.Result, "균열", "")

						if result == "" {
							continue
						}

						results = append(results, result)
					}
					crack := fmt.Sprintf("%v균열", strings.Join(results, "/"))

					title := ""

					if i != 3 {
						part := ""
						if check[i*2] == true && check[i*2+1] == true {
							part = "벽체, 슬래브"
						} else if check[i*2] == true {
							part = "벽체"
						} else {
							part = "슬래브"
						}

						title = fmt.Sprintf("%v %v", positionTitles[i], part)
					} else {
						title = positionTitles[i]
					}

					item.Crack = append(item.Crack, fmt.Sprintf("%v - %v", title, crack))
				}
			}
		}

		{
			check := []bool{false, false, false, false, false, false, false, false}

			items := make([][]SummaryResultList, 4)

			for i := 0; i < 4; i++ {
				items[i] = make([]SummaryResultList, 0)
			}

			for _, v := range leaks {
				check[v.Originaltype-1] = true

				pos := (v.Originaltype - 1) / 2
				items[pos] = append(items[pos], v)
			}

			for i := 0; i <= 3; i++ {

				if check[i*2] == true || check[i*2+1] == true {
					leak1 := false
					leak2 := false
					leak3 := false
					leak4 := false

					for _, v := range items[i] {
						results := strings.Split(v.Item.Result, "/")

						for _, v2 := range results {
							if v2 == "균열" {
								leak1 = true
							} else if v2 == "누수" {
								leak2 = true
							} else if v2 == "누수흔적" {
								leak3 = true
							} else if v2 == "백태" {
								leak4 = true
							}
						}
					}

					results := make([]string, 0)

					if leak1 == true {
						results = append(results, "균열")
					}

					if leak2 == true {
						results = append(results, "누수")
					} else if leak3 == true {
						results = append(results, "누수흔적")
					}

					if leak4 == true {
						results = append(results, "백태")
					}

					leak := strings.Join(results, "/")

					title := ""

					if i != 3 {
						part := ""
						if check[i*2] == true && check[i*2+1] == true {
							part = "벽체, 슬래브"
						} else if check[i*2] == true {
							part = "벽체"
						} else {
							part = "슬래브"
						}

						title = fmt.Sprintf("%v %v", positionTitles[i], part)
					} else {
						title = positionTitles[i]
					}

					item.Leak = append(item.Leak, fmt.Sprintf("%v - %v", title, leak))
				}
			}
		}

		{
			check := []bool{false, false, false, false, false, false, false, false}

			items := make([][]SummaryResultList, 4)

			for i := 0; i < 4; i++ {
				items[i] = make([]SummaryResultList, 0)
			}

			for _, v := range concrets {
				check[v.Originaltype-1] = true

				pos := (v.Originaltype - 1) / 2
				items[pos] = append(items[pos], v)
			}

			for i := 0; i <= 3; i++ {

				if check[i*2] == true || check[i*2+1] == true {
					concrete := false
					concrete1 := false
					concrete2 := false
					concrete3 := false
					rebar := false
					rebar1 := false
					rebar2 := false

					for _, v := range items[i] {
						if strings.Contains(v.Item.Result, "콘크리트") {
							if strings.Contains(v.Item.Result, "박리") {
								concrete1 = true
							}

							if strings.Contains(v.Item.Result, "박락") {
								concrete2 = true
							}

							if strings.Contains(v.Item.Result, "건조수축") {
								concrete3 = true
							}

							concrete = true
						}

						if strings.Contains(v.Item.Result, "철근") {
							if strings.Contains(v.Item.Result, "노출") {
								rebar1 = true
							}

							if strings.Contains(v.Item.Result, "부식") {
								rebar2 = true
							}

							rebar = true
						}
					}

					results := make([]string, 0)

					if concrete == true {
						temp := make([]string, 0)
						if concrete1 == true {
							temp = append(temp, "박리")
						}

						if concrete2 == true {
							temp = append(temp, "박락")
						}

						if concrete3 == true {
							temp = append(temp, "건조수축")
						}

						results = append(results, fmt.Sprintf("콘크리트 %v", strings.Join(temp, "/")))
					}

					if rebar == true {
						temp := make([]string, 0)
						if rebar1 == true {
							temp = append(temp, "노출")
						}

						if rebar2 == true {
							temp = append(temp, "부식")
						}

						results = append(results, fmt.Sprintf("철근 %v", strings.Join(temp, "/")))
					}

					concret := strings.Join(results, ", ")

					title := ""

					if i != 3 {
						part := ""
						if check[i*2] == true && check[i*2+1] == true {
							part = "벽체, 슬래브"
						} else if check[i*2] == true {
							part = "벽체"
						} else {
							part = "슬래브"
						}

						title = fmt.Sprintf("%v %v", positionTitles[i], part)
					} else {
						title = positionTitles[i]
					}

					item.Concret = append(item.Concret, fmt.Sprintf("%v - %v", title, concret))
				}
			}
		}

		{
			check := []bool{false, false, false, false, false, false, false, false}

			items := make([][]SummaryResultList, 4)

			for i := 0; i < 4; i++ {
				items[i] = make([]SummaryResultList, 0)
			}

			for _, v := range materials {
				check[v.Originaltype-1] = true

				pos := (v.Originaltype - 1) / 2
				items[pos] = append(items[pos], v)
			}

			for i := 0; i <= 3; i++ {
				if check[i*2] == true || check[i*2+1] == true {
					material1 := false
					material2 := false

					for _, v := range items[i] {
						if strings.Contains(v.Item.Result, "오염") {
							material1 = true
						}

						if strings.Contains(v.Item.Result, "파손") {
							material2 = true
						}

					}

					temp := make([]string, 0)

					if material1 == true {
						temp = append(temp, "오염")
					}

					if material2 == true {
						temp = append(temp, "파손")
					}

					material := fmt.Sprintf("마감재 %v", strings.Join(temp, "/"))

					title := ""

					if i != 3 {
						part := ""
						if check[i*2] == true && check[i*2+1] == true {
							part = "벽체, 슬래브"
						} else if check[i*2] == true {
							part = "벽체"
						} else {
							part = "슬래브"
						}

						title = fmt.Sprintf("%v %v", positionTitles[i], part)
					} else {
						title = positionTitles[i]
					}

					item.Material = append(item.Material, fmt.Sprintf("%v - %v", title, material))
				}
			}
		}

		item.CrackCount = len(item.Crack)
		item.LeakCount = len(item.Leak)
		item.ConcretCount = len(item.Concret)
		item.MaterialCount = len(item.Material)
		summaryResult.Check = item

		positionTitles[0] = "지붕 및 옥탑층"
		if topcount == 0 {
			positionTitles[0] = "지붕층"
		}

		numbers := []string{"①", "②", "③", "④", "⑤", "⑥", "⑦", "⑧", "⑨", "⑩", "⑪", "⑫", "⑬", "⑭", "⑮", "⑯", "⑰", "⑱", "⑲", "⑳"}
		numberPos := 0
		resultItems := make([]string, 0)

		items := make([][]SummaryResultList, 4)

		for i := 0; i < 4; i++ {
			items[i] = make([]SummaryResultList, 0)
		}

		check := []bool{false, false, false, false, false, false, false, false}

		for _, v := range summaryResult.Items {
			check[v.Originaltype-1] = true

			pos := (v.Originaltype - 1) / 2
			items[pos] = append(items[pos], v)
		}

		outerwallMaterial := "수성페인트"
		for _, v := range otherMap[14] {
			if v.Order == 141 {
				if v.Status != "" {
					outerwallMaterial = strings.ReplaceAll(v.Status, ",", ", ")
					break
				}
			}
		}

		for i := 0; i < 4; i++ {
			if check[i*2] == true || check[i*2+1] == true {
				parts := make([]string, 0)

				results := make([]string, 0)
				allresults := make([]string, 0)
				allresults2 := make([]string, 0)

				var maxWidth float64 = -9999.0
				var minWidth float64 = 9999.0

				for _, v := range items[i] {
					temp := strings.Split(v.Item.Part, ", ")
					for _, v2 := range temp {
						find := false
						for _, v3 := range parts {
							if v2 == v3 {
								find = true
								break
							}
						}

						if find == false {
							parts = append(parts, v2)
						}
					}

					if v.Originaltype == i*2+1 {
						if v.Item.Type == 1 || v.Item.Type == 2 {
							if v.Item.MaxWidth > maxWidth {
								maxWidth = v.Item.MaxWidth
							}

							if v.Item.MinWidth < minWidth {
								minWidth = v.Item.MinWidth
							}

							result := strings.ReplaceAll(v.Item.Result, "균열", "")

							if result != "" {
								results = append(results, result)
							}
						} else {
							allresults = append(allresults, v.Item.Result)
						}
					} else {
						allresults2 = append(allresults2, v.Item.Result)
					}
				}

				text := ""

				if len(results) > 0 {
					width := ""
					if minWidth == maxWidth {
						width = fmt.Sprintf("%vmm ", minWidth)
					} else if minWidth < 9000 {
						if maxWidth > 0 {
							width = fmt.Sprintf("%v ~ %vmm ", minWidth, maxWidth)
						} else {
							width = fmt.Sprintf("%vmm ", minWidth)
						}
					} else {
						if maxWidth > 0 {
							width = fmt.Sprintf("%vmm ", maxWidth)
						}
					}

					crack := fmt.Sprintf("%v균열(%v)", width, strings.Join(results, "/"))

					if len(allresults) > 0 {
						text = fmt.Sprintf("벽체(기둥 포함) %v과 %v", crack, strings.Join(allresults, ", "))
					} else {
						text = fmt.Sprintf("벽체(기둥 포함) %v", crack)
					}
				} else {
					if len(allresults) > 0 {
						text = fmt.Sprintf("벽체(기둥 포함) %v", strings.Join(allresults, ", "))
					}
				}

				if len(allresults2) > 0 {
					if text != "" {
						text = fmt.Sprintf("%v, 슬래브(보 포함) %v", text, strings.Join(allresults2, ", "))
					} else {
						text = fmt.Sprintf("슬래브(보 포함) %v", strings.Join(allresults2, ", "))
					}
				}

				if i == 3 {
					text = ""

					log.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
					other := others[14]

					allresults = append(allresults, other.Items...)
					if len(results) > 0 {
						crack := fmt.Sprintf("균열(%v)", strings.Join(results, "/"))

						if len(allresults) > 0 {
							text = fmt.Sprintf("외부벽체 %v 마감부위에서 %v과 %v", outerwallMaterial, crack, strings.Join(allresults, ", "))
						} else {
							text = fmt.Sprintf("외부벽체 %v 마감부위에서 %v", outerwallMaterial, crack)
						}
					} else {
						if len(allresults) > 0 {
							text = fmt.Sprintf("외부벽체 %v 마감부위에서 %v", outerwallMaterial, strings.Join(allresults, ", "))
						}
					}

					if len(allresults2) > 0 {
						if text != "" {
							text += " "
						}

						text = fmt.Sprintf("%v, %v", text, strings.Join(allresults2, ", "))
					}

					resultItems = append(resultItems, strings.ReplaceAll(fmt.Sprintf("%v %v", numbers[numberPos], text), " ,", ","))
					numberPos++
				} else {
					resultItems = append(resultItems, strings.ReplaceAll(fmt.Sprintf("%v %v %v", numbers[numberPos], positionTitles[i], text), " ,", ","))
					numberPos++
				}

			} else {
				if i == 3 {
					other := others[14]

					if other.Head != "" {
						resultItems = append(resultItems, strings.ReplaceAll(fmt.Sprintf("%v %v", numbers[numberPos], other.Head), " ,", ","))
						numberPos++
					} else {
						if len(other.Items) > 0 {
							resultItems = append(resultItems, fmt.Sprintf("%v 외부벽체 %v", numbers[numberPos], strings.Join(other.Items, ", ")))
							numberPos++
						}
					}
				}
			}

		}

		resultItems = append(resultItems, fmt.Sprintf("%v 전체적으로 구조체의 균열은 대부분 비진행성인 것으로 판단됨", numbers[numberPos]))
		numberPos++

		changeCount := 0
		for _, v := range periodicchanges {
			if v.Type == 1 {
				changeCount++
			}
		}

		if changeCount > 0 {
			resultItems = append(resultItems, fmt.Sprintf("%v 이전 사용검사일 이후 용도변경 및 증.개축은 %v건이 있는 것으로 조사되었으며, 설계하중을 초과하는 특기할 만한 하중변화는 없는 것으로 판단됨.", numbers[numberPos], changeCount))
		} else {
			resultItems = append(resultItems, fmt.Sprintf("%v 이전 사용검사일 이후 용도변경 및 증.개축은 없는 것으로 조사되었으며, 설계하중을 초과하는 특기할 만한 하중변화는 없는 것으로 판단됨.", numbers[numberPos]))
		}
		numberPos++

		{
			items := make([]models.Periodicdata, 0)
			for i := 2; i < 5; i++ {
				items = append(items, summarys[i].Data...)
			}

			result := Data(items)
			summaryResult.ResultBottom = result.ResultsForResult()
		}

		summaryResult.Result = resultItems

		summaryResults = append(summaryResults, summaryResult)
	}

	return summaryResults
}

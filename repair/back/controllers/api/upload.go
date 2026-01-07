package api

import (
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"repair/controllers"
	"repair/global"
	"repair/global/config"
	"repair/models"

	"github.com/xuri/excelize/v2"
)

type UploadController struct {
	controllers.Controller
}

// @POST()
func (c *UploadController) Index() {
	param := c.Get("param")
	path := c.Get("path")
	path = strings.ReplaceAll(path, "..", "")
	path = strings.ReplaceAll(path, "/", "")

	originalfilename, file := c.GetUpload(path, "file")

	c.Set("originalfilename", originalfilename)
	c.Set("filename", file)
	c.Set("param", param)

	if path == "periodic" {
		global.SendNotify(global.Notify{Type: global.NotifyImage, Filename: file})
	}
}

// @POST()
func (c *UploadController) Periodic() {
	originalfilename, file := c.GetUpload("periodic", "file")

	c.Set("originalfilename", originalfilename)
	c.Set("filename", file)

	global.SendNotify(global.Notify{Type: global.NotifyImage, Filename: file})
}

func (c *UploadController) Diff(id int64, filename string) {
	DiffProcess(c, id, filename)
}

// @POST()
func (c *UploadController) Diffupdate(diffs *models.Diff) {
	log.Println("diffupdate filename", diffs.Filename)
	log.Println(diffs)
	DiffupdateProcess(diffs)
}

func (c *UploadController) Excel(id int64, filename string, historydel int, breakdowndel int) {
	ExcelProcess(id, filename, historydel, breakdowndel)
}

func ExcelProcess(id int64, filename string, historydel int, breakdowndel int) {
	conn := models.NewConnection()
	defer conn.Close()

	conn.Begin()
	defer conn.Rollback()

	//aptManager := models.NewAptManager(conn)
	repairManager := models.NewRepairManager(conn)
	areaManager := models.NewAreaManager(conn)
	dongManager := models.NewDongManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	historyManager := models.NewHistoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)
	standardManager := models.NewStandardManager(conn)

	//apt := aptManager.Get(id)

	categorys := categoryManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id")})
	standards := standardManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("s_order,s_id")})
	//historys := historyManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("h_id")})
	//areas := areaManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("ar_order,ar_id")})
	dongs := dongManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("d_order,d_id")})
	//breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_id")})

	categoryMap := make(map[int64]models.Category)
	standardMap := make(map[int64]models.Standard)
	//dongMap := make(map[int64]models.Dong)

	level2Title := []string{"", "가", "나", "다", "라", "마", "바", "사", "아", "자", "차", "카", "타", "파", "하", "거", "너", "더", "러", "머", "버", "서", "어", "저", "처", "커", "터", "퍼", "허"}

	for _, v := range categorys {
		categoryMap[v.Id] = v
	}

	for _, v := range standards {
		standardMap[v.Id] = v
	}

	f, err := excelize.OpenFile(path.Join(config.UploadPath, filename))
	if err != nil {
		log.Println(err)
		return
	}

	areaManager.DeleteByApt(id)
	dongManager.DeleteByApt(id)

	if historydel == 1 {
		historyManager.DeleteByApt(id)
	}

	if breakdowndel == 1 {
		breakdownManager.DeleteByApt(id)
	}

	pos := 160
	sheet := "표지.목차.개요"

	_, err = f.GetCellValue(sheet, GetCell("B", 9))
	if err != nil {
		sheet = "표지.목차.개요(2)"
	}

	for i := 160; i < 170; i++ {
		str, _ := f.GetCellValue(sheet, GetCell("A", i))
		if str == "1) 부대복리시설 면적 현황" {
			pos = i + 3
			break
		}
	}

	repair := repairManager.Get(id)
	str1, _ := f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure1 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure2 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure3 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure4 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure5 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure6 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure7 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure8 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure9 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure10 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure11 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure12 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure13 = str1
	pos++
	str1, _ = f.GetCellValue(sheet, GetCell("K", pos))
	repair.Structure14 = str1
	pos++

	repairManager.Update(repair)

	pos += 3

	for {
		no, _ := f.GetCellValue(sheet, GetCell("A", pos))

		if no == "" {
			break
		}

		familycount, _ := f.GetCellValue(sheet, GetCell("E", pos))
		sizeStr, _ := f.GetCellValue(sheet, GetCell("G", pos))
		remark, _ := f.GetCellValue(sheet, GetCell("M", pos))

		size := global.Atof(sizeStr)

		var item models.Area
		item.Apt = id
		item.Familycount = global.Atoi(familycount)
		item.Size = models.Double(size)
		item.Remark = remark

		if item.Familycount > 0 && item.Size > 0 {
			areaManager.Insert(&item)
		}

		pos++
	}

	pos += 2

	titleCheck, _ := f.GetCellValue(sheet, GetCell("A", pos))
	switch titleCheck {
	case "3) 동별 층수 및 세대수":
		pos += 2
	case "구분":
		pos += 1
	}

	dongs = make([]models.Dong, 0)
	for {
		no, _ := f.GetCellValue(sheet, GetCell("A", pos))

		if no == "" {
			break
		}

		name, _ := f.GetCellValue(sheet, GetCell("C", pos))
		ground, _ := f.GetCellValue(sheet, GetCell("E", pos))
		underground, _ := f.GetCellValue(sheet, GetCell("G", pos))
		familycount, _ := f.GetCellValue(sheet, GetCell("J", pos))

		var item models.Dong
		item.Apt = id
		item.Name = strings.ReplaceAll(name, "동", "") + "동"
		item.Ground = global.Atoi(ground)
		item.Underground = global.Atoi(underground)
		item.Familycount = global.Atoi(familycount)
		item.Basic = 1

		dongManager.Insert(&item)
		item.Id = dongManager.GetIdentity()
		pos++

		dongs = append(dongs, item)
	}

	sheet = "장기수선충당금 사용현황"

	_, err = f.GetCellValue(sheet, GetCell("A", 3))
	if err != nil {
		sheet = "장기수선충당금 사용현황 "
	}

	emptyCount := 0
	row := 3
	for {
		date, err := f.GetCellValue(sheet, GetCell("A", row))
		if err != nil {
			log.Println(err)
		}

		if date == "" || date == "장기수선 충당금 사용 계" || date == "합계" {
			if date == "" {
				row++

				emptyCount++

				if emptyCount == 4 {
					break
				}

				continue
			}

			break
		}

		topcategoryStr, _ := f.GetCellValue(sheet, GetCell("B", row))
		subcategoryStr, _ := f.GetCellValue(sheet, GetCell("C", row))
		categoryStr, _ := f.GetCellValue(sheet, GetCell("D", row))
		//

		if topcategoryStr == "6. 옥외부대복리시설" {
			topcategoryStr = "6. 옥외 부대시설 및 옥외 복리시설"
		}

		if topcategoryStr == "3. 전기소화승강기 및 지능형 홈 네트워크설비" {
			topcategoryStr = "3. 전기·소화·승강기 및 지능형 홈네트워크 설비"
		}

		if topcategoryStr == "4. 급수·가스·배수 및 환기시설" {
			topcategoryStr = "4. 급수·가스·배수 및 환기설비"
		}

		if subcategoryStr == "가. 옥외부대시설 및 복리시설" {
			subcategoryStr = "가. 옥외부대시설 및 옥외 복리시설"
		}

		if subcategoryStr == "다. 외부창.문" {
			subcategoryStr = "다. 외부 창·문"
		}

		if subcategoryStr == "라. 환기팬" {
			subcategoryStr = "라. 환기설비"
		}

		if subcategoryStr == "자. 보안.방범시설" {
			subcategoryStr = "자. 보안․방범시설"
		}

		if subcategoryStr == "차. 지능형 홈네워크 설비" {
			subcategoryStr = "차. 지능형 홈네트워크 설비"
		}

		var topcategory models.Category
		var subcategory models.Category
		var category models.Category

		if topcategoryStr != "" {
			flag := false

			for _, v := range categorys {
				if v.Name == topcategoryStr && v.Level == 1 {
					topcategory = v
					flag = true
					break
				}
			}

			if flag == false {
				strs := strings.Split(topcategoryStr, ". ")
				if len(strs) == 2 {
					strPos := global.Atoi(strs[0])
					current := 1

					for _, v := range categorys {
						if v.Level != 1 {
							continue
						}

						if current == strPos {
							topcategory = v
							flag = true
							break
						}

						current++
					}
				}
			}

			if flag == true {
				flag = false

				for _, v := range categorys {
					if v.Level != 2 {
						continue
					}

					if v.Parent != topcategory.Id {
						continue
					}

					s1 := strings.SplitN(v.Name, " ", 2)
					s2 := strings.SplitN(subcategoryStr, " ", 2)
					if v.Name == subcategoryStr {
						subcategory = v
						flag = true
						break
					}

					if len(s1) >= 2 && len(s2) >= 2 && s1[1] == s2[1] {
						subcategory = v
						flag = true
						break
					}
				}

				if flag == false {
					strs := strings.Split(subcategoryStr, ". ")
					if len(strs) == 2 && strs[0] != "" {
						strPos := 0
						for i, v := range level2Title {
							if v == strs[0] {
								strPos = i
								break
							}
						}

						current := 1

						for _, v := range categorys {
							if v.Level != 2 {
								continue
							}

							if current == strPos {
								subcategory = v
								flag = true
								break
							}

							current++
						}
					}
				}

				if flag == true {
					flag = false

					for _, v := range categorys {
						if v.Level != 3 {
							continue
						}

						if v.Parent != subcategory.Id {
							continue
						}

						if v.Name == categoryStr {
							category = v
							flag = true
							break
						}
					}

					if flag == false {
						for _, v := range categorys {
							if v.Level != 3 {
								continue
							}

							if v.Parent != subcategory.Id {
								continue
							}

							strs := strings.Split(v.Name, " ")
							if len(strs) == 2 && strs[1] == "기타" {
								category = v
								flag = true
								break
							}
						}
					}

					if flag == false {
						strs := strings.Split(categoryStr, ") ")
						if len(strs) == 2 && strs[0] != "" {
							strs[0] = strings.ReplaceAll(strs[0], "(", "")
							strPos := global.Atoi(strs[0])
							current := 1

							for _, v := range categorys {
								if v.Level != 3 {
									continue
								}

								if current == strPos {
									category = v
									flag = true
									break
								}

								current++
							}
						}
					}
				}
			}
		}

		content, _ := f.GetCellValue(sheet, GetCell("E", row))
		priceStr, _ := f.GetCellValue(sheet, GetCell("F", row))
		price := global.Atol(strings.ReplaceAll(priceStr, ",", ""))

		dates := strings.Split(strings.ReplaceAll(date, "월", ""), "년 ")

		year := 0
		month := 0
		if len(dates) == 2 {
			year = global.Atoi(dates[0])
			month = global.Atoi(dates[1])
		}

		var item models.History
		item.Apt = id
		item.Year = year
		item.Month = month
		item.Topcategory = topcategory.Id
		item.Subcategory = subcategory.Id
		item.Category = category.Id
		item.Content = content
		item.Price = price

		historyManager.Insert(&item)
		row++

	}

	sheet = "통합세부자료"

	{
		_, err := f.GetCellValue(sheet, GetCell("A", 1))
		if err != nil {
			sheet = "통합세부자료 "

			_, err := f.GetCellValue(sheet, GetCell("A", 1))
			if err != nil {
				sheet = "내역세부"

				_, err := f.GetCellValue(sheet, GetCell("A", 1))
				if err != nil {
					sheet = "항목세부자료"
				}
			}
		}
	}

	pos = 2
	emptyCount = 1

	var topcategory models.Category
	var subcategory models.Category
	var category models.Category

	cell1 := "I"
	cell2 := "L"
	cell3 := "M"
	cell4 := "N"

	cellCount := "G"

	for i := range 5 {
		dong, _ := f.GetCellValue(sheet, GetCell("A", pos+i))
		dong = strings.ReplaceAll(dong, "\n", "")
		dong = strings.ReplaceAll(dong, "\r", "")
		dong = strings.ReplaceAll(dong, "\t", "")
		dong = strings.ReplaceAll(dong, " ", "")
		if dong == "시설물위치" {
			pos = pos + i + -1

			ptitle, _ := f.GetCellValue(sheet, GetCell("W", pos+i))

			if ptitle == "단가" {
				cellCount = "U"
				cell1 = "W"
				cell2 = "Z"
				cell3 = "AA"
				cell4 = "AB"

				break
			}

			ptitle, _ = f.GetCellValue(sheet, GetCell("I", pos+i))

			if ptitle == "단가" {
			} else {
				cell1 = "H"
				cell2 = "K"
				cell3 = "L"
				cell4 = "M"
			}

			ptitle, _ = f.GetCellValue(sheet, GetCell(cellCount, pos+i))

			if ptitle == "수량" {
			} else {
				titles := []string{"G", "H", "I", "J"}

				for _, v := range titles {
					ptitle, _ = f.GetCellValue(sheet, GetCell(v, pos+i))

					if ptitle == "수량" {
						cellCount = v
						break
					}
				}
			}
			break
		}
	}

	log.Println("pos = ", pos, cell1, cell2, cell3, cell4, cellCount)

	for {
		dong, _ := f.GetCellValue(sheet, GetCell("A", pos))

		var item models.Breakdown
		if dong == "" {
			t, _ := f.GetCellValue(sheet, GetCell("B", pos))
			if t != "" {
				emptyCount = 0
				pos++
				continue
			}

			emptyCount++

			if emptyCount == 5 {
				break
			}

			pos++
			continue
		} else {
			lineCheck1, _ := f.GetCellValue(sheet, GetCell("A", pos))
			lineCheck2, _ := f.GetCellValue(sheet, GetCell("B", pos))

			if emptyCount > 0 && lineCheck1 == lineCheck2 {
				emptyCount = 1
			}

			if emptyCount == 1 {
				// 찾고

				strs := strings.Split(dong, " > ")

				topcategory.Id = 0
				subcategory.Id = 0
				category.Id = 0

				flag := false
				step := 0
				if len(strs) == 3 {
					if strs[0] == "6. 옥외부대복리시설" {
						strs[0] = "6. 옥외 부대시설 및 옥외 복리시설"
					}

					if strs[0] == "3. 전기소화승강기 및 지능형 홈 네트워크설비" {
						strs[0] = "3. 전기·소화·승강기 및 지능형 홈네트워크 설비"
					}

					if strs[0] == "4. 급수·가스·배수 및 환기시설" {
						strs[0] = "4. 급수·가스·배수 및 환기설비"
					}

					if strs[1] == "가. 옥외부대시설 및 복리시설" {
						strs[1] = "가. 옥외부대시설 및 옥외 복리시설"
					}

					if strs[1] == "가. 천정" {
						strs[1] = "가. 천장"
					}

					if strs[1] == "다. 외부창.문" {
						strs[1] = "다. 외부 창·문"
					}

					if strs[1] == "라. 환기팬" {
						strs[1] = "라. 환기설비"
					}

					if strs[1] == "자. 보안.방범시설" {
						strs[1] = "자. 보안․방범시설"
					}

					if strs[1] == "차. 지능형 홈네워크 설비" {
						strs[1] = "차. 지능형 홈네트워크 설비"
					}

					for _, t := range categorys {
						if t.Level == 1 && t.Name == strings.TrimSpace(strs[0]) {
							step = 1
							topcategory = t

							for _, s := range categorys {
								s1 := strings.SplitN(s.Name, " ", 2)
								s2 := strings.SplitN(strings.TrimSpace(strs[1]), " ", 2)
								if s.Level == 2 && s.Parent == t.Id && len(s1) >= 2 && len(s2) >= 2 && s1[1] == s2[1] {
									step = 2
									subcategory = s

									for _, c := range categorys {
										step = 3
										if c.Level == 3 && c.Parent == s.Id && c.Name == strings.TrimSpace(strs[2]) {
											category = c

											flag = true
											break
										}
									}

									if flag == false {
										for _, c := range categorys {
											if c.Level != 3 {
												continue
											}

											if c.Parent != s.Id {
												continue
											}

											s1 := strings.SplitN(c.Name, " ", 2)
											if s1[1] == "기타" {
												step = 4
												category = c

												flag = true
												break
											}
										}
									}

									break
								}
							}

							break
						}
					}
				}

				if flag == false {
					log.Println("세부 못찾음", step, strs)
				}

				t1, _ := f.GetCellValue(sheet, GetCell("A", pos+1))
				t2, _ := f.GetCellValue(sheet, GetCell("A", pos+2))

				if t1 == t2 {
					pos += 3
				} else {
					pos += 2
				}

				emptyCount = 0
				continue

			} else {
				flag := false
				for _, v := range dongs {
					if v.Apt != id {
						continue
					}

					if v.Name == dong {
						item.Dong = v.Id
						flag = true
						break
					}
				}

				if flag == false {
					dong = strings.ReplaceAll(dong, "호기", "")
					strs := strings.Split(dong, "동 ")

					dong = strs[0]

					if len(strs) == 2 {
						dong += "동"

						for _, v := range dongs {
							if v.Apt != id {
								continue
							}

							if v.Name == dong {
								item.Dong = v.Id
								flag = true
								break
							}
						}

						if flag == true {
							item.Elevator = global.Atoi(strs[1])
						}
					} else {
						for _, v := range dongs {
							if v.Apt != id {
								continue
							}

							if v.Name == dong {
								item.Dong = v.Id
								flag = true
								break
							}
						}
					}
				}

				if flag == false {
					var nitem models.Dong
					nitem.Apt = id
					nitem.Name = dong
					nitem.Ground = 0
					nitem.Underground = 0
					nitem.Familycount = 0
					nitem.Basic = 1

					dongManager.Insert(&nitem)
					nitem.Id = dongManager.GetIdentity()

					dongs = append(dongs, nitem)
					item.Dong = nitem.Id
				}
			}
		}

		standard, _ := f.GetCellValue(sheet, GetCell("B", pos))
		method, _ := f.GetCellValue(sheet, GetCell("C", pos))
		//cycle, _ := f.GetCellValue(sheet, GetCell("D", pos))
		percent, _ := f.GetCellValue(sheet, GetCell("E", pos))
		unit, _ := f.GetCellValue(sheet, GetCell("F", pos))
		count, _ := f.GetCellValue(sheet, GetCell(cellCount, pos))
		price, _ := f.GetCellValue(sheet, GetCell(cell1, pos))
		lastdate, _ := f.GetCellValue(sheet, GetCell(cell2, pos))
		duedate, _ := f.GetCellValue(sheet, GetCell(cell3, pos))
		remark, _ := f.GetCellValue(sheet, GetCell(cell4, pos))

		item.Apt = id

		for _, v2 := range standards {
			if v2.Apt != id {
				continue
			}

			if v2.Category != category.Id {
				continue
			}

			if v2.Name == standard {
				item.Standard = v2.Id
				originalPrice := CalculatePriceRate(v2.Direct, v2.Labor, v2.Cost, 0, 0.0)

				iprice := global.Atoi(strings.ReplaceAll(price, ",", ""))
				if iprice != int(originalPrice) {

					direct := reversPrice(global.Atol(strings.ReplaceAll(price, ",", "")))
					calcuratePrice := CalculatePriceRate(direct, 0, 0, 0, 0.0)

					if direct != calcuratePrice {
						if direct > calcuratePrice {
							v2.Direct--
						} else {
							v2.Direct++
						}
					}

					standardManager.Update(&v2)
				}
				break
			}
		}

		if item.Standard == 0 {
			var n models.Standard

			n.Apt = id
			n.Category = category.Id
			n.Name = standard

			n.Direct = reversPrice(global.Atol(strings.ReplaceAll(price, ",", "")))
			n.Unit = unit

			standardManager.Insert(&n)

			item.Standard = standardManager.GetIdentity()
			n.Id = item.Standard

			standards = append(standards, n)
		}

		methods := make([]models.Category, 0)
		for _, v2 := range categorys {
			if v2.Apt != id {
				continue
			}

			if v2.Level != 4 {
				continue
			}

			if v2.Parent != category.Id {
				continue
			}

			if v2.Name == method {
				item.Method = v2.Id
				break
			}

			methods = append(methods, v2)
		}

		if item.Method == 0 {
			if len(methods) == 1 {
				item.Method = methods[0].Id
			} else if len(methods) == 2 {
				if global.Atoi(percent) == 100 {
					if methods[0].Name[:2] == "전면" {
						item.Method = methods[0].Id
					} else {
						item.Method = methods[1].Id
					}
				} else {
					if methods[0].Name[:2] == "부분" {
						item.Method = methods[0].Id
					} else {
						item.Method = methods[1].Id
					}
				}
			}
		}

		item.Topcategory = topcategory.Id
		item.Subcategory = subcategory.Id
		item.Category = category.Id

		item.Count = global.Atoi(count)
		item.Lastdate = global.Atoi(lastdate)
		item.Duedate = global.Atoi(duedate)
		item.Remark = remark

		breakdownManager.Insert(&item)

		pos++

		emptyCount = 0
	}

	conn.Commit()

	f.Close()

	os.Remove(path.Join(config.UploadPath, filename))
}

func reversPrice(directInt int64) int64 {
	value := float64(directInt)

	value = value * 10.0 / 11.0
	direct := value * 100.0 / (100.0 + 5.5 + 3.09 + 0.3 + 0.07 + 6.0 + 0.9)

	return int64(direct)
}

// @POST()
func (c *UploadController) Assistance(id int64, filenames []string, historydel int) {
	conn := models.NewConnection()
	defer conn.Close()

	conn.Begin()
	defer conn.Rollback()

	savingManager := models.NewSavingManager(conn)
	historyManager := models.NewHistoryManager(conn)

	if historydel == 1 {
		historyManager.DeleteByApt(id)
	}

	historys := make([]models.History, 0)
	savings := make([]models.Saving, 0)

	for _, filename := range filenames {
		fullFilename := path.Join(config.UploadPath, filename)
		f, err := excelize.OpenFile(fullFilename)
		if err != nil {
			log.Println(err)
			return
		}

		sheet := f.GetSheetName(0)

		pos := 0
		empty := 0
		olddate := ""

		year := 0
		var totalForward int64 = 0
		var totalInterest int64 = 0
		var totalSurplus int64 = 0
		var totalSaving int64 = 0
		var totalUse int64 = 0

		sign1, _ := f.GetCellValue(sheet, GetCell("M", 6))
		sign2, _ := f.GetCellValue(sheet, GetCell("Q", 5))

		position1 := "H"
		position2 := "L"

		datePosition1 := "C"
		datePosition2 := "E"

		if sign1 == "담당" {
		} else if sign2 == "담당" {
			position1 = "K"
			position2 = "O"
		} else {

		}

		cols := []string{"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U"}
		for i := 4; i < 15; i++ {
			str, _ := f.GetCellValue(sheet, GetCell("C", i))
			str2, _ := f.GetCellValue(sheet, GetCell("D", i))
			str3, _ := f.GetCellValue(sheet, GetCell("B", i))
			str4, _ := f.GetCellValue(sheet, GetCell("A", i))
			str = strings.TrimSpace(str)
			str2 = strings.TrimSpace(str2)
			str3 = strings.TrimSpace(str3)
			str4 = strings.TrimSpace(str4)

			s, _ := f.GetCellValue(sheet, GetCell("E", i))
			s2, _ := f.GetCellValue(sheet, GetCell("F", i))
			s3, _ := f.GetCellValue(sheet, GetCell("D", i))
			s = strings.TrimSpace(s)
			s2 = strings.TrimSpace(s2)
			s3 = strings.TrimSpace(s3)

			if str == "전표일자" || str2 == "전표일자" || str3 == "전표일자" || str4 == "전표일자" {
				if str3 == "전표일자" {
					datePosition1 = "B"
					datePosition2 = "D"
				} else if str == "전표일자" {
				} else if str2 == "전표일자" {
					datePosition1 = "D"
					datePosition2 = "F"
				} else if str4 == "전표일자" {
					datePosition1 = "A"
					datePosition2 = "D"
				}

				if s == "적요" {
					datePosition2 = "E"
				} else if s2 == "적요" {
					datePosition2 = "F"
				} else if s3 == "적요" {
					datePosition2 = "D"
				}

				find := ""

				for _, v := range cols {
					find, _ = f.GetCellValue(sheet, GetCell(v, i))
					find = strings.TrimSpace(find)
					if find == "차변" {
						position1 = v
						break
					}
				}

				for _, v := range cols {
					find, _ = f.GetCellValue(sheet, GetCell(v, i))
					find = strings.TrimSpace(find)
					if find == "대변" {
						position2 = v
						break
					}
				}

				break
			}
		}

		for {
			date, _ := f.GetCellValue(sheet, GetCell(datePosition1, pos))
			title, _ := f.GetCellValue(sheet, GetCell(datePosition2, pos))
			useStr, _ := f.GetCellValue(sheet, GetCell(position1, pos))
			inStr, _ := f.GetCellValue(sheet, GetCell(position2, pos))

			inStr = strings.ReplaceAll(inStr, "(", "")
			inStr = strings.ReplaceAll(inStr, ")", "")

			use := global.Atol(useStr)
			in := global.Atol(inStr)

			pos++

			if date == "" && strings.ReplaceAll(title, " ", "") == "전기이월" {
				totalForward = in
			} else if date != title {
				if date == "" {
					date = olddate
				}

				date = strings.ReplaceAll(date, "-", ".")
				dates := strings.Split(date, ".")

				if len(dates) == 3 {
					year = global.Atoi(dates[0])

					if use > 0 {
						var history models.History
						history.Year = year
						history.Month = global.Atoi(dates[1])

						history.Content = title
						history.Apt = id
						history.Price = use
						historys = append(historys, history)

						totalUse += use
					} else if in > 0 {
						if strings.Contains(title, "만기이자") || strings.Contains(title, "만기 이자") || strings.Contains(title, "예치만기") || strings.Contains(title, "만기해지") || strings.Contains(title, "중도해지") {
							totalInterest += in
						} else if strings.Contains(title, "월분 장기수선충당전입액") || strings.Contains(title, "장기수선충당금으로 예치") || strings.Contains(title, "장기수선충당예치") || strings.Contains(title, "장기수선충당금적립") || strings.Contains(title, "장기수선충당금 적립") || (strings.Contains(title, "장기수선충당금") && (strings.Contains(title, "회") || strings.Contains(title, "월분") || strings.Contains(title, "전입액") || strings.Contains(title, "적립"))) {
							totalSaving += in
						} else if strings.Contains(title, "이자수익 처리오류") {
							in = global.Atol(inStr)
							totalInterest -= in
						} else {
							totalSurplus += in
						}
					}

				}
			} else if date == "" {
				empty++

				if empty > 10 {
					break
				}

				continue
			}

			olddate = date
			empty = 0
		}

		f.Close()

		var saving models.Saving
		saving.Year = year
		saving.Forward = totalForward
		saving.Interest = totalInterest
		saving.Surplus = totalSurplus
		saving.Saving = totalSaving
		saving.Use = totalUse
		saving.Apt = id

		savings = append(savings, saving)

		f.Close()
		os.Remove(path.Join(config.UploadPath, filename))
	}

	sort.Slice(historys, func(i, j int) bool {
		if historys[i].Year == historys[j].Year {
			return historys[i].Month < historys[j].Month
		} else {
			return historys[i].Year < historys[j].Year
		}
	})

	for _, v := range historys {
		historyManager.Insert(&v)
	}

	sort.Slice(savings, func(i, j int) bool {
		return savings[i].Year < savings[j].Year
	})

	if historydel == 1 {
		savingManager.DeleteByApt(id)
	}
	for _, v := range savings {
		savingManager.Insert(&v)
	}

	conn.Commit()
}

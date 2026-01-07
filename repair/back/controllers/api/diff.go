package api

import (
	"fmt"
	"log"
	"os"
	"path"
	"repair/global/config"
	"repair/global"
	"repair/models"
	"strings"

	"github.com/xuri/excelize/v2"
)

func DiffProcess(c *UploadController, id int64, filename string) {
	conn := models.NewConnection()
	defer conn.Close()

	update := false

	var dongManager *models.DongManager
	var categoryManager *models.CategoryManager
	var breakdownManager *models.BreakdownManager
	var standardManager *models.StandardManager

	if update == true {
		conn.Begin()
		defer conn.Rollback()

		dongManager = models.NewDongManager(conn)
		categoryManager = models.NewCategoryManager(conn)
		breakdownManager = models.NewBreakdownManager(conn)
		standardManager = models.NewStandardManager(conn)

	} else {
		dongManager = models.NewDongManager(conn)
		categoryManager = models.NewCategoryManager(conn)
		breakdownManager = models.NewBreakdownManager(conn)
		standardManager = models.NewStandardManager(conn)
	}

	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}})
	categorys := categoryManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id")})
	standards := standardManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("s_order,s_id")})
	dongs := dongManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("d_order,d_id")})

	categoryMap := make(map[int64]models.Category)
	standardMap := make(map[int64]models.Standard)

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

	pos := 160
	sheet := "통합세부자료"

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

	changeStandards := make([]models.Standard, 0)
	newStandards := make([]models.Standard, 0)
	newBreakdowns := make([]models.Breakdown, 0)

	pos = 2
	emptyCount := 1

	var topcategory models.Category
	var subcategory models.Category
	var category models.Category

	cell1 := "I"
	cell2 := "L"
	cell3 := "M"
	cell4 := "N"

	cellCount := "G"

	for i := 0; i < 5; i++ {
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

	log.Println("pos 4 = ", pos, cell1, cell2, cell3, cell4, cellCount)

	for {
		dong, _ := f.GetCellValue(sheet, GetCell("A", pos))
		log.Println(dong)
		for i := 1; i < 20; i++ {
			dong = strings.ReplaceAll(dong, fmt.Sprintf(" %v호기", i), "")
		}

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

		diffStandards := make([]models.Standard, 0)
		for _, v2 := range standards {
			if v2.Apt != id {
				continue
			}

			if v2.Category != category.Id {
				continue
			}

			if v2.Name == standard {
				diffStandards = append(diffStandards, v2)
			}

		}

		findStandard := false
		for _, v2 := range diffStandards {
			originalPrice := CalculatePriceRate(v2.Direct, v2.Labor, v2.Cost, 0, 0.0)

			iprice := global.Atoi(strings.ReplaceAll(price, ",", ""))
			if iprice == int(originalPrice) {
				item.Standard = v2.Id
				findStandard = true
				break
			}
		}

		if findStandard == false && len(diffStandards) > 0 {
			changeStandards = append(changeStandards, diffStandards[0])
			item.Standard = diffStandards[0].Id
		}

		if item.Standard == 0 {
			var n models.Standard

			n.Apt = id
			n.Category = category.Id
			n.Name = standard

			n.Direct = reversPrice(global.Atol(strings.ReplaceAll(price, ",", "")))
			n.Unit = unit

			newStandards = append(newStandards, n)
			//standardManager.Insert(&n)
			/*

				item.Standard = standardManager.GetIdentity()
				n.Id = item.Standard

				standards = append(standards, n)
			*/
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

		find := false
		findIndex := 0
		for i, v := range breakdowns {
			if item.Dong == v.Dong &&
				item.Method == v.Method &&
				item.Topcategory == v.Topcategory &&
				item.Subcategory == v.Subcategory &&
				item.Category == v.Category &&
				item.Count == v.Count &&
				item.Lastdate == v.Lastdate &&
				item.Duedate == v.Duedate &&
				item.Remark == v.Remark {

				if item.Standard == v.Standard {
					find = true
					findIndex = i
					break
				}

				for _, v2 := range diffStandards {
					if v.Standard == v2.Id {
						find = true
						findIndex = i
						break
					}
				}

				if find == true {
					break
				}
			}
		}

		if find == false {
			//breakdownManager.Insert(&item)
			newBreakdowns = append(newBreakdowns, item)
		} else {
			breakdowns = append(breakdowns[:findIndex], breakdowns[findIndex+1:]...)
		}

		pos++

		emptyCount = 0
	}

	if update == true {
		conn.Commit()
	}

	f.Close()

	//os.Remove(path.Join(config.UploadPath, filename))

	c.Set("changeStandard", changeStandards)
	c.Set("newStandard", newStandards)
	c.Set("newBreakdown", newBreakdowns)
	c.Set("remainBreakdown", breakdowns)
}

func DiffupdateProcess(diffs *models.Diff) {
	id := diffs.Id
	filename := diffs.Filename

	conn := models.NewConnection()
	defer conn.Close()

	conn.Begin()
	defer conn.Rollback()

	dongManager := models.NewDongManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)
	standardManager := models.NewStandardManager(conn)

	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}})
	categorys := categoryManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id")})
	standards := standardManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("s_order,s_id")})
	dongs := dongManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("d_order,d_id")})

	categoryMap := make(map[int64]models.Category)
	standardMap := make(map[int64]models.Standard)

	for _, v := range categorys {
		categoryMap[v.Id] = v
	}

	for _, v := range standards {
		standardMap[v.Id] = v
	}

	log.Println("filename", path.Join(config.UploadPath, filename))
	f, err := excelize.OpenFile(path.Join(config.UploadPath, filename))
	if err != nil {
		log.Println(err)
		return
	}

	pos := 160
	sheet := "통합세부자료"

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
	emptyCount := 1

	var topcategory models.Category
	var subcategory models.Category
	var category models.Category

	cell1 := "I"
	cell2 := "L"
	cell3 := "M"
	cell4 := "N"

	cellCount := "G"

	for i := 0; i < 5; i++ {
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

	log.Println("pos find = ", pos, cell1, cell2, cell3, cell4, cellCount)

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

					log.Println("price", iprice, int(originalPrice))

					find := false
					for _, v3 := range diffs.ChangeStandard {
						if v2.Name != v3.Name {
							continue
						}

						if v2.Category != v3.Category {
							continue
						}

						if v2.Direct != v3.Direct {
							continue
						}

						find = true
						break
					}

					if find == true {
						standardManager.Update(&v2)
					}
				}
				break
			}
		}

		diffStandards := make([]models.Standard, 0)
		for _, v2 := range standards {
			if v2.Apt != id {
				continue
			}

			if v2.Category != category.Id {
				continue
			}

			if v2.Name == standard {
				diffStandards = append(diffStandards, v2)
			}

		}

		findStandard := false
		for _, v2 := range diffStandards {
			originalPrice := CalculatePriceRate(v2.Direct, v2.Labor, v2.Cost, 0, 0.0)

			iprice := global.Atoi(strings.ReplaceAll(price, ",", ""))
			if iprice == int(originalPrice) {
				item.Standard = v2.Id
				findStandard = true
				break
			}
		}

		if findStandard == false && len(diffStandards) > 0 {
			v2 := diffStandards[0]

			find := false
			for _, v3 := range diffs.ChangeStandard {
				if v2.Name != v3.Name {
					continue
				}

				if v2.Category != v3.Category {
					continue
				}

				if v2.Direct != v3.Direct {
					continue
				}

				find = true
				break
			}

			if find == true {
				standardManager.Update(&v2)
			}

			item.Standard = diffStandards[0].Id
		}

		if item.Standard == 0 {
			var n models.Standard

			n.Apt = id
			n.Category = category.Id
			n.Name = standard

			n.Direct = reversPrice(global.Atol(strings.ReplaceAll(price, ",", "")))
			n.Unit = unit

			find := false
			for _, v3 := range diffs.ChangeStandard {
				if n.Name != v3.Name {
					continue
				}

				if n.Category != v3.Category {
					continue
				}

				if n.Direct != v3.Direct {
					continue
				}

				find = true
				break
			}

			if find == true {
				standardManager.Insert(&n)

				item.Standard = standardManager.GetIdentity()
				n.Id = item.Standard

				standards = append(standards, n)
			}
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

		find := false
		findIndex := 0
		for i, v := range breakdowns {
			if item.Dong == v.Dong &&
				item.Standard == v.Standard &&
				item.Method == v.Method &&
				item.Topcategory == v.Topcategory &&
				item.Subcategory == v.Subcategory &&
				item.Category == v.Category &&
				item.Count == v.Count &&
				item.Lastdate == v.Lastdate &&
				item.Duedate == v.Duedate &&
				item.Remark == v.Remark {

				if item.Standard == v.Standard {
					find = true
					findIndex = i
					break
				}

				for _, v2 := range diffStandards {
					if v.Standard == v2.Id {
						find = true
						findIndex = i
						break
					}
				}

				if find == true {
					break
				}
			}
		}

		if find == false {
			for _, v3 := range diffs.NewBreakdown {
				if item.Dong == v3.Dong &&
					item.Standard == v3.Standard &&
					item.Method == v3.Method &&
					item.Topcategory == v3.Topcategory &&
					item.Subcategory == v3.Subcategory &&
					item.Category == v3.Category &&
					item.Count == v3.Count &&
					item.Lastdate == v3.Lastdate &&
					item.Duedate == v3.Duedate &&
					item.Remark == v3.Remark {

					log.Println("insert", item)
					breakdownManager.Insert(&item)
					break
				}
			}
		} else {
			breakdowns = append(breakdowns[:findIndex], breakdowns[findIndex+1:]...)
		}

		pos++

		emptyCount = 0
	}

	conn.Commit()

	f.Close()

	os.Remove(path.Join(config.UploadPath, filename))
}

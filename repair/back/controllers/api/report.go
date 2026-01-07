package api

import (
	"fmt"
	"repair/controllers"
	"repair/global"
	"repair/models"
	"strings"
)

type ReportController struct {
	controllers.Controller
}

func (c *ReportController) Total(id int64) {
	conn := c.NewConnection()

	areaManager := models.NewAreaManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)

	areas := areaManager.FindByApt(id)
	breakdowns := breakdownManager.FindByApt(id)
	repair := repairManager.Get(id)

	var totalPrice int64 = 0
	var totalSaveprice int64 = 0
	for _, v := range breakdowns {
		standard := v.Extra["standard"].(models.Standard)
		category := v.Extra["category"].(models.Category)

		price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repair.Parcelrate), v.Count, float64(category.Percent))
		var saveprice int64 = 0

		if category.Cycle > 0 {
			saveprice = int64(float64(price) / float64(category.Cycle))
		}

		totalPrice += price
		totalSaveprice += saveprice
	}

	c.Result["price"] = totalPrice
	c.Result["saveprice"] = totalSaveprice

	var totalSize float64 = 0
	for _, v := range areas {
		totalSize += float64(v.Familycount) * float64(v.Size)
	}

	c.Result["totalsize"] = totalSize
}

type SummaryItem struct {
	Title          string `json:"title"`
	Method         string `json:"method"`
	Cycle          string `json:"cycle"`
	Percent        string `json:"percent"`
	Price          int64  `json:"price"`
	Saveprice      int64  `json:"saveprice"`
	Totalsaveprice int64  `json:"totalsaveprice"`
}

type Summarys struct {
	Items []SummaryItem `json:"items"`
}

func (c *ReportController) Summary(id int64) {
	conn := c.NewConnection()

	repairManager := models.NewRepairManager(conn)
	areaManager := models.NewAreaManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)

	repair := repairManager.Get(id)
	categorys := categoryManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id")})
	areas := areaManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("ar_order,ar_id")})
	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_dong,b_standard,b_method,b_id")})

	categoryMap := make(map[int64]models.Category)

	for _, v := range categorys {
		categoryMap[v.Id] = v
	}

	var areaTotal float64 = 0.0

	for _, v := range areas {
		size := float64(v.Size) * float64(v.Familycount)
		areaTotal += size
	}

	//startTop := 3
	topSummary := make(map[int64]TopSummary)
	summary := make(map[int64]Summary)

	totals := make([]SummaryItem, 0)
	items := make([]Summarys, 0)

	for i := 0; i < 7; i++ {
		var item Summarys
		item.Items = make([]SummaryItem, 0)

		items = append(items, item)
	}

	for _, v := range breakdowns {
		standard := v.Extra["standard"].(models.Standard)
		category := v.Extra["category"].(models.Category)

		price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repair.Parcelrate), v.Count, float64(category.Percent))
		var saveprice int64 = 0

		if category.Cycle > 0 {
			saveprice = int64(float64(price) / float64(category.Cycle))
		}

		if value, ok := topSummary[v.Topcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			topSummary[v.Topcategory] = value
		} else {
			var newValue TopSummary
			newValue.Price = price
			newValue.Saveprice = saveprice
			topSummary[v.Topcategory] = newValue
		}

		if value, ok := summary[v.Topcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			value.Cycle = append(value.Cycle, category.Cycle)
			value.Percent = append(value.Percent, category.Percent)
			if category.Percent == 100 {
				value.Method = append(value.Method, "전면")
			} else {
				value.Method = append(value.Method, "부분")
			}

			summary[v.Topcategory] = value
		} else {
			var newValue Summary
			newValue.Cycle = make([]int, 0)
			newValue.Percent = make([]int, 0)
			newValue.Method = make([]string, 0)

			newValue.Price += price
			newValue.Saveprice += saveprice
			newValue.Cycle = append(newValue.Cycle, category.Cycle)
			newValue.Percent = append(newValue.Percent, category.Percent)
			if category.Percent == 100 {
				newValue.Method = append(newValue.Method, "전면")
			} else {
				newValue.Method = append(newValue.Method, "부분")
			}

			summary[v.Topcategory] = newValue
		}

		if value, ok := summary[v.Subcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			value.Cycle = append(value.Cycle, category.Cycle)
			value.Percent = append(value.Percent, category.Percent)
			if category.Percent == 100 {
				value.Method = append(value.Method, "전면")
			} else {
				value.Method = append(value.Method, "부분")
			}

			summary[v.Subcategory] = value
		} else {
			var newValue Summary
			newValue.Cycle = make([]int, 0)
			newValue.Percent = make([]int, 0)
			newValue.Method = make([]string, 0)

			newValue.Price += price
			newValue.Saveprice += saveprice
			newValue.Cycle = append(newValue.Cycle, category.Cycle)
			newValue.Percent = append(newValue.Percent, category.Percent)
			if category.Percent == 100 {
				newValue.Method = append(newValue.Method, "전면")
			} else {
				newValue.Method = append(newValue.Method, "부분")
			}

			summary[v.Subcategory] = newValue
		}
	}

	pos := 0

	var topTotalPrice int64 = 0
	var topTotalSaveprice int64 = 0

	for _, v := range categorys {
		if v.Level != 1 {
			continue
		}

		var price int64 = 0
		var saveprice int64 = 0

		if value, ok := topSummary[v.Id]; ok {
			price = value.Price
			saveprice = value.Saveprice
		}

		topTotalPrice += price
		topTotalSaveprice += saveprice

		var item SummaryItem
		item.Price = price
		item.Saveprice = saveprice
		item.Totalsaveprice = saveprice * int64(GetPlanyears(repair.Planyears))
		item.Title = categoryMap[v.Id].Name
		totals = append(totals, item)

		/*
			f.SetCellDefault(sheet, GetCell("E", subtopPos[pos]), fmt.Sprintf("%v", price))
			f.SetCellDefault(sheet, GetCell("F", subtopPos[pos]), fmt.Sprintf("%v", saveprice))
		*/

		pos2 := 1

		for _, v2 := range categorys {
			typeid := 0

			if v2.Parent == v.Id && v2.Level == 2 {
				typeid = 2
			} else if v2.Id == v.Id {
				typeid = 1
			} else {
				continue
			}

			var price int64 = 0
			var saveprice int64 = 0
			cycle := ""
			percent := ""
			method := ""

			if value, ok := summary[v2.Id]; ok {
				price = value.Price
				saveprice = value.Saveprice

				cycles := global.Unique(value.Cycle)

				if len(cycles) == 1 {
					cycle = fmt.Sprintf("%v년", cycles[0])
				} else if len(cycles) == 2 {
					cycle = fmt.Sprintf("%v/%v년", cycles[0], cycles[1])
				} else {
					cycle = fmt.Sprintf("%v~%v년", cycles[0], cycles[len(cycles)-1])
				}

				percents := global.Unique(value.Percent)

				if len(percents) == 1 {
					percent = fmt.Sprintf("%v%%", percents[0])
				} else if len(percents) == 2 {
					percent = fmt.Sprintf("%v/%v%%", percents[0], percents[1])
				} else {
					percent = fmt.Sprintf("%v~%v%%", percents[0], percents[len(percents)-1])
				}

				methods := global.UniqueString(value.Method)

				if len(methods) == 1 {
					method = fmt.Sprintf("%v", methods[0])
				} else if len(methods) == 2 {
					method = fmt.Sprintf("%v/%v", methods[0], methods[1])
				}
			}

			var item SummaryItem
			if typeid == 2 {
				item.Method = method
			}
			item.Cycle = cycle
			item.Percent = percent
			item.Price = price
			item.Saveprice = saveprice

			if typeid == 1 {
				item.Title = categoryMap[v.Id].Name
				items[pos].Items = append([]SummaryItem{item}, items[pos].Items...)
			} else {
				item.Title = categoryMap[v2.Id].Name
				items[pos].Items = append(items[pos].Items, item)
			}

			if typeid == 2 {
				pos2++
			}
		}

		pos++

		if pos >= 7 {
			break
		}
	}

	c.Set("totals", totals)
	c.Set("items", items)

}

func (c *ReportController) Plan(id int64) {
	conn := c.NewConnection()

	repairManager := models.NewRepairManager(conn)
	areaManager := models.NewAreaManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)

	repair := repairManager.Get(id)

	categorys := categoryManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id")})
	areas := areaManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("ar_order,ar_id")})
	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_dong,b_standard,b_method,b_id")})

	categoryMap := make(map[int64]models.Category)

	for _, v := range categorys {
		categoryMap[v.Id] = v
	}

	var areaTotal float64 = 0.0

	for _, v := range areas {
		size := float64(v.Size) * float64(v.Familycount)
		areaTotal += size
	}

	if repair.Parcelrate != 0.0 && repair.Parcelrate != 100.0 {
		areaTotal *= float64(repair.Parcelrate) / 100.0
	}

	//startTop := 3
	topSummary := make(map[int64]TopSummary)
	summary := make(map[int64]Summary)

	items := make([]Summarys, 0)

	for i := 0; i < 7; i++ {
		var item Summarys
		item.Items = make([]SummaryItem, 0)

		items = append(items, item)
	}

	for _, v := range breakdowns {
		standard := v.Extra["standard"].(models.Standard)
		category := v.Extra["category"].(models.Category)

		price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repair.Parcelrate), v.Count, float64(category.Percent))
		var saveprice int64 = 0

		if category.Cycle > 0 {
			saveprice = int64(float64(price) / float64(category.Cycle))
		}

		if value, ok := topSummary[v.Topcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			topSummary[v.Topcategory] = value
		} else {
			var newValue TopSummary
			newValue.Price = price
			newValue.Saveprice = saveprice
			topSummary[v.Topcategory] = newValue
		}

		if value, ok := summary[v.Topcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			value.Cycle = append(value.Cycle, category.Cycle)
			value.Percent = append(value.Percent, category.Percent)
			if category.Percent == 100 {
				value.Method = append(value.Method, "전면")
			} else {
				value.Method = append(value.Method, "부분")
			}

			summary[v.Topcategory] = value
		} else {
			var newValue Summary
			newValue.Cycle = make([]int, 0)
			newValue.Percent = make([]int, 0)
			newValue.Method = make([]string, 0)

			newValue.Price += price
			newValue.Saveprice += saveprice
			newValue.Cycle = append(newValue.Cycle, category.Cycle)
			newValue.Percent = append(newValue.Percent, category.Percent)
			if category.Percent == 100 {
				newValue.Method = append(newValue.Method, "전면")
			} else {
				newValue.Method = append(newValue.Method, "부분")
			}

			summary[v.Topcategory] = newValue
		}

		if value, ok := summary[v.Subcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			value.Cycle = append(value.Cycle, category.Cycle)
			value.Percent = append(value.Percent, category.Percent)
			if category.Percent == 100 {
				value.Method = append(value.Method, "전면")
			} else {
				value.Method = append(value.Method, "부분")
			}

			summary[v.Subcategory] = value
		} else {
			var newValue Summary
			newValue.Cycle = make([]int, 0)
			newValue.Percent = make([]int, 0)
			newValue.Method = make([]string, 0)

			newValue.Price += price
			newValue.Saveprice += saveprice
			newValue.Cycle = append(newValue.Cycle, category.Cycle)
			newValue.Percent = append(newValue.Percent, category.Percent)
			if category.Percent == 100 {
				newValue.Method = append(newValue.Method, "전면")
			} else {
				newValue.Method = append(newValue.Method, "부분")
			}

			summary[v.Subcategory] = newValue
		}
	}

	pos := 0

	var topTotalPrice int64 = 0
	var topTotalSaveprice int64 = 0

	for _, v := range categorys {
		if v.Level != 1 {
			continue
		}

		var price int64 = 0
		var saveprice int64 = 0

		if value, ok := topSummary[v.Id]; ok {
			price = value.Price
			saveprice = value.Saveprice
		}

		topTotalPrice += price
		topTotalSaveprice += saveprice

		pos++

		if pos >= 7 {
			break
		}
	}

	c.Set("area", areaTotal)
	c.Set("total", topTotalSaveprice*int64(GetPlanyears(repair.Planyears)))

	c.Set("startyear", repair.Completionyear+1)

	temps := strings.Split(repair.Reportdate, "-")
	reportyear := temps[0]

	c.Set("reportyear", global.Atoi(reportyear))
}

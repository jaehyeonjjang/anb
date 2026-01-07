package api

import (
	"repair/controllers"
	"repair/models"
	"repair/repair"
)

type RepairController struct {
	controllers.Controller
}

func (c *RepairController) Lastdate(id int64) {
	conn := c.NewConnection()

	manager := models.NewRepairManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)

	item := manager.Get(id)

	repairs := manager.Find([]any{models.Where{Column: "apt", Value: item.Apt, Compare: "="}, models.Ordering("r_reportdate,r_id")})
	var last int64 = 0
	for _, v := range repairs {
		if v.Id == id {
			break
		}

		last = v.Id
	}

	prevLastdate := -1
	lastdate := -1
	if last > 0 {
		items := breakdownManager.FindByApt(last)

		for _, v := range items {
			if v.Lastdate > prevLastdate {
				prevLastdate = v.Lastdate
			}
		}
	}

	items := breakdownManager.FindByApt(id)

	for _, v := range items {
		if v.Lastdate > lastdate {
			lastdate = v.Lastdate
		}
	}

	if prevLastdate == -1 {
		c.Set("prev", "")
	} else {
		c.Set("prev", prevLastdate)
	}

	if lastdate == -1 {
		c.Set("current", "")
	} else {
		c.Set("current", lastdate)
	}
}

func (c *RepairController) Pre_Update(item *models.Repair) {
	conn := c.NewConnection()

	repairManager := models.NewRepairManager(conn)
	outlineManager := models.NewOutlineManager(conn)
	outlineplanManager := models.NewOutlineplanManager(conn)
	areaManager := models.NewAreaManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)

	repair := repairManager.Get(item.Id)

	if item.Parcelrate == repair.Parcelrate {
		return
	}

	outlines := outlineManager.Find([]any{
		models.Where{Column: "apt", Value: item.Id, Compare: "="},
		models.Ordering("o_id"),
	})

	outlineplans := outlineplanManager.Find([]any{
		models.Where{Column: "apt", Value: item.Id, Compare: "="},
		models.Ordering("op_id"),
	})

	areas := areaManager.Find([]any{models.Where{Column: "apt", Value: item.Id, Compare: "="}, models.Ordering("ar_order,ar_id")})

	var totalSize float64 = 0.0
	for _, v := range areas {
		totalSize += float64(v.Familycount) * float64(v.Size)
	}

	categorys := categoryManager.Find([]any{models.Where{Column: "apt", Value: item.Id, Compare: "="}, models.Ordering("c_order,c_id")})
	breakdowns := breakdownManager.Find([]any{models.Where{Column: "apt", Value: item.Id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_dong,b_standard,b_method,b_id")})

	var topTotalSaveprice int64 = 0

	topSummary := make(map[int64]TopSummary)
	summary := make(map[int64]Summary)

	for _, v := range breakdowns {
		standard := v.Extra["standard"].(models.Standard)
		category := v.Extra["category"].(models.Category)

		price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(item.Parcelrate), v.Count, float64(category.Percent))
		saveprice := int64(float64(price) / float64(category.Cycle))

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

	for _, v := range categorys {
		if v.Level != 1 {
			continue
		}

		var saveprice int64 = 0

		if value, ok := topSummary[v.Id]; ok {
			saveprice = value.Saveprice
		}

		topTotalSaveprice += saveprice
	}

	for _, v := range outlines {
		totalprice := float64(topTotalSaveprice) * float64(GetPlanyears(repair.Planyears)) / 100.0 * float64(v.Rate)

		duration := getDurationMonth(v.Endyear, v.Endmonth, v.Startyear, v.Startmonth)
		price := totalprice / (totalSize * float64(duration))

		v.Price = models.Double(price)
		outlineManager.Update(&v)
	}

	for _, v := range outlineplans {
		totalprice := float64(topTotalSaveprice) * float64(GetPlanyears(repair.Planyears)) / 100.0 * float64(v.Rate)

		duration := getDurationMonth(v.Endyear, v.Endmonth, v.Startyear, v.Startmonth)
		price := totalprice / (totalSize * float64(duration))

		v.Price = models.Double(price)
		outlineplanManager.Update(&v)
	}
}

// @POST()
func (c *RepairController) Change(id int64) {
	repair.Change(id)
}

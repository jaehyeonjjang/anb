package api

import (
	"log"
	"repair/controllers"
	"repair/models"
)

type BreakdownController struct {
	controllers.Controller
}

// @Post()
func (c *BreakdownController) Deduplication(apt int64) {
	conn := c.NewConnection()

	manager := models.NewBreakdownManager(conn)
	items := manager.FindByApt(apt)

	for _, item := range items {
		if item.Id == 0 {
			continue
		}

		for i, v := range items {
			if v.Id == 0 {
				continue
			}

			if item.Id == v.Id {
				continue
			}

			if item.Topcategory == v.Topcategory &&
				item.Subcategory == v.Subcategory &&
				item.Category == v.Category &&
				item.Method == v.Method &&
				item.Count == v.Count &&
				item.Lastdate == v.Lastdate &&
				item.Duedate == v.Duedate &&
				item.Elevator == v.Elevator &&
				item.Dong == v.Dong &&
				item.Standard == v.Standard {

				manager.Delete(v.Id)

				items[i].Id = 0
			}
		}
	}
}

func (c *BreakdownController) Pre_Deletebatch(items *[]models.Breakdown) {
	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewBreakdownhistoryManager(conn)

	for _, item := range *items {
		old := breakdownManager.Get(item.Id)
		repair := repairManager.Get(old.Apt)

		if repair.Type != 2 {
			continue
		}

		history := manager.GetByBreakdown(item.Id)

		if history != nil && history.Id != 0 {
			if history.Type&8 == 0 {
				history.Type += 8
				manager.Update(history)
			}
		} else {
			var n models.Breakdownhistory
			n.Id = old.Id
			n.Topcategory = old.Topcategory
			n.Subcategory = old.Subcategory
			n.Category = old.Category
			n.Method = old.Method
			n.Count = old.Count
			n.Lastdate = old.Lastdate
			n.Duedate = old.Duedate
			n.Remark = old.Remark
			n.Elevator = old.Elevator
			n.Percent = old.Percent
			n.Rate = old.Rate
			n.Dong = old.Dong
			n.Standard = old.Standard
			n.Breakdown = old.Id
			n.Apt = old.Apt
			n.Type = 8

			manager.Insert(&n)
		}
	}
}

func (c *BreakdownController) Pre_Delete(item *models.Breakdown) {
	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewBreakdownhistoryManager(conn)

	old := breakdownManager.Get(item.Id)
	repair := repairManager.Get(old.Apt)

	if repair.Type != 2 {
		return
	}

	history := manager.GetByBreakdown(item.Id)

	if history != nil && history.Id != 0 {
		if history.Type&8 == 0 {
			history.Type += 8
			manager.Update(history)
		}
	} else {
		var n models.Breakdownhistory
		n.Id = old.Id
		n.Topcategory = old.Topcategory
		n.Subcategory = old.Subcategory
		n.Category = old.Category
		n.Method = old.Method
		n.Count = old.Count
		n.Lastdate = old.Lastdate
		n.Duedate = old.Duedate
		n.Remark = old.Remark
		n.Elevator = old.Elevator
		n.Percent = old.Percent
		n.Rate = old.Rate
		n.Dong = old.Dong
		n.Standard = old.Standard
		n.Breakdown = old.Id
		n.Apt = old.Apt
		n.Type = 8

		manager.Insert(&n)
	}
}

func (c *BreakdownController) Pre_Insertbatch(items *[]models.Breakdown) {
	if items == nil || len(*items) == 0 {
		return
	}

	conn := c.NewConnection()

	adjustManager := models.NewAdjustManager(conn)
	categoryManager := models.NewCategoryManager(conn)

	adjusts := adjustManager.Find([]interface{}{models.Where{Column: "apt", Value: (*items)[0].Apt, Compare: "="}, models.Ordering("aj_order,aj_id")})

	for i, item := range *items {
		var rate float64 = 0
		for _, v := range adjusts {
			if v.Category != 0 {
				category := categoryManager.Get(v.Category)

				if category.Level == 1 {
					if item.Topcategory != category.Id {
						continue
					}
				} else if category.Level == 2 {
					if item.Subcategory != category.Id {
						continue
					}
				} else if category.Level == 3 {
					if item.Category != category.Id {
						continue
					}
				}
			}

			if v.Standard != 0 {
				if item.Standard != v.Standard {
					continue
				}
			}

			rate = float64(v.Rate)
		}

		(*items)[i].Rate = models.Double(rate)
	}
}

func (c *BreakdownController) Pre_Insert(item *models.Breakdown) {
	conn := c.NewConnection()

	adjustManager := models.NewAdjustManager(conn)
	categoryManager := models.NewCategoryManager(conn)

	adjusts := adjustManager.Find([]interface{}{models.Where{Column: "apt", Value: item.Apt, Compare: "="}, models.Ordering("aj_order,aj_id")})

	var rate float64 = 0
	for _, v := range adjusts {
		if v.Category != 0 {
			category := categoryManager.Get(v.Category)

			if category.Level == 1 {
				if item.Topcategory != category.Id {
					continue
				}
			} else if category.Level == 2 {
				if item.Subcategory != category.Id {
					continue
				}
			} else if category.Level == 3 {
				if item.Category != category.Id {
					continue
				}
			}
		}

		if v.Standard != 0 {
			if item.Standard != v.Standard {
				continue
			}
		}

		rate = float64(v.Rate)
	}

	item.Rate = models.Double(rate)
}

func (c *BreakdownController) Pre_Update(item *models.Breakdown) {
	log.Println("Pre_Update")
	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewBreakdownhistoryManager(conn)
	standardhistoryManager := models.NewStandardhistoryManager(conn)
	adjustManager := models.NewAdjustManager(conn)
	categoryManager := models.NewCategoryManager(conn)

	repair := repairManager.Get(item.Apt)

	adjusts := adjustManager.Find([]interface{}{models.Where{Column: "apt", Value: item.Apt, Compare: "="}, models.Ordering("aj_order,aj_id")})

	var rate float64 = 0
	for _, v := range adjusts {
		if v.Category != 0 {
			category := categoryManager.Get(v.Category)

			if category.Level == 1 {
				if item.Topcategory != category.Id {
					continue
				}
			} else if category.Level == 2 {
				if item.Subcategory != category.Id {
					continue
				}
			} else if category.Level == 3 {
				if item.Category != category.Id {
					continue
				}
			}
		}

		if v.Standard != 0 {
			if item.Standard != v.Standard {
				continue
			}
		}

		rate = float64(v.Rate)
	}

	item.Rate = models.Double(rate)

	if repair.Type != 2 {
		return
	}

	old := breakdownManager.Get(item.Id)

	history := manager.GetByBreakdown(item.Id)

	if history != nil && history.Id != 0 {
		old.Topcategory = history.Topcategory
		old.Subcategory = history.Subcategory
		old.Category = history.Category
		old.Method = history.Method
		old.Count = history.Count
		old.Lastdate = history.Lastdate
		old.Duedate = history.Duedate
		old.Remark = history.Remark
		old.Elevator = history.Elevator
		old.Percent = history.Percent
		old.Rate = history.Rate
		old.Dong = history.Dong
		old.Standard = history.Standard
	}

	if old.Count == item.Count && old.Duedate == item.Duedate && old.Standard == item.Standard {
		if history != nil && history.Id != 0 {
			if history.Type < 16 {
				manager.Delete(history.Id)
			} else {
				history.Type = 16
				manager.Update(history)
			}
		}
		return
	}

	typeid := 0
	if old.Count != item.Count {
		typeid += 1
	}

	if old.Duedate != item.Duedate {
		typeid += 2
	}

	if old.Standard != item.Standard {
		typeid += 4
	}

	standardhistory := standardhistoryManager.GetByStandard(item.Standard)

	if standardhistory != nil && standardhistory.Id != 0 {
		typeid += 16
	}

	if history != nil && history.Id != 0 {
		history.Type = typeid
		history.Originalduedate = item.Duedate
		manager.Update(history)
	} else {
		var n models.Breakdownhistory
		n.Id = old.Id
		n.Topcategory = old.Topcategory
		n.Subcategory = old.Subcategory
		n.Category = old.Category
		n.Method = old.Method
		n.Count = old.Count
		n.Lastdate = old.Lastdate
		n.Duedate = old.Duedate
		n.Remark = old.Remark
		n.Elevator = old.Elevator
		n.Percent = old.Percent
		n.Rate = old.Rate
		n.Dong = old.Dong
		n.Standard = old.Standard
		n.Breakdown = old.Id
		n.Apt = old.Apt
		n.Type = typeid
		n.Originalcount = item.Count
		n.Originalduedate = item.Duedate

		manager.Insert(&n)

		history = &n
	}
}

func (c *BreakdownController) Pre_UpdateDuedateById(duedate int, id int64) {
	log.Println("Pre_UpdateDuedateById")

	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewBreakdownhistoryManager(conn)
	standardhistoryManager := models.NewStandardhistoryManager(conn)

	item := breakdownManager.Get(id)
	item.Duedate = duedate

	repair := repairManager.Get(item.Apt)

	if repair.Type != 2 {
		return
	}

	old := breakdownManager.Get(item.Id)

	history := manager.GetByBreakdown(item.Id)

	if history != nil && history.Id != 0 {
		old.Topcategory = history.Topcategory
		old.Subcategory = history.Subcategory
		old.Category = history.Category
		old.Method = history.Method
		old.Count = history.Count
		old.Lastdate = history.Lastdate
		old.Duedate = history.Duedate
		old.Remark = history.Remark
		old.Elevator = history.Elevator
		old.Percent = history.Percent
		old.Rate = history.Rate
		old.Dong = history.Dong
		old.Standard = history.Standard
	}

	if old.Count == item.Count && old.Duedate == item.Duedate && old.Standard == item.Standard {
		if history != nil && history.Id != 0 {
			if history.Type < 16 {
				manager.Delete(history.Id)
			} else {
				history.Type = 16
				manager.Update(history)
			}
		}
		return
	}

	typeid := 0
	if old.Count != item.Count {
		typeid += 1
	}

	if old.Duedate != item.Duedate {
		typeid += 2
	}

	if old.Standard != item.Standard {
		typeid += 4
	}

	standardhistory := standardhistoryManager.GetByStandard(item.Standard)

	if standardhistory != nil && standardhistory.Id != 0 {
		typeid += 16
	}

	if history != nil && history.Id != 0 {
		history.Type = typeid
		history.Originalduedate = item.Duedate
		manager.Update(history)
	} else {
		var n models.Breakdownhistory
		n.Id = old.Id
		n.Topcategory = old.Topcategory
		n.Subcategory = old.Subcategory
		n.Category = old.Category
		n.Method = old.Method
		n.Count = old.Count
		n.Lastdate = old.Lastdate
		n.Duedate = old.Duedate
		n.Remark = old.Remark
		n.Elevator = old.Elevator
		n.Percent = old.Percent
		n.Rate = old.Rate
		n.Dong = old.Dong
		n.Standard = old.Standard
		n.Breakdown = old.Id
		n.Apt = old.Apt
		n.Type = typeid
		n.Originalcount = item.Count
		n.Originalduedate = item.Duedate

		manager.Insert(&n)

		history = &n
	}
}

// @Post()
func (c *BreakdownController) UpdateLastdate(date int, ids []int64) {
	log.Println(date)
	log.Println(ids)

	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)

	for _, v := range ids {
		breakdownManager.UpdateLastdateById(date, v)
	}
}

// @Post()
func (c *BreakdownController) UpdateDuedate(apt int64, date int, ids []int64) {
	conn := c.NewConnection()

	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewBreakdownhistoryManager(conn)
	standardhistoryManager := models.NewStandardhistoryManager(conn)

	repair := repairManager.Get(apt)

	for _, id := range ids {
		old := breakdownManager.Get(id)

		item := breakdownManager.Get(id)
		item.Duedate = date
		breakdownManager.Update(item)

		if repair.Type != 2 {
			continue
		}

		history := manager.GetByBreakdown(item.Id)

		if history != nil && history.Id != 0 {
			old.Topcategory = history.Topcategory
			old.Subcategory = history.Subcategory
			old.Category = history.Category
			old.Method = history.Method
			old.Count = history.Count
			old.Lastdate = history.Lastdate
			old.Duedate = history.Duedate
			old.Remark = history.Remark
			old.Elevator = history.Elevator
			old.Percent = history.Percent
			old.Rate = history.Rate
			old.Dong = history.Dong
			old.Standard = history.Standard
		}

		if old.Count == item.Count && old.Duedate == item.Duedate && old.Standard == item.Standard {
			if history != nil && history.Id != 0 {
				if history.Type < 16 {
					manager.Delete(history.Id)
				} else {
					history.Type = 16
					manager.Update(history)
				}
			}
			continue
		}

		typeid := 0
		if old.Count != item.Count {
			typeid += 1
		}

		if old.Duedate != item.Duedate {
			typeid += 2
		}

		if old.Standard != item.Standard {
			typeid += 4
		}

		standardhistory := standardhistoryManager.GetByStandard(item.Standard)

		if standardhistory != nil && standardhistory.Id != 0 {
			typeid += 16
		}

		if history != nil && history.Id != 0 {
			history.Type = typeid
			history.Originalduedate = item.Duedate
			manager.Update(history)
		} else {
			var n models.Breakdownhistory
			n.Id = old.Id
			n.Topcategory = old.Topcategory
			n.Subcategory = old.Subcategory
			n.Category = old.Category
			n.Method = old.Method
			n.Count = old.Count
			n.Lastdate = old.Lastdate
			n.Duedate = old.Duedate
			n.Remark = old.Remark
			n.Elevator = old.Elevator
			n.Percent = old.Percent
			n.Rate = old.Rate
			n.Dong = old.Dong
			n.Standard = old.Standard
			n.Breakdown = old.Id
			n.Apt = old.Apt
			n.Type = typeid
			n.Originalcount = item.Count
			n.Originalduedate = item.Duedate

			manager.Insert(&n)

			history = &n
		}
	}
}

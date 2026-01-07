package api

import (
	"repair/controllers"
	"repair/models"
)

type StandardController struct {
	controllers.Controller
}

// @Post()
func (c *StandardController) Insertall(item *models.Standard) {
	if item == nil {
		return
	}

	conn := c.NewConnection()

	categoryManager := models.NewCategoryManager(conn)
	manager := models.NewStandardManager(conn)

	originalCategory := categoryManager.Get(item.Category)
	originalSubcategory := categoryManager.Get(originalCategory.Parent)
	originalTopcategory := categoryManager.Get(originalSubcategory.Parent)

	var originalId int64 = 0

	apt := item.Apt
	category := item.Category

	topcategory := categoryManager.GetByAptLevelParentName(-1, 1, 0, originalTopcategory.Name)
	if topcategory == nil || topcategory.Id == 0 {
	} else {
		subcategory := categoryManager.GetByAptLevelParentName(-1, 2, topcategory.Id, originalSubcategory.Name)
		if subcategory == nil || subcategory.Id == 0 {
		} else {
			category := categoryManager.GetByAptLevelParentName(-1, 3, subcategory.Id, originalCategory.Name)
			if category == nil || category.Id == 0 {
			} else {
				item.Id = 0
				item.Apt = -1
				item.Category = category.Id

				manager.Insert(item)

				originalId = manager.GetIdentity()
			}
		}
	}

	if apt != -1 {
		item.Id = 0
		item.Apt = apt
		item.Category = category
		item.Original = originalId
		manager.Insert(item)
	}

	/*
		id := manager.GetIdentity()
		c.Result["id"] = id
		item.Id = id
	*/

	go func() {
		conn := models.NewConnection()
		defer conn.Close()

		categoryManager := models.NewCategoryManager(conn)
		repairManager := models.NewRepairManager(conn)
		manager := models.NewStandardManager(conn)

		var args []interface{}
		args = append(args, models.Where{Column: "type", Value: 1, Compare: "="})
		repairs := repairManager.Find(args)

		for _, v := range repairs {
			if v.Id == apt {
				continue
			}

			topcategory := categoryManager.GetByAptLevelParentName(v.Id, 1, 0, originalTopcategory.Name)

			if topcategory == nil || topcategory.Id == 0 {
				continue
			}

			subcategory := categoryManager.GetByAptLevelParentName(v.Id, 2, topcategory.Id, originalSubcategory.Name)

			if subcategory == nil || subcategory.Id == 0 {
				continue
			}

			category := categoryManager.GetByAptLevelParentName(v.Id, 3, subcategory.Id, originalCategory.Name)

			if category == nil || category.Id == 0 {
				continue
			}

			item.Id = 0
			item.Apt = v.Id
			item.Category = category.Id

			manager.Insert(item)
		}
	}()
}

// @Put()
func (c *StandardController) Updateall(item *models.Standard) {
	if item == nil {
		return
	}

	conn := c.NewConnection()

	categoryManager := models.NewCategoryManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewStandardManager(conn)

	originalCategory := categoryManager.Get(item.Category)
	originalSubcategory := categoryManager.Get(originalCategory.Parent)
	originalTopcategory := categoryManager.Get(originalSubcategory.Parent)

	old := manager.Get(item.Id)

	c.Pre_Update(item)

	manager.Update(item)

	apt := item.Apt

	if old.Original != 0 {
		items := manager.Find([]any{models.Where{Column: "original", Value: old.Original, Compare: "="}})

		for _, newItem := range items {
			newItem.Direct = item.Direct
			newItem.Labor = item.Labor
			newItem.Cost = item.Cost
			newItem.Unit = item.Unit

			manager.Update(&newItem)
		}
	} else if old.Name == item.Name {
		topcategory := categoryManager.GetByAptLevelParentName(-1, 1, 0, originalTopcategory.Name)
		if topcategory == nil || topcategory.Id == 0 {
		} else {
			subcategory := categoryManager.GetByAptLevelParentName(-1, 2, topcategory.Id, originalSubcategory.Name)
			if subcategory == nil || subcategory.Id == 0 {
			} else {
				category := categoryManager.GetByAptLevelParentName(-1, 3, subcategory.Id, originalCategory.Name)
				if category == nil || category.Id == 0 {
				} else {
					newItem := manager.GetByAptCategoryName(-1, category.Id, item.Name)

					if newItem != nil {
						newItem.Direct = item.Direct
						newItem.Labor = item.Labor
						newItem.Cost = item.Cost
						newItem.Unit = item.Unit

						manager.Update(newItem)
					}
				}
			}
		}

		repairs := repairManager.Find(nil)

		for _, v := range repairs {
			if v.Id == apt {
				continue
			}

			topcategory := categoryManager.GetByAptLevelParentName(v.Id, 1, 0, originalTopcategory.Name)
			if topcategory == nil || topcategory.Id == 0 {
				continue
			}

			subcategory := categoryManager.GetByAptLevelParentName(v.Id, 2, topcategory.Id, originalSubcategory.Name)
			if subcategory == nil || subcategory.Id == 0 {
				continue
			}

			category := categoryManager.GetByAptLevelParentName(v.Id, 3, subcategory.Id, originalCategory.Name)
			if category == nil || category.Id == 0 {
				continue
			}

			newItem := manager.GetByAptCategoryName(-1, category.Id, item.Name)

			if newItem != nil {
				newItem.Direct = item.Direct
				newItem.Labor = item.Labor
				newItem.Cost = item.Cost
				newItem.Unit = item.Unit

				manager.Update(newItem)
			}
		}
	}
}

func (c *StandardController) Pre_Update(item *models.Standard) {
	conn := c.NewConnection()

	standardbackupManager := models.NewStandardbackupManager(conn)
	standardManager := models.NewStandardManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)
	repairManager := models.NewRepairManager(conn)
	manager := models.NewBreakdownhistoryManager(conn)
	standardhistoryManager := models.NewStandardhistoryManager(conn)

	repair := repairManager.Get(item.Apt)

	if repair.Type != 2 {
		return
	}

	old := standardManager.Get(item.Id)

	standardbackup := models.Standardbackup{}
	standardbackup.Standard = old.Id
	standardbackup.Name = old.Name
	standardbackup.Direct = old.Direct
	standardbackup.Labor = old.Labor
	standardbackup.Cost = old.Cost
	standardbackup.Unit = old.Unit
	standardbackup.Order = old.Order
	standardbackup.Original = old.Original
	standardbackup.Category = old.Category
	standardbackup.Apt = old.Apt
	standardbackupManager.Insert(&standardbackup)

	history := standardhistoryManager.GetByStandard(item.Id)

	if history != nil && history.Id != 0 {
		old.Direct = int64(history.Direct)
		old.Labor = history.Labor
		old.Cost = history.Cost
	}

	if item.Direct == old.Direct && item.Labor == old.Labor && item.Cost == old.Cost {
		if history != nil && history.Id != 0 {
			standardhistoryManager.Delete(history.Id)

			historys := manager.Find([]interface{}{models.Where{Column: "apt", Value: item.Apt, Compare: "="}, models.Where{Column: "standard", Value: item.Id, Compare: "="}})

			for _, v := range historys {
				if v.Type < 16 {
					continue
				}

				v.Type -= 16
				manager.Update(&v)
			}
		}

		return
	}

	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: item.Apt, Compare: "="}, models.Where{Column: "standard", Value: item.Id, Compare: "="}})
	historys := manager.Find([]interface{}{models.Where{Column: "apt", Value: item.Apt, Compare: "="}, models.Where{Column: "standard", Value: item.Id, Compare: "="}})
	for _, breakdown := range breakdowns {
		flag := false
		for _, v := range historys {
			if breakdown.Id == v.Breakdown {
				if v.Type >= 16 {
					continue
				}

				v.Type += 16
				manager.Update(&v)

				flag = true
				break
			}
		}

		if flag == false {
			old := breakdown
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
			n.Type = 16
			n.Originalcount = breakdown.Count
			n.Originalduedate = breakdown.Duedate

			manager.Insert(&n)
		}
	}

	if history != nil && history.Id != 0 {
		return
	}

	var n models.Standardhistory

	n.Name = old.Name
	n.Direct = int(old.Direct)
	n.Labor = old.Labor
	n.Cost = old.Cost
	n.Unit = old.Unit
	n.Order = old.Order
	n.Original = old.Original
	n.Category = old.Category
	n.Standard = item.Id
	n.Apt = old.Apt

	standardhistoryManager.Insert(&n)

}

func (c *StandardController) Pre_Delete(item *models.Standard) {
	conn := c.NewConnection()

	standardManager := models.NewStandardManager(conn)
	standardbackupManager := models.NewStandardbackupManager(conn)

	old := standardManager.Get(item.Id)

	standardbackup := models.Standardbackup{}
	standardbackup.Standard = old.Id
	standardbackup.Name = old.Name
	standardbackup.Direct = old.Direct
	standardbackup.Labor = old.Labor
	standardbackup.Cost = old.Cost
	standardbackup.Unit = old.Unit
	standardbackup.Order = old.Order
	standardbackup.Original = old.Original
	standardbackup.Category = old.Category
	standardbackup.Apt = old.Apt
	standardbackupManager.Insert(&standardbackup)
}

func (c *StandardController) Pre_Deletebatch(items *[]models.Standard) {
	conn := c.NewConnection()

	standardManager := models.NewStandardManager(conn)
	standardbackupManager := models.NewStandardbackupManager(conn)

	for _, v := range *items {
		old := standardManager.Get(v.Id)

		standardbackup := models.Standardbackup{}
		standardbackup.Standard = old.Id
		standardbackup.Name = old.Name
		standardbackup.Direct = old.Direct
		standardbackup.Labor = old.Labor
		standardbackup.Cost = old.Cost
		standardbackup.Unit = old.Unit
		standardbackup.Order = old.Order
		standardbackup.Original = old.Original
		standardbackup.Category = old.Category
		standardbackup.Apt = old.Apt
		standardbackupManager.Insert(&standardbackup)
	}
}

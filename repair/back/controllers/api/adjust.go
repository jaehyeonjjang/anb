package api

import (
	"repair/controllers"
	"repair/models"
)

type AdjustController struct {
	controllers.Controller
}

func (c *AdjustController) Post_Insert(item *models.Adjust) {
	updatePrice(item)
}

func (c *AdjustController) Post_Update(item *models.Adjust) {
	updatePrice(item)
}

func (c *AdjustController) Post_Delete(item *models.Adjust) {
	updatePrice(item)
}

func (c *AdjustController) Post_Deletebatch(item *[]models.Adjust) {
	if len(*item) == 0 {
		return
	}

	updatePrice(&(*item)[0])
}

func updatePrice(item *models.Adjust) {
	id := item.Apt

	conn := models.NewConnection()
	defer conn.Close()

	manager := models.NewAdjustManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)

	adjusts := manager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("aj_order,aj_id")})
	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_dong,b_standard,b_method,b_id")})

	conn.Begin()
	defer conn.Rollback()

	for _, v := range breakdowns {
		var rate float64 = 0
		flag := false
		for _, adjust := range adjusts {
			if adjust.Standard > 0 {
				if v.Standard != adjust.Standard {
					continue
				}
			} else if adjust.Category != 0 {
				category := categoryManager.Get(adjust.Category)
				if category.Level == 1 {
					if v.Topcategory != category.Id {
						continue
					}
				} else if category.Level == 2 {
					if v.Subcategory != category.Id {
						continue
					}
				} else if category.Level == 3 {
					if v.Category != category.Id {
						continue
					}
				}
			}

			rate = float64(adjust.Rate)
			flag = true
		}

		if flag == true {
			if v.Rate != models.Double(rate) {
				v.Rate = models.Double(rate)
				breakdownManager.Update(&v)
			}
		} else {
			if v.Rate != 0.0 {
				v.Rate = models.Double(0.0)
				breakdownManager.Update(&v)
			}
		}
	}

	conn.Commit()
}

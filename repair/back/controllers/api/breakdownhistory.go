package api

import (
	"fmt"
	"repair/controllers"
	"repair/models"
	"sort"
	"strings"
)

type BreakdownhistoryController struct {
	controllers.Controller
}

func (c *BreakdownhistoryController) Auto(id int64) {
	items := MakeReview(id)
	c.Set("items", items)
}

func MakeReview(id int64) []models.Review {
	conn := models.NewConnection()
	defer conn.Close()

	manager := models.NewBreakdownhistoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	standardManager := models.NewStandardManager(conn)
	repairManager := models.NewRepairManager(conn)

	repair := repairManager.Get(id)

	rets := make([]models.Review, 0)

	if repair.Type != 2 {
		return rets
	}

	historys := manager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("bh_id")})
	breakdowns := breakdownManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_dong,b_standard,b_method,b_id")})

	alreadyDelete := make(map[int64]models.Breakdownhistory)
	already := make(map[string]models.Breakdownhistory)

	items := make([]models.Breakdownhistory, 0)

	for _, v := range historys {
		if v.Type == 8 {
			if _, ok := alreadyDelete[v.Standard]; ok {
				continue
			}

			alreadyDelete[v.Standard] = v

			continue
		}

		key := fmt.Sprintf("%v-%v", v.Standard, v.Method)

		if item, ok := already[key]; ok {
			v.Type = item.Type | v.Type
		}

		already[key] = v
	}

	for _, v := range breakdowns {
		key := fmt.Sprintf("%v-%v", v.Standard, v.Method)

		if _, ok := already[key]; ok {
			continue
		}

		if _, ok := alreadyDelete[v.Standard]; ok {
			continue
		}

		var item models.Breakdownhistory
		item.Id = 0
		item.Type = 0
		item.Topcategory = v.Topcategory
		item.Subcategory = v.Subcategory
		item.Category = v.Category
		item.Method = v.Method
		item.Count = v.Count
		item.Lastdate = v.Lastdate
		item.Duedate = v.Duedate
		item.Remark = v.Remark
		item.Elevator = v.Elevator
		item.Percent = v.Percent
		item.Rate = v.Rate
		item.Standard = v.Standard
		item.Breakdown = v.Id
		item.Originalduedate = v.Duedate
		item.Apt = v.Apt
		item.Date = v.Date

		category := categoryManager.Get(v.Method)
		standard := standardManager.Get(v.Standard)
		item.Extra = map[string]interface{}{
			"category": *category,
			"standard": *standard,
		}
		already[key] = item
	}

	for _, v := range already {
		var args []interface{}
		args = append(args, models.Where{Column: "apt", Value: id, Compare: "="})
		if v.Standard > 0 {
			args = append(args, models.Where{Column: "standard", Value: v.Standard, Compare: "="})
		}

		if v.Method > 0 {
			args = append(args, models.Where{Column: "method", Value: v.Method, Compare: "="})
		}

		breakdowns := breakdownManager.Find(args)

		var rate models.Double

		count := 0
		var price int64 = 0
		for _, breakdown := range breakdowns {
			count += breakdown.Count
			standard := breakdown.Extra["standard"].(models.Standard)

			price += CalculatePriceRate(standard.Direct, standard.Labor, standard.Cost, float64(breakdown.Rate), float64(repair.Parcelrate))

			rate = breakdown.Rate
		}

		v.Totalcount = count
		v.Totalprice = price
		v.Rate = rate

		items = append(items, v)
	}

	for _, v := range alreadyDelete {
		items = append(items, v)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Method < items[j].Method
	})

	for _, item := range items {
		count := false
		duedate := false
		standard := false
		deleted := false
		price := false

		if item.Type&1 > 0 {
			count = true
		}

		if item.Type&2 > 0 {
			duedate = true
		}

		if item.Type&4 > 0 {
			standard = true
		}

		if item.Type&8 > 0 {
			deleted = true
		}

		if item.Type&16 > 0 {
			price = true
		}

		content := ""
		adjust := ""

		if deleted == true {
			content = "우리단지 필요없는 항목으로 검토"
			adjust = "우리단지에 필요없는 항목으로 검토되어 삭제하는 것으로 조정"
		} else {
			if duedate == true {
				title := ""
				part := ""

				if count == true || standard == true || price == true {
					title += " 및 "

					subtitle := make([]string, 0)
					if count == true {
						subtitle = append(subtitle, "수량(물량)")
					}

					if standard == true {
						subtitle = append(subtitle, "규격(용량)")
					}

					if price == true {
						subtitle = append(subtitle, "금액")
					}

					title += strings.Join(subtitle, ", ")
					part = "이"
				} else {
					part = "가"
				}

				content = fmt.Sprintf("수선예정년도%v%v 현실에 맞지 않아 조정이 필요함", title, part)
				adjust = fmt.Sprintf("수선예정년도(%v)%v을 현실에 맞게 조정함", item.Originalduedate, title)
			} else {
				if count == true {
					if standard == true {
						if price == true {
							content = "규격(용량) 및 금액, 수량(물량)이 현실에 맞지 않는 관계로 조정이 필요함"
							adjust = "도면 및 견적을 참조하여 새롭게 규격(용량) 및 금액, 수량(물량)을 산출함"
						} else {
							content = "규격(용량) 및 수량(물량)이 현실에 맞지 않는 관계로 조정이 필요함"
							adjust = "도면을 참조하여 새롭게 규격(용량) 및 수량(물량)을 산출함"
						}
					} else {
						if price == true {
							content = "금액 및 수량(물량)이 현실에 맞지 않는 관계로 조정이 필요하다고 검토"
							adjust = "도면 및 견적을 참조하여 새롭게 금액, 수량(물량)을 산출하는 것으로 조정"
						} else {
							content = "수량(물량)이 현실에 맞지 않는 관계로 조정이 필요하다고 검토"
							adjust = "도면을 참조하여 새롭게 수량(물량)을 산출하는 것으로 조정"
						}
					}
				} else {
					if standard == true {
						if price == true {
							content = "규격(용량) 및 금액이 현실에 맞지 않는 관계로 조정이 필요함"
							adjust = "도면 및 견적을 참조하여 새롭게 규격(용량) 및 금액을 산출함"
						} else {
							content = "규격(용량)이 현실에 맞지 않는 관계로 조정이 필요함"
							adjust = "도면을 참조하여 새롭게 규격(용량)을 산출함"
						}
					} else {
						if price == true {
							content = "금액이 현실에 맞지 않는 관계로 조정이 필요하다고 검토"
							adjust = "견적 및 물가정보를 참조하여 새롭게 금액을 산출하는 것으로 조정"
						} else {
							content = "기존의 계획서대로 진행하기로 하여 검토사유가 없음"
							adjust = fmt.Sprintf("기존의 계획서(%v)대로 검토되어 조정사유가 없는 것으로 조정.", item.Originalduedate)
						}
					}
				}
			}
		}

		var review models.Review

		review.Id = 0

		method := item.Extra["category"].(models.Category)
		s := item.Extra["standard"].(models.Standard)

		review.Topcategory = item.Topcategory
		review.Subcategory = item.Subcategory
		review.Category = item.Category
		review.Standard = item.Standard
		review.Method = item.Method
		review.Content = content
		review.Adjust = adjust
		review.Cycle = fmt.Sprintf("%v", method.Cycle)
		review.Percent = method.Percent
		review.Count = item.Totalcount
		review.Price = CalculateRepair(s.Direct, s.Labor, s.Cost, float64(item.Rate), float64(repair.Parcelrate), item.Totalcount, float64(method.Percent))
		review.Extra = item.Extra

		rets = append(rets, review)
	}

	return rets
}

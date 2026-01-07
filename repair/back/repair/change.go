package repair

import (
	"log"
	"repair/models"
	"repair/models/adjust"
	"strings"
)

func MoveMethod(conn *models.Connection, apt int64, old int64, newer int64) {
	breakdownManager := models.NewBreakdownManager(conn)
	breakdownManager.Isolation = false
	reviewManager := models.NewReviewManager(conn)
	reviewManager.Isolation = false
	breakdownhistoryManager := models.NewBreakdownhistoryManager(conn)
	breakdownhistoryManager.Isolation = false

	breakdowns := breakdownManager.FindByMethod(old)
	for _, v := range breakdowns {
		v.Method = newer
		breakdownManager.Update(&v)
	}

	reviews := reviewManager.FindByMethod(old)
	for _, v := range reviews {
		v.Method = newer
		reviewManager.Update(&v)
	}

	breakdownhistorys := breakdownhistoryManager.FindByMethod(old)
	for _, v := range breakdownhistorys {
		v.Method = newer
		breakdownhistoryManager.Update(&v)
	}
}

func DeleteMethod(conn *models.Connection, apt int64, old int64) {
	breakdownManager := models.NewBreakdownManager(conn)
	breakdownManager.Isolation = false
	reviewManager := models.NewReviewManager(conn)
	reviewManager.Isolation = false
	breakdownhistoryManager := models.NewBreakdownhistoryManager(conn)
	breakdownhistoryManager.Isolation = false

	breakdowns := breakdownManager.FindByMethod(old)
	for _, v := range breakdowns {
		breakdownManager.Delete(v.Id)
	}

	reviews := reviewManager.FindByMethod(old)
	for _, v := range reviews {
		reviewManager.Delete(v.Id)
	}

	breakdownhistorys := breakdownhistoryManager.FindByMethod(old)
	for _, v := range breakdownhistorys {
		breakdownhistoryManager.Delete(v.Id)
	}
}

func MoveCategory(conn *models.Connection, apt int64, old int64, newer int64) {
	standardManager := models.NewStandardManager(conn)
	standardManager.Isolation = false
	breakdownManager := models.NewBreakdownManager(conn)
	breakdownManager.Isolation = false
	historyManager := models.NewHistoryManager(conn)
	historyManager.Isolation = false
	reviewManager := models.NewReviewManager(conn)
	reviewManager.Isolation = false
	adjustManager := models.NewAdjustManager(conn)
	adjustManager.Isolation = false
	breakdownhistoryManager := models.NewBreakdownhistoryManager(conn)
	breakdownhistoryManager.Isolation = false
	categoryManager := models.NewCategoryManager(conn)
	categoryManager.Isolation = false

	standards := standardManager.FindByCategory(old)
	for _, v := range standards {
		v.Category = newer
		standardManager.Update(&v)
	}

	breakdowns := breakdownManager.FindByCategory(old)
	for _, v := range breakdowns {
		v.Category = newer
		breakdownManager.Update(&v)
	}

	category := categoryManager.Get(newer)
	subcategory := categoryManager.Get(category.Parent)
	historys := historyManager.FindByCategory(old)
	for _, v := range historys {
		v.Topcategory = subcategory.Parent
		v.Subcategory = category.Parent
		v.Category = newer
		historyManager.Update(&v)
	}

	reviews := reviewManager.FindByCategory(old)
	for _, v := range reviews {
		v.Category = newer
		reviewManager.Update(&v)
	}

	adjusts := adjustManager.FindByCategory(old)
	for _, v := range adjusts {
		v.Category = newer
		adjustManager.Update(&v)
	}

	breakdownhistorys := breakdownhistoryManager.FindByCategory(old)
	for _, v := range breakdownhistorys {
		v.Category = newer
		breakdownhistoryManager.Update(&v)
	}

	items := categoryManager.FindByAptParent(apt, old)
	for _, v := range items {
		v.Parent = newer
		categoryManager.Update(&v)
	}
}

func DeleteCategory(conn *models.Connection, apt int64, id int64) {
	standardManager := models.NewStandardManager(conn)
	standardManager.Isolation = false
	breakdownManager := models.NewBreakdownManager(conn)
	breakdownManager.Isolation = false
	historyManager := models.NewHistoryManager(conn)
	historyManager.Isolation = false
	reviewManager := models.NewReviewManager(conn)
	reviewManager.Isolation = false
	adjustManager := models.NewAdjustManager(conn)
	adjustManager.Isolation = false
	breakdownhistoryManager := models.NewBreakdownhistoryManager(conn)
	breakdownhistoryManager.Isolation = false
	categoryManager := models.NewCategoryManager(conn)
	categoryManager.Isolation = false

	standardManager.DeleteByCategory(id)
	breakdownManager.DeleteByCategory(id)
	historyManager.DeleteByCategory(id)
	reviewManager.DeleteByCategory(id)
	adjustManager.DeleteByCategory(id)
	breakdownhistoryManager.DeleteByCategory(id)
	categoryManager.Delete(id)
}

func UpdateCategoryName(conn *models.Connection, apt int64, old string, newer string) {
	categoryManager := models.NewCategoryManager(conn)
	categoryManager.Isolation = false
	item := categoryManager.GetByAptName(apt, old)
	if item == nil {
		return
	}
	item.Name = (newer)
	categoryManager.Update(item)
}

func GetCategory(conn *models.Connection, apt int64, name string, newName string, level int, cycle int, percent int, parent int64, order int) *models.Category {
	categoryManager := models.NewCategoryManager(conn)
	categoryManager.Isolation = false

	items := categoryManager.FindByAptParent(apt, parent)
	var item *models.Category = nil
	for _, v := range items {
		if v.Name == name {
			item = &v
			break
		}
	}
	if item == nil {
		if newName == "" {
			newName = name
		}
		item = &models.Category{
			Name:     newName,
			Level:    level,
			Parent:   parent,
			Cycle:    cycle,
			Percent:  percent,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    order,
			Apt:      apt,
		}
		categoryManager.Insert(item)
		item.Id = categoryManager.GetIdentity()
	} else {
		if newName != "" {
			item.Name = newName
			categoryManager.Update(item)
		}
	}

	return item
}

func Change(id int64) {
	conn := models.NewConnection()
	defer conn.Close()

	conn.Begin()
	defer conn.Rollback()

	categoryManager := models.NewCategoryManager(conn)
	categoryManager.Isolation = false
	standardManager := models.NewStandardManager(conn)
	standardManager.Isolation = false
	breakdownManager := models.NewBreakdownManager(conn)
	breakdownManager.Isolation = false
	adjustManager := models.NewAdjustManager(conn)
	adjustManager.Isolation = false

	apt := id

	category1 := categoryManager.GetByAptName(apt, "(1) 모르타르 마감")
	category3 := categoryManager.GetByAptName(apt, "(3) 고분자시트방수")
	UpdateCategoryName(conn, apt, "(4) 금속기와 잇기", "(2) 금속기와 잇기")
	UpdateCategoryName(conn, apt, "(5) 아스팔트 슁글잇기", "(3) 아스팔트 슁글잇기")

	roof := categoryManager.GetByAptName(apt, "가. 지붕")
	log.Println("========================")
	log.Println(roof)
	if roof != nil {
		item := categoryManager.GetByAptLevelParentName(apt, 3, roof.Id, "(6) 기타")
		if item != nil {
			item.Name = "(4) 기타"
			categoryManager.Update(item)
		}
	}

	var proof *models.Category
	root1 := categoryManager.GetByAptName(apt, "가. 지붕")
	category2 := categoryManager.GetByAptName(apt, "(2) 고분자도막방수")
	if category2 != nil {
		category2.Name = "(1) 방수"
		category2.Order = 101010000
		categoryManager.Update(category2)

		items := categoryManager.FindByAptParent(apt, category2.Id)
		for i, v := range items {
			if i > 0 {
				categoryManager.Delete(v.Id)
			}
			v.Name = "전면수리"
			v.Cycle = 15
			v.Percent = 100
			v.Order = 101010100
			categoryManager.Update(&v)
			proof = &v
		}
	} else {
		category2 = &models.Category{
			Name:     "(1) 방수",
			Level:    3,
			Parent:   root1.Id,
			Cycle:    0,
			Percent:  0,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    101010000,
			Apt:      apt,
		}
		categoryManager.Insert(category2)
		category2.Id = categoryManager.GetIdentity()

		item := &models.Category{
			Name:     "전면수리",
			Level:    4,
			Parent:   category2.Id,
			Cycle:    15,
			Percent:  100,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    101010100,
			Apt:      apt,
		}
		categoryManager.Insert(item)
		item.Id = categoryManager.GetIdentity()
		proof = item
	}

	var args []interface{}
	args = append(args, models.Ordering("c_order"))

	if category1 != nil {
		items := categoryManager.FindByAptParent(apt, category1.Id)
		for _, v := range items {
			categoryManager.Delete(v.Id)
			MoveMethod(conn, apt, v.Id, proof.Id)
		}

		MoveCategory(conn, apt, category1.Id, category2.Id)
		categoryManager.Delete(category1.Id)
	}

	if category3 != nil {
		items := categoryManager.FindByAptParent(apt, category3.Id)
		for _, v := range items {
			categoryManager.Delete(v.Id)
			MoveMethod(conn, apt, v.Id, proof.Id)
		}

		MoveCategory(conn, apt, category3.Id, category2.Id)
		categoryManager.Delete(category3.Id)
	}

	pp := categoryManager.GetByAptName(apt, "(2) 수성페인트칠")
	if pp != nil {
		items := categoryManager.FindByAptParent(apt, pp.Id)
		for _, v := range items {
			if v.Name == "전면도장" {
				v.Cycle = 8
				v.Percent = 100
				categoryManager.Update(&v)
			}
		}
	}

	UpdateCategoryName(conn, apt, "(2) 수성페인트칠", "(2) 페인트칠")

	shelter := categoryManager.GetByAptName(apt, "7. 피난시설")
	var shelter1 *models.Category
	if shelter == nil {
		shelter = GetCategory(conn, apt, "7. 피난시설", "", 1, 0, 0, 0, 700000000)
		shelter1 = GetCategory(conn, apt, "가. 피난시설", "", 2, 0, 0, shelter.Id, 701000000)
	}

	var door *models.Category
	var doorDetail *models.Category
	window := categoryManager.GetByAptName(apt, "다. 외부 창·문")
	if window != nil {
		item := categoryManager.GetByAptLevelParentName(apt, 3, window.Id, "(3) 방화문")
		if item == nil {
			item = &models.Category{
				Name:     "(1) 방화문",
				Level:    3,
				Parent:   shelter1.Id,
				Cycle:    0,
				Percent:  0,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    701010000,
				Apt:      apt,
			}
			categoryManager.Insert(item)
			item.Id = categoryManager.GetIdentity()

			item2 := &models.Category{
				Name:     "전면교체",
				Level:    4,
				Parent:   item.Id,
				Cycle:    25,
				Percent:  100,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    701010100,
				Apt:      apt,
			}
			categoryManager.Insert(item2)
			doorDetail = item2
		} else {
			item.Name = "(1) 방화문"
			item.Level = 3
			item.Parent = shelter1.Id
			item.Cycle = 0
			item.Percent = 0
			item.Unit = ""
			item.Elevator = 2
			item.Remark = ""
			item.Order = 701010000
			categoryManager.Update(item)

			items := categoryManager.FindByAptParent(apt, item.Id, args)
			for i, v := range items {
				v.Parent = item.Id
				v.Cycle = 15
				v.Percent = 100
				v.Order = item.Order + (i+1)*100
				categoryManager.Update(&v)

				if i == 0 {
					doorDetail = &v
				}
			}
		}
		door = item

		item2 := categoryManager.GetByAptLevelParentName(apt, 3, window.Id, "(4) 기타")
		if item2 != nil {
			item2.Name = "(3) 기타"
			categoryManager.Update(item2)
		}
	}

	top := categoryManager.GetByAptName(apt, "가. 천장")
	var paint *models.Category
	var paint1 *models.Category

	var etc *models.Category
	var etc1 *models.Category
	var etc2 *models.Category
	if top != nil {
		top.Name = "가. 내부"
		categoryManager.Update(top)
		item := categoryManager.GetByAptLevelParentName(apt, 3, top.Id, "(1) 수성도료칠")
		if item == nil {
			item = &models.Category{
				Name:     "(1) 페인트칠",
				Level:    3,
				Parent:   top.Id,
				Cycle:    0,
				Percent:  0,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201010000,
				Apt:      apt,
			}
			categoryManager.Insert(item)
			item.Id = categoryManager.GetIdentity()

			item2 := &models.Category{
				Name:     "전면도장",
				Level:    4,
				Parent:   item.Id,
				Cycle:    8,
				Percent:  100,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201010100,
				Apt:      apt,
			}
			categoryManager.Insert(item2)
			item2.Id = categoryManager.GetIdentity()

			paint1 = item2
		} else {
			item.Name = "(1) 페인트칠"
			item.Level = 3
			item.Parent = top.Id
			item.Cycle = 0
			item.Percent = 0
			item.Unit = ""
			item.Elevator = 2
			item.Remark = ""
			item.Order = 201010000
			categoryManager.Update(item)

			items := categoryManager.FindByAptParent(apt, item.Id, args)
			var id int64
			for i, v := range items {
				if i > 0 {
					MoveMethod(conn, apt, v.Id, id)
					categoryManager.Delete(v.Id)
				}
				v.Name = "전면도장"
				v.Cycle = 8
				v.Percent = 100
				v.Order = item.Order + (i+1)*100
				categoryManager.Update(&v)
				id = v.Id
				paint1 = &v
			}

			items2 := standardManager.FindByCategory(item.Id)
			for _, v := range items2 {
				v.Name = "천장" + strings.TrimSpace(strings.ReplaceAll(v.Name, "천장", ""))
				standardManager.Update(&v)
			}
		}

		paint = item

		item2 := categoryManager.GetByAptLevelParentName(apt, 3, top.Id, "(2) 유성도료칠")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "천장" + strings.TrimSpace(strings.ReplaceAll(v.Name, "천장", ""))
				standardManager.Update(&v)
			}

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)
		}

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, top.Id, "(3) 합성수지도료칠")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "천장" + strings.TrimSpace(strings.ReplaceAll(v.Name, "천장", ""))
				standardManager.Update(&v)
			}

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)
		}

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, top.Id, "(4) 무늬코트")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "천장" + strings.TrimSpace(strings.ReplaceAll(v.Name, "천장", ""))
				standardManager.Update(&v)
			}

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)
		}

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, top.Id, "(5) 균열누수공사")
		if item2 == nil {
			item2 = &models.Category{
				Name:     "(2) 균열누수공사",
				Level:    3,
				Parent:   top.Id,
				Cycle:    0,
				Percent:  0,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201020000,
				Apt:      apt,
			}
			categoryManager.Insert(item2)
			item2.Id = categoryManager.GetIdentity()

			item3 := &models.Category{
				Name:     "전면보수",
				Level:    4,
				Parent:   item2.Id,
				Cycle:    25,
				Percent:  100,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201020100,
				Apt:      apt,
			}
			categoryManager.Insert(item3)
		} else {
			item2.Name = "(2) 균열누수공사"
			item2.Level = 3
			item2.Parent = top.Id
			item2.Cycle = 0
			item2.Percent = 0
			item2.Unit = ""
			item2.Elevator = 2
			item2.Remark = ""
			item2.Order = 201020000
			categoryManager.Update(item2)

			items := categoryManager.FindByAptParent(apt, item2.Id, args)
			var id int64
			for i, v := range items {
				if i == 0 {
					v.Name = "전면보수"
					v.Cycle = 25
					v.Percent = 100
					v.Order = item2.Order + (i+1)*100
					categoryManager.Update(&v)
					id = v.Id
				} else {
					MoveMethod(conn, apt, v.Id, id)
					categoryManager.Delete(v.Id)
				}
			}

			items2 := standardManager.FindByCategory(item.Id)
			for _, v := range items2 {
				v.Name = "천장" + strings.TrimSpace(strings.ReplaceAll(v.Name, "천장", ""))
				standardManager.Update(&v)
			}
		}

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, top.Id, "(6) 기타")
		if item2 == nil {
			item2 = &models.Category{
				Name:     "(3) 기타",
				Level:    3,
				Parent:   top.Id,
				Cycle:    0,
				Percent:  0,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201030000,
				Apt:      apt,
			}
			categoryManager.Insert(item2)
			item2.Id = categoryManager.GetIdentity()
			etc = item2

			item3 := &models.Category{
				Name:     "부분보수",
				Level:    4,
				Parent:   item2.Id,
				Cycle:    10,
				Percent:  20,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201030100,
				Apt:      apt,
			}
			categoryManager.Insert(item3)
			item3.Id = categoryManager.GetIdentity()

			etc1 = item3

			item3 = &models.Category{
				Name:     "전면교체",
				Level:    4,
				Parent:   item2.Id,
				Cycle:    25,
				Percent:  100,
				Unit:     "",
				Elevator: 2,
				Remark:   "",
				Order:    201030200,
				Apt:      apt,
			}
			categoryManager.Insert(item3)
			item3.Id = categoryManager.GetIdentity()

			etc2 = item3
		} else {
			item2.Name = "(3) 기타"
			item2.Level = 3
			item2.Parent = top.Id
			item2.Cycle = 0
			item2.Percent = 0
			item2.Unit = ""
			item2.Elevator = 2
			item2.Remark = ""
			item2.Order = 201030000
			categoryManager.Update(item2)
			etc = item2

			items := categoryManager.FindByAptParent(apt, item2.Id, args)
			for i, v := range items {
				if v.Name == "부분수선" {
					v.Name = "부분보수"
					v.Order = item2.Order + (i+1)*100
					categoryManager.Update(&v)

					etc1 = &v
				} else {
					v.Order = item2.Order + (i+1)*100
					categoryManager.Update(&v)

					etc2 = &v
				}
			}

			items2 := standardManager.FindByCategory(item.Id)
			for _, v := range items2 {
				v.Name = "천장" + strings.TrimSpace(strings.ReplaceAll(v.Name, "천장", ""))
				standardManager.Update(&v)
			}
		}
	}

	items := categoryManager.FindByAptOrder(apt, 202010100)
	for _, v := range items {
		MoveMethod(conn, apt, v.Id, paint1.Id)
		categoryManager.Delete(v.Id)
	}
	items = categoryManager.FindByAptOrder(apt, 202020100)
	for _, v := range items {
		MoveMethod(conn, apt, v.Id, paint1.Id)
		categoryManager.Delete(v.Id)
	}
	items = categoryManager.FindByAptOrder(apt, 202030100)
	for _, v := range items {
		MoveMethod(conn, apt, v.Id, paint1.Id)
		categoryManager.Delete(v.Id)
	}
	items = categoryManager.FindByAptOrder(apt, 202040100)
	for _, v := range items {
		MoveMethod(conn, apt, v.Id, paint1.Id)
		categoryManager.Delete(v.Id)
	}
	items = categoryManager.FindByAptOrder(apt, 202050100)
	for _, v := range items {
		MoveMethod(conn, apt, v.Id, paint1.Id)
		categoryManager.Delete(v.Id)
	}

	wall := categoryManager.GetByAptName(apt, "나. 내벽")
	if wall != nil {
		categoryManager.Delete(wall.Id)
		item := paint
		item2 := categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(1) 수성도료칠")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "벽체" + strings.TrimSpace(strings.ReplaceAll(v.Name, "벽체", ""))
				standardManager.Update(&v)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
		}
		categoryManager.Delete(item2.Id)

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(2) 유성도료칠")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "벽체" + strings.TrimSpace(strings.ReplaceAll(v.Name, "벽체", ""))
				standardManager.Update(&v)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
		}
		categoryManager.Delete(item2.Id)

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(3) 합성수지도료칠")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "벽체" + strings.TrimSpace(strings.ReplaceAll(v.Name, "벽체", ""))
				standardManager.Update(&v)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
		}
		categoryManager.Delete(item2.Id)

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(4) 무늬코트")
		if item2 != nil {
			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "벽체" + strings.TrimSpace(strings.ReplaceAll(v.Name, "벽체", ""))
				standardManager.Update(&v)
			}
			MoveCategory(conn, apt, item2.Id, item.Id)
			categoryManager.Delete(item2.Id)

			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				MoveMethod(conn, apt, v.Id, paint1.Id)
				categoryManager.Delete(v.Id)
			}
		}
		categoryManager.Delete(item2.Id)

		item2 = categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(5) 기타")
		if item2 != nil {
			items2 := categoryManager.FindByAptParent(apt, item2.Id)
			for _, v := range items2 {
				if v.Name == "부분수리" || v.Name == "부분수선" {
					MoveMethod(conn, apt, v.Id, etc1.Id)
				} else {
					MoveMethod(conn, apt, v.Id, etc2.Id)
				}
				categoryManager.Delete(v.Id)
			}

			items := standardManager.FindByCategory(item2.Id)
			for _, v := range items {
				v.Name = "벽체" + strings.TrimSpace(strings.ReplaceAll(v.Name, "벽체", ""))
				standardManager.Update(&v)
			}
			MoveCategory(conn, apt, item2.Id, etc.Id)
			categoryManager.Delete(item2.Id)
		}
	}

	paintBreakdowns := breakdownManager.Find([]interface{}{
		models.Where{Column: "category", Value: paint.Id, Compare: "="},
	})

	for _, v := range paintBreakdowns {
		adjustManager.UpdateWhere(
			[]adjust.Params{
				{Column: adjust.ColumnCategory, Value: paint.Parent},
			},
			[]interface{}{
				models.Where{Column: "category", Value: v.Subcategory, Compare: "="},
			})

		v.Subcategory = paint.Parent

		breakdownManager.Update(&v)

	}

	land := categoryManager.GetByAptName(apt, "다. 바닥")
	if land != nil {
		land.Name = "나. 바닥"
		categoryManager.Update(land)

		item := categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(1) 바닥 마감재")
		if item != nil {
			item.Name = "(1) 자하주차장(바닥)"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(2) 타일붙이기")
		if item != nil {
			items2 := categoryManager.FindByAptParent(apt, item.Id)
			for _, v2 := range items2 {
				categoryManager.Delete(v2.Id)
			}
			DeleteCategory(conn, apt, item.Id)
			categoryManager.Delete(item.Id)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, wall.Id, "(3) 기타")
		if item != nil {
			item.Name = "(2) 기타"
			categoryManager.Update(item)
		}
	}

	stairs := categoryManager.GetByAptName(apt, "라. 계단")
	if stairs != nil {
		items := categoryManager.FindByAptParent(apt, stairs.Id)
		for _, v := range items {
			items2 := categoryManager.FindByAptParent(apt, v.Id)
			for _, v2 := range items2 {
				categoryManager.Delete(v2.Id)
			}
			DeleteCategory(conn, apt, v.Id)
			categoryManager.Delete(v.Id)
		}

		DeleteCategory(conn, apt, stairs.Id)
		categoryManager.Delete(stairs.Id)
	}

	trans := categoryManager.GetByAptName(apt, "나. 변전설비")
	if trans != nil {
		v := categoryManager.GetByAptLevelParentName(apt, 3, trans.Id, "(5) 소방시설유지보수")
		if v != nil {
			items2 := categoryManager.FindByAptParent(apt, v.Id)
			for _, v2 := range items2 {
				categoryManager.Delete(v2.Id)
			}
			DeleteCategory(conn, apt, v.Id)
			categoryManager.Delete(v.Id)
		}

		item := categoryManager.GetByAptLevelParentName(apt, 3, trans.Id, "(6) 옥상비상문자동개폐장치")
		if item != nil {
			item.Name = "(2) 옥상 비상문 자동 개폐장치"
			item.Level = 3
			item.Parent = shelter1.Id
			item.Cycle = 0
			item.Percent = 0
			item.Unit = ""
			item.Elevator = 2
			item.Remark = ""
			item.Order = 701020000
			categoryManager.Update(item)

			items2 := categoryManager.FindByAptParent(apt, item.Id)
			for _, v2 := range items2 {
				v2.Name = "전면교체"
				v2.Cycle = 15
				v2.Percent = 100
				v2.Order = 701020200
				categoryManager.Update(&v2)
			}

			GetCategory(conn, apt, "부분수선", "", 4, 5, 30, item.Id, 701020100)
		} else {
			shelter12 := GetCategory(conn, apt, "(2) 옥상 비상문 자동 개폐장치", "", 3, 0, 0, shelter1.Id, 701020000)
			GetCategory(conn, apt, "부분수선", "", 4, 5, 30, shelter12.Id, 701020100)
			GetCategory(conn, apt, "전면교체", "", 4, 15, 100, shelter12.Id, 701020200)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, trans.Id, "(7) 기타")
		if item != nil {
			item.Name = "(5) 기타"
			item.Order = 302050000
			categoryManager.Update(item)
		}
	}

	fire := categoryManager.GetByAptName(apt, "다. 자동화재감지설비")
	if fire != nil {
		v := categoryManager.GetByAptLevelParentName(apt, 3, fire.Id, "(4) 방화문")
		if v != nil {
			MoveCategory(conn, apt, v.Id, door.Id)
			items2 := categoryManager.FindByAptParent(apt, v.Id)
			for _, v2 := range items2 {
				MoveMethod(conn, apt, v2.Id, doorDetail.Id)
				categoryManager.Delete(v2.Id)
			}
			DeleteCategory(conn, apt, v.Id)
			categoryManager.Delete(v.Id)
		}

		// v = categoryManager.GetByAptLevelParentName(apt, 3, fire.Id, "(3) 기타")
		// if v != nil {
		// 	items2 := categoryManager.FindByAptParent(apt, v.Id)
		// 	for _, v2 := range items2 {
		// 		categoryManager.Delete(v2.Id)
		// 	}
		// 	DeleteCategory(conn, apt, v.Id)
		// }

		items = categoryManager.FindByAptOrder(apt, 303040100)
		for _, v2 := range items {
			DeleteCategory(conn, apt, v2.Id)
			categoryManager.Delete(v2.Id)
		}
		items = categoryManager.FindByAptOrder(apt, 303040050)
		for _, v2 := range items {
			DeleteCategory(conn, apt, v2.Id)
			categoryManager.Delete(v2.Id)
		}
	}

	thunther := categoryManager.GetByAptName(apt, "(1) 피뢰설비")
	if thunther != nil {
		items := categoryManager.FindByAptParent(apt, thunther.Id)
		for _, v := range items {
			if v.Name == "전면교체" {
				v.Name = "부분수선"
				v.Cycle = 10
				v.Percent = 30
				categoryManager.Update(&v)
			}
		}
	}

	water := categoryManager.GetByAptName(apt, "가. 급수설비")
	if water != nil {
		item := categoryManager.GetByAptLevelParentName(apt, 3, water.Id, "(9) 기타")
		if item != nil {
			item.Name = "(10) 기타"
			item.Order = 401100000
			categoryManager.Update(item)

			items2 := categoryManager.FindByAptParent(apt, item.Id)
			for i, v2 := range items2 {
				v2.Order = 401100100 + i*100
				categoryManager.Update(&v2)
			}
		}

		item2 := &models.Category{
			Name:     "(9) 저수조 방수",
			Level:    3,
			Parent:   water.Id,
			Cycle:    0,
			Percent:  0,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    401090000,
			Apt:      apt,
		}
		categoryManager.Insert(item2)
		item2.Id = categoryManager.GetIdentity()

		item3 := &models.Category{
			Name:     "전면교체",
			Level:    4,
			Parent:   item2.Id,
			Cycle:    25,
			Percent:  100,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    401090100,
			Apt:      apt,
		}
		categoryManager.Insert(item3)
	}

	gas := categoryManager.GetByAptName(apt, "나. 가스설비")
	if gas != nil {
		items := categoryManager.FindByAptParent(apt, gas.Id)
		for _, v := range items {
			if v.Name == "(1) 배관" {
				items2 := categoryManager.FindByAptParent(apt, v.Id)
				for _, v2 := range items2 {
					v2.Name = "부분수선"
					v2.Cycle = 10
					v2.Percent = 10
					categoryManager.Update(&v2)
				}
			}
			if v.Name == "(2) 밸브" {
				items2 := categoryManager.FindByAptParent(apt, v.Id)
				for _, v2 := range items2 {
					v2.Name = "부분수선"
					v2.Cycle = 10
					v2.Percent = 30
					categoryManager.Update(&v2)
				}
			}
		}
	}

	drain := categoryManager.GetByAptName(apt, "다. 배수설비")
	if drain != nil {
		items := categoryManager.FindByAptParent(apt, drain.Id)
		for _, v := range items {
			if v.Name == "(3) 오배수관(주철)" {
				items2 := categoryManager.FindByAptParent(apt, v.Id)
				for _, v2 := range items2 {
					v2.Name = "부분수선"
					v2.Cycle = 10
					v2.Percent = 10
					categoryManager.Update(&v2)
				}
			}
			if v.Name == "(4) 오배수관(PVC)" {
				items2 := categoryManager.FindByAptParent(apt, v.Id)
				for _, v2 := range items2 {
					v2.Name = "부분수선"
					v2.Cycle = 10
					v2.Percent = 10
					categoryManager.Update(&v2)
				}
			}
		}
	}

	wind := categoryManager.GetByAptName(apt, "(1) 환기팬")
	if wind != nil {
		items := categoryManager.FindByAptParent(apt, wind.Id)
		for i, v := range items {
			if i == 0 {
				v.Name = "부분수선"
				v.Cycle = 10
				v.Percent = 10
				categoryManager.Update(&v)
			} else {
				DeleteCategory(conn, apt, v.Id)
				categoryManager.Delete(v.Id)
			}
		}
	}

	boiler := categoryManager.GetByAptName(apt, "(3) 보일러수관")
	if boiler != nil {
		items := categoryManager.FindByAptParent(apt, boiler.Id)
		for _, v := range items {
			categoryManager.Delete(v.Id)
		}
		DeleteCategory(conn, apt, boiler.Id)
		categoryManager.Delete(boiler.Id)
	}

	fire2 := categoryManager.GetByAptName(apt, "가. 난방설비")
	if fire2 != nil {
		item := categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(4) 난방순환펌프")
		if item != nil {
			item.Name = "(3) 난방순환펌프"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(5) 난방관(강관)")
		if item != nil {
			item.Name = "(4) 난방관(강관)"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(6) 난방관(동관)")
		if item != nil {
			item.Name = "(5) 난방관(동관)"
			categoryManager.Update(item)
			items := categoryManager.FindByAptParent(apt, item.Id)
			for _, v := range items {
				item.Cycle = 10
				item.Percent = 20
				categoryManager.Update(&v)
			}
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(7) 난방관(STS)")
		if item != nil {
			item.Name = "(6) 난방관(STS)"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(8) 자동제어 기기")
		if item != nil {
			item.Name = "(7) 자동제어 기기"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(9) 열교환기")
		if item != nil {
			item.Name = "(8) 열교환기"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(10) 밸브류")
		if item != nil {
			item.Name = "(9) 밸브류"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(11) 계량기")
		if item != nil {
			item.Name = "(10) 계량기"
			categoryManager.Update(item)
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(12) 기타")
		if item != nil {
			item.Name = "(11) 기타"
			categoryManager.Update(item)
		}
	}

	water2 := categoryManager.GetByAptName(apt, "나. 급탕설비")
	if water2 != nil {
		item := categoryManager.GetByAptLevelParentName(apt, 3, water2.Id, "(4) 급탕관(동관)")
		if item != nil {
			items := categoryManager.FindByAptParent(apt, item.Id)
			for _, v := range items {
				item.Cycle = 10
				item.Percent = 20
				categoryManager.Update(&v)
			}
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(5) 급탕관(STS)")
		if item != nil {
			items := categoryManager.FindByAptParent(apt, item.Id)
			for _, v := range items {
				item.Cycle = 10
				item.Percent = 20
				categoryManager.Update(&v)
			}
		}

		item = categoryManager.GetByAptLevelParentName(apt, 3, water2.Id, "(9) 기타")
		if item != nil {
			item.Name = "(10) 기타"
			item.Order = 502100000
			categoryManager.Update(item)

			items2 := categoryManager.FindByAptParent(apt, item.Id)
			for i, v2 := range items2 {
				v2.Order = 502100100 + i*100
				categoryManager.Update(&v2)
			}
		}

		item2 := &models.Category{
			Name:     "(9) 자동제어기기",
			Level:    3,
			Parent:   water2.Id,
			Cycle:    0,
			Percent:  0,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    502090000,
			Apt:      apt,
		}
		categoryManager.Insert(item2)
		item2.Id = categoryManager.GetIdentity()

		item3 := &models.Category{
			Name:     "전면교체",
			Level:    4,
			Parent:   item2.Id,
			Cycle:    20,
			Percent:  100,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    502090100,
			Apt:      apt,
		}
		categoryManager.Insert(item3)

		item = categoryManager.GetByAptLevelParentName(apt, 3, fire2.Id, "(9) 기타")
		if item != nil {
			item.Name = "(10) 기타"
			item.Order = 502100000
			categoryManager.Update(item)

			items := categoryManager.FindByAptParent(apt, item.Id)
			for _, v := range items {
				if v.Name == "부분수선" {
					v.Order = 502100100
				} else {
					v.Order = 502100200
				}
				categoryManager.Update(&v)
			}
		}
	}

	item := categoryManager.GetByAptName(apt, "(6) 소화밸브")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "전면교체" {
				v.Cycle = 25
				v.Percent = 100
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(7) 소방시설유지보수")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "전면교체" {
				v.Cycle = 10
				v.Percent = 100
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(1) 아스팔트포장")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "부분수리" {
				v.Cycle = 5
				v.Percent = 10
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(3) 어린이놀이시설")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "부분수리" {
				v.Cycle = 5
				v.Percent = 10
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(4) 보도블록")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "부분수리" {
				v.Cycle = 5
				v.Percent = 10
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(5) 정화조")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "부분수리" {
				v.Cycle = 5
				v.Percent = 15
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(10) 조경시설물")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			v.Name = "부분수선"
			v.Cycle = 10
			v.Percent = 10
			categoryManager.Update(&v)
		}
	}

	item = categoryManager.GetByAptName(apt, "(11) 안내표지판")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			v.Name = "부분수선"
			v.Cycle = 10
			v.Percent = 30
			categoryManager.Update(&v)
		}
	}

	item = categoryManager.GetByAptName(apt, "(14) 저수조 방수")
	if item != nil {
		items2 := categoryManager.FindByAptParent(apt, item.Id)
		for _, v2 := range items2 {
			categoryManager.Delete(v2.Id)
		}
		DeleteCategory(conn, apt, item.Id)
		categoryManager.Delete(item.Id)
	}

	item = categoryManager.GetByAptName(apt, "가. 옥외부대시설 및 옥외복리시설")
	if item == nil {
		item = categoryManager.GetByAptName(apt, "가. 옥외부대시설 및 옥외 복리시설")
		if item == nil {
			item = categoryManager.GetByAptName(apt, "가. 옥외 부대시설 및 옥외 복리시설")
			if item == nil {
				items := categoryManager.FindByAptOrder(apt, 601000000)
				for _, v := range items {
					item = &v
				}
			}
		}
	}

	if item != nil {
		item2 := &models.Category{
			Name:     "(14) 전기자동차의 고정형 충전기",
			Level:    3,
			Parent:   item.Id,
			Cycle:    0,
			Percent:  0,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    601140000,
			Apt:      apt,
		}
		categoryManager.Insert(item2)
		item2.Id = categoryManager.GetIdentity()

		item3 := &models.Category{
			Name:     "부분수선",
			Level:    4,
			Parent:   item2.Id,
			Cycle:    5,
			Percent:  10,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    601140100,
			Apt:      apt,
		}

		categoryManager.Insert(item3)

		item3 = &models.Category{
			Name:     "전면교체",
			Level:    4,
			Parent:   item2.Id,
			Cycle:    10,
			Percent:  100,
			Unit:     "",
			Elevator: 2,
			Remark:   "",
			Order:    601140200,
			Apt:      apt,
		}
		categoryManager.Insert(item3)
	}

	item = categoryManager.GetByAptName(apt, "다. 냉난방 및 공조기")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "(5) 기타" {
				v.Name = ("(7) 기타")
				v.Order += 20000
				categoryManager.Update(&v)
				items2 := categoryManager.FindByAptParent(apt, v.Id)

				for _, v := range items2 {
					v.Order += 20000
					if v.Name == "전면수리" {
						v.Cycle = 25
						v.Percent = 100
					}
					categoryManager.Update(&v)
				}
			}

			if v.Name == "(6) 배관(STS)" {
				v.Name = ("(5) 배관(STS)")
				v.Order -= 10000
				categoryManager.Update(&v)
				items2 := categoryManager.FindByAptParent(apt, v.Id)
				for _, v := range items2 {
					v.Order -= 10000
					categoryManager.Update(&v)
				}
			}

			if v.Name == "(7) 계량기" {
				v.Name = ("(6) 계량기")
				v.Order -= 10000
				categoryManager.Update(&v)
				items2 := categoryManager.FindByAptParent(apt, v.Id)
				for _, v := range items2 {
					v.Order -= 10000
					if v.Name == "전면교체" {
						v.Cycle = 8
						v.Percent = 100
					}
					categoryManager.Update(&v)
				}
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(8) 자전거보관소")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "전면교체" {
				v.Cycle = 15
				v.Percent = 100
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(3) 판넬 붙이기")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			v.Cycle = 10
			v.Percent = 5
			categoryManager.Update(&v)
		}
	}

	item = categoryManager.GetByAptName(apt, "(1) 바닥 마감재")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "부분수리" {
				v.Cycle = 5
				v.Percent = 10
				categoryManager.Update(&v)
			}
		}
	}

	item = categoryManager.GetByAptName(apt, "(1) 발전기")
	if item != nil {
		items := categoryManager.FindByAptParent(apt, item.Id)
		for _, v := range items {
			if v.Name == "부분수선" {
				v.Cycle = 10
				v.Percent = 10
				categoryManager.Update(&v)
			}
		}
	}

	item2 := &models.Category{
		Name:     "(3) 기타",
		Level:    3,
		Parent:   shelter1.Id,
		Cycle:    0,
		Percent:  0,
		Unit:     "",
		Elevator: 2,
		Remark:   "",
		Order:    701030000,
		Apt:      apt,
	}
	categoryManager.Insert(item2)
	item2.Id = categoryManager.GetIdentity()

	item3 := &models.Category{
		Name:     "부분수선",
		Level:    4,
		Parent:   item2.Id,
		Cycle:    10,
		Percent:  20,
		Unit:     "",
		Elevator: 2,
		Remark:   "",
		Order:    701030100,
		Apt:      apt,
	}

	categoryManager.Insert(item3)

	item3 = &models.Category{
		Name:     "전면교체",
		Level:    4,
		Parent:   item2.Id,
		Cycle:    10,
		Percent:  100,
		Unit:     "",
		Elevator: 2,
		Remark:   "",
		Order:    701030200,
		Apt:      apt,
	}
	categoryManager.Insert(item3)
	conn.Commit()

	cate := categoryManager.GetByAptName(apt, "(14) 전기자동차의 고정형 충전기")
	if cate != nil {
		item := models.Standard{
			Name:     "전기차 충전소",
			Direct:   392321,
			Labor:    0,
			Cost:     2,
			Unit:     "ea",
			Order:    0,
			Category: cate.Id,
			Apt:      apt,
		}

		standardManager.Insert(&item)
	}

	cate = categoryManager.GetByAptName(apt, "가. 피난시설")
	if cate != nil {
		cate2 := categoryManager.GetByAptLevelParentName(apt, 3, cate.Id, "(1) 방화문")
		if cate2 != nil {
			item := models.Standard{
				Name:     "방화문",
				Direct:   353091,
				Labor:    0,
				Cost:     0,
				Unit:     "ea",
				Order:    0,
				Category: cate2.Id,
				Apt:      apt,
			}

			standardManager.Insert(&item)
		}

		cate2 = categoryManager.GetByAptLevelParentName(apt, 3, cate.Id, "(2) 옥상 비상문 자동 개폐장치")
		if cate2 != nil {
			item := models.Standard{
				Name:     "옥상 비상문 자동 개폐장치",
				Direct:   235394,
				Labor:    0,
				Cost:     0,
				Unit:     "ea",
				Order:    0,
				Category: cate2.Id,
				Apt:      apt,
			}

			standardManager.Insert(&item)
		}
	}

	standards := standardManager.FindByCategory(category2.Id)
	find1 := false
	find2 := false
	for _, v := range standards {
		if v.Name == "우레탄방수" {
			find1 = true
		}

		if v.Name == "시트복합방수" {
			find2 = true
		}
	}

	if find1 == false {
		item := models.Standard{
			Name:     "우레탄방수",
			Direct:   43570,
			Labor:    0,
			Cost:     0,
			Unit:     "㎡",
			Order:    0,
			Category: category2.Id,
			Apt:      apt,
		}

		standardManager.Insert(&item)
	}

	if find2 == false {
		item := models.Standard{
			Name:     "시트복합방수",
			Direct:   67293,
			Labor:    0,
			Cost:     0,
			Unit:     "㎡",
			Order:    0,
			Category: category2.Id,
			Apt:      apt,
		}

		standardManager.Insert(&item)
	}
}

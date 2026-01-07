package api

import (
	"fmt"
	"log"
	"repair/controllers"
	"repair/models"
	"repair/models/blueprint"
)

type AptdongController struct {
	controllers.Controller
}

func (c *AptdongController) Post_Insert(item *models.Aptdong) {
	updateBlueprint(item)
}

func (c *AptdongController) Pre_Update(item *models.Aptdong) {
	conn := c.NewConnection()

	aptdongManager := models.NewAptdongManager(conn)
	aptdongetcManager := models.NewAptdongetcManager(conn)
	old := aptdongManager.Get(item.Id)

	if old.Parkingcount != item.Parkingcount ||
		old.Undergroundcount != item.Undergroundcount ||
		old.Groundcount != item.Groundcount ||
		old.Roofcount != item.Roofcount ||
		old.Topcount != item.Topcount {
		aptdongetcManager.DeleteByAptdong(item.Id)
	}

}

func (c *AptdongController) Post_Update(item *models.Aptdong) {
	updateBlueprint(item)
}

func (c *AptdongController) Post_Delete(item *models.Aptdong) {
	updateBlueprint(item)
}

func (c *AptdongController) Post_Deletebatch(item *[]models.Aptdong) {
	if len(*item) == 0 {
		return
	}

	updateBlueprint(&(*item)[0])
}

type Floor struct {
	Order             int
	Floortype         int
	OriginalFloortype int
	Aptdongetc        models.Aptdongetc
}

// @POST()
func (c *AptdongController) Blueprint(item *models.Aptdongblueprint) {
	conn := c.NewConnection()

	conn.Begin()
	defer conn.Rollback()

	aptdongetcManager := models.NewAptdongetcManager(conn)

	oldAptdongetc := aptdongetcManager.Find([]interface{}{
		models.Where{Column: "apt", Value: item.Apt, Compare: "="},
		models.Where{Column: "aptdong", Value: item.Aptdong, Compare: "="},
	})

	ids := make([]int64, 0)
	for _, v := range item.Items {
		aptdongetc := aptdongetcManager.GetByAptAptdongParentName(item.Apt, v.Aptdong, v.Parent, v.Name)
		if aptdongetc == nil || aptdongetc.Id == 0 {
			aptdongetcManager.Insert(&models.Aptdongetc{Name: v.Name, Floortype: v.Floortype, Parent: v.Parent, Order: v.Order, Aptdong: v.Aptdong, Apt: item.Apt})
		} else {
			aptdongetc.Order = v.Order
			aptdongetcManager.Update(aptdongetc)

			ids = append(ids, aptdongetc.Id)
		}
	}

	for _, v := range oldAptdongetc {
		flag := false
		for _, v2 := range ids {
			if v.Id == v2 {
				flag = true
				break
			}
		}

		if flag == false {
			aptdongetcManager.Delete(v.Id)
		}
	}

	conn.Commit()

	apt := models.Aptdong{Apt: item.Apt}
	updateBlueprint(&apt)
}

func updateBlueprint(item *models.Aptdong) {
	id := item.Apt

	conn := models.NewConnection()
	defer conn.Close()

	aptdongManager := models.NewAptdongManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)
	aptdongetcManager := models.NewAptdongetcManager(conn)

	aptdongs := aptdongManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("au_order,au_id")})
	blueprints := blueprintManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("bp_order desc,bp_id")})
	aptdongetc := aptdongetcManager.Find([]interface{}{
		models.Where{Column: "apt", Value: item.Apt, Compare: "="},
		models.Ordering("ae_order desc,ae_id"),
	})

	var privateDong models.Aptdong
	for _, v := range aptdongs {
		if v.Private == 3 {
			privateDong = v
			break
		}
	}

	var commonDong models.Aptdong
	for _, v := range aptdongs {
		if v.Private == 4 {
			commonDong = v
			break
		}
	}

	findLayout := false

	findCommons := []bool{false, false, false}
	commonTitles := []string{"기계실", "전기실", "발전기실"}
	for _, v := range blueprints {
		flag := false

		if v.Name == "배치도" {
			findLayout = true
			flag = true
		}

		if commonDong.Id > 0 {
			for i := 0; i < len(commonTitles); i++ {
				if v.Name == commonTitles[i] {
					findCommons[i] = true
					flag = true
				}
			}
		}

		for _, dong := range aptdongs {
			if v.Aptdong != dong.Id {
				continue
			}

			if v.Parent == 0 {
				flag = true
				break
			}

			if v.Order%10 != 0 {
				for _, v2 := range aptdongetc {
					if v.Order == v2.Order && v.Name == v2.Name {
						log.Println("FIND")
						flag = true
						break
					}
				}
			} else {
				order := v.Order / 10
				if order >= (dong.Parkingcount+dong.Undergroundcount)*-1 && order < dong.Undergroundcount*-1 {
					if v.Floortype == 1 {
						flag = true
					} else {
						flag = false
					}
					break
				}

				if order >= (dong.Undergroundcount)*-1 && order < 0 {
					if v.Floortype == 2 {
						flag = true
					} else {
						flag = false
					}
					break
				}

				start := 0
				if dong.Id != privateDong.Id {
					start = privateDong.Groundcount
				}

				if order <= dong.Groundcount && order > start {
					if v.Floortype == 3 {
						flag = true
					} else {
						flag = false
					}
					break
				}

				roofcount := 0

				if dong.Topcount == 0 {
					roofcount = 1
				}

				if dong.Groundcount == 0 {
					roofcount = 0
				}

				if order <= roofcount+dong.Groundcount && order > dong.Groundcount {
					if v.Floortype == 4 {
						flag = true
					} else {
						flag = false
					}
					break
				}

				if order <= dong.Topcount+dong.Groundcount && order > dong.Groundcount {
					if v.Floortype == 5 {
						flag = true
					} else {
						flag = false
					}
					break
				}
			}
		}

		if flag == false {
			blueprintManager.Delete(v.Id)
		}
	}

	if findLayout == false {
		item := models.Blueprint{Name: "배치도", Level: 1, Parent: 0, Floortype: 0, Filename: "", Upload: 1, Parentorder: 0, Order: 0, Aptdong: 0, Category: 1, Apt: id}
		blueprintManager.Insert(&item)
	}

	//var commonBlueprint models.Blueprint

	for _, dong := range aptdongs {
		var parent int64 = 0
		var parentOrder int = 0
		flag := false

		for _, v := range blueprints {
			if v.Aptdong != dong.Id {
				continue
			}

			if v.Parent == 0 {
				parent = v.Id

				v.Parent = 0
				v.Parentorder = dong.Order
				v.Order = 1000
				v.Name = dong.Dong

				if dong.Parkingcount+dong.Undergroundcount+dong.Groundcount+dong.Topcount+dong.Roofcount == 0 {
					v.Upload = 1
				} else {
					v.Upload = 0
				}

				/*
					if dong.Private == 4 {
						commonBlueprint = v
					}
				*/

				blueprintManager.Update(&v)

				flag = true
				break
			}
		}

		if flag == false {
			upload := 0
			if dong.Parkingcount+dong.Undergroundcount+dong.Groundcount+dong.Topcount+dong.Roofcount == 0 {
				upload = 1
			}

			item := models.Blueprint{Name: dong.Dong, Level: 1, Parent: 0, Floortype: 0, Filename: "", Upload: upload, Parentorder: dong.Order, Order: 1000, Aptdong: dong.Id, Category: 1, Apt: id}
			blueprintManager.Insert(&item)

			parent = blueprintManager.GetIdentity()

			if dong.Private == 4 {
				item.Id = parent
				//commonBlueprint = item
			}
		}

		parentOrder = dong.Order

		floors := make([]Floor, 0)
		for i := (dong.Parkingcount + dong.Undergroundcount) * -1; i < dong.Undergroundcount*-1; i++ {
			floor := Floor{Order: i, Floortype: 1}
			floors = append(floors, floor)

			for _, v2 := range aptdongetc {
				if dong.Id == v2.Aptdong && v2.Parent/10 == i {
					log.Println("INSERT 1")
					floor := Floor{Order: i, Floortype: 6, OriginalFloortype: v2.Floortype, Aptdongetc: v2}
					floors = append(floors, floor)
				}
			}
		}

		for i := (dong.Undergroundcount) * -1; i < 0; i++ {
			floor := Floor{Order: i, Floortype: 2}
			floors = append(floors, floor)

			for _, v2 := range aptdongetc {
				if dong.Id == v2.Aptdong && v2.Parent/10 == i {
					log.Println("INSERT 2")
					floor := Floor{Order: i, Floortype: 6, OriginalFloortype: v2.Floortype, Aptdongetc: v2}
					floors = append(floors, floor)
				}
			}
		}

		start := 1

		if dong.Private == 1 {
			start = privateDong.Groundcount + 1
		}

		for i := start; i <= dong.Groundcount; i++ {
			floor := Floor{Order: i, Floortype: 3}
			floors = append(floors, floor)

			for _, v2 := range aptdongetc {
				if dong.Id == v2.Aptdong && v2.Parent/10 == i {
					log.Println("INSERT 3")
					floor := Floor{Order: i, Floortype: 6, OriginalFloortype: v2.Floortype, Aptdongetc: v2}
					floors = append(floors, floor)
				}
			}
		}

		if dong.Private == 1 {
			roofcount := 0

			if dong.Topcount == 0 {
				roofcount = 1
			}

			if dong.Groundcount == 0 {
				roofcount = 0
			}

			for i := dong.Groundcount + 1; i <= roofcount+dong.Groundcount; i++ {
				floor := Floor{Order: i, Floortype: 4}
				floors = append(floors, floor)

				for _, v2 := range aptdongetc {
					if dong.Id == v2.Aptdong && v2.Parent/10 == i {
						log.Println("INSERT 4")
						floor := Floor{Order: i, Floortype: 6, OriginalFloortype: v2.Floortype, Aptdongetc: v2}
						floors = append(floors, floor)
					}
				}
			}

			for i := dong.Groundcount + 1; i <= dong.Topcount+dong.Groundcount; i++ {
				floor := Floor{Order: i, Floortype: 5}
				floors = append(floors, floor)

				for _, v2 := range aptdongetc {
					if dong.Id == v2.Aptdong && v2.Parent/10 == i {
						log.Println("INSERT 5")
						floor := Floor{Order: i, Floortype: 6, OriginalFloortype: v2.Floortype, Aptdongetc: v2}
						floors = append(floors, floor)
					}
				}
			}
		}

		for _, f := range floors {
			i := f.Order

			order := i * 10
			var item models.Blueprint
			floortype := blueprint.Floortype(f.Floortype)
			switch f.Floortype {
			case 1:
				item.Name = fmt.Sprintf("주차장 지하%v층", i*-1)
			case 2:
				item.Name = fmt.Sprintf("지하%v층", i*-1)
			case 3:
				item.Name = fmt.Sprintf("%v층", i)
			case 4:
				item.Name = "지붕층"
			case 5:
				if i-dong.Groundcount == 1 {
					item.Name = "옥탑/지붕층"
				} else {
					item.Name = fmt.Sprintf("옥탑 %v층", i-dong.Groundcount)
				}
			case 6:
				item.Name = f.Aptdongetc.Name
				order = f.Aptdongetc.Order
				floortype = blueprint.Floortype(f.Aptdongetc.Floortype)
			}

			item.Level = 2
			item.Parent = parent
			item.Floortype = floortype
			item.Upload = 1
			item.Parentorder = parentOrder
			item.Order = order
			item.Aptdong = dong.Id
			item.Category = 1
			item.Apt = id

			flag := false
			for _, v := range blueprints {
				if v.Aptdong != dong.Id {
					continue
				}

				if v.Order == order {
					item.Id = v.Id
					item.Filename = v.Filename
					item.Date = v.Date

					flag = true
					break
				}
			}

			if flag == true {
				blueprintManager.Update(&item)
			} else {
				blueprintManager.Insert(&item)
			}
		}
	}

	/*
		if commonDong.Id > 0 {
			for i := 0; i < len(commonTitles); i++ {
				if findCommons[i] == false {
					item := models.Blueprint{Name: commonTitles[i], Level: 2, Parent: commonBlueprint.Id, Floortype: 1, Filename: "", Upload: 1, Parentorder: commonDong.Order, Order: (1000 + i*10) * -1, Aptdong: commonDong.Id, Category: 1, Apt: id}
					blueprintManager.Insert(&item)
				}
			}
		}
	*/
}

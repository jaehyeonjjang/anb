package detail

import (
	"fmt"
	"log"
	"repair/global"
	"repair/global/config"
	"repair/models"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/LoperLee/golang-hangul-toolkit/hangul"
)

func Detail0(id int64, conn *models.Connection) string {
	facilitycategoryManager := models.NewFacilitycategoryManager(conn)
	periodicManager := models.NewPeriodicManager(conn)
	periodictechnicianManager := models.NewPeriodictechnicianManager(conn)
	aptManager := models.NewAptManager(conn)
	aptperiodicManager := models.NewAptperiodicManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	aptusagefloorManager := models.NewAptusagefloorManager(conn)
	periodiccheckManager := models.NewPeriodiccheckManager(conn)
	periodicpastManager := models.NewPeriodicpastManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)
	periodicimageManager := models.NewPeriodicimageManager(conn)
	periodicincidentalManager := models.NewPeriodicincidentalManager(conn)
	periodicouterwallManager := models.NewPeriodicouterwallManager(conn)
	periodicopinionManager := models.NewPeriodicopinionManager(conn)
	periodicchangeManager := models.NewPeriodicchangeManager(conn)

	periodicotherManager := models.NewPeriodicotherManager(conn)
	periodicotheretcManager := models.NewPeriodicotheretcManager(conn)

	others := periodicotherManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("po_order,po_id"),
	})

	periodicotheretc := periodicotheretcManager.GetByPeriodic(id)

	otherMap := make(map[int][]models.Periodicother)
	for _, v := range others {
		if _, ok := otherMap[v.Category]; !ok {
			otherMap[v.Category] = make([]models.Periodicother, 0)
		}

		otherMap[v.Category] = append(otherMap[v.Category], v)
	}

	facilitycategorys := facilitycategoryManager.Find([]any{
		models.Ordering("fc_order,fc_id"),
	})

	periodic := periodicManager.Get(id)
	aptperiodic := aptperiodicManager.Get(periodic.Apt)
	if aptperiodic == nil {
		aptperiodic = &models.Aptperiodic{}
	}

	apt := aptManager.Get(periodic.Apt)
	if periodic.Aptname != "" {
		apt.Name = periodic.Aptname
	}
	incidental := periodicincidentalManager.GetByPeriodic(id)
	if incidental == nil {
		incidental = &models.Periodicincidental{}
	}

	outerwall := periodicouterwallManager.GetByPeriodic(id)
	opinion := periodicopinionManager.GetByPeriodic(id)

	aptdongs := aptdongManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("au_order,au_id"),
	})

	blueprints := blueprintManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id"),
	})

	aptusagefloors := aptusagefloorManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("af_order,af_id"),
	})

	periodicchecks := periodiccheckManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
	})

	periodicpasts := periodicpastManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("pp_repairstartdate desc,pp_id desc"),
	})

	technicians := periodictechnicianManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("dt_order,dt_id"),
	})

	allPeriodicdatas := periodicdataManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "group", Value: 1, Compare: ">="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id,pd_order,pd_id"),
	})

	for i, v := range allPeriodicdatas {
		switch v.Type {
		case 3, 5:
			allPeriodicdatas[i].Type = 1
		case 4, 6:
			allPeriodicdatas[i].Type = 2
		}

		if v.Remark == "-" {
			allPeriodicdatas[i].Remark = ""
		}

		if v.Width == "-" {
			allPeriodicdatas[i].Width = ""
		}

		if v.Length == "-" {
			allPeriodicdatas[i].Length = ""
		}
	}

	periodicdatas := make([]models.Periodicdata, 0)

	inclinationPeriodicdatas := make([][]models.Periodicdata, 0)
	inclinationPeriodicdatasCount := 0
	inclinationPeriodicdatasPart := make([]models.Periodicdata, 0)

	fiberPeriodicdatas := make([][]models.Periodicdata, 0)
	fiberPeriodicdatasCount := 0
	fiberPeriodicdatasPart := make([]models.Periodicdata, 0)

	meterialPeriodicdatas := make([][]models.Periodicdata, 0)
	meterialPeriodicdatasCount := 0
	meterialPeriodicdatasPart := make([]models.Periodicdata, 0)

	for _, v := range allPeriodicdatas {
		if v.Type < 200 {
			periodicdatas = append(periodicdatas, v)
		} else if v.Type < 300 {
			inclinationPeriodicdatasPart = append(inclinationPeriodicdatasPart, v)
			inclinationPeriodicdatasCount++
			if inclinationPeriodicdatasCount == 20 {
				inclinationPeriodicdatas = append(inclinationPeriodicdatas, inclinationPeriodicdatasPart)
				inclinationPeriodicdatasPart = make([]models.Periodicdata, 0)
				inclinationPeriodicdatasCount = 0
			}
		} else if v.Type < 400 {
			fiberPeriodicdatasPart = append(fiberPeriodicdatasPart, v)
			fiberPeriodicdatasCount++
			if fiberPeriodicdatasCount == 24 {
				fiberPeriodicdatas = append(fiberPeriodicdatas, fiberPeriodicdatasPart)
				fiberPeriodicdatasPart = make([]models.Periodicdata, 0)
				fiberPeriodicdatasCount = 0
			}
		} else if v.Type < 500 {
			meterialPeriodicdatasPart = append(meterialPeriodicdatasPart, v)
			meterialPeriodicdatasCount++
			if meterialPeriodicdatasCount == 18 {
				meterialPeriodicdatas = append(meterialPeriodicdatas, meterialPeriodicdatasPart)
				meterialPeriodicdatasPart = make([]models.Periodicdata, 0)
				meterialPeriodicdatasCount = 0
			}
		}
	}

	if inclinationPeriodicdatasCount > 0 {
		inclinationPeriodicdatas = append(inclinationPeriodicdatas, inclinationPeriodicdatasPart)
		inclinationPeriodicdatasPart = make([]models.Periodicdata, 0)
	}

	if fiberPeriodicdatasCount > 0 {
		fiberPeriodicdatas = append(fiberPeriodicdatas, fiberPeriodicdatasPart)
		fiberPeriodicdatasPart = make([]models.Periodicdata, 0)
	}

	if meterialPeriodicdatasCount > 0 {
		meterialPeriodicdatas = append(meterialPeriodicdatas, meterialPeriodicdatasPart)
		meterialPeriodicdatasPart = make([]models.Periodicdata, 0)
	}

	for i, v := range meterialPeriodicdatas {
		var blueprint int64
		count := -1
		for j, v2 := range v {
			if v2.Blueprint == blueprint {
				meterialPeriodicdatas[i][j].Status = 2
				count++
			} else {
				meterialPeriodicdatas[i][j].Status = 1

				count++
				for k := 0; k < count; k++ {
					meterialPeriodicdatas[i][j-k-1].Progress = count
				}

				count = 0
			}

			blueprint = v2.Blueprint
		}

		if count > 0 {
			count++
			l := len(meterialPeriodicdatas[i])
			for k := 0; k < count; k++ {
				meterialPeriodicdatas[i][l-1-k].Progress = count
			}
		}
	}

	periodicimages := periodicimageManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "use", Value: 1, Compare: "="},
		models.Ordering("pi_order,pi_id"),
	})

	periodicchanges := periodicchangeManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("pc_order,pc_id"),
	})

	v := make(jet.VarMap)
	v.Set("name", apt.Name)
	v.Set("nameArray", strings.Split(apt.Name, ""))

	v.Set("apt", apt)
	v.Set("periodic", periodic)
	v.Set("aptperiodic", aptperiodic)
	v.Set("facilitycategorys", facilitycategorys)
	v.Set("aptdongs", aptdongs)

	v.Set("meterialPeriodicdatas", meterialPeriodicdatas)
	v.Set("inclinationPeriodicdatas", inclinationPeriodicdatas)

	aptusagefloorsPage := len(aptusagefloors) / 13

	if len(aptusagefloors)%13 > 0 {
		aptusagefloorsPage++
	}

	aptusagePages := make([]AptusagefloorPage, aptusagefloorsPage)

	cnt := 0

	for i := 0; i < aptusagefloorsPage; i++ {
		aptusagePages[i] = AptusagefloorPage{Count: 0, Items: make([]models.Aptusagefloor, 0)}

		for range 13 {
			aptusagePages[i].Items = append(aptusagePages[i].Items, aptusagefloors[cnt])
			aptusagePages[i].Count++
			cnt++

			if cnt == len(aptusagefloors) {
				break
			}
		}

		if cnt == len(aptusagefloors) {
			break
		}
	}

	v.Set("aptusagePageCount", aptusagefloorsPage)
	v.Set("aptusagePages", aptusagePages)
	v.Set("aptusagefloors", aptusagefloors)

	if aptusagefloors != nil {
		v.Set("aptusagefloorCount", len(aptusagefloors))
	} else {
		v.Set("aptusagefloorCount", 0)
	}

	var privateDong models.Aptdong
	onlyaptdongs := make([]models.Aptdong, 0)

	for _, v := range aptdongs {
		if v.Private == 3 {
			privateDong = v
		}

		if v.Private == 1 {
			onlyaptdongs = append(onlyaptdongs, v)
		}
	}

	v.Set("privateDong", privateDong)
	v.Set("onlyaptdongs", onlyaptdongs)
	v.Set("onlyaptdongCount", len(onlyaptdongs))

	resultaptdongs := make([]models.Aptdong, 0)

	if privateDong.Id == 0 {
		for _, v := range aptdongs {
			if v.Private != 1 {
				continue
			}

			resultaptdongs = append(resultaptdongs, v)
		}

	} else {
		maxFloor := 0
		for _, v := range aptdongs {
			if v.Groundcount > maxFloor {
				maxFloor = v.Groundcount
			}

			if v.Private != 3 {
				continue
			}

			resultaptdongs = append(resultaptdongs, v)
		}

		if len(resultaptdongs) > 0 {
			resultaptdongs[0].Groundcount = maxFloor
		}
	}

	v.Set("resultaptdongs", resultaptdongs)

	aptusagefloorCount := 0
	if aptusagefloors != nil {
		aptusagefloorCount = len(aptusagefloors)
	}
	v.Set("aptusagefloorCount", aptusagefloorCount)

	privateCount := 0
	for _, v2 := range aptdongs {
		if v2.Private == 1 {
			privateCount++
		}
	}

	v.Set("privateCount", privateCount)

	for _, v2 := range aptdongs {
		flag := false

		for _, v := range periodicchecks {
			if v.Aptdong == v2.Id {
				flag = true
				break
			}
		}

		if flag == false {
			var item models.Periodiccheck
			item.Periodic = id
			item.Aptdong = v2.Id
			item.Content1 = "⦁용도 변경사항 없음(이전 점검일 이후)"
			item.Content2 = "⦁구조부재 변경사항 없음"
			item.Content3 = "⦁증.개축 없음"
			item.Content4 = "⦁주변조건 변경사항 없음"
			item.Content5 = ""
			item.Content6 = ""
			item.Content7 = ""
			item.Content8 = ""
			item.Content9 = ""
			item.Content10 = "⦁전반적으로 양호함"
			item.Content11 = "⦁해당사항 없음"
			item.Content12 = "⦁해당사항 없음"
			item.Content13 = "⦁해당사항 없음"
			item.Content14 = "⦁전반적으로 양호함"
			item.Content15 = "⦁특기사항 없음"
			item.Content16 = "⦁발생된 결함 사항에 대하여 보수·보강 우선순위 선정 후 유지관리계획에 따라 보수하여 관리하는 것이 바람직함"
			item.Use1 = 1
			item.Use2 = 1
			item.Use3 = 1
			item.Use4 = 1
			item.Need1 = 2
			item.Need2 = 2
			item.Need3 = 2
			item.Need4 = 2

			periodicchecks = append(periodicchecks, item)
		}
	}

	v.Set("periodicchecks", periodicchecks)

	pastdate := ""
	if len(periodicpasts) > 0 {
		pastdate = periodicpasts[0].Repairenddate
	}
	v.Set("periodicpasts", periodicpasts)
	v.Set("pastdate", pastdate)
	v.Set("pastCounts", len(periodicpasts))

	mastertechnician := ""
	mastertechnicianname := ""
	if len(technicians) > 0 {
		v := technicians[0]

		typeid := ""
		if v.Type == 1 {
			typeid = "책임기술자"
		} else {
			typeid = "참여기술자"
		}

		grade := ""

		user := v.Extra["technician"].(models.Technician)
		switch user.Grade {
		case 1:
			grade = "건축초급기술자"
		case 2:
			grade = "건축중급기술자"
		case 3:
			grade = "건축고급기술자"
		case 4:
			grade = "건축특급기술자"
		}

		if v.Type == 1 {
			mastertechnicianname = strings.Join(strings.Split(user.Name, ""), "   ")
			name := strings.Join(strings.Split(user.Name, ""), " ")
			mastertechnician = fmt.Sprintf("%v - %v (인) (자격 : %v)", typeid, name, grade)
		}
	}

	v.Set("mastertechnician", mastertechnician)
	v.Set("mastertechnicianname", mastertechnicianname)

	dongCount := 0

	for _, v := range aptdongs {
		if v.Private != 1 {
			continue
		}

		dongCount++
	}

	if dongCount == 0 {
		dongCount = 1
	}

	underground := make([]int, 0)
	ground := make([]int, 0)
	undergroundFlag := false

	for _, v := range aptdongs {
		if v.Parkingcount+v.Undergroundcount > 0 {
			undergroundFlag = true
		}

		if v.Parkingcount+v.Undergroundcount > 0 {
			underground = append(underground, v.Parkingcount+v.Undergroundcount)
		}

		if v.Private == 1 {
			if v.Groundcount > 0 {
				ground = append(ground, v.Groundcount)
			}
		}
	}

	underground = global.Unique(underground)
	ground = global.Unique(ground)

	dongText := ""
	if undergroundFlag == true {
		dongText = fmt.Sprintf("(지하%v층~지상%v층)", JoinInt(underground), JoinInt(ground))
	} else {
		dongText = fmt.Sprintf("(지상%v층)", JoinInt(ground))
	}

	v.Set("dongCount", dongCount)
	v.Set("dongText", dongText)

	useapproval := global.ParseDatetime(fmt.Sprintf("%v 00:00:00", apt.Useapproval))
	rdate := global.ParseDatetime(fmt.Sprintf("%v 00:00:00", periodic.Reportdate))

	dyear := 0
	dmonth := 0

	if useapproval != nil && rdate != nil {
		dyear, dmonth, _, _, _, _ = global.DiffTime(*useapproval, *rdate)
	}

	progressDate := ""
	if dmonth == 0 {
		progressDate = fmt.Sprintf("%v년", dyear)
	} else {
		progressDate = fmt.Sprintf("%v년 %v개월", dyear, dmonth)
	}

	v.Set("progressDate", progressDate)

	reportdateArray := make([]string, 0)
	reportdateTemp := strings.Split(periodic.Reportdate, "-")
	reportdate := reportdateTemp[0]
	if reportdate != "" {
		reportdate += "."
	}

	reportdateHalf := reportdateTemp[0] + "년도"

	reportdateArray = append(reportdateArray, strings.Split(reportdateTemp[0], "")...)
	reportdateArray = append(reportdateArray, ".")
	if len(reportdateTemp) > 1 {
		if global.Atoi(reportdateTemp[1]) < 7 {
			reportdateHalf += " 상반기"
		} else {
			reportdateHalf += " 하반기"
		}
		reportdate += " " + reportdateTemp[1] + "."
		reportdateArray = append(reportdateArray, reportdateTemp[1])
	}

	v.Set("reportdate", reportdate)
	v.Set("reportdateArray", reportdateArray)
	v.Set("manager", periodic.Manager)
	v.Set("position", apt.Position)
	v.Set("reportdateHalf", reportdateHalf)
	v.Set("reportdateHalf2", strings.ReplaceAll(reportdateHalf, "년도", "년"))

	if technicians != nil {
		v.Set("technicianCount", len(technicians)+1)
	} else {
		v.Set("technicianCount", 1)
	}
	v.Set("technicians", technicians)

	resulttechnicians := make([]models.Periodictechnician, 0)
	for i, v := range technicians {
		resulttechnicians = append(resulttechnicians, v)

		if i == 2 {
			break
		}
	}

	for i := 0; i < 3-len(resulttechnicians); i++ {
		v := models.Periodictechnician{}
		t := models.Technician{}
		v.Extra = make(map[string]any)
		v.Extra["technician"] = t
		resulttechnicians = append(resulttechnicians, v)
	}

	v.Set("resulttechnicians", resulttechnicians)
	v.Set("resultTechnicianCount", len(resulttechnicians)+15)

	otherData := Other(others)
	summaryItems := GetSummary(periodic.Id, *periodic, aptdongs, periodicdatas, blueprints, periodicchanges, otherData, otherMap)
	v.Set("summarys", summaryItems)

	// 강재 노후
	steelstructureCount := 0
	steelstructures := make([]string, 0)
	for _, v := range otherMap[15] {
		if v.Order == 150 {
			if v.Status == "해당사항 없음" {
				steelstructures = append(steelstructures, "⦁"+v.Status)
				break
			}

			continue
		}

		if v.Status == "" {
			continue
		}

		str := fmt.Sprintf("⦁%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", "))

		if steelstructureCount%2 == 1 {
			steelstructures[len(steelstructures)-1] = steelstructures[len(steelstructures)-1] + "             " + str
		} else {
			steelstructures = append(steelstructures, str)
		}
		steelstructureCount++
	}

	if len(steelstructures) == 0 {
		steelstructures = append(steelstructures, "⦁전반적으로 양호함")
	}
	v.Set("steelstructures", steelstructures)

	// 도로포장
	roadCount := 0
	roads := make([]string, 0)
	for _, v := range otherMap[11] {
		if v.Order == 110 {
			continue
		}

		if v.Status == "" {
			continue
		}

		str := fmt.Sprintf("⦁%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", "))

		if roadCount%2 == 1 {
			roads[len(roads)-1] = roads[len(roads)-1] + "             " + str
		} else {
			roads = append(roads, str)
		}
		roadCount++
	}

	// 옹벽, 담장
	retailwall := "⦁해당사항 없음"
	fence := "⦁해당사항 없음"
	for _, v := range otherMap[3] {
		switch v.Order {
		case 203:
			if v.Position == "" {
				if v.Status != "" {
					retailwall = fmt.Sprintf("⦁%v", v.Status)
				}
			} else {
				retailwall = fmt.Sprintf("⦁%v - %v", v.Position, v.Status)
			}
		case 218:
			if v.Position == "" {
				if v.Status != "" {
					fence = fmt.Sprintf("⦁%v", v.Status)
				}
			} else {
				fence = fmt.Sprintf("⦁%v - %v", v.Position, v.Status)
			}
		}
	}

	log.Println("############## STEP 1")

	v.Set("retailwall", retailwall)
	v.Set("fence", fence)

	if len(roads) == 0 {
		roads = append(roads, "⦁전반적으로 양호함")
	}
	v.Set("roads", roads)

	endopinion := Endopinion(aptdongs, periodicdatas)
	v.Set("endopinion", endopinion)
	causes, otherResult := GetCause(periodicdatas, aptdongs, otherData, otherMap)
	v.Set("causes", causes)

	v.Set("other10", otherResult[10])
	v.Set("other11", otherResult[11])
	v.Set("other12", otherResult[12])
	v.Set("other13", otherResult[13])
	v.Set("other14", otherResult[14])

	otherFlag := true
	for i := 10; i <= 14; i++ {
		if otherResult[i].Good == 0 {
			otherFlag = false
			break
		}
	}

	log.Println("############## STEP 2")

	otherStrs := make([]string, 0)
	endStrs := make([]string, 0)
	if otherFlag == true {
		otherStrs = append(otherStrs, "부대시설 조사결과 외부 D/A, 보도블럭, 목재데크, 휀스 등의 상태는 양호한 것으로 확인되었다.")
		otherStrs = append(otherStrs, "관리주체는 현재와 같은 상태를 유지하기 위해 지속적인 점검 및 유지관리계획 등의 노력을 진행해야 할 것이다.")
	} else {
		type otherItem struct {
			Position []string
			Result   []string
			Title    string
			Content  string
		}

		otherItems := make([]otherItem, 0)

		titleHeads := []string{"추락방지시설", "지반(포장)", "도로부 신축 이음부", "D/A"}
		for i := 10; i < 14; i++ {
			positions := make([]string, 0)
			titles := make([]string, 0)
			results := make([]string, 0)
			for _, v := range others {
				if v.Category != i {
					continue
				}
				if v.Status == "" {
					continue
				}

				if v.Type != 2 {
					continue
				}

				titles = append(titles, v.Position)

				if i == 10 && len(positions) == 0 {
					v.Position = "일부 " + v.Position
				}
				positions = append(positions, v.Position)

				results = append(results, strings.Split(v.Status, ",")...)
			}

			if len(positions) > 0 && len(results) > 0 {
				title := fmt.Sprintf("%v의 (%v)", titleHeads[i-10], strings.Join(titles, "/"))
				content := fmt.Sprintf("(%v)", strings.Join(global.UniqueStringWithoutSort(results), ", "))

				if i == 10 {
					content = "일부 " + content
				}

				item := otherItem{Position: positions, Result: results, Title: title, Content: content}
				otherItems = append(otherItems, item)
			}
		}

		for _, v := range others {
			if v.Category != 3 {
				continue
			}
			if v.Status == "" {
				continue
			}

			item := otherItem{Position: []string{v.Position}, Result: []string{v.Status}, Title: v.Position, Content: v.Status}
			otherItems = append(otherItems, item)
		}

		strs := "부대시설 조사결과 "
		for _, v := range otherItems {
			strs += fmt.Sprintf("%v에서 %v가 ", strings.Join(v.Position, ", "), strings.Join(global.UniqueStringWithoutSort(v.Result), "/"))
		}

		endStr := ""
		if len(otherItems) == 1 {
			v := otherItems[0]
			endStr = fmt.Sprintf("%v에서 %v 조사되었다.", v.Title, Josa(v.Content, hangul.I_GA))
		} else {
			for i, v := range otherItems {
				if i == len(otherItems)-1 {
					endStr += fmt.Sprintf("%v에서 %v 조사되었다.", v.Title, Josa(v.Content, hangul.I_GA))
				} else {
					endStr += fmt.Sprintf("%v에서 %v ", v.Title, Josa(v.Content, hangul.GWA_WA))
				}

				if i == 0 {
					endStr = "바. " + endStr
				}

			}
		}

		endStrs = append(endStrs, endStr)
		endStrs = append(endStrs, "   관리주체는 결함이 발생된 부분에 대해서는 적절한 보수를 진행해야 할 것이다.")

		strs += "확인되었다."

		otherStrs = append(otherStrs, strs)
		otherStrs = append(otherStrs, "관리주체는 결함이 발생한 부분에 대해서는 적절한 보수를 진행해야 할 것이며, 양호한 부분에 대해서는 현재와 같은 상태를 유지하기 위해 지속적인 점검 및 유지관리계획 등의 노력을 진행해야 할 것이다.")
	}

	v.Set("endStrs", endStrs)
	v.Set("otherStrs", otherStrs)

	log.Println("############## STEP 3")

	periodicchangeCount := []int{0, 0, 0, 0, 0, 0}

	for _, v := range periodicchanges {
		periodicchangeCount[v.Type]++
	}

	changeRows := []int{6, 6, 3, 3, 1}

	for i, v := range changeRows {
		pos := i + 1
		if periodicchangeCount[pos] < v {
			for i := 0; i < v-periodicchangeCount[pos]; i++ {
				periodicchanges = append(periodicchanges, models.Periodicchange{Type: pos})
			}
		}
	}

	periodicchangeCount = []int{0, 0, 0, 0, 0, 0}

	log.Println("############## STEP 4")
	for _, v := range periodicchanges {
		periodicchangeCount[v.Type]++
	}

	for i := 1; i <= 5; i++ {
		items := make([]models.Periodicchange, 0)
		v.Set(fmt.Sprintf("periodicchangeCount%v", i), periodicchangeCount[i])

		for _, v := range periodicchanges {
			if v.Type == i {
				items = append(items, v)
			}
		}

		if i == 3 {
			if items[0].Content1 == "" {
				items[0].Content1 = "사용하중"
			}

			if items[1].Content1 == "" {
				items[1].Content1 = "기초 및 지반조건"
			}

			if items[2].Content1 == "" {
				items[2].Content1 = "주변환경"
			}
		}

		v.Set(fmt.Sprintf("periodicchanges%v", i), items)
	}

	datalists := make([]PeriodicDataList, 0)
	datas := make([]PeriodicData, 0)

	dataFirst := true
	dataCount := 0

	for _, blueprint := range blueprints {
		if blueprint.Upload != 1 {
			continue
		}

		subdatas := make([]PeriodicData, 0)
		for _, v := range periodicdatas {
			if v.Blueprint != blueprint.Id {
				continue
			}

			subdatas = append(subdatas, PeriodicData{Type: 2, Title: "", Item: v})
		}

		if len(subdatas) > 0 {
			var parent models.Blueprint

			for _, v2 := range blueprints {
				if blueprint.Parent == v2.Id {
					parent = v2
					break
				}
			}

			title := fmt.Sprintf("%v %v", parent.Name, blueprint.Name)
			if privateDong.Id > 0 {
				for _, aptdong := range aptdongs {
					if blueprint.Aptdong != aptdong.Id {
						continue
					}

					if aptdong.Private == 3 {
						title = blueprint.Name
					}

					break
				}
			}
			datas = append(datas, PeriodicData{Type: 1, Title: title, Item: models.Periodicdata{}})

			dataCount++

			for i, v2 := range subdatas {
				datas = append(datas, v2)
				dataCount++

				last := false
				if i == len(subdatas)-1 {
					last = true
				}
				if (dataCount >= 19 && dataFirst == true) || dataCount == 20 || (dataCount >= 18 && dataFirst == true && last == true) || (dataCount == 19 && last == true) {
					datas = append(datas, PeriodicData{Type: 3, Title: "", Item: models.Periodicdata{}})
					datalists = append(datalists, PeriodicDataList{Items: datas})
					datas = make([]PeriodicData, 0)

					dataFirst = false

					dataCount = 0
				}
			}

		}
	}

	if dataCount > 0 {
		datas = append(datas, PeriodicData{Type: 3, Title: "", Item: models.Periodicdata{}})
		datalists = append(datalists, PeriodicDataList{Items: datas})
	}

	log.Println("############## STEP 5")
	v.Set("periodicdatas", datalists)

	images := make([]ImageInfo, 0)

	imageCount := 0
	for _, v := range periodicimages {
		if v.Type != 1 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))

		imageCount++
		if imageCount == 2 {
			break
		}
	}

	for i := 0; i < 2-imageCount; i++ {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	// 위치도
	positionimages := make([]ImageInfo, 0)
	positionimageCount := 0
	for _, v := range periodicimages {
		if v.Type != 1 {
			continue
		}

		info := GetImageInfo(v.Filename)
		info.Text = v.Name
		positionimages = append(positionimages, info)

		positionimageCount++

		if positionimageCount == 2 {
			break
		}
	}

	if positionimageCount == 0 {
		positionimages = append(positionimages, GetImageInfo("empty.jpg"))
		positionimageCount++
	}

	log.Println("############## STEP 6")
	// 전경 사진

	sightimages := make([]ImageInfo, 0)
	sightimageCount := 0
	for _, v := range periodicimages {
		if v.Type != 2 {
			continue
		}

		info := GetImageInfo(v.Filename)
		info.Text = v.Name
		sightimages = append(sightimages, info)

		sightimageCount++
	}

	if sightimageCount%2 > 0 {
		for i := 0; i < 2-(sightimageCount%2); i++ {
			sightimages = append(sightimages, GetImageInfo("empty.jpg"))
		}

		sightimageCount += 2 - (sightimageCount % 2)
	}

	if sightimageCount == 0 {
		for range 2 {
			sightimages = append(sightimages, GetImageInfo("empty.jpg"))
		}

		sightimageCount = 2
	}

	groupSightimages := make([]PartImageInfo, 0)

	for i := 0; i < sightimageCount/2; i++ {
		var group PartImageInfo
		group.Image1 = sightimages[i*2+0]
		group.Image2 = sightimages[i*2+1]

		groupSightimages = append(groupSightimages, group)
	}

	v.Set("groupsightimages", groupSightimages)
	v.Set("positionimages", positionimages)
	v.Set("positionimageCount", positionimageCount)
	v.Set("sightimages", sightimages)
	v.Set("sightimageCount", sightimageCount)

	log.Println("############## STEP 7")
	// 부위별 사진

	partimages := make([]ImageInfo, 0)
	imageCount = 0
	for _, v := range periodicimages {
		if v.Type != 3 {
			continue
		}

		info := GetImageInfo(v.Filename)
		info.Text = v.Name
		partimages = append(partimages, info)

		imageCount++
	}

	if imageCount%6 > 0 {
		for i := 0; i < 6-(imageCount%6); i++ {
			partimages = append(partimages, GetImageInfo("empty.jpg"))
		}

		imageCount += 6 - (imageCount % 6)
	}

	if imageCount == 0 {
		for range 6 {
			partimages = append(partimages, GetImageInfo("empty.jpg"))
		}

		imageCount = 6
	}

	groupPartimages := make([]PartImageInfo, 0)

	for i := 0; i < imageCount/6; i++ {
		var group PartImageInfo
		group.Image1 = partimages[i*6+0]
		group.Image2 = partimages[i*6+1]
		group.Image3 = partimages[i*6+2]
		group.Image4 = partimages[i*6+3]
		group.Image5 = partimages[i*6+4]
		group.Image6 = partimages[i*6+5]

		groupPartimages = append(groupPartimages, group)
	}

	v.Set("grouppartimages", groupPartimages)
	v.Set("partimages", partimages)
	v.Set("partimageCount", imageCount)

	privateDongPages := 0
	publicDongPages := 0
	oneDongFlag := false
	for _, v := range aptdongs {
		switch v.Private {
		case 3:
			oneDongFlag = true
		case 2:
			publicDongPages = 1
		default:
			privateDongPages++
		}
	}

	if oneDongFlag == true {
		privateDongPages = 1
	}

	dongPages := publicDongPages + privateDongPages
	v.Set("startDongPages", 2+len(groupSightimages)+len(groupPartimages))
	v.Set("dongPages", dongPages)

	stampFlag := false
	for _, v := range technicians {
		if v.Type != 1 {
			continue
		}

		images = append(images, GetImageInfo(v.Extra["technician"].(models.Technician).Stamp))
		images = append(images, GetImageInfo(v.Extra["technician"].(models.Technician).Stamp))

		stampFlag = true
		break
	}

	if stampFlag == false {
		images = append(images, GetImageInfo("empty.jpg"))
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("images", images)
	v.Set("incidental", incidental)
	v.Set("outerwall", outerwall)
	v.Set("opinion", opinion)

	v.Set("periodicotheretc", periodicotheretc)
	v.Set("others1", otherMap[1])
	v.Set("others2", otherMap[2])
	v.Set("others3", otherMap[3])
	v.Set("others3Count", len(otherMap[3]))

	if privateDong.Id > 0 || len(onlyaptdongs) <= 1 {
		v.Set("sizetype", 1)
	} else {
		v.Set("sizetype", 2)
	}

	for _, item := range others {
		if item.Category == 3 && item.Order == 201 {

			n := make([]string, 0)
			d := make([]string, 0)
			for _, v := range otherData[11].Items {
				temp := strings.Split(v, " - ")

				if len(temp) < 2 {
					continue
				}

				n = append(n, temp[0])
				d = append(d, strings.ReplaceAll(temp[1], ", ", "/"))
			}

			str := ""

			if len(n) > 0 {
				str = fmt.Sprintf("%v %v", strings.Join(n, "/"), strings.Join(global.UniqueStringWithoutSort(d), "/"))
			}

			v.Set("other201", str)
		}

		if item.Category == 3 && item.Order == 203 {
			v.Set("other203", item)
		}

		if item.Category == 3 && item.Order == 218 {
			v.Set("other218", item)
		}
	}

	nameWithJosa := ""
	if apt.Name != "" {
		nameWithJosa = global.GetJosa(apt.Name, hangul.EUN_NEUN)
	}

	v.Set("namejosa", nameWithJosa)

	log.Println("############## STEP 8")
	template := GetTemplate("detail/detail-00.jet", v)

	filename := fmt.Sprintf("detail/detail-00-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	err := global.WriteFile(fullFilename, template)
	if err != nil {
		log.Println(err)
	}

	log.Println("############## STEP 9")
	log.Println(fullFilename)
	return filename
}

package periodic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"repair/global"
	"repair/global/config"
	"repair/models"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/LoperLee/golang-hangul-toolkit/hangul"
)

type PeriodicData struct {
	Type  int
	Title string
	Item  models.Periodicdata
}

type PeriodicDataList struct {
	Items []PeriodicData
}

type ImageInfo struct {
	Size     int64
	Data     string
	Text     string
	Status   string
	Number   string
	Filename string
	Ext      string
	Width    int
	Height   int
}

type PartImageInfo struct {
	Image1 ImageInfo
	Image2 ImageInfo
	Image3 ImageInfo
	Image4 ImageInfo
	Image5 ImageInfo
	Image6 ImageInfo
}

type ImageDouble struct {
	First  bool
	Second bool
}

type ImageSizeInfo struct {
	Width   int
	Height  int
	Top     int
	Left    int
	Width2  int
	Height2 int
}

type ResultMap struct {
	Dong  string
	Items []string
}

type AptusagefloorPage struct {
	Count int
	Items []models.Aptusagefloor
}

func GetThumbnail(filename string) string {
	ext := global.GetExt(filename)

	newFilename := fmt.Sprintf("temp/%v-thumbnail.%v", strings.ReplaceAll(filepath.Base(filename), "."+ext, ""), ext)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, newFilename)
	_, err := os.Stat(fullFilename)
	if err != nil {
		global.MakeThumbnail(21780/50, 16335/50, fmt.Sprintf("%v/%v", config.UploadPath, filename), fmt.Sprintf("%v/%v", config.UploadPath, newFilename))
	}
	return newFilename
}

func Josa(str string, josa hangul.Josa) string {
	if str == "" {
		return ""
	}
	temp := strings.Split(str, "(")

	if len(temp) > 1 {
		if temp[0] == "" {
			temp = strings.Split(temp[1], ")")
		}

		t := global.GetJosa(temp[0], josa)

		josaStr := strings.ReplaceAll(t, temp[0], "")
		return str + josaStr
	}

	return global.GetJosa(str, josa)
}

func JoinInt(items []int) string {
	strs := make([]string, 0)

	for _, v := range items {
		strs = append(strs, fmt.Sprintf("%v", v))
	}

	return strings.Join(strs, "/")
}

func GetImageSize(filename string) (int, int) {
	fullFilename := filepath.Join(config.UploadPath, filename)
	return global.GetImageSize(fullFilename)
}

func GetImageInfo(filename string) ImageInfo {
	if filename == "" || filename == "webdata" {
		filename = "empty.jpg"
	}

	fullFilename := filepath.Join(config.UploadPath, filename)

	fi, err := os.Stat(fullFilename)
	if err != nil {
		log.Print("not found")
		fullFilename = filepath.Join(config.UploadPath, "empty.jpg")
		fi, _ = os.Stat(fullFilename)
	}
	/*
		b64Filename := fmt.Sprintf("%v.b64", filename)
		b64fullFilename := filepath.Join(config.UploadPath, b64Filename)

		log.Println("b64filenaem", b64fullFilename)

		dec := ""
		_, err2 := os.Stat(b64fullFilename)
		if err2 != nil {
			dat, _ := ioutil.ReadFile(fullFilename)
			dec := base64.StdEncoding.EncodeToString(dat)
			global.WriteFile(b64fullFilename, dec)
		} else {
			dat, _ := ioutil.ReadFile(b64fullFilename)
			dec = string(dat)
		}
	*/

	dat, _ := os.ReadFile(fullFilename)
	dec := base64.StdEncoding.EncodeToString(dat)

	width, height := GetImageSize(filename)
	info := ImageInfo{Filename: filename, Size: fi.Size(), Data: dec, Width: width, Height: height, Ext: global.GetExt(filename)}
	return info
}

func GetStructName(value int) string {
	if value == 1 {
		return "철근콘크리트 구조"
	} else if value == 2 {
		return "철골.철근콘크리트 구조"
	} else if value == 3 {
		return "철근콘크리트라멘조+벽식 구조"
	} else if value == 4 {
		return "철골.철근콘크리트구조 및 복합구조"
	} else if value == 5 {
		return "프리케스트콘크리트 구조"
	} else if value == 6 {
		return "철근콘크리트 PC구조"
	}

	return ""
}

func Periodic0(id int64, conn *models.Connection) string {
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

	others := periodicotherManager.Find([]interface{}{
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

	facilitycategorys := facilitycategoryManager.Find([]interface{}{
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
	if opinion == nil {
		opinion = &models.Periodicopinion{}
	}

	aptdongs := aptdongManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("au_order,au_id"),
	})

	blueprints := blueprintManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id"),
	})

	aptusagefloors := aptusagefloorManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("af_order,af_id"),
	})

	periodicchecks := periodiccheckManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
	})

	periodicpasts := periodicpastManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("pp_repairstartdate desc,pp_id desc"),
	})

	technicians := periodictechnicianManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("dt_order,dt_id"),
	})

	allPeriodicdatas := periodicdataManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "group", Value: 1, Compare: ">="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id,pd_order,pd_id"),
	})

	for i, v := range allPeriodicdatas {
		if v.Type == 3 || v.Type == 5 {
			allPeriodicdatas[i].Type = 1
		} else if v.Type == 4 || v.Type == 6 {
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

	for _, v := range allPeriodicdatas {
		if v.Type >= 200 {
			continue
		}

		periodicdatas = append(periodicdatas, v)
	}

	periodicimages := periodicimageManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "use", Value: 1, Compare: "="},
		models.Ordering("pi_order,pi_id"),
	})

	periodicchanges := periodicchangeManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("pc_order,pc_id"),
	})

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	view.AddGlobal("date", func(str string) string {
		return strings.ReplaceAll(str, "-", ".")
	})

	view.AddGlobal("year", func(str string) string {
		strs := strings.Split(str, "-")
		if strs[0] != "" {
			return strs[0] + "."
		}

		return ""
	})

	view.AddGlobal("monthday", func(str string) string {
		strs := strings.Split(str, "-")
		if len(strs) == 3 {
			return strs[1] + ". " + strs[2] + "."
		}

		return ""
	})

	view.AddGlobal("split", func(str string) []string {
		return strings.Split(str, "\n")
	})

	view.AddGlobal("humandate", func(str string) string {
		temp := strings.Split(str, "-")

		if len(temp) != 3 {
			return str
		}

		return fmt.Sprintf("%04d년 %02d월 %02d일", global.Atoi(temp[0]), global.Atoi(temp[1]), global.Atoi(temp[2]))
	})

	view.AddGlobal("without", func(str string) string {
		return strings.ReplaceAll(str, "동", "")
	})

	view.AddGlobal("dataCount", func(items []PeriodicData) int {
		return len(items) + 2
	})

	view.AddGlobal("pastCount", func(items []models.Periodicpast) int {
		return (len(items) + 1) * 2
	})

	view.AddGlobal("usageCount", func(items []models.Aptdong) int {
		count := 0

		for _, v := range items {
			if v.Private != 1 {
				continue
			}

			count++
		}
		return count + 1
	})

	view.AddGlobal("usagefloorCount", func(items []models.Aptusagefloor) int {
		return len(items) + 1
	})

	view.AddGlobal("dash", func(str string) string {
		if str == "" {
			return "-"
		}

		return str
	})

	view.AddGlobal("text", func(str string) string {
		if str == "" {
			return " "
		}

		return str
	})

	v := make(jet.VarMap)
	v.Set("name", apt.Name)
	v.Set("nameArray", strings.Split(apt.Name, ""))

	v.Set("apt", apt)
	v.Set("periodic", periodic)
	v.Set("aptperiodic", aptperiodic)
	v.Set("facilitycategorys", facilitycategorys)
	v.Set("aptdongs", aptdongs)

	aptusagefloorsPage := len(aptusagefloors) / 13

	if len(aptusagefloors)%13 > 0 {
		aptusagefloorsPage++
	}

	aptusagePages := make([]AptusagefloorPage, aptusagefloorsPage)

	cnt := 0

	for i := 0; i < aptusagefloorsPage; i++ {
		aptusagePages[i] = AptusagefloorPage{Count: 0, Items: make([]models.Aptusagefloor, 0)}

		for j := 0; j < 13; j++ {
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
		if user.Grade == 1 {
			grade = "건축초급기술자"
		} else if user.Grade == 2 {
			grade = "건축중급기술자"
		} else if user.Grade == 3 {
			grade = "건축고급기술자"
		} else if user.Grade == 4 {
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

	v.Set("resultTechnicianCount", len(technicians)+15)

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
		if v.Order == 203 {
			if v.Position == "" {
				if v.Status != "" {
					retailwall = fmt.Sprintf("⦁%v", v.Status)
				}
			} else {
				retailwall = fmt.Sprintf("⦁%v - %v", v.Position, v.Status)
			}
		} else if v.Order == 218 {
			if v.Position == "" {
				if v.Status != "" {
					fence = fmt.Sprintf("⦁%v", v.Status)
				}
			} else {
				fence = fmt.Sprintf("⦁%v - %v", v.Position, v.Status)
			}
		}
	}

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

	imageCount = 0
	for _, v := range periodicimages {
		if v.Type != 2 {
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

	for i := 0; i < 6; i++ {
		images = append(images, GetImageInfo("empty.jpg"))
	}

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
		for i := 0; i < 2; i++ {
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
	v.Set("sightimages", sightimages)
	v.Set("sightimageCount", sightimageCount)

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
		for i := 0; i < 6; i++ {
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
		if v.Private == 3 {
			oneDongFlag = true
		} else if v.Private == 2 {
			publicDongPages = 1
		} else {
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

	var b bytes.Buffer
	t, err := view.GetTemplate("periodic/periodic-00.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("periodic/periodic-00-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Periodic1(id int64, conn *models.Connection) string {
	aptManager := models.NewAptManager(conn)
	periodicManager := models.NewPeriodicManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	aptperiodicManager := models.NewAptperiodicManager(conn)

	periodic := periodicManager.Get(id)
	apt := aptManager.Get(periodic.Apt)
	if periodic.Aptname != "" {
		apt.Name = periodic.Aptname
	}
	aptperiodic := aptperiodicManager.Get(periodic.Apt)
	aptdongs := aptdongManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
	})

	v := make(jet.VarMap)

	v.Set("apt", apt)
	v.Set("aptperiodic", aptperiodic)
	v.Set("periodic", periodic)

	underground := 0
	ground := 0
	for _, v := range aptdongs {
		if v.Private == 2 {
			continue
		}
		if v.Undergroundcount+v.Parkingcount > underground {
			underground = v.Undergroundcount + v.Parkingcount
		}

		if v.Groundcount > ground {
			ground = v.Groundcount
		}
	}

	floor := ""
	if underground > 0 && ground > 0 {
		floor = fmt.Sprintf("지하%v층, 지상%v층", underground, ground)
	} else if underground > 0 {
		floor = fmt.Sprintf("지하%v층", underground)
	} else {
		floor = fmt.Sprintf("지상%v층", ground)
	}

	v.Set("floor", floor)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	view.AddGlobal("humandate", func(str string) string {
		temp := strings.Split(str, "-")

		if len(temp) != 3 {
			return str
		}

		return fmt.Sprintf("%04d년 %02d월 %02d일", global.Atoi(temp[0]), global.Atoi(temp[1]), global.Atoi(temp[2]))
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("periodic/periodic-01.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("periodic/periodic-01-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Periodic2(id int64, conn *models.Connection) string {
	periodicManager := models.NewPeriodicManager(conn)
	periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)

	periodic := periodicManager.Get(id)

	blueprints := blueprintManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id"),
	})

	zooms := periodicblueprintzoomManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
	})

	images := make([]ImageInfo, 0)
	titles := make([]string, 0)
	imagesizes := make([]ImageSizeInfo, 0)

	for _, v := range blueprints {
		if v.Upload != 1 {
			continue
		}

		find := false
		for _, v2 := range zooms {
			if v.Id == v2.Blueprint {
				find = true
				break
			}
		}

		filename := ""
		if find == true {
			filename = fmt.Sprintf("periodicresult/%v/%v.jpg", id, v.Id)
		} else {
			filename = v.Filename
		}

		if filename == "" || filename == "webata" {
			continue
		}

		images = append(images, GetImageInfo(filename))

		width, height := GetImageSize(filename)

		w := 397600
		h := 281200

		w2 := 40479
		h2 := 28629

		newWidth := 0
		newHeight := 0

		newWidth2 := 0
		newHeight2 := 0

		rate := float64(w) / float64(h)
		target := float64(width) / float64(height)

		if rate > target {
			newWidth = int(float64(h) * target)
			newHeight = h

			newWidth2 = int(float64(h2) * target)
			newHeight2 = h2
		} else if rate < target {
			newWidth = w
			newHeight = int(float64(w) / target)

			newWidth2 = w2
			newHeight2 = int(float64(w2) / target)
		} else {
			newWidth = w
			newHeight = h

			newWidth2 = w2
			newHeight2 = h2
		}

		var info ImageSizeInfo

		info.Width = newWidth
		info.Height = newHeight
		info.Width2 = newWidth2
		info.Height2 = newHeight2
		info.Left = 0
		info.Top = 0

		imagesizes = append(imagesizes, info)

		title := v.Name

		if v.Level == 2 {
			for _, v3 := range blueprints {
				if v3.Id == v.Parent {
					title = fmt.Sprintf("%v %v", v3.Name, v.Name)
					break
				}
			}
		}

		titles = append(titles, title)
	}

	v := make(jet.VarMap)

	v.Set("imagecount", len(images))
	v.Set("images", images)

	items := make([]ImageDouble, 0)

	count := int(len(images) / 2)
	for i := 0; i < count; i++ {
		items = append(items, ImageDouble{First: true, Second: true})
	}

	if len(images)%2 != 0 {
		items = append(items, ImageDouble{First: true, Second: false})
		titles = append(titles, "")
	}

	v.Set("imagesizes", imagesizes)
	v.Set("itemcount", len(items))
	v.Set("items", items)
	v.Set("titles", titles)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("periodic/periodic-02.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("periodic/periodic-02-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Periodic3(id int64, conn *models.Connection) string {
	periodicManager := models.NewPeriodicManager(conn)
	managebookManager := models.NewManagebookManager(conn)

	periodic := periodicManager.Get(id)

	managebooks := managebookManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("mc_order,mc_name,mb_order,mb_id"),
	})

	images := make([]ImageInfo, 0)

	v := make(jet.VarMap)

	for _, v := range managebooks {
		filename := fmt.Sprintf("periodicresult/%v/%v", id, v.Filename)
		images = append(images, GetImageInfo(filename))
	}

	v.Set("imagecount", len(images))
	v.Set("images", images)

	items := make([]ImageDouble, 0)

	count := int(len(images) / 2)
	for i := 0; i < count; i++ {
		items = append(items, ImageDouble{First: true, Second: true})
	}

	if len(images)%2 != 0 {
		items = append(items, ImageDouble{First: true, Second: false})
	}

	v.Set("itemcount", len(items))
	v.Set("items", items)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("periodic/periodic-03.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("periodic/periodic-03-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Periodic4(id int64, conn *models.Connection) string {
	periodicManager := models.NewPeriodicManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	periodicotherManager := models.NewPeriodicotherManager(conn)

	periodic := periodicManager.Get(id)

	aptdongs := aptdongManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("au_order,au_id"),
	})

	allPeriodicdatas := periodicdataManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "group", Value: 1, Compare: ">="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id,pd_order,pd_id"),
	})

	for i, v := range allPeriodicdatas {
		if v.Type == 3 || v.Type == 5 {
			allPeriodicdatas[i].Type = 1
		} else if v.Type == 4 || v.Type == 6 {
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

	for _, v := range allPeriodicdatas {
		if v.Type >= 200 {
			continue
		}

		periodicdatas = append(periodicdatas, v)
	}

	periodicothers := periodicotherManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("po_order,po_id"),
	})

	v := make(jet.VarMap)

	images := make([]ImageInfo, 0)

	number := 1
	for _, v := range periodicdatas {
		if v.Filename == "" {
			continue
		}

		blueprint := v.Extra["blueprint"].(models.Blueprint)

		var aptdong models.Aptdong
		for _, v2 := range aptdongs {
			if v2.Id == blueprint.Aptdong {
				aptdong = v2
				break
			}
		}

		title := ""
		if blueprint.Level == 1 {
			title = fmt.Sprintf("%v %v %v", blueprint.Name, v.Part, v.Member)
		} else {
			if aptdong.Private == 3 {
				title = fmt.Sprintf("%v %v %v", blueprint.Name, v.Part, v.Member)
			} else {
				title = fmt.Sprintf("%v %v %v %v", aptdong.Dong, blueprint.Name, v.Part, v.Member)
			}
		}

		strs := strings.Split(v.Filename, ",")

		for _, v2 := range strs {
			v2 = GetThumbnail(v2)
			info := GetImageInfo(v2)
			info.Text = title
			info.Status = v.Shape
			info.Number = fmt.Sprintf("%v", number)
			images = append(images, info)

			number++
		}
	}

	otherCategorys := []int{10, 11, 12, 13, 2, 14, 1, 3}

	for _, category := range otherCategorys {
		for _, v := range periodicothers {
			if category != v.Category {
				continue
			}

			if v.Filename == "" {
				continue
			}

			title := v.Position

			strs := strings.Split(v.Filename, ",")

			for _, v2 := range strs {
				v2 = GetThumbnail(v2)
				info := GetImageInfo(v2)
				info.Text = title

				if v.Status == "" {
					info.Status = "상태 양호"
					/*
						if v.Type == 1 {
							if v.Result == 1 {
								info.Status = "상태 양호"
							} else {
								info.Status = "상태 보통"
							}
						} else if v.Type == 2 {
							if v.Result == 1 {
								info.Status = "없음"
							} else {
								info.Status = "있음"
							}
						}
					*/
				} else {
					info.Status = v.Status
				}

				info.Number = fmt.Sprintf("%v", number)

				images = append(images, info)

				number++
			}
		}
	}

	if len(images)%6 > 0 {
		remain := len(images) % 6
		for i := 0; i < 6-remain; i++ {
			info := GetImageInfo("empty.jpg")
			images = append(images, info)
		}
	}

	v.Set("imagecount", len(images))
	v.Set("images", images)

	items := make([]PartImageInfo, 0)

	for i := 0; i < len(images)/6; i++ {
		var part PartImageInfo

		part.Image1 = images[i*6+0]
		part.Image2 = images[i*6+1]
		part.Image3 = images[i*6+2]
		part.Image4 = images[i*6+3]
		part.Image5 = images[i*6+4]
		part.Image6 = images[i*6+5]

		items = append(items, part)
	}

	v.Set("itemcount", len(items))
	v.Set("items", items)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("periodic/periodic-04.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("periodic/periodic-04-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Periodic5(id int64, conn *models.Connection) string {
	aptManager := models.NewAptManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	periodicManager := models.NewPeriodicManager(conn)
	aptperiodicManager := models.NewAptperiodicManager(conn)
	periodicimageManager := models.NewPeriodicimageManager(conn)
	periodictechnicianManager := models.NewPeriodictechnicianManager(conn)
	periodickeepManager := models.NewPeriodickeepManager(conn)

	periodic := periodicManager.Get(id)
	apt := aptManager.Get(periodic.Apt)
	if periodic.Aptname != "" {
		apt.Name = periodic.Aptname
	}

	aptperiodic := aptperiodicManager.Get(periodic.Apt)

	periodicimages := periodicimageManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "use", Value: 1, Compare: "="},
		models.Ordering("pi_order,pi_id"),
	})

	technicians := periodictechnicianManager.Find([]interface{}{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("dt_order,dt_id"),
	})

	aptdongs := aptdongManager.Find([]interface{}{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("au_order,au_id"),
	})

	periodickeep := periodickeepManager.GetByPeriodic(id)
	if periodickeep == nil {
		periodickeep = &models.Periodickeep{}
	}

	v := make(jet.VarMap)

	technicianCount := 0
	if technicians != nil {
		technicianCount = len(technicians)
	}

	v.Set("apt", apt)
	v.Set("aptperiodic", aptperiodic)
	v.Set("periodic", periodic)
	v.Set("technicianCount", technicianCount+1)
	v.Set("technicians", technicians)

	v.Set("keep1", strings.Split(periodickeep.Content1, "\n"))
	v.Set("keep2", strings.Split(periodickeep.Content2, "\n"))
	v.Set("keep3", strings.Split(periodickeep.Content3, "\n"))
	v.Set("keep4", strings.Split(periodickeep.Content4, "\n"))
	v.Set("keep5", strings.Split(periodickeep.Content5, "\n"))
	v.Set("keep6", strings.Split(periodickeep.Content6, "\n"))
	v.Set("periodickeep", periodickeep)

	temp := strings.Split(strings.ReplaceAll(periodickeep.Content2, "-", ""), ",")

	titles := []string{"건축", "구조", "기계", "전기", "토목", "통신"}
	strs := make([]string, 0)

	for _, v := range temp {
		str := strings.TrimSpace(v)

		find := false
		for _, title := range titles {
			if str == title {
				find = true
				break
			}
		}

		if find == true {
			strs = append(strs, str)
		}
	}

	keepstr := ""
	if len(strs) > 0 {
		keepstr = fmt.Sprintf("[%v]", strings.Join(strs, "/"))
	}

	v.Set("keepstr", keepstr)

	images := make([]ImageInfo, 0)
	basicimages := make([]ImageInfo, 0)

	for i := 1; i <= 51; i++ {
		basicimages = append(basicimages, GetImageInfo(fmt.Sprintf("../doc/img/%v.jpg", i)))
	}

	sign := false
	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 4 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
		sign = true

		break
	}

	if sign == false {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("sign", sign)

	license := false
	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 6 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
		license = true

		break
	}

	if license == false {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("license", license)

	cert := false
	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 7 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
		cert = true

		break
	}

	if cert == false {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("cert", cert)

	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 5 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
	}

	v.Set("imagecount", len(images))
	v.Set("images", images)
	v.Set("basicimages", basicimages)

	underground := 0
	ground := 0
	dongcount := 0
	for _, v := range aptdongs {
		if v.Private == 2 {
			continue
		}

		if v.Private == 1 {
			dongcount++
		}

		if v.Undergroundcount+v.Parkingcount > underground {
			underground = v.Undergroundcount + v.Parkingcount
		}

		if v.Groundcount > ground {
			ground = v.Groundcount
		}
	}

	floor := ""
	if underground > 0 && ground > 0 {
		floor = fmt.Sprintf("지하%v층, 지상%v층", underground, ground)
	} else if underground > 0 {
		floor = fmt.Sprintf("지하%v층", underground)
	} else {
		floor = fmt.Sprintf("지상%v층", ground)
	}

	v.Set("floor", floor)
	v.Set("dongcount", dongcount)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	view.AddGlobal("humandate", func(str string) string {
		temp := strings.Split(str, "-")

		if len(temp) != 3 {
			return str
		}

		return fmt.Sprintf("%04d년 %02d월 %02d일", global.Atoi(temp[0]), global.Atoi(temp[1]), global.Atoi(temp[2]))
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("periodic/periodic-05.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("periodic/periodic-05-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

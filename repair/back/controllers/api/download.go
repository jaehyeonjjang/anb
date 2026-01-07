package api

import (
	"errors"
	"fmt"
	"log"
	_ "net/url"
	"os"
	"repair/controllers"
	"repair/controllers/api/document/detail"
	"repair/controllers/api/document/periodic"
	"repair/estimate"
	"repair/global"
	"repair/global/config"
	"repair/models"
	re "repair/models/repair"
	"sort"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/xuri/excelize/v2"
)

type SummaryData struct {
}

type Summary struct {
	Category  int64
	Price     int64
	Saveprice int64
	Method    []string
	Cycle     []int
	Percent   []int
}

type TopSummary struct {
	Category        int64
	Price           int64
	Saveprice       int64
	SavepriceParcel int64
}

type TotalreportItem struct {
	Type int

	Index   int
	Report  models.Totalreport
	History models.History
	Span    int
}

type TotalreportCategory struct {
	Topcategory int64
	Subcategory int64

	Data []TotalreportItem
}

type DownloadController struct {
	controllers.Controller
}

func GetCell(col string, row int) string {
	return fmt.Sprintf("%v%v", col, row)
}

func GetPlanyears(value int) int {
	if value == 0 {
		return 49
	}

	return value - 1
}

func CalculatePriceRate(directInt int64, laborInt int, costInt int, rate float64, parcelrate float64) int64 {
	direct := float64(directInt)
	labor := float64(laborInt)
	cost := float64(costInt)

	k := labor * 7.9 / 100
	l := (direct + labor) * 5.5 / 100
	m := labor * 3.75 / 100
	n := labor * 0.79 / 100
	o := labor * 3.23 / 100
	p := labor * 4.5 / 100

	q := o * 8.51 / 100
	r := labor * 2.3 / 100
	s := (direct + labor) * 3.09 / 100
	t := (direct + labor + cost) * 0.3 / 100
	u := (direct + labor + cost) * 0.07 / 100
	v := (direct + labor + cost) * 6 / 100
	w := (labor + cost + v) * 15 / 100

	x := direct + labor + cost + k + l + m + n + o + p + q + r + s + t + u + v + w
	y := x * 10 / 100

	var ret float64 = x + y
	if rate != 0.0 && rate != 100.0 {
		ret *= rate / 100.0
	}

	if parcelrate != 0.0 && parcelrate != 100.0 {
		ret *= parcelrate / 100.0
	}

	return int64(ret)
}

func CalculateRepair(directInt int64, laborInt int, costInt int, rate float64, parcelrate float64, count int, percent float64) int64 {
	retI := CalculatePriceRate(directInt, laborInt, costInt, rate, parcelrate)
	ret := float64(retI)
	ret *= float64(count)

	if percent != 0.0 {
		ret *= percent / 100.0
	}

	return int64(ret)
}

func (c *DownloadController) File(id int64) {
	conn := c.NewConnection()

	manager := models.NewFileManager(conn)
	item := manager.Get(id)

	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, item.Filename)
	c.Download(fullFilename, item.Originalfilename)
}

func getDurationMonth(endyear int, endmonth int, startyear int, startmonth int) int {
	if startyear > endyear {
		return 0
	}

	if startyear == endyear {
		if startmonth > endmonth {
			return 0
		}

		return endmonth - startmonth + 1
	}

	total := 12 - startmonth + 1
	total += (endyear - startyear - 1) * 12
	total += endmonth

	return total
}

func (c *DownloadController) Report(id int64) {
	conn := c.NewConnection()

	aptManager := models.NewAptManager(conn)
	repairManager := models.NewRepairManager(conn)
	areaManager := models.NewAreaManager(conn)
	dongManager := models.NewDongManager(conn)
	categoryManager := models.NewCategoryManager(conn)
	historyManager := models.NewHistoryManager(conn)
	breakdownManager := models.NewBreakdownManager(conn)
	totalreportManager := models.NewTotalreportManager(conn)
	totalyearreportManager := models.NewTotalyearreportManager(conn)
	yearreportManager := models.NewYearreportManager(conn)
	standardManager := models.NewStandardManager(conn)
	reviewcontentManager := models.NewReviewcontentManager(conn)
	//reviewdateManager := models.NewReviewdateManager(conn)
	reviewManager := models.NewReviewManager(conn)
	outlineManager := models.NewOutlineManager(conn)
	outlineplanManager := models.NewOutlineplanManager(conn)
	savingManager := models.NewSavingManager(conn)
	approvalManager := models.NewApprovalManager(conn)

	repair := repairManager.Get(id)
	apt := aptManager.Get(repair.Apt)

	repairs := repairManager.Find([]any{models.Where{Column: "apt", Value: repair.Apt, Compare: "="}, models.Ordering("r_reportdate,r_id")})

	categorys := categoryManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id")})
	standards := standardManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("s_order,s_id")})

	if repair.Provision == 2 {
		for i, v := range categorys {
			categorys[i].Name = ReplaceWord(2, v.Name)
		}

		for i, v := range standards {
			standards[i].Name = ReplaceWord(2, v.Name)
		}
	}

	areas := areaManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("ar_order,ar_id")})
	dongs := dongManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Where{Column: "basic", Value: 1, Compare: "="}, models.Ordering("d_order,d_id")})
	alldongs := dongManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("d_order,d_id")})
	historys := historyManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("h_year,h_month,h_id")})
	breakdowns := breakdownManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_dong,b_standard,b_method,b_id")})
	totalreports := totalreportManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_standard,b_method")})
	totalyearreports := totalyearreportManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_topcategory,b_subcategory,b_category,b_standard,b_method")})
	yearreports := yearreportManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("b_duedate,b_topcategory,b_subcategory,b_category,b_standard,b_method")})
	//reviewdates := reviewdateManager.Find([]interface{}{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("rd_id")})
	reviewcontents := reviewcontentManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("rc_id")})
	reviewsList := reviewManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("c_order,c_id,re_method,re_id")})

	repairParcelrate := 0.0

	reviews := MakeReview(id)

	for _, v := range reviewsList {
		find := false
		pos := 0
		for i, v2 := range reviews {
			if v.Topcategory == v2.Topcategory &&
				v.Subcategory == v2.Subcategory &&
				v.Category == v2.Category &&
				v.Standard == v2.Standard &&
				v.Method == v2.Method {
				find = true
				pos = i
				break
			}
		}

		if find == true {
			reviews[pos] = v
		} else {
			reviews = append([]models.Review{v}, reviews...)
		}
	}

	//reviews := append(reviewsList, MakeReview(id)...)
	sort.Slice(reviews, func(i, j int) bool {
		return reviews[i].Method < reviews[j].Method
	})

	outlines := outlineManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("o_id")})
	outlineplans := outlineplanManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("op_id")})
	savings := savingManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("sa_year,sa_id")})
	approvals := approvalManager.Find([]any{models.Where{Column: "apt", Value: id, Compare: "="}, models.Ordering("ap_order,ap_id")})

	categoryMap := make(map[int64]models.Category)
	standardMap := make(map[int64]models.Standard)
	dongMap := make(map[int64]models.Dong)

	/*
		level2Title := []string{"", "가", "나", "다", "라", "마", "바", "사", "아", "자", "차", "카", "타", "파", "하", "거", "너", "더", "러", "머", "버", "서", "어", "저", "처", "커", "터", "퍼", "허"}
		level1 := 1

		for i, v := range *categorys {
			if v.Level == 1 {
				(*categorys)[i].Name = fmt.Sprintf("%v. %v", level1, v.Name)
				level1++

				level2 := 1

				for i2, v2 := range *categorys {
					if v2.Level == 2 && v2.Parent == v.Id {
						(*categorys)[i2].Name = fmt.Sprintf("%v. %v", level2Title[level2], v2.Name)
						level2++

						level3 := 1

						for i3, v3 := range *categorys {
							if v3.Level == 3 && v3.Parent == v2.Id {
								(*categorys)[i3].Name = fmt.Sprintf("(%v) %v", level3, v3.Name)
								level3++
							}
						}
					}
				}
			}
		}
	*/

	for _, v := range categorys {
		categoryMap[v.Id] = v
	}

	for _, v := range standards {
		standardMap[v.Id] = v
	}

	for _, v := range alldongs {
		dongMap[v.Id] = v
	}

	excelFilename := "./doc/repair/basic.xlsx"
	if repair.Provision == 2 {
		excelFilename = "./doc/repair/basic2.xlsx"
	}

	f, err := excelize.OpenFile(excelFilename)
	if err != nil {
		log.Println(err)
		return
	}

	temps := strings.Split(repair.Reportdate, "-")
	reportyear := ""
	reportmonth := ""

	reportyear = temps[0]

	if len(temps) > 1 {
		reportmonth = temps[1]
	}

	sheet := "표지.목차.개요(2)"

	tt := ReplaceWord(repair.Provision, "장기수선계획서(정기조정)")

	if repair.Periodtype == re.PeriodtypeSometime {
		tt = ReplaceWord(repair.Provision, "장기수선계획서(수시조정)")
	}

	// if repair.Remark == "정기" {
	// 	tt = ReplaceWord(repair.Provision, "장기수선계획서(정기조정)")
	// } else if repair.Remark == "수시" {
	// 	tt = ReplaceWord(repair.Provision, "장기수선계획서(수시조정)")
	// } else if strings.Contains(repair.Remark, "정기") {
	// 	tt = ReplaceWord(repair.Provision, "장기수선계획서(정기조정)")
	// } else if strings.Contains(repair.Remark, "수시") {
	// 	tt = ReplaceWord(repair.Provision, "장기수선계획서(수시조정)")
	// }

	f.SetCellStr(sheet, "A5", tt)
	f.SetCellStr(sheet, "A65", tt)

	f.SetCellStr(sheet, "A9", apt.Name)
	f.SetCellStr(sheet, "A93", fmt.Sprintf("%v 관리사무소", apt.Name))
	f.SetCellStr(sheet, "E143", apt.Name)
	f.SetCellStr(sheet, "E144", fmt.Sprintf("%v세대 (동수 : %v개동)", apt.Familycount, apt.Flatcount))
	f.SetCellStr(sheet, "E145", apt.Address)
	f.SetCellStr(sheet, "E146", repair.Complex1)
	f.SetCellStr(sheet, "E147", repair.Complex2)
	f.SetCellStr(sheet, "E148", fmt.Sprintf("%v년 %v월 %v일", repair.Completionyear, repair.Completionmonth, repair.Completionday))

	f.SetCellStr(sheet, "E151", apt.Name)
	f.SetCellStr(sheet, "E152", apt.Address)
	f.SetCellStr(sheet, "E153", repair.Info1)
	f.SetCellStr(sheet, "E154", repair.Info2)
	f.SetCellStr(sheet, "E155", repair.Info3)
	f.SetCellStr(sheet, "E156", repair.Info4)
	f.SetCellStr(sheet, "E157", repair.Info5)
	f.SetCellStr(sheet, "E158", repair.Info6)
	f.SetCellStr(sheet, "E159", repair.Info7)
	f.SetCellStr(sheet, "E160", repair.Info8)
	f.SetCellStr(sheet, "E161", repair.Info9)
	f.SetCellStr(sheet, "E162", repair.Info10)
	f.SetCellStr(sheet, "E163", repair.Info11)

	f.SetCellStr(sheet, "K169", repair.Structure1)
	f.SetCellStr(sheet, "K170", repair.Structure2)
	f.SetCellStr(sheet, "K171", repair.Structure3)
	f.SetCellStr(sheet, "K172", repair.Structure4)
	f.SetCellStr(sheet, "K173", repair.Structure5)
	f.SetCellStr(sheet, "K174", repair.Structure6)
	f.SetCellStr(sheet, "K175", repair.Structure7)
	f.SetCellStr(sheet, "K176", repair.Structure8)
	f.SetCellStr(sheet, "K177", repair.Structure9)
	f.SetCellStr(sheet, "K178", repair.Structure10)
	f.SetCellStr(sheet, "K179", repair.Structure11)
	f.SetCellStr(sheet, "K180", repair.Structure12)
	f.SetCellStr(sheet, "K181", repair.Structure13)
	f.SetCellStr(sheet, "K182", repair.Structure14)

	reportdateStr := fmt.Sprintf("%v년 %v월", reportyear, reportmonth)
	f.SetCellStr(sheet, "A13", reportdateStr)
	f.SetCellStr(sheet, "A46", reportdateStr)
	f.SetCellStr(sheet, "A89", reportdateStr)

	if repair.Provision == 2 {
		f.SetCellStr(sheet, "B127", "04. 용어정리")
		f.SetCellStr(sheet, "B128", "")
		f.SetCellStr(sheet, "B129", "")
		f.SetCellStr(sheet, "B130", "")
	}

	{
		cnt := 0
		for _, v := range repairs {
			if v.Reportdate == "" {
				continue
			}

			if v.Reportdate > repair.Reportdate {
				continue
			}

			cnt++
		}

		pos := 82
		step := 1

		f.SetCellStr(sheet, "E82", fmt.Sprintf("○ 최초 수립 : %v년 %v월", repair.Completionyear, repair.Completionmonth))

		for _, v := range repairs {
			if v.Reportdate == "" {
				continue
			}

			if v.Reportdate > repair.Reportdate {
				continue
			}

			strs := strings.Split(v.Reportdate, "-")
			year := ""
			month := ""
			if len(strs) >= 2 {
				year = strs[0]
				month = strs[1]
			}

			if step == cnt {
				f.SetCellStr(sheet, GetCell("E", pos+step), fmt.Sprintf("○ 금회 조정 : %v년 %v월", year, global.Atoi(month)))
			} else {
				if v.Periodtype == re.PeriodtypeSometime {
					f.SetCellStr(sheet, GetCell("E", pos+step), fmt.Sprintf("○ 수시 조정 : %v년 %v월", year, global.Atoi(month)))
				} else {
					f.SetCellStr(sheet, GetCell("E", pos+step), fmt.Sprintf("○ 정기 조정 : %v년 %v월", year, global.Atoi(month)))
				}
			}

			step++
		}
	}

	var areaTotal float64 = 0.0
	// area
	if len(areas) == 0 {
		f.RemoveRow(sheet, 186)
	} else {
		rows := len(areas)

		for i := 0; i < rows-1; i++ {
			f.DuplicateRow(sheet, 186)
		}

		for i, v := range areas {
			f.SetCellInt(sheet, GetCell("A", 186+i), i+1)
			f.SetCellStr(sheet, GetCell("C", 186+i), v.Name)
			f.SetCellInt(sheet, GetCell("E", 186+i), v.Familycount)
			f.SetCellStr(sheet, GetCell("G", 186+i), humanize.CommafWithDigits(float64(v.Size), 2))

			size := float64(v.Size) * float64(v.Familycount)
			f.SetCellStr(sheet, GetCell("J", 186+i), humanize.CommafWithDigits(size, 3))

			areaTotal += size
		}

		if repair.Parcelrate != 0.0 && repair.Parcelrate != 100.0 {
			areaTotal *= float64(repair.Parcelrate) / 100.0
		}

		f.SetCellStr(sheet, GetCell("J", 186+rows), humanize.CommafWithDigits(areaTotal, 3))
	}

	// dong
	if len(dongs) == 0 {
		f.RemoveRow(sheet, 191+len(areas)-1)
	} else {
		start := 191 + len(areas) - 1
		rows := len(dongs)

		for i := 0; i < rows-1; i++ {
			f.DuplicateRow(sheet, start)
		}

		for i, v := range dongs {
			f.SetCellInt(sheet, GetCell("A", start+i), i+1)
			f.SetCellStr(sheet, GetCell("C", start+i), v.Name)
			f.SetCellInt(sheet, GetCell("E", start+i), v.Ground)
			f.SetCellInt(sheet, GetCell("G", start+i), v.Underground)
			f.SetCellInt(sheet, GetCell("J", start+i), v.Familycount)
			f.SetCellStr(sheet, GetCell("M", start+i), v.Remark)
		}
	}

	{
		cols := []string{"C", "E", "G", "I", "K", "M", "C", "E", "G", "I", "K", "M"}
		rows := []int{73, 73, 73, 73, 73, 73, 77, 77, 77, 77, 77, 77}

		for i, v := range approvals {
			if i >= 12 {
				continue
			}
			f.SetCellStr(sheet, GetCell(cols[i], rows[i]), v.Duty)
			f.SetCellStr(sheet, GetCell(cols[i], rows[i]+1), v.Name)
		}
	}

	// 장기수선충당금 사용현황

	sheet = ReplaceWord(repair.Provision, "장기수선충당금 사용현황")
	if len(historys) == 0 {
		f.RemoveRow(sheet, 2)
	} else {
		start := 3
		rows := len(historys)

		for i := 0; i < rows-1; i++ {
			f.DuplicateRow(sheet, start)
		}

		var total int64 = 0

		for i, v := range historys {
			topcategory := ""
			if val, ok := categoryMap[v.Topcategory]; ok {
				topcategory = val.Name
			}

			subcategory := ""
			if val, ok := categoryMap[v.Subcategory]; ok {
				subcategory = val.Name
			}

			category := ""
			if val, ok := categoryMap[v.Category]; ok {
				category = val.Name
			}

			f.SetCellStr(sheet, GetCell("A", start+i), fmt.Sprintf("%v년 %v월", v.Year, v.Month))
			f.SetCellStr(sheet, GetCell("B", start+i), topcategory)
			f.SetCellStr(sheet, GetCell("C", start+i), subcategory)
			f.SetCellStr(sheet, GetCell("D", start+i), category)
			f.SetCellStr(sheet, GetCell("E", start+i), v.Content)
			f.SetCellDefault(sheet, GetCell("F", start+i), humanize.Comma(int64(v.Price)))

			total += int64(v.Price)
		}

		f.SetCellDefault(sheet, GetCell("F", start+rows), humanize.Comma(total))
	}

	// 공사종별 수선계획금액 집계표
	sheet = "공사종별 수선계획금액 집계표"

	startTop := 3
	topSummary := make(map[int64]TopSummary)
	summary := make(map[int64]Summary)

	for _, v := range breakdowns {
		standard := v.Extra["standard"].(models.Standard)
		category := v.Extra["category"].(models.Category)

		price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repairParcelrate), v.Count, float64(category.Percent))
		saveprice := int64(float64(price) / float64(category.Cycle))
		priceParcel := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repair.Parcelrate), v.Count, float64(category.Percent))
		savepriceParcel := int64(float64(priceParcel) / float64(category.Cycle))

		if value, ok := topSummary[v.Topcategory]; ok {
			value.Price += price
			value.Saveprice += saveprice
			value.SavepriceParcel += savepriceParcel
			topSummary[v.Topcategory] = value
		} else {
			var newValue TopSummary
			newValue.Price = price
			newValue.Saveprice = saveprice
			newValue.SavepriceParcel += savepriceParcel
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
	var topTotalSavepriceParcel int64 = 0

	subtopPos := []int{13, 18, 24, 36, 42, 47, 50}

	maxPos := 7
	convert := categoryManager.GetByAptName(id, "7. 피난시설")
	if convert == nil {
		maxPos = 6
	}

	for _, v := range categorys {
		if v.Level != 1 {
			continue
		}

		var price int64 = 0
		var saveprice int64 = 0
		var savepriceParcel int64 = 0

		if value, ok := topSummary[v.Id]; ok {
			price = value.Price
			saveprice = value.Saveprice
			savepriceParcel = value.SavepriceParcel
		}

		topTotalPrice += price
		topTotalSaveprice += saveprice
		topTotalSavepriceParcel += savepriceParcel

		f.SetCellDefault(sheet, GetCell("B", startTop+pos), fmt.Sprintf("%v", price))
		f.SetCellDefault(sheet, GetCell("D", startTop+pos), fmt.Sprintf("%v", saveprice))
		f.SetCellDefault(sheet, GetCell("F", startTop+pos), fmt.Sprintf("%v", saveprice*int64(GetPlanyears(repair.Planyears))))

		f.SetCellDefault(sheet, GetCell("E", subtopPos[pos]), fmt.Sprintf("%v", price))
		f.SetCellDefault(sheet, GetCell("F", subtopPos[pos]), fmt.Sprintf("%v", saveprice))

		pos2 := 0

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

			target := 0

			if typeid == 1 {
				target = subtopPos[pos]
			} else {
				target = subtopPos[pos] + 1 + pos2
			}

			if typeid == 2 {
				f.SetCellDefault(sheet, GetCell("B", target), method)
			}

			f.SetCellDefault(sheet, GetCell("C", target), cycle)
			f.SetCellDefault(sheet, GetCell("D", target), percent)

			f.SetCellDefault(sheet, GetCell("E", target), humanize.Comma(price))
			f.SetCellDefault(sheet, GetCell("F", target), humanize.Comma(saveprice))

			if typeid == 2 {
				pos2++
			}
		}

		pos++

		if pos >= maxPos {
			break
		}
	}

	if maxPos == 6 {
		pos++
	}

	f.SetCellDefault(sheet, GetCell("B", startTop+pos), fmt.Sprintf("%v", topTotalPrice))
	f.SetCellDefault(sheet, GetCell("D", startTop+pos), fmt.Sprintf("%v", topTotalSaveprice))
	f.SetCellDefault(sheet, GetCell("F", startTop+pos), fmt.Sprintf("%v", topTotalSaveprice*int64(GetPlanyears(repair.Planyears))))

	f.SetCellDefault(sheet, GetCell("E", 53), fmt.Sprintf("%v", topTotalPrice))
	f.SetCellDefault(sheet, GetCell("F", 53), fmt.Sprintf("%v", topTotalSaveprice))

	if maxPos == 6 {
		f.RemoveRow(sheet, 52)
		f.RemoveRow(sheet, 51)
		f.RemoveRow(sheet, 50)
		f.RemoveRow(sheet, 9)
	}

	// 항목총괄자료
	sheet = "항목총괄자료"

	totalreportCategorys := make([]TotalreportCategory, 0)

	for _, v := range categorys {
		for _, v2 := range totalreports {
			if v.Id == int64(v2.Subcategory) {
				flag := false
				for _, v3 := range totalreportCategorys {
					if v3.Subcategory == int64(v2.Subcategory) {
						flag = true
						break
					}
				}

				if flag == false {
					var item TotalreportCategory
					item.Topcategory = int64(v2.Topcategory)
					item.Subcategory = int64(v2.Subcategory)
					item.Data = make([]TotalreportItem, 0)

					totalreportCategorys = append(totalreportCategorys, item)
				}
			}
		}
	}

	for i, v := range totalreportCategorys {
		var lastCategory int64 = -1
		reportIndex := 1

		for _, v2 := range totalreports {
			if v.Subcategory == int64(v2.Subcategory) {
				if lastCategory != -1 && int64(v2.Category) != lastCategory {
					// history 넣어야 함

					historyPos := 0
					historyLast := -1
					for _, v3 := range historys {
						if lastCategory == v3.Category {
							if historyPos == 0 {
								var item TotalreportItem
								item.Type = 3
								totalreportCategorys[i].Data = append(totalreportCategorys[i].Data, item)
							}

							var item TotalreportItem
							item.Type = 2
							item.History = v3
							totalreportCategorys[i].Data = append(totalreportCategorys[i].Data, item)

							historyLast = len(totalreportCategorys[i].Data) - 1
							historyPos++
						}
					}

					if historyLast != -1 {
						totalreportCategorys[i].Data[historyLast].Span = historyPos
					}
				}

				var item TotalreportItem
				item.Type = 1
				item.Report = v2
				item.Index = reportIndex
				totalreportCategorys[i].Data = append(totalreportCategorys[i].Data, item)

				lastCategory = int64(v2.Category)
				reportIndex++
			}
		}

		if lastCategory != -1 {
			// history 넣어야 함
			historyPos := 0
			historyLast := -1
			for _, v3 := range historys {
				if lastCategory == v3.Category {
					if historyPos == 0 {
						var item TotalreportItem
						item.Type = 3
						totalreportCategorys[i].Data = append(totalreportCategorys[i].Data, item)
					}

					var item TotalreportItem
					item.Type = 2
					item.History = v3
					totalreportCategorys[i].Data = append(totalreportCategorys[i].Data, item)

					historyLast = len(totalreportCategorys[i].Data) - 1
					historyPos++
				}
			}

			if historyLast != -1 {
				totalreportCategorys[i].Data[historyLast].Span = historyPos
			}
		}
	}

	pos = 13

	for _, v := range totalreportCategorys {
		f.DuplicateRowTo(sheet, 3, pos)
		f.DuplicateRowTo(sheet, 4, pos+1)

		title := fmt.Sprintf("%v > %v", categoryMap[v.Topcategory].Name, categoryMap[v.Subcategory].Name)
		f.SetCellDefault(sheet, GetCell("A", pos+1), title)
		f.DuplicateRowTo(sheet, 5, pos+2)
		f.DuplicateRowTo(sheet, 6, pos+3)

		f.MergeCell(sheet, GetCell("A", pos+2), GetCell("A", pos+3))
		f.MergeCell(sheet, GetCell("B", pos+2), GetCell("B", pos+3))
		f.MergeCell(sheet, GetCell("C", pos+2), GetCell("C", pos+3))
		f.MergeCell(sheet, GetCell("D", pos+2), GetCell("D", pos+3))
		f.MergeCell(sheet, GetCell("G", pos+2), GetCell("G", pos+3))
		f.MergeCell(sheet, GetCell("H", pos+2), GetCell("H", pos+3))
		f.MergeCell(sheet, GetCell("I", pos+2), GetCell("I", pos+3))
		f.MergeCell(sheet, GetCell("J", pos+2), GetCell("J", pos+3))
		f.MergeCell(sheet, GetCell("K", pos+2), GetCell("K", pos+3))
		f.MergeCell(sheet, GetCell("L", pos+2), GetCell("L", pos+3))
		f.MergeCell(sheet, GetCell("M", pos+2), GetCell("M", pos+3))
		f.MergeCell(sheet, GetCell("N", pos+2), GetCell("N", pos+3))

		pos += 4

		var partTotal int64 = 0
		var partSave int64 = 0

		var allTotal int64 = 0
		var allSave int64 = 0

		for _, v2 := range v.Data {
			switch v2.Type {
			case 1:
				standard := standardMap[int64(v2.Report.Standard)]
				method := categoryMap[int64(v2.Report.Method)]
				f.DuplicateRowTo(sheet, 7, pos)
				f.SetCellInt(sheet, GetCell("A", pos), v2.Index)
				f.SetCellDefault(sheet, GetCell("B", pos), categoryMap[int64(v2.Report.Category)].Name)
				f.SetCellDefault(sheet, GetCell("C", pos), standardMap[int64(v2.Report.Standard)].Name)
				f.SetCellDefault(sheet, GetCell("D", pos), method.Name)
				f.SetCellInt(sheet, GetCell("E", pos), method.Cycle)
				f.SetCellInt(sheet, GetCell("F", pos), method.Percent)
				f.SetCellDefault(sheet, GetCell("G", pos), standard.Unit)
				f.SetCellDefault(sheet, GetCell("H", pos), humanize.Comma(int64(v2.Report.Count)))

				priceOne := CalculatePriceRate(standard.Direct, standard.Labor, standard.Cost, float64(v2.Report.Rate), float64(repairParcelrate))
				f.SetCellDefault(sheet, GetCell("I", pos), humanize.Comma(priceOne))

				price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v2.Report.Rate), float64(repairParcelrate), v2.Report.Count, float64(method.Percent))
				f.SetCellDefault(sheet, GetCell("J", pos), humanize.Comma(price))

				saveprice := int64(float64(price) / float64(method.Cycle))

				f.SetCellDefault(sheet, GetCell("K", pos), humanize.Comma(saveprice))
				f.SetCellInt(sheet, GetCell("L", pos), v2.Report.Lastdate)
				f.SetCellInt(sheet, GetCell("M", pos), v2.Report.Duedate)

				if method.Percent == 100 {
					allTotal += price
					allSave += saveprice
				} else {
					partTotal += price
					partSave += saveprice
				}
			case 2:
				f.DuplicateRowTo(sheet, 9, pos)
				f.SetCellDefault(sheet, GetCell("C", pos), fmt.Sprintf("%v년 %v월", v2.History.Year, v2.History.Month))
				f.SetCellDefault(sheet, GetCell("D", pos), v2.History.Content)
			default:
				f.DuplicateRowTo(sheet, 8, pos)
			}

			if v2.Span > 0 {
				f.MergeCell(sheet, GetCell("A", pos-v2.Span), GetCell("B", pos))
			}
			pos++
		}

		f.DuplicateRowTo(sheet, 10, pos)
		f.DuplicateRowTo(sheet, 11, pos+1)
		f.DuplicateRowTo(sheet, 12, pos+2)

		f.SetCellDefault(sheet, GetCell("J", pos), humanize.Comma(partTotal))
		f.SetCellDefault(sheet, GetCell("K", pos), humanize.Comma(partSave))

		f.SetCellDefault(sheet, GetCell("J", pos+1), humanize.Comma(allTotal))
		f.SetCellDefault(sheet, GetCell("K", pos+1), humanize.Comma(allSave))

		f.SetCellDefault(sheet, GetCell("J", pos+2), humanize.Comma(partTotal+allTotal))
		f.SetCellDefault(sheet, GetCell("K", pos+2), humanize.Comma(partSave+allSave))

		pos += 3

	}

	for i := 12; i >= 3; i-- {
		f.RemoveRow(sheet, i)
	}

	// 총론2
	sheet = "총론2"

	year := global.Atoi(reportyear)

	years := make([]int, 0)

	currentYear := time.Now().Year()
	for _, v := range yearreports {
		/*
			if v.Duedate < year {
				continue
			}
		*/
		if v.Duedate <= 0 {
			continue
		}

		if v.Duedate >= currentYear+100 {
			continue
		}

		years = append(years, v.Duedate)
	}

	years = global.Unique(years)

	if repair.Reportdate != "" {
		strs := strings.Split(repair.Reportdate, "-")
		year = global.Atoi(strs[0])
	}
	pos = 7
	no := 1
	for i := year; i < year+4; i++ {
		cnt := 0
		for _, v := range totalyearreports {
			if v.Duedate == i {
				cnt++
			}
		}

		if cnt == 0 {
			continue
		}

		if pos != 7 {
			f.DuplicateRowTo(sheet, 8, pos+1)
		}
		f.DuplicateRowTo(sheet, 3, pos+2)
		f.DuplicateRowTo(sheet, 4, pos+3)
		f.DuplicateRowTo(sheet, 5, pos+4)

		f.MergeCell(sheet, GetCell("A", pos+3), GetCell("A", pos+4))
		f.MergeCell(sheet, GetCell("B", pos+3), GetCell("B", pos+4))
		f.MergeCell(sheet, GetCell("C", pos+3), GetCell("C", pos+4))
		f.MergeCell(sheet, GetCell("D", pos+3), GetCell("D", pos+4))
		f.MergeCell(sheet, GetCell("E", pos+3), GetCell("E", pos+4))
		f.MergeCell(sheet, GetCell("F", pos+3), GetCell("F", pos+4))
		f.MergeCell(sheet, GetCell("G", pos+3), GetCell("G", pos+4))

		f.SetCellDefault(sheet, GetCell("A", pos+2), fmt.Sprintf("(%v) 수립대상시설 %v년", no, i))

		switch cnt {
		case 0:
			f.DuplicateRowTo(sheet, 8, pos+5)
		case 1:
			f.DuplicateRowTo(sheet, 7, pos+5)
			if i == year+3 {
				f.DuplicateRowTo(sheet, 8, pos+6)
			}
		default:
			for j := 0; j < cnt-1; j++ {
				f.DuplicateRowTo(sheet, 6, pos+5+j)
			}

			f.DuplicateRowTo(sheet, 7, pos+5+cnt-1)
			if i == year+3 {
				f.DuplicateRowTo(sheet, 8, pos+6+cnt-1)
			}
		}

		if cnt > 0 {
			cnt = 0
			for _, v := range totalyearreports {
				if v.Duedate == i {
					f.SetCellInt(sheet, GetCell("A", pos+5+cnt), cnt+1)
					f.SetCellDefault(sheet, GetCell("B", pos+5+cnt), categoryMap[int64(v.Topcategory)].Name)
					f.SetCellDefault(sheet, GetCell("C", pos+5+cnt), categoryMap[int64(v.Subcategory)].Name)
					f.SetCellDefault(sheet, GetCell("D", pos+5+cnt), categoryMap[int64(v.Category)].Name)
					f.SetCellDefault(sheet, GetCell("E", pos+5+cnt), standardMap[int64(v.Standard)].Name)
					f.SetCellDefault(sheet, GetCell("F", pos+5+cnt), categoryMap[int64(v.Method)].Name)

					cnt++
				}
			}
		}

		pos += 4 + cnt
		no++
	}

	for i := 8; i >= 2; i-- {
		f.RemoveRow(sheet, i)
	}

	// 항목세부자료
	sheet = "항목세부자료"

	breakdownCategorys := make([]int64, 0)
	for _, v := range breakdowns {
		flag := false
		for _, v2 := range breakdownCategorys {
			if v.Category == v2 {
				flag = true
				break
			}
		}

		if flag == false {
			breakdownCategorys = append(breakdownCategorys, v.Category)
		}
	}

	pos = 10

	for _, v := range breakdownCategorys {
		subcategory := categoryMap[v].Parent
		topcategory := categoryMap[subcategory].Parent

		title := fmt.Sprintf("%v > %v > %v", categoryMap[topcategory].Name, categoryMap[subcategory].Name, categoryMap[v].Name)

		f.DuplicateRowTo(sheet, 3, pos)
		f.DuplicateRowTo(sheet, 4, pos+1)
		f.SetCellDefault(sheet, GetCell("A", pos+1), title)
		f.DuplicateRowTo(sheet, 5, pos+2)
		pos += 3

		var partTotal int64 = 0
		var partSave int64 = 0

		var allTotal int64 = 0
		var allSave int64 = 0

		for _, v2 := range breakdowns {
			if v == v2.Category {
				standard := standardMap[v2.Standard]
				method := categoryMap[v2.Method]
				f.DuplicateRowTo(sheet, 6, pos)

				dongName := dongMap[v2.Dong].Name
				if v2.Elevator > 0 {
					dongName = fmt.Sprintf("%v %v호기", dongMap[v2.Dong].Name, v2.Elevator)
				}
				f.SetCellDefault(sheet, GetCell("A", pos), dongName)

				f.SetCellDefault(sheet, GetCell("B", pos), standardMap[v2.Standard].Name)
				f.SetCellDefault(sheet, GetCell("C", pos), method.Name)
				f.SetCellInt(sheet, GetCell("D", pos), method.Cycle)
				f.SetCellInt(sheet, GetCell("E", pos), method.Percent)
				f.SetCellDefault(sheet, GetCell("F", pos), standard.Unit)
				f.SetCellDefault(sheet, GetCell("G", pos), humanize.Comma(int64(v2.Count)))

				priceOne := CalculatePriceRate(standard.Direct, standard.Labor, standard.Cost, float64(v2.Rate), float64(repairParcelrate))
				price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v2.Rate), float64(repairParcelrate), v2.Count, float64(method.Percent))
				f.SetCellDefault(sheet, GetCell("H", pos), humanize.Comma(priceOne))

				f.SetCellDefault(sheet, GetCell("I", pos), humanize.Comma(price))

				saveprice := int64(float64(price) / float64(method.Cycle))

				f.SetCellDefault(sheet, GetCell("J", pos), humanize.Comma(saveprice))
				f.SetCellInt(sheet, GetCell("K", pos), v2.Lastdate)
				f.SetCellInt(sheet, GetCell("L", pos), v2.Duedate)

				f.SetCellStr(sheet, GetCell("M", pos), v2.Remark)

				if method.Percent == 100 {
					allTotal += price
					allSave += saveprice
				} else {
					partTotal += price
					partSave += saveprice
				}

				pos++
			}

		}

		f.DuplicateRowTo(sheet, 7, pos)
		f.DuplicateRowTo(sheet, 8, pos+1)
		f.DuplicateRowTo(sheet, 9, pos+2)

		f.SetCellDefault(sheet, GetCell("I", pos), humanize.Comma(partTotal))
		f.SetCellDefault(sheet, GetCell("J", pos), humanize.Comma(partSave))

		f.SetCellDefault(sheet, GetCell("I", pos+1), humanize.Comma(allTotal))
		f.SetCellDefault(sheet, GetCell("J", pos+1), humanize.Comma(allSave))

		f.SetCellDefault(sheet, GetCell("I", pos+2), humanize.Comma(partTotal+allTotal))
		f.SetCellDefault(sheet, GetCell("J", pos+2), humanize.Comma(partSave+allSave))

		pos += 3
	}

	for i := 9; i >= 3; i-- {
		f.RemoveRow(sheet, i)
	}

	// 연도별공사예정현황
	sheet = "연도별공사예정현황"

	pos = 10
	no = 1

	yearValue := year

	if len(years) > 0 {
		startyear := years[0]
		lastyear := years[len(years)-1]

		for i := startyear; i <= lastyear; i++ {
			year := 0
			for _, v := range years {
				if v == i {
					year = i
					break
				}
			}

			if year == 0 {
				if i > yearValue+2 {
					continue
				}

				f.DuplicateRowTo(sheet, 2, pos)
				f.DuplicateRowTo(sheet, 3, pos+1)
				f.DuplicateRowTo(sheet, 4, pos+2)
				f.DuplicateRowTo(sheet, 5, pos+3)

				f.MergeCell(sheet, GetCell("A", pos+2), GetCell("A", pos+3))
				f.MergeCell(sheet, GetCell("B", pos+2), GetCell("B", pos+3))
				f.MergeCell(sheet, GetCell("C", pos+2), GetCell("C", pos+3))
				f.MergeCell(sheet, GetCell("D", pos+2), GetCell("D", pos+3))
				f.MergeCell(sheet, GetCell("E", pos+2), GetCell("E", pos+3))
				f.MergeCell(sheet, GetCell("F", pos+2), GetCell("F", pos+3))
				f.MergeCell(sheet, GetCell("I", pos+2), GetCell("I", pos+3))
				f.MergeCell(sheet, GetCell("J", pos+2), GetCell("J", pos+3))
				f.MergeCell(sheet, GetCell("K", pos+2), GetCell("K", pos+3))
				f.MergeCell(sheet, GetCell("L", pos+2), GetCell("L", pos+3))
				f.MergeCell(sheet, GetCell("M", pos+2), GetCell("M", pos+3))

				f.SetCellDefault(sheet, GetCell("A", pos+1), fmt.Sprintf("●%v년도", i))
				f.DuplicateRowTo(sheet, 9, pos+4)
				f.MergeCell(sheet, GetCell("A", pos+4), GetCell("M", pos+4))
				f.SetCellDefault(sheet, GetCell("A", pos+4), "수선계획 없음")

				pos += 5

				continue
			}

			cnt := 0
			for _, v := range yearreports {
				if v.Duedate == year {
					cnt++
				}
			}

			if cnt == 0 {

				f.DuplicateRowTo(sheet, 2, pos)
				f.DuplicateRowTo(sheet, 3, pos+1)
				f.DuplicateRowTo(sheet, 4, pos+2)
				f.DuplicateRowTo(sheet, 5, pos+3)

				f.MergeCell(sheet, GetCell("A", pos+2), GetCell("A", pos+3))
				f.MergeCell(sheet, GetCell("B", pos+2), GetCell("B", pos+3))
				f.MergeCell(sheet, GetCell("C", pos+2), GetCell("C", pos+3))
				f.MergeCell(sheet, GetCell("D", pos+2), GetCell("D", pos+3))
				f.MergeCell(sheet, GetCell("E", pos+2), GetCell("E", pos+3))
				f.MergeCell(sheet, GetCell("F", pos+2), GetCell("F", pos+3))
				f.MergeCell(sheet, GetCell("I", pos+2), GetCell("I", pos+3))
				f.MergeCell(sheet, GetCell("J", pos+2), GetCell("J", pos+3))
				f.MergeCell(sheet, GetCell("K", pos+2), GetCell("K", pos+3))
				f.MergeCell(sheet, GetCell("L", pos+2), GetCell("L", pos+3))
				f.MergeCell(sheet, GetCell("M", pos+2), GetCell("M", pos+3))

				f.SetCellDefault(sheet, GetCell("A", pos+1), fmt.Sprintf("●%v년도", i))
				f.DuplicateRowTo(sheet, 9, pos+4)
				f.MergeCell(sheet, GetCell("A", pos+4), GetCell("M", pos+4))
				f.SetCellDefault(sheet, GetCell("A", pos+4), "수선계획 없음")

				pos += 5

				continue
			}

			f.DuplicateRowTo(sheet, 2, pos)
			f.DuplicateRowTo(sheet, 3, pos+1)
			f.DuplicateRowTo(sheet, 4, pos+2)
			f.DuplicateRowTo(sheet, 5, pos+3)

			f.MergeCell(sheet, GetCell("A", pos+2), GetCell("A", pos+3))
			f.MergeCell(sheet, GetCell("B", pos+2), GetCell("B", pos+3))
			f.MergeCell(sheet, GetCell("C", pos+2), GetCell("C", pos+3))
			f.MergeCell(sheet, GetCell("D", pos+2), GetCell("D", pos+3))
			f.MergeCell(sheet, GetCell("E", pos+2), GetCell("E", pos+3))
			f.MergeCell(sheet, GetCell("F", pos+2), GetCell("F", pos+3))
			f.MergeCell(sheet, GetCell("I", pos+2), GetCell("I", pos+3))
			f.MergeCell(sheet, GetCell("J", pos+2), GetCell("J", pos+3))
			f.MergeCell(sheet, GetCell("K", pos+2), GetCell("K", pos+3))
			f.MergeCell(sheet, GetCell("L", pos+2), GetCell("L", pos+3))
			f.MergeCell(sheet, GetCell("M", pos+2), GetCell("M", pos+3))

			f.SetCellDefault(sheet, GetCell("A", pos+1), fmt.Sprintf("●%v년도", year))

			pos += 4

			for j := 0; j < cnt; j++ {
				f.DuplicateRowTo(sheet, 6, pos+j)
			}

			f.DuplicateRowTo(sheet, 7, pos+cnt)
			f.DuplicateRowTo(sheet, 8, pos+cnt+1)
			f.DuplicateRowTo(sheet, 9, pos+cnt+2)

			var partTotal int64 = 0
			var allTotal int64 = 0
			cnt = 0
			for _, v := range yearreports {
				if v.Duedate == year {
					category := categoryMap[int64(v.Method)]
					standard := standardMap[int64(v.Standard)]
					f.SetCellInt(sheet, GetCell("A", pos+cnt), cnt+1)
					f.SetCellDefault(sheet, GetCell("B", pos+cnt), categoryMap[int64(v.Topcategory)].Name)
					f.SetCellDefault(sheet, GetCell("C", pos+cnt), categoryMap[int64(v.Subcategory)].Name)
					f.SetCellDefault(sheet, GetCell("D", pos+cnt), categoryMap[int64(v.Category)].Name)
					f.SetCellDefault(sheet, GetCell("E", pos+cnt), standard.Name)
					f.SetCellDefault(sheet, GetCell("F", pos+cnt), categoryMap[int64(v.Method)].Name)
					f.SetCellInt(sheet, GetCell("G", pos+cnt), category.Cycle)
					f.SetCellInt(sheet, GetCell("H", pos+cnt), category.Percent)
					f.SetCellDefault(sheet, GetCell("I", pos+cnt), standard.Unit)

					f.SetCellInt(sheet, GetCell("J", pos+cnt), int(v.Count))
					priceOne := CalculatePriceRate(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repairParcelrate))
					price := CalculateRepair(standard.Direct, standard.Labor, standard.Cost, float64(v.Rate), float64(repairParcelrate), v.Count, float64(category.Percent))
					if category.Percent == 100 {
						allTotal += price
					} else {
						partTotal += price
					}

					f.SetCellDefault(sheet, GetCell("K", pos+cnt), humanize.Comma(priceOne))
					f.SetCellDefault(sheet, GetCell("L", pos+cnt), humanize.Comma(price))

					cnt++
				}
			}

			f.SetCellDefault(sheet, GetCell("L", pos+cnt), humanize.Comma(partTotal))
			f.SetCellDefault(sheet, GetCell("L", pos+cnt+1), humanize.Comma(allTotal))
			f.SetCellDefault(sheet, GetCell("L", pos+cnt+2), humanize.Comma(allTotal+partTotal))

			pos += 3 + cnt
			no++
		}
	}

	for i := 10; i >= 2; i-- {
		f.RemoveRow(sheet, i)
	}

	if repair.Provision == 2 {
		f.DeleteSheet("법조문")
		f.SetSheetName("법조문 (2)", "법조문")
		f.SetCellStr("용어정리", "A1", "04. 용어 정리")
	} else {
		f.DeleteSheet("법조문 (2)")
		// 법조문
		sheet = "법조문"
		f.SetCellDefault(sheet, GetCell("A", 204), fmt.Sprintf("%v 아파트   관리사무소장                     (인)", apt.Name))
		f.SetCellDefault(sheet, GetCell("A", 246), fmt.Sprintf("%v 아파트 관리소장 ○ ○ ○   (인)", apt.Name))
	}

	// 검토보고서
	sheet = "검토보고서"

	f.SetCellStr(sheet, "A19", fmt.Sprintf("%v년 %v월", reportyear, reportmonth))
	f.SetCellStr(sheet, "F52", repair.Complex1)
	f.SetCellStr(sheet, "F58", fmt.Sprintf("%v년", repair.Completionyear))
	f.SetCellStr(sheet, "G58", fmt.Sprintf("%v월", repair.Completionmonth))

	strs := strings.Split(repair.Reviewcontent1, "-")
	if len(strs) == 2 {
		f.SetCellStr(sheet, "F53", fmt.Sprintf("%v년 %v월", strs[0], strs[1]))
	} else {
		f.SetCellStr(sheet, "F53", repair.Reviewcontent1)
	}

	/*
		cnt := 0
		for i, v := range *reviewdates {
			if i == 0 {
				continue
			}

			if v.Year == 0 || v.Month == 0 {
				continue
			}

			cnt++
		}

		pos = 57
		step := 0
		for i, v := range *reviewdates {
			if i == 0 {
				continue
			}

			if v.Year == 0 || v.Month == 0 {
				continue
			}

			if step == cnt-1 {
				f.SetCellStr(sheet, GetCell("B", pos+step), "○ 금  회   조  정    :")
			} else {
				f.SetCellStr(sheet, GetCell("B", pos+step), "○ 이  전   조  정    :")
			}

			f.SetCellStr(sheet, GetCell("F", pos+step), fmt.Sprintf("%v.", v.Year))
			f.SetCellStr(sheet, GetCell("G", pos+step), fmt.Sprintf("%v.", v.Month))

			step++
		}
	*/

	cnt := 0
	for _, v := range repairs {
		if v.Reportdate == "" {
			continue
		}

		if v.Reportdate > repair.Reportdate {
			continue
		}

		cnt++
	}

	pos = 59
	step := 0
	for _, v := range repairs {
		if v.Reportdate == "" {
			continue
		}

		if v.Reportdate > repair.Reportdate {
			continue
		}

		if step == cnt-1 {
			f.SetCellStr(sheet, GetCell("B", pos+step), "○ 금  회   조  정    :")
		} else {
			if v.Periodtype == re.PeriodtypeSometime {
				f.SetCellStr(sheet, GetCell("B", pos+step), "○ 수  시   조  정    :")
			} else {
				f.SetCellStr(sheet, GetCell("B", pos+step), "○ 정  기   조  정    :")
			}
		}

		strs := strings.Split(v.Reportdate, "-")

		year := ""
		month := ""

		if len(strs) >= 2 {
			year = strs[0]
			month = fmt.Sprintf("%v", global.Atoi(strs[1]))
		}
		f.SetCellStr(sheet, GetCell("F", pos+step), fmt.Sprintf("%v년", year))
		f.SetCellStr(sheet, GetCell("G", pos+step), fmt.Sprintf("%v월", month))

		step++
	}

	pos = 21
	for i, v := range reviewcontents {
		f.SetCellStr(sheet, GetCell("C", pos+i), v.Content)
	}

	// 검토보고서(2)
	sheet = "검토보고서(2)"

	{
		categorys := make([]models.Review, 0)
		for _, v := range reviews {
			flag := false
			for _, v2 := range categorys {
				if v.Subcategory == v2.Subcategory {
					flag = true
					break
				}
			}

			if flag == false {
				categorys = append(categorys, v)
			}
		}

		pos := 11

		for _, v := range categorys {
			f.DuplicateRowTo(sheet, 2, pos)
			f.DuplicateRowTo(sheet, 3, pos+1)

			category := categoryMap[v.Category]
			subcategory := categoryMap[category.Parent]
			topcategory := categoryMap[subcategory.Parent]
			f.SetCellStr(sheet, GetCell("A", pos), fmt.Sprintf("%v << %v", topcategory.Name, subcategory.Name))

			pos += 2

			rows := 0

			for _, v2 := range reviews {
				if v.Subcategory != v2.Subcategory {
					continue
				}

				rows++
			}

			i := 0
			for _, v2 := range reviews {
				if v.Subcategory != v2.Subcategory {
					continue
				}

				if i == rows-1 {
					f.DuplicateRowTo(sheet, 7, pos)
					f.DuplicateRowTo(sheet, 8, pos+1)
					f.DuplicateRowTo(sheet, 9, pos+2)
					f.DuplicateRowTo(sheet, 10, pos+3)

				} else {
					f.DuplicateRowTo(sheet, 4, pos)
					f.DuplicateRowTo(sheet, 5, pos+1)
					f.DuplicateRowTo(sheet, 6, pos+2)
				}

				f.SetCellStr(sheet, GetCell("A", pos), categoryMap[v2.Category].Name)

				if v2.Standard != 0 {
					f.SetCellStr(sheet, GetCell("B", pos), standardMap[v2.Standard].Name)
				}

				if v2.Method != 0 {
					f.SetCellStr(sheet, GetCell("C", pos), categoryMap[v2.Method].Name)
				}

				f.SetCellStr(sheet, GetCell("D", pos), v2.Cycle)
				f.SetCellStr(sheet, GetCell("E", pos), fmt.Sprintf("%v", v2.Percent))
				f.SetCellStr(sheet, GetCell("F", pos), humanize.Comma(int64(v2.Count)))
				f.SetCellStr(sheet, GetCell("G", pos), humanize.Comma(v2.Price))

				f.SetCellStr(sheet, GetCell("C", pos+1), v2.Content)
				f.SetCellStr(sheet, GetCell("C", pos+2), v2.Adjust)

				f.MergeCell(sheet, GetCell("A", pos), GetCell("A", pos+2))
				f.MergeCell(sheet, GetCell("C", pos+1), GetCell("G", pos+1))
				f.MergeCell(sheet, GetCell("C", pos+2), GetCell("G", pos+2))

				if i == rows-1 {
					pos += 4
				} else {
					pos += 3
				}

				i++
			}

		}

		f.RemoveRow(sheet, pos)

		for i := 10; i >= 2; i-- {
			f.RemoveRow(sheet, i)
		}
	}

	// 총론
	sheet = "총론"

	{
		f.SetCellStr(sheet, "A16", fmt.Sprintf("    3) 계획기간의 결정 : 수선계획상의 계획기간을 %v년으로 조정", repair.Planyears))

		var totalSize float64 = 0.0
		for _, v := range areas {
			totalSize += float64(v.Familycount) * float64(v.Size)
		}

		f.SetCellDefault(sheet, GetCell("I", 40), humanize.CommafWithDigits(areaTotal, 3))
		f.SetCellDefault(sheet, GetCell("K", 29), fmt.Sprintf("%v", topTotalSavepriceParcel*int64(GetPlanyears(repair.Planyears))))

		f.SetCellStr(sheet, "F4", repair.Complex1)
		strs := strings.Split(repair.Reviewcontent1, "-")
		if len(strs) == 2 {
			f.SetCellStr(sheet, "F5", fmt.Sprintf("%v년 %v월", strs[0], strs[1]))
		} else {
			f.SetCellStr(sheet, "F5", repair.Reviewcontent1)
		}

		f.SetCellStr(sheet, "C8", fmt.Sprintf("%v년 %v월", repair.Completionyear, repair.Completionmonth))

		f.SetCellInt(sheet, "E29", repair.Completionyear+1)
		f.SetCellInt(sheet, "I29", repair.Completionyear+(repair.Planyears-1))
		f.SetCellInt(sheet, "I30", repair.Planyears-1)

		f.SetCellInt(sheet, "I38", (repair.Planyears-1)*12)

		f.SetCellStr(sheet, "I37", fmt.Sprintf("적립개월수(%v년*12개월)", repair.Planyears-1))

		f.SetCellInt(sheet, "I31", global.Atoi(reportyear))

		cnt := 0
		for _, v := range repairs {
			if v.Reportdate == "" {
				continue
			}

			if v.Reportdate > repair.Reportdate {
				continue
			}

			cnt++
		}

		cellTitles := []string{"A", "F", "J"}
		cellDateTitles := []string{"C", "H", "L"}
		pos := 8
		step := 1

		for _, v := range repairs {
			if v.Reportdate == "" {
				continue
			}

			if v.Reportdate > repair.Reportdate {
				continue
			}

			if step == cnt {
				f.SetCellStr(sheet, GetCell(cellTitles[step%3], pos), fmt.Sprintf("    %v) 금회 조정 :", step+1))
			} else {
				if v.Periodtype == re.PeriodtypeSometime {
					f.SetCellStr(sheet, GetCell(cellTitles[step%3], pos), fmt.Sprintf("    %v) 수시 조정 :", step+1))
				} else {
					f.SetCellStr(sheet, GetCell(cellTitles[step%3], pos), fmt.Sprintf("    %v) 정기 조정 :", step+1))
				}
			}

			strs := strings.Split(v.Reportdate, "-")
			year := ""
			month := ""
			if len(strs) >= 2 {
				year = strs[0]
				month = strs[1]
			}
			f.SetCellStr(sheet, GetCell(cellDateTitles[step%3], pos), fmt.Sprintf("%v년 %v월", year, month))

			step++

			if step%3 == 0 {
				pos++
			}
		}

		f.SetCellDefault(sheet, GetCell("K", 54), fmt.Sprintf("%v", repair.Savingprice))

		outlineLen := len(outlines)
		outlineplanLen := len(outlineplans)

		for i := 0; i < 20-outlineLen; i++ {
			f.RemoveRow(sheet, 67)
		}

		for i := 0; i < 20-outlineplanLen; i++ {
			f.RemoveRow(sheet, 91-(20-outlineLen))
		}

		for i := 0; i < 20-outlineplanLen; i++ {
			f.RemoveRow(sheet, 116-(20-outlineLen)-(20-outlineplanLen))
			f.RemoveRow(sheet, 116-(20-outlineLen)-(20-outlineplanLen))
			f.RemoveRow(sheet, 116-(20-outlineLen)-(20-outlineplanLen))
			f.RemoveRow(sheet, 116-(20-outlineLen)-(20-outlineplanLen))
		}

		pos = 67
		var alltotalprice int64 = 0
		var rateSum float64 = 0.0
		for i, v := range outlines {
			f.SetCellDefault(sheet, GetCell("A", pos+i), fmt.Sprintf("%v구간", i+1))
			f.SetCellStr(sheet, GetCell("B", pos+i), fmt.Sprintf("%v.%02d", v.Startyear, v.Startmonth))
			f.SetCellStr(sheet, GetCell("D", pos+i), fmt.Sprintf("%v.%02d", v.Endyear, v.Endmonth))
			totalprice := float64(topTotalSaveprice) * float64(GetPlanyears(repair.Planyears)) / 100.0 * float64(v.Rate)

			if i == len(outlines)-1 {
				if rateSum == 0.0 {
					f.SetCellDefault(sheet, GetCell("E", pos+i), humanize.CommafWithDigits(totalprice, 2))
				} else {
					f.SetCellDefault(sheet, GetCell("E", pos+i), humanize.CommafWithDigits(float64(topTotalSaveprice)*float64(GetPlanyears(repair.Planyears))-float64(alltotalprice), 2))
				}
			} else {
				f.SetCellDefault(sheet, GetCell("E", pos+i), humanize.CommafWithDigits(totalprice, 2))
			}
			f.SetCellDefault(sheet, GetCell("G", pos+i), fmt.Sprintf("%v", v.Rate))

			//duration := getDurationMonth(v.Endyear, v.Endmonth, v.Startyear, v.Startmonth)
			//price := totalprice / (totalSize * float64(duration))
			//f.SetCellDefault(sheet, GetCell("I", pos+i), humanize.CommafWithDigits(float64(price), 2))
			f.SetCellDefault(sheet, GetCell("I", pos+i), humanize.CommafWithDigits(float64(v.Price), 2))
			rateSum += float64(v.Rate)
			f.SetCellDefault(sheet, GetCell("K", pos+i), fmt.Sprintf("%v", rateSum))
			f.SetCellDefault(sheet, GetCell("M", pos+i), fmt.Sprintf("%v", v.Remark))

			alltotalprice += int64(totalprice)
		}

		alltotalprice = 0
		pos = 91 - (20 - outlineLen)
		rateSum = 0.0
		for i, v := range outlineplans {
			f.SetCellDefault(sheet, GetCell("A", pos+i), fmt.Sprintf("%v구간", i+1))
			f.SetCellStr(sheet, GetCell("B", pos+i), fmt.Sprintf("%v.%02d", v.Startyear, v.Startmonth))
			f.SetCellStr(sheet, GetCell("D", pos+i), fmt.Sprintf("%v.%02d", v.Endyear, v.Endmonth))
			totalprice := float64(topTotalSaveprice) * float64(GetPlanyears(repair.Planyears)) / 100.0 * float64(v.Rate)

			if i == len(outlineplans)-1 {
				f.SetCellDefault(sheet, GetCell("E", pos+i), humanize.CommafWithDigits(float64(topTotalSaveprice)*float64(GetPlanyears(repair.Planyears))-float64(alltotalprice), 2))
			} else {
				f.SetCellDefault(sheet, GetCell("E", pos+i), humanize.CommafWithDigits(totalprice, 2))
			}
			f.SetCellDefault(sheet, GetCell("G", pos+i), fmt.Sprintf("%v", v.Rate))

			//duration := getDurationMonth(v.Endyear, v.Endmonth, v.Startyear, v.Startmonth)
			//price := totalprice / (totalSize * float64(duration))
			//f.SetCellDefault(sheet, GetCell("I", pos+i), humanize.CommafWithDigits(float64(price), 2))
			f.SetCellDefault(sheet, GetCell("I", pos+i), humanize.CommafWithDigits(float64(v.Price), 2))

			rateSum += float64(v.Rate)
			f.SetCellDefault(sheet, GetCell("K", pos+i), fmt.Sprintf("%v", rateSum))
			f.SetCellDefault(sheet, GetCell("M", pos+i), fmt.Sprintf("%v", v.Remark))

			alltotalprice += int64(totalprice)
		}

		pos = 116 - (20 - outlineLen) - (20 - outlineplanLen)
		for i, v := range outlineplans {
			totalprice := float64(topTotalSaveprice) * float64(GetPlanyears(repair.Planyears)) / 100.0 * float64(v.Rate)

			f.SetCellDefault(sheet, GetCell("A", pos+i*4), fmt.Sprintf("%v구간", i+1))
			f.SetCellDefault(sheet, GetCell("B", pos+i*4), fmt.Sprintf("%v", v.Startyear))
			f.SetCellDefault(sheet, GetCell("B", pos+i*4+3), fmt.Sprintf("%v", v.Endyear))
			f.SetCellDefault(sheet, GetCell("F", pos+i*4), humanize.CommafWithDigits(float64(topTotalSaveprice*int64(GetPlanyears(repair.Planyears))), 2))
			f.SetCellDefault(sheet, GetCell("F", pos+i*4+2), humanize.CommafWithDigits(areaTotal, 2))
			f.SetCellDefault(sheet, GetCell("J", pos+i*4+2), fmt.Sprintf("%v", v.Endyear-v.Startyear+1))
			f.SetCellDefault(sheet, GetCell("K", pos+i*4), fmt.Sprintf("%v", v.Rate))

			//result := (float64(topTotalSaveprice) * 49.0 * float64(v.Rate)) / 100.0 / (areaTotal * float64(v.Endyear-v.Startyear+1) * 12.0)
			//f.SetCellDefault(sheet, GetCell("C", pos+i*4), humanize.CommafWithDigits(result, 2))

			duration := getDurationMonth(v.Endyear, v.Endmonth, v.Startyear, v.Startmonth)
			price := totalprice / (totalSize * float64(duration))
			f.SetCellDefault(sheet, GetCell("C", pos+i*4), humanize.CommafWithDigits(float64(price), 2))
		}

		pos = 199 - (20 - outlineLen) - (20 - outlineplanLen) - (20-outlineplanLen)*4
		for i := 0; i < 20-len(areas); i++ {
			f.RemoveRow(sheet, pos)
		}

		for i, v := range areas {
			f.SetCellInt(sheet, GetCell("A", pos+i), i+1)
			f.SetCellDefault(sheet, GetCell("B", pos+i), v.Name)
			f.SetCellDefault(sheet, GetCell("C", pos+i), humanize.CommafWithDigits(float64(v.Size), 2))
			f.SetCellDefault(sheet, GetCell("E", pos+i), fmt.Sprintf("%v", v.Familycount))

			f.SetCellDefault(sheet, GetCell("G", pos+i), humanize.CommafWithDigits(float64(repair.Savingprice), 2))

			price := float64(v.Size) * float64(repair.Savingprice)
			f.SetCellDefault(sheet, GetCell("I", pos+i), humanize.CommafWithDigits(float64(price), 2))
			f.SetCellDefault(sheet, GetCell("L", pos+i), humanize.CommafWithDigits(price*float64(v.Familycount), 2))
		}

		/*
			f.SetCellDefault(sheet, GetCell("F", 10), humanize.CommafWithDigits(float64(topTotalSaveprice*49), 2))
			f.SetCellDefault(sheet, GetCell("F", 11), humanize.CommafWithDigits(areaTotal, 3))

			result := float64(topTotalSaveprice) * 49.0 / (areaTotal * 49.0 * 12.0)
			f.SetCellDefault(sheet, GetCell("A", 10), humanize.CommafWithDigits(result, 2))
		*/

		f.UpdateLinkedValue()

	}

	// 충당금적립
	sheet = "충당금적립"
	if repair.Provision == 2 {
		sheet = "수선적립금 적립"
	}

	{
		pos := 3
		height := 7

		for i, v := range savings {
			f.SetCellStr(sheet, GetCell("A", pos+i*height), fmt.Sprintf("%v년", v.Year))
			f.SetCellDefault(sheet, GetCell("F", pos+i*height), humanize.Comma(v.Forward))
			f.SetCellDefault(sheet, GetCell("F", pos+i*height+1), humanize.Comma(v.Interest))
			f.SetCellDefault(sheet, GetCell("F", pos+i*height+2), humanize.Comma(v.Surplus))
			f.SetCellDefault(sheet, GetCell("F", pos+i*height+3), humanize.Comma(v.Etc))
			f.SetCellDefault(sheet, GetCell("F", pos+i*height+4), humanize.Comma(v.Saving))
			f.SetCellDefault(sheet, GetCell("H", pos+i*height+5), humanize.Comma(v.Use))
			f.SetCellDefault(sheet, GetCell("F", pos+i*height+6), humanize.Comma(v.Forward+v.Interest+v.Surplus+v.Etc+v.Saving))
			f.SetCellDefault(sheet, GetCell("H", pos+i*height+6), humanize.Comma(v.Use))
			f.SetCellDefault(sheet, GetCell("J", pos+i*height+6), humanize.Comma(v.Forward+v.Interest+v.Surplus+v.Etc+v.Saving-v.Use))
		}
	}

	// 총론1
	sheet = "총론1"
	{
		pos := 5

		repair.Content1 = ReplaceWord(repair.Provision, repair.Content1)
		repair.Content2 = ReplaceWord(repair.Provision, repair.Content2)
		repair.Content3 = ReplaceWord(repair.Provision, repair.Content3)
		repair.Content4 = ReplaceWord(repair.Provision, repair.Content4)
		repair.Content5 = ReplaceWord(repair.Provision, repair.Content5)
		repair.Content6 = ReplaceWord(repair.Provision, repair.Content6)
		repair.Content7 = ReplaceWord(repair.Provision, repair.Content7)
		repair.Content8 = ReplaceWord(repair.Provision, repair.Content8)
		repair.Content9 = ReplaceWord(repair.Provision, repair.Content9)
		repair.Content10 = ReplaceWord(repair.Provision, repair.Content10)
		repair.Content11 = ReplaceWord(repair.Provision, repair.Content11)

		// 42 13
		length := global.Strlen(ReplaceWord(repair.Provision, "가 소액인 항목을 일부 교체 또는 수리가 필요할 경우 장기수선공사 관리대장에 기록 후 장기수선충당금"))

		f.DuplicateRowTo(sheet, 1, pos)

		f.SetCellDefault(sheet, GetCell("A", pos), ReplaceWord(repair.Provision, "03. 장기수선충당금 사용시 예외적 사항의 인정"))

		pos++

		l := global.Strlen(repair.Content1)
		rows := l / length
		if l%length > 0 {
			rows++
		}

		start := 0
		end := 0
		for i := 0; i < rows; i++ {
			f.DuplicateRowTo(sheet, 3, pos+i)

			start = i * length
			end = (i + 1) * length

			if end > l {
				end = l
			}

			str := global.Substr(repair.Content1, start, end)
			if str[0] == ' ' {
				str = str[1:]
			}
			f.SetCellStr(sheet, GetCell("A", pos+i), str)
		}

		pos += rows
		f.DuplicateRowTo(sheet, 4, pos)
		pos++

		f.DuplicateRowTo(sheet, 2, pos)
		f.SetCellDefault(sheet, GetCell("A", pos), "3.1 긴급공사 및 소액지출(예측불가한 사고)")
		pos++

		titles := []string{"(1) 내용", "(2) 수선대상", "(3) 근거마련", ReplaceWord(repair.Provision, "(4) 장기수선계획 조정시기"), "(5) 사용절차", "(6) 긴급공사 대상 시설물", "(7) 소액지출 사용요건"}
		contents := []string{repair.Content2, repair.Content3, repair.Content4, repair.Content5, repair.Content6, repair.Content7, repair.Content8}
		for i, title := range titles {
			f.DuplicateRowTo(sheet, 2, pos)
			f.SetCellDefault(sheet, GetCell("A", pos), title)
			pos++

			strs := strings.Split(contents[i], "\n")
			for _, v := range strs {
				pos = SplitString(f, sheet, pos, "● "+v)
			}

			f.DuplicateRowTo(sheet, 4, pos)
			pos++
		}

		f.DuplicateRowTo(sheet, 2, pos)
		f.SetCellDefault(sheet, GetCell("A", pos), ReplaceWord(repair.Provision, "3.2 장기수선충당금 사용 방법 및 절차"))
		pos++

		strs := strings.Split(repair.Content9, "\n")
		for i, v := range strs {
			pos = SplitString(f, sheet, pos, fmt.Sprintf("(%v) %v", i+1, v))

			f.DuplicateRowTo(sheet, 4, pos)
			pos++
		}

		f.DuplicateRowTo(sheet, 2, pos)
		f.SetCellDefault(sheet, GetCell("A", pos), ReplaceWord(repair.Provision, "3.3 장기수선충당금 사용 금액의 범위"))
		pos++

		strs = strings.Split(repair.Content10, "\n")
		for _, v := range strs {
			pos = SplitString(f, sheet, pos, "● "+v)
		}

		f.DuplicateRowTo(sheet, 4, pos)
		pos++

		f.DuplicateRowTo(sheet, 1, pos)
		f.SetCellDefault(sheet, GetCell("A", pos), ReplaceWord(repair.Provision, "04. 향후 장기수선충당금에 대한 고찰."))
		pos++

		pos = SplitString(f, sheet, pos, repair.Content11)

		f.RemoveRow(sheet, 1)
		f.RemoveRow(sheet, 1)
		f.RemoveRow(sheet, 1)
		f.RemoveRow(sheet, 1)
	}

	uniq := time.Now().UnixNano() / (1 << 22)
	now := time.Now()
	date := now.Format("20060102150405")
	filename := fmt.Sprintf("%v_%v.xlsx", date, uniq)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)

	e := f.SaveAs(fullFilename)
	if e != nil {
		log.Println(e)
	}
	f.Close()

	if _, err := os.Stat(fullFilename); errors.Is(err, os.ErrNotExist) {
		log.Println("file not found")
	} else {
		c.Download(fullFilename, fmt.Sprintf("장기수선계획-%v.xlsx", apt.Name))
		os.Remove(fullFilename)
	}
}

func SplitString(f *excelize.File, sheet string, pos int, str string) int {
	length := global.Strlen("가 소액인 항목을 일부 교체 또는 수리가 필요할 경우 장기수선공사 관리대장에 기록 후 장기수선충당금")

	l := global.Strlen(str)
	rows := l / length
	if l%length > 0 {
		rows++
	}

	sum := 0

	if length > l {
		f.DuplicateRowTo(sheet, 3, pos)
		f.SetCellStr(sheet, GetCell("A", pos), str)

		sum++
	} else {
		start := 0
		end := 0

		for i := 0; i < rows; i++ {
			start = i * length
			end = (i + 1) * length

			if end > l {
				end = l
			}

			if i == rows-1 {
				end += 3
			}

			nstr := ""

			nstr = strings.TrimSpace(global.Substr(str, start, end))

			if len(nstr) == 0 {
				continue
			}

			f.DuplicateRowTo(sheet, 3, pos+i)
			f.SetCellStr(sheet, GetCell("A", pos+i), nstr)

			sum++
		}
	}

	pos += sum

	return pos
}

func (c *DownloadController) Address() {
	conn := c.NewConnection()

	manager := models.NewAptlistManager(conn)
	items := manager.Find([]any{models.Ordering("a_id desc")})
	f, err := excelize.OpenFile("./doc/repair/address.xlsx")
	if err != nil {
		log.Println(err)
		return
	}

	sheet := "Sheet1"

	for i, item := range items {
		f.SetCellStr(sheet, GetCell("A", i+2), item.Name)
		f.SetCellStr(sheet, GetCell("B", i+2), item.Tel)
		f.SetCellStr(sheet, GetCell("C", i+2), item.Fax)
		f.SetCellStr(sheet, GetCell("D", i+2), item.Testdate)

		f.SetCellStr(sheet, GetCell("E", i+2), item.Repairdate)
		f.SetCellStr(sheet, GetCell("F", i+2), item.Periodicdate)
		f.SetCellStr(sheet, GetCell("G", i+2), item.Email)
		f.SetCellStr(sheet, GetCell("H", i+2), item.Personalemail)
		f.SetCellStr(sheet, GetCell("I", i+2), item.Zip)
		f.SetCellStr(sheet, GetCell("J", i+2), item.Address)
		f.SetCellStr(sheet, GetCell("K", i+2), item.Address2)

		apttype := item.Type

		if apttype == "0" {
			apttype = ""
		}
		f.SetCellStr(sheet, GetCell("L", i+2), item.Completeyear)
		f.SetCellStr(sheet, GetCell("M", i+2), apttype)
		f.SetCellStr(sheet, GetCell("N", i+2), item.Flatcount)
		f.SetCellStr(sheet, GetCell("O", i+2), item.Familycount)
		f.SetCellStr(sheet, GetCell("P", i+2), item.Floor)
		f.SetCellStr(sheet, GetCell("Q", i+2), item.Fmsloginid)
		f.SetCellStr(sheet, GetCell("R", i+2), item.Fmspasswd)
	}

	uniq := time.Now().UnixNano() / (1 << 22)
	now := time.Now()
	date := now.Format("20060102150405")
	filename := fmt.Sprintf("%v_%v.xlsx", date, uniq)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)

	e := f.SaveAs(fullFilename)
	if e != nil {
		log.Println(e)
	}
	f.Close()

	if _, err := os.Stat(fullFilename); errors.Is(err, os.ErrNotExist) {
		log.Println("file not found")
	} else {
		c.Download(fullFilename, "주소록.xlsx")
		os.Remove(fullFilename)
	}
}

func (c *DownloadController) Periodic0(id int64) {
	conn := c.NewConnection()

	filename := periodic.Periodic0(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Periodic1(id int64) {
	conn := c.NewConnection()

	filename := periodic.Periodic1(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Periodic2(id int64) {
	conn := c.NewConnection()

	filename := periodic.Periodic2(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Periodic3(id int64) {
	conn := c.NewConnection()

	filename := periodic.Periodic3(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Periodic4(id int64) {
	conn := c.NewConnection()

	filename := periodic.Periodic4(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Periodic5(id int64) {
	conn := c.NewConnection()

	filename := periodic.Periodic5(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Periodic(id int64) {
	conn := c.NewConnection()

	filenames := make([]string, 0)
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, periodic.Periodic0(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, periodic.Periodic1(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, periodic.Periodic2(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, periodic.Periodic3(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, periodic.Periodic4(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, periodic.Periodic5(id, conn)))

	files := []string{"00.본보고서.hml", "01.과업지시서.hml", "02.외관조사망도.hml", "03.시설물관리대장.hml", "04.사진자료.hml", "05.사전자료일체 및 기타참고자료.hml"}

	filename := fmt.Sprintf("periodic-%v-%v.zip", id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)

	global.MakeZipfile(fullFilename, files, filenames)
	c.Download(fullFilename, filename)
	for _, v := range filenames {
		os.Remove(v)
	}
	os.Remove(fullFilename)
}

func (c *DownloadController) Estimate(id int64, typeid int) {
	conn := c.NewConnection()

	filename := estimate.Estimate(id, typeid, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	log.Println(fullFilename)

	c.Download(fullFilename, filename)
	//os.Remove(fullFilename)
}

func (c *DownloadController) Contract(id int64, typeid int) {
	conn := c.NewConnection()

	filename := estimate.Contract(id, typeid, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	log.Println(fullFilename)

	c.Download(fullFilename, filename)
	//os.Remove(fullFilename)
}

func ReplaceWord(typeid int, str string) string {
	if typeid == 1 {
		return str
	}

	str = strings.ReplaceAll(str, "장기수선충당금", "수선적립금")
	str = strings.ReplaceAll(str, "장기수선계획서", "수선계획서")
	str = strings.ReplaceAll(str, "입주자대표회의", "관리단")
	str = strings.ReplaceAll(str, "장기수선", "수선")

	return str
}

func (c *DownloadController) Addressrepair() {
	conn := c.NewConnection()

	manager := models.NewAptrepairlistManager(conn)
	items := manager.Find([]any{models.Ordering("a_name")})
	f, err := excelize.OpenFile("./doc/repair/address-repair.xlsx")
	if err != nil {
		log.Println(err)
		return
	}

	sheet := "Sheet1"

	for i, item := range items {
		f.SetCellStr(sheet, GetCell("A", i+2), item.Name)
		f.SetCellStr(sheet, GetCell("B", i+2), item.Tel)
		f.SetCellStr(sheet, GetCell("C", i+2), item.Fax)
		f.SetCellStr(sheet, GetCell("D", i+2), item.Testdate)
		f.SetCellStr(sheet, GetCell("E", i+2), item.Email)
		f.SetCellStr(sheet, GetCell("F", i+2), item.Personalemail)
		f.SetCellStr(sheet, GetCell("G", i+2), item.Zip)
		f.SetCellStr(sheet, GetCell("H", i+2), item.Address)
		f.SetCellStr(sheet, GetCell("I", i+2), item.Address2)

		apttype := item.Type

		if apttype == "0" {
			apttype = ""
		}
		f.SetCellStr(sheet, GetCell("J", i+2), item.Completeyear)
		f.SetCellStr(sheet, GetCell("K", i+2), apttype)
		f.SetCellStr(sheet, GetCell("L", i+2), item.Flatcount)
		f.SetCellStr(sheet, GetCell("M", i+2), item.Familycount)
		f.SetCellStr(sheet, GetCell("N", i+2), item.Floor)
		f.SetCellStr(sheet, GetCell("O", i+2), item.Fmsloginid)
		f.SetCellStr(sheet, GetCell("P", i+2), item.Fmspasswd)
		f.SetCellStr(sheet, GetCell("Q", i+2), item.Reportdate)
	}

	uniq := time.Now().UnixNano() / (1 << 22)
	now := time.Now()
	date := now.Format("20060102150405")
	filename := fmt.Sprintf("%v_%v.xlsx", date, uniq)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)

	e := f.SaveAs(fullFilename)
	if e != nil {
		log.Println(e)
	}
	f.Close()

	if _, err := os.Stat(fullFilename); errors.Is(err, os.ErrNotExist) {
		log.Println("file not found")
	} else {
		c.Download(fullFilename, "장기수선주소록.xlsx")
		os.Remove(fullFilename)
	}
}

func (c *DownloadController) Detail0(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail0(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	log.Println(fullFilename, filename)
	c.Download(fullFilename, filename)
	// os.Remove(fullFilename)
}

func (c *DownloadController) Detail1(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail1(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail2(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail2(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail3(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail3(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail4(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail4(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail5(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail5(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail6(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail6(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail7(id int64) {
	conn := c.NewConnection()

	filename := detail.Detail7(id, conn)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	c.Download(fullFilename, filename)
	os.Remove(fullFilename)
}

func (c *DownloadController) Detail(id int64) {
	conn := c.NewConnection()

	filenames := make([]string, 0)
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail0(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail1(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail2(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail3(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail4(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail5(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail6(id, conn)))
	filenames = append(filenames, fmt.Sprintf("%v/%v", config.UploadPath, detail.Detail7(id, conn)))

	files := []string{"00.본보고서.hml", "01.과업지시서.hml", "02.외관조사망도.hml", "03.측정시험결과표.hml", "04.상대평가결과자료.hml", "05.시설물관리대장.hml", "06.사진자료.hml", "07.사전자료일체.hml"}

	filename := fmt.Sprintf("Detail-%v-%v.zip", id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)

	global.MakeZipfile(fullFilename, files, filenames)
	c.Download(fullFilename, filename)
	for _, v := range filenames {
		os.Remove(v)
	}
	os.Remove(fullFilename)
}

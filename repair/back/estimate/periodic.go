package estimate

import (
	"fmt"
	"log"
	"repair/global"
	"repair/global/config"
	"repair/global/time"
	"repair/models"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/xuri/excelize/v2"
)

func Periodic(id int64, typeid int, conn *models.Connection, estimate *models.Estimate, compareestimates []models.Compareestimate, apt *models.Apt) string {
	excelFilename := fmt.Sprintf("periodic%v.xlsx", estimate.Subtype)

	var complex int64 = 1
	if estimate.Subtype == 3 {
		complex = 2
	}

	switch typeid {
	case 1:
		if complex == 1 {
			excelFilename = "detail-compare1.xlsx"
		} else {
			excelFilename = "detail-compare1-3.xlsx"
		}
	case 2:
		excelFilename = "detail-compare2.xlsx"
	}

	f, err := excelize.OpenFile(fmt.Sprintf("./doc/estimate/%v", excelFilename))
	if err != nil {
		log.Println(err)
		return ""
	}

	sheet := "개요"

	t := time.ParseDate(estimate.Writedate)
	//no := fmt.Sprintf("ANB-%04d%02d%02d-%v", t.Year(), t.Month(), t.Day(), id)

	writedate := strings.Split(estimate.Writedate, "-")

	complateyear := ""
	str := strings.Split(apt.Completeyear, "-")

	if len(str) == 3 {
		complateyear = fmt.Sprintf("%v년 %v월", str[0], str[1])
	} else if len(str) == 2 {
		complateyear = fmt.Sprintf("%v년 %v월", str[0], str[1])
	} else {
		str := strings.Split(apt.Completeyear, " ")

		if len(str) == 3 {
			complateyear = fmt.Sprintf("%v %v", str[0], str[1])
		} else if len(str) == 2 {
			complateyear = fmt.Sprintf("%v %v", str[0], str[1])
		} else {
			complateyear = apt.Completeyear
		}
	}

	buildingSize := ""

	if apt.Type == "아파트" || apt.Familycount3 > 0 {
		flatcount := strings.Split(apt.Flatcount, "(")
		if len(flatcount) == 2 {
			all := strings.TrimSpace(flatcount[0])
			target := strings.TrimSpace(strings.ReplaceAll(flatcount[1], ")", ""))

			buildingSize = fmt.Sprintf("아파트 %v개동 %v세대 (해당동 : %v개동)", all, apt.Familycount, target)
		} else {
			buildingSize = fmt.Sprintf("아파트 %v개동 %v세대", flatcount[0], apt.Familycount)
		}
	} else {
		floor := ""
		if apt.Undergroundfloor > 0 {
			floor = fmt.Sprintf("지하%v층~", apt.Undergroundfloor)
		}

		if apt.Groundfloor != 0 {
			apt.Floor = fmt.Sprintf("%v", apt.Groundfloor)
		}

		floor += fmt.Sprintf("지상%v층", apt.Floor)

		if apt.Area != "" {
			floor += fmt.Sprintf("(연면적 : %v)", apt.Area)
		}

		buildingSize = floor
	}

	tel := ""
	fax := ""

	if apt.Tel != "" {
		tel = fmt.Sprintf("전화 : %v", apt.Tel)
	}

	if apt.Fax != "" {
		fax = fmt.Sprintf("팩스 : %v", apt.Fax)
	}

	if fax != "" {
		tel += ",      " + fax
	}

	if estimate.Days == 0 {
		estimate.Days = 1
	}

	if typeid == 0 {
		if apt.Type == "아파트" || apt.Familycount3 > 0 {
			f.SetCellStr(sheet, "B12", "공동주택")
		} else {
			f.SetCellStr(sheet, "B12", "공동주택외 건축물")
		}

		f.SetCellStr(sheet, "B4", apt.Name)
		switch estimate.Subtype {
		case 1:
			f.SetCellStr(sheet, "C5", fmt.Sprintf("(%v년 상반기)", t.Year()))
		case 2:
			f.SetCellStr(sheet, "C5", fmt.Sprintf("(%v년 하반기)", t.Year()))
		case 3:
			f.SetCellStr(sheet, "C5", fmt.Sprintf("(%v년 연간)", t.Year()))
		default:
			f.SetCellStr(sheet, "C5", fmt.Sprintf("(%v년)", t.Year()))
		}

		switch estimate.Subtype {
		case 1:
			f.SetCellStr(sheet, "B18", fmt.Sprintf("%v년 상반기 정기안전점검 1회 금액", t.Year()))
		case 2:
			f.SetCellStr(sheet, "B18", fmt.Sprintf("%v년 하반기 정기안전점검 1회 금액", t.Year()))
		case 3:
			f.SetCellStr(sheet, "B18", fmt.Sprintf("%v년 연간 정기안전점검 금액", t.Year()))
		default:
			f.SetCellStr(sheet, "B18", fmt.Sprintf("%v년 정기안전점검 금액", t.Year()))
		}

		f.SetCellStr(sheet, "B6", fmt.Sprintf("%v 정기안전점검 용역", apt.Name))
		f.SetCellStr(sheet, "B8", apt.Address)
		f.SetCellStr(sheet, "B9", t.Humandate())
		f.SetCellStr(sheet, "B10", fmt.Sprintf("%v년", writedate[0]))

		f.SetCellStr(sheet, "B11", buildingSize)

		f.SetCellStr(sheet, "B13", complateyear)

		f.SetCellStr(sheet, "B14", tel)
		//f.SetCellStr(sheet, "B15", fax)

		f.SetCellValue(sheet, "B19", estimate.Price*int(complex))
		f.SetCellValue(sheet, "H29", estimate.Price)

		for i, v := range compareestimates {
			col := "C3"
			if i == 1 {
				col = "D3"
			}
			f.SetCellStr(sheet, col, v.Extra["comparecompany"].(models.Comparecompany).Name)

			col = "C19"
			if i == 1 {
				col = "D19"
			}
			f.SetCellValue(sheet, col, v.Price*int(complex))
		}

		sheet = "대가산출"

		f.SetCellValue(sheet, "E9", estimate.Personprice7)
		f.SetCellValue(sheet, "E10", estimate.Personprice8)
		f.SetCellValue(sheet, "E11", estimate.Personprice9)
		f.SetCellValue(sheet, "E12", estimate.Personprice10)

		f.SetCellValue(sheet, "E13", estimate.Personprice2)
		f.SetCellValue(sheet, "E14", estimate.Personprice3)
		f.SetCellValue(sheet, "E15", estimate.Personprice4)
		f.SetCellValue(sheet, "E16", estimate.Personprice5)

		f.SetCellValue(sheet, "G9", estimate.Person7*estimate.Days)
		f.SetCellValue(sheet, "G10", estimate.Person8*estimate.Days)
		f.SetCellValue(sheet, "G11", estimate.Person9*estimate.Days)
		f.SetCellValue(sheet, "G12", estimate.Person10*estimate.Days)

		f.SetCellValue(sheet, "G13", estimate.Person2)
		f.SetCellValue(sheet, "G14", estimate.Person3)
		f.SetCellValue(sheet, "G15", estimate.Person4)
		f.SetCellValue(sheet, "G16", estimate.Person5)

		f.SetCellValue(sheet, "F21", 1)
		f.SetCellValue(sheet, "F22", 1)

		outPersons := (estimate.Person7 + estimate.Person8 + estimate.Person9 + estimate.Person10) * estimate.Days
		f.SetCellValue(sheet, "G21", outPersons)
		f.SetCellValue(sheet, "G22", outPersons)

		if estimate.Person7 > 0 {
			f.SetCellValue(sheet, "I9", fmt.Sprintf("%v인 * %v일", estimate.Person7, estimate.Days))
		}
		if estimate.Person8 > 0 {
			f.SetCellValue(sheet, "I10", fmt.Sprintf("%v인 * %v일", estimate.Person8, estimate.Days))
		}
		if estimate.Person9 > 0 {
			f.SetCellValue(sheet, "I11", fmt.Sprintf("%v인 * %v일", estimate.Person9, estimate.Days))
		}
		if estimate.Person10 > 0 {
			f.SetCellValue(sheet, "I12", fmt.Sprintf("%v인 * %v일", estimate.Person10, estimate.Days))
		}

		// if estimate.Person2 > 0 {
		// 	f.SetCellValue(sheet, "I13", fmt.Sprintf("%v인 * 1일", estimate.Person2))
		// }
		// if estimate.Person3 > 0 {
		// 	f.SetCellValue(sheet, "I14", fmt.Sprintf("%v인 * 1일", estimate.Person3))
		// }
		// if estimate.Person4 > 0 {
		// 	f.SetCellValue(sheet, "I15", fmt.Sprintf("%v인 * 1일", estimate.Person4))
		// }
		// if estimate.Person5 > 0 {
		// 	f.SetCellValue(sheet, "I16", fmt.Sprintf("%v인 * 1일", estimate.Person5))
		// }

		f.SetCellValue(sheet, "G17", estimate.Financialprice)
		f.SetCellValue(sheet, "G18", estimate.Techprice)

		f.SetCellValue(sheet, "I17", fmt.Sprintf("직접인건비 * %v%%", estimate.Financialprice))
		f.SetCellValue(sheet, "I18", fmt.Sprintf("(직접인건비 + 제경비) * %v%%", estimate.Techprice))

		f.SetCellValue(sheet, "E21", estimate.Travelprice)
		f.SetCellValue(sheet, "E22", estimate.Carprice)

		f.SetCellValue(sheet, "G23", estimate.Danger)
		f.SetCellValue(sheet, "I23", fmt.Sprintf("외업인건비의 %v%%", estimate.Danger))

		f.SetCellValue(sheet, "G24", estimate.Machine)
		f.SetCellValue(sheet, "I24", fmt.Sprintf("직접인건비의 %v%%", estimate.Machine))

		f.SetCellValue(sheet, "E25", estimate.Printprice)
		f.SetCellValue(sheet, "G25", estimate.Print)

		f.SetCellValue(sheet, "H28", estimate.Saleprice)

		f.SetCellStr(sheet, "D3", fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price))))

		sheet = "에이앤비 갑지"

		switch estimate.Subtype {
		case 3:
			f.SetCellStr(sheet, "D45", fmt.Sprintf("      상반기 정기안전점검 : %v원 (VAT 별도)", humanize.Comma(int64(estimate.Price))))
			f.SetCellStr(sheet, "D46", fmt.Sprintf("      하반기 정기안전점검 : %v원 (VAT 별도)", humanize.Comma(int64(estimate.Price))))
			f.SetCellStr(sheet, "D47", fmt.Sprintf("※ 총 액 (연2회) : 일금 %v원정(₩ %v)- VAT 별도 ※", global.HumanMoney(int64(estimate.Price*2)), humanize.Comma(int64(estimate.Price*2))))
		case 4:
			f.SetCellStr(sheet, "D43", fmt.Sprintf("(단, 계약 위반시 원수립금액 %v원-VAT 별도 청구)", humanize.Comma(int64(estimate.Price*2))))

			temp := strings.Split(estimate.Start, "-")
			year := global.Atoi(temp[0])
			month := global.Atoi(temp[1])

			if month == 1 {
				f.SetCellStr(sheet, "D45", fmt.Sprintf("라) %v년 상반기/하반기, %v년 상반기/하반기, %v년 상반기 (정기점검 5회)", year, year+1, year+2))
				f.SetCellStr(sheet, "D46", fmt.Sprintf("     %v년 하반기 정밀점검 (1회) (추후 견적금액 상의 예정)", year+2))
			} else {
				f.SetCellStr(sheet, "D45", fmt.Sprintf("라) %v년 하반기, %v년 상반기/하반기, %v년 상반기/하반기 (정기점검 5회)", year, year+1, year+2))
				f.SetCellStr(sheet, "D46", fmt.Sprintf("     %v년 상반기 정밀점검 (1회) (추후 견적금액 상의 예정)", year+3))
			}

			f.SetCellStr(sheet, "D50", fmt.Sprintf("※ 총 액 (정기점검 5회) : 일금 %v원정(₩ %v)- VAT 별도 ※", global.HumanMoney(int64(estimate.Price*5)), humanize.Comma(int64(estimate.Price*5))))
		}

		subtitle := ""
		switch estimate.Subtype {
		case 1:
			f.SetCellStr(sheet, "I43", fmt.Sprintf("(%v년 상반기 정기안전점검 1회 금액)", t.Year()))
			subtitle = fmt.Sprintf("%v년 상반기", t.Year())
		case 2:
			f.SetCellStr(sheet, "I43", fmt.Sprintf("(%v년 하반기 정기안전점검 1회 금액)", t.Year()))
			subtitle = fmt.Sprintf("%v년 하반기", t.Year())
		case 3:
			f.SetCellStr(sheet, "I43", fmt.Sprintf("(%v년 연간 정기안전점검 금액)", t.Year()))
			subtitle = fmt.Sprintf("%v년 연간", t.Year())
		default:
			f.SetCellStr(sheet, "I43", fmt.Sprintf("(%v년 정기안전점검 금액)", t.Year()))
			subtitle = fmt.Sprintf("%v년", t.Year())
		}

		f.SetCellStr(sheet, "G16", subtitle)

		sheet = "계약서1"

		priceStr := fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))
		priceStr2 := fmt.Sprintf("일금 %v원정(₩%v)", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))
		priceStr3 := fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))

		f.SetCellStr(sheet, "B9", apt.Name)
		switch estimate.Subtype {
		case 1:
			f.SetCellStr(sheet, "B16", fmt.Sprintf("%v년 상반기", t.Year()))
		case 2:
			f.SetCellStr(sheet, "B16", fmt.Sprintf("%v년 하반기", t.Year()))
		case 3:
			f.SetCellStr(sheet, "B16", fmt.Sprintf("%v년 연간", t.Year()))
			priceStr = fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(estimate.Price)*2), humanize.Comma(int64(estimate.Price)*2))
			priceStr2 = fmt.Sprintf("일금 %v원정(₩%v)", global.HumanMoney(int64(estimate.Price)*2), humanize.Comma(int64(estimate.Price)*2))
		default:
			f.SetCellStr(sheet, "B16", fmt.Sprintf("%v년", t.Year()))
		}

		f.SetCellStr(sheet, "G37", fmt.Sprintf("%v 정기안전점검 용역", apt.Name))
		f.SetCellStr(sheet, "G38", priceStr)
		f.SetCellStr(sheet, "G39", priceStr2)

		contractManager := models.NewContractManager(conn)
		contract := contractManager.GetByEstimate(estimate.Id)

		if contract != nil {
			startdate := time.ParseDate(contract.Contractstartdate)
			enddate := time.ParseDate(contract.Contractenddate)
			contractDate := time.ParseDate(contract.Contractdate)

			if startdate != nil && enddate != nil {
				f.SetCellStr(sheet, "G41", fmt.Sprintf("%04d   .  %02d .  %02d .   ~   %04d .  %02d .  %02d . ", startdate.Year(), startdate.Month(), startdate.Day(), enddate.Year(), enddate.Month(), enddate.Day()))
			} else if enddate != nil {
				f.SetCellStr(sheet, "G41", fmt.Sprintf("%04d   .     .     .   ~   %04d .  %02d .  %02d . ", t.Year(), enddate.Year(), enddate.Month(), enddate.Day()))
			} else {
				f.SetCellStr(sheet, "G41", fmt.Sprintf("%04d   .     .     .   ~   %04d .     .     . ", t.Year(), t.Year()))
			}
			if contractDate != nil {
				f.SetCellStr(sheet, "E53", contractDate.Humandate())
			} else {
				f.SetCellStr(sheet, "E53", fmt.Sprintf("%v년", t.Year()))
			}
		} else {
			f.SetCellStr(sheet, "G41", fmt.Sprintf("%04d   .     .     .   ~   %04d .     .     . ", t.Year(), t.Year()))
			f.SetCellStr(sheet, "E53", fmt.Sprintf("%v년", t.Year()))
		}

		f.SetCellStr(sheet, "G42", apt.Address)

		f.SetCellStr(sheet, "G43", tel)

		if estimate.Subtype == 3 {
			f.SetCellStr(sheet, "G45", fmt.Sprintf("상기 금액은 %v 정기안전점검 용역대가임. (연 2회)\n  - 1회 : %v", subtitle, priceStr3))
		} else {
			f.SetCellStr(sheet, "G45", fmt.Sprintf("상기 금액은 %v 정기안전점검 용역대가임.", subtitle))
		}
		if apt.Type == "아파트" || apt.Familycount3 > 0 {
			f.SetCellStr(sheet, "F71", "공동주택")
		} else {
			f.SetCellStr(sheet, "F71", "공동주택외 건축물")
		}
		f.SetCellStr(sheet, "K72", buildingSize)

		f.SetCellValue(sheet, "I151", estimate.Personprice7)
		f.SetCellValue(sheet, "I152", estimate.Personprice8)
		f.SetCellValue(sheet, "I153", estimate.Personprice9)
		f.SetCellValue(sheet, "I154", estimate.Personprice10)

		f.SetCellValue(sheet, "I155", estimate.Personprice2)
		f.SetCellValue(sheet, "I156", estimate.Personprice3)
		f.SetCellValue(sheet, "I157", estimate.Personprice4)
		f.SetCellValue(sheet, "I158", estimate.Personprice5)

		f.SetCellValue(sheet, "P151", estimate.Person7*estimate.Days)
		f.SetCellValue(sheet, "P152", estimate.Person8*estimate.Days)
		f.SetCellValue(sheet, "P153", estimate.Person9*estimate.Days)
		f.SetCellValue(sheet, "P154", estimate.Person10*estimate.Days)

		f.SetCellValue(sheet, "P155", estimate.Person2)
		f.SetCellValue(sheet, "P156", estimate.Person3)
		f.SetCellValue(sheet, "P157", estimate.Person4)
		f.SetCellValue(sheet, "P158", estimate.Person5)

		f.SetCellValue(sheet, "P159", estimate.Financialprice)
		f.SetCellValue(sheet, "P160", estimate.Techprice)
		f.SetCellValue(sheet, "I164", estimate.Carprice)
		f.SetCellValue(sheet, "I163", estimate.Travelprice)
		f.SetCellValue(sheet, "I165", fmt.Sprintf("외업인건비의 %v%%", estimate.Danger))
		f.SetCellFormula(sheet, "R165", fmt.Sprintf("=ROUND(SUM(R151:R154)*%v%%, 0)", estimate.Danger))
		f.SetCellValue(sheet, "I166", fmt.Sprintf("직접인건비의 %v%%", estimate.Machine))
		f.SetCellFormula(sheet, "R166", fmt.Sprintf("=ROUND(R150*%v%%, 0)", estimate.Machine))
		f.SetCellValue(sheet, "R167", estimate.Printprice)

		f.SetCellValue(sheet, "R170", estimate.Saleprice)
		f.SetCellValue(sheet, "R171", estimate.Price)
		f.SetCellValue(sheet, "R173", estimate.Price)

		f.UpdateLinkedValue()
	} else if typeid == 1 {

		// 비교 견적
		var compareestimate models.Compareestimate
		for _, v := range compareestimates {
			if v.Comparecompany == 1 {
				compareestimate = v
				break
			}
		}
		sheet = "엘림공문-비교견적"

		title := ""
		switch estimate.Subtype {
		case 1:
			title = fmt.Sprintf("%v %v년 상반기 정기안전점검", apt.Name, t.Year())
		case 2:
			title = fmt.Sprintf("%v %v년 하반기 정기안전점검", apt.Name, t.Year())
		case 3:
			title = fmt.Sprintf("%v %v년 연간 정기안전점검", apt.Name, t.Year())
		default:
			title = fmt.Sprintf("%v %v년 정기안전점검", apt.Name, t.Year())
		}

		priceStr := fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(compareestimate.Price)*complex), humanize.Comma(int64(compareestimate.Price)*complex))
		price1Str := fmt.Sprintf("일금 %v원정", global.HumanMoney(int64(compareestimate.Price)*complex))
		price2Str := fmt.Sprintf("(₩%v) - VAT 별도", humanize.Comma(int64(compareestimate.Price)*complex))

		t = time.ParseDate(compareestimate.Writedate)
		f.SetCellStr(sheet, "C8", fmt.Sprintf("ELIM-%04d%02d%02d", t.Year(), t.Month(), t.Day()))
		f.SetCellStr(sheet, "C9", t.Humandate())
		f.SetCellStr(sheet, "C11", apt.Name)
		f.SetCellStr(sheet, "C18", title)
		f.SetCellStr(sheet, "C29", title)

		f.SetCellStr(sheet, "C30", "- 육안조사 통한 현장조사\n- 시특법에 의한 정기점검 실시 및 보고서 작성\n- FMS등록 제출업무")

		sheet = "엘림대가내역서 갑지"

		f.SetCellStr(sheet, "B7", title)
		f.SetCellStr(sheet, "B10", buildingSize)
		f.SetCellStr(sheet, "B13", price1Str)
		f.SetCellStr(sheet, "D13", price2Str)
		f.SetCellValue(sheet, "F18", compareestimate.Price)

		if complex == 1 {
			sheet = "엘림대가내역서"
		} else {
			sheet = "엘림 상하반기 대가내역서"

			f.SetCellStr(sheet, "B3", fmt.Sprintf("%v %v년 건축물 정기안전점검(상반기) 용역", apt.Name, t.Year()))
			f.SetCellStr(sheet, "B31", fmt.Sprintf("%v %v년 건축물 정기안전점검(하반기) 용역", apt.Name, t.Year()))
		}

		priceStr = fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(compareestimate.Price)), humanize.Comma(int64(compareestimate.Price)))
		f.SetCellStr(sheet, "B5", priceStr)

		f.SetCellValue(sheet, "F9", compareestimate.Personprice7)
		f.SetCellValue(sheet, "F10", compareestimate.Personprice8)
		f.SetCellValue(sheet, "F11", compareestimate.Personprice9)
		f.SetCellValue(sheet, "F12", compareestimate.Personprice10)

		f.SetCellValue(sheet, "F13", compareestimate.Personprice2)
		f.SetCellValue(sheet, "F14", compareestimate.Personprice3)
		f.SetCellValue(sheet, "F15", compareestimate.Personprice4)
		f.SetCellValue(sheet, "F16", compareestimate.Personprice5)

		if compareestimate.Person7 > 0 {
			f.SetCellValue(sheet, "D9", compareestimate.Person7)
			f.SetCellValue(sheet, "E9", estimate.Days)
		}
		if compareestimate.Person8 > 0 {
			f.SetCellValue(sheet, "D10", compareestimate.Person8)
			f.SetCellValue(sheet, "E10", estimate.Days)
		}
		if compareestimate.Person9 > 0 {
			f.SetCellValue(sheet, "D11", compareestimate.Person9)
			f.SetCellValue(sheet, "E11", estimate.Days)
		}
		if compareestimate.Person10 > 0 {
			f.SetCellValue(sheet, "D12", compareestimate.Person10)
			f.SetCellValue(sheet, "E12", estimate.Days)
		}

		if compareestimate.Person2 > 0 {
			f.SetCellValue(sheet, "E13", compareestimate.Person2)
			f.SetCellValue(sheet, "D13", 1)
		}
		if compareestimate.Person3 > 0 {
			f.SetCellValue(sheet, "E14", compareestimate.Person3)
			f.SetCellValue(sheet, "D14", 1)
		}
		if compareestimate.Person4 > 0 {
			f.SetCellValue(sheet, "E15", compareestimate.Person4)
			f.SetCellValue(sheet, "D15", 1)
		}
		if compareestimate.Person5 > 0 {
			f.SetCellValue(sheet, "E16", compareestimate.Person5)
			f.SetCellValue(sheet, "D16", 1)
		}

		f.SetCellValue(sheet, "E20", estimate.Days)
		f.SetCellValue(sheet, "E21", estimate.Days)

		f.SetCellValue(sheet, "C18", compareestimate.Financialprice)
		f.SetCellValue(sheet, "C19", compareestimate.Techprice)
		f.SetCellValue(sheet, "F20", compareestimate.Carprice)
		f.SetCellValue(sheet, "F21", compareestimate.Travelprice)
		f.SetCellValue(sheet, "C22", compareestimate.Danger)
		f.SetCellValue(sheet, "C23", compareestimate.Machine)
		f.SetCellValue(sheet, "G24", compareestimate.Printprice)

		f.SetCellValue(sheet, "G28", compareestimate.Price)

		f.UpdateLinkedValue()
	} else if typeid == 2 {
		// 비교 견적
		var compareestimate models.Compareestimate
		for _, v := range compareestimates {
			if v.Comparecompany == 2 {
				compareestimate = v
				break
			}
		}
		sheet = "갑지"

		title := ""

		switch estimate.Subtype {
		case 1:
			title = fmt.Sprintf("%v년 상반기 정기안전점검", t.Year())
		case 2:
			title = fmt.Sprintf("%v년 하반기 정기안전점검", t.Year())
		case 3:
			title = fmt.Sprintf("%v년 연간 정기안전점검", t.Year())
		default:
			title = fmt.Sprintf("%v년 정기안전점검", t.Year())
		}

		t = time.ParseDate(compareestimate.Writedate)
		f.SetCellStr(sheet, "A6", apt.Name)

		if estimate.Subtype == 3 {
			title = fmt.Sprintf("%v년 상반기 정기안전점검", t.Year())
		}

		f.SetCellStr(sheet, "A9", title)
		f.SetCellStr(sheet, "D24", buildingSize)
		f.SetCellStr(sheet, "I10", t.Humandate())

		sheet = "을지"

		f.SetCellValue(sheet, "G6", compareestimate.Personprice7)
		f.SetCellValue(sheet, "G7", compareestimate.Personprice8)
		f.SetCellValue(sheet, "G8", compareestimate.Personprice9)
		f.SetCellValue(sheet, "G9", compareestimate.Personprice10)

		if compareestimate.Person7 > 0 {
			f.SetCellValue(sheet, "C6", compareestimate.Person7*estimate.Days)
		}
		if compareestimate.Person8 > 0 {
			f.SetCellValue(sheet, "C7", compareestimate.Person8*estimate.Days)
		}
		if compareestimate.Person9 > 0 {
			f.SetCellValue(sheet, "C8", compareestimate.Person9*estimate.Days)
		}
		if compareestimate.Person10 > 0 {
			f.SetCellValue(sheet, "C9", compareestimate.Person10*estimate.Days)
		}

		if compareestimate.Person2 > 0 {
			f.SetCellValue(sheet, "E6", compareestimate.Person2)
		}
		if compareestimate.Person3 > 0 {
			f.SetCellValue(sheet, "E7", compareestimate.Person3)
		}
		if compareestimate.Person4 > 0 {
			f.SetCellValue(sheet, "E8", compareestimate.Person4)
		}
		if compareestimate.Person5 > 0 {
			f.SetCellValue(sheet, "E9", compareestimate.Person5)
		}

		f.SetCellValue(sheet, "G11", compareestimate.Travelprice)
		f.SetCellValue(sheet, "G13", compareestimate.Carprice)
		f.SetCellValue(sheet, "F15", compareestimate.Danger)
		f.SetCellValue(sheet, "F16", compareestimate.Machine)
		f.SetCellValue(sheet, "G17", compareestimate.Printprice)

		f.SetCellValue(sheet, "B19", compareestimate.Financialprice)
		f.SetCellValue(sheet, "B20", compareestimate.Techprice)

		f.SetCellValue(sheet, "F27", compareestimate.Saleprice)
		f.SetCellValue(sheet, "H27", compareestimate.Price)

		f.UpdateLinkedValue()
	}

	filename := fmt.Sprintf("%v.xlsx", global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	e := f.SaveAs(fullFilename)
	if e != nil {
		log.Println(e)
	}
	f.Close()

	if estimate.Subtype == 3 && typeid == 2 {
		f, err = excelize.OpenFile(fmt.Sprintf("./doc/estimate/%v", excelFilename))
		if err != nil {
			log.Println(err)
			return ""
		}

		var compareestimate models.Compareestimate
		for _, v := range compareestimates {
			if v.Comparecompany == 2 {
				compareestimate = v
				break
			}
		}
		sheet = "갑지"

		title := ""

		switch estimate.Subtype {
		case 1:
			title = fmt.Sprintf("%v년 상반기 정기안전점검", t.Year())
		case 2:
			title = fmt.Sprintf("%v년 하반기 정기안전점검", t.Year())
		case 3:
			title = fmt.Sprintf("%v년 연간 정기안전점검", t.Year())
		default:
			title = fmt.Sprintf("%v년 정기안전점검", t.Year())
		}

		t = time.ParseDate(compareestimate.Writedate)
		f.SetCellStr(sheet, "A6", apt.Name)

		if estimate.Subtype == 3 {
			title = fmt.Sprintf("%v년 하반기 정기안전점검", t.Year())
		}

		f.SetCellStr(sheet, "A9", title)
		f.SetCellStr(sheet, "D24", buildingSize)
		f.SetCellStr(sheet, "I10", t.Humandate())

		sheet = "을지"

		f.SetCellValue(sheet, "G6", compareestimate.Personprice7)
		f.SetCellValue(sheet, "G7", compareestimate.Personprice8)
		f.SetCellValue(sheet, "G8", compareestimate.Personprice9)
		f.SetCellValue(sheet, "G9", compareestimate.Personprice10)

		if estimate.Person7 > 0 {
			f.SetCellValue(sheet, "C6", compareestimate.Person7*estimate.Days)
		}
		if estimate.Person8 > 0 {
			f.SetCellValue(sheet, "C7", compareestimate.Person8*estimate.Days)
		}
		if estimate.Person9 > 0 {
			f.SetCellValue(sheet, "C8", compareestimate.Person9*estimate.Days)
		}
		if estimate.Person10 > 0 {
			f.SetCellValue(sheet, "C9", compareestimate.Person10*estimate.Days)
		}

		if estimate.Person2 > 0 {
			f.SetCellValue(sheet, "E6", compareestimate.Person2)
		}
		if estimate.Person3 > 0 {
			f.SetCellValue(sheet, "E7", compareestimate.Person3)
		}
		if estimate.Person4 > 0 {
			f.SetCellValue(sheet, "E8", compareestimate.Person4)
		}
		if estimate.Person5 > 0 {
			f.SetCellValue(sheet, "E9", compareestimate.Person5)
		}

		f.SetCellValue(sheet, "G11", compareestimate.Travelprice)
		f.SetCellValue(sheet, "G13", compareestimate.Carprice)
		f.SetCellValue(sheet, "F15", compareestimate.Danger)
		f.SetCellValue(sheet, "F16", compareestimate.Machine)
		f.SetCellValue(sheet, "G17", compareestimate.Printprice)

		f.SetCellValue(sheet, "B19", compareestimate.Financialprice)
		f.SetCellValue(sheet, "B20", compareestimate.Techprice)

		f.SetCellValue(sheet, "F27", compareestimate.Saleprice)
		f.SetCellValue(sheet, "H27", compareestimate.Price)

		f.UpdateLinkedValue()

		filename = fmt.Sprintf("%v.xlsx", global.UniqueId())
		fullFilename2 := fmt.Sprintf("%v/%v", config.UploadPath, filename)
		e := f.SaveAs(fullFilename2)
		if e != nil {
			log.Println(e)
		}
		f.Close()

		fname := []string{"상반기 정기안전점검.xlsx", "하반기 정기안전점검.xlsx"}
		files := []string{fullFilename, fullFilename2}
		filename = fmt.Sprintf("%v.zip", global.UniqueId())
		fullFilename = fmt.Sprintf("%v/%v", config.UploadPath, filename)
		global.MakeZipfile(fullFilename, fname, files)
	}

	return filename
}

package estimate

import (
	"fmt"
	"log"
	"repair/global"
	"repair/global/config"
	"repair/global/time"
	"repair/models"
	"strings"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
	"github.com/dustin/go-humanize"
	"github.com/xuri/excelize/v2"
)

func Detail(estimateType int, id int64, typeid int, conn *models.Connection, estimate *models.Estimate, compareestimates []models.Compareestimate, apt *models.Apt) string {
	excelFilename := "detail.xlsx"

	switch typeid {
	case 1:
		excelFilename = "detail-compare1.xlsx"
	case 2:
		excelFilename = "detail-compare2.xlsx"
	}

	switch estimateType {
	case 5:
		excelFilename = "defect.xlsx"
	case 6:
		excelFilename = "diagnosis.xlsx"
	}

	f, err := excelize.OpenFile(fmt.Sprintf("./doc/estimate/%v", excelFilename))
	if err != nil {
		log.Println(err)
		return ""
	}

	sheet := "개요"

	t := time.ParseDate(estimate.Writedate)
	writedate := strings.Split(estimate.Writedate, "-")

	complateyear := ""

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

	start := strings.Split(estimate.Start, "-")
	part := ""

	if estimateType == 2 {
		switch start[1] {
		case "0":
			part = "연간 "
		case "1":
			part = "상반기 "
		default:
			part = "하반기 "
		}
	}

	typeStr := "정밀안전점검"
	duration := 25
	name := "시설물 "
	name2 := "시설물의 "

	switch estimateType {
	case 4:
		typeStr = "하자보수"
		part = ""
	case 5:
		typeStr = "하자조사"
		part = ""
	case 6:
		typeStr = "구조안전진단"
		part = ""
		duration = 60
	case 7:
		typeStr = "감리"
		part = ""
	case 8:
		if estimate.Name != "" {
			typeStr = fmt.Sprintf("%v 기술자문", estimate.Name)
		} else {
			typeStr = "기술자문"
		}
		part = ""
		duration = 30
		name = ""
	case 9:
		typeStr = "순찰"
		part = ""
	case 10:
		typeStr = "점검프로그램 사용"
		part = ""
	}

	title := fmt.Sprintf("(%v년 %v%v)", start[0], part, typeStr)
	priceStr := fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))

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
		f.SetCellStr(sheet, "B6", fmt.Sprintf("%v %v 용역", apt.Name, typeStr))
		f.SetCellStr(sheet, "B8", apt.Address)
		f.SetCellStr(sheet, "B9", t.Humandate())
		f.SetCellStr(sheet, "B10", fmt.Sprintf("%v년", writedate[0]))

		f.SetCellStr(sheet, "B11", buildingSize)
		f.SetCellStr(sheet, "B13", complateyear)

		f.SetCellStr(sheet, "B14", tel)
		//f.SetCellStr(sheet, "B15", fax)

		f.SetCellValue(sheet, "B19", estimate.Price)
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
			f.SetCellValue(sheet, col, v.Price)
		}

		f.SetCellStr(sheet, "B5", fmt.Sprintf("%v%v 견적서", name, typeStr))
		f.SetCellStr(sheet, "B7", fmt.Sprintf("상기 금액은 %v 용역대가임.", typeStr))

		sheet = "에이앤비 갑지"

		if estimateType == 2 {
			f.SetCellStr(sheet, "D40", fmt.Sprintf("가) 시설물의 안전 및 유지관리에 관한 특별법 제11조에 의거한 건축물 %v", typeStr))
		} else {
			f.SetCellStr(sheet, "D40", fmt.Sprintf("가) %v%v 용역", name, typeStr))
			f.SetCellStr(sheet, "D41", fmt.Sprintf("나) %v 용역 보고서 1부", typeStr))
			f.SetCellStr(sheet, "D42", "")
			f.SetCellStr(sheet, "D43", "")
			f.SetCellStr(sheet, "D44", "")
			f.SetCellStr(sheet, "I46", "")
			f.SetCellStr(sheet, "G48", fmt.Sprintf("%v 보고서 납품후 10일내 지불조건", typeStr))

			f.SetCellStr(sheet, "D51", fmt.Sprintf("과업 수행 기간 : %v일", duration))
		}

		sheet = "대가산출"

		f.SetCellStr(sheet, "B1", fmt.Sprintf("%v 용역", typeStr))

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

		if estimateType == 6 {
			f.SetCellValue(sheet, "H26", estimate.Stability)
			f.SetCellValue(sheet, "H27", estimate.Earthquake)
		}

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

		if estimateType != 6 {
			f.SetCellValue(sheet, "H28", estimate.Saleprice)
		} else {
			f.SetCellValue(sheet, "H29", estimate.Saleprice)
		}

		f.SetCellStr(sheet, "D3", priceStr)

		sheet = "에이앤비 갑지"

		if estimateType == 2 {
			f.SetCellStr(sheet, "I46", title)
		}

		sheet = "계약서1"

		f.SetCellStr(sheet, "G38", priceStr)
		priceStr = fmt.Sprintf("일금 %v원정(₩%v)", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))
		f.SetCellStr(sheet, "G39", priceStr)

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

		f.SetCellStr(sheet, "G45", fmt.Sprintf("상기 금액은 %v년 %v%v 용역대가임.", start[0], part, typeStr))
		if apt.Type == "아파트" || apt.Familycount3 > 0 {
			f.SetCellStr(sheet, "F71", "공동주택")
		} else {
			f.SetCellStr(sheet, "F71", "공동주택외 건축물")
		}
		f.SetCellStr(sheet, "J72", buildingSize)

		f.SetCellStr(sheet, "B8", fmt.Sprintf("%v 용역 계약서", typeStr))
		if estimateType != 2 {
			f.SetCellStr(sheet, "G44", fmt.Sprintf("%v 보고서 1부", typeStr))
			f.SetCellStr(sheet, "E49", fmt.Sprintf("1. %v%v 표준계약조건 1부.", name2, typeStr))
			f.SetCellStr(sheet, "E50", fmt.Sprintf("2. %v%v 과업내용 1부.", name2, typeStr))
			f.SetCellStr(sheet, "E51", fmt.Sprintf("3. %v%v 대가산출 내역서 1부.", name2, typeStr))

			f.SetCellStr(sheet, "B59", fmt.Sprintf("%v%v 표준계약조건", name2, typeStr))
			f.SetCellStr(sheet, "B63", fmt.Sprintf("%v의 실시자인 ㈜에이앤비(이하 \"수주자\"라 한다)는 다음과 같이 계약을 체결한다.", typeStr))
			f.SetCellStr(sheet, "C66", fmt.Sprintf("이 계약은 \"발주자\"가 관리하고 있는 시설물에 대하여 \"수주자\"는 %v 실시하여 구조적・기능적 결함을 발견하고, 그에 대한 신속하고 적절한 조치를 취하기 위하여 구조적 안전성 및 결함의 원인 등을 조사・측정・평가하고 보수・보강 등의 방법을 제시함으로써 재해 및 재난을 예방하고 시설물의 효용증진과 공공의 안전을 확보하는데 그 목적이 있다.", global.GetJosa(typeStr, hangul.EUL_REUL)))
			f.SetCellStr(sheet, "C72", fmt.Sprintf("④ %v 대상시설물의 범위 : ", typeStr))

			f.SetCellStr(sheet, "C76", fmt.Sprintf("② 다만, 천재지변 및 부득이한 사유로 인하여 %v의 일부 또는 전부의 변경이 불가피한 경우에는 \"발주자\"와 \"수주자\"는 협의하여 그 기간을 변경할 수 있다.", typeStr))
			f.SetCellStr(sheet, "C81", fmt.Sprintf("③ 제3조 제2항에 따라 %v 업무의 일부 또는 전부의 변경이 불가피한 경우 \"발주자\"와 \"수주자\"는 협의하여 변경할 수 있다.", typeStr))
			f.SetCellStr(sheet, "C84", fmt.Sprintf("%v 대가 지급은 기획재정부가 회계예규로 정한 기술․용역 계약 일반조건 지급 요령에 의한다. ", typeStr))
			f.SetCellStr(sheet, "C90", fmt.Sprintf("② \"수주자\"는 시설물안전법 제21조의 규정에 의한 안전점검 등에 관한 지침에 의거 %v 성실하게 실시하여야 한다. ", global.GetJosa(typeStr, hangul.EUL_REUL)))
			f.SetCellStr(sheet, "C91", fmt.Sprintf("③ \"수주자\"는 %v 실시함에 있어 시설물안전법 제11조의 규정에 따라 실시한 %v 결과를 기초로 수행하여야 하며, \"발주자\"는 주요결함사항 및 보수・보강 내용 등에 대하여 \"수주자\"에게 미리 통보하여야 한다. ", global.GetJosa(typeStr, hangul.EUL_REUL), typeStr))
			f.SetCellStr(sheet, "C95", fmt.Sprintf("① \"수주자\"는 시설물안전법 제17조 제1항에 따라 \"발주자\"에게 %v 실시결과를 통보하여야 한다. ", typeStr))
			f.SetCellStr(sheet, "C97", "")
			f.SetCellStr(sheet, "C100", fmt.Sprintf("① \"수주자\"는 %v 수행함에 있어서 필요한 경우에는 시설물안전법 제10조에 따라 \"발주자\"에게 해당 시설물의 설계・시공 및 감리와 관련된 서류의 열람이나 그 사본의 교부를 요청할 수 있으며 이 경우 \"발주자\"는 이에 응하여야 한다. ", global.GetJosa(typeStr, hangul.EUL_REUL)))
			f.SetCellStr(sheet, "C104", fmt.Sprintf("① \"발주자\"는 \"수주자\"가 %v 공정하게 실시할 수 있도록 제반여건을 조성하여야 한다. ", global.GetJosa(typeStr, hangul.EUL_REUL)))
			f.SetCellStr(sheet, "C105", fmt.Sprintf("② \"수주자\"는 시설물의 상태평가・안전성평가 등에 있어서 %v 성과를 일관성 있게 비교평가하기 위하여 \"발주자\"가 소유하고 있는 진단관련장비(사다리,줄자) 등의 사용 및 운영관리자의 지원을 요청할 수 있으며 \"발주자\"는 특별한 사정이 없는 한 이에 협조하여야 한다. ", typeStr))
			f.SetCellStr(sheet, "C106", fmt.Sprintf("③ \"발주자\"는 %v 실시함에 있어서 시설물의 근접조사 및 상세조사가 가능하도록 \"수주자\"에게 협조하여야 한다.", global.GetJosa(typeStr, hangul.EUL_REUL)))
			f.SetCellStr(sheet, "C109", fmt.Sprintf("\"수주자\"가 %v 업무와 관련하여 인・허가 등의 사항이 필요한 경우에는 \"발주자\"는 %v 업무수행에 지장이 없도록 협조하여야 한다.", typeStr, typeStr))
			f.SetCellStr(sheet, "C114", fmt.Sprintf("2. \"발주자\"와 \"수주자\"가 %v 업무가 필요치 않다고 협의되었을 때", typeStr))
			f.SetCellStr(sheet, "C119", fmt.Sprintf("\"수주자\"는 %v 수행시 얻은 정보 또는 성과품 및 자료를 계약이행의 전・후를 막론하고 \"발주자\"의 사전승인 없이 외부에 누설할 수 없다.", typeStr))

			f.SetCellStr(sheet, "B125", fmt.Sprintf("%v%v 과업내용", name2, typeStr))
			f.SetCellStr(sheet, "B150", fmt.Sprintf("%v%v 대가산출 내역서", name2, typeStr))
		}

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
		title = fmt.Sprintf("%v %v년 %v건축물 %v", apt.Name, start[0], part, typeStr)
		priceStr = fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(compareestimate.Price)), humanize.Comma(int64(compareestimate.Price)))
		price1Str := fmt.Sprintf("일금 %v원정", global.HumanMoney(int64(compareestimate.Price)))
		price2Str := fmt.Sprintf("(₩%v) - VAT 별도", humanize.Comma(int64(compareestimate.Price)))

		t = time.ParseDate(compareestimate.Writedate)
		f.SetCellStr(sheet, "C8", fmt.Sprintf("ELIM-%04d%02d%02d", t.Year(), t.Month(), t.Day()))
		f.SetCellStr(sheet, "C9", t.Humandate())
		f.SetCellStr(sheet, "C11", apt.Name)
		f.SetCellStr(sheet, "C18", title)
		f.SetCellStr(sheet, "C29", title)

		if estimateType == 2 {
			f.SetCellStr(sheet, "C30", fmt.Sprintf("- 육안조사 통한 현장조사\n- 시특법에 의한 %v 실시 및 보고서 작성\n- FMS등록 제출업무", typeStr))
		} else {
			f.SetCellStr(sheet, "C30", fmt.Sprintf("- 육안조사 통한 현장조사\n- %v 실시 및 보고서 작성", typeStr))
		}

		sheet = "엘림대가내역서 갑지"

		f.SetCellStr(sheet, "B7", title)
		f.SetCellStr(sheet, "B10", buildingSize)
		f.SetCellStr(sheet, "B13", price1Str)
		f.SetCellStr(sheet, "D13", price2Str)
		f.SetCellValue(sheet, "F18", compareestimate.Price)

		f.SetCellStr(sheet, "A18", "정밀점검")

		sheet = "엘림대가내역서"

		f.SetCellStr(sheet, "B3", fmt.Sprintf("%v %v년 %v건축물 %v", apt.Name, t.Year(), part, typeStr))
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
		title = fmt.Sprintf("%v년 %v 용역", t.Year(), typeStr)

		t = time.ParseDate(compareestimate.Writedate)
		f.SetCellStr(sheet, "A6", apt.Name)
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

	return filename
}

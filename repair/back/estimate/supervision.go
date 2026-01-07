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

func Supervision(id int64, typeid int, conn *models.Connection, estimate *models.Estimate, compareestimates []models.Compareestimate, apt *models.Apt) string {
	excelFilename := "supervision.xlsx"

	switch typeid {
	case 1:
		excelFilename = "detail-compare1.xlsx"
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

	typeStr := "감리"

	if estimate.Name != "" {
		typeStr = fmt.Sprintf("%v 감리", estimate.Name)
	} else {
		typeStr = "감리"
	}

	title := fmt.Sprintf("(%v년 %v%v)", start[0], part, typeStr)
	priceStr := fmt.Sprintf("일금 %v원정(₩%v) - VAT 별도", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))

	if estimate.Days == 0 {
		estimate.Days = 1
	}

	if typeid == 0 {
		sheet = "갑지"

		no := GetEstimateNo(typeid, t, estimate.Date, conn)
		f.SetCellStr(sheet, "E6", no)
		f.SetCellStr(sheet, "E7", t.Humandate())
		f.SetCellStr(sheet, "E8", fmt.Sprintf("%v 입주자대표회의", apt.Name))
		f.SetCellStr(sheet, "E10", fmt.Sprintf("%v %v 견적건", apt.Name, typeStr))

		f.SetCellStr(sheet, "E15", fmt.Sprintf("2. 귀 아파트에서 의뢰하신 %v건에 대한 견적서를 ", typeStr))

		f.SetCellStr(sheet, "F19", fmt.Sprintf("가.조사대상: %v", apt.Name))
		f.SetCellStr(sheet, "F20", fmt.Sprintf("나.점검범위: %v %v", buildingSize, estimate.Name))
		f.SetCellStr(sheet, "F21", fmt.Sprintf("다.소 재 지: %v", apt.Address))
		f.SetCellStr(sheet, "F22", fmt.Sprintf("라.준 공 일: %v", complateyear))

		sheet = "산출내역(감리)"

		f.SetCellStr(sheet, "A5", t.Humandate())
		f.SetCellStr(sheet, "A6", apt.Name)

		f.SetCellValue(sheet, "H14", estimate.Personprice7)
		f.SetCellValue(sheet, "H15", estimate.Personprice8)
		f.SetCellValue(sheet, "H16", estimate.Personprice9)
		f.SetCellValue(sheet, "H17", estimate.Personprice10)

		f.SetCellValue(sheet, "H19", estimate.Personprice2)
		f.SetCellValue(sheet, "H20", estimate.Personprice3)
		f.SetCellValue(sheet, "H21", estimate.Personprice4)
		f.SetCellValue(sheet, "H22", estimate.Personprice5)

		f.SetCellValue(sheet, "F14", estimate.Person7)
		f.SetCellValue(sheet, "F15", estimate.Person8)
		f.SetCellValue(sheet, "F16", estimate.Person9)
		f.SetCellValue(sheet, "F17", estimate.Person10)

		outPersons := (estimate.Person7 + estimate.Person8 + estimate.Person9 + estimate.Person10) * estimate.Days
		f.SetCellValue(sheet, "F27", outPersons)
		f.SetCellValue(sheet, "F28", outPersons)

		if estimate.Person7 > 0 {
			f.SetCellValue(sheet, "D14", estimate.Days)
		}
		if estimate.Person8 > 0 {
			f.SetCellValue(sheet, "D15", estimate.Days)
		}
		if estimate.Person9 > 0 {
			f.SetCellValue(sheet, "D16", estimate.Days)
		}
		if estimate.Person10 > 0 {
			f.SetCellValue(sheet, "D17", estimate.Days)
		}

		if estimate.Person2 > 0 {
			f.SetCellValue(sheet, "D19", estimate.Person2)
			f.SetCellValue(sheet, "F19", 1)
		}
		if estimate.Person3 > 0 {
			f.SetCellValue(sheet, "D20", estimate.Person3)
			f.SetCellValue(sheet, "F20", 1)
		}
		if estimate.Person4 > 0 {
			f.SetCellValue(sheet, "D21", estimate.Person4)
			f.SetCellValue(sheet, "F21", 1)
		}
		if estimate.Person5 > 0 {
			f.SetCellValue(sheet, "D22", estimate.Person5)
			f.SetCellValue(sheet, "F22", 1)
		}

		f.SetCellFormula(sheet, "I24", fmt.Sprintf("=ROUND(I13*%v%%, 0)", estimate.Financialprice))
		f.SetCellFormula(sheet, "I25", fmt.Sprintf("=ROUND((I13+I24)*%v%%, 0)", estimate.Techprice))

		f.SetCellValue(sheet, "D24", fmt.Sprintf("직접인건비 * %v%%", estimate.Financialprice))
		f.SetCellValue(sheet, "D25", fmt.Sprintf("(직접인건비 + 제경비) * %v%%", estimate.Techprice))

		f.SetCellValue(sheet, "H27", estimate.Travelprice)
		f.SetCellValue(sheet, "H28", estimate.Carprice)

		f.SetCellFormula(sheet, "I29", fmt.Sprintf("=ROUND(I18*%v%%, 0)", estimate.Danger))
		f.SetCellValue(sheet, "D29", fmt.Sprintf("외업인건비의 %v%%", estimate.Danger))

		f.SetCellFormula(sheet, "I30", fmt.Sprintf("=ROUND(I13*%v%%, 0)", estimate.Machine))
		f.SetCellValue(sheet, "D30", fmt.Sprintf("직접인건비의 %v%%", estimate.Machine))

		f.SetCellValue(sheet, "I31", estimate.Printprice)
		f.SetCellValue(sheet, "I35", estimate.Saleprice)

		f.SetCellValue(sheet, "D11", priceStr)

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

		f.SetCellStr(sheet, "C30", fmt.Sprintf("- 육안조사 통한 현장조사\n- %v 실시 및 보고서 작성", typeStr))

		sheet = "엘림대가내역서 갑지"

		f.SetCellStr(sheet, "B7", title)
		f.SetCellStr(sheet, "B10", buildingSize)
		f.SetCellStr(sheet, "B13", price1Str)
		f.SetCellStr(sheet, "D13", price2Str)
		f.SetCellValue(sheet, "F18", compareestimate.Price)

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

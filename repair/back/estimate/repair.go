package estimate

import (
	"fmt"
	"log"
	"os"
	"repair/global"
	"repair/global/config"
	"repair/global/time"
	"repair/models"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/xuri/excelize/v2"
)

func Repair(id int64, typeid int, conn *models.Connection, estimate *models.Estimate, compareestimates []models.Compareestimate, apt *models.Apt) string {
	excelFilename := ""
	type1 := ""
	type2 := ""
	switch estimate.Subtype {
	case 1:
		excelFilename = "repair1.xlsx"
		type1 = "조정"
		type2 = "조정"
	case 2:
		excelFilename = "repair2.xlsx"
		type1 = "수립"
		type2 = "수립(조정 포함)"
	}

	log.Println("typeid", typeid)
	switch typeid {
	case 3:
		excelFilename = "repair-compare.xlsx"
	case 4:
		excelFilename = "repair-compare2.xlsx"
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	f, err := excelize.OpenFile(fmt.Sprintf("./doc/estimate/%v", excelFilename))
	if err != nil {
		log.Println(err)
		return ""
	}

	flatcount := strings.Split(apt.Flatcount, "(")

	sheet := "갑지"

	if typeid == 0 {
		t := time.ParseDate(estimate.Writedate)

		// estimateManager := models.NewEstimateManager(conn)
		// duration := time.Now().GetDurationArray()
		// count := estimateManager.Count([]interface{}{
		// 	// models.Where{Column: "type", Value: 1, Compare: "="},
		// 	models.Where{Column: "date", Value: duration[0], Compare: ">="},
		// 	models.Where{Column: "date", Value: estimate.Date, Compare: "<"},
		// })

		// no := fmt.Sprintf("ANB-%v-%v", t.DateAsOnlyNumber(), count+1)
		no := GetEstimateNo(typeid, t, estimate.Date, conn)

		str := strings.Split(estimate.Writedate, "-")
		f.SetCellStr(sheet, "E6", no)
		f.SetCellStr(sheet, "E7", fmt.Sprintf("%v년 %v월 %v일", str[0], str[1], str[2]))
		f.SetCellStr(sheet, "E8", fmt.Sprintf("%v 입주자대표회장님", apt.Name))

		if estimate.Event == 1 {
			f.SetCellStr(sheet, "E10", fmt.Sprintf("%v 장기수선계획서 %v 견적건 (이벤트 견적)", apt.Name, type2))
		} else {
			f.SetCellStr(sheet, "E10", fmt.Sprintf("%v 장기수선계획서 %v 견적건", apt.Name, type2))
		}

		f.SetCellStr(sheet, "A21", fmt.Sprintf("를 드리며, 귀 아파트에서 %v하고자 하는 장기수선계획서 작성에 관한 견적서를", type1))
		f.SetCellStr(sheet, "A22", fmt.Sprintf("아래와 같이 제출하오니 검토하시어 %v 대행업무를 위임하여 주시기 바랍니다.", type1))

		f.SetCellStr(sheet, "A37", fmt.Sprintf("첨    부 :  1. 장기수선계획서 %v 견적서 1부 끝.", type1))

		f.SetCellStr(sheet, "M25", apt.Name)

		complateyear := ""
		str = strings.Split(apt.Completeyear, "-")

		if len(str) == 3 {
			complateyear = fmt.Sprintf("%v년 %v월 %v일", str[0], str[1], str[2])
		} else if len(str) == 2 {
			complateyear = fmt.Sprintf("%v년 %v월", str[0], str[1])
		} else {
			complateyear = apt.Completeyear
		}

		f.SetCellStr(sheet, "M26", apt.Address)
		f.SetCellStr(sheet, "M27", fmt.Sprintf("아파트 %v개동 %v세대", flatcount[0], apt.Familycount))
		f.SetCellStr(sheet, "M28", complateyear)
		f.SetCellStr(sheet, "M29", apt.Tel)

		f.SetCellStr(sheet, "M30", fmt.Sprintf("장기수선계획 %v", type2))

		if estimate.Parcel == 1 {
			f.SetCellStr(sheet, "M34", "4. 장기수선계획 관련 도면, 서류 수령 및 보고서 납품은")
			f.SetCellStr(sheet, "M35", "   택배활용으로 함.")
		}

		sheet = fmt.Sprintf("산출내역( 장기수선계획 %v)", type1)

		f.SetCellStr(sheet, "I3", no)
		f.SetCellStr(sheet, "A3", fmt.Sprintf("%04d/%02d/%02d", t.Year(), t.Month(), t.Day()))

		f.SetCellStr(sheet, "A4", apt.Name)
		f.SetCellStr(sheet, "A8", fmt.Sprintf("아래와 같이 장기수선계획서 %v", type1))

		f.SetCellValue(sheet, "D13", float64(estimate.Person2))
		f.SetCellValue(sheet, "D14", float64(estimate.Person3))
		f.SetCellValue(sheet, "D15", float64(estimate.Person4))
		f.SetCellValue(sheet, "D16", float64(estimate.Person5))

		f.SetCellValue(sheet, "F13", 1)
		f.SetCellValue(sheet, "F14", 1)
		f.SetCellValue(sheet, "F15", 1)
		f.SetCellValue(sheet, "F16", 1)

		f.SetCellValue(sheet, "H13", estimate.Personprice2)
		f.SetCellValue(sheet, "H14", estimate.Personprice3)
		f.SetCellValue(sheet, "H15", estimate.Personprice4)
		f.SetCellValue(sheet, "H16", estimate.Personprice5)

		f.SetCellValue(sheet, "D18", fmt.Sprintf("직접인건비 × %v%%", estimate.Financialprice))
		f.SetCellValue(sheet, "D19", fmt.Sprintf("(직접인건비+제경비)×%v%%", estimate.Techprice))

		f.SetCellFormula(sheet, "I18", fmt.Sprintf("=ROUND(I17*%v%%, 0)", estimate.Financialprice))
		f.SetCellFormula(sheet, "I19", fmt.Sprintf("=ROUND((I17+I18)*%v%%, 0)", estimate.Techprice))
		f.SetCellValue(sheet, "I20", estimate.Directprice)
		f.SetCellValue(sheet, "I21", estimate.Printprice)
		f.SetCellValue(sheet, "I22", estimate.Extraprice)

		f.SetCellValue(sheet, "I26", estimate.Saleprice)
		f.SetCellValue(sheet, "I27", estimate.Price)

		f.SetCellValue(sheet, "F10", fmt.Sprintf("%v원정(₩%v)", global.HumanMoney(estimate.Price), humanize.Comma(int64(estimate.Price))))

		if estimate.Event == 1 {
			f.SetCellValue(sheet, "A26", "이벤트 할인 금액")
		}

		f.UpdateLinkedValue()
	} else {
		// 비교 견적 -
		compareestimate := models.Compareestimate{}
		for _, v := range compareestimates {
			if v.Comparecompany == int64(typeid) {
				compareestimate = v
			}
		}
		log.Println("COM", compareestimate.Id)
		log.Println("comparecompany", compareestimate.Comparecompany)
		log.Println("COM", compareestimate.Type)
		t := time.ParseDate(compareestimate.Writedate)
		sheet = "표지"
		f.SetCellStr(sheet, "A5", fmt.Sprintf("%v 귀하", apt.Name))
		f.SetCellStr(sheet, "A11", t.Humandate())
		f.SetCellStr(sheet, "B9", fmt.Sprintf("일금%v원정", global.HumanMoney(compareestimate.Price)))
		f.SetCellValue(sheet, "G13", compareestimate.Price)
		f.SetCellStr(sheet, "A26", fmt.Sprintf("[ 견적조건 및 특기사항 ]    %v개동 %v세대", flatcount[0], apt.Familycount))

		if estimate.Subtype == 1 {
			f.SetCellValue(sheet, "A28", "2) 장기수선계획 보고서 1권 제출")
			f.SetCellValue(sheet, "A29", "")
		}

		sheet = "내역서"
		f.SetCellValue(sheet, "F12", compareestimate.Saleprice)
		f.SetCellValue(sheet, "F10", compareestimate.Printprice)

		if compareestimate.Person2 > 0 {
			f.SetCellValue(sheet, "B4", float64(compareestimate.Person2))
			f.SetCellValue(sheet, "C4", 1)
		}
		if compareestimate.Person3 > 0 {
			f.SetCellValue(sheet, "B5", float64(compareestimate.Person3))
			f.SetCellValue(sheet, "C5", 1)
		}
		if compareestimate.Person4 > 0 {
			f.SetCellValue(sheet, "B6", float64(compareestimate.Person4))
			f.SetCellValue(sheet, "C6", 1)
		}
		if compareestimate.Person5 > 0 {
			f.SetCellValue(sheet, "B7", float64(compareestimate.Person5))
			f.SetCellValue(sheet, "C7", 1)
		}

		f.SetCellValue(sheet, "E4", compareestimate.Personprice2)
		f.SetCellValue(sheet, "E5", compareestimate.Personprice3)
		f.SetCellValue(sheet, "E6", compareestimate.Personprice4)
		f.SetCellValue(sheet, "E7", compareestimate.Personprice5)

		f.SetCellValue(sheet, "B9", fmt.Sprintf("직접인건비*%v%%", compareestimate.Financialprice))
		f.SetCellFormula(sheet, "F9", fmt.Sprintf("=F8*%v%%", compareestimate.Financialprice))
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

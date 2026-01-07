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

func Program(id int64, typeid int, conn *models.Connection, estimate *models.Estimate, compareestimates []models.Compareestimate, apt *models.Apt) string {
	excelFilename := fmt.Sprintf("program-periodic%v.xlsx", estimate.Subtype)

	f, err := excelize.OpenFile(fmt.Sprintf("./doc/estimate/%v", excelFilename))
	if err != nil {
		log.Println(err)
		return ""
	}

	sheet := "갑지"

	t := time.ParseDate(estimate.Writedate)

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

	no := GetEstimateNo(typeid, t, estimate.Date, conn)

	// buildingType := "공동주택"
	// if apt.Type == "아파트" || apt.Familycount3 > 0 {
	// 	buildingType = "공동주택"
	// } else {
	// 	buildingType = "공동주택외 건축물"
	// }

	part := ""

	switch estimate.Subtype {
	case 1:
		part = "상반기 "
	case 2:
		part = "하반기 "
	default:
		part = "연간 "
	}

	title := fmt.Sprintf("%v년 %v %v 점검 프로그램 사용 견적건", t.Year(), part, apt.Name)
	subtitle := fmt.Sprintf("   %v년 %v 정기안전점검", t.Year(), part)
	subtitle2 := fmt.Sprintf("   %v년 상반기 / 하반기 정기안전점검", t.Year())

	str := strings.Split(apt.Completeyear, "-")
	complateyear := ""
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

	humanPrice := fmt.Sprintf("%v원정(₩%v)", global.HumanMoney(estimate.Price), humanize.Comma(int64(estimate.Price)))

	f.SetCellStr(sheet, "E6", no)
	f.SetCellStr(sheet, "E7", t.Humandate())
	f.SetCellStr(sheet, "E8", fmt.Sprintf("%v 입주자대표회장님", apt.Name))
	f.SetCellStr(sheet, "E10", title)
	f.SetCellStr(sheet, "M24", apt.Name)
	f.SetCellStr(sheet, "M25", apt.Address)
	f.SetCellStr(sheet, "M26", buildingSize)
	f.SetCellStr(sheet, "M27", complateyear)
	f.SetCellStr(sheet, "M28", apt.Tel)
	switch estimate.Subtype {
	case 1, 2:
		f.SetCellStr(sheet, "M33", subtitle)
	case 3:
		f.SetCellStr(sheet, "M33", subtitle2)
	}

	sheet = "산출내역 (점검프로그램)"
	f.SetCellStr(sheet, "A4", apt.Name)
	f.SetCellStr(sheet, "F10", humanPrice)
	if estimate.Subtype == 1 || estimate.Subtype == 2 {
		f.SetCellValue(sheet, "H13", estimate.Saleprice+estimate.Price)
	} else {
		f.SetCellValue(sheet, "H13", (estimate.Saleprice+estimate.Price)/2)
	}
	f.SetCellValue(sheet, "I27", estimate.Saleprice)

	f.UpdateLinkedValue()

	filename := fmt.Sprintf("%v.xlsx", global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	e := f.SaveAs(fullFilename)
	if e != nil {
		log.Println(e)
	}
	f.Close()

	return filename
}

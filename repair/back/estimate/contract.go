package estimate

import (
	"fmt"
	"regexp"
	"repair/global"
	"repair/global/config"
	"repair/global/time"
	"repair/models"
	"strings"

	"github.com/dustin/go-humanize"
)

func Contract(id int64, typeid int, conn *models.Connection) string {
	contractManager := models.NewContractManager(conn)
	estimateManager := models.NewEstimateManager(conn)

	contract := contractManager.Get(id)
	if contract == nil {
		return ""
	}

	var estimate *models.Estimate

	estimateType := 1
	estimateSubtype := 1
	if contract.Estimate > 0 {
		estimate = estimateManager.Get(contract.Estimate)
		if estimate != nil {
			estimateType = estimate.Type
			estimateSubtype = estimate.Subtype
		}
	}

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(estimate.Apt)

	if apt == nil {
		return ""
	}

	apttype := 1
	if apt.Type == "아파트" || apt.Familycount3 > 0 {
	} else {
		apttype = 2
	}

	address := strings.Split(apt.Address, "(")
	temp := strings.Split(apt.Address2, " ")
	r, _ := regexp.Compile("^[^0-9a-zA-Z]+동$")
	dong := ""
	for i := len(temp) - 1; i >= 0; i-- {
		str := temp[i]

		find := r.FindString(str)
		if find != "" {
			dong = fmt.Sprintf("(%v)", find)
			break
		}
	}

	t := time.ParseDate(contract.Contractdate)
	if t == nil {
		t = time.Now()
	}
	startdate := time.ParseDate(contract.Contractstartdate)
	enddate := time.ParseDate(contract.Contractenddate)
	contractDate := time.ParseDate(contract.Contractdate)

	price := fmt.Sprintf("일금 %v원정(₩%v) - ", global.HumanMoney(int64(estimate.Price)), humanize.Comma(int64(estimate.Price)))
	divprice := fmt.Sprintf("일금%v원(₩%v)", global.HumanMoney(int64(estimate.Price)/2), humanize.Comma(int64(estimate.Price)/2))
	oneprice := fmt.Sprintf("%v원 - ", humanize.Comma(int64(estimate.Price/2)))

	if contract.Vat == 1 {
		price += " VAT 포함"
		oneprice += " VAT 포함"
	} else {
		price += " VAT 별도"
		oneprice += " VAT 별도"
	}

	hwpFilename := ""
	switch estimateType {
	case 1:
		hwpFilename = fmt.Sprintf("./doc/estimate/repair-contract%v-%v.hml", estimateSubtype, apttype)
	case 7:
		hwpFilename = "./doc/estimate/supervision-contract.hml"
	case 10:
		hwpFilename = fmt.Sprintf("./doc/estimate/program-periodic-contract%v.hml", estimateSubtype)
	}

	content := global.ReadFile(hwpFilename)
	content = strings.ReplaceAll(content, "{{title}}", fmt.Sprintf("%v %v 감리용역", apt.Name, estimate.Name))
	content = strings.ReplaceAll(content, "{{subtitle}}", estimate.Name)
	content = strings.ReplaceAll(content, "{{name}}", apt.Name)
	content = strings.ReplaceAll(content, "{{fulladdress}}", address[0])
	content = strings.ReplaceAll(content, "{{address}}", address[0])
	content = strings.ReplaceAll(content, "{{addressetc}}", dong)
	content = strings.ReplaceAll(content, "{{dong}}", dong)
	content = strings.ReplaceAll(content, "{{tel}}", apt.Tel)

	content = strings.ReplaceAll(content, "{{year}}", fmt.Sprintf("%v", t.Year()))
	duration := ""
	startdateStr := ""
	enddateStr := ""
	if startdate != nil && enddate != nil {
		duration = fmt.Sprintf("%v ~ %v", startdate.Humandate(), enddate.Humandate())
		startdateStr = startdate.Humandate()
		enddateStr = enddate.Humandate()
	} else if startdate != nil {
		enddateStr = ""
		duration = fmt.Sprintf("%v ~ ", startdate.Humandate())
	} else if enddate != nil {
		startdateStr = ""
		duration = fmt.Sprintf("               ~ %v", enddate.Humandate())
	}

	content = strings.ReplaceAll(content, "{{duration}}", duration)
	if contractDate != nil {
		content = strings.ReplaceAll(content, "{{contractdate}}", contractDate.Humandate())
	} else {
		content = strings.ReplaceAll(content, "{{contractdate}}", "")
	}
	content = strings.ReplaceAll(content, "{{startdate}}", startdateStr)
	content = strings.ReplaceAll(content, "{{enddate}}", enddateStr)

	content = strings.ReplaceAll(content, "{{price}}", price)
	content = strings.ReplaceAll(content, "{{divprice}}", divprice)
	content = strings.ReplaceAll(content, "{{oneprice}}", oneprice)

	filename := fmt.Sprintf("%v.hml", global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, content)

	return filename
}

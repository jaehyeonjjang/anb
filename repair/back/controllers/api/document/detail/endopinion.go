package detail

import (
	"fmt"
	"repair/models"
	"strings"
)

func Endopinion(dongs []models.Aptdong, datas []models.Periodicdata) string {
	item := Data(datas)
	areas := item.Areas(dongs)
	//part := item.Types()

	strs := make([]string, 0)

	crack := []bool{false, false, false}
	results := item.Result()
	for _, v := range results {
		switch v.Type {
		case 1:
			crack[1] = true
		case 2:
			crack[1] = true
		case 3:
			crack[2] = true
		}
	}

	if crack[1] {
		strs = append(strs, "균열(건식)")
	}

	if crack[2] {
		strs = append(strs, "균열누수흔적(습식)")
	}

	for _, v := range results {
		if v.Type > 3 {
			strs = append(strs, v.Result)
		}
	}

	//ret := fmt.Sprintf("%v의 %v에서 %v", strings.Join(areas, ", "), part, strings.Join(strs, " 및 "))
	ret := fmt.Sprintf("%v에서 %v", strings.Join(areas, ", "), strings.Join(strs, " 및 "))
	return ret
}

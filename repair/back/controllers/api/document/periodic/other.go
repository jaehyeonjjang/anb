package periodic

import (
	"fmt"
	"repair/models"
	"strings"
)

type OtherResult struct {
	Grade    string
	Items    []string
	Head     string
	Position string
	Good     int
}

func GetOtherByCategory(category int, items []models.Periodicother) []models.Periodicother {
	lists := make([]models.Periodicother, 0)

	for _, v := range items {
		if v.Category != category {
			continue
		}

		lists = append(lists, v)
	}

	return lists
}

func GetOtherByType(typeid int, items []models.Periodicother) []models.Periodicother {
	lists := make([]models.Periodicother, 0)

	for _, v := range items {
		if v.Type != typeid {
			continue
		}

		lists = append(lists, v)
	}

	return lists
}

func CheckStatus(datas []models.Periodicother) bool {
	for _, v := range datas {
		if v.Status != "" {
			return true
		}
	}

	return false
}

func GetItemByPosition(name string, items []models.Periodicother) *models.Periodicother {
	for _, v := range items {
		if v.Position == name {
			return &v
		}
	}

	return nil
}

func GetGrade(name string, items []models.Periodicother) string {
	for _, v := range items {
		if v.Position == name {
			return v.Status
		}
	}

	return ""
}

func Other(datas []models.Periodicother) map[int]OtherResult {
	ret := make(map[int]OtherResult)

	{
		pos := 10
		items := GetOtherByCategory(pos, datas)
		typeitems := GetOtherByType(2, items)
		check := CheckStatus(typeitems)

		strs := make([]string, 0)

		good := 0
		if check == false {
			strs = append(strs, "해당 대지면적 내 추락방지시설은 전반적으로 양호한 상태로 조사되었다.")
			good = 1
		} else {
			for _, v := range typeitems {
				if v.Status == "" {
					continue
				}
				strs = append(strs, fmt.Sprintf("%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", ")))
			}
		}

		ret[pos] = OtherResult{Grade: GetGrade("추락방지시설", items), Items: strs, Good: good}
	}

	{
		pos := 11
		items := GetOtherByCategory(pos, datas)
		typeitems := GetOtherByType(2, items)
		check := CheckStatus(typeitems)

		strs := make([]string, 0)

		good := 0
		if check == false {
			strs = append(strs, "해당 대지면적 내 도로포장은 전반적으로 양호한 상태로 조사되었다.")
			good = 1
		} else {
			for _, v := range typeitems {
				if v.Status == "" {
					continue
				}
				strs = append(strs, fmt.Sprintf("%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", ")))
			}
		}

		ret[pos] = OtherResult{Grade: GetGrade("도로포장", items), Items: strs, Good: good}
	}

	{
		pos := 12
		items := GetOtherByCategory(pos, datas)
		typeitems := GetOtherByType(2, items)
		check := CheckStatus(typeitems)

		strs := make([]string, 0)

		work := GetItemByPosition("시공", items)
		if work == nil {
			work = &models.Periodicother{}
		}

		good := 0
		if work.Status == "미시공" {
			strs = append(strs, "해당 대지면적 내 시공되어있지 않아 조사에서 제외하였다.")
			good = 2
		} else {
			if check == false {
				if work.Status == "" {
					strs = append(strs, "신축 이음부는 전반적으로 양호한 상태로 조사되었다.")
				} else {
					strs = append(strs, fmt.Sprintf("신축 이음부는 %v로 시공되어 있으며 전반적으로 양호한 상태로 조사되었다.", work.Status))
				}
				good = 1
			} else {
				for _, v := range typeitems {
					if v.Status == "" {
						continue
					}

					strs = append(strs, fmt.Sprintf("%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", ")))

				}
			}
		}

		ret[pos] = OtherResult{Grade: GetGrade("도로부 신축 이음부", items), Items: strs, Good: good}
	}

	{
		pos := 13
		items := GetOtherByCategory(pos, datas)
		typeitems := GetOtherByType(2, items)
		check := CheckStatus(typeitems)

		strs := make([]string, 0)

		work := GetItemByPosition("환기구", items)
		if work == nil {
			work = &models.Periodicother{}
		}

		good := 0
		if check == false {
			str := ""
			if work.Status != "" {
				str = fmt.Sprintf("(%v) ", strings.ReplaceAll(work.Status, ",", ", "))
			}
			strs = append(strs, fmt.Sprintf("해당 건축물 내 %v환기구의 상태는 전반적으로 양호한 상태로 조사되었다.", str))
			good = 1
		} else {
			for _, v := range typeitems {
				if v.Status == "" {
					continue
				}

				strs = append(strs, fmt.Sprintf("%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", ")))

			}
		}

		ret[pos] = OtherResult{Grade: GetGrade("환기구 등의 덮개", items), Items: strs, Good: good}
	}

	{
		pos := 14
		items := GetOtherByCategory(pos, datas)
		typeitems := GetOtherByType(2, items)
		check := CheckStatus(typeitems)

		strs := make([]string, 0)

		work := GetItemByPosition("건물외부 벽체", items)
		if work == nil {
			work = &models.Periodicother{}
		}

		head := ""
		position := ""
		good := 0
		if check == false {
			str := ""
			if work.Status != "" {
				position = fmt.Sprintf("(%v 등)", strings.ReplaceAll(work.Status, ",", ", "))
				str = fmt.Sprintf("%v ", position)
			}
			head = fmt.Sprintf("건물외부 벽체 %v마감부위 상태는 양호한 것으로 확인됨.", str)
			good = 1
		} else {
			for _, v := range typeitems {
				if v.Order == 141 {
					// strs = append(strs, " ")
					continue
				}

				if v.Status == "" {
					continue
				}

				strs = append(strs, fmt.Sprintf("%v - %v", v.Position, strings.ReplaceAll(v.Status, ",", ", ")))
			}

			if len(strs) == 0 {
				str := ""
				if work.Status != "" {
					position = fmt.Sprintf("(%v 등)", strings.ReplaceAll(work.Status, ",", ", "))
					str = fmt.Sprintf("%v ", position)
				}
				head = fmt.Sprintf("건물외부 벽체 %v마감부위 상태는 양호한 것으로 확인됨.", str)
				good = 1
			}
		}

		ret[pos] = OtherResult{Grade: GetGrade("외벽 마감재", items), Items: strs, Head: head, Position: position, Good: good}
	}

	return ret
}

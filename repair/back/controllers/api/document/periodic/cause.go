package periodic

import (
	"fmt"
	"repair/global"
	"repair/models"
	"strings"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
)

func GetOuterwallMaterial(otherMap map[int][]models.Periodicother) string {
	for _, v := range otherMap[14] {
		if v.Order == 141 {
			if v.Status != "" {
				return strings.ReplaceAll(v.Status, ",", ", ")
			}
		}
	}

	return ""
}
func GetCause(periodicdatas []models.Periodicdata, dongs []models.Aptdong, others map[int]OtherResult, otherMap map[int][]models.Periodicother) ([][]string, map[int]OtherResult) {
	titles := []string{"지붕", "지상층", "지하층", "외부벽체"}

	for _, v := range dongs {
		if v.Topcount > 0 {
			titles[0] = "지붕,옥탑층"
			break
		}
	}

	floorCount := 7
	summarys := make([]DataSummary, floorCount)

	for _, v := range periodicdatas {
		blueprint := v.Extra["blueprint"].(models.Blueprint)

		position := 3
		if blueprint.Floortype == 4 || blueprint.Floortype == 5 {
			position = 1
		} else if blueprint.Floortype == 1 || blueprint.Floortype == 2 {
			position = 5
		}

		if v.Type == 2 {
			position++
		}

		if v.Part == "외부" {
			position = 7
		}

		summarys[position-1].Data = append(summarys[position-1].Data, v)
	}

	datas := make([]DataItem, 0)

	for i := 0; i < floorCount; i++ {
		di := Data(summarys[i].Data)
		datas = append(datas, di)
	}

	items := make([][]string, 5)
	for i := 0; i < 5; i++ {
		items[i] = make([]string, 0)
	}

	first := datas[0]
	second := datas[1]

	if first.IsValid() || second.IsValid() {
		str := ""
		cause := ""

		if first.IsValid() && second.IsValid() {
			parts := strings.Join(global.UniqueStringWithoutSort(append(first.Part, second.Part...)), ", ")
			head := fmt.Sprintf("%v(%v 등) 점검결과", titles[0], parts)
			str = fmt.Sprintf("%v 벽체(기둥포함)에서 %v 조사되었고, 슬래브(보 포함)에서 %v 조사되었다.", head, global.GetJosa(first.Results(), hangul.I_GA), global.GetJosa(second.Results(), hangul.I_GA))

			cause = strings.Join(global.UniqueStringWithoutSort(append(first.Cause, second.Cause...)), ", ")
		} else if first.IsValid() {
			parts := first.Parts()
			head := fmt.Sprintf("%v(%v 등) 점검결과", titles[0], parts)
			str = fmt.Sprintf("%v 벽체(기둥포함)에서 %v 조사되었다", head, global.GetJosa(first.Results(), hangul.I_GA))

			cause = first.Causes()
		} else {
			parts := second.Parts()
			head := fmt.Sprintf("%v(%v 등) 점검결과", titles[0], parts)
			str = fmt.Sprintf("%v 슬래브(보 포함)에서 %v 조사되었다", head, global.GetJosa(second.Results(), hangul.I_GA))

			cause = second.Causes()
		}

		items[0] = append(items[0], str)
		str = fmt.Sprintf("조사된 결함은 %v에 의하여 발생한 결함으로, 구조적인 문제는 없는 것으로 보이나 장기간 방치시 내구성 저하의 원인이 되므로 적절한 보수가 요구된다.", cause)
		items[0] = append(items[0], str)
	} else {
		items[0] = append(items[0], fmt.Sprintf("%v 점검결과 대체적으로 양호한 것으로 확인되었다.", titles[0]))
	}

	first = datas[2]
	second = datas[3]

	if first.IsValid() || second.IsValid() {
		str := ""
		cause := ""

		if first.IsValid() && second.IsValid() {
			parts := strings.Join(global.UniqueStringWithoutSort(append(first.Part, second.Part...)), ", ")
			head := fmt.Sprintf("지상층의 경우 전용부분은 육안조사에 제한이 있어 공용(%v) 부위를 중점으로 조사하였고, 조사결과", parts)
			str = fmt.Sprintf("%v 벽체(기둥포함)에서 %v 조사되었고, 슬래브(보 포함)에서 %v 조사되었다.", head, global.GetJosa(first.Results(), hangul.I_GA), global.GetJosa(second.Results(), hangul.I_GA))

			cause = strings.Join(global.UniqueStringWithoutSort(append(first.Cause, second.Cause...)), ", ")
		} else if first.IsValid() {
			parts := first.Parts()
			head := fmt.Sprintf("지상층의 경우 전용부분은 육안조사에 제한이 있어 공용(%v) 부위를 중점으로 조사하였고, 조사결과", parts)
			str = fmt.Sprintf("%v 벽체(기둥포함)에서 %v 조사되었다", head, global.GetJosa(first.Results(), hangul.I_GA))

			cause = first.Causes()
		} else {
			parts := second.Parts()
			head := fmt.Sprintf("지상층의 경우 전용부분은 육안조사에 제한이 있어 공용(%v) 부위를 중점으로 조사하였고, 조사결과", parts)
			str = fmt.Sprintf("%v 슬래브(보 포함)에서 %v 조사되었다", head, global.GetJosa(second.Results(), hangul.I_GA))

			cause = second.Causes()
		}

		items[1] = append(items[1], str)
		str = fmt.Sprintf("조사된 결함은 %v에 의하여 발생한 것으로 판단되며, 해당 결함을 장기간 방치시 내구성저하의 원인이 되므로 관리주체는 적절한 보수조치가 필요할 것으로 사료된다.", cause)
		items[1] = append(items[1], str)
	} else {
		items[1] = append(items[1], fmt.Sprintf("%v 점검결과 대체적으로 양호한 것으로 확인되었다.", titles[1]))
	}

	first = datas[4]
	second = datas[5]

	if first.IsValid() || second.IsValid() {
		str := ""
		cause := ""

		firstResult := first.Results()
		secondResult := second.Results()

		if first.IsValid() && second.IsValid() {
			parts := strings.Join(global.UniqueStringWithoutSort(append(first.Part, second.Part...)), ", ")
			head := fmt.Sprintf("%v(%v 등) 점검결과", titles[2], parts)
			str = fmt.Sprintf("%v 벽체(기둥포함)에서 %v 조사되었고, 슬래브(보 포함)에서 %v 조사되었다.", head, global.GetJosa(firstResult, hangul.I_GA), global.GetJosa(secondResult, hangul.I_GA))

			cause = strings.Join(global.UniqueStringWithoutSort(append(first.Cause, second.Cause...)), ", ")
		} else if first.IsValid() {
			parts := first.Parts()
			head := fmt.Sprintf("%v(%v 등) 점검결과", titles[2], parts)
			str = fmt.Sprintf("%v 벽체(기둥포함)에서 %v 조사되었다", head, global.GetJosa(firstResult, hangul.I_GA))

			cause = first.Causes()
		} else {
			parts := second.Parts()
			head := fmt.Sprintf("%v(%v 등) 점검결과", titles[2], parts)
			str = fmt.Sprintf("%v 슬래브(보 포함)에서 %v 조사되었다", head, global.GetJosa(secondResult, hangul.I_GA))

			cause = second.Causes()
		}

		items[2] = append(items[2], str)

		if strings.Contains(firstResult, "누수") || strings.Contains(secondResult, "누수") {
			str = fmt.Sprintf("조사된 결함은 %v 등으로 인해 발생된 것으로 판단되며, 특히 균열과 누수(흔적)를 동반한 결함부위는 침투된 수분으로 인하여 철근을 부식시키고, 콘크리트 탄산화를 촉진시킬 우려가 있으므로 관리주체는 현재 누수가 발생되는 구조체에 대하여 적극적인 보수를 진행해야 할 것으로 판단한다.", cause)
		} else {
			str = fmt.Sprintf("조사된 결함은 %v 등으로 인해 발생된 것으로 판단되며, 구조적인 문제는 없는 것으로 보이나 장기간 방치시 내구성 저하의 원인이 되므로 적절한 보수가 요구된다.", cause)
		}

		items[2] = append(items[2], str)
	} else {
		items[2] = append(items[2], fmt.Sprintf("%v 점검결과 대체적으로 양호한 것으로 확인되었다.", titles[2]))
	}

	first = datas[6]
	other := others[14]

	outerwallMaterial := GetOuterwallMaterial(otherMap)
	if first.IsValid() || len(other.Items) > 0 {
		str := ""

		if outerwallMaterial != "" {
			str += fmt.Sprintf("외부벽체는 %v이고 ", outerwallMaterial)
		}

		causes := first.Cause

		if len(other.Items) == 0 {
			str += fmt.Sprintf("%v 조사되었다", global.GetJosa(first.Results(), hangul.I_GA))
		} else {
			if len(first.Result()) > 0 {
				str += fmt.Sprintf("%v 조사되었고, 또한 ", global.GetJosa(first.Results(), hangul.I_GA))
			}

			for i, v := range other.Items {
				if i == len(other.Items)-1 {
					str += fmt.Sprintf("%v 조사되었다.", global.GetJosa(v, hangul.I_GA))
				} else {
					str += fmt.Sprintf("%v 조사되었고, ", global.GetJosa(v, hangul.I_GA))
				}

				if strings.Contains(v, "백화현상") {
					causes = append(causes, "결함부위 누수")
				}

			}

			causes = append(causes, "마감재 파손")
		}

		items[3] = append(items[3], str)

		cause := strings.Join(global.UniqueStringWithoutSort(causes), ", ")
		str = fmt.Sprintf("조사된 결함은 %v 등으로 인해 발생된 것으로 판단되며,", cause)

		resultFlag := false
		if len(other.Items) > 0 && other.Position != "" {
			resultFlag = true
			str += fmt.Sprintf(" 또한 %v 도 결함사항에 대하여 적절한 보수가 이루어져야 할 것으로 판단된다.", strings.ReplaceAll(strings.ReplaceAll(other.Position, "(", ""), ")", ""))
		}

		if strings.Contains(cause, "누수") || strings.Contains(cause, "누수") {
			resultFlag = true
			str += " 특히 균열과 누수(흔적)를 동반한 결함부위는 침투된 수분으로 인해 발생된 것으로 판단되며, 해당 결함의 경우 장기간 방치 시 결함 확대 및 전유세대 누수로 이어지므로 적절한 보수조치가 필요한 것으로 사료된다. 또한, 향후 외벽 도장공사 진행시 해당결함이 재 발생 되지 않도록 철저한 관리감독이 이루어져야 할 것이다."
		} else {
			if len(other.Items) == 0 {
				resultFlag = true
				str += fmt.Sprintf(" 구조적인 문제는 없는 것으로 보이나 장기간 방치시 내구성 저하의 원인이 되므로 적절한 보수가 요구된다.")
			}
		}

		if resultFlag == false {
			str += fmt.Sprintf(" 결함사항에 대하여 적절한 보수가 이루어져야 할 것으로 판단된다.")
		}

		items[3] = append(items[3], str)
	} else {
		if outerwallMaterial == "" {
			items[3] = append(items[3], fmt.Sprintf("외부벽체 %v 점검결과 %v으로 마감되어 있으며 상태는 양호한 것으로 조사되었다.", other.Position, strings.ReplaceAll(strings.ReplaceAll(other.Position, "(", ""), ")", "")))
		} else {
			items[3] = append(items[3], fmt.Sprintf("외부벽체는 %v이고 상태는 양호한 것으로 조사되었다.", outerwallMaterial))
		}
	}

	{
		items := make([]string, 0)
		if len(first.Result()) > 0 {
			add := fmt.Sprintf("%v 마감부위에서 %v 등이 확인됨", outerwallMaterial, first.Results())

			items = append(items, add)

			other.Good = 1
		} else {
			if other.Head != "" {
				items = append(items, other.Head)
			}
		}

		items = append(items, other.Items...)
		other.Items = items
		others[14] = other
	}

	return items, others
}

package periodic

import (
	"fmt"
	"repair/global"
	"repair/models"
	"strings"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
)

type DataResult struct {
	Type   int
	Result string
	Method []string
}

type DataItem struct {
	Part     []string
	Cause    []string
	MaxWidth float64
	MinWidth float64
	Height   int

	Group    []bool
	Position []bool
	Check    []bool
	Type     []bool

	Width2   bool
	Width3   bool
	Maxwidth float32
	Minwidth float32

	Crack1 bool
	Crack2 bool
	Crack3 bool
	Crack4 bool
	Crack5 bool

	Othercrack1 bool
	Othercrack2 bool

	Leak1 bool
	Leak2 bool
	Leak3 bool

	Concrete1 bool
	Concrete2 bool
	Concrete3 bool

	Rebar1 bool
	Rebar2 bool

	Material1 bool
	Material2 bool
}

func Data(datas []models.Periodicdata) DataItem {
	groups := []bool{false, false, false, false, false}
	positions := []bool{false, false, false, false, false, false, false, false}
	check := []bool{false, false, false, false, false, false}
	types := []bool{false, false, false}

	width2 := false
	width3 := false
	maxWidth := 0.0
	minWidth := 9999.0

	crack1 := false
	crack2 := false
	crack3 := false
	crack4 := false
	crack5 := false

	othercrack1 := false
	othercrack2 := false

	leak1 := false
	leak2 := false
	leak3 := false

	concrete1 := false
	concrete2 := false
	concrete3 := false

	rebar1 := false
	rebar2 := false

	material1 := false
	material2 := false

	parts := make([]string, 0)
	causes := make([]string, 0)

	for _, data := range datas {
		if data.Type != 1 && data.Type != 2 {
			continue
		}

		blueprint := data.Extra["blueprint"].(models.Blueprint)
		position := 3
		group := 2

		if blueprint.Floortype == 4 || blueprint.Floortype == 5 {
			position = 1
			group = 1
		} else if blueprint.Floortype == 1 || blueprint.Floortype == 2 {
			position = 5
			group = 3
		}

		if data.Type == 2 {
			position++
		}

		if data.Part == "외부" {
			position = 7
			group = 4
		}

		positions[position] = true
		groups[group] = true
		types[data.Type] = true

		if strings.Contains(data.Shape, "조적") {
			othercrack1 = true
			check[1] = true
		}

		if strings.Contains(data.Shape, "이질재") {
			othercrack2 = true
			check[1] = true
		}

		if strings.Contains(data.Shape, "마감재") {
			if strings.Contains(data.Shape, "오염") {
				material1 = true
			}

			if strings.Contains(data.Shape, "파손") {
				material2 = true
			}

			check[5] = true
		}

		if strings.Contains(data.Shape, "콘크리트") {
			if strings.Contains(data.Shape, "박리") {
				concrete1 = true
			}

			if strings.Contains(data.Shape, "박락") {
				concrete2 = true
			}

			if strings.Contains(data.Shape, "건조수축") {
				concrete3 = true
			}

			check[3] = true
		}

		if strings.Contains(data.Shape, "철근") {
			if strings.Contains(data.Shape, "노출") {
				rebar1 = true
			}

			if strings.Contains(data.Shape, "부식") {
				rebar2 = true
			}

			check[4] = true
		}

		if strings.Contains(data.Shape, "균열") {
			width := global.Atof(data.Width)
			if width >= 0.3 {
				width2 = true
			} else {
				width3 = true
			}

			if width > maxWidth {
				maxWidth = width
			}

			if width < minWidth {
				minWidth = width
			}

			if strings.Contains(data.Shape, "수직") {
				crack1 = true
			}

			if strings.Contains(data.Shape, "수평") {
				crack2 = true
			}

			if strings.Contains(data.Shape, "경사") {
				crack3 = true
			}

			if strings.Contains(data.Shape, "층간") {
				crack4 = true
			}

			if strings.Contains(data.Shape, "망상") {
				crack5 = true
			}

			check[0] = true
		}

		if strings.Contains(data.Shape, "누수") {
			strs := strings.Split(data.Shape, "/")
			for _, name := range strs {
				if name == "누수" {
					leak1 = true
				} else if name == "누수흔적" {
					leak2 = true
				} else if name == "백태" {
					leak3 = true
				}
			}

			check[2] = true
		}

		parts = append(parts, data.Part)
		causes = append(causes, data.Remark)
	}

	var item DataItem
	item.Part = global.UniqueStringWithoutSort(parts)
	item.Cause = global.UniqueStringWithoutSort(causes)
	item.MaxWidth = maxWidth
	item.MinWidth = minWidth

	item.Group = groups
	item.Position = positions
	item.Type = types
	item.Check = check

	item.Width2 = width2
	item.Width3 = width3

	item.Crack1 = crack1
	item.Crack2 = crack2
	item.Crack3 = crack3
	item.Crack4 = crack4
	item.Crack5 = crack5

	item.Othercrack1 = othercrack1
	item.Othercrack2 = othercrack2

	item.Leak1 = leak1
	item.Leak2 = leak2
	item.Leak3 = leak3

	item.Concrete1 = concrete1
	item.Concrete2 = concrete2
	item.Concrete3 = concrete3

	item.Rebar1 = rebar1
	item.Rebar2 = rebar2

	item.Material1 = material1
	item.Material2 = material2

	return item
}

func (c *DataItem) Parts() string {
	return strings.Join(c.Part, ", ")
}

func (c *DataItem) Results() string {
	items := make([]string, 0)
	result := c.Result()
	for _, v := range result {
		items = append(items, v.Result)
	}
	return strings.Join(items, ", ")
}

func (c *DataItem) ResultsForResult() string {
	items := make([]string, 0)
	result := c.Result()
	for _, v := range result {
		if strings.Contains(v.Result, "균열(") {
			items = append(items, "균열")
		} else {
			items = append(items, v.Result)
		}
	}

	items = global.UniqueStringWithoutSort(items)

	if len(items) > 2 {
		return fmt.Sprintf("%v 및 %v %v", items[0], global.GetJosa(items[1], hangul.GWA_WA), strings.Join(items[2:], ", "))
	} else if len(items) == 2 {
		return fmt.Sprintf("%v 및 %v", items[0], items[1])
	} else if len(items) == 1 {
		return items[0]
	} else {
		return ""
	}
}

func (c *DataItem) Causes() string {
	return strings.Join(c.Cause, ", ")
}

func (c *DataItem) Methods() string {
	items := make([]string, 0)
	result := c.Result()
	for _, v := range result {
		for _, v2 := range v.Method {
			items = append(items, v2)
		}
	}

	return strings.Join(global.UniqueStringWithoutSort(items), ", ")
}

func (c *DataItem) Width() string {
	width := ""
	if c.MinWidth == c.MaxWidth {
		width = fmt.Sprintf("%vmm ", c.MinWidth)
	} else if c.MinWidth < 9000 {
		if c.MaxWidth > 0 {
			width = fmt.Sprintf("%v ~ %vmm ", c.MinWidth, c.MaxWidth)
		} else {
			width = fmt.Sprintf("%vmm ", c.MinWidth)
		}
	} else {
		if c.MaxWidth > 0 {
			width = fmt.Sprintf("%vmm ", c.MaxWidth)
		}
	}

	return width
}

func (c *DataItem) Result() []DataResult {

	items := make([]DataResult, 0)

	if c.Check[0] == true {
		results := make([]string, 0)

		if c.Crack1 == true {
			results = append(results, "수직")
		}

		if c.Crack2 == true {
			results = append(results, "수평")
		}

		if c.Crack3 == true {
			results = append(results, "경사")
		}

		if c.Crack4 == true {
			results = append(results, "층간")
		}

		if c.Crack5 == true {
			results = append(results, "망상")
		}

		title := fmt.Sprintf("균열(%v)", strings.Join(results, "/"))
		if len(results) == 0 {
			title = "균열"
		}
		result := DataResult{Type: 1, Result: title, Method: make([]string, 0)}

		if c.Width2 == true {
			result.Method = append(result.Method, "건식주입 공법(0.3mm 이상 균열)")
		}

		if c.Width3 == true {
			result.Method = append(result.Method, "표면처리 공법(0.3mm 미만 균열)")
		}

		items = append(items, result)
	}

	if c.Check[1] == true {
		results := make([]string, 0)

		if c.Othercrack1 == true {
			results = append(results, "조적")
		}

		if c.Othercrack2 == true {
			results = append(results, "이질재")
		}

		title := fmt.Sprintf("%v균열", strings.Join(results, "/"))
		result := DataResult{Type: 2, Result: title, Method: make([]string, 0)}

		result.Method = append(result.Method, "충전식 공법")

		items = append(items, result)
	}

	if c.Check[2] == true {
		results := make([]string, 0)

		results = append(results, "균열")
		if c.Leak1 == true {
			results = append(results, "누수")
		} else if c.Leak2 == true {
			results = append(results, "누수흔적")
		}

		if c.Leak3 == true {
			results = append(results, "백태")
		}

		title := strings.Join(results, "/")
		result := DataResult{Type: 3, Result: title, Method: make([]string, 0)}

		result.Method = append(result.Method, "습식주입 공법")

		items = append(items, result)
	}

	if c.Check[3] == true {
		results := make([]string, 0)

		if c.Concrete1 == true {
			results = append(results, "박리")
		}

		if c.Concrete2 == true {
			results = append(results, "박락")
		}

		if c.Concrete3 == true {
			results = append(results, "건조수축")
		}

		title := fmt.Sprintf("콘크리트 %v", strings.Join(results, "/"))
		result := DataResult{Type: 4, Result: title, Method: make([]string, 0)}

		result.Method = append(result.Method, "단면복구 공법")

		items = append(items, result)
	}

	if c.Check[4] == true {
		results := make([]string, 0)

		if c.Rebar1 == true {
			results = append(results, "노출")
		}

		if c.Rebar2 == true {
			results = append(results, "부식")
		}

		title := fmt.Sprintf("철근 %v", strings.Join(results, "/"))
		result := DataResult{Type: 5, Result: title, Method: make([]string, 0)}

		result.Method = append(result.Method, "철근노출 보수공법")

		items = append(items, result)
	}

	if c.Check[5] == true {
		results := make([]string, 0)

		if c.Material1 == true {
			results = append(results, "오염")
		}

		if c.Material2 == true {
			results = append(results, "파손")
		}

		title := fmt.Sprintf("마감재 %v", strings.Join(results, "/"))
		result := DataResult{Type: 6, Result: title, Method: make([]string, 0)}

		result.Method = append(result.Method, "외부 우수유입 및 충격")

		items = append(items, result)
	}

	return items
}

func (c *DataItem) Areas(dongs []models.Aptdong) []string {
	topcount := 0
	for _, dong := range dongs {
		if dong.Topcount > 0 {
			topcount = dong.Topcount
		}
	}

	titles := []string{"", "지붕층", "지상층", "지하층", "외부벽체"}
	if topcount > 0 {
		titles[1] = "지붕층·옥탑층"
	}

	items := make([]string, 0)

	if c.Position[1] == true || c.Position[2] == true {
		if c.Position[1] != true {
			items = append(items, fmt.Sprintf("%v의 슬래브(보 포함)", titles[1]))
		} else if c.Position[2] != true {
			items = append(items, fmt.Sprintf("%v의 벽체(기둥 포함)", titles[1]))
		} else {
			items = append(items, fmt.Sprintf("%v의 벽체(기둥 포함) 및 슬래브(보 포함)", titles[1]))
		}
	}

	if c.Position[3] == true || c.Position[4] == true {
		if c.Position[3] != true {
			items = append(items, fmt.Sprintf("%v의 슬래브(보 포함)", titles[2]))
		} else if c.Position[4] != true {
			items = append(items, fmt.Sprintf("%v의 벽체(기둥 포함)", titles[2]))
		} else {
			items = append(items, fmt.Sprintf("%v의 벽체(기둥 포함) 및 슬래브(보 포함)", titles[2]))
		}
	}

	if c.Position[5] == true || c.Position[6] == true {
		if c.Position[5] != true {
			items = append(items, fmt.Sprintf("%v의 슬래브(보 포함)", titles[3]))
		} else if c.Position[6] != true {
			items = append(items, fmt.Sprintf("%v의 벽체(기둥 포함)", titles[3]))
		} else {
			items = append(items, fmt.Sprintf("%v의 벽체(기둥 포함) 및 슬래브(보 포함)", titles[3]))
		}
	}

	if c.Position[7] == true {
		items = append(items, fmt.Sprintf("외벽의 벽체"))
	}

	return items
}

func (c *DataItem) Types() string {
	ret := ""
	if c.Type[1] && c.Type[2] {
		ret = "벽체(기둥 포함), 슬래브(보 포함)"
	} else if c.Type[1] {
		ret = "벽체(기둥 포함)"
	} else if c.Type[2] {
		ret = "슬래브(보 포함)"
	}

	return ret
}

func (c *DataItem) IsValid() bool {
	if len(c.Part) > 0 && len(c.Cause) > 0 {
		for _, v := range c.Check {
			if v == true {
				return true
			}
		}
	}

	return false
}

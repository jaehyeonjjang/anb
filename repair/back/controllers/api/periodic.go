package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"repair/controllers"
	"repair/global"
	"repair/global/config"
	"repair/models"
	"strings"
	"time"
)

type PeriodicController struct {
	controllers.Controller
}

func (c *PeriodicController) Pre_Insert(item *models.Periodic) {
	if item.Category == 1 {
		item.Result1 = 2
		item.Result2 = 2
		item.Result3 = 2
		item.Result4 = 2
		item.Result5 = 2

		item.Resulttext1 = "※ 이전 점검일 이후 건축물관리대장 열람 및 관리주체 청문, 현장 조사결과 없는 것으로 조사됨."
		item.Resulttext2 = "※ 건축물관리대장 열람 및 관리주체 청문, 현장 조사결과 없는 것으로 조사됨."
		item.Resulttext3 = "※ 관리주체 청문, 현장 조사결과 없는 것으로 조사됨."
		item.Resulttext4 = "※ 건축물관리대장 열람 및 관리주체 청문, 현장 조사결과 없는 것으로 조사됨."
		item.Resulttext5 = "※ 시설물관리대장의 보수·보강이력 사항 확인 및 청문결과 없는 것으로 조사됨."
		item.Past = ""
	} else {
		if c.Session != nil {
			item.User = c.Session.Id
		}
	}
}

func (c *PeriodicController) Post_Insert(item *models.Periodic) {
	if item.Category != 1 {
		return
	}

	conn := c.NewConnection()

	aptperiodicManager := models.NewAptperiodicManager(conn)
	periodicincidentalManager := models.NewPeriodicincidentalManager(conn)
	periodicouterwallManager := models.NewPeriodicouterwallManager(conn)
	periodicopinionManager := models.NewPeriodicopinionManager(conn)
	periodicotherManager := models.NewPeriodicotherManager(conn)
	periodicotheretcManager := models.NewPeriodicotheretcManager(conn)
	periodickeepManager := models.NewPeriodickeepManager(conn)

	aptperiodicItem := aptperiodicManager.Get(item.Apt)
	if aptperiodicItem == nil || aptperiodicItem.Id == 0 {
		var aptperiodic models.Aptperiodic
		aptperiodic.Id = item.Apt
		aptperiodicManager.Insert(&aptperiodic)
	}

	var periodicincidental models.Periodicincidental
	periodicincidental.Periodic = item.Id
	periodicincidental.Result1 = "양호"
	periodicincidental.Result2 = "없음"
	periodicincidental.Result3 = "없음"
	periodicincidental.Result4 = "없음"
	periodicincidental.Result5 = "양호"
	periodicincidental.Result6 = "없음"
	periodicincidental.Result7 = "없음"
	periodicincidental.Result8 = "없음"
	periodicincidental.Result9 = "양호"
	periodicincidental.Result10 = "없음"
	periodicincidental.Result11 = "양호"
	periodicincidental.Result12 = "없음"
	periodicincidental.Result13 = "없음"
	periodicincidental.Result14 = "없음"
	periodicincidental.Result15 = "없음"
	periodicincidental.Result16 = "없음"
	periodicincidental.Result17 = "없음"
	periodicincidental.Result18 = "없음"
	periodicincidental.Result19 = "없음"
	periodicincidental.Result20 = "양호"
	periodicincidental.Result21 = "양호"
	periodicincidentalManager.Insert(&periodicincidental)

	var periodicouterwall models.Periodicouterwall
	periodicouterwall.Periodic = item.Id
	periodicouterwall.Result1 = "양호"
	periodicouterwall.Result2 = "양호"
	periodicouterwall.Result3 = "양호"
	periodicouterwall.Result4 = "양호"
	periodicouterwall.Result5 = "양호"
	periodicouterwall.Result6 = "양호"
	periodicouterwallManager.Insert(&periodicouterwall)

	var periodicopinion models.Periodicopinion
	periodicopinion.Periodic = item.Id
	periodicopinionManager.Insert(&periodicopinion)

	var periodicotheretc models.Periodicotheretc
	periodicotheretc.Periodic = item.Id
	periodicotheretcManager.Insert(&periodicotheretc)

	titles1 := []string{
		"균열발생 상태",
		"붙임모르타르 상태",
		"연결철물 시공 상태",
		"균열방지 조치 상태",
		"기울기 및 배부름",
	}

	/*
		titles2 := []string{
			"지지구조 철물 및 연결재의 규격",
			"지지구조 철물 및 연결재의 용접 및 볼트상태",
			"지지구조 철물 및 연결재의 탈락 유무",
			"지지구조 철물 및 연결재의 부식 유무",
			"걸침턱, 추락방지시설의 상태 및 유무",
			"환기구 덮개의 처짐 및 변형 유무",
		}
	*/

	titles3 := []string{
		"바닥포장부위 침하 및 균열현상",
		"건물전체의 부등침하 현상",
		"외부 옹벽(축대)의 균열 현상",
		"건물주변 토량 침하현상",
		"하수관로 및 맨홀의 배수, 청소 상태",
		"외벽의 전도 위험부위",
		"외벽 모르터 또는 콘크리트의 탈락부위",
		"외벽 창문 유리의 파손",
		"ROOF DRAIN의 상태",
		"옥상에 하중(물건)의 과재 여부",
		"내부 창, 문의 작동상태",
		"건물 내부의 진동 여부",
		"천정재(텍스류)의 탈락 및 갈라짐 상태",
		"벽지 및 천정지가 찢어진 곳 유무",
		"실내의 하중(물건)의 과적 여부",
		"건물에서 뚝뚝하는 소리",
		"코킹이 갑자기 떨어진 곳의 유무",
		"담장의 전도 징후",
		"돌출물(간판, 안테나 등)의 탈락현상",
		"지하수 배수펌프 작동 상태",
		"안전난간의 견고성",
	}

	types := []int{1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1}

	id := item.Id

	for i, v := range titles1 {
		var item models.Periodicother
		item.Name = v
		item.Type = 1
		item.Category = 1
		item.Order = i + 1
		item.Periodic = id

		periodicotherManager.Insert(&item)
	}

	/*
		for i, v := range titles2 {
			var item models.Periodicother
			item.Name = v
			item.Type = 1
			item.Category = 2
			item.Order = i + 1 + 100
			item.Periodic = id

			periodicotherManager.Insert(&item)
		}
	*/

	for i, v := range titles3 {
		var item models.Periodicother
		item.Name = v
		item.Type = types[i]
		item.Category = 3
		item.Order = i + 1 + 200
		item.Periodic = id

		periodicotherManager.Insert(&item)
	}

	periodicotherManager.Insert(&models.Periodicother{Name: "a,b,c,d,e", Type: 1, Position: "추락방지시설", Category: 10, Order: 100, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "전도,변형,고정부 및 연결부 파손,이탈,부식,노후화", Type: 2, Position: "안전난간대", Category: 10, Order: 101, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "변형,고정부 및 연결부 파손,이탈,부식,노후화", Type: 2, Position: "점검사다리 등받이 보호망", Category: 10, Order: 102, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "변형,파손,탈락", Type: 2, Position: "추락방호망", Category: 10, Order: 103, Periodic: id})

	periodicotherManager.Insert(&models.Periodicother{Name: "a,b,c,d,e", Type: 1, Position: "도로포장", Category: 11, Order: 110, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,함몰,단차(요청),블리딩,마모", Type: 2, Position: "아스팔트", Category: 11, Order: 111, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,마모,박락,파손", Type: 2, Position: "바닥 콘크리트", Category: 11, Order: 112, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,파손,침하", Type: 2, Position: "보도블럭", Category: 11, Order: 113, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,파손,침하", Type: 2, Position: "보도판석", Category: 11, Order: 114, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,파손,침하,이탈", Type: 2, Position: "경계석", Category: 11, Order: 115, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "파손,침하", Type: 2, Position: "트렌치", Category: 11, Order: 116, Periodic: id})

	periodicotherManager.Insert(&models.Periodicother{Name: "a,b,c,d,e", Type: 1, Position: "도로부 신축 이음부", Category: 12, Order: 120, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "미시공,고무재 및 강재,후타재", Type: 3, Position: "시공", Category: 12, Order: 121, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "마모,강판노출,부식,누수,단차,파손,이격,이물질 퇴적", Type: 2, Position: "신축 이음부 고무재 및 강재", Category: 12, Order: 122, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "콘크리트 균열PeriodiType:,파손", Type: 2, Position: "신축 이음부 후타재", Category: 12, Order: 123, Periodic: id})

	periodicotherManager.Insert(&models.Periodicother{Name: "a,b,c,d,e", Type: 1, Position: "환기구 등의 덮개", Category: 13, Order: 130, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "벽부형,입상형,바닥형,벤츄레이터", Type: 4, Position: "환기구", Category: 13, Order: 131, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "변형,파손,부식", Type: 2, Position: "그릴창", Category: 13, Order: 132, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "탈락,변형,부식", Type: 2, Position: "연결부", Category: 13, Order: 133, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "파손,균열", Type: 2, Position: "지지구조", Category: 13, Order: 134, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "파손,균열", Type: 2, Position: "걸침턱", Category: 13, Order: 135, Periodic: id})

	periodicotherManager.Insert(&models.Periodicother{Name: "a,b,c,d,e", Type: 1, Position: "외벽 마감재", Category: 14, Order: 140, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "수성페인트,A/L 판넬,커튼월,석재,적벽돌,드라이비트", Type: 4, Position: "건물외부 벽체", Category: 14, Order: 141, Periodic: id})
	//periodicotherManager.Insert(&models.Periodicother{Name: "(층간/경사/수직)균열,철근노출,콘크리트 (박리/박락)", Type: 2, Position: "수성페인트 마감부위", Category: 14, Order: 142, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,파손,탈락,이격", Type: 2, Position: "석재(화강석, 대리석) 마감", Category: 14, Order: 143, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "변형,파손,탈락,이격,부식", Type: 2, Position: "알루미늄 판넬 및 강재 마감", Category: 14, Order: 144, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "변형,균열,파손,탈락", Type: 2, Position: "드라이비트", Category: 14, Order: 145, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "균열,백화현상,파손,탈락", Type: 2, Position: "벽돌", Category: 14, Order: 146, Periodic: id})

	periodicotherManager.Insert(&models.Periodicother{Name: "해당사항 없음,해당사항 있음", Status: "해당사항 없음", Type: 1, Position: "강재구조 노후", Category: 15, Order: 150, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "부식,파손,이격", Type: 2, Position: "비구조형강", Category: 15, Order: 151, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "부식,파손,볼트풀림", Type: 2, Position: "철골구조물접합부위", Category: 15, Order: 152, Periodic: id})
	periodicotherManager.Insert(&models.Periodicother{Name: "들뜸,탈락", Type: 2, Position: "내화피복", Category: 15, Order: 153, Periodic: id})

	var periodickeep models.Periodickeep
	periodickeep.Periodic = item.Id
	periodickeep.Status1 = 1
	periodickeep.Status2 = 1
	periodickeep.Status3 = 1
	periodickeep.Status4 = 2
	periodickeep.Status5 = 1
	periodickeep.Status6 = 2

	periodickeep.Content1 = "- 준공내역서\n- 공사시방서\n- 각종계산서\n- 토질조사 보고서\n- 기타 특이사항 보고서"
	periodickeep.Content2 = "- 건축, 구조, 기계, 전기, 통신, 토목"
	periodickeep.Content3 = "◦기본현황\n◦상제제원\n◦유지관리 이력"
	periodickeep.Content4 = "- 관리소 미 보유"
	periodickeep.Content5 = "- 2022년 하반기 정기안전점검\n- 2022년 상반기 정밀안전점검"
	periodickeep.Content6 = "- 보수이력 없음"

	periodickeep.Remark1 = "FMS 등록 되어 있음"
	periodickeep.Remark2 = "FMS 등록 되어 있음"
	periodickeep.Remark3 = "FMS 등록 되어 있음"
	periodickeep.Remark4 = "FMS 미등록 되어 있음"
	periodickeep.Remark5 = "FMS 등록 되어 있음"
	periodickeep.Remark6 = "FMS 미등록 되어 있음"

	periodickeepManager.Insert(&periodickeep)

	os.Mkdir(fmt.Sprintf("%v/periodicresult/%v", config.UploadPath, item.Id), 0755)
}

type BlueprintData struct {
	Blueprint []models.Blueprint `json:"blueprint"`
}

func (c *PeriodicController) Data(id int64) {
	conn := c.NewConnection()

	blueprintManager := models.NewBlueprintManager(conn)
	blueprints := blueprintManager.Find([]any{
		models.Where{Column: "apt", Value: id, Compare: "="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id"),
	})

	data := BlueprintData{Blueprint: blueprints}

	c.Set("item", data)
}

// @POST()
func (c *PeriodicController) Upload(item *models.Data) {
	if c.Session != nil {
		item.User = c.Session.Id
	}

	conn := c.NewConnection()

	conn.Begin()
	defer conn.Rollback()

	periodicManager := models.NewPeriodicManager(conn)
	periodicimageManager := models.NewPeriodicimageManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	periodicdataimageManager := models.NewPeriodicdataimageManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)
	periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)

	periodicItem := periodicManager.Get(item.Id)
	log.Println("periodic not found", item.Id)
	if periodicItem == nil {
		return
	}
	blueprintItems := blueprintManager.FindByApt(periodicItem.Apt)

	now := global.GetDatetime(time.Now())

	for _, v := range item.Images {
		if periodicimageManager.CountByOfflinefilename(v.Offlinefilename) > 0 {
			continue
		}

		v.Use = 1
		v.Periodic = item.Id
		v.Date = now

		periodicimageManager.Insert(&v)
	}

	blueprints := make(map[int64][]models.Periodicdata, 0)

	for _, v := range item.Datas {
		if v.Numberzoom <= 0.0 {
			v.Numberzoom = v.Iconzoom
		}
		if v.Crackzoom <= 0.0 {
			v.Crackzoom = v.Iconzoom
		}
		periodicblueprintzoom := periodicblueprintzoomManager.GetByPeriodicBlueprint(item.Id, v.Id)
		if periodicblueprintzoom == nil || periodicblueprintzoom.Id == 0 {
			periodicblueprintzoomManager.Insert(&models.Periodicblueprintzoom{Iconzoom: models.Double(v.Iconzoom), Numberzoom: models.Double(v.Numberzoom), Crackzoom: models.Double(v.Crackzoom), Zoom: models.Double(v.Zoom), Status: 1, Periodic: item.Id, Blueprint: v.Id, Date: now})
		} else {
			periodicblueprintzoom.Iconzoom = models.Double(v.Iconzoom)
			periodicblueprintzoom.Numberzoom = models.Double(v.Numberzoom)
			periodicblueprintzoom.Crackzoom = models.Double(v.Crackzoom)
			periodicblueprintzoom.Zoom = models.Double(v.Zoom)
			periodicblueprintzoom.Status = 1
			periodicblueprintzoom.Date = now
			periodicblueprintzoomManager.Update(periodicblueprintzoom)
		}

		blueprints[v.Id] = make([]models.Periodicdata, 0)

		periodicdatas := periodicdataManager.Find([]any{
			models.Where{Column: "periodic", Value: item.Id, Compare: "="},
			models.Where{Column: "blueprint", Value: v.Id, Compare: "="},
		})

		for _, v2 := range periodicdatas {
			//periodicdataimageManager.DeleteByPeriodicdata(v2.Id)
			periodicdataimageManager.FakeDelete(item.Id, v2.Id)
		}

		//periodicdataManager.DeleteByPeriodicBlueprint(item.Id, v.Id)
		periodicdataManager.FakeDelete(item.Id, v.Id)

		for j, v2 := range v.Points {
			count := global.Atoi(v2.Count)
			progress := 1
			if v2.Progress == "X" {
				progress = 2
			}

			filenames := make([]string, 0)
			offlinefilenames := make([]string, 0)
			for i, v3 := range v2.Images {
				filenames = append(filenames, v2.Onlineimages[i])
				offlinefilenames = append(offlinefilenames, v3)
			}

			filename := strings.Join(filenames, ",")
			offlinefilename := strings.Join(offlinefilenames, ",")

			b, _ := json.Marshal(v2.Items)
			periodicdata := models.Periodicdata{Group: v2.Number, Type: v2.Icon, Part: v2.Part, Member: v2.Member, Shape: v2.Shape, Width: v2.Weight, Length: v2.Length, Count: count, Progress: progress, Remark: v2.Remark, Blueprint: v.Id, Periodic: item.Id, Content: string(b), Order: j + 1, Status: 1, Filename: filename, Offlinefilename: offlinefilename, User: item.User, Date: now}
			periodicdataManager.Insert(&periodicdata)
			id := periodicdataManager.GetIdentity()

			periodicdata.Id = id

			blueprints[v.Id] = append(blueprints[v.Id], periodicdata)

			for i, v3 := range v2.Images {
				dataimage := models.Periodicdataimage{Filename: v2.Onlineimages[i], Offlinefilename: v3, Order: i + 1, Periodicdata: id, Periodic: item.Id, Date: now}
				periodicdataimageManager.Insert(&dataimage)
			}
		}
	}

	periodicotherManager := models.NewPeriodicotherManager(conn)

	for _, v := range item.Periodicothers {
		if v.Change == 0 {
			continue
		}

		periodicotherManager.Update(&v)
	}

	for k := range blueprints {
		for _, v2 := range blueprintItems {
			if k == v2.Id {
				global.SendNotify(global.Notify{Type: global.NotifyBlueprint, Periodic: item.Id, Blueprint: k})
				break
			}
		}
	}
	conn.Commit()
}

// @POST()
func (c *PeriodicController) Duplication(id int64) {
	conn := c.NewConnection()

	conn.Begin()
	defer conn.Rollback()
	conn.Isolation = false

	periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
	periodicchangeManager := models.NewPeriodicchangeManager(conn)
	periodiccheckManager := models.NewPeriodiccheckManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	periodicimageManager := models.NewPeriodicimageManager(conn)
	periodicincidentalManager := models.NewPeriodicincidentalManager(conn)
	periodicopinionManager := models.NewPeriodicopinionManager(conn)
	periodicouterwallManager := models.NewPeriodicouterwallManager(conn)
	periodicpastManager := models.NewPeriodicpastManager(conn)
	periodicresultManager := models.NewPeriodicresultManager(conn)
	periodictechnicianManager := models.NewPeriodictechnicianManager(conn)
	periodicManager := models.NewPeriodicManager(conn)

	managebookManager := models.NewManagebookManager(conn)
	periodicotherManager := models.NewPeriodicotherManager(conn)
	periodicotheretcManager := models.NewPeriodicotheretcManager(conn)
	periodickeepManager := models.NewPeriodickeepManager(conn)

	periodic := periodicManager.Get(id)
	if periodic == nil || periodic.Id == 0 {
		return
	}

	now := global.GetDatetime(time.Now())
	periodic.Id = 0
	periodic.Date = now

	periodicManager.Insert(periodic)
	newId := periodicManager.GetIdentity()

	{
		items := managebookManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("mb_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			managebookManager.Insert(&v)
		}
	}

	{
		items := periodicotherManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("po_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Filename = ""
			v.Offlinefilename = ""
			v.Date = now

			periodicotherManager.Insert(&v)
		}
	}

	{
		items := periodicotheretcManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pe_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicotheretcManager.Insert(&v)
		}
	}

	{
		items := periodicchangeManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pc_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicchangeManager.Insert(&v)
		}
	}

	{
		items := periodiccheckManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pc_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodiccheckManager.Insert(&v)
		}
	}

	{
		items := periodicincidentalManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pi_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicincidentalManager.Insert(&v)
		}
	}

	{
		items := periodicopinionManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("po_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicopinionManager.Insert(&v)
		}
	}

	{
		items := periodicouterwallManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("po_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicouterwallManager.Insert(&v)
		}
	}

	{
		items := periodicpastManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pp_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicpastManager.Insert(&v)
		}
	}

	{
		items := periodicresultManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pr_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicresultManager.Insert(&v)
		}
	}

	{
		items := periodictechnicianManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("dt_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodictechnicianManager.Insert(&v)
		}
	}

	{
		items := periodicimageManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pi_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodicimageManager.Insert(&v)
		}
	}

	{
		items := periodickeepManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pk_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Date = now

			periodickeepManager.Insert(&v)
		}
	}

	os.Mkdir(fmt.Sprintf("%v/periodicresult/%v", config.UploadPath, newId), 0755)

	{
		items := periodicdataManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pd_id"),
		})

		for _, v := range items {
			// oldId := v.Id

			v.Id = 0
			v.Periodic = newId
			v.Filename = ""
			v.Offlinefilename = ""
			v.Date = now

			periodicdataManager.Insert(&v)

			// dataId := periodicdataManager.GetIdentity()

			// images := periodicdataimageManager.Find([]interface{}{
			// 	models.Where{Column: "periodic", Value: id, Compare: "="},
			// 	models.Where{Column: "periodicdata", Value: oldId, Compare: "="},
			// 	models.Ordering("pi_id"),
			// })

			// for _, v2 := range images {
			// 	v2.Id = 0
			// 	v2.Periodicdata = dataId
			// 	v2.Periodic = newId
			// 	v2.Date = now
			// 	periodicdataimageManager.Insert(&v2)
			// }

			// source := fmt.Sprintf("%v/periodicresult/%v/%v.jpg", config.UploadPath, id, v.Blueprint)
			// target := fmt.Sprintf("%v/periodicresult/%v/%v.jpg", config.UploadPath, newId, v.Blueprint)

			// global.CopyFile(source, target)
		}
	}

	{
		items := periodicblueprintzoomManager.Find([]any{
			models.Where{Column: "periodic", Value: id, Compare: "="},
			models.Ordering("pb_id"),
		})

		for _, v := range items {
			v.Id = 0
			v.Periodic = newId
			v.Status = 1
			v.Date = now

			periodicblueprintzoomManager.Insert(&v)
		}
	}

	conn.Commit()
}

func (c *PeriodicController) Search() {
	apt := c.Geti("apt")
	category := c.Geti("category")
	categorys := c.Get("categorys")
	status := c.Geti("status")
	page := c.Geti("page")
	pagesize := c.Geti("pagesize")

	if c.Session == nil {
		c.Result["code"] = "auth error"
		return
	}

	if c.Session.Level < 3 {

		if c.Session.Apt == 0 {
			c.Result["code"] = "auth error"
			return
		}

	}

	conn := c.NewConnection()

	manager := models.NewPeriodicManager(conn)

	var args []any

	_title := c.Get("title")
	if _title != "" {
		args = append(args, models.Custom{Query: fmt.Sprintf("(a_name like '%%%v%%' or d_name like '%%%v%%')", _title, _title)})
	}

	if c.Session.Level < 3 {
		args = append(args, models.Where{Column: "apt", Value: c.Session.Apt, Compare: "="})
	}

	if category != 0 {
		args = append(args, models.Where{Column: "category", Value: category, Compare: "="})
	}

	if apt != 0 {
		args = append(args, models.Where{Column: "apt", Value: apt, Compare: "="})
	}

	if categorys != "" {
		values := global.StringToIntArray(categorys)
		args = append(args, models.Where{Column: "category", Value: values, Compare: "in"})
	}

	log.Println("status", status)
	if status != 0 {
		args = append(args, models.Where{Column: "status", Value: status, Compare: "="})
	}

	if page != 0 && pagesize != 0 {
		args = append(args, models.Paging(page, pagesize))
	}

	orderby := c.Get("orderby")
	if orderby == "" {
		if page != 0 && pagesize != 0 {
			orderby = "id desc"
			args = append(args, models.Ordering(orderby))
		}
	} else {
		orderbys := strings.Split(orderby, ",")

		str := ""
		for i, v := range orderbys {
			if i == 0 {
				str += v
			} else {
				if strings.Contains(v, "_") {
					str += ", " + strings.Trim(v, " ")
				} else {
					str += ", d_" + strings.Trim(v, " ")
				}
			}
		}

		args = append(args, models.Ordering(str))
	}

	items := manager.Find(args)
	c.Set("items", items)

	total := manager.Count(args)
	c.Set("total", total)
}

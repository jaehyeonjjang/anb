package api

import (
	"repair/controllers"
	"repair/global"
	"repair/models"
	"repair/models/repair"
	"time"
)

type CategoryController struct {
	controllers.Controller
}

// @Post()
func (c *CategoryController) InitData(apt int64) {
	conn := c.NewConnection()

	repairManager := models.NewRepairManager(conn)
	standardManager := models.NewStandardManager(conn)
	dongManager := models.NewDongManager(conn)
	manager := models.NewCategoryManager(conn)

	var item models.Repair
	item.Apt = apt
	item.Type = 1
	item.Provision = 1
	item.Planyears = 50
	item.Reportdate = global.GetDate(time.Now().Add(time.Hour * 9))

	item.Content1 = `국토교통부에서는 유권해석으로 전기, 승강기, 급수시설 등 중요시설로서 예측이 불가한 사고에 의해 불가피한 긴급공사가 필요한 경우에 한해 장기수선계획의 조정, 장기수선충당금의 사용에 대해 법령에서 정한 예외적 집행을 인정하고 있다.`

	item.Content2 = `긴급공사 공용부위 주요시설이나 설비가 갑작스런 사고 등 예기치 못한 사정에 따라 수선주기가 도래하지 않았음에도 장기수선충당금을 사용해야 할 경우 총론에 마련한 장기수선충당금 사용근거에 따라 입주자대표회의 의결 후 장기수선충당금을 사용하고 추후 장기수선계획을 변경한다.
소액지출 : 갑작스런 배관의 누수, 배수펌프의 고장, 승강기 고장 등 예기치 못한 경우로써 장기수선계획의 수선주기에 따라 수선이나 보수를 할 수 없는 경우 예외적으로 소액의 범위내에서 장기수선 대상범위와 사용요건, 금액한도를 정하여 사용이 가능하도록 규정한다.`

	item.Content3 = `긴급공사 / 소액지출 : 장기수선계획에 반영되어 있는 항목 또는 공동주택관리법 시행규칙 별표1의 항목이 있는 경우 『공동주택관리법 시행규칙』 별표1 공사항목 중에서 단지에 설치된 단위 개수가 많고 단가가 소액인 항목을 일부 교체 또는 수리가 필요할 경우 장기수선공사관리대장에 기록 후 장기수선충당금 소액지출을 선 집행하고, 차기 장기수선계획 검토 및 조정 시 장기수선계획에 반영한다.`

	item.Content4 = `긴급공사 / 소액지출 : 장기수선충당금의 사용은 장기수선계획에 따른다(공동주택관리법 제30조 제2항)`

	item.Content5 = `긴급공사 / 소액지출 : 교체나 보수를 먼저 실시하고 최초 도래하는 정기 또는 수시검토 시 장기수선계획에 반영한다.`

	item.Content6 = `긴급공사 / 소액지출 : 장기수선계획서 총론에 소액지출로 집행이 가능한 긴급공사 내용, 처리방법 및 절차를 명시하여 입주자대표회의 의결 후 사용한다.`

	item.Content7 = `건물 내∙외부의 지붕이나 외부 등의 문제로 누수가 발생한 경우
예비전원 설비의 고장으로 단전에 대비할 수 없는 경우(발전기, 배전반 등)
변전설비의 고장으로 수전할 수 없거나 전원공급이 어려운 경우(변압기, 수전반<ACB, LBS, VCB>, 배전반<MCCB>, 고장구간 자동개폐기<ASS> 등)
전기설비의 지적사항이 발생하여 보수가 긴급한 경우
소방설비 고장으로 화재진압이나 경보가 이루어지지 않아 긴급히 보수(교체 포함)해야 하는 경우(감지기, 수신반<P형1급, R형>, 소화펌프, 스프링쿨러, 소화수관 등)
소방시설 점검 지적 사항으로 긴급히 보수(교체 포함)를 해야 하는 경우. 특히 소방점검 지적사항 중 수선비 항목과 장기수선충당금 항목을 구분하여 사용한다. 소화기 충압불량, 램프파손, 각 종 기구 탈락 및 미부착, 전원불량, 복합수신기 프로그램 리셋 등은 수선유지비로 보수를 하여야 하고, 오작동, 소화기 비치, 연동 불량, 선로단선 불량, 각 종 밸브류 교체 등은 장기수선충당금으로 보수한다.
승강기 고장, 사용이 불가능하여 안전에 문제가 발생되어 긴급히 보수(교체 포함)해야 하는 경우. 승강기 부품이 정확하게 구분이 불가한 관계로 다음과 같은 항목은 장기수선충당금으로 보수를 한다. 권상기, 전동기, 오일씰, 포텐셜 메타, 브레이크 쉬 어쉬, 속도 검출장치, 베아링 등은 기계장치에 속하는 부품 메인쉬브(도르래), 세컨쉬브, 샤클로드, 바빗메탈 등은 쉬브에 속하는 부품, 메인로브, 세컨로프, 연도로프, 카가이드(슈타입, 롤러타입), 컴펜로프 등은 와이어로프에 속하는 부품. 인버터, PCB보드, 마그네트, 비상정지 장치, 엠코더, 냉각 휀, 전자 접촉기, 전자 릴레이, 전원 공급장치, 파워팩 등은 제어반에 속하는 부품, 본체, 거버너, 균형체인 등은 조속기에 속하는 부품, 도어모터, 도어센서, 카도어, 싱크로밸트, 도어씰, 인터록, 행거플랜, 도어스위치, 도어슈, 카 고정장치, 도어 클러치, 도어 드라이맘 등은 도어개폐장치에 속하는 부품
승강기 점검 지적 사항으로 긴급히 보수(교체 포함)를 해야 하는 경우
승강기 부품교체 공사시 종합 유지보수계약 후 장기수선계획에 따른 부품교체 공사비용은 장기수선충당금(월 : 만원)으로 집행하고,유지보수비용은 관리비(월 : 만원 )로 집행
피뢰설비, 통신·방송수신 설비, 보안·방범시설의 고장으로 피해가 우려되는 경우(피뢰기, 비상방송용 앰프, 녹화기, 모니터, CCTV카메라 등)
홈네트워크 설비의 고장 등으로 입주자 등의 안전과 불편을 초래하는 경우(로비폰, 인터폰 등)
급수·소방·배수·난방·급탕시설 등의 고장으로 급수 공급 및 배수 문제로 긴급히  보수(교체 포함)을 해야 하는 경우(배관 및 펌프)
급수 · 소방 · 배수 · 난방 · 급탕시설 등의 배관 및 밸브류 등의 파열로 누수가 발생한 경우(감압, 게이트, 글로브, 버터플라이, 체크밸브, 스트레아나 등 각종 밸브류)
어린이놀이시설 정기검사에서 지적사항으로 보수를 요하는 경우
옥외 부대시설·복리시설에 문제가 발생하여 안전에 위험이 발생한 경우
장기수선계획의 조정(재수립 포함)을 외부업체에게 대행하여 용역비가 발생한 경우
각종 공사의 재해예방기술지도비, 기술지도점검비, 시공 및 설계감리비 등은 총론에 의거하여 장기수선충당금으로 집행한다.
수시조정을 위한 주민들에게 동의 받는 전자투표 수수료는 장기수선충당금으로 사용할 수 있다.`

	item.Content8 = `긴급한 고장이나 문제가 발생하였으나 장기수선 계획의 주기에 이르지 않았을 것.						
긴급한 보수를 요하는 경우로서 장기수선 계획 주기를 기다릴 수 없는 긴급성이 있을 것.
그 긴급성이 장기수선 계획 조정을 기다릴 수 없을 것.
신변 안전이나 그로 인해 시설물 또는 입주민에게 2차 피해가 현존하고 중 할 것.
위에서 언급한 항목들이 공용부분의 주요시설로 장기수선계획에 포함되어 있어야 하고, 소액인 경우
소액이고 지속적이고 반복적일 것.`

	item.Content9 = `관리주체는 주요시설이나 설비 등의 갑작스런 고장이나 파손으로 긴급보수(교체 포함)를 해야 할 상황이 발생하면 유선·사진 전송 등을 통하여 입주자대표회의에 보고한다.
보고를 받은 입주자대표회의는 긴급성 여부를 판단하여 회의를 소집할 필요가 없을 정도(입주자 등의 안전을 위협하고 2차 피해 등의 발생으로 더 큰 피해가 예상될 경우 등)의 긴급을 요하는 사항은 선 집행하고 입주자대표회의의 추인을 받도록 한다.
기타 긴급한 사항에 대해서는 입주자대표회의에 정식 안건 상정 후 입주자대표회의의 의결을 받아 집행한다.
관리주체가 장기수선계획 총론을 입주자대표회의에 보고하여 의결을 받은 사항에 대하여 소액 지출(시재금 지출)로 보수를 하여야 한다. 이때 장기수선충담금 사용계획서를 작성하여 같이 의결을 받는다.
관리주체가 장기수선충당금(시재금)을 사용 후 보수이력과 사용금액을 결산 시에 보고 하여야 한다.
위에서 언급한 사항에 대하여 의결 후 집행한 보수(공사)건들을 모았다가 정기검토 주기에 반영한다.
장기수선계획에 없으나 보수가 필요한 경우는 우선 입주자대표회의 의결 후 집행하고 부정기 검토 및 조정(입주자 과반 이상 동의)을 실시하여 반영한다.`

	item.Content10 = `긴급의 단일공사 2000만원 이내이고 년간공사 9000만원 이내로 하고 소액은 500만원 수의계약가능하나 그 이상은 입찰로 진행한다.
긴급 및 소액 지출건에 대한 부분은 보통예금 통장으로 관리한다.`

	item.Content11 = `위에서 언급한 바와 같이 시설물의 내구연한을 연장하고 입주민의 안전을 도모하기 위해서는 보다 철저한 유지 보수가 필요한 부분이 많은 관계로 장기수선충담금 적립을 현실화하는 부분이 마땅하나, 갑작스런 인상으로 입주민에게 피해를 최소화하기 위해서는 좀 더 현실성 있는 인상안이 필요하며, 또한 적립요율에 대하여는 현재 관리규약의 요율과 부과단가로 본 요율이 상이한 관계로 차후 관리규약 개정시 적립요율 조정이 필요하다고 사료됨.`

	item.Periodtype = repair.PeriodtypePeriodic

	repairManager.Insert(&item)
	id := repairManager.GetIdentity()

	ids := make(map[int64]int64)

	items := manager.Find([]interface{}{models.Where{Column: "apt", Value: -1, Compare: "="}, models.Ordering("c_order,c_id")})

	for _, level1 := range items {
		if level1.Apt != -1 {
			continue
		}

		if level1.Level == 1 {
			oldLevel1Id := level1.Id

			level1.Id = 0
			level1.Apt = id

			manager.Insert(&level1)
			level1Id := manager.GetIdentity()

			ids[oldLevel1Id] = level1Id

			for _, level2 := range items {
				if level2.Level == 2 && level2.Parent == oldLevel1Id {
					oldLevel2Id := level2.Id

					level2.Id = 0
					level2.Apt = id
					level2.Parent = level1Id

					manager.Insert(&level2)
					level2Id := manager.GetIdentity()

					ids[oldLevel2Id] = level2Id

					for _, level3 := range items {
						if level3.Level == 3 && level3.Parent == oldLevel2Id {
							oldLevel3Id := level3.Id

							level3.Id = 0
							level3.Apt = id
							level3.Parent = level2Id

							manager.Insert(&level3)
							level3Id := manager.GetIdentity()

							ids[oldLevel3Id] = level3Id

							for _, level4 := range items {
								if level4.Level == 4 && level4.Parent == oldLevel3Id {
									level4.Id = 0
									level4.Apt = id
									level4.Parent = level3Id

									manager.Insert(&level4)
								}
							}
						}
					}
				}
			}
		}
	}

	standards := standardManager.Find([]interface{}{models.Where{Column: "apt", Value: -1, Compare: "="}, models.Ordering("s_order,s_id")})

	for _, v := range standards {
		if v.Id == 0 {
			continue
		}

		if v.Apt != -1 {
			continue
		}

		v.Id = 0
		v.Apt = id
		v.Category = ids[v.Category]

		standardManager.Insert(&v)
	}

	dongs := dongManager.Find([]interface{}{models.Where{Column: "apt", Value: -1, Compare: "="}, models.Ordering("d_order,d_id")})

	for _, v := range dongs {
		if v.Id == 0 {
			continue
		}

		if v.Apt != -1 {
			continue
		}

		v.Id = 0
		v.Apt = id

		dongManager.Insert(&v)
	}
}

// @Post()
func (c *CategoryController) DuplicationData(apt int64) {
	conn := c.NewConnection()

	conn.Begin()
	defer conn.Rollback()

	repairManager := models.NewRepairManager(conn)
	standardManager := models.NewStandardManager(conn)
	dongManager := models.NewDongManager(conn)
	adjustManager := models.NewAdjustManager(conn)
	manager := models.NewCategoryManager(conn)

	repair := repairManager.Get(apt)

	if repair == nil {
		return
	}

	repair.Id = 0
	repair.Date = ""
	repair.Type = 2
	repair.Reportdate = global.GetDate(time.Now().Add(time.Hour * 9))
	repairManager.Insert(repair)
	id := repairManager.GetIdentity()

	ids := make(map[int64]int64)

	items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}, models.Ordering("c_order,c_id")})
	adjusts := adjustManager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}, models.Ordering("aj_order,aj_id")})

	for _, level1 := range items {
		if level1.Level == 1 {
			oldLevel1Id := level1.Id

			level1.Id = 0
			level1.Apt = id
			level1.Date = ""

			manager.Insert(&level1)
			level1Id := manager.GetIdentity()

			ids[oldLevel1Id] = level1Id

			for i, adjust := range adjusts {
				if oldLevel1Id == adjust.Category {
					adjusts[i].Category = level1Id
				}
			}

			for _, level2 := range items {
				if level2.Level == 2 && level2.Parent == oldLevel1Id {
					oldLevel2Id := level2.Id

					level2.Id = 0
					level2.Apt = id
					level2.Parent = level1Id
					level2.Date = ""

					manager.Insert(&level2)
					level2Id := manager.GetIdentity()

					ids[oldLevel2Id] = level2Id

					for i, adjust := range adjusts {
						if oldLevel2Id == adjust.Category {
							adjusts[i].Category = level2Id
						}
					}

					for _, level3 := range items {
						if level3.Level == 3 && level3.Parent == oldLevel2Id {
							oldLevel3Id := level3.Id

							level3.Id = 0
							level3.Apt = id
							level3.Parent = level2Id
							level3.Date = ""

							manager.Insert(&level3)
							level3Id := manager.GetIdentity()

							ids[oldLevel3Id] = level3Id

							for i, adjust := range adjusts {
								if oldLevel3Id == adjust.Category {
									adjusts[i].Category = level3Id
								}
							}

							for _, level4 := range items {
								if level4.Level == 4 && level4.Parent == oldLevel3Id {
									oldLevel4Id := level4.Id

									level4.Id = 0
									level4.Apt = id
									level4.Parent = level3Id
									level4.Date = ""

									manager.Insert(&level4)
									level4Id := manager.GetIdentity()

									ids[oldLevel4Id] = level4Id
								}
							}
						}
					}
				}
			}
		}
	}

	standardIds := make(map[int64]int64)

	standards := standardManager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}, models.Ordering("s_order,s_id")})

	for _, v := range standards {
		if v.Id == 0 {
			continue
		}

		oldId := v.Id

		v.Id = 0
		v.Apt = id
		v.Category = ids[v.Category]

		standardManager.Insert(&v)
		newId := standardManager.GetIdentity()

		standardIds[oldId] = newId

		for i, adjust := range adjusts {
			if oldId == adjust.Standard {
				adjusts[i].Standard = newId
			}
		}
	}

	dongIds := make(map[int64]int64)

	dongs := dongManager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}, models.Ordering("d_order,d_id")})

	for _, v := range dongs {
		if v.Id == 0 {
			continue
		}

		oldId := v.Id

		v.Id = 0
		v.Apt = id

		dongManager.Insert(&v)
		newId := dongManager.GetIdentity()

		dongIds[oldId] = newId
	}

	{
		manager := models.NewApprovalManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id

			manager.Insert(&v)
		}
	}

	{
		manager := models.NewAreaManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id

			manager.Insert(&v)
		}
	}

	{
		manager := models.NewBreakdownManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			v.Topcategory = ids[v.Topcategory]
			v.Subcategory = ids[v.Subcategory]
			v.Category = ids[v.Category]
			v.Method = ids[v.Method]
			v.Standard = standardIds[v.Standard]
			v.Dong = dongIds[v.Dong]

			manager.Insert(&v)
		}
	}

	{
		manager := models.NewHistoryManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			v.Topcategory = ids[v.Topcategory]
			v.Subcategory = ids[v.Subcategory]
			v.Category = ids[v.Category]

			manager.Insert(&v)
		}
	}

	{
		manager := models.NewOutlineplanManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			manager.Insert(&v)
		}
	}

	{
		manager := models.NewOutlineManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			manager.Insert(&v)
		}
	}

	{
		manager := models.NewReviewcontentManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			manager.Insert(&v)
		}
	}

	{
		manager := models.NewReviewdateManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			manager.Insert(&v)
		}
	}

	/*
		{
			manager := models.NewReviewManager(conn)
			items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

			for _, v := range *items {
				if v.Id == 0 {
					continue
				}

				v.Id = 0
				v.Apt = id
				manager.Insert(&v)
			}
		}
	*/

	{
		manager := models.NewSavingManager(conn)
		items := manager.Find([]interface{}{models.Where{Column: "apt", Value: apt, Compare: "="}})

		for _, v := range items {
			if v.Id == 0 {
				continue
			}

			v.Id = 0
			v.Apt = id
			manager.Insert(&v)
		}
	}

	for _, adjust := range adjusts {
		adjust.Id = 0
		adjust.Apt = id
		adjustManager.Insert(&adjust)
	}

	conn.Commit()
}

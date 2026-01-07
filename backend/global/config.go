package global

type Level int

const (
	_ Level = iota
	User
	Manager
	Admin
	RootAdmin
)

var Levels = []string{"", "작업자", "매니저", "관리자", "총 관리자"}

type AptStatus int

const (
	_ AptStatus = iota
	Ing
	Complete
	Close
)

var AptStatuss = []string{"", "진행", "완료", "마감"}

type UseType int

const (
	_ UseType = iota
	Use
	NotUse
)

var Uses = []string{"", "사용", "사용 안함"}

type DefaultStatus int

const (
	_ DefaultStatus = iota
	Default
	NotDefault
)

var Defaults = []string{"", "디폴트", "디폴트 아님"}

var Statuss = []string{"", "부재", "유형", "폭", "길이", "개소", "진행사항", "비고", "부재열", "부재명"}

package models

type Offset struct {
	Dx float64 `json:"dx"`
	Dy float64 `json:"dy"`
}

type Point struct {
	Items []Offset `json:"items"`

	Color  int `json:"color"`
	Type   int `json:"type"`
	Icon   int `json:"icon"`
	Number int `json:"number"`

	Part     string `json:"part"`
	Member   string `json:"member"`
	Shape    string `json:"shape"`
	Weight   string `json:"weight"`
	Length   string `json:"length"`
	Count    string `json:"count"`
	Progress string `json:"progress"`
	Remark   string `json:"remark"`
	Order    int    `json:"order"`

	Images       []string `json:"images"`
	Onlineimages []string `json:"onlineimages"`
}

type DataItem struct {
	Id         int64   `json:"id"`
	Points     []Point `json:"points"`
	Zoom       float64 `json:"zoom"`
	Iconzoom   float64 `json:"iconzoom"`
	Numberzoom float64 `json:"numberzoom"`
	Crackzoom  float64 `json:"crackzoom"`
}

type Data struct {
	Id             int64           `json:"id"`
	User           int64           `json:"user"`
	Datas          []DataItem      `json:"datas"`
	Images         []Periodicimage `json:"images"`
	Periodicothers []Periodicother `json:"periodicothers"`
}

type Aptdongblueprint struct {
	Apt     int64        `json:"apt"`
	Aptdong int64        `json:"aptdong"`
	Items   []Aptdongetc `json:"items"`
}

type Diff struct {
	Id             int64       `json:"id"`
	NewStandard    []Standard  `json:"newStandard"`
	ChangeStandard []Standard  `json:"changeStandard"`
	NewBreakdown   []Breakdown `json:"newBreakdown"`
	Filename       string      `json:"filename"`
}

type ConvertOldapt struct {
	Items []Oldapt `json:"items"`
}

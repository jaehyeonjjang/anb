package detail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"repair/global"
	"repair/global/config"
	"repair/models"
	"slices"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/LoperLee/golang-hangul-toolkit/hangul"
)

type PeriodicData struct {
	Type  int
	Title string
	Item  models.Periodicdata
}

type PeriodicDataList struct {
	Items []PeriodicData
}

type ImageInfo struct {
	Size     int64
	Data     string
	Text     string
	Status   string
	Number   string
	Filename string
	Ext      string
	Width    int
	Height   int
}

type PartImageInfo struct {
	Image1 ImageInfo
	Image2 ImageInfo
	Image3 ImageInfo
	Image4 ImageInfo
	Image5 ImageInfo
	Image6 ImageInfo
}

type ImageDouble struct {
	First  bool
	Second bool
}

type ImageSizeInfo struct {
	Width   int
	Height  int
	Top     int
	Left    int
	Width2  int
	Height2 int
}

type ResultMap struct {
	Dong  string
	Items []string
}

type AptusagefloorPage struct {
	Count int
	Items []models.Aptusagefloor
}

func GetThumbnail(filename string) string {
	ext := global.GetExt(filename)

	newFilename := fmt.Sprintf("temp/%v-thumbnail.%v", strings.ReplaceAll(filepath.Base(filename), "."+ext, ""), ext)
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, newFilename)
	_, err := os.Stat(fullFilename)
	if err != nil {
		global.MakeThumbnail(21780/50, 16335/50, fmt.Sprintf("%v/%v", config.UploadPath, filename), fmt.Sprintf("%v/%v", config.UploadPath, newFilename))
	}
	return newFilename
}

func Josa(str string, josa hangul.Josa) string {
	if str == "" {
		return ""
	}
	temp := strings.Split(str, "(")

	if len(temp) > 1 {
		if temp[0] == "" {
			temp = strings.Split(temp[1], ")")
		}

		t := global.GetJosa(temp[0], josa)

		josaStr := strings.ReplaceAll(t, temp[0], "")
		return str + josaStr
	}

	return global.GetJosa(str, josa)
}

func JoinInt(items []int) string {
	strs := make([]string, 0)

	for _, v := range items {
		strs = append(strs, fmt.Sprintf("%v", v))
	}

	return strings.Join(strs, "/")
}

func GetImageSize(filename string) (int, int) {
	fullFilename := filepath.Join(config.UploadPath, filename)
	return global.GetImageSize(fullFilename)
}

func GetImageInfo(filename string) ImageInfo {
	if filename == "" || filename == "webdata" {
		filename = "empty.jpg"
	}

	fullFilename := filepath.Join(config.UploadPath, filename)

	fi, err := os.Stat(fullFilename)
	if err != nil {
		log.Print("not found")
		fullFilename = filepath.Join(config.UploadPath, "empty.jpg")
		fi, _ = os.Stat(fullFilename)
	}
	/*
		b64Filename := fmt.Sprintf("%v.b64", filename)
		b64fullFilename := filepath.Join(config.UploadPath, b64Filename)

		log.Println("b64filenaem", b64fullFilename)

		dec := ""
		_, err2 := os.Stat(b64fullFilename)
		if err2 != nil {
			dat, _ := ioutil.ReadFile(fullFilename)
			dec := base64.StdEncoding.EncodeToString(dat)
			global.WriteFile(b64fullFilename, dec)
		} else {
			dat, _ := ioutil.ReadFile(b64fullFilename)
			dec = string(dat)
		}
	*/

	dat, _ := os.ReadFile(fullFilename)
	dec := base64.StdEncoding.EncodeToString(dat)

	width, height := GetImageSize(filename)
	info := ImageInfo{Filename: filename, Size: fi.Size(), Data: dec, Width: width, Height: height, Ext: global.GetExt(filename)}
	return info
}

func GetStructName(value int) string {
	switch value {
	case 1:
		return "철근콘크리트 구조"
	case 2:
		return "철골.철근콘크리트 구조"
	case 3:
		return "철근콘크리트라멘조+벽식 구조"
	case 4:
		return "철골.철근콘크리트구조 및 복합구조"
	case 5:
		return "프리케스트콘크리트 구조"
	case 6:
		return "철근콘크리트 PC구조"
	}

	return ""
}

func Detail1(id int64, conn *models.Connection) string {
	aptManager := models.NewAptManager(conn)
	periodicManager := models.NewPeriodicManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	aptperiodicManager := models.NewAptperiodicManager(conn)

	periodic := periodicManager.Get(id)
	apt := aptManager.Get(periodic.Apt)
	if periodic.Aptname != "" {
		apt.Name = periodic.Aptname
	}
	aptperiodic := aptperiodicManager.Get(periodic.Apt)
	aptdongs := aptdongManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
	})

	v := make(jet.VarMap)

	v.Set("apt", apt)
	v.Set("aptperiodic", aptperiodic)
	v.Set("periodic", periodic)

	underground := 0
	ground := 0
	for _, v := range aptdongs {
		if v.Private == 2 {
			continue
		}
		if v.Undergroundcount+v.Parkingcount > underground {
			underground = v.Undergroundcount + v.Parkingcount
		}

		if v.Groundcount > ground {
			ground = v.Groundcount
		}
	}

	floor := ""
	if underground > 0 && ground > 0 {
		floor = fmt.Sprintf("지하%v층, 지상%v층", underground, ground)
	} else if underground > 0 {
		floor = fmt.Sprintf("지하%v층", underground)
	} else {
		floor = fmt.Sprintf("지상%v층", ground)
	}

	v.Set("floor", floor)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	view.AddGlobal("humandate", func(str string) string {
		temp := strings.Split(str, "-")

		if len(temp) != 3 {
			return str
		}

		return fmt.Sprintf("%04d년 %02d월 %02d일", global.Atoi(temp[0]), global.Atoi(temp[1]), global.Atoi(temp[2]))
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("detail/detail-01.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("detail/detail-01-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Detail2(id int64, conn *models.Connection) string {
	periodicManager := models.NewPeriodicManager(conn)
	periodicblueprintzoomManager := models.NewPeriodicblueprintzoomManager(conn)
	blueprintManager := models.NewBlueprintManager(conn)

	periodic := periodicManager.Get(id)

	blueprints := blueprintManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id"),
	})

	zooms := periodicblueprintzoomManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
	})

	images := make([]ImageInfo, 0)
	titles := make([]string, 0)
	imagesizes := make([]ImageSizeInfo, 0)

	for _, v := range blueprints {
		if v.Upload != 1 {
			continue
		}

		find := false
		for _, v2 := range zooms {
			if v.Id == v2.Blueprint {
				find = true
				break
			}
		}

		filename := ""
		if find == true {
			filename = fmt.Sprintf("detailresult/%v/%v.jpg", id, v.Id)
		} else {
			filename = v.Filename
		}

		if filename == "" || filename == "webata" {
			continue
		}

		images = append(images, GetImageInfo(filename))

		width, height := GetImageSize(filename)

		w := 397600
		h := 281200

		w2 := 40479
		h2 := 28629

		newWidth := 0
		newHeight := 0

		newWidth2 := 0
		newHeight2 := 0

		rate := float64(w) / float64(h)
		target := float64(width) / float64(height)

		if rate > target {
			newWidth = int(float64(h) * target)
			newHeight = h

			newWidth2 = int(float64(h2) * target)
			newHeight2 = h2
		} else if rate < target {
			newWidth = w
			newHeight = int(float64(w) / target)

			newWidth2 = w2
			newHeight2 = int(float64(w2) / target)
		} else {
			newWidth = w
			newHeight = h

			newWidth2 = w2
			newHeight2 = h2
		}

		var info ImageSizeInfo

		info.Width = newWidth
		info.Height = newHeight
		info.Width2 = newWidth2
		info.Height2 = newHeight2
		info.Left = 0
		info.Top = 0

		imagesizes = append(imagesizes, info)

		title := v.Name

		if v.Level == 2 {
			for _, v3 := range blueprints {
				if v3.Id == v.Parent {
					title = fmt.Sprintf("%v %v", v3.Name, v.Name)
					break
				}
			}
		}

		titles = append(titles, title)
	}

	v := make(jet.VarMap)

	v.Set("imagecount", len(images))
	v.Set("images", images)

	items := make([]ImageDouble, 0)

	count := int(len(images) / 2)
	for range count {
		items = append(items, ImageDouble{First: true, Second: true})
	}

	if len(images)%2 != 0 {
		items = append(items, ImageDouble{First: true, Second: false})
		titles = append(titles, "")
	}

	v.Set("imagesizes", imagesizes)
	v.Set("itemcount", len(items))
	v.Set("items", items)
	v.Set("titles", titles)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("detail/detail-02.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("detail/detail-02-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Detail3(id int64, conn *models.Connection) string {
	return ""
}

func Detail4(id int64, conn *models.Connection) string {
	return ""
}

func Detail5(id int64, conn *models.Connection) string {
	periodicManager := models.NewPeriodicManager(conn)
	managebookManager := models.NewManagebookManager(conn)

	periodic := periodicManager.Get(id)

	managebooks := managebookManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("mc_order,mc_name,mb_order,mb_id"),
	})

	images := make([]ImageInfo, 0)

	v := make(jet.VarMap)

	for _, v := range managebooks {
		filename := fmt.Sprintf("detailresult/%v/%v", id, v.Filename)
		images = append(images, GetImageInfo(filename))
	}

	v.Set("imagecount", len(images))
	v.Set("images", images)

	items := make([]ImageDouble, 0)

	count := int(len(images) / 2)
	for range count {
		items = append(items, ImageDouble{First: true, Second: true})
	}

	if len(images)%2 != 0 {
		items = append(items, ImageDouble{First: true, Second: false})
	}

	v.Set("itemcount", len(items))
	v.Set("items", items)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("detail/detail-05.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("detail/detail-05-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Detail6(id int64, conn *models.Connection) string {
	periodicManager := models.NewPeriodicManager(conn)
	periodicdataManager := models.NewPeriodicdataManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	periodicotherManager := models.NewPeriodicotherManager(conn)

	periodic := periodicManager.Get(id)

	aptdongs := aptdongManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("au_order,au_id"),
	})

	allPeriodicdatas := periodicdataManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "group", Value: 1, Compare: ">="},
		models.Ordering("bp_parentorder,bp_order desc,bp_id,pd_order,pd_id"),
	})

	for i, v := range allPeriodicdatas {
		switch v.Type {
		case 3, 5:
			allPeriodicdatas[i].Type = 1
		case 4, 6:
			allPeriodicdatas[i].Type = 2
		}
	}

	periodicdatas := make([]models.Periodicdata, 0)

	for _, v := range allPeriodicdatas {
		if v.Type >= 200 {
			continue
		}

		periodicdatas = append(periodicdatas, v)
	}

	log.Println("periodicdatas len", len(periodicdatas))

	periodicothers := periodicotherManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("po_order,po_id"),
	})

	v := make(jet.VarMap)

	images := make([]ImageInfo, 0)

	number := 1
	for _, v := range periodicdatas {
		if v.Filename == "" {
			continue
		}

		blueprint := v.Extra["blueprint"].(models.Blueprint)

		var aptdong models.Aptdong
		for _, v2 := range aptdongs {
			if v2.Id == blueprint.Aptdong {
				aptdong = v2
				break
			}
		}

		title := ""
		if blueprint.Level == 1 {
			title = fmt.Sprintf("%v %v %v", blueprint.Name, v.Part, v.Member)
		} else {
			if aptdong.Private == 3 {
				title = fmt.Sprintf("%v %v %v", blueprint.Name, v.Part, v.Member)
			} else {
				title = fmt.Sprintf("%v %v %v %v", aptdong.Dong, blueprint.Name, v.Part, v.Member)
			}
		}

		strs := strings.SplitSeq(v.Filename, ",")

		for v2 := range strs {
			v2 = GetThumbnail(v2)
			info := GetImageInfo(v2)
			info.Text = title
			info.Status = v.Shape
			info.Number = fmt.Sprintf("%v", number)
			images = append(images, info)

			number++
		}
	}

	otherCategorys := []int{10, 11, 12, 13, 2, 14, 1, 3}

	for _, category := range otherCategorys {
		for _, v := range periodicothers {
			if category != v.Category {
				continue
			}

			if v.Filename == "" {
				continue
			}

			title := v.Position

			strs := strings.SplitSeq(v.Filename, ",")

			for v2 := range strs {
				v2 = GetThumbnail(v2)
				info := GetImageInfo(v2)
				info.Text = title

				if v.Status == "" {
					info.Status = "상태 양호"
					/*
						if v.Type == 1 {
							if v.Result == 1 {
								info.Status = "상태 양호"
							} else {
								info.Status = "상태 보통"
							}
						} else if v.Type == 2 {
							if v.Result == 1 {
								info.Status = "없음"
							} else {
								info.Status = "있음"
							}
						}
					*/
				} else {
					info.Status = v.Status
				}

				info.Number = fmt.Sprintf("%v", number)

				images = append(images, info)

				number++
			}
		}
	}

	if len(images)%6 > 0 {
		remain := len(images) % 6
		for i := 0; i < 6-remain; i++ {
			info := GetImageInfo("empty.jpg")
			images = append(images, info)
		}
	}

	v.Set("imagecount", len(images))
	v.Set("images", images)

	items := make([]PartImageInfo, 0)

	for i := 0; i < len(images)/6; i++ {
		var part PartImageInfo

		part.Image1 = images[i*6+0]
		part.Image2 = images[i*6+1]
		part.Image3 = images[i*6+2]
		part.Image4 = images[i*6+3]
		part.Image5 = images[i*6+4]
		part.Image6 = images[i*6+5]

		items = append(items, part)
	}

	v.Set("itemcount", len(items))
	v.Set("items", items)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("detail/detail-06.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("detail/detail-06-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

func Detail7(id int64, conn *models.Connection) string {
	aptManager := models.NewAptManager(conn)
	aptdongManager := models.NewAptdongManager(conn)
	periodicManager := models.NewPeriodicManager(conn)
	aptperiodicManager := models.NewAptperiodicManager(conn)
	periodicimageManager := models.NewPeriodicimageManager(conn)
	periodictechnicianManager := models.NewPeriodictechnicianManager(conn)
	periodickeepManager := models.NewPeriodickeepManager(conn)

	periodic := periodicManager.Get(id)
	apt := aptManager.Get(periodic.Apt)
	if periodic.Aptname != "" {
		apt.Name = periodic.Aptname
	}

	aptperiodic := aptperiodicManager.Get(periodic.Apt)

	periodicimages := periodicimageManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Where{Column: "use", Value: 1, Compare: "="},
		models.Ordering("pi_order,pi_id"),
	})

	technicians := periodictechnicianManager.Find([]any{
		models.Where{Column: "periodic", Value: id, Compare: "="},
		models.Ordering("dt_order,dt_id"),
	})

	aptdongs := aptdongManager.Find([]any{
		models.Where{Column: "apt", Value: periodic.Apt, Compare: "="},
		models.Ordering("au_order,au_id"),
	})

	periodickeep := periodickeepManager.GetByPeriodic(id)

	v := make(jet.VarMap)

	technicianCount := 0
	if technicians != nil {
		technicianCount = len(technicians)
	}

	v.Set("apt", apt)
	v.Set("aptperiodic", aptperiodic)
	v.Set("periodic", periodic)
	v.Set("technicianCount", technicianCount+1)
	v.Set("technicians", technicians)

	v.Set("keep1", strings.Split(periodickeep.Content1, "\n"))
	v.Set("keep2", strings.Split(periodickeep.Content2, "\n"))
	v.Set("keep3", strings.Split(periodickeep.Content3, "\n"))
	v.Set("keep4", strings.Split(periodickeep.Content4, "\n"))
	v.Set("keep5", strings.Split(periodickeep.Content5, "\n"))
	v.Set("keep6", strings.Split(periodickeep.Content6, "\n"))
	v.Set("periodickeep", periodickeep)

	temp := strings.Split(strings.ReplaceAll(periodickeep.Content2, "-", ""), ",")

	titles := []string{"건축", "구조", "기계", "전기", "토목", "통신"}
	strs := make([]string, 0)

	for _, v := range temp {
		str := strings.TrimSpace(v)

		find := slices.Contains(titles, str)

		if find == true {
			strs = append(strs, str)
		}
	}

	keepstr := ""
	if len(strs) > 0 {
		keepstr = fmt.Sprintf("[%v]", strings.Join(strs, "/"))
	}

	v.Set("keepstr", keepstr)

	images := make([]ImageInfo, 0)
	basicimages := make([]ImageInfo, 0)

	for i := 1; i <= 51; i++ {
		basicimages = append(basicimages, GetImageInfo(fmt.Sprintf("../doc/img/%v.jpg", i)))
	}

	sign := false
	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 4 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
		sign = true

		break
	}

	if sign == false {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("sign", sign)

	license := false
	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 6 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
		license = true

		break
	}

	if license == false {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("license", license)

	cert := false
	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 7 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
		cert = true

		break
	}

	if cert == false {
		images = append(images, GetImageInfo("empty.jpg"))
	}

	v.Set("cert", cert)

	for _, v := range periodicimages {
		if v.Filename == "" {
			continue
		}

		if v.Type != 5 {
			continue
		}

		images = append(images, GetImageInfo(v.Filename))
	}

	v.Set("imagecount", len(images))
	v.Set("images", images)
	v.Set("basicimages", basicimages)

	underground := 0
	ground := 0
	dongcount := 0
	for _, v := range aptdongs {
		if v.Private == 2 {
			continue
		}

		if v.Private == 1 {
			dongcount++
		}

		if v.Undergroundcount+v.Parkingcount > underground {
			underground = v.Undergroundcount + v.Parkingcount
		}

		if v.Groundcount > ground {
			ground = v.Groundcount
		}
	}

	floor := ""
	if underground > 0 && ground > 0 {
		floor = fmt.Sprintf("지하%v층, 지상%v층", underground, ground)
	} else if underground > 0 {
		floor = fmt.Sprintf("지하%v층", underground)
	} else {
		floor = fmt.Sprintf("지상%v층", ground)
	}

	v.Set("floor", floor)
	v.Set("dongcount", dongcount)

	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	view.AddGlobal("humandate", func(str string) string {
		temp := strings.Split(str, "-")

		if len(temp) != 3 {
			return str
		}

		return fmt.Sprintf("%04d년 %02d월 %02d일", global.Atoi(temp[0]), global.Atoi(temp[1]), global.Atoi(temp[2]))
	})

	var b bytes.Buffer
	t, err := view.GetTemplate("detail/detail-07.jet")
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}

	filename := fmt.Sprintf("detail/detail-07-%v-%v.hml", periodic.Id, global.UniqueId())
	fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, filename)
	global.WriteFile(fullFilename, b.String())

	return filename
}

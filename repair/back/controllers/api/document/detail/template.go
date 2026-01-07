package detail

import (
	"bytes"
	"fmt"
	"log"
	"repair/global"
	"repair/models"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dustin/go-humanize"
)

func GetTemplate(filename string, v jet.VarMap) string {
	var view = jet.NewSet(jet.NewOSFileSystemLoader("./doc"), jet.InDevelopmentMode())

	view.AddGlobal("structure", func(value int) string {
		return GetStructName(value)
	})

	view.AddGlobal("date", func(str string) string {
		return strings.ReplaceAll(str, "-", ".")
	})

	view.AddGlobal("year", func(str string) string {
		strs := strings.Split(str, "-")
		if strs[0] != "" {
			return strs[0] + "."
		}

		return ""
	})

	view.AddGlobal("monthday", func(str string) string {
		strs := strings.Split(str, "-")
		if len(strs) == 3 {
			return strs[1] + ". " + strs[2] + "."
		}

		return ""
	})

	view.AddGlobal("split", func(str string) []string {
		return strings.Split(str, "\n")
	})

	view.AddGlobal("humandate", func(str string) string {
		temp := strings.Split(str, "-")

		if len(temp) != 3 {
			return str
		}

		return fmt.Sprintf("%04d년 %02d월 %02d일", global.Atoi(temp[0]), global.Atoi(temp[1]), global.Atoi(temp[2]))
	})

	view.AddGlobal("without", func(str string) string {
		return strings.ReplaceAll(str, "동", "")
	})

	view.AddGlobal("dataCount", func(items []PeriodicData) int {
		return len(items) + 2
	})

	view.AddGlobal("pastCount", func(items []models.Periodicpast) int {
		return (len(items) + 1) * 2
	})

	view.AddGlobal("usageCount", func(items []models.Aptdong) int {
		count := 0

		for _, v := range items {
			if v.Private != 1 {
				continue
			}

			count++
		}
		return count + 1
	})

	view.AddGlobal("usagefloorCount", func(items []models.Aptusagefloor) int {
		return len(items) + 1
	})

	view.AddGlobal("dash", func(str string) string {
		if str == "" {
			return "-"
		}

		return str
	})

	view.AddGlobal("text", func(str string) string {
		if str == "" {
			return " "
		}

		return str
	})

	view.AddGlobal("comma", func(str string) string {
		if str == "" {
			str = "0"
		}
		v := global.Atol(str)
		return humanize.Comma(v)
	})

	view.AddGlobal("inclination", func(str string) string {
		temp := strings.Split(str, "/")

		if len(temp) < 2 {
			return ""
		}

		value := global.Atoi(temp[1])

		if value >= 750 {
			return "a"
		} else if value >= 500 {
			return "b"
		} else if value >= 250 {
			return "c"
		} else if value >= 150 {
			return "d"
		} else {
			return "e"
		}
	})

	view.AddGlobal("border", func(count int, pos int, head int, body int, tail int) string {
		value := body

		switch pos {
		case 0:
			value = head
		case count - 1:
			value = tail
		}

		return fmt.Sprintf("%v", value)
	})

	var b bytes.Buffer
	t, err := view.GetTemplate(filename)
	if err == nil {
		if err = t.Execute(&b, v, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	} else {
		log.Println("##################")
		log.Println(err)
	}

	return b.String()
}

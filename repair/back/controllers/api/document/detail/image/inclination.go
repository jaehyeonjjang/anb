package image

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"repair/global"
	"repair/global/config"
	"repair/models"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

func MakeInclinationImage(periodic int64, blueprint models.Blueprint, items []models.Periodicdata, iconZoom float64) {
	draw2d.SetFontFolder("./doc")

	os.Mkdir(fmt.Sprintf("%v/periodicresult/%v", config.UploadPath, periodic), 0755)

	filename := fmt.Sprintf("%v/%v", config.UploadPath, blueprint.Filename)
	img, _ := global.LoadImageFile(filename)

	if img == nil {
		log.Println("not found", filename)
		return
	}

	log.Println(filename)
	dest := global.ImageToRGBA(img)
	gc := draw2dimg.NewGraphicContext(dest)

	step := iconZoom / 2.0

	red := color.RGBA{0xf4, 0x43, 0x37, 0xff}

	for _, v := range items {
		log.Println("type", v.Type)
		var results []global.Offset
		json.Unmarshal([]byte(v.Content), &results)

		if len(results) == 0 {
			continue
		}

		point := results[len(results)-1]
		point2 := results[0]
		x := point.Dx
		y := point.Dy

		w := 4.0 * step / 50.0

		if w < 1 {
			w = 1
		}

		gc.SetLineWidth(w)

		if v.Type == 201 || v.Type == 202 || v.Type == 203 {
			gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
			gc.SetStrokeColor(red)

			gc.BeginPath()

			if v.Type == 201 {
				for i, point := range results {
					if i == 0 {
						gc.MoveTo(point.Dx, point.Dy)
					} else {
						gc.LineTo(point.Dx, point.Dy)
					}
				}
			} else if v.Type == 202 {
				gc.MoveTo(point2.Dx, point2.Dy)
				gc.LineTo(point.Dx, point2.Dy)
				gc.LineTo(point.Dx, point.Dy)
			} else if v.Type == 203 {
				gc.MoveTo(point2.Dx, point2.Dy)
				gc.LineTo(point2.Dx, point.Dy)
				gc.LineTo(point.Dx, point.Dy)
			}

			gc.Stroke()

			gc.BeginPath()
			gc.ArcTo(x, y, step, step, 0, math.Pi*2)
			gc.FillStroke()

			gc.SetFillColor(red)

			gc.SetFontSize(50 * iconZoom / 100)
			gc.SetFontData(draw2d.FontData{
				Name:   "luxi",
				Family: draw2d.FontFamilyMono,
				Style:  draw2d.FontStyleBold,
			})

			if v.Group >= 10 {
				gc.FillStringAt(fmt.Sprintf("%v", v.Group), point.Dx-step*0.7, point.Dy+step/2)
			} else {
				gc.FillStringAt(fmt.Sprintf("%v", v.Group), point.Dx-step*0.35, point.Dy+step/2)
			}
			gc.Stroke()

			gc.BeginPath()

			y1 := results[0].Dy
			y2 := results[1].Dy
			x1 := results[0].Dx
			x2 := results[1].Dx

			var dx float64 = 0.0
			var dy float64 = 0.0

			for i := 2; i < len(results); i++ {
				if x1 == x2 && y1 == y2 {
					y2 = results[i].Dy
					x2 = results[i].Dx
				}
			}

			if v.Type == 202 {
				y2 = y1
			} else if v.Type == 203 {
				x2 = x1
			}

			sy := y2 - y1
			sx := x2 - x1

			angle := math.Atan(sy/sx) * (180.0 / 3.14)

			angle += 180

			if sx < 0.0 {
				angle += 180.0
			} else {
				if sy < 0.0 {
					angle += 360.0
				}
			}

			r := (iconZoom / 2.0) * 0.4
			zoom := 1.0

			if v.Type == 202 {
				if point2.Dx > point.Dx {
					angle = 0
				} else {
					angle = 180
				}
			} else if v.Type == 203 {
				if point2.Dy > point.Dy {
					angle = 90
				} else {
					angle = 270
				}
			}

			xinc := math.Cos(angle*3.14/180) * r
			yinc := math.Sin(angle*3.14/180) * r

			x1 -= xinc
			y1 -= yinc

			log.Println(x1, xinc, zoom, dx)
			dx = 0
			dy = 0
			log.Println("start", (x1+xinc)*zoom+dx, (y1+yinc)*zoom+dy)
			gc.MoveTo((x1+xinc)*zoom+dx, (y1+yinc)*zoom+dy)

			angle += 120

			xinc = math.Cos(angle*3.14/180) * r
			yinc = math.Sin(angle*3.14/180) * r

			gc.LineTo((x1+xinc)*zoom+dx, (y1+yinc)*zoom+dy)

			gc.LineTo((x1)*zoom+dx, (y1)*zoom+dy)

			angle += 120

			xinc = math.Cos(angle*3.14/180) * r
			yinc = math.Sin(angle*3.14/180) * r

			gc.LineTo((x1+xinc)*zoom+dx, (y1+yinc)*zoom+dy)

			angle += 120

			xinc = math.Cos(angle*3.14/180) * r
			yinc = math.Sin(angle*3.14/180) * r

			gc.LineTo((x1+xinc)*zoom+dx, (y1+yinc)*zoom+dy)

			gc.FillStroke()
		}
	}

	log.Println("make image")

	targetFilename := fmt.Sprintf("%v/periodicresult/%v/%v_200.jpg", config.UploadPath, periodic, blueprint.Id)
	log.Println("targetFilename", targetFilename)
	global.SaveToJpegFile(targetFilename, dest)
}

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

func MakeMeterialImage(periodic int64, blueprint models.Blueprint, items []models.Periodicdata, iconZoom float64) {
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

	blue := color.RGBA{0x00, 0x00, 0xff, 0xff}
	red := color.RGBA{0xf4, 0x43, 0x37, 0xff}

	for _, v := range items {
		log.Println("type", v.Type)
		var results []global.Offset
		json.Unmarshal([]byte(v.Content), &results)

		if len(results) == 0 {
			continue
		}

		point := results[0]
		x := point.Dx
		y := point.Dy

		w := 4.0 * step / 50.0

		if w < 1 {
			w = 1
		}

		gc.SetLineWidth(w)

		if v.Type == 401 || v.Type == 402 {
			gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
			if v.Type == 401 {
				gc.SetStrokeColor(red)
			} else {
				gc.SetStrokeColor(blue)
			}

			gc.BeginPath()
			gc.ArcTo(x, y, step, step, 0, math.Pi*2)
			gc.FillStroke()

			if v.Type == 401 {
				gc.SetFillColor(red)
			} else {
				gc.SetFillColor(blue)
			}

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

		}

	}

	log.Println("make image")

	targetFilename := fmt.Sprintf("%v/periodicresult/%v/%v_400.jpg", config.UploadPath, periodic, blueprint.Id)
	log.Println("targetFilename", targetFilename)
	global.SaveToJpegFile(targetFilename, dest)
}

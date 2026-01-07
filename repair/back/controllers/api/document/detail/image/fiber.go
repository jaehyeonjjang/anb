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

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/font/gofont/goregular"
)

func MakeFiberImage(periodic int64, blueprint models.Blueprint, items []models.Periodicdata, iconZoom float64) {
	draw2d.SetFontFolder("./doc")

	os.Mkdir(fmt.Sprintf("%v/periodicresult/%v", config.UploadPath, periodic), 0755)

	filename := fmt.Sprintf("%v/%v", config.UploadPath, blueprint.Filename)
	img, _ := global.LoadImageFile(filename)

	if img == nil {
		log.Println("not found", filename)
		return
	}

	extra := 150
	dest, width, height := global.ImageToRGBAWithResize(img, extra, 0)
	gc := draw2dimg.NewGraphicContext(dest)

	step := iconZoom / 2.0
	stepMiddle := step * 0.66

	white := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	blue := color.RGBA{0x00, 0x00, 0xff, 0xff}
	red := color.RGBA{0xf4, 0x43, 0x37, 0xff}

	gc.BeginPath()
	gc.SetStrokeColor(white)
	gc.SetFillColor(white)
	gc.MoveTo(float64(width), 0)
	gc.LineTo(float64(width+extra), 0)
	gc.LineTo(float64(width+extra), float64(height))
	gc.LineTo(float64(width), float64(height))
	gc.LineTo(float64(width), 0)
	gc.FillStroke()
	gc.BeginPath()
	gc.SetFillColor(red)
	gc.SetStrokeColor(red)

	gc.SetFontSize(20)
	fontData := draw2d.FontData(draw2d.FontData{
		Name:  "NotoSans",
		Style: draw2d.FontStyleNormal,
	})

	// fontPath := "./doc/NotoSansKR-Regular.ttf"
	// fontBytes, err := os.ReadFile(fontPath)
	fontBytes := goregular.TTF
	log.Println("length", len(fontBytes))
	// if err != nil {
	// 	panic(err)
	// }
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("Failed to parse font: %v", err)
	}
	draw2d.RegisterFont(fontData, font)

	gc.SetFontData(fontData)

	gc.FillStringAt("ABC", float64(extra), 0)
	gc.FillStringAt("ABC", 0, 0)
	gc.Stroke()

	for _, v := range items {
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

		if v.Type == 301 {
			gc.BeginPath()

			gc.SetStrokeColor(red)
			gc.SetFillColor(red)

			gc.ArcTo(x, y, stepMiddle, stepMiddle, 0, math.Pi*2)

			gc.FillStroke()
		} else if v.Type == 302 {
			gc.BeginPath()

			gc.SetStrokeColor(blue)
			gc.SetFillColor(blue)

			gc.ArcTo(x, y, stepMiddle, stepMiddle, 0, math.Pi*2)

			gc.FillStroke()
		}

	}

	targetFilename := fmt.Sprintf("%v/periodicresult/%v/%v_300.jpg", config.UploadPath, periodic, blueprint.Id)
	global.SaveToJpegFile(targetFilename, dest)
}

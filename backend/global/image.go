package global

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func GetImageSize(filename string) (int, int) {
	imgfile, err := os.Open(filename)

	if err != nil {
		return 0, 0
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		return 0, 0
	}

	return imgCfg.Width, imgCfg.Height
}

func MakeThumbnail(w int, h int, filename string, targetFilename string) {
	imgfile, err := os.Open(filename)

	if err != nil {
		return
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		return
	}

	width := imgCfg.Width
	height := imgCfg.Height

	rate := float64(w) / float64(h)
	target := float64(width) / float64(height)

	newWidth := 0
	newHeight := 0

	x := 0
	y := 0

	if rate > target {
		newWidth = int(float64(h) * target)
		newHeight = h

		x = (w - newWidth) / 2
	} else if rate < target {
		newWidth = w
		newHeight = int(float64(w) / target)

		y = (h - newHeight) / 2
	} else {
		newWidth = w
		newHeight = h
	}

	src, err := imaging.Open(filename)
	if err != nil {
		return
	}

	src = imaging.Resize(src, newWidth, newHeight, imaging.Lanczos)

	dst := imaging.New(w, h, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, src, image.Pt(x, y))

	err = imaging.Save(dst, targetFilename)
}

func MakeThumbnailFault(w int, h int, originalW int, originalH int, filename string, targetFilename string) {
	imgfile, err := os.Open(filename)

	if err != nil {
		log.Println("MakeThumbnailFault file open error", filename)
		return
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		log.Println("MakeThumbnailFault decodeconfig error")
		return
	}

	width := imgCfg.Width
	height := imgCfg.Height

	rate := float64(w) / float64(h)
	target := float64(width) / float64(height)

	newWidth := 0
	newHeight := 0

	x := 0
	y := 0

	if rate > target {
		newWidth = int(float64(h) * target)
		newHeight = h

		x = (w - newWidth) / 2
	} else if rate < target {
		newWidth = w
		newHeight = int(float64(w) / target)

		y = (h - newHeight) / 2
	} else {
		newWidth = w
		newHeight = h
	}

	original, err := imaging.Open(filename)
	if err != nil {
		log.Println("MakeThumbnailFault imageing.open error")
		return
	}

	src := imaging.New(originalW, originalH, color.NRGBA{0, 0, 0, 0})
	src = imaging.Paste(src, original, image.Pt(x, y))

	src = imaging.Resize(src, newWidth, newHeight, imaging.Lanczos)

	dst := imaging.New(w, h, color.White)
	dst = imaging.Paste(dst, src, image.Pt(x, y))

	err = imaging.Save(dst, targetFilename)
}

func MakeThumbnailPicture(w int, h int, filename string, targetFilename string) {
	imgfile, err := os.Open(filename)

	if err != nil {
		return
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		return
	}

	width := imgCfg.Width
	height := imgCfg.Height

	rate := float64(w) / float64(h)
	target := float64(width) / float64(height)

	newWidth := 0
	newHeight := 0

	x := 0
	y := 0

	if rate > target {
		newWidth = w
		newHeight = int(float64(w) / target)

		y = (newHeight - h) / 2
	} else if rate < target {
		newWidth = int(float64(h) * target)
		newHeight = h

		x = (newWidth - w) / 2
	} else {
		newWidth = w
		newHeight = h
	}

	src, err := imaging.Open(filename)
	if err != nil {
		return
	}

	src = imaging.Resize(src, newWidth, newHeight, imaging.Lanczos)

	dst := imaging.New(w, h, color.White)
	dst = imaging.Paste(dst, src, image.Pt(-1*x, -1*y))

	err = imaging.Save(dst, targetFilename)
}

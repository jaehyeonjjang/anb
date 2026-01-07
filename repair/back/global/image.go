package global

import (
	"bufio"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"repair/global/log"

	"github.com/disintegration/imaging"
)

func GetImageSize(filename string) (int, int) {
	reader, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	bounds := m.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	return w, h
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

func ImageToRGBA(src image.Image) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

func ImageToRGBAWithResize(src image.Image, width int, height int) (*image.RGBA, int, int) {
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx()+width, b.Dy()+height))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst, b.Dx(), b.Dy()
}

type Offset struct {
	Dx float64 `json:"dx"`
	Dy float64 `json:"dy"`
}

func SaveToPngFile(filePath string, m image.Image) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create Writer from file
	b := bufio.NewWriter(f)
	// Write the image into the buffer
	err = png.Encode(b, m)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

func SaveToJpegFile(filePath string, m image.Image) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create Writer from file
	b := bufio.NewWriter(f)
	// Write the image into the buffer
	err = jpeg.Encode(b, m, nil)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

func LoadFromPngFile(filePath string) (image.Image, error) {
	// Open file
	f, err := os.OpenFile(filePath, 0, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b := bufio.NewReader(f)
	img, err := png.Decode(b)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func LoadFromJpegFile(filePath string) (image.Image, error) {
	// Open file
	f, err := os.OpenFile(filePath, 0, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b := bufio.NewReader(f)
	img, err := jpeg.Decode(b)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func LoadImageFile(filename string) (image.Image, error) {
	ext := filepath.Ext(filename)
	if ext == ".png" {
		return LoadFromPngFile(filename)
	} else if ext == ".jpg" || ext == ".jpeg" {
		return LoadFromJpegFile(filename)
	}

	return nil, errors.New("file format error")
}

func DownloadImage(url string, filename string) int64 {
	file, err := os.Create(filename)

	if err != nil {
		log.Error().Msgf("download image error: %v", url)
		return 0
	}

	defer file.Close()

	resp, err := http.Get(url)

	if err != nil {
		log.Error().Msg(err.Error())
		return 0
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		log.Error().Msg(err.Error())
		return 0
	}

	if size == 0 {
		os.Remove(filename)
	}

	return size
}

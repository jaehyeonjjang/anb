package main

import "github.com/fogleman/gg"

func main() {
    const S = 256
    dc := gg.NewContext(S, S)
    dc.SetRGB(1, 1, 1)
    dc.Clear()
    if err := dc.LoadFontFace("gulim.ttc", 36); err != nil {
        panic(err)
    }
    dc.SetRGB(0, 0, 0)
    s := "한글ABC"
    /*
    n := 3 // "stroke" size

    for dy := -n; dy <= n; dy++ {
        for dx := -n; dx <= n; dx++ {
            if dx*dx+dy*dy >= n*n {
                // give it rounded corners
                continue
            }
            x := S/2 + float64(dx)
            y := S/2 + float64(dy)
            dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
        }
    }
    */
    dc.SetRGB(0, 0, 1)
    dc.DrawString(s, S/2, S/2)
    dc.SavePNG("out.png")

    /*
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}

    dc.DrawString("테스트입니다", 10, 10)
	dc.SavePNG("out.png")
    */
}

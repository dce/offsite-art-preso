package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
)

const WIDTH = 800
const HEIGHT = 600

type Circle struct {
	x, y, radius int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	srcImg := image.NewRGBA(image.Rect(0, 0, WIDTH/2, HEIGHT/2))

	for x := 0; x < WIDTH/2; x++ {
		for y := 0; y < HEIGHT/2; y++ {
			setPixel(srcImg, x, y)
		}
	}

	drawText(srcImg)

	for x := 0; x < WIDTH/2; x++ {
		for y := 0; y < HEIGHT/2; y++ {
			l := rand.Intn(30) - 15

			srcImg.Set(x, y, lighten(srcImg.At(x, y).(color.RGBA), l))
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	for x := 0; x < WIDTH/2; x++ {
		for y := 0; y < HEIGHT/2; y++ {
			img.Set(2*x, 2*y, srcImg.At(x, y))
			img.Set(2*x+1, 2*y, srcImg.At(x, y))
			img.Set(2*x, 2*y+1, srcImg.At(x, y))
			img.Set(2*x+1, 2*y+1, srcImg.At(x, y))
		}
	}

	out, _ := os.Create("result.png")
	png.Encode(out, img)
	out.Close()
}

func setPixel(img *image.RGBA, x int, y int) {
	lblue := color.RGBA{R: 68, G: 146, B: 179, A: 255}
	dblue := color.RGBA{R: 21, G: 45, B: 62, A: 255}
	orange := color.RGBA{R: 222, G: 114, B: 56, A: 255}

	lg_circle := Circle{x: 165, y: 165, radius: 75}
	sm_circle := Circle{x: 290, y: 90, radius: 40}

	col := orange

	n := float64(x*y) / float64(WIDTH*HEIGHT/4) * 100

	if withinCircle(lg_circle, x, y) {
		col = lblue
	} else if onCircle(lg_circle, x, y) && rand.Intn(4) == 0 {
		col = lblue
	} else if withinCircle(sm_circle, x, y) {
		col = orange
	} else if onCircle(sm_circle, x, y) && rand.Intn(4) == 0 {
		col = orange
	} else if rand.Intn(100) > int(n) {
		col = dblue
	}

	img.Set(x, y, col)
}

func withinCircle(c Circle, x int, y int) bool {
	return distance(x, y, c.x, c.y) < float64(c.radius)
}

func onCircle(c Circle, x int, y int) bool {
	return math.Abs(distance(x, y, c.x, c.y)-float64(c.radius)) < 5
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func lighten(c color.RGBA, n int) color.RGBA {
	adjust := func(v uint8) uint8 {
		nv := int(v) + n

		if nv < 0 {
			nv = 0
		} else if nv > 255 {
			nv = 255
		}

		return uint8(nv)
	}

	return color.RGBA{R: adjust(c.R), G: adjust(c.G), B: adjust(c.B), A: 255}
}

func drawText(img *image.RGBA) {
	fontBytes, err := ioutil.ReadFile("./fonts/Victor Mono Regular Nerd Font Complete Mono.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(24)
	c.SetClip(img.Bounds())
	c.SetDst(img)

	c.SetSrc(image.Black)

	pt := freetype.Pt(11, 31)
	c.DrawString("Viget Dev Offsite", pt)

	c.SetFontSize(18)
	pt = freetype.Pt(11, 61)
	c.DrawString("Fall 2022", pt)

	c.SetFontSize(24)
	pt = freetype.Pt(66, 281)
	c.DrawString("“Friendship is Magic”", pt)

	c.SetSrc(image.White)

	pt = freetype.Pt(10, 30)
	c.DrawString("Viget Dev Offsite", pt)

	c.SetFontSize(18)
	pt = freetype.Pt(10, 60)
	c.DrawString("Fall 2022", pt)

	c.SetFontSize(24)
	pt = freetype.Pt(65, 280)
	c.DrawString("“Friendship is Magic”", pt)
}

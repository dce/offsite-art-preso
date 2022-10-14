package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const WIDTH = 800
const HEIGHT = 600

func main() {
	rand.Seed(time.Now().UnixNano())

	srcImg := image.NewRGBA(image.Rect(0, 0, WIDTH/2, HEIGHT/2))

	for x := 0; x < WIDTH/2; x++ {
		for y := 0; y < HEIGHT/2; y++ {
			setPixel(srcImg, x, y)
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
	LBLUE := color.RGBA{R: 68, G: 146, B: 179, A: 255}
	DBLUE := color.RGBA{R: 21, G: 45, B: 62, A: 255}
	ORANGE := color.RGBA{R: 222, G: 114, B: 56, A: 255}

	colors := [3]color.RGBA{ORANGE, LBLUE, DBLUE}

	col := colors[rand.Intn(3)]

	// lighten := uint8(255 * (float64(x) / float64(WIDTH)))

	img.Set(x, y, col)
}

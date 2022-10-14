package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const WIDTH = 800
const HEIGHT = 600

func main() {
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			setPixel(img, x, y)
		}
	}

	out, _ := os.Create("result.png")
	png.Encode(out, img)
	out.Close()
}

func setPixel(img *image.RGBA, x int, y int) {
	img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
}

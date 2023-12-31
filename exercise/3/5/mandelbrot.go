package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	file, _ := os.Create("./mandelbrot.png")

	defer file.Close()

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(file, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			if n > 50 {
				return color.RGBA{100, 0, 0, 255}
			}
			logScale := math.Log(float64(n)) / math.Log(float64(iterations))

			return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
		}
	}

	return color.Black
}

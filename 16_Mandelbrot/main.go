package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func Abs(c complex128) float64 {
	return math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
}

func analizeSequence(N int, x, y float64) bool {
	z := complex(0, 0)
	p := complex(x, y)
	for i := 1; i < N; i++ {
		z = z*z + p
		if Abs(z) >= 2.0 {
			return false
		}
	}
	return true
}

func visualizeData(size, N int) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{3 * size, 2 * size}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i < 2*size; i++ {
		y := 2.0*float64(i)/float64(2*size) - 1.0
		for j := 0; j < 3*size; j++ {
			x := 3.0*float64(j)/float64(3*size) - 2.0
			if analizeSequence(N, x, y) {
				img.Set(j, i, color.White)
			} else {
				img.Set(j, i, color.Transparent)
			}
		}

	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func main() {
	visualizeData(600, 200)
}

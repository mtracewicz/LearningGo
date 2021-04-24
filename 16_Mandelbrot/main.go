package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

func analyzeSequence(N int, x, y float64) uint8 {
	z := complex(0, 0)
	p := complex(x, y)
	for i := 1; i < N; i++ {
		z = z*z + p
		if cmplx.Abs(z) >= 2.0 {
			return uint8(255 * i / N)
		}
	}
	return 255
}

func createImage(size, N, cpuCount int) (time.Duration, image.Image) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{3 * size, 2 * size}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	runtime.GOMAXPROCS(cpuCount)
	//6*size*size is bringing 2D image of width 3*size and height 2*size into 1D
	workPerCore := 6 * size * size / cpuCount
	var wg sync.WaitGroup
	wg.Add(cpuCount)

	startTime := time.Now()
	for cpu := 0; cpu < cpuCount; cpu++ {
		go func(currentCpu int) {
			start := currentCpu * workPerCore
			stop := start + workPerCore

			for it := start; it < stop; it++ {
				i := it % (3 * size)
				j := it / (3 * size)

				x := 3.0*float64(i)/float64(3*size) - 2.0
				y := 2.0*float64(j)/float64(2*size) - 1.0
				a := analyzeSequence(N, x, y)
				img.Set(i, j, color.RGBA{R: 0xf0, G: 0x3a, B: 0x47, A: a})
			}
			wg.Done()
		}(cpu)
	}
	wg.Wait()
	stopTime := time.Since(startTime)

	return stopTime, img
}

func visualizeData(size, N, cpuCount int) {
	fmt.Println("Calculating mandelbrot set with parameters: size =", size, "accuracy =", N, "CPU count = ", cpuCount)
	_, img := createImage(size, N, cpuCount)
	f, _ := os.Create("image.png")
	png.Encode(f, img)
	fmt.Println("Compleated. Checkout image.png!")
}

func main() {
	size := flag.Int("size", 60, "Base on which image size will be calculated: Width=3*size, Height=2*size")
	accuracy := flag.Int("accuracy", 200, "How many iterations take place during calculation for single point")
	cpuCount := flag.Int("cpuCount", runtime.NumCPU(), "How many goroutines should be used")
	benchmarkMode := flag.Bool("benchmark", false, "If enabled no image will be generated, instead info about computation time will be provided.")
	flag.Parse()

	if *benchmarkMode {
		calculationsTime, _ := createImage(*size, *accuracy, *cpuCount)
		fmt.Println("size =", *size, "accuracy =", *accuracy, "cpuCount =", *cpuCount, "calculationsTime =", calculationsTime)
	} else {
		visualizeData(*size, *accuracy, *cpuCount)
	}
}

package main

import (
	"fmt"
	"math"
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

func visualizeData(size,N int) {
	for i := 0; i < 2*size; i++ {
		var y float64
		y = 2.0 * float64(i)/float64(2*size) - 1.0
		for j := 0; j < 3*size; j++ {
			var x float64
			x = 3.0 * float64(j)/float64(3*size) - 2.0
			if(analizeSequence(N,x,y)){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func main() {
	visualizeData(60,200)
}

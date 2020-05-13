package main

import (
	"fmt"
	"sync"
	"time"
)

func iterativeFib(n int) (fib int) {
	x := 0
	y := 1
	z := 1
	for i := 0; i < (n / 3); i++ {
		x = y + z
		y = x + z
		z = y + x
	}
	if n%3 == 0 {
		fib = x
	} else if n%3 == 1 {
		fib = y
	} else {
		fib = z
	}
	return
}
func iterativeFibV2(n int) int {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		y += x
		x = y - x
	}
	return x
}
func recursiveFib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return recursiveFib(n-1) + recursiveFib(n-2)
	}
}
func slicesFib(n int) int {
	var tab = []int{0, 1}
	for i := 2; i <= n; i++ {
		tab = append(tab, tab[i-1]+tab[i-2])
	}
	return tab[n]
}
func main() {
	n := 0
	fmt.Println("\033[96mWhich Fibonacci number?\033[0m")
	x, err := fmt.Scanf("%d", &n)
	if err != nil || x < 1 {
		fmt.Println("\033[91mNazleży podać liczbę naturalną!\033[0m")
		return
	}
	var wg sync.WaitGroup
	wg.Add(4)
	startI := time.Now()
	go func() {
		fmt.Printf("\033[92mIterative in place: Fib(%d) = %d it took me:\033[93m %s\n\033[0m", n, iterativeFib(n), time.Since(startI))
		wg.Done()
	}()
	startI2 := time.Now()
	go func() {
		fmt.Printf("\033[92mIterative v2 in place: Fib(%d) = %d it took me:\033[93m %s\n\033[0m", n, iterativeFibV2(n), time.Since(startI2))
		wg.Done()
	}()
	startR := time.Now()
	go func() {
		fmt.Printf("\033[92mRecursive: Fib(%d) = %d it took me: \033[94m%s\n\033[0m", n, recursiveFib(n), time.Since(startR))
		wg.Done()
	}()
	startS := time.Now()
	go func() {
		fmt.Printf("\033[92mIterative storing all earlier answers: Fib(%d) = %d it took me: \033[95m%s\n\033[0m", n, slicesFib(n), time.Since(startS))
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("\033[96mCompleted!")
}

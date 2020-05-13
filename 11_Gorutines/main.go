package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	quitT1 := make(chan bool)
	quitT2 := make(chan bool)
	ansT1 := make(chan int)
	ansT2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i:=0;;i++{
			select {
			case <-quitT1:
				ansT1 <- i
				wg.Done()
				return
			default:
				//d, _ := time.ParseDuration("2s")
				d, _ := time.ParseDuration("5ms")
				time.Sleep(d)
//				fmt.Println("T1")
			}
		}
	}()

	go func() {
		for i := 0; ;i++{
			select {
			case <-quitT2:
				ansT2 <- i
				wg.Done()
				return
			default:
				//d, _ := time.ParseDuration("1s")
				d, _ := time.ParseDuration("5ms")
				time.Sleep(d)
//				fmt.Println("T2")
			}
		}
	}()

	go func() {
		a := <-ansT1
		b := <-ansT2
		fmt.Printf("%d - %d = %d ms\n",5*b,5*a,5*(b-a))
		wg.Done()
		return
	}()

	_, err := fmt.Scanln()
	if err != nil {
		fmt.Println(err)
	} else {
		quitT1 <- true
	}
	_, err = fmt.Scanln()
	if err != nil {
		fmt.Println(err)
	} else {
		quitT2 <- true
	}
	wg.Wait()
}

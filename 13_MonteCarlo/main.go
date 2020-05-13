package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
	"math/rand"
)

func isInsideCircle(x,y,r float64)bool{
	return (x*x + y*y) < (r*r)
}

func countPiMonteCarloMethod(numberOfGoRoutine,howManyRandomNumbers int)float64{
	operationsPerThread := howManyRandomNumbers/numberOfGoRoutine
	var sum int
	var lock sync.Mutex
	var wg sync.WaitGroup
	seed := time.Now().UnixNano()
	wg.Add(numberOfGoRoutine)
	for i := 0;i < numberOfGoRoutine;i++{
		go func(){
			r := rand.New(rand.NewSource(seed))
			counter := 0
			for j := 0;j < operationsPerThread;j++{
				x := r.Float64()
				y := r.Float64()
				if(isInsideCircle(x,y,1.0)){
					counter++
				}
			}
			runtime.Gosched()
			lock.Lock()
			sum += counter
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return float64(4.0*float64(sum)/float64(howManyRandomNumbers))
}

func countPiAndExecutionTime(numberOfGoRoutine,howManyRandomNumbers int){
	startTime := time.Now()
	pi := countPiMonteCarloMethod(numberOfGoRoutine,howManyRandomNumbers)
	fmt.Printf("For %d routines: Pi = %f and Time = %s\n",numberOfGoRoutine,pi,time.Since(startTime));
}

func main(){
	N := 100000000000
	fmt.Printf("N = %d\n",N)
	countPiAndExecutionTime(16,N)
	countPiAndExecutionTime(4,N)
	countPiAndExecutionTime(2,N)
	countPiAndExecutionTime(1,N)
}
package main
import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func main(){
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var number int
	counter := 0
	waiting := make(chan bool)
	wg.Add(2)
	go func(){
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for ;; {
		select{
		case <- waiting:
			mutex.Lock()
			number = r.Int()
			waiting <- false
			mutex.Unlock()
		default:
		}
		if (counter == 5){
			wg.Done()
			break
		}
	}
	}()
	go func(){
		for ;;{
			select{
			case <- waiting:
				mutex.Lock()
				fmt.Println(number)
				waiting <- true
				mutex.Unlock()
				counter++
			default:
				waiting <- true
			}
			if (counter == 5){
				wg.Done()
				break
			}
		}
	}()
	wg.Wait()
}

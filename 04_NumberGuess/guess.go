package main
import (
	"fmt"
	"time"
	"math/rand"
)
func guess(randomNumber int){
	fmt.Printf("Take your guess: ");
	var guessedNumber int;
	fmt.Scanf("%d",&guessedNumber);
	if(guessedNumber == randomNumber){
		fmt.Println("Congratulations you won!")
	}else if(guessedNumber < randomNumber){
		fmt.Println("You went too small.")
		guess(randomNumber)
	}else{
		fmt.Println("You went too big.")
		guess(randomNumber)
	}

}
func main(){
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)
	fmt.Println("Number is between 0 and 100.")
	guess(randomNumber)
}

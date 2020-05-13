package main
import(
	"fmt"
	"math/rand"
	"time"
)
func rollDice()int{
	return (rand.Intn(6)+1)
}

func main(){
	rand.Seed(time.Now().UnixNano())
	answer := "y"
	for answer != "n"{
		if(answer == "y"){
			fmt.Println("Wylosowałeś: ",rollDice())
			fmt.Println("Losować jeszcze raz (y/n)")
		}else{
			fmt.Println("Musisz wybrać y albo n")
		}
		fmt.Scanf("%s",&answer)
	}
}

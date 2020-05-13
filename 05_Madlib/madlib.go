package main
import (
	"fmt"
	"strings"
)
func createStory(words []string)string{
	var story string
	var builder strings.Builder
	for _,element :=range words{
		builder.WriteString(element)
		builder.WriteString(" ")
	}
	story = builder.String()
	return story
}

func getInput()[]string{
	var userInput = []string{}
	userInput = append(userInput,promptForSingleWord("osoba"))
	userInput = append(userInput,"ubrała się w")
	userInput = append(userInput,promptForSingleWord("pzrymiotnik"))
	userInput = append(userInput,promptForSingleWord("rzeczownik")+".")
	userInput = append(userInput,promptForSingleWord("czasownik"))
	userInput = append(userInput,"zęby i poszła do")
	userInput = append(userInput,promptForSingleWord("miasto")+".")
	userInput = append(userInput,"Gdzie spotkała ")
	userInput = append(userInput,promptForSingleWord("zwierze")+".")
	userInput = append(userInput,"Tadam!")
	return userInput
}
func promptForSingleWord(what string)(input string){
	fmt.Println("Podaj ",what)
	fmt.Scanf("%s",&input)
	return
}

func main(){
	fmt.Println(createStory(getInput()))
}

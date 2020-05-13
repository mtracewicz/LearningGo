package main
import (
	"fmt"
	"strings"
)
func main(){
	var txt string;
	var tmp string;
	fmt.Println("Tell me the text you want to eradicate question marks from.")
	for tmp != "exit"{
		fmt.Scanf("%s",&tmp)
		txt += tmp
	}
	txt = strings.Replace(txt,"?","-",-1)
	fmt.Println(txt)
}

package messages

import "fmt"

// displays a message about the developer 
func Whoami() {
	name := "Henrique Guido"

	Hello()
	fmt.Println("*Developed by", name,"*")
}
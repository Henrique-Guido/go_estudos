package main

import "fmt"

func main() {

	name, age, email, pass := getData("Henrique Guido", 24, "henriqueguido.tech@gmail.com", true)
	
	fmt.Printf(" Nome: %v \n Idade: %v \n E-mail: %v \n Passe: %v", name, age, email, pass)

}
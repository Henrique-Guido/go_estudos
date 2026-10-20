package main

import "fmt"

func main() {

	// struct User
	type User struct {
		Name string
		Age int8
		Email string
		AgeCheck bool
	}

	// criação da variável do tipo
	cliente := User {
		Name: "Henrique Guido",
		Age: 24,
		Email: "henriqueguido.tech@gmail.com",
		AgeCheck: true,
	}

	fmt.Println(cliente.AgeCheck)

	// outras formas de definição de campos
	usuario2 := User{"Orlando", 68, "orlando@gmail.com", false}

	fmt.Println(usuario2.Name)

}
package main

import "fmt"

func main() {

	dev := User {
		name: "Henrique Guido",
		age: 24,
		email: "henriqueguido.tech@gmail.com",
	}

	fmt.Println("===== Ficha do Funça =====")
	fmt.Println("Nome:", dev.name,)
	fmt.Println("Idade:", dev.age)
	fmt.Println("E-mail:", dev.email)
}
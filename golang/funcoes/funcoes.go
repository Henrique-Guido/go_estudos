package main

import "fmt"

func id(name, email string) (string, string) {
	return name, email // para cada valor retornado, podemos usar uma variável diferente
}

func main() {

	name, email := id("Henrique Guido", "henrique@gmail.com")

	fmt.Println(name)
	fmt.Println(email)

}
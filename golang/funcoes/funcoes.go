package main

import "fmt"

// funções  // se tiver retorno precisa especificar o tipo do retorno
func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	result := somar(10, 100)
	fmt.Println(result)
}
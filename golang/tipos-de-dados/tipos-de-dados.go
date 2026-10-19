package main

import (
	"errors"
	"fmt"
)

func main() {
	var idade int = 240000
	fmt.Println(idade)

	//uint unsigned int (int sem sinal)
	var num uint = 200
	fmt.Println(num)

	// float32, float64
	var altura float32 = 1.70
	fmt.Println(altura)

	// string
	var nome string = "Henrique Guido"
	fmt.Println(nome)

	// boolean
	var booleano bool = true
	fmt.Println(booleano)

	// erro também é um tipo de dado em GO
	var erro error = errors.New("=== ERROR RUNNING THE APPLICATION ===")
	fmt.Println(erro)


}
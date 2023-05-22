package main

import (
	"errors"
	"fmt"
)
		
func main() {

	//numero inteiros
    var numero int64 = - 1000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	var numero3 rune = 0 
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)
	//fim numero inteiros


	// Números reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)
	//Fim numero reais

	//Strings
	var str string = "Números reais"
	fmt.Println(str)

	str2 := "Thalison"
	fmt.Println(str2)

	char := 'a'
	fmt.Println(char)

	//Fim strings

    var numero5 bool = true 
	fmt.Println(numero5)

	var texto string 
	fmt.Println(texto)

	var booleano1 bool = false
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro) 

}
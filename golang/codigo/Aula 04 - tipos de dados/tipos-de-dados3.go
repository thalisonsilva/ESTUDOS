package main

import (
	"errors"
	"fmt"
)
		
func main() {

	//numero inteiros 
	//int8, int16, int32, int64 -suporta número negativos 
    var numero int64 = - 1000000
	fmt.Println(numero)

	//uint8, uint16, uint32, uint64 - suporta número negativos - um int sem sinal (-)
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//E igual ao int32
	var numero3 rune = 0 
	fmt.Println(numero3)

	//Igual ao uint8
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
	fmt.Println(char) //Ele pega o número da tabela spa da cadeia de caracte

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
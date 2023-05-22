package main

import (
	"errors"
	"fmt"
)

//Os tipos de ponto flutuante representam os números “reais”, ou seja, aqueles que podem possuir uma parte fracionária.

func main() {
 
	//Ex: float32	±1,5 x 10−45 para ±3,4 x 1038	~6 a 9 dígitos	: 32 bits
	var numeroReal1 float32 = 230.60
	fmt.Println(numeroReal1)
	//Ex: float64	±5.0 × 10−324 para ±1.7 × 10308	~15 a 17 dígitos	64 bits
	var numeroReal2 float64 = 230.50
	fmt.Println(numeroReal2)

	//
	char := 'a' 
	fmt.Println(char)

	//No go todo valor inicial sempre será 0
	var numero int 
	fmt.Println(numero)

	var booleano1 bool = false
	fmt.Println(booleano1)

	//tipo de dados 
	var erro error = errors.New("Erro interno")
	fmt.Println(erro) 
}
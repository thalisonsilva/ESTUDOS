package main

import (
	"fmt"

)

//Podemos declara variaveis da seguintes formas!!
func main() {
	var nome1 string = "Thalison 1"   // Metodo 1 
	nome2 := "Thalison 2"			// Metodo 2 
	
	fmt.Println(nome1)
	fmt.Println(nome2)

	//Outro tipo de declaração de variaveis!! 

	var (
		variaveis3 string = "Outro tipo 1"  //Metodo 3
		variaveis4 string = "Outro tipo 2" //Metodo 4
	)
	fmt.Println(variaveis3, variaveis4)

	//Outro tipo de declaração de variaveis!!

	variaveis5, variaveis6 := "Outro tipo 3", "Outro tipo 4" //Metodo 5
	fmt.Println(variaveis5, variaveis6)

	//Outro tipo de declaração de variaveis!!

	const constantel string = "Outro tipo 5" //Metodo 6
	fmt.Println(constantel)

	//Enverte valores de variaveis (para fazer trocas de valores de variaveis)!! 
	variaveis5, variaveis5 = variaveis3, variaveis4
	fmt.Println(variaveis5, variaveis6)

	//Variaveis com números
	var idade int = 25
	ano := 1998
	salario := 1360.90
	fmt.Println(idade, ano)
	fmt.Println(salario)
}
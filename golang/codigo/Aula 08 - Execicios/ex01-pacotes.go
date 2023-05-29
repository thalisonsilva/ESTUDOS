package main

//Modlo e basicamente um conjutos de pacotes (centraliza todas a dependencia e compilat tudo em lugar só!)
import (
		"fmt"
	  	"modulo/auxiliar"
)

func main() { 
	fmt.Println("Olá thalison") 
	auxiliar.Escrever()
}
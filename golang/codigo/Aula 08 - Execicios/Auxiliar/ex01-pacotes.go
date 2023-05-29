package main

//Modlo e basicamente um conjutos de pacotes (centraliza todas a dependencia e compilat tudo em lugar só!)
import (
		"fmt"
	  	"modulo/auxiliar"
		"github.com/badoux/checkmail"  //utilizamos o comando "go get URL da importação"
)

func main() { 
	fmt.Println("Olá thalison") 
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devthalison@gmail.com")
	fmt.Println(erro)

}
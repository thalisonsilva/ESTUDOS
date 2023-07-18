package main // Todos o pacotes serão executados no pacote main!

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"

)


func main() {
	fmt.Println("Olá thalison")
	auxiliar.Escrever()  
	auxiliar.Escrever2()
	
	error :=  checkmail.ValidateFormat("contato@thalisonsilva.com.br")
	fmt.Println(error)
}



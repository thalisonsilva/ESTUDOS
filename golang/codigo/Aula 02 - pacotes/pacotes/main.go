package main

import (
	"modulo/auxiliar"
	"fmt" 
	"github.com/badoux/checkmail"
) 
	// Por padrão no go Toda função quer for exportada - deve ter um comentario o que ela faz!
func main() { 
	fmt.Println("Hello World")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("devthalison@gmail.com")
	fmt.Println(erro)
}
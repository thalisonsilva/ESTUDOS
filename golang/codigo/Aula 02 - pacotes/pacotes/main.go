package main

import (
	"modulo/auxiliar"
	"fmt" 
	"github.com/badoux/checkmail"
) 
	// Por padrão no GO toda função  exportada - deve ter um comentario explicando oque ela faz!
func main() { 
	fmt.Println("Hello World")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("devthalison@gmail.com")
	fmt.Println(erro)
}
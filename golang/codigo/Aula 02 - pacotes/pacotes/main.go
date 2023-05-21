package main

import (
	"modulo/auxiliar"
	"fmt" 
) 
	// Por padrão no go Toda função quer for exportada - deve ter um comentario o que ela faz!
func main() { 
	fmt.Println("Hello World")
	auxiliar.Escrever()
}
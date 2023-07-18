package main // Todos o pacotes serão executados no pacote main!

import (
	"fmt"
	"modulo/auxiliar"
)

// Toda função que será exportada, dever ter um comntario em cima!
func main() {
	fmt.Println("Olá thalison")
	auxiliar.Escrever()
	auxiliar.Escrever2()
}

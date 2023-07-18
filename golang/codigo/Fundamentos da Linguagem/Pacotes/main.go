package main // Todos o pacotes serão executados no pacote main!

import (
	"fmt"
	"modulo/auxiliar"
)

// Toda função que será exportada, dever ter um comntario em cima!
func main() {
	fmt.Println("Olá thalison")
	auxiliar.Escrever()  //Para cria mais de um pacote devemos cria mais de uma pastas.
	auxiliar.Escrever2()
}


// Quando estamos lidando com mais de um pacotes no GO devemos cria modulos para compilar! 
// comando: "go mod init modulo" -> você vai centralizar as depedenciar e pacotes.
// comando: "go build" -> você cria um "modulo.exe" um arquivo binario executavel.

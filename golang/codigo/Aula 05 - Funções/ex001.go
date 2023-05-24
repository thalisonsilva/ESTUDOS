package main

import "fmt"

//Funções sem parâmetros e sem retorno!!

// Sempre que escrevemos um programa em Go nós utilizamos esse tipo de função, pois sempre precisamos definir a função main.
func boasVindas() {
	fmt.Println("Boas-vindas ao meu programa escrito em Go!")
}
// Porém, fora do caso da função main esse tipo de função não é muito comum de ser utilizada.
// Pois possuem seu funcionamento estático, ou seja, independente de quando ou como sejam executadas sempre irão produzir o mesmo resultado.
func main() {
	boasVindas()
	boasVindas()
	boasVindas()
	boasVindas()
}

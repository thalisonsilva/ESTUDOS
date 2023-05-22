package main

import (
	"fmt"
)

//Regras de nomeação de variáveis e constantes!! 
func main() {
	//Padrão de formato camelCase (primeira palavra tudo munuscula, as demias a primeira letra maiuscula)
	//Usando o padrão em "variaveis = var" 
	var idadeDoAluno = 25   
	//Usando o padrão em "constante = const " 
	const idadeDoAluno2 = 30  

	fmt.Println(idadeDoAluno, idadeDoAluno2)
}
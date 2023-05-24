package main

import "fmt"
//Funções com parâmetros! 


func boasVindas(nome string) {
	fmt.Println("Bem vindo", nome)  //Os parâmetros de uma função são os dados que uma função precisa receber para possa funcionar corretamente.
}
func main() {
	boasVindas("Thalison")  //Os parâmetros de uma função são definidos entre os parênteses dessa função e podem ser acessados em seu corpo como se fossem simples variáveis.
}
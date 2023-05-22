package main

import  "fmt"


//No GO existe 4 tipos de número inteiros ex: "int8, int16, int32, int64"
//int: é um tipo inteiro que pode conter 32 ou 64 bits, a definição de qual o tamanho em memória será utilizado vai depender da arquitetura da máquina em que o código está sendo executado;
// byte: é um atalho para o tipo uint8;
func main() {
	 
	var numero int = 1000000000

	fmt.Println(numero)

		//uint: tem o mesmo comportamento do tipo int, mas só recebe valores sem sinal;
	var numero2 uint32 = 1000
	fmt.Println(numero2)

		// rune: é um atalho para o tipo int32, geralmente o tipo "rune" é utilizado para armazenar valores numéricos que representam caracteres Unicode.
		// int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)
}
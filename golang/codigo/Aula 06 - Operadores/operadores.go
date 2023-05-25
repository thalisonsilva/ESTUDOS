package main

import "fmt"

func main() {
	//Operadores de Aritimeticos
	//+ - * / %

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 4
	restoDaDivisao := 10 % 4

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//No Go para fazer operação de tipo diferente, ele não permite! 
	var numero1 int16 = 10 
	var numero2 int16 = 25

	soma1 := numero1 + numero2
	fmt.Println(soma1)
	//Fim operadores Aritimeticos

	//Atribuição
	var variavel1 string = "Thalison"
	variavel2 := "Thalison Moreira"
	fmt.Println(variavel1, variavel2)
	//Fim Atribuição

	//Operadores RELACIONAIS 
	fmt.Println(1 > 2) //Maior que
	fmt.Println(1 < 2) //Menor que
	fmt.Println(1 >= 2) //Menor ou igual
	fmt.Println(1 <= 2) //Menor ou igual
	fmt.Println(1 == 2) //Igual
	fmt.Println(1!= 2) //Não igual
	//Fim Operadores Relacionais

	//Operadores Lógicos
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) 
	fmt.Println(verdadeiro || falso)	
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//Fim Operadores Lógicos

	//Operadores UNÁRIOS
	fmt.Println("--------------")
	numero := 10
	numero++
	numero+= 15  //numero = numero + 15
	fmt.Println(numero)

	numero-- //
	numero-= 15  //numero = numero - 15
	numero *= 10 //numero = numero * 10
	numero /= 2  
	numero %= 2 
	fmt.Println(numero)
	//Fim Operadores UNÁRIOS

	//Operadores ternarios
	var texto string
	if numero > 10 {
			texto = "Maior que 10"
		} else {
			texto = "Menor que 10"
		}
		fmt.Println(texto)

}
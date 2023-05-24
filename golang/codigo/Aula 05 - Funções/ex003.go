package main

import "fmt" 

func main() {
	resultado := func(x ...int) func() int{
		res := 0
		for _, v := range x{
			res += v
		}
			return func() int {
				return res * res
			}
		}
		fmt.Println(resultado(1, 2, 3, 4, 5)())
}

// O result nomeado, você chama a variavel
func soma (a int, b int) (result int){	
	result =  a + b
	return

}

// Funções variadicas 

func somaTudo(x ...int) int { //Posso passa todo valores que quiser!
	resultado := 0

	for _, v := range x { // Vai fazer um loop de todos os valores colocando no x! 
        resultado += v
    }
	return resultado
}
// Funções anonimas 
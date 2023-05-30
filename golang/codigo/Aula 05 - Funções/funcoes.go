package main	

import "fmt"

// Funções e uma serie de passos que será seguidos pelo programa!
// São basicamente o dados que pomos para funcionar 
// Retono e o que a função retorna !!

func somar(n1 int8, n2 int8) int8{
	return n1 + n2 
}
func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2 
	subtracao := n1 - n2
	return soma, subtracao //No go a orde dos fatores autera o produto (Ela Importa)
}



func main() {
  soma := somar(10, 20)
  fmt.Println(soma)

  var f = func(txt string) string{
	  fmt.Println(txt)
	  return txt
	
  }
  resultado := f("Texto da função 1")
  fmt.Println(resultado)

  resultadosSoma, resultadoSubtracao:=calculosMatematicos(10, 15) // Lembre-se, no go toda variavel declarada deve ser usada! 
  fmt.Println(resultadosSoma, resultadoSubtracao)
}
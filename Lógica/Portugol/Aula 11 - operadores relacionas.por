programa
{
	logico x, y, z
	inteiro n1, n2
	
	funcao inicio()
	{
		escreva("digite um número: ")
		leia(n1)
		escreva("Digite outro número: ")
		leia(n2)

		//Operadores relacionais!! 

		x = n1 == n2 //igual

		escreva(" São valores iguais? ", x ,"\n")

		z = n1 > n2 //maior
		
		escreva(n1 , " é maior que ", n2 , " ? ", z, "\n")
		
		y = n1 != n2 // diferente 
		
		escreva("São diferentes? " , y)
	}
} 
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 405; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */
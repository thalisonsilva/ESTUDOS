programa         //O programa ela vai toma uma decisão (condicional simples)
{
	
	real n1, n2
	real media
	
	funcao inicio()
	{
		escreva("Informa a primeira nota: ")
		leia(n1)
		escreva("Informa a segunda nota: ")
		leia(n2)

		media = (n1 + n2)/2

		se (( media >= 1) e (media < 7)) {           // Nesse exemplos colocamos outra condição, se a media estiver entre 1 e 7 o aluno estara de recuperação!! 
			escreva("Você está de recuperação\n")
		} 
		escreva("Sua nota média foi: " + media)

	}
}

/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 19; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */
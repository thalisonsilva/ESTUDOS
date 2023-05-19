programa         //O programa ela vai toma uma decisão (condicional simples)
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

		se( media >= 7) {
			escreva("Você foi Aprovado\n")
		} 
		escreva("Sua nota média foi: " + media)

	}
}

/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 315; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */
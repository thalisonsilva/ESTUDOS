programa
{	
	real n1, n2 ,n3, n4
	real media 
	
	funcao inicio() //calculando a média de aluno 
	{
		escreva("Informe sua nota 01: ")
		leia(n1)
		escreva("Informe a nota 02: ")
		leia(n2)
		escreva("Informe a nota 03: ")
		leia(n3)
		escreva("Informe a nota 04: ")
		leia(n4)


		media = (n1 + n2 + n3 + n4)/4

		se (media >= 6.5) {
			escreva("Você foi APROVADO!!")
			escreva("Sua nota foi: " + media)
		} senao {
			escreva("Você foi REPROVADO!!\n")
			escreva("Sua nota foi: " + media)
		}
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 506; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */
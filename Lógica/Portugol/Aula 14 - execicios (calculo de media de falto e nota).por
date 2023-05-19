programa
{	
	real frequencia, media
	
	funcao inicio() //calculando a frequência do aulo 
	{
		escreva("Informe sua frequência: ")
		leia(frequencia)
		escreva("Informe sua média: ")
		leia(media)
	

		se (frequencia >= 0.75) {			
			escreva("Você foi APROVADO!!\n")	
		}senao {
			escreva("Você foi REPROVADO POR FALTA!!\n")		
		} se (media >= 7.0) {
			escreva("Você foi APROVADO!!\n")
			
		}senao {
			escreva("Você foi REPROVADO\n")
			
		}
		escreva("Sua média de frequencia foi: "+ frequencia)
		escreva("Sua média foi: " + media)
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 542; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */
programa
  
{
 
	
real numero= 10.0 //variáveis globais
	
	funcao inicio()
  
	{
		real numLocal = 12.0 //variáveis locais 

		escreva (numero +  " \n")
		escreva (numLocal + " \n")  // O ("\n") se refere o salto de linha
		escreva (calcula() + "\n")
	}
	funcao real calcula(){        // Podemos acessa a varável local em todo o escopo! 
							// Não podemos acessar a variavel fora do seu escopo de origem! 

		real numCalcula = 5.0 // variavel local
		retorne numCalcula * numero
		

		
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 371; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */

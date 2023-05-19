programa 
{		caracter  j1, j2, j3  //caracter porque o valor atribuido é uma letra apenas.
		logico estado  // podia apenas ser uma letra.
	
	funcao inicio()  //Operadores Lógicos 
	{
		j1 = 'a'  //'f' de fechado, 'a' de aberto
		j2 = 'f'
		j3 = 'f'

		escreva("jabela 01 aberta?: ", j1 == 'a')
		
		estado = j1  == 'a' ou j2 == 'a' ou j3 == 'a' //Todos os operadores são verdade
		
		escreva("\nAlguma janela aberta? " + estado)
		escreva("\nAlame desligado? " + nao estado) 

		estado = j1  == 'a' e j2 == 'a' e j3 == 'a'  // Todos precisa está (verdade) 
		escreva("\nTodas as jabelas estão abertas? " + estado)
		
		
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 625; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */
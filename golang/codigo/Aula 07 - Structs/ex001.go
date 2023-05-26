package main	

import "fmt"	

type Aluno struct {
	Nome string
	Idade int
	Turma int
}

func main() {    
	 
	Aluno:= Aluno{
		Nome: "João",
        Idade: 20,
        Turma: 1,
	}
	
    fmt.Println(Aluno) 
	// Outro jeito de chamar a função fmt.Println
    fmt.Println(Aluno.Nome)
	fmt.Println(Aluno.Idade)
	fmt.Println(Aluno.Turma)
}
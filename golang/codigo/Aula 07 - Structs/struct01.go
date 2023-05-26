package main

import "fmt"
// Structs e basicamente um selção de campos! 
// Usamos o Structs para salver dados
type usuarios struct {  // struct e basicamente um class em orientação ao objetos
	nome string
	idade int8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuarios 
	u.nome = "João"	
	u.idade = 18 
    fmt.Println(u)

	usuarios2 := usuarios{"Davi", 21}
	fmt.Println(usuarios2)

	usuarios3 := usuarios{nome: "João", idade: 18}
	fmt.Println(usuarios3)

}
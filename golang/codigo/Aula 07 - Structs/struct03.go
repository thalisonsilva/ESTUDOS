package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome string
	Email string
	CPF int

}
func (c Cliente) Exibe(){
	fmt.Println("Exibindo o nome do Cliente", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`

}

func main() {
    
 cliente := Cliente{
	Nome : "Thalison",
	Email : "test@example",
	CPF: 70460098136,
	}
	fmt.Println(cliente)

    cliente2 := Cliente{"Thalison Moreira", "test@example.com", 98691389109}
	fmt.Println("Nome: %s. Email: %s. CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
		Nome : "Thalison Moreira",
        Email : "test@example",
        CPF: 70460098136,
		},
		Pais: "Brasil",
	}
	 fmt.Println("Nome: %s. Email: %s. Pais %s\n", cliente3.Nome, cliente3.Email, cliente3.Pais)
	 cliente3.Exibe()
	 cliente2.Exibe()

	 clienteJson, err := json.Marshal(cliente3)
	 if err != nil {
         log.Fatal(err.Error())
     }
	 fmt.Println(string(clienteJson))
}
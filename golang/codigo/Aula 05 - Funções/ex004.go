package main

import "fmt"
// Utilização de metodos em GO, paracendo um orientação ao objetos!! 
type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "Andou")
}
func main() {
	carro := Carro {
		Nome: "BMW", 
	}
	carro.andar()
}
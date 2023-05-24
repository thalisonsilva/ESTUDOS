package main

import "fmt"

func mediaDeAluno(n1 int, n2 int, n3 int) int {
	soma := (n1 + n2 + n3) * 3
	return soma
}
func main() {

	resutadoSoma := mediaDeAluno(10, 20, 40)
	fmt.Println(resutadoSoma)
}

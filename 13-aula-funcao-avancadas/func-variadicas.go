package main

import "fmt"

//variadico e um slice no argumentos da func
// nao podemos ter mais de um argumento varidico
func variadcas(nu ...int) int {
	total := 0
	for _, valor := range nu {
		total += valor
	}
	return total
}

func main() {
	fmt.Println("func variadicas...")
	valor := variadcas(11, 36, 55, 55, 10)
	fmt.Println(valor)
}

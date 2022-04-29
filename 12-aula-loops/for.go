package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("aula de for loops...")

	var i int8 = 0
	for i < 3 {
		i++
		fmt.Println("incrementando i :")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 3; j++ {
		fmt.Println("valor de J : ", j)
		time.Sleep(time.Second)
	}
	//range interage com os dados inclusive string, retorna os numeros da tabela ascii
	//podemos excluir os indice colocando anderline
	array := [3]string{"allan", "joao", "lucas"}
	for _, nome := range array {
		fmt.Println(nome)
	}
	for indice, nome := range "allan" {
		fmt.Println(indice, string(nome))
	}
}

package main

import (
	"fmt"
)

//fuc nomeado e quando colocamos um nome para a declaracao de seu retorno
// nao necessitando de fazer um return assim : return var1, var2
func retorno_nomeado(n, n2 int) (nu int, nu2 int) {
	nu = n + 15
	nu2 = n2 - 10
	return
}

func main() {
	fmt.Println("retorno nomeado...")

	fmt.Println(retorno_nomeado(10, 20))

}

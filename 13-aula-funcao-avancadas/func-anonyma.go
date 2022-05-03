package main

import "fmt"

//func anonymo nao tem nome, e ja e executada logo apos a criacao

func main() {
	fmt.Println("func anonyma...")

	func() {
		fmt.Println("Func criada e executada automaticamente.")
	}()

	//com argumentos
	func(nu int) {
		res := nu + 20
		fmt.Println("Com argumentos -> ", res)
	}(5)

	//com retorno
	retorno := func(nu int) int {
		res := nu + 20
		return res
	}(8)
	fmt.Println("Com retorn -> ", retorno)
}

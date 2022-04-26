package main

import (
	"fmt"
)

//funcao simpes
func olar() {
	fmt.Println("olar funcao simples... ")
}

//com parametro
func olar2(paran string) {
	fmt.Println("olar func com", paran)
}

//com parametros e retorno
func calc(soma, soma2 int8) int8 {
	return soma + soma2
}

//com parametros e retornos multiplos
func calc2(soma, soma2 int8, texto string) (int8, string) {
	return soma - soma2, texto

}

func main() {
	//funcoa simples
	olar()

	//com parametro
	olar2("paranmetro")

	//com parametros e retorno
	fmt.Println(calc(10, 12))
	resu := calc(11, 15)
	fmt.Println(resu)

	//com parametros e retornos multiplos
	fmt.Println(calc2(25, 20, "valor do calculo "))
	res1, res2 := calc2(25, 10, "valor do calculo")
	fmt.Println(res2, res1)

	//passando uma func pra uma variavel
	var f = func() {
		fmt.Println("funcoe por variavel")
	}
	f()
	var ff = func(n1, n2 int, texto string) (int, string) {
		return n1 + n2, texto
	}
	re1, re2 := ff(50, 30, "Func por variavel 2")
	fmt.Println(re1, re2)
}

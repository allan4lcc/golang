package main

import (
	"errors"
	"fmt"
)

func main() {
	//int - aceita sinal de menos - int puro assume a arquitetura do pc
	var a int = 100
	var b int16 = 200
	var c int32 = 3000
	var d int64 = 30000
	fmt.Printf("Valores Int: \n%d\n%d\n%d\n%d\n", a, b, c, d)
	//int Fin

	//uint - nao aceita sinal de  menos ( - )
	var aa uint16 = 14
	fmt.Println(aa)

	// os tipos tem apelidos/alias - type rune = int32 - type byte = uint8
	var a1 rune = 10
	var a2 byte = 101
	fmt.Printf("Tipos alias/apelidos:\n%T\n%T\n", a1, a2)

	//float -
	var a3 float32 = 100.32
	var a4 float64 = 1000.35
	fmt.Printf("Tipod Float :\n%f\n%f\n", a3, a4)

	//String - posso usar apenas um caracter  de uma string como se fosse um char
	//representado por numero
	var a5 string = "Tipo String :\nqialquer coisa\n"
	//a6 := 'A'
	var a6 int8 = 'D'
	fmt.Println(a5, a6)

	//tipos de zeros, quando declara uma variavel com var sem atruir nenhum valor
	var str string
	var intt int
	var bo bool
	var err error
	var err2 error = errors.New("Erro interno golang")
	fmt.Printf("Valor de zer :\nString Vazia -> %s\n", str)
	fmt.Printf("int Vazia -> %d\n", intt)
	fmt.Printf("Bool Vazia -> %t\n", bo)
	fmt.Printf("Erro Vazia -> %T\n", err)
	fmt.Printf("Erro Vazia -> %s\n", err2)
}

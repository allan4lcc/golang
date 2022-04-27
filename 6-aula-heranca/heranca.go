package main

import (
	"fmt"
)

type pessoa struct {
	nome    string
	idade   uint8
	emprego string
	salario float32
}

type heraca struct {
	pessoa
	estuda string
	luta   string
	corre  float32
}

func main() {
	fmt.Println("olar heranca!!!")
	p := pessoa{"paloma", 15, "light", 1500.00}
	p2 := heraca{p, "Colegio Jardim Alvorada", "judo", 32.5}
	fmt.Printf("Acessando struct pessoa : %v\n", p2.emprego)
	fmt.Printf("Acessando struct heraca : %v\n", p2.estuda)
}

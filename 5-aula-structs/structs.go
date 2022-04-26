package main

import (
	"fmt"
)

type aluno struct {
	nome    string
	idade   uint8
	salario float32
}

//struct dentro de struct
type aluno2 struct {
	nome     string
	idade    uint8
	salario  float32
	getCarro carro
}
type carro struct {
	marca      string
	velocidade uint8
	preco      float32
}

//struct dentro fim

func main() {
	fmt.Println("Aula de structs ...")
	//chmando structs metodo 1
	var dados aluno
	dados.nome = "allan ferreira"
	dados.idade = 35
	dados.salario = 1500.00

	fmt.Println(dados.nome, dados.salario)

	//chmada 2
	dado := aluno{"allan santana", 35, 1700.00}
	fmt.Println(dado.idade, dados.salario)

	// chmada 3 - atrubir dados apenas a um elemento especifico do struct
	dad := aluno{nome: "Bruno souza"}
	fmt.Println(dad)

	//structs dentro de structs
	carro := carro{"Fiat", 120, 60.000}
	alu := aluno2{"joao lucas", 44, 1589.00, carro}
	fmt.Println(alu)
	fmt.Println(alu.nome)
	fmt.Println(alu.getCarro.marca, alu.getCarro.velocidade)
}

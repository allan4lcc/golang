package main

import (
	"fmt"
)

func main() {
	fmt.Println("array interno...")
	// tipo, tamanho, capacidade
	//make/array-interno se os dados colocado for maior que seu tamanho
	//ele dobra sua capacidade de armazenamento
	slace := make([]float32, 9, 10)
	fmt.Println("Tamnaho -> ", len(slace))
	fmt.Println("capacidade -> ", cap(slace))
	fmt.Println(slace)
	slace = append(slace, 12)
	slace = append(slace, 15)
	fmt.Println(slace)
	fmt.Println("Tamnaho Atualizado-> ", len(slace))
	fmt.Println("capacidade atuazada-> ", cap(slace))
}

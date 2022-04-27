package main

import "fmt"

func main() {
	fmt.Println("aula de ponteiro")

	// * cria ponteiro, e exibe valor de uma variavel que foi atribuida a um pontero
	// & atrubui a um ponteiro a posicao em memoria de uma variavel
	// & mostra posicao em memoria da variavel atual
	var ponteiro *int //guarda endereco de memoria do int
	var memoria_b int //guarda o valor de  int

	memoria_b = 100
	ponteiro = &memoria_b

	//repare, atribuir valor pelo ponteiro, a variavel (memoria_b foi alterado tambem)

	*ponteiro = 150
	fmt.Println("Agora valor de ponteiro : ", *ponteiro)
	fmt.Println("Agora valor de memoria_b : ", memoria_b)

}

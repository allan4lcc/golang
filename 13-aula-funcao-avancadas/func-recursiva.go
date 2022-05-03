package main

import "fmt"

//uma func que chama ela memso, nao e muito ultilizada
//como se fosse um loop da propria funcao
//nao a a necessidade de incremntar nenhum parmetro para o loop funcionar
//o if nu == 10{return nu}, e o freio que faz a func para com a repeticao

func recurs(nu int) int {
	if nu >= 10 {
		return nu
	}
	nu++
	fmt.Println("repetindo a func ", nu)
	return recurs(nu)
}
func main() {
	fmt.Println("func recurssiva...")
	get := recurs(5)
	fmt.Println(get)

}

package main

import "fmt"

func dia(dia int) string {
	switch dia {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda"
	case 3:
		return "Terca"

	case 4:
		return "Quarta"
	case 5:
		return "Quinta"

	case 6:
		return "Sexta"
	default:
		return "Escolha Invalida!!!"
	}
}

//segunda maneira de declarar switch
var semana string

func dia2(dia int) string {
	switch {
	case dia == 1:
		semana = "Domingo"
		//fallthrough sai da condicao atual e entra na proxima sem fazer comparacao
		fallthrough

	case dia == 2:
		semana = "Segunda"

	case dia == 3:
		semana = "Terca"

	case dia == 4:
		semana = "Quarta"

	case dia == 5:
		semana = "Quinta"

	case dia == 6:
		semana = "Sexta"

	default:
		semana = "Escolha Invalida!!!"
	}
	return semana
}
func main() {
	fmt.Println("aula de Switch ...")
	var dia string = dia(5)
	fmt.Println("O dia da semana e : ", dia)

	//chmando a segunda maneira de declaracao
	var dia2 string = dia2(1)
	fmt.Println("Dia2 da semana e : ", dia2)

}

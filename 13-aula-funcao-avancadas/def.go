package main

import "fmt"

//sem DEFER
func um1() {
	fmt.Println("func 1")
}

func dois2() {
	fmt.Println("func 2")
}

//com DEFER
func tres3() {
	fmt.Println("func 3")
}

func quatro4() {
	fmt.Println("func 4")
}

//DEFER - executa uma func depois de todas as outras
//se a func tiver um (retunr) ele execurata depois do return
//ele executa uma tarefa por ultimo de todas as outras nao somente func
func main() {
	fmt.Println("aula de - DEFER...")
	//sem DEFER
	um1()
	dois2()
	fmt.Println("...............................")
	//com DEFER
	defer tres3()
	quatro4()
}

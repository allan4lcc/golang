package main

import "fmt"

var vari int8 = 10
var vari2 *int8 = &vari

func tes1() {

	fmt.Println(vari)
}
func tes2() {

	fmt.Println(vari)
}

func main() {
	fmt.Println("aula de ponteiro")
	*vari2 = 20
	tes1()
	tes2()
	fmt.Println(vari)
}

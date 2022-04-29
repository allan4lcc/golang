package main

import (
	"fmt"
)

func main() {
	fmt.Println("aula de slace e array.")
	//array tem tamanho fixo nao e dinamico
	var array [3]int
	fmt.Println(array)
	array[0] = 10
	fmt.Println(array)

	array2 := [2]int{10, 25}
	fmt.Println(array2)

	array3 := [...]int8{25, 30, 50, 85, 59}
	fmt.Println(array3)
	//slice e dinamico
	var slice []string
	fmt.Println(slice)
	slice = append(slice, "allan")
	fmt.Println(slice)

}

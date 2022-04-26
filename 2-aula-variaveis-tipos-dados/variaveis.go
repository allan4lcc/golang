package main

import (
	"fmt"
)

func main() {
	//declarar variaveis
	var a string
	a = " olar mundo\n"

	b := "olar mundo 2\n"

	var (
		c string
		d string
	)
	c = "olar mundo 3\n"
	d = "olar mundo 4\n"

	var (
		e string = "olar mundo 5\n"
		f string = "olar mundo 6\n"
	)

	var g, h string
	g = "olar mundo 7\n"
	h = "olar mundo 8\n"

	var i, j string = "olar mundo 9\n", "olar mundo 10\n"

	l, m := "olar mundo 11\n", "olar mundo 12\n"
	fmt.Println(a, b, c, d, e, f, g, h, i, j, l, m)
}

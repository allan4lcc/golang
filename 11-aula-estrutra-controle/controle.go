package main

import (
	"fmt"
)

func main() {
	fmt.Println("estrutura de controle...")

	var nu int = 5

	if nu > 11 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	//if init, declara uma varavel somente no escopo do if/else
	// nu2 so pode ser acessada no escopo de if/else
	if nu2 := nu; nu2 < 5 {
		fmt.Println("maior")
	} else if nu2 == 5 {
		fmt.Println("Igual")
	} else {
		fmt.Println("menor")
	}

}

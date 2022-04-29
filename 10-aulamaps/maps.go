package main

import "fmt"

func main() {
	fmt.Println("aula de maps...")
	//maps sua chave tem um tipo e seus dados tambem

	user := map[string]string{
		"nome":  "pedro",
		"sobre": "morais",
	}

	fmt.Println(user["nome"])

	//podemos criar maps dentro de maps
	user2 := map[string]map[string]string{

		"carro2": {
			"nome": "chevrolet",
			"ano":  "1987",
		},

		"carro": {
			"nome":   "fiat",
			"modelo": "bmw",
		},
	}
	fmt.Println(user2["carro"]["nome"])

}

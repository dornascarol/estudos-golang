package main

import "fmt"

func main() {
	posicao := 3
	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}

	nome := "Péricles"
	switch nome {
	case "Dornas":
		fmt.Println("É meu nome")
	case "Bruce":
		fmt.Println("É o cachorro da minha amiga")
	case "Péricles":
		fmt.Println("É meu cantor brasileiro preferido")
	}
}

package main

import "fmt"

func main() {
	fmt.Println(soma1(42, 13))

	fmt.Println(soma2(18, 74))

	somaDosInteiros := soma1(31, 20)
	fmt.Println(somaDosInteiros)

	fmt.Println(subtracao(80, 20))

	subtracaoValores := subtracao(48, 10)
	fmt.Println(subtracaoValores)

	fmt.Println(printNome("Dornas"))

	nome := printNome("Carol")
	fmt.Println(nome)

	esporte1, esporte2 := printEsportes("Muay Thai")
	fmt.Println(esporte1)
	fmt.Println(esporte2)

	comida1, comida2, comida3, _ := printComida("Lasanha")
	fmt.Println(comida1)
	fmt.Println(comida2)
	fmt.Println(comida3)

	nomeCompleto, sobrenomeCompleto := printNomeCompleto("Andréa", "Figueiredo")
	fmt.Println(nomeCompleto)
	fmt.Println(sobrenomeCompleto)
}

func soma1(x int, y int) int {
	return x + y
}

func soma2(x int, y int) int {
	somaDosValores := x + y
	return somaDosValores
}

func subtracao(x int, y int) int {
	return x - y
}

func printNome(nome string) string {
	return nome
}

func printEsportes(esporte string) (string, string) {
	return esporte, esporte
}

func printComida(comida string) (string, string, string, string) {
	return comida, comida, comida, comida
}

func printNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// Função começando com letra minúscula --> função PRIVADA
// Função começando com letra maiúscula --> função PÚBLICA

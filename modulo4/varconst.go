package main

import "fmt"

func main() {
	// var + nome da variável + tipo
	// nome da variável + = (recebe) + valor atribuído
	var nome string
	nome = "carol"
	fmt.Println(nome)

	nome = "dornas"
	fmt.Println(nome)

	var idade int
	idade = 28
	fmt.Println(idade)

	// var + nome da variável + = (recebe) + valor atribuído
	var a = "livro"
	fmt.Println(a)

	var b = true
	fmt.Println(b)

	// var + nome da variável 1 + nome da variável 2 + tipo + = (recebe) + valores atribuídos
	var c, d int = 1, 2
	fmt.Println(c)
	fmt.Println(d)

	// nome da variável + := (declarar e atribuir) + valor atribuído
	e := "patins"
	fmt.Println(e)

	f := 1996
	fmt.Println(f)

	g := 1.56
	fmt.Println(g)
}

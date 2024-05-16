// Atrelar o método a um tipo específico de struct
// Quando atrelo o método a uma struct, eu tenho acesso a todos os campos desta struct

package main

import "fmt"

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

func (p Pessoa) ListaNomeEIdade() string {
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func main() {
	pessoa := Pessoa{
		Nome:      "Dornas",
		Idade:     28,
		Profissao: "Dev",
	}

	pessoa2 := Pessoa{
		Nome:      "Monica",
		Idade:     33,
		Profissao: "Analista financeiro",
	}

	println(pessoa.ListaNomeEIdade())
	println(pessoa2.ListaNomeEIdade())
}

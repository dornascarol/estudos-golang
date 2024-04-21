package main

import "fmt"

// type + nome da estrutura + struct + { Campos }
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Carol", "Dornas", 28})
	fmt.Println(Pessoa{Nome: "Andréa", Sobrenome: "Nascimento", Idade: 55})
	fmt.Println(Pessoa{Sobrenome: "Dornas", Idade: 28})

	p1 := Pessoa{Nome: "Maurilio", Sobrenome: "Dornas", Idade: 56}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Sobrenome)
	fmt.Println(p1.Idade)

	p1.Idade = 57
	fmt.Println(p1)

	p2 := Pessoa{Nome: "Andréa", Sobrenome: "Nascimento", Idade: 55}
	fmt.Println(p2)
	p3 := Pessoa{Nome: "Carol", Sobrenome: "Dornas", Idade: 28}
	fmt.Println(p3)
	p4 := Pessoa{Nome: "Hugo", Sobrenome: "Fonseca", Idade: 29}
	fmt.Println(p4)

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2, p3, p4)
	fmt.Println(pessoas)

	// structs + map
	alunosTi := map[string][]Pessoa{}
	alunosTi["programação"] = pessoas
	fmt.Println(alunosTi)

	var alunosCursos = map[string][]Pessoa{
		"Programação":   {{Nome: "Carol", Sobrenome: "Dornas", Idade: 28}},
		"Contabilidade": {{Nome: "Priscila", Sobrenome: "Vasconcellos", Idade: 25}},
	}
	fmt.Println(alunosCursos)

	// struct herdando campos de outra struct
	profis := Profissao{p4, "Software Engineer"}
	fmt.Println(profis)
	fmt.Println(profis.Pessoa.Nome, profis.Pessoa.Sobrenome)
	fmt.Println(profis.Tipo)
}

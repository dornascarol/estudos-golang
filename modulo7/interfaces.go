package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (c Cachorro) NumeroDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "miau"
}

func (g Gato) NumeroDePatas() int {
	return g.Patas
}

type Animal interface {
	Barulho() string
	NumeroDePatas() int
}

func FazBarulho(animal Animal) {
	fmt.Println(animal.Barulho())
}

type PessoaAdulta struct {
	Nome             string
	Idade            int
	Profissao        string
	TempoDeProfissao int
}

type Crianca struct {
	Nome  string
	Idade int
}

func (c Crianca) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia para você", c.Nome)
}

func (p PessoaAdulta) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia para você", p.Nome)
}

func (p PessoaAdulta) FalaProfissao() string {
	return fmt.Sprintf("%s tem %d anos como %s", p.Nome, p.TempoDeProfissao, p.Profissao)
}

type Adulto interface {
	FalaBomDia() string
	FalaProfissao() string
}

func BomDia(adulto Adulto) string {
	return adulto.FalaBomDia()
}

func main() {

	adulto := PessoaAdulta{
		Nome:             "Dornas",
		Idade:            28,
		Profissao:        "Desenvolvedora",
		TempoDeProfissao: 1,
	}

	crianca := Crianca{
		Nome:  "Serena",
		Idade: 5,
	}

	crianca.FalaBomDia()
	adulto.FalaBomDia()
	BomDia(adulto)

	cachorro := Cachorro{
		Raca:  "Spitz alemão",
		Cor:   "Preto",
		Patas: 4,
	}

	gato := Gato{
		Cor:   "Laranja, branco e preto",
		Patas: 4,
	}

	FazBarulho(cachorro)
	FazBarulho(gato)
}

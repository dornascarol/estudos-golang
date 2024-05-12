// Camel case => primeira letra minúscula e a primeira letra de cada nova palavra subsequente maiúscula. Só pode ser usada dentro do pacote que foi criada.
// Exemplo: var carolDornas string       (privada dentro do pacote)

// Pascal case => todas as palavras com a primeira letra maiúscula. Pode ser utilizada fora do pacote.
// Exemplo: var CarolDornas string       (pública fora do pacote)

package escola

type Aluno struct {
	Nome  string
	Email string
	Notas []int
}

type professor struct {
	nome     string
	email    string
	materias []string
}

func Notas() []int {
	return []int{10, 9, 8, 10}
}

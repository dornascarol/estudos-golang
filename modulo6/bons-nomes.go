// Escolher bons nomes para variáveis, funções e constantes é essencial para ter um código limpo e de fácil leitura
// Evitar repetir o nome do pacote
// Quando for índices, procure ser curto com apenas uma letra
// Quando for nomear pacotes, prefira nomes curtos e que sejam claros

package main

import "fmt"

func main() {
	var comidaItaliana = []string{"Pizza", "Ravioli", "Lasanha", "Risotto"}
	fmt.Println(comidaItaliana)

	var comidaPortuguesa = []string{"Bacalhau", "Pastéis de Belém", "Cozido à portuguesa", "Francesinha"}
	fmt.Println(comidaPortuguesa)

	var comidaFrancesa = []string{"Foie Gras", "Boeuf Bourguignon", "Ratatouille", "Crème brûlée"}
	fmt.Println(comidaFrancesa)

	var comidaCoreana = []string{"Kimchi", "Kimbap", "Tteokbokki", "Chikin coreano"}
	fmt.Println(comidaCoreana)
}

// Total amount of points
// Our football team has finished the championship.
// Our team's match results are recorded in a collection of strings. Each match is represented by a string in the format "x:y", where x is our team's score and y is our opponents score.

// Lógica:
// 01) Receber uma lista e percorrer com for range os elementos do placar.
// 02) Separar os pontos do nosso time e do adversário.
// 03) Entender se ganhamos, perdemos ou empatamos.
// 04) Domar os pontos acumulados por cada situação dos jogos.
// 05) Retorna a soma dos pontos.

package kata

import "strings"

func Points(games []string) int {
	var pontos int

	for _, game := range games {
		placar := strings.Split(game, ":")

		if placar[0] == placar[1] {
			pontos = pontos + 1
		}

		if placar[0] > placar[1] {
			pontos = pontos + 3
		} else {
			pontos = pontos + 0
		}

	}
	return pontos
}

// Sum of positive
// You get an array of numbers, return the sum of all of the positives ones.

// Lógica:
// 01) Pegar todos os números, para isso precisamos percorrer toda a lista => for range
// 02) Somar um a um
// 03) Retornar a soma

package kata

func PositiveSum(numbers []int) int {
	var soma int
	for _, number := range numbers {
		if number >= 0 {
			soma = soma + number
		}
	}
	return soma
}

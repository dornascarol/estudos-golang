// Count of positives / sum of negatives

// Given an array of integers.
// Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.
// If the input is an empty array or is null, return an empty array.

// Lógica:
// 01) Pegar todos os números, para isso precisamos percorrer toda a lista => for range
// 02) Se o número for maior do que zero, vamos contar a quantidade
// 03) Se o número for menor do que o zero, vamos somar os valores
// 04) Retornar uma lista: [quantidade dos positivos, soma dos negativos]

package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var quantidadePositiva int
	var somaNegativa int

	for _, number := range numbers {
		if number > 0 {
			quantidadePositiva = quantidadePositiva + 1
		}
		if number < 0 {
			somaNegativa = somaNegativa + number
		}
	}

	return []int{quantidadePositiva, somaNegativa}
}

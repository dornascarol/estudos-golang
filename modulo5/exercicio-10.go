// Cat years, Dog years
// I have a cat and a dog. I got them at the same time as kitten/puppy. That was humanYears years ago.
// Return their respective ages now as [humanYears,catYears,dogYears]

// Lógica:
// 01) Para o primeiro ano, o retorno será => [1, 15, 15]
// 02) Para o segundo ano, o retorno será => [2, 24, 24]
// 03) Para o terceiro ano, o retorno será => [3, (24+4), (24+5)]
// 04) Para o quarto ano, o retorno será => [4, (28+4), (29+5)]
// 05) Tirei os dois primeiros anos do cálculo

package kata

func CalculateYears(years int) (result [3]int) {
	if years == 1 {
		return [3]int{years, 15, 15}
	}
	if years == 2 {
		return [3]int{years, 24, 24}
	}

	catYears := 24 + (years-2)*4
	dogYears := 24 + (years-2)*5

	result = [3]int{years, catYears, dogYears}
	return
}

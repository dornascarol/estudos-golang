// Invert values
// Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

// Lógica:
// 01) Multiplicar por -1 => n * (-1)
// 02) Percorrer o array usando o for range
// 03) Usar o append para adicionar mais um item ao array

package kata

func Invert(arr []int) []int {
	var invertNumbers []int

	for _, number := range arr {
		invertNumbers = append(invertNumbers, number*(-1))
	}

	return invertNumbers
}

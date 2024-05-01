// Convert a Number to a String!
// We need a function that can transform a number (integer) into a string.

// Lógica:
// 01) Pesquisar a forma de converter o inteiro para string

package kata

import "strconv"

func NumberToString(n int) string {
	string := strconv.Itoa(n)

	return string
}

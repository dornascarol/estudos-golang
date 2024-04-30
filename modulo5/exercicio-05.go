// Convert a String to a Number!
// We need a function that can transform a string into a number. What ways of achieving this do you know?

// Lógica:
// 01) Pesquisar a forma de converter a string para inteiro
// 02) Ignorar o parâmetro que não vou usar

package kata

import "strconv"

func StringToNumber(str string) int {
	number, _ := strconv.Atoi(str)

	return number
}

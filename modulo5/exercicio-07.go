// String repeat
// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// Lógica:
// 01) Usar o for i := 0 ...
// 02) Criar uma variável para guardar os valores da string.

package kata

func RepeatStr(repetitions int, value string) string {
	var repeteString string

	for i := 0; i < repetitions; i++ {
		repeteString = repeteString + value
	}

	return repeteString
}

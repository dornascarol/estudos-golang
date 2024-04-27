// Summation
// Write a program that finds the summation of every number from 1 to num. The number will always be a positive integer greater than 0. Your function only needs to return the result, what is shown between parentheses in the example below is how you reach that result and it's not part of it, see the sample tests.

// Lógica:
// 01) Forma de incrementar os números de 1 até n
// 02) 1 + 2 + 3 + 4 + ... incrementando (somando)
// Exemplo: Quando i = 1 => 0 + 1 = 1
// Exemplo: Quando i = 2 => 1 + 2 = 3
// Exemplo: Quando i = 3 => 3 + 3 = 6
// Exemplo: Quando i = 4 => 6 + 4 = 10

package kata

func Summation(n int) int {
	var num int
	for i := 1; i <= n; i++ {
		num = num + i
	}
	return num
}

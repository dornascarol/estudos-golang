// Even or Odd
// Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

// Lógica:
// 01) Quando o resto da divisão por 2 for igual a zero, então o número é par
// 02) Operador de resto é o %

package kata

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

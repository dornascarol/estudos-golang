// Return Negative
// In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

//Lógica:
// 01) recebe um número
// 02) verificar se o número é negativo ou zero, e se for já retorna
// 03) se não for, transformar ele em negativo => 7 * (-1) = -7

package kata

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	} else {
		negativeNumber := x * (-1)
		return negativeNumber
	}
}

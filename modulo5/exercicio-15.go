// Odd-Even String Sort
// Given a string s, your task is to return another string such that even-indexed and odd-indexed characters of s are grouped and the groups are space-separated. Even-indexed group comes as first, followed by a space, and then by the odd-indexed part.

// Lógica:
// 01) Receber uma lista e percorrer com for i todas as letras das palavras. Vamos usar o tamanho da lista.
// 02) Para um número ser par o resto da divisão por dois tem que ser igual a zero.
// 03) Importar o pacote "fmt".
// 04) Usar o Sprintf para formatar a string do jeito que queremos.

package kata

import "fmt"

func SortMyString(s string) string {
	var letraPar string
	var letraImpar string

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			letraPar = letraPar + string(s[i])
		} else {
			letraImpar = letraImpar + string(s[i])
		}
	}

	return fmt.Sprintf("%s %s", letraPar, letraImpar)
}

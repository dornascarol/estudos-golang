// Parts of a list
// Write a function partlist that gives all the ways to divide a list (an array) of at least two elements into two non-empty parts.

// Lógica:
// 01) Receber uma lista e percorrer com for i. Vamos usar o tamanho da lista.
// 02) Usar o Sprintf para formatar a string do jeito que queremos.
// 03) Importar os pacotes "fmt" e "strings".
// 04) Usar o Join para concatenar as strings.

package kata

import (
	"fmt"
	"strings"
)

func PartList(arr []string) string {
	var finalWord string

	for i := 1; i < len(arr); i++ {
		word := fmt.Sprintf("(%s, %s)", strings.Join(arr[:i], " "), strings.Join(arr[i:], " "))
		finalWord = finalWord + word
	}

	return finalWord
}

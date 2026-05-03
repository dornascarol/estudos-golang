package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func removerAcentos(s string) string {
    t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
        return unicode.Is(unicode.Mn, r)
    }), norm.NFC)
    resultado, _, _ := transform.String(t, s)
    return resultado
}

func contarLetras(titulo string) map[string]int {

    titulo = strings.ToLower(titulo)
    titulo = removerAcentos(titulo)

    contagem := make(map[string]int)

    for _, letra := range titulo {
        if letra == ' ' {
            continue
        }
        contagem[string(letra)]++
    }

    return contagem
}

func main() {
    titulo := "iPhone Pró"
    resultado := contarLetras(titulo)
    fmt.Println(resultado)
}

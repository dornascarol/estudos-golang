package main

import "fmt"

func encontrarDuplicatas(titulos []string) []string {

    contagem := make(map[string]int)

    for _, titulo := range titulos {
        contagem[titulo]++
    }

    duplicatas := []string{}

    for titulo, quantidade := range contagem {
        if quantidade > 1 {
            duplicatas = append(duplicatas, titulo)
        }
    }

    return duplicatas
}

func main() {
    titulos := []string{"iPhone 13", "Notebook Dell", "iPhone 13", "Galaxy S22", "Notebook Dell", "iPhone 13"}
    fmt.Println(encontrarDuplicatas(titulos))
}

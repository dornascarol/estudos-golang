// Um ponteiro nada mais é do que uma variável que ao invés de armazenar um valor, ela armazena um endereço de memória
// Memória => essa memória tem um endereço -> esse endereço guarda um valor
// &var => apontando para o endereço de memória da variável
// Quando uso *, estou falando do endereço e não do valor

package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Valor inicial: ", i)
	fmt.Println("Valor do endereço de memória: ", &i)

	a := &i
	fmt.Println("Variável a: ", a)
	fmt.Println("Variável *a: ", *a)
	fmt.Println("Variável &a: ", &a)
}

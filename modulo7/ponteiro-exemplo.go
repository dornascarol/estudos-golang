// Um ponteiro nada mais é do que uma variável que ao invés de armazenar um valor, ela armazena um endereço de memória
// Memória => essa memória tem um endereço -> esse endereço guarda um valor
// &var => apontando para o endereço de memória da variável
// Quando uso *, estou falando do endereço e não do valor

package main

import "fmt"

func zeroValue(b int) {
	b = 0
	fmt.Println("Endereço de memória do b dentro func zeroValue: ", &b)
}

func zeroPointer(b *int) {
	*b = 0
}

func main() {
	b := 1
	fmt.Println("Valor inicial de b dentro main: ", b)
	fmt.Println("Valor endereço de memória de b dentro main: ", &b)

	zeroValue(b)
	fmt.Println("zeroValue dentro de main: ", b)

	zeroPointer(&b)
	fmt.Println("zeroPointer dentro de main: ", b)

	fmt.Println("Pointer endereço de memória de b dentro main: ", &b)
}

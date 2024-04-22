// IF => se
// ELSE => senão

package main

import "fmt"

func main() {
	numero := 7
	// if + condição + { ação } + else + { outra ação }
	if numero == 1 {
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor não é igual a 1")
	}

	if numero == 2 {
		fmt.Println("Valor é igual a 2")
	} else if numero == 8 {
		fmt.Println("Valor é igual a 8")
	} else {
		fmt.Println("Valor é diferente de 2 e 8")
	}

	fmt.Println(46 + 92)
	fmt.Println(55 - 34)
	fmt.Println(2 * 3)
	fmt.Println(10 / 2)
	fmt.Println(10 % 2)

	if numero%2 == 0 {
		fmt.Printf("%d é par\n", numero)
	} else {
		fmt.Printf("%d é impar\n", numero)
	}

	numero2 := 8
	if numero2%2 == 0 {
		fmt.Printf("%d é par\n", numero2)
	} else {
		fmt.Printf("%d é impar\n", numero2)
	}

	usuario := "Dornas"
	if usuario == "Dornas" {
		fmt.Println("Usuário logado!")
	} else {
		fmt.Println("Tente novamente um usuário válido.")
	}
}

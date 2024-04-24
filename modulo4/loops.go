// Laços de repetição => para repetir tarefas
// i++ => somando mais 1 => quer dizer: i = i + 1
// i-- => subtraindo menos 1 => quer dizer: i = i - 1
// sum += i => quer dizer: sum = sum + i
// sum -= i => quer dizer: sum = sum - i
// loop infinito
// for range

package main

import (
	"fmt"
)

func main() {
	soma1 := 0
	for i := 0; i < 10; i++ {
		soma1 += i
	}
	fmt.Println(soma1)

	soma2 := 0
	for i := 90; i < 100; i++ {
		soma2 += i
	}
	fmt.Println(soma2)

	//for {
	//	fmt.Println("Golang do zero")
	//	time.Sleep(2 * time.Second)
	//}

	frutas := []string{"Laranja", "Maça", "Banana", "Uva"}
	for i, fruta := range frutas {
		fmt.Println("Fruta:", fruta, "|", "índice:", i)
	}

	comidasItalianas := []string{"Pizza", "Macarrão", "Lasanha"}
	for _, pastas := range comidasItalianas {
		fmt.Println(pastas)
	}
}

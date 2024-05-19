package main

import (
	"calculator/math"
	"fmt"
)

func main() {
	soma := math.Soma(3, 7)
	fmt.Println(soma)

	subtracao := math.Subtracao(22, 15)
	fmt.Println(subtracao)
}

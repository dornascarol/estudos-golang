package main

import (
	"calculator/math"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	soma := math.Soma(3, 7)
	fmt.Println(soma)

	subtracao := math.Subtracao(22, 15)
	fmt.Println(subtracao)
}

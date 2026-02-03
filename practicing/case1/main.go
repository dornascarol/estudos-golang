package main

import "fmt"

func count(words []string) map[string]int {
	result := make(map[string]int)

	for _, wd := range words {
		result[wd]++
	}

	return result
}

func main() {
	input := []string{"usd", "brl", "usd", "eur", "usd", "brl"}

	output := count(input)

	fmt.Println(output)
}

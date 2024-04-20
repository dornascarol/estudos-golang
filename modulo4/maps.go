// Listas

// 2) Maps - heterogêneos (pode misturar os tipos)
// var := map[key]value{}

package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["dornas"] = 28
	idade["andrea"] = 55
	fmt.Println(idade)
	fmt.Println(idade["dornas"])
	fmt.Println(idade["andrea"])

	anoNasc := map[string]int{
		"dornas": 1996,
		"andrea": 1968,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["dornas"])
	fmt.Println(anoNasc["andrea"])

	anoNasc["maurilio"] = 1966
	anoNasc["curso golang"] = 2024
	fmt.Println(anoNasc)

}

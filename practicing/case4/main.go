package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complemento := target - num

		if index, found := seen[complemento]; found {
			return []int{index, i}
		}

		seen[num] = i
	}

	return nil 
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9

    result := twoSum(nums, target)

    fmt.Println("Índices encontrados:", result)
}

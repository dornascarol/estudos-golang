package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(number int) []string {
    result := make([]string, 0, number)

    for i := 1; i <= number; i++ {
        if i%3 == 0 && i%5 == 0 {
            result = append(result, "FizzBuzz")
        } else if i%3 == 0 {
            result = append(result, "Fizz")
        } else if i%5 == 0 {
            result = append(result, "Buzz")
        } else {
            result = append(result, strconv.Itoa(i))
        }
    }

    return result
}

func main() {
    output := fizzBuzz(19)
    fmt.Println(output)
}

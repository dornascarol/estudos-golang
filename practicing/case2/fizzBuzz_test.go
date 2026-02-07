package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected []string
    }{
        {
            name:  "n igual a 3",
            input: 3,
            expected: []string{
                "1", "2", "Fizz",
            },
        },
        {
            name:  "n igual a 5",
            input: 5,
            expected: []string{
                "1", "2", "Fizz", "4", "Buzz",
            },
        },
        {
            name:  "n igual a 15",
            input: 15,
            expected: []string{
                "1", "2", "Fizz", "4", "Buzz",
                "Fizz", "7", "8", "Fizz", "Buzz",
                "11", "Fizz", "13", "14", "FizzBuzz",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := fizzBuzz(tt.input)

            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf(
                    "esperado %v, mas obteve %v",
                    tt.expected,
                    result,
                )
            }
        })
    }
}

package main

import "testing"

func TestCount(t *testing.T) {
    input := []string{"usd", "brl", "usd", "eur", "usd", "brl"}

    expected := map[string]int{
        "usd": 3,
        "brl": 2,
        "eur": 1,
    }

    result := count(input)

    for key, expectedValue := range expected {
        if result[key] != expectedValue {
            t.Errorf("expected %d for %s, but got %d",
                expectedValue,
                key,
                result[key],
            )
        }
    }
}

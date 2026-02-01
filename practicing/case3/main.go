package main

import "fmt"


func IsAnagram(word1, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	count := make(map[rune]int)

	for _, letter := range word1 {
		count[letter]++
	}

	for _, letter2 := range word2 {
		count[letter2]--
		if count[letter2] < 0 {
			return false
		}
	}

	return true
}

func main() {

	word1 := "UVA"

	word2 := "AVU"

	result := IsAnagram(word1, word2)
	fmt.Println("Are the words anagrams?", result)
}

//    word1 = UVA       word2 = AVU
//      U,V,A             A,V,U	
//     [1,1,1]           [0,0,0]

// Do jeito que está, a função é case-sensitive. 
// Se o requisito for ignorar maiúsculas e minúsculas, eu normalizaria as strings antes de processar, por exemplo usando 'strings.ToLower'

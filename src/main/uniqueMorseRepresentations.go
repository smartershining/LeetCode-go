package main

import (
	"fmt"
	"strings"
)

func main()  {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

var table = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
func uniqueMorseRepresentations(words []string) int {
	res := make(map[string]bool, len(words))
	for _, word := range words{
		var b strings.Builder
		for i := 0; i < len(word); i++ {
			b.WriteString(table[word[i] - 'a'])
		}
		res[b.String()] = true
	}
	return len(res)
}


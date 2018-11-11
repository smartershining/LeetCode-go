package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(93, 73))

}
func hammingDistance(x int, y int) int {
	result, count := x^y, 0
	for result > 0 {
		count += result & 1
		result >>= 1
	}
	return count
}

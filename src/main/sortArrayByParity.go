package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{0, 2}))
}
func sortArrayByParity(A []int) []int {
	i, j := 0, len(A) - 1
	for i < j {
		for i < j && A[i]%2 == 0 {
			i++
		}
		for i < j && A[j]%2 == 1 {
			j--
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
	}
	return A
}

package main

import "fmt"

func main() {
	fmt.Print(media([]int{5, 2, 7, 8, 3, 4}))
}

func media(numeros []int) int {
	soma := 0
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}
	return soma / len(numeros)
}

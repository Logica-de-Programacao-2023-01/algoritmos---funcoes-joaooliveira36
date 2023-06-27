package main

import "fmt"

func main() {
	fmt.Print(Posiçãonum([]int{1, 3, 2, 7, 5, 4, 8, 6}, 5))
}

func Posiçãonum(lista_num []int, numero_desejado int) int {
	for i, numero := range lista_num {
		if numero == numero_desejado {
			return i
		}
	}
	return -1
}

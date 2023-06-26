package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := []int{1, 10, 9, 6, 4, 5, 3, 7, 8}
	ordenação, err := ordem_crescente(numeros)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(ordenação)
	}
}

func ordem_crescente(numeros []int) ([]int, error) {
	if len(numeros) == 0 {
		return nil, fmt.Errorf("Sua slice está vaziaa")
	}

	sort.Ints(numeros)

	return numeros, nil
}

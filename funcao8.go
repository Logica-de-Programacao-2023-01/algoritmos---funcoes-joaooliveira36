package main

import "fmt"

func main() {
	numeros := []int{2, 3, 7, 4, 6, 9, 4, 1, 5}
	num_pares, err := num_par(numeros)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(num_pares)
	}
}
func num_par(numeros []int) ([]int, error) {
	lista_pares := []int{}

	if len(numeros) == 0 {
		return []int{}, fmt.Errorf("Sua lista está vazia")
	}

	for _, num := range numeros {
		if num%2 == 0 {
			lista_pares = append(lista_pares, num)
		}
	}

	return lista_pares, nil
}

package main

import "fmt"

func main() {
	lista1 := []int{2, 1, 6, 4, 7}
	lista2 := []int{5, 3, 2, 7, 4}
	numeros_iguais, err := numeros_iguais(lista1, lista2)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(numeros_iguais)
	}
}

func numeros_iguais(lista1 []int, lista2 []int) ([]int, error) {
	if len(lista1) == 0 || len(lista2) == 0 {
		return nil, fmt.Errorf("Alguma de suas listas está vazia")
	}
	num_coinhecidentes := []int{}

	for _, numero1 := range lista1 {
		for _, numero2 := range lista2 {
			if numero1 == numero2 {
				num_coinhecidentes = append(num_coinhecidentes, numero1)
				break
			}
		}
	}

	return num_coinhecidentes, nil
}

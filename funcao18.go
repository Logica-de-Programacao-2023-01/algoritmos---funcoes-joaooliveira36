package main

import "fmt"

func main() {
	numeros := []int{5, 3, 2, 7, 6, 8, 4}
	soma, err := função2(numeros, função1)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(soma)
	}
}
func função1(numero int) int {
	return numero * numero
}

func função2(numeros []int, funcao func(int) int) (int, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("Sua lista está vazia")
	}
	soma_resultado := 0

	for _, numero := range numeros {
		soma_resultado += funcao(numero)
	}

	return soma_resultado, nil
}

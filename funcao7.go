package main

import (
	"fmt"
)

func main() {
	lista_numeros := []int{2, 5, 9, 4, 6, 3, 7}
	lista_duplicada, err := aplicar_multiplicação(lista_numeros, mult_por_2)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(lista_duplicada)
	}
}
func mult_por_2(num int) int {
	return num * 2
}

func aplicar_multiplicação(numeros []int, funcao func(int) int) ([]int, error) {
	if len(numeros) == 0 {
		return nil, fmt.Errorf("Lista de numeros vazia")
	}
	resultado := []int{}
	for _, numero := range numeros {
		resultado = append(resultado, funcao(numero))
	}
	return resultado, nil
}

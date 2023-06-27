package main

import "fmt"

func main() {
	numero := 9
	lista_numeros := []int{2, 5, 6, 3, 9}
	verificação, err := number_check(numero, lista_numeros)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(verificação)
	}

}
func number_check(numero int, lista []int) (bool, error) {
	if len(lista) == 0 {
		return false, fmt.Errorf("Sua lista de números está vazia")
	}
	for _, numeros := range lista {
		if numeros == numero {
			return true, nil
		}
	}
	return false, nil
}

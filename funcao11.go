package main

import (
	"fmt"
)

func main() {
	numero := 11
	check_primos, err := numeros_primos(numero)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(check_primos)
	}

}

func numeros_primos(numero int) (bool, error) {
	if numero <= 0 {
		return false, fmt.Errorf("Seu número é negativo")
	}
	if numero <= 1 {
		return false, fmt.Errorf("Seu número é igual a 1")
	}
	if numero == 2 {
		return false, nil
	}

	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

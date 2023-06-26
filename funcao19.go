package main

import "fmt"

func main() {
	numero := 15
	primos, err := menores_primos(numero)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(primos)
	}
}

func menores_primos(numero int) ([]int, error) {
	if numero < 0 {
		return nil, fmt.Errorf("Seu número é menor que 0")
	}
	numeros_primos := []int{}

	for i := 2; i <= numero; i++ {
		isPrimo := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrimo = false
				break
			}
		}
		if isPrimo {
			numeros_primos = append(numeros_primos, i)
		}
	}

	return numeros_primos, nil
}

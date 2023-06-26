package main

import (
	"fmt"
	"strconv"
)

func main() {
	numero := 15
	soma, err := soma_digitos(numero)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(soma)
	}
}

func soma_digitos(numero int) (int, error) {
	soma := 0
	digitos := strconv.Itoa(numero)

	if numero < 0 {
		return 0, fmt.Errorf("Seu número é menor que 0")
	}

	for _, digito := range digitos {
		numero, error := strconv.Atoi(string(digito))
		if error != nil {
			return 0, fmt.Errorf("Tem um erro no seu número")
		}
		soma += numero
	}

	return soma, nil
}

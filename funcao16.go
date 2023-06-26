package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra := "senhorita"
	substituição, err := vogais_check(palavra)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(substituição)
	}
}

func vogais_check(palavra string) (string, error) {
	if len(palavra) == 0 {
		return "", fmt.Errorf("Sua palavra tem tamanho 0")
	}
	vogais := "AEIOUaeiou"
	vogais_substituidas := ""

	for _, caracter := range palavra {
		if strings.Contains(vogais, string(caracter)) {
			vogais_substituidas += "*"
		} else {
			vogais_substituidas += string(caracter)
		}
	}
	return vogais_substituidas, nil
}

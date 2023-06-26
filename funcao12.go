package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	lista := []string{"Ola", "coMo", "voce", "Esta"}
	palavras, err := letras_maiusculas(lista)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(palavras)
	}
}

func letras_maiusculas(palavras []string) (string, error) {
	var palavras_maiusculas []string

	if len(palavras) == 0 {
		return "", fmt.Errorf("Sua lista está vazia")
	}

	for _, palavra := range palavras {
		for _, letra := range palavra {
			if unicode.IsUpper(letra) {
				palavras_maiusculas = append(palavras_maiusculas, palavra)
				break
			}
		}
	}

	if len(palavras_maiusculas) == 0 {
		return "", fmt.Errorf("Não tem palavras com letras maiusculas na sua lista")
	}
	lista_final := strings.Join(palavras_maiusculas, ", ")

	return lista_final, nil
}

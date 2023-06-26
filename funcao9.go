package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Olá como você está?"
	palavras, err := palavras(frase)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(palavras)
	}
}

func palavras(string string) ([]string, error) {
	if len(string) == 0 {
		return nil, fmt.Errorf("Sua string está vazia!")
	}
	palavras := strings.Split(string, " ")

	return palavras, nil
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	lista_palavras := []string{"ola", "como", "vai"}
	junção, err := junção_palavras(lista_palavras)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(junção)
	}
}

func junção_palavras(palavras []string) (string, error) {
	if len(palavras) == 0 {
		return "", fmt.Errorf("Sua string esta vazia")
	}
	palavras_com_virgula := strings.Join(palavras, ",")

	return palavras_com_virgula, nil
}

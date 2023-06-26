package main

import (
	"fmt"
	"sort"
)

func main() {
	palavras := []string{"marcia", "alo", "dona"}
	ordem_alfabetica, err := ordem_alfabetica(palavras)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(ordem_alfabetica)
	}
}
func ordem_alfabetica(palavras []string) ([]string, error) {
	if len(palavras) == 0 {
		return nil, fmt.Errorf("Sua lista está vazia")
	}
	sort.Strings(palavras)

	return palavras, nil
}

package main

import "fmt"

func main() {
	lista := []string{"sabemos", "como", "chegar", "ao", "nosso", "destino"}
	validas, err := maior_5(lista)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(validas)
	}

}

func maior_5(palavras []string) ([]string, error) {
	if len(palavras) == 0 {
		return nil, fmt.Errorf("Lista vazia!")
	}
	palvras_validas := []string{}

	for _, palavra := range palavras {
		if len(palavra) > 5 {
			palvras_validas = append(palvras_validas, palavra)
		}
	}

	return palvras_validas, nil
}

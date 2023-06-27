package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(vog_check("comovai"))
}

func vog_check(palavra string) int {
	vogais := "AEIOUaeiou"
	contador_vogais := 0

	for i := 0; i < len(palavra); i++ {
		if strings.Contains(vogais, string(palavra[i])) {
			contador_vogais++
		}
	}

	return contador_vogais
}

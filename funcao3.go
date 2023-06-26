package main

import "fmt"

func main() {
	fmt.Print(concatenação([]string{"ola", "como", "vai"}))
}

func concatenação(palavras []string) string {
	var palavra_concatenada string

	for i := 0; i < len(palavras); i++ {
		palavra_concatenada += palavras[i]
	}
	return palavra_concatenada
}

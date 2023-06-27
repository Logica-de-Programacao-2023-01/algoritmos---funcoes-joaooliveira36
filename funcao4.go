package main

import "fmt"

func main() {
	fmt.Print(segundo_maior([]int{1, 6, 2, 3, 7, 3, 8}))
}

func segundo_maior(numeros []int) int {
	maior_num := numeros[0]
	segundo_maior_num := numeros[1]

	if segundo_maior_num > maior_num {
		maior_num = segundo_maior_num
		segundo_maior_num = maior_num
	}

	for _, numero := range numeros[2:] {
		if numero > maior_num {
			segundo_maior_num = maior_num
			maior_num = numero
		} else if numero > segundo_maior_num {
			segundo_maior_num = numero
		}

	}

	return segundo_maior_num
}

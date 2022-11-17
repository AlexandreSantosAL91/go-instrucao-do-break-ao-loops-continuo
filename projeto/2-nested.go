package main

import "fmt"

func main() {
	for outer := 0; outer < 5; outer++ {
		if outer == 3 {
			fmt.Println("Saindo do loop externo")
			break // break here
		}
		fmt.Println("O valor de exterior é", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 2 {
				fmt.Println("Saindo do loop interno")
				break // break here
			}
			fmt.Println("O valor de interno é", inner)
		}
	}
	fmt.Println("Tarefa concluida")
}

/*
Resultado
O valor de exterior é 0
O valor de interno é 0
O valor de interno é 1
Saindo do loop interno

O valor de exterior é 1
O valor de interno é 0
O valor de interno é 1
Saindo do loop interno

O valor de exterior é 2
O valor de interno é 0
O valor de interno é 1
Saindo do loop interno
Saindo do loop externo

Saindo do programa
*/

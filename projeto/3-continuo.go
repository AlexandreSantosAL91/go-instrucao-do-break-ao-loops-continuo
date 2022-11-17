package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Loop contínuo")
			continue // break here
		}
		fmt.Println("O valor de i é", i)
	}
	fmt.Println("Tarefa concluida")
}

/*
Resultado:
O valor de i é 0
O valor de i é 1
O valor de i é 2
O valor de i é 3
O valor de i é 4
Loop contínuo
O valor de i é 6
O valor de i é 7
O valor de i é 8
O valor de i é 9

Saindo do programa
*/

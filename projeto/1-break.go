package main

import "fmt"

/*
Utilizando uma instrução break em um for loop:
*/

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			fmt.Println("Saindo do loop")
			break // break here
		}
		fmt.Println("O valor de i é: ", i)
	}
	fmt.Println("Tarefa concluida")
}

/*
Resultado:
O valor de i é: 0
O valor de i é: 1
O valor de i é: 2
O valor de i é: 3
O valor de i é: 4
O valor de i é: 5
Saindo do Loop

Tarefa concluida
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num1, num2 float64
	var operation string

	// Solicita ao usuário para inserir o primeiro número
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	// Solicita ao usuário para inserir o segundo número
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	// Solicita ao usuário para inserir a operação desejada
	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&operation)

	var result float64
	var err error

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Erro: divisão por zero não é permitida.")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Erro: operação inválida.")
		os.Exit(1)
	}

	// Exibe o resultado na tela
	fmt.Printf("Resultado: %.2f\n", result)
}

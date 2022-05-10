package main

import "fmt" // É necessário importar o módulo de função para utilizar funções

func main(){

	// Declaração das variáveis e seus tipos
	var X, Y float64
	var op string

	// Recebimento de valores do teclado
	fmt.Print("Primeiro número: ")
	fmt.Scanln(&X)
	fmt.Print("Segundo número: ")
	fmt.Scanln(&Y)

	// Menu de operações
	fmt.Println("Escolha a operação: ")
	fmt.Println("+")
	fmt.Println("-")
	fmt.Println("/")
	fmt.Println("*")

	// Leitura do valor digitado
	fmt.Scanln(&op)

	// Switch que seleciona a operação e imprime o resultado
	switch op {

		case "+":
			fmt.Printf("%f %s %f = %f", X, op, Y, X+Y)
		
		case "-":
			fmt.Printf("%f %s %f = %f", X, op, Y, X-Y)

		case "*":
			fmt.Printf("%f %s %f = %f", X, op, Y, X*Y)

		case "/":
			if Y == 0.0{
				fmt.Println("Não é possível fazer uma divisão por 0!")
			} else {
				fmt.Printf("%f %s %f = %f", X, op, Y, X/Y)
			}	
		
		default:
			fmt.Println("Opção Inválida!")
	}
}
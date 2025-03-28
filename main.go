package main

import "fmt"

func main(){

	var num1 int
    var num2 int

	fmt.Println("Digite um número: ")
	fmt.Scan(&num1)
	fmt.Println("Digite outro número: ")
	fmt.Scan(&num2)
	fmt.Println("A soma é ", num1 + num2)
	fmt.Println("A subtração é ", num1 - num2)
	fmt.Println("A multiplicação é ", num1 * num2)
	fmt.Println("O resto da divisão é ", num1 % num2)
}
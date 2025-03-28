package main

import "fmt"

func main(){

	a, b := 10, 3


	a += 1
	fmt.Println("Incrementar a", a)

	if a > 0 && b > 0 {
		fmt.Println("Positivos")
	}
}
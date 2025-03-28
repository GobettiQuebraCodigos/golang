package main

import "fmt"

func main(){

	var user string
	var senha string

    fmt.Println("Digite o usuário:")
	fmt.Scan(&user)
	fmt.Println("Digite a senha: ")
	fmt.Scan(&senha)

	if user == "admin" && senha == "1234" {
		fmt.Println("Acesso liberado")
	
 
	} else {
		fmt.Println("Acesso negado")
	}
}
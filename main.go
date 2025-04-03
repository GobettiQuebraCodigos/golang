package main

import "fmt"

func main(){

	var num = []int{10, 20, 30, 40, 50}
    fmt.Println(num)
	num = append(num, 60, 70, 80)
	fmt.Println(num)
	fmt.Println(len(num), cap(num))
	

}
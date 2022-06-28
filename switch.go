package main

import "fmt"

func main(){
	switch modulo := 4 % 2; modulo{
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
}
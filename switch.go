package main

import "fmt"

func main(){
	modulo := 4 % 2
	
	switch modulo = 5 % 2; modulo{
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//sin condición
	value := 200
	switch{
	case value > 100:
		fmt.Println("Es mayor a 0")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No aplica a ninguno")
	}
}
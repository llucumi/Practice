package main

import "fmt"

func main(){
	sum :=0 
	//Standar for
	for i := 0; i < 100; i++ {
		if i == 2{
			fmt.Println("continue 2")
			continue
		}
	}
	
	//For range
	frutas := []string{"manzana","banana","pera"}
	for i, fruta := range frutas {
		fmt.Println(i,fruta)
	}

	//Bucle infinito
	for{
		sum++
	}

	//Bucle "while"
	suma := 1

	for suma < 10 {
		suma += suma
	}

	fmt.Println(suma)
}
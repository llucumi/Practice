package main

import "fmt"

func main(){
	numberOdd(2)
	fmt.Println(userPassword("luz","luz123"))
}

func numberOdd(a int){
	if a % 2 == 0{
		fmt.Println("Es par")
	}else{
		fmt.Println("Es Impar")
	}
}

func userPassword(name, pass string) bool{
	if name == "luz" && pass == "luz123"{
		return true
	}else{
		return false
	}
}
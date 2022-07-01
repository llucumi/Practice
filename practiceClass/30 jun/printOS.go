package main

import(
	"fmt"
	"os"
)

func main(){
	//Setenv
	//La variable NAME tiene el valor gopher
	//key, value
	err := os.Setenv("NAME", "gopher")
	fmt.Printf("Error: %v\n",err)
	//Getenv
	//Si name no está nos devuelve un error
	value := os.Getenv("NAME")
	fmt.Println(value)
	//LookupEnv
	//Nos dice si existe, y si existe nos retorna
	//el valor
	value, ok := os.LookupEnv("NAME")
	fmt.Println(ok)
	//Exit
	//
	os.Exit(1)
}
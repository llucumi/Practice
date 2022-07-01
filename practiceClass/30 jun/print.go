package main

import(
	"fmt"
	"os"
	//"io"
)

func main(){
	nombre, edad := "Kim", 22
	//fmt

	//Print
	fmt.Print(nombre, " tiene ", edad, " años de edad.\n")
	fmt.Println(nombre, " tiene ", edad, " años de edad.")
	//Printf
	fmt.Printf("%10.2f \n", 12222.123123)
	fmt.Printf("%.2f \n", 12222.123123)
	//Print
	fmt.Print(nombre, " tiene ", edad, " años de edad.\n")
	fmt.Println(nombre, " tiene ", edad, " años de edad.")
	//Printf
	fmt.Printf("%s tiene %d años de edad.\n", nombre, edad)
	fmt.Printf("%10d", 12222)
	fmt.Printf("%10s", "aa\n")
	//Sprint
	//Concatena la info de res (retorna string)
	res := fmt.Sprint(nombre, " tiene ", edad, " años de edad.\n")
	fmt.Print(res)
	//Sprintf
	//Concatena la info de res (retorna string)
	res = fmt.Sprintf("%s tiene %d años de edad.\n", nombre, edad)
	fmt.Print(res)
}
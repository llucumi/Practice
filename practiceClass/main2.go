package main

import "fmt"

func main(){
	var i,j int= 10, 20
	j += i
	fmt.Printf("resultado=%d \n", j)
	fmt.Printf("resultado=%v \n", i != j)
	fmt.Println(i == j)

	text := "Hola Mundo"
	var pText *string

	pText = &text

	fmt.Println(pText) /*0xc0000a0a0*/
	fmt.Println(*pText) /* Hola Mundo */

	//
	/*var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Prinln(a)
	fmt.Printf(a[0],a[1])*/

	/*var s = []bool{true, false}
	fmt.Println(s[0])*/

	//Array a[0],a[1]
	a := make([]int, 5) //len(a) = 5
	fmt.Println(a)

	//Slice
	var slice = []int{1,2,3,4,5,6,7,8}
	fmt.Printf("resultado = %v",slice[1:4])
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	fmt.Println(append(slice,2,3,4))

	//Map
	//myMap := map[string]int{}
	myMap := make(map[string]string)
	fmt.Println(len(myMap))

	var students = map[string]int{"Benjamin": 20, "Nahuel": 26}
	fmt.Println(students)
	students["Brenda"] = 19
	students["Marcos"] = 22
	fmt.Println(students)

	delete(students,"Benjamin")	
	fmt.Println(students)

	//recorrer elementos
	//range recorre el mapa valor por valor
	for key, element := range students {
		fmt.Println("Key:", key,"=>","Element:",element)
	}

}
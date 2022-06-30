package main

import "fmt"

func main() {
	//tipeAsertion es como abrir una caja
	//"ahh, es un string"
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //panic: errores que
	//ocurren en tiempo de ejecucion
	fmt.Println(f, ok)
}

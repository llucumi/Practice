package main

import "fmt"

func main() {
	//Array
	//Inmutable
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, "elementos:", len(array), "Capacidad", cap(array))

	//Slice
	//Mutable
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, "elementos:", len(slice), cap(slice))

	//Slicing
	//Para interactuar entre elementos
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Adicionar elementos al array (append)
	slice = append(slice, 7)

	//Agregar una lista
	newSlice := []int{8, 9, 10}
	//Los tres puntos porque cada elemento se agrega
	//de manera independiente
	slice = append(slice, newSlice...)
}

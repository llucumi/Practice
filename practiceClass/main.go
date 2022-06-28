package main

import "fmt"

func main(){
	fmt.Println("Practice at class! :3")

	//VARIABLES
	var nombre string
	var horas int

	horas = 20
	horasSinDeclaracion := 20

	fmt.Println(horas,horasSinDeclaracion)

	//Dos variables
	product, price := "Jean", 10.5

	fmt.Println(product,price)

	//agrupación de variables
	var(
		product = "Course"
		quantity = 25
		price = 40.50
		inStock = true
	)

	fmt.Println(product, quantity, price, inStock)

	//Encasillamiento o casting
	var i int = 123
	var f float64 = float64(i)

	fmt.Println(i, f)

	//CONSTANTES
	const Nombre = "valor"

	const(
		Product = "Course"
		Quantity = 25
		Price = 40.50
	)

	fmt.Println(Product, Quantity, Price)
}
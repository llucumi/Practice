package main

import "fmt"

//Definiendo la struct
type Car struct{
	brand string
	modelo int
}

func main(){
	myCar := Car{brand:"Ford",modelo:2020}
	fmt.Println(myCar)

	var otherCar Car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	myThirdCar := Car{"Mazda",2021}
	fmt.Println(myThirdCar)
}
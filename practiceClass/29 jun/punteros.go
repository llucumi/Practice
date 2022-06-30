package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circulo{radio: 5}
	fmt.Println(c.area())
	fmt.Println(c.perim())
	c.setRadio(10)
	//Cambia porque el radio cambió
	fmt.Println(c.area())
	fmt.Println(c.perim())
}

type Circulo struct {
	radio float64
}

//No usa puntero porque no necesita cambiar el
//valor de la variable
func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

//Si va a modificar el radio
func (c *Circulo) setRadio(r float64) {
	c.radio = r
}

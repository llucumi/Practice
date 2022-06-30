package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

//Recibe una interfaz geometry
//Voy a buscar una estructura que esté
//Asociada a geometry
func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

//estructura circulo con sus metodos area
// y perimetro
type circle struct {
	radio float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radio
}

type rectangle struct {
	width, height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

func (r *rectangle) perim() float64 {
	return 2 * r.width * r.height
}

func newCircle(values float64) geometry {
	//Intanscia un circulo y devuelve el puntero a este circulo
	//que el circulo de geometria (area y perim)
	return &circle{radio: values}
}

func newRectangle(h, w float64) geometry {
	//Intanscia un rectangulo y devuelve el puntero a este circulo
	//que el circulo de geometria (area y perim)
	return &rectangle{height: h, width: w}
}

func main() {
	c := newCircle(3)
	fmt.Println("Circle")
	fmt.Printf("area: %.3f\n", c.area())
	fmt.Printf("perimetro: %.3f\n", c.perim())
	//Nos obliga a crear las funciones de la
	//interfaz
	r := newRectangle(2, 5)
	fmt.Println("Rectangle")
	fmt.Printf("area: %.3f\n", r.area())
	fmt.Printf("perimetro: %.3f\n", r.perim())
}

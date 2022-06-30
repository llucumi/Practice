package main

import "fmt"

func main() {
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()
}

type Vehiculo struct {
	km     float64
	tiempo float64
}

type Auto struct {
	v Vehiculo
}

func (a *Auto) Correr(minutos int) {
	//auto tiene vehiculo y vehiculo tiene tiempo
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	//auto tiene vehiculo y vehiculo tiene detalle
	a.v.detalle()
}

func (v Vehiculo) detalle() {
	fmt.Printf("km:\t%f\ntiempo:\t%f\n", v.km, v.tiempo)
}

type Moto struct {
	v Vehiculo
}

//Correr y Detalle ahora para moto
func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 80
}

func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

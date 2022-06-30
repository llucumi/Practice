package main

import "fmt"

func main() {
	type Preferencias struct {
		Comidas, Pelicula string
	}

	type Persona struct {
		Nombre    string
		Edad      int
		Profesion string
		Peso      float64
		Gustos    Preferencias
	}

	p1 := Persona{"Luz", 26, "Ingeniera", 55.0, Preferencias{"Titanic", "Asado"}}

	fmt.Println(p1)

	p2 := Persona{
		Nombre:    "Maria",
		Edad:      24,
		Profesion: "Maestra",
		Peso:      30,
		Gustos: Preferencias{
			Pelicula: "Terminator",
			Comidas:  "Lasagna",
		},
	}

	fmt.Printf("p2: %v", p2)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Persona struct {
		Nombre   string `bd:"nombre"`
		Apellido string `bd:"apellido"`
		Edad     int
	}

	persona := Persona{"Luz", "Lucumi Hernandez", 26}
	personaReflect := reflect.TypeOf(persona)

	//Se usa cuando se necesita saber el tipo de dato en tiempo
	//de ejecución
	fmt.Printf("Type: %v\n", personaReflect.Name()) //Persona
	fmt.Printf("Kind: %v\n", personaReflect.Kind()) //struct

	for i := 0; i < personaReflect.NumField(); i++ {
		field := personaReflect.Field(i)
		//trae todo lo que traiga bd
		fmt.Printf("Tag %v", field.Tag.Get("bd"))
	}
}

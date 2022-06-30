package main

import "fmt"

//En el campo Data voy a tener un array de algun tipo
//de dato que quiera, y como los tipos de datos
//pueden generar uno o n datos
type ListaHeterogenea struct {
	Data []interface{}
}

//NO ES UNA BUENA PRÁCTICA USARLO
func main() {
	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "hola")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)
}

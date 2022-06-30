package main

import "fmt"

func main(){
	s := suma(4,5)
	fmt.Println(s)

	var(
		Suma = "+"
		Resta = "-"
		Multiplicacion = "*"
		Division = "/"
	)

	a, b := 4, 2
	fmt.Println("Suma %.2f",a,b,"Suma")
}

func operacion(valor1, valor2 float64, operador string) float64{
	switch operador {
	case Suma:
		return valor1 + valor2
	case Resta:
		return valor1 - valor2
	case Multiplicacion:
		return valor1 * valor2
	case division:
		if valor2 != 0{
			return valor1/valor2
		}
	}
	return 0
}
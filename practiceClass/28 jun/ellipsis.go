package main

func main(){
	const(
		Suma = "+"
		Resta = "-"
		Multip = "*"
		Divis = "/"
	)

	fmt.Println(operacionAritmetica(Suma, 2, 3, 4, 5, 6, 1, 2))
}

func opSuma(valor1, valor2 float64) float64{
	return valor1 + valor2
}

func opResta(valor1, valor2 float64) float64{
	return valor1 - valor2
}

func opMultip(valor1, valor2 float64) float64{
	return valor1 * valor2
}

func opDivis(valor1, valor2 float64) float64{
	if valor2 == 0{
		return 0
	}
	return valor1 / valor2
}

func orquestadorOperaciones(valores []float64, operacion func(value1, value2 float64) float64) float64{
	var resultado float64
	for i, valor := range valores{
		if i == 0{
			resultado = valor
		}else{
			resultado = operacion(resultado,valor)
		}
	}
	return resultado
}

func operacionAritmetica(operador string, valores ...float64) float64{
	length = len(valores)
	fmt.Printf("len: %d\n", length)

	switch operador {
		case Suma:
			return orquestadorOperaciones(valores,opSuma)
		case Resta:
			return orquestadorOperaciones(valores,opResta)
		case Multiplicacion:
			return orquestadorOperaciones(valores,opMultip)
		case division:
			if valor2 != 0{
				return orquestadorOperaciones(valores,opDivis)
			}
		}
		return 0
	}
}
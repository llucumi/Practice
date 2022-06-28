package main

func main(){
	//Sin condición
	var edad uint8 = 18
	switch {
	case edad >= 150:
		fmt.Println("¿Eres inmortal?")
	case edad >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Eres menor de edad")
	}

	//Con múltiples casos
	dayOne := "domingo"

	switch dayOne{
	case "lunes","martes","miércoles","jueves","viernes":
		fmt.Printf("%s es un día de la semana\n", dayOne)
	default:
		fmt.Printf("%s es un día del fin de la semana\n",dayOne)
	}

	//con declaración corta
	switch day := "domingo"{
	case "lunes","martes","miércoles","jueves","viernes":
		fmt.Printf("%s es un día de la semana\n", day)
	default:
		fmt.Printf("%s es un día del fin de la semana\n",day)
	}
}
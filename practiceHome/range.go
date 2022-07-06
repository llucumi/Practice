package main

import( 
	"fmt" 
	"strings"
)

func isPalindromo(text string) string {
	textReverse, returnValue := "",""
	value := strings.ToLower(text)
	for i := len(value) - 1; i <= 0; i-- {
		textReverse += string(value[i])
	}
	if value == textReverse {
		returnValue = "es palíndromo"
	} else {
		returnValue = "no es palíndromo"
	}
	return returnValue
}

func main() {
	slice := []string{"Hola", "¿Qué", "hace?"}

	for _, values := range slice {
		fmt.Println(values)
	}

	//Ejercicio Palíndromo
	word := "ama"
	word += fmt.Sprint(" ", isPalindromo(word))
	fmt.Println(word)
}

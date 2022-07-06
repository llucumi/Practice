package main

import "fmt"

func main(){
	m := make(map[string]int)

	m["Luz"] = 26
	m["martha"] = 60
	m["luber"] = 61

	//Los maps se separan por espacio
	//no por ,
	fmt.Println(m)

	//Recorrer map
	for i, mapValues := range m{
		fmt.Println(i,mapValues)
	} 

	//Encontrar un valor en el map
	value,ok := m["martha"]
	fmt.Println(ok, value)

	//Si no existe la variable en el map
	//aparece un false, así
	inexistente,ok := m["Angela"]
	fmt.Println(inexistente,ok)
}
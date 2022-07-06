package main

import (
	"fmt"
	"net/http"
)

//PAQUETE SIMPLE PARA GENERAR SERVIDOR WEB
//MANEJADOR llamado handler
func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola Luz\n")
}

func main() {
	//Iniciamos, patern y manejador
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}

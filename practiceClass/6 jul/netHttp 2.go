package main

import (
	"fmt"
	"net/http"
)

//PAQUETE SIMPLE PARA GENERAR SERVIDOR WEB
//MANEJADOR llamado handler
func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

//localhost:8080/hola
//ctr+C para salir de ejecución
func main() {
	//Iniciamos, patern y manejador
	http.HandleFunc("/ping", holaHandler)
	http.HandleFunc("/pang", holaHandler)
	http.ListenAndServe(":8080", nil)
}

package main

import( 
	"fmt"
	"time"
)

func Incrementar(value *int) {
	// Desreferenciamos la variable v para obtener
	// su valor e incrementarlo en 1
	*value++
 }

func procesar(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
} 

//Recibir
func procesarRecibirChan(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	//fmt.Println(i, "-Termina")
	c <- i
	//c<-i
 }
 
//Enviar valor a canal
func procesarEnviarChan(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	//fmt.Println(i, "-Termina")
	c <- i
 }
 

func main(){
	//Desrreferenciación
	var v int = 19
	var p *int
	//Hacemos que el puntero p, referencie la dirección de memoria de la 
	//variable v.

	//puntero es una variable y es un dato que tiene la dirección
	//*traer la información de ese valor
	//p apunta a a dirección de la variable v pero no se puede 
	//operar con el puntero, sólo se puede asignar cosas de go
	p = &v
	fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n",p)
	fmt.Printf("Al desreferenciar el puntero p obtengo el valor: %d \n",*p)

	var value int = 19
	// La función Increase recibe un puntero
	// utilizamos el operador de dirección &
	// para pasar la dirección de memoria de v

	//Para cambiar la variable e incrementar una direccion
	//de memoria
	Incrementar(&value)
	fmt.Println("El valor de v ahora vale:", value)

	// fmt.Println("SIN GO ROUTINES")
	// for i := 0; i < 10; i++{
	// 	procesar(i)
	// }
	
	// time.Sleep(5000 * time.Millisecond)
    // fmt.Println("Termino el programa")

	//LAS GOROUTINES SE EJECUTAN DE MANERA CONCURRENTE

	//La función main es una goroutine
	//Imprime TERMINO EL PROGRAMA
	//murió la goroutine de main y chao a
	//procesar
	// fmt.Println("CON GO ROUTINES")
	// for j := 0; j < 10; j++{
	// 	go procesar(j)
	// }
	// time.Sleep(5000 * time.Millisecond)
	// fmt.Println("Termino el programa")
	 
	//CANALES
	//Una manera de enviar valores entre
	//goroutines, hace las veces de tubo de
	//comunicación

	//apunta hacia canal escribiendo
	//apunta hacia afuera imprimiendo
	c := make(chan int)
	for l := 0; l < 5; l++{
    	go procesarRecibirChan(l, c)
	}

	// c = make(chan int)
    // go procesarEnviarChan(1, c)
    // fmt.Println("Termino el programa")

	// //<-c // recibimos el valor del canal
	// variable := <-c // recibimos y lo asignamos a una variable
	// fmt.Println("Termino el programa en ", variable) // recibimos y lo imprimimos

	for k := 0; k < 5; k++{
		fmt.Println(<-c,"Termina")
	}
    fmt.Println("Termino el programa")
 
}
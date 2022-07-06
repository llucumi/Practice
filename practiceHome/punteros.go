package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

//agregando funciones a los struct
func (p pc) ping() {
	fmt.Println(p.brand, " Pong")
}

func (p *pc) duplicationRAM() {
	p.ram = p.ram * 2
}

func main() {
	a := 50
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{16, 200, "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicationRAM()

	fmt.Println(myPc)
	myPc.duplicationRAM()

	fmt.Println(myPc)
}

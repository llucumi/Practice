package main

import(
	"io"
	"os"
	"fmt"
	"strings"
)

func main(){
	//IO
	//Copy
	//Genera una estructura de lectura
	r := strings.NewReader("some io.Reader stream to be read\n")
	//primero destino, luego origen
	//ignoramos cantidad de bytes copiados
	//pero quedamos con el error por si lo necesitamos
	bytes,err := io.Copy(os.Stdout, r);
	fmt.Printf("error %v\n",err)
	fmt.Printf("Bytes escritos %v\n",bytes)
	//ReadAll
	//Genera una estructura de lectura y la lee
	//LEE TODOS HASTA ENDOFF
	r = strings.NewReader("some io.Reader stream to be read")
	b,err := io.ReadAll(r); //b,err
	fmt.Printf("bytes: %v\n",string(b))
	fmt.Printf("error %v\n",err)
	//WriteString
	//toma la cadena y la escribe en el output de pantalla
	io.WriteString(os.Stdout, "Hello world!\n")
}
package main

import(
	"fmt"
	"os"
)

func main(){
	//ReadDir
	//files, _ := os.ReadDir(".")
	//fmt.Printf("files:%v",string(files))

	//ReadFile
	data, err := os.ReadFile("./myFile.txt")
	fmt.Println(err)
	fmt.Printf("file %v\n",string(data))
	//WriteFile
	d1 := []byte("hello, gophers!")
	err = os.WriteFile("./myFile.txt", d1, 0644)

	//Unsetenv
	err = os.Unsetenv("MI_VARIABLE")
	fmt.Println(err)

	
}
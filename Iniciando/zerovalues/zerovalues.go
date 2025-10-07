package main

import "fmt"

func main() { //quando nao atribuo um valor, ele da um valor default
	var i int     //valor inicial 0
	var f float64 //valor inicial 0
	var b bool    //valor inicial false
	var s string  //valor inicial ""

	fmt.Printf("inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %q\n", s)
}

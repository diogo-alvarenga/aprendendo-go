package main

import "fmt"

//bool = true/false
//string = sequencia de bytes - diferente de java ***
//int inteiros
//float (float64/float32) decimal

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, true) //%T = tipo, %v = valor
	fmt.Printf("Type: %T - Value: %v\n", "steph", "steph")
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
}

///*** Mterial de estudo
//s := "olá" // O 'á' usa mais de um byte em UTF-8
// 1. Contagem de Bytes (usando a função padrão len())
//fmt.Println(len(s))
// Output: 4
// (os bytes são: 'o' (1 byte) + 'l' (1 byte) + 'á' (2 bytes) = 4 bytes)

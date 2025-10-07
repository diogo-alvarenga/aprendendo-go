package main

import "fmt"

func main() {
	//Variaveis
	//declarando string
	var nome string
	nome = "bento"
	fmt.Println(nome)

	nome = "steph"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "Testando" //ele pega o tipo automaticamente
	fmt.Println(a)

	var b, c int = 2, 3
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	num := 2
	fmt.Println(num)

	//Constantes
	const idadeBento = 4
	fmt.Println(idadeBento)
}

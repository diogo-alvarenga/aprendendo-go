package main

//LISTAS
//Arrays e Slices: homogeneos
//todos os elementos tem o MESMO tipo
//[1,2,3,4,5,6] - int
//["arthur","vanesssa","caio"] - []string

//Array

//Tamanho fixo, de zero ou mais elementos do mesmo tipo
//Acessamos os valores com indice: a[0], a[1]...
//Funçao embutida len() retorna o tamanho do array
//Por conta do tamanho fixo, só se usa em casos especificos

//Slice

//Igual o array, mas sem tamanho fixo
//Acessamos os valores com indice a[0], a[1]...
//Funçao embutida len() retorna o tamanho do slice
//Funcao append() adiciona itens

import "fmt"

func main() {
	//Array -tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	//Outra forma de criar o array
	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0])
	fmt.Println(numPrimos[1])
	fmt.Println(numPrimos[0:3]) //tudo antes da posição 3
	fmt.Println(numPrimos[2:])  //tudo a partir da posição 2

	//Slices
	//Eu preciso dar um tamanho previo para o slice antes de
	//adicionar coisas dentro. Se eu comentar a linha 47
	//e descomentar a 46, irá gerar erro
	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])

	fmt.Println(slice[2])
	slice[2] = "Olá"
	fmt.Println(slice[2])

	fmt.Println(len(slice))
	fmt.Println(slice)

	//Criando um slice ja com valores
	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)
	//adicionando mais conteudo a esse slice
	numPares = append(numPares, 14)
	fmt.Println(numPares)
}

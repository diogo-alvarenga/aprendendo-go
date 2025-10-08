package main

import "fmt"

func main() {
	var nome string

	fmt.Println("Digite seu nome: ")

	fmt.Scanln(&nome)
	fmt.Println("Ola, ", nome)

	soma := soma(42, 13)
	fmt.Println("A soma dos valores é: ", soma)

	sub := subtracao(10, 5)
	fmt.Println("A subtração resulta em: ", sub)

	div := divisao(10, 3)
	fmt.Println("O resultado da divisão é: ", div)

	mult := multiplicacao(2, 10)
	fmt.Println("O resultado da multiplicação é: ", mult)

	fmt.Println("Tchau, ", printaNome(nome))
}

/*
	FUNÇÃO QUE COMEÇA COM LETRA MINUSCULA
	são funções PRIVADAS
	ela só pode ser usada nesse pacote
*/
/*
	FUNÇÃO QUE COMEÇA COM LETRA MAIUSCULA
	são funções PUBLICAS
	ela só pode ser usada fora do pacote

*/

func soma(x int, y int) int {
	resultado := x + y
	return resultado
}

func subtracao(x int, y int) int {
	resultado := x - y
	return resultado
}

func divisao(x int, y int) int {
	return x / y
}

func multiplicacao(x int, y int) int {
	return x * y
}

func printaNome(nome string) string {
	return nome
}

func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

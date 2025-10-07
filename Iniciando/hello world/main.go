package main

//import "fmt"

import (
	"fmt"
	steph "strings" //estou dando o nome para o pacote de steph, assim preciso chama-lo dessa forma
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Print(steph.Split("pasta/pasta2/pasta3/arquivo.go", "/")) //o metodo corta fora o caractere escolhido(segunda string)
}

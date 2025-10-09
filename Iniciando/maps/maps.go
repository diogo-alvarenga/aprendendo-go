package main

import "fmt"

func main() {

	//Maps: Heterogeneos
	//pode misturar tipos
	//estrutura chave - valor
	//[key] - value
	//chave tem um tipo, e o valor pode ter outro
	//map[k]v -> k = chave, v = valor

	idade := map[string]int{} //string é a chave, int o valor
	idade["Ana"] = 21
	idade["Ronaldo"] = 4
	fmt.Println(idade)
	fmt.Println(idade["Ana"])
	fmt.Println(idade["Ronaldo"])

	anoNasc := map[string]int{
		"Ana":     2004,
		"Ronaldo": 2021,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Ana"])
	fmt.Println(anoNasc["Ronaldo"])
	anoNasc["Rogerio"] = 2025 //adicionando novos elementos
	fmt.Println(anoNasc)

}

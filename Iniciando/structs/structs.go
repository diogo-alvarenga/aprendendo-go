package main

import "fmt"

//Structs
//Forma de criar sau propria estrutura de dados
//Personalizar de acordo com a sua necessidade
//Podemos usar varios tipos diferentes

//type <nome da estrutura> struct { <campos> }
type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {

	fmt.Println(Pessoa{"Marcos", 21})
	fmt.Println(Pessoa{Nome: "Roger", Idade: 4})
	fmt.Println(Pessoa{Nome: "Daniel"}) //como nao preenchi idade, ele automaticamente vai ficar 0

	//Pessoa é um tipo, entao o tipo da variavel é Pessoa
	p1 := Pessoa{Nome: "Bob", Idade: 2}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 3
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 2}

	//slice de Pessoa
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	//Map de pessoas
	alunos := map[string][]Pessoa{}
	alunos["Programação"] = pessoas
	fmt.Println(alunos)

	var outrosAlunos = map[string][]Pessoa{
		"Programação": {{Nome: "Arthur", Idade: 20}, {Nome: "Bento", Idade: 19}},
		"Engenharia":  {{Nome: "Vinicius", Idade: 21}, {Nome: "Joao", Idade: 22}},
	}

	fmt.Println(outrosAlunos)

	//struct profissao herdando campos da struct pessoa
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)

}

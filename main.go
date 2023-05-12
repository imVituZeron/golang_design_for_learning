package main

import (
	"fmt"
)

type Pessoa struct {
	Id        int
	nome      string
	sobrenome string
	profissao string
}

var (
	name     string
	lastname string
	work     string
)

func main() {
	i := 0
	list := []Pessoa{}

	for {
		fmt.Println("Digite seu nome: ")
		fmt.Scan(&name)
		fmt.Println("Digite seu sobrenome: ")
		fmt.Scan(&lastname)
		fmt.Println("Digite seu profissao: ")
		fmt.Scan(&work)

		people := Pessoa{i, name, lastname, work}

		list = append(list, people)

		fmt.Println("Deseja Continuar? [y/n]")
		var code string
		fmt.Scan(&code)
		if code == "y" {
			i++
		} else if code == "n" {
			fmt.Println("Finalizando!")
			break
		}
	}

	fmt.Println(list)
}

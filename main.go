package main

import (
	"fmt"
	P "project/packages/models"
	create "project/packages/utils"
)

var (
	name     string
	lastname string
	work     string
)

func main() {
	i := 0
	listPeople := []P.Pessoa{}

	for {
		fmt.Println("Digite seu nome: ")
		fmt.Scan(&name)
		fmt.Println("Digite seu sobrenome: ")
		fmt.Scan(&lastname)
		fmt.Println("Digite seu profissao: ")
		fmt.Scan(&work)

		people := P.Pessoa{i, name, lastname, work} // create instance of Pessoa
		listPeople = append(listPeople, people)     // add instance in the Slice

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

	exit := create.Createfile(listPeople)
	fmt.Println(exit)
}

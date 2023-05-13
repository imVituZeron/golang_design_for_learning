package main

import (
	"fmt"
	P "project/packages/models"
	utils "project/packages/utils"
)

var (
	name, lastname, work, code string
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
		fmt.Scan(&code)
		if code == "y" {
			i++
		} else if code == "n" {
			break
		}
	}

	exit := utils.Createfile(listPeople)
	fmt.Println(exit)
}

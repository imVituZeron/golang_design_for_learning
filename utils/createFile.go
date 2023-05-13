package utils

import (
	"os"
	P "project/packages/models"
)

// Created function for create file with name of person
func Createfile(value []P.Pessoa) string {
	for _, v := range value {
		file, _ := os.OpenFile("files/"+v.Nome+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		file.WriteString("Name: " + v.Nome)
		file.WriteString("\nLastname: " + v.Sobrenome)
		file.WriteString("\nProfession: " + v.Profissao)
		file.Close()
	}

	return "Finish!"
}

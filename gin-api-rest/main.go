package main

import (
	"gin-api-rest/database"
	"gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Gui Lima", CPF: "00000000000", RG: "4700000000"},
	// 	{Nome: "Ana", CPF: "11111111111", RG: "4800000000"},
	// }
	routes.HandleRequest()
}

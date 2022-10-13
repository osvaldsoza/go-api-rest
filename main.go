package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{1, "Nome 1", "História 1"},
		{2, "Nome 2", "História 2"},
	}
	database.OpenConnection()
	fmt.Println("Inciando o servidor rest com go")
	routes.HandleRequest()
}

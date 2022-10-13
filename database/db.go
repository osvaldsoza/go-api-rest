package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func OpenConnection() {
	stringConexao := "host=localhost user=postgres password=postgres dbname=go_rest port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panicln("Erro ao conectar com o banco")
	}
}

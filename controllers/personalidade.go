package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api-rest/database"
	"go-api-rest/models"
	"net/http"
)

func TodasPersonalidades(resp http.ResponseWriter, req *http.Request) {
	var personalidade []models.Personalidade
	database.DB.Find(&personalidade)
	json.NewEncoder(resp).Encode(personalidade)
}

func RetornaUmaPersonalidade(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(resp).Encode(personalidade)
}

func CriaNovaPersonalidade(resp http.ResponseWriter, req *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(req.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(resp).Encode(personalidade)
}

func DeletaUmaPersonalidade(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(resp).Encode(personalidade)
}

func EditaUmaPersonalidade(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Find(&personalidade, id)
	json.NewDecoder(req.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(resp).Encode(personalidade)
}

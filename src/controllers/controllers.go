package controllers

import (
	"encoding/json"
	"net/http"

	"go/src/models"
	"go/src/services"
)

func GetAmount(w http.ResponseWriter, r *http.Request) {

	requestBody := models.GetAmountBody{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Error":       "Error with request body",
			"status code": 300,
		})
		return
	}

	user, err := services.GetUserById(requestBody.Cpf)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Error":       "Error in getting user, try with another CPF",
			"status code": 300,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Name":        user.Name,
		"Amount":      user.Amount,
		"status code": 200,
	})
}

func TransferMoney(w http.ResponseWriter, r *http.Request) {

	requestBody := models.TransferBody{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Error":       "Error with request body",
			"status code": 400,
		})
		return
	}
	messageReturn, err := services.Debit(requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Error":       messageReturn,
			"status code": 400,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Sucess":      messageReturn,
		"status code": 200,
	})

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	requestBody := models.User{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Error":       "Error with request body",
			"status code": 400,
		})
		return
	}
	if services.RegisterUser(requestBody) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Sucess":      "User created successfully",
			"Name":        requestBody.Name,
			"CPF":         requestBody.Cpf,
			"status code": 200,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Error":       "Error in connection database",
		"status code": 400,
	})
}

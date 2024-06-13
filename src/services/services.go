package services

import (
	"context"
	"encoding/json"

	"go/conection"
	"go/models"
)

func GetUserById(id string) (models.User, error) {
	user := models.User{}
	client := conection.ConectionDbUser()
	val, err := client.Get(context.Background(), id).Result()

	if err != nil {
		return user, err
	}
	err = json.Unmarshal([]byte(val), &user)

	if err != nil {
		return user, err
	}
	return user, nil

}

func RegisterUser(user models.User) bool {

	jsonObject, err := json.Marshal(user)
	if err != nil {
		return false
	}

	dbUser := conection.ConectionDbUser()
	err = dbUser.Set(context.Background(), user.Cpf, jsonObject, 0).Err()

	if err != nil {
		return false
	}

	return true

}

func Debit(requestBody models.TransferBody) (string, error) {

	userPayer, err := GetUserById(requestBody.Payer)
	if err != nil {
		return "with connection", err
	}

	userPayee, err := GetUserById(requestBody.Payee)

	if err != nil {
		return "with connection", err
	}

	if userPayer.Amount >= requestBody.Value && !userPayer.Shopper {

		userPayer.Amount = userPayer.Amount - requestBody.Value
		userPayee.Amount = userPayee.Amount + requestBody.Value

		if RegisterUser(userPayee) && RegisterUser(userPayer) {

			return "Sucess", nil
		}

		return "with connection", err
	}

	return "Insuficient Amount or User is a shopper", err

}

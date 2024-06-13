package models

type User struct {
	Name     string  `json:"name"`
	Cpf      string  `json:"cpf"`
	Mail     string  `json:"mail"`
	Password string  `json:"password"`
	Amount   float64 `json:"amount"`
	Shopper  bool    `json:"shopper"`
}

type TransferBody struct {
	Value float64 `json:"value"`
	Payer string  `json:"payer"`
	Payee string  `json:"payee"`
}

type GetAmountBody struct {
	Cpf string `json:"cpf"`
}

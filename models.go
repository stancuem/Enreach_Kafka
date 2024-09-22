package main

// User представляет модель данных для пользователя.

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Job       string `json:"job"`
	Age       int    `json:"age"`
}

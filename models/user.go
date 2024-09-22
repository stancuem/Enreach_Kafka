package models

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Job       string `json:"job"`
	Age       int    `json:"age"`
}

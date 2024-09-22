package postgres

import (
	"database/sql"
	"log"

	"Go_REST_Kafka/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", "user=madara password=123 dbname=madaradb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserById(id int) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, first_name, last_name, job, age FROM users WHERE id = $1", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Job, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

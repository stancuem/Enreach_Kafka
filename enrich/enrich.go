package enrich

import (
	"Go_REST_Kafka/models"
	"Go_REST_Kafka/postgres"
)

func EnrichUser(user *models.User) (*models.User, error) {
	enrichedUser, err := postgres.GetUserById(user.Id)
	if err != nil {
		return nil, err
	}
	return enrichedUser, nil
}

package tests

import (
	"testing"
)

func TestDatabaseSave(t *testing.T) {
	db, _ := data_enrichment_service.ConnectDatabase("your_database_url")
	defer db.conn.Close()

	user := data_enrichment_service.User{FirstName: "John", LastName: "Doe", Job: "Developer", Age: 30}
	err := db.Save(user)

	if err != nil {
		t.Errorf("Failed to save to database: %v", err)
	}
}

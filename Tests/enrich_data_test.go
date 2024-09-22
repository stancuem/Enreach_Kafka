package tests

import (
	"testing"
)

func TestEnrichData(t *testing.T) {
	raw := data_enrichment_service.User{FirstName: "John", LastName: "Doe"}
	enriched := data_enrichment_service.enrichData(raw)

	if enriched.FirstName != raw.FirstName || enriched.LastName != raw.LastName {
		t.Error("Enriched data does not match raw data")
	}
}

package helpers

import (
	"slices"
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config("config_fixture.json")
	if !slices.Contains(config.BaseUrls, "https://geodienste.ch/stac") ||
		!slices.Contains(config.BaseUrls, "https://data.geo.admin.ch/api/stac/v1") {
		t.Errorf("Base URLs are not correct")
	}
}

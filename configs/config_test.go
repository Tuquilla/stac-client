package configs

import (
	"slices"
	"testing"
)

func TestConfig(t *testing.T) {
	config := ReadConfig()
	if !slices.Contains(config.BaseUrls, "https://geodienste.ch/stac") ||
		!slices.Contains(config.BaseUrls, "https://data.geo.admin.ch/api/stac/v1") {
		t.Errorf("Base URLs are not correct")
	}
}

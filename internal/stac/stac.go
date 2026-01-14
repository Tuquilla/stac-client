package stac

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kglaus/geodienste-cli/internal/stac/models"
)

func GetCollections(baseUrl string) models.Collections {
	if baseUrl == "" {
		var collections models.Collections
		return collections
	}

	resp, err := http.Get(baseUrl + "/collections")
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error calling service")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading Response Body")
	}

	collections := createCollections(body)
	return collections
}

func GetItems(itemUrl string) models.FeatureCollection {
	resp, err := http.Get(itemUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error calling service")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading Response Body")
	}

	featureCollection := createFeatureCollection(body)
	return featureCollection
}

func createCollections(body []byte) models.Collections {
	var collections models.Collections
	err := json.Unmarshal(body, &collections)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s", err)
	}
	return collections
}

func createFeatureCollection(body []byte) models.FeatureCollection {
	var featureCollections models.FeatureCollection
	err := json.Unmarshal(body, &featureCollections)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s", err)
	}
	return featureCollections
}

package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/YoungsoonLee/design-pattern-go/models"
)

// CatBreedsInterface is an adapter that adapts the data to the CatBreed struct
type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
}

// RemoteService is a struct that represents a remote service
type RemoteService struct {
	Remote CatBreedsInterface
}

// GetAllCatBreeds calls the remote service
func (rs *RemoteService) GetAllCatBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

// JSONBackend is a struct that represents a JSON backend
type JSONBackend struct{}

// GetAllCatBreeds gets the data from the JSON backend
func (j *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("https://localhost:8080/api/cat-breeds/all/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var catBreeds []*models.CatBreed
	err = json.Unmarshal(body, &catBreeds)
	if err != nil {
		return nil, err
	}

	return catBreeds, nil
}

// XMLBackend is a struct that represents an XML backend
type XMLBackend struct{}

// GetAllCatBreeds gets the data from the XML backend
func (x *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("https://localhost:8080/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-Breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var c catBreeds
	err = xml.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return c.Breeds, nil
}

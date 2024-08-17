package adapters

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
	GetCatBreedByName(b string) (*models.CatBreed, error)
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

// GetCatBreedByName gets the data from the JSON backend
func (j *JSONBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	resp, err := http.Get("https://localhost:8080/api/cat-breeds/" + b + "/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var catBreed *models.CatBreed
	err = json.Unmarshal(body, &catBreed)
	if err != nil {
		return nil, err
	}

	return catBreed, nil
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

// GetCatBreedByName gets the data from the XML backend
func (x *XMLBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	resp, err := http.Get("https://localhost:8080/api/cat-breeds/" + b + "/xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var catBreed *models.CatBreed
	err = xml.Unmarshal(body, &catBreed)
	if err != nil {
		return nil, err
	}

	return catBreed, nil
}

type TestBackend struct{}

func (j *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{
			ID:      1,
			Breed:   "Siamese",
			Details: "The Siamese cat is one of the first distinctly recognized breeds of Asian cat. Derived from the Wichianmat landrace, one of several varieties of cat native to Thailand (formerly known as Siam), the Siamese became one of the most popular breeds in Europe and North America in the 19th century.",
		},
	}

	return breeds, nil
}

func (j *TestBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	return nil, nil
}

package main

import (
	"os"
	"testing"

	"github.com/YoungsoonLee/design-pattern-go/adapters"
	"github.com/YoungsoonLee/design-pattern-go/config"
)

var testApp application

func TestMain(m *testing.M) {

	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}

	testApp = application{
		App: config.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}

// type TestBackend struct{}

// func (j *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
// 	breeds := []*models.CatBreed{
// 		&models.CatBreed{
// 			ID:      1,
// 			Breed:   "Siamese",
// 			Details: "The Siamese cat is one of the first distinctly recognized breeds of Asian cat. Derived from the Wichianmat landrace, one of several varieties of cat native to Thailand (formerly known as Siam), the Siamese became one of the most popular breeds in Europe and North America in the 19th century.",
// 		},
// 	}

// 	return breeds, nil
// }

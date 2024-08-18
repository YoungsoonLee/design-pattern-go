package pets

import (
	"fmt"

	"github.com/YoungsoonLee/design-pattern-go/config"
	"github.com/YoungsoonLee/design-pattern-go/models"
)

// AnimalInterface is an interface that defines the Show method
type AnimalInterface interface {
	Show() string
}

// DogFromFactory is a struct that represents a dog from the factory
type DogFromFactory struct {
	Pet *models.Dog
}

// Show is a method that returns a string
func (d *DogFromFactory) Show() string {
	return fmt.Sprintf("this animal is a %s", d.Pet.Breed.Breed)
}

// CatFromFactory is a struct that represents a cat from the factory
type CatFromFactory struct {
	Pet *models.Cat
}

// Show is a method that returns a string
func (c *CatFromFactory) Show() string {
	return fmt.Sprintf("this animal is a %s", c.Pet.Breed.Breed)
}

// PetFactoryInterface is an interface that defines the newPet method
type PetFactoryInterface interface {
	newPet() AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

// DogAbstarctFactory is a struct that represents a dog abstract factory
type DogAbstarctFactory struct{}

func (d *DogAbstarctFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (d *DogAbstarctFactory) newPetWithBreed(b string) AnimalInterface {
	app := config.GetInstance()
	breed, _ := app.Models.DogBreed.GetBreedByName(b)
	return &DogFromFactory{
		Pet: &models.Dog{
			Breed: *breed,
		},
	}
}

// CatAbstarctFactory is a struct that represents a cat abstract factory
type CatAbstarctFactory struct{}

func (c *CatAbstarctFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (c *CatAbstarctFactory) newPetWithBreed(b string) AnimalInterface {
	app := config.GetInstance()
	breed, err := app.CatService.Remote.GetCatBreedByName(b)
	if err != nil {
		return nil
	}

	return &CatFromFactory{
		Pet: &models.Cat{
			Breed: *breed,
		},
	}
}

// NewPetFromAbstractFactory is a function that returns a new pet from the factory
func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	var factory PetFactoryInterface

	switch species {
	case "dog":
		factory = &DogAbstarctFactory{}
	case "cat":
		factory = &CatAbstarctFactory{}
	default:
		return nil, fmt.Errorf("invalid species")
	}

	return factory.newPet(), nil
}

// NewPetWithBreedFromAbstractFactory is a function that returns a new pet with a breed from the factory
func NewPetWithBreedFromAbstractFactory(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstarctFactory
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		var catFactory CatAbstarctFactory
		cat := catFactory.newPetWithBreed(breed)
		return cat, nil
	default:
		return nil, fmt.Errorf("invalid species")
	}
}

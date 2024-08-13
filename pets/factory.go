package pets

import "github.com/YoungsoonLee/design-pattern-go/models"

// NewPet is a factory function that returns a new Pet
func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "",
		LifeSpan:    0,
	}

	return &pet
}

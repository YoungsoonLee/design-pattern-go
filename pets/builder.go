package pets

import "errors"

// PetInterface is an interface for a pet
type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(i int) *Pet
	SetMaxWeight(i int) *Pet
	SetWeight(i int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(i int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(i int) *Pet
	SetAgeEstimated(s bool) *Pet
	Build() (*Pet, error)
}

// NewPetBuilder creates a new pet builder
func NewPetBuilder() PetInterface {
	return &Pet{}
}

// SetSpecies sets the species of a pet
func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

// SetBreed sets the breed of a pet
func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}

// SetMinWeight sets the minimum weight of a pet
func (p *Pet) SetMinWeight(i int) *Pet {
	p.MinWeight = i
	return p
}

// SetMaxWeight sets the maximum weight of a pet
func (p *Pet) SetMaxWeight(i int) *Pet {
	p.MaxWeight = i
	return p
}

// SetWeight sets the weight of a pet
func (p *Pet) SetWeight(i int) *Pet {
	p.Weight = i
	return p
}

// SetDescription sets the description of a pet
func (p *Pet) SetDescription(s string) *Pet {
	p.Description = s
	return p
}

// SetLifeSpan sets the lifespan of a pet
func (p *Pet) SetLifeSpan(i int) *Pet {
	p.LifeSpan = i
	return p
}

// SetGeographicOrigin sets the geographic origin of a pet
func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}

// SetColor sets the color of a pet
func (p *Pet) SetColor(s string) *Pet {
	p.Color = s
	return p
}

// SetAge sets the age of a pet
func (p *Pet) SetAge(i int) *Pet {
	p.Age = i
	return p
}

// SetAgeEstimated sets the estimated age of a pet
func (p *Pet) SetAgeEstimated(s bool) *Pet {
	p.AgeEstimated = s
	return p
}

// Build builds a pet
func (p *Pet) Build() (*Pet, error) {
	// validation
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("min weight cannot be greater than max weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}

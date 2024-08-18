package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/YoungsoonLee/design-pattern-go/models"
	"github.com/YoungsoonLee/design-pattern-go/pets"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) DogOfMOnth(w http.ResponseWriter, r *http.Request) {
	// Get the breed
	breed, _ := app.App.Models.DogBreed.GetBreedByName("Golden Retriever")

	// Get the dog of the month from database
	dom, _ := app.App.Models.Dog.GetDogOfMonthByID(breed.ID)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2014-11-01")

	// Create dog and decorate it
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			Name:             "Rex",
			BreedID:          breed.ID,
			Color:            "Golden",
			DateOfBirth:      dob,
			SpayedOrNeutered: true,
			Description:      "A friendly, loyal dog",
			Weight:           65,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	// Serve the web page
	data := make(map[string]any)
	data["dog"] = dog

	app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	// dog := pets.NewPet("dog")

	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	// dog := pets.NewPet("dog")

	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))

}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	breeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, breeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// create a dog using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("Golden Retriever").
		SetWeight(65).
		SetLifeSpan(12).
		SetDescription("A friendly, loyal dog").
		SetColor("Golden").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// create a cat using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("Siamese").
		SetWeight(10).
		SetLifeSpan(15).
		SetDescription("A friendly, loyal  cat").
		SetColor("White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	catBreeds, err := app.App.CatService.GetAllCatBreeds()
	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// Get species from URL itself.
	species := chi.URLParam(r, "species")

	// Get breed from URL itself.
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	// fmt.Println("species:", species, "breed:", breed)

	// Create a pet from the abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	// Return the pet as JSON
	_ = t.WriteJSON(w, http.StatusOK, pet)
}

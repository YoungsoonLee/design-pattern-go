package main

import (
	"fmt"
	"net/http"

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
	// var t toolbox.Tools

	// Get species from URL itself.

	// Get breed from URL itself.

	// Create a pet from the abstract factory

	// Return the pet as JSON
}

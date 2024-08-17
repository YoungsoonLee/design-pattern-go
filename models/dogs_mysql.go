package models

import (
	"context"
	"time"
)

// AllDogBreeds returns all dog breeds
func (d *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs, 
				cast(((weight_low_lbs + weight_high_lbs)/2) as unsigned) as average_weight, 
				lifespan, coalesce(details, ''),  coalesce(alternate_names, ''),
				coalesce(geographic_origin, '') 
				FROM dog_breeds ORDER BY breed`

	var breeds []*DogBreed

	rows, err := d.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var breed DogBreed

		err = rows.Scan(
			&breed.ID,
			&breed.Breed,
			&breed.WwirghtLowLbs,
			&breed.WeightHighLbs,
			&breed.AverageWeight,
			&breed.LifeSpan,
			&breed.Details,
			&breed.AlternateNames,
			&breed.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}

		breeds = append(breeds, &breed)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return breeds, nil
}

// GetBreedByName returns a dog breed by name
func (d *mysqlRepository) GetBreedByName(name string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs, 
				cast(((weight_low_lbs + weight_high_lbs)/2) as unsigned) as average_weight, 
				lifespan, coalesce(details, ''),  coalesce(alternate_names, ''),
				coalesce(geographic_origin, '') 
				FROM dog_breeds WHERE breed = ?`

	var breed DogBreed

	err := d.DB.QueryRowContext(ctx, query, name).Scan(
		&breed.ID,
		&breed.Breed,
		&breed.WwirghtLowLbs,
		&breed.WeightHighLbs,
		&breed.AverageWeight,
		&breed.LifeSpan,
		&breed.Details,
		&breed.AlternateNames,
		&breed.GeographicOrigin,
	)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}

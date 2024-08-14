package config

import (
	"database/sql"
	"sync"

	"github.com/YoungsoonLee/design-pattern-go/models"
)

// AppConfig is the struct for the application configuration
type AppConfig struct {
	Models *models.Models
}

var instance *AppConfig
var once sync.Once
var db *sql.DB

// New creates a new AppConfig instance
func New(pool *sql.DB) *AppConfig {
	db = pool
	return GetInstance()
}

// GetInstance returns the singleton instance of AppConfig
func GetInstance() *AppConfig {
	once.Do(func() {
		instance = &AppConfig{Models: models.New(db)}
	})
	return instance
}

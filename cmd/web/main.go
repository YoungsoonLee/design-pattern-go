package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/YoungsoonLee/design-pattern-go/config"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *config.AppConfig
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := &application{
		templateMap: map[string]*template.Template{},
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use a template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:mypassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "MySQL data source name")
	flag.Parse()

	// get a connection to the database
	db, err := initMySQL(app.config.dsn)
	if err != nil {
		log.Fatal(err)
	}

	app.App = config.New(db)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting server on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

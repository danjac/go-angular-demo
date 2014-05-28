package main

import (
	"github.com/danjac/angular-react-compare/api/csrf"
	"github.com/danjac/angular-react-compare/api/models"
	"github.com/danjac/angular-react-compare/api/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	// DATABASE

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatal("DB_NAME is missing")
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		log.Fatal("DB_USER is missing")
	}

	dbpass := os.Getenv("DB_PASS")
	if dbpass == "" {
		log.Fatal("DB_PASS is missing")
	}

	dbMap, err := models.InitDb(dbname, dbuser, dbpass)
	if err != nil {
		log.Fatal(err)
	}
	defer dbMap.Db.Close()

	// SERVER
	http.Handle("/", csrf.NewCSRF(os.Getenv("SECRET_KEY"), routes.NewRouter()))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}

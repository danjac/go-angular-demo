package main

import (
	"fmt"
	"github.com/danjac/angular-react-compare/api/csrf"
	"github.com/danjac/angular-react-compare/api/models"
	"github.com/danjac/angular-react-compare/api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func getEnvOrDie(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatal(fmt.Sprintf("%s is missing", name))
	}
	return value
}

func getEnvOrDefault(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {

	// get all our env variables

	dbname := getEnvOrDie("DB_NAME")
	dbuser := getEnvOrDie("DB_USER")
	dbpass := getEnvOrDie("DB_PASS")

	logPrefix := getEnvOrDefault("LOG_PREFIX", "myapp")

	secretKey := getEnvOrDie("SECRET_KEY")

	dbMap, err := models.InitDb(dbname, dbuser, dbpass, logPrefix)
	if err != nil {
		log.Fatal(err)
	}
	defer dbMap.Db.Close()

	r := mux.NewRouter()

	// API

	routes.Configure(r, "/api")

	// STATIC FILES

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", csrf.NewCSRF(secretKey, r))

	// SERVER

	port := getEnvOrDefault("PORT", "3000")

	http.ListenAndServe(":"+port, nil)
}

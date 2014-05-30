package main

import (
	"fmt"
	"github.com/danjac/angular-react-compare/api"
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

	app, err := api.NewApp(dbname, dbuser, dbpass, logPrefix,
		secretKey, "/api", "/", "./public/")

	if err != nil {
		log.Fatal(err)
	}
	defer app.Shutdown()

	http.Handle("/", app.Handler)

	// SERVER

	port := getEnvOrDefault("PORT", "3000")

	http.ListenAndServe(":"+port, nil)
}

package main

import (
	"fmt"
	"github.com/danjac/go-angular-demo/api"
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

	config := &api.Config{
		DbName:       getEnvOrDie("DB_NAME"),
		DbUser:       getEnvOrDie("DB_USER"),
		DbPassword:   getEnvOrDie("DB_PASS"),
		LogPrefix:    getEnvOrDefault("LOG_PREFIX", "myapp"),
		SecretKey:    getEnvOrDie("SECRET_KEY"),
		ApiPrefix:    "/api",
		StaticPrefix: "/",
		StaticDir:    "./public/",
	}

	app, err := api.NewApp(config)

	if err != nil {
		log.Fatal(err)
	}
	defer app.Shutdown()

	http.Handle("/", app.Handler)

	// SERVER

	port := getEnvOrDefault("PORT", "3000")

	http.ListenAndServe(":"+port, nil)
}

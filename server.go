package main

import (
	"flag"
	"log"
	"myapp/csrf"
	"myapp/models"
	"myapp/routes"
	"net/http"
	"os"
)

func main() {

	dbName := flag.String("database", "/tmp/tweets.db", "Sqlite database name")
	flag.Parse()

	// DATABASE
	dbMap, err := models.InitDb(*dbName)
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

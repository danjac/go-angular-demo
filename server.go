package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	dbName := flag.String("database", "/tmp/tweets.db", "Sqlite database name")
	flag.Parse()

	// DATABASE
	dbMap, err := InitDb(*dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer dbMap.Db.Close()

	// SERVER
	http.Handle("/", NewCSRF(os.Getenv("SECRET_KEY"), SetupRoutes()))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}

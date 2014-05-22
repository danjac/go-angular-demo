package main

import (
	"flag"
	"github.com/gorilla/context"
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
	r := context.ClearHandler(SetupRoutes())
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}

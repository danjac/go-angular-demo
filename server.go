package main

import (
	"flag"
	"log"
	"net/http"
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
	r := SetupRoutes()
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}

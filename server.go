package main

import (
	"flag"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	dbName := flag.String("database", "/tmp/tweets.db", "Sqlite database name")
	flag.Parse()

	// DATABASE
	dbMap := InitDb(*dbName)
	defer dbMap.Db.Close()

	// SERVER
	r := SetupRoutes()
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}

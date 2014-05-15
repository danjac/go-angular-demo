package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
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
	m := martini.Classic()
	m.Use(render.Renderer())

	SetupRoutes(m)

	m.Run()
}

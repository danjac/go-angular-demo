package main

import (
	"flag"
	"github.com/coopernurse/gorp"
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
	dbMap := initDb(*dbName)
	defer dbMap.Db.Close()

	// SERVER
	m := martini.Classic()
	m.Use(render.Renderer())
	m.MapTo(dbMap, (*gorp.SqlExecutor)(nil))

	addRoutes(m)

	m.Run()
}

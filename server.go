package main

import (
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

	// DATABASE
	dbMap := initDb()
	defer dbMap.Db.Close()

	// SERVER
	m := martini.Classic()
	m.Use(render.Renderer())
	m.MapTo(dbMap, (*gorp.SqlExecutor)(nil))

	addRoutes(m)

	m.Run()
}

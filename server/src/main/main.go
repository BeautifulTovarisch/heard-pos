package main

import (
	"flag"
	"fmt"
	"net/http"

	"heard/database"
	"heard/routes"
)

func main() {
	run_migration := flag.Bool("migrate", false, "Run schema migration.")
	flag.Parse()

	if *run_migration {
		database.SetupSchema()
	} else {
		http.ListenAndServe("0.0.0.0:2305", routes.Routes())



	}
}

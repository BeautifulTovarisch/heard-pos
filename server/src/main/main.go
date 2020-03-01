package main

import (
	"flag"
	"net/http"
	"os"

	"heard/database"
	"heard/routes"
)

func main() {
	run_migration := flag.Bool("migrate", false, "Run schema migration.")
	flag.Parse()

	if *run_migration {
		database.SetupSchema()
	} else {
		conn := database.Connect("pos", "pos", os.Getenv("USER_PASSWORD"))
		defer conn.Close()

		http.ListenAndServe("0.0.0.0:2305", routes.Routes(conn))
	}
}

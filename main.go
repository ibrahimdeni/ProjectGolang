package main

import (
	"dumbflix/routes"
	"fmt"
	"net/http"

	"dumbflix/database"
	"dumbflix/pkg/mysql"

	"github.com/gorilla/mux"
)

func main() {

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

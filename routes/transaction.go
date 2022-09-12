package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	r.HandleFunc("/films", h.FindTransacations).Methods("GET")         //get alll
	r.HandleFunc("/film/{id}", h.GetTransaction).Methods("GET")       //select
	r.HandleFunc("/film", h.CreateTransaction).Methods("POST")        // add
	r.HandleFunc("/film/{id}", h.UpdateTransaction).Methods("PATCH")  // edite
	r.HandleFunc("/film/{id}", h.DeleteTransaction).Methods("DELETE") // delete

}

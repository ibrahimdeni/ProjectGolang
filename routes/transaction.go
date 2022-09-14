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

	r.HandleFunc("/transactions", h.FindTransacations).Methods("GET")         //get alll
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")       //select
	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")        // add
	r.HandleFunc("/transaction/{id}", h.UpdateTransaction).Methods("PATCH")  // edite
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE") // delete

}

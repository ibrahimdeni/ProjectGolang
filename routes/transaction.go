package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	r.HandleFunc("/transactions", h.FindTransacations).Methods("GET")         					//get alll
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")       					//select
	r.HandleFunc("/transaction",   middleware.Auth(middleware.UploadFile(h.CreateTransaction))).Methods("POST")        	// add
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")  	// edite
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE") 	// delete

}

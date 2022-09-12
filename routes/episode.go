package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(EpisodeRepository)

	r.HandleFunc("/films", h.FindEpisodes).Methods("GET")         //get alll
	r.HandleFunc("/film/{id}", h.GetEpisode).Methods("GET")       //select
	r.HandleFunc("/film", h.CreateEpisode).Methods("POST")        // add
	r.HandleFunc("/film/{id}", h.UpdateEpisode).Methods("PATCH")  // edite
	r.HandleFunc("/film/{id}", h.DeleteEpisode).Methods("DELETE") // delete

}

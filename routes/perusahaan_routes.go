package routes

import (
	"github.com/gorilla/mux"
	"github.com/gusetiawn/syamil-group/handlers"
	"github.com/gusetiawn/syamil-group/services"
)

func SetupPerusahaanRoutes(service services.PerusahaanService) *mux.Router {
	router := mux.NewRouter()

	perusahaanHandler := handlers.NewPerusahaanHandler(service)

	router.HandleFunc("/perusahaan", perusahaanHandler.CreatePerusahaan).Methods("POST")
	router.HandleFunc("/perusahaan", perusahaanHandler.GetPerusahaans).Methods("GET")
	router.HandleFunc("/perusahaan/{id:[0-9]+}", perusahaanHandler.UpdatePerusahaan).Methods("PUT")
	router.HandleFunc("/perusahaan/{id:[0-9]+}", perusahaanHandler.DeletePerusahaan).Methods("DELETE")

	return router
}

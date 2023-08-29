package routes

import (
	"github.com/gorilla/mux"
	"github.com/gusetiawn/syamil-group/handlers"
	"github.com/gusetiawn/syamil-group/services"
)

func SetupRoutes(service services.BarangService) *mux.Router {
	router := mux.NewRouter()

	barangHandler := handlers.NewBarangHandler(service)

	router.HandleFunc("/barang", barangHandler.CreateBarang).Methods("POST")
	router.HandleFunc("/barang", barangHandler.GetBarangs).Methods("GET")
	router.HandleFunc("/barang/{id:[0-9]+}", barangHandler.UpdateBarang).Methods("PUT")
	router.HandleFunc("/barang/{id:[0-9]+}", barangHandler.DeleteBarang).Methods("DELETE")

	return router
}

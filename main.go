package main

import (
	"fmt"
	"github.com/gusetiawn/syamil-group/db"
	"github.com/gusetiawn/syamil-group/repositories"
	"github.com/gusetiawn/syamil-group/routes"
	"github.com/gusetiawn/syamil-group/services"
	"net/http"
)

func main() {
	db.InitDB()

	barangRepo := repositories.NewBarangRepository(db.DB)
	barangService := services.NewBarangService(*barangRepo)
	barangRouter := routes.SetupRoutes(*barangService)

	perusahaanRepo := repositories.NewPerusahaanRepository(db.DB)
	perusahaanService := services.NewPerusahaanService(*perusahaanRepo)
	perusahaanRouter := routes.SetupPerusahaanRoutes(*perusahaanService)

	router := http.NewServeMux()
	router.Handle("/barang", barangRouter)
	router.Handle("/perusahaan", perusahaanRouter)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
	fmt.Println("Server Stopped")
}

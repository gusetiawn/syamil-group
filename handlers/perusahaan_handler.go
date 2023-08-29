package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gusetiawn/syamil-group/models"
	"github.com/gusetiawn/syamil-group/services"
	"net/http"
	"strconv"
)

type PerusahaanHandler struct {
	PerusahaanService services.PerusahaanService
}

func NewPerusahaanHandler(service services.PerusahaanService) *PerusahaanHandler {
	return &PerusahaanHandler{PerusahaanService: service}
}

func (h *PerusahaanHandler) CreatePerusahaan(w http.ResponseWriter, r *http.Request) {
	var perusahaan models.Perusahaan
	json.NewDecoder(r.Body).Decode(&perusahaan)

	err := h.PerusahaanService.CreatePerusahaan(perusahaan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PerusahaanHandler) GetPerusahaans(w http.ResponseWriter, r *http.Request) {
	perusahaans, err := h.PerusahaanService.GetPerusahaans()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(perusahaans)
}

func (h *PerusahaanHandler) UpdatePerusahaan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	perusahaanID, _ := strconv.Atoi(vars["id"])

	var perusahaan models.Perusahaan
	json.NewDecoder(r.Body).Decode(&perusahaan)

	err := h.PerusahaanService.UpdatePerusahaan(perusahaanID, perusahaan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PerusahaanHandler) DeletePerusahaan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	perusahaanID, _ := strconv.Atoi(vars["id"])

	err := h.PerusahaanService.DeletePerusahaan(perusahaanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

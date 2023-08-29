package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gusetiawn/syamil-group/models"
	"github.com/gusetiawn/syamil-group/services"
	"net/http"
	"strconv"
)

type BarangHandler struct {
	BarangService services.BarangService
}

func NewBarangHandler(service services.BarangService) *BarangHandler {
	return &BarangHandler{BarangService: service}
}

func (h *BarangHandler) CreateBarang(w http.ResponseWriter, r *http.Request) {
	var barang models.Barang
	json.NewDecoder(r.Body).Decode(&barang)

	err := h.BarangService.CreateBarang(barang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BarangHandler) GetBarangs(w http.ResponseWriter, r *http.Request) {
	barangs, err := h.BarangService.GetBarangs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangs)
}

func (h *BarangHandler) UpdateBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barangID, _ := strconv.Atoi(vars["id"])

	var barang models.Barang
	json.NewDecoder(r.Body).Decode(&barang)

	err := h.BarangService.UpdateBarang(barangID, barang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BarangHandler) DeleteBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barangID, _ := strconv.Atoi(vars["id"])

	err := h.BarangService.DeleteBarang(barangID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

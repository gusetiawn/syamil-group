// services/barang_service.go
package services

import (
	"github.com/gusetiawn/syamil-group/models"
	"github.com/gusetiawn/syamil-group/repositories"
)

type BarangService struct {
	BarangRepo repositories.BarangRepository
}

func NewBarangService(repo repositories.BarangRepository) *BarangService {
	return &BarangService{BarangRepo: repo}
}

func (s *BarangService) CreateBarang(barang models.Barang) error {
	return s.BarangRepo.CreateBarang(barang)
}

func (s *BarangService) GetBarangs() ([]models.Barang, error) {
	return s.BarangRepo.GetBarangs()
}

func (s *BarangService) UpdateBarang(barangID int, barang models.Barang) error {
	return s.BarangRepo.UpdateBarang(barangID, barang)
}

func (s *BarangService) DeleteBarang(barangID int) error {
	return s.BarangRepo.DeleteBarang(barangID)
}

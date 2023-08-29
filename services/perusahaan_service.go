package services

import (
	"github.com/gusetiawn/syamil-group/models"
	"github.com/gusetiawn/syamil-group/repositories"
)

type PerusahaanService struct {
	PerusahaanRepo repositories.PerusahaanRepository
}

func NewPerusahaanService(repo repositories.PerusahaanRepository) *PerusahaanService {
	return &PerusahaanService{PerusahaanRepo: repo}
}

func (s *PerusahaanService) CreatePerusahaan(perusahaan models.Perusahaan) error {
	return s.PerusahaanRepo.CreatePerusahaan(perusahaan)
}

func (s *PerusahaanService) GetPerusahaans() ([]models.Perusahaan, error) {
	return s.PerusahaanRepo.GetPerusahaans()
}

func (s *PerusahaanService) UpdatePerusahaan(perusahaanID int, perusahaan models.Perusahaan) error {
	return s.PerusahaanRepo.UpdatePerusahaan(perusahaanID, perusahaan)
}

func (s *PerusahaanService) DeletePerusahaan(perusahaanID int) error {
	return s.PerusahaanRepo.DeletePerusahaan(perusahaanID)
}

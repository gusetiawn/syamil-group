package repositories

import (
	"database/sql"
	"github.com/gusetiawn/syamil-group/models"
)

type PerusahaanRepository struct {
	DB *sql.DB
}

func NewPerusahaanRepository(db *sql.DB) *PerusahaanRepository {
	return &PerusahaanRepository{DB: db}
}

func (r *PerusahaanRepository) CreatePerusahaan(perusahaan models.Perusahaan) error {
	_, err := r.DB.Exec("INSERT INTO Perusahaan (namaPerusahaan, alamat) VALUES ($1, $2)",
		perusahaan.NamaPerusahaan, perusahaan.Alamat)
	return err
}

func (r *PerusahaanRepository) GetPerusahaans() ([]models.Perusahaan, error) {
	rows, err := r.DB.Query("SELECT * FROM Perusahaan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var perusahaans []models.Perusahaan
	for rows.Next() {
		var perusahaan models.Perusahaan
		err := rows.Scan(&perusahaan.PerusahaanID, &perusahaan.NamaPerusahaan, &perusahaan.Alamat)
		if err != nil {
			return nil, err
		}
		perusahaans = append(perusahaans, perusahaan)
	}
	return perusahaans, nil
}

func (r *PerusahaanRepository) UpdatePerusahaan(perusahaanID int, perusahaan models.Perusahaan) error {
	_, err := r.DB.Exec("UPDATE Perusahaan SET namaPerusahaan=$1, alamat=$2 WHERE perusahaanID=$3",
		perusahaan.NamaPerusahaan, perusahaan.Alamat, perusahaanID)
	return err
}

func (r *PerusahaanRepository) DeletePerusahaan(perusahaanID int) error {
	_, err := r.DB.Exec("DELETE FROM Perusahaan WHERE perusahaanID=$1", perusahaanID)
	return err
}

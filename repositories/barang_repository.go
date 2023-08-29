package repositories

import (
	"database/sql"
	"github.com/gusetiawn/syamil-group/models"
)

type BarangRepository struct {
	DB *sql.DB
}

func NewBarangRepository(db *sql.DB) *BarangRepository {
	return &BarangRepository{DB: db}
}

func (r *BarangRepository) CreateBarang(barang models.Barang) error {
	_, err := r.DB.Exec("INSERT INTO Barang (namaBarang, harga, stok) VALUES ($1, $2, $3)",
		barang.NamaBarang, barang.Harga, barang.Stok)
	return err
}

func (r *BarangRepository) GetBarangs() ([]models.Barang, error) {
	rows, err := r.DB.Query("SELECT * FROM Barang")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barangs []models.Barang
	for rows.Next() {
		var barang models.Barang
		err := rows.Scan(&barang.BarangID, &barang.NamaBarang, &barang.Harga, &barang.Stok)
		if err != nil {
			return nil, err
		}
		barangs = append(barangs, barang)
	}
	return barangs, nil
}

func (r *BarangRepository) UpdateBarang(barangID int, barang models.Barang) error {
	_, err := r.DB.Exec("UPDATE Barang SET namaBarang=$1, harga=$2, stok=$3 WHERE barangID=$4",
		barang.NamaBarang, barang.Harga, barang.Stok, barangID)
	return err
}

func (r *BarangRepository) DeleteBarang(barangID int) error {
	_, err := r.DB.Exec("DELETE FROM Barang WHERE barangID=$1", barangID)
	return err
}

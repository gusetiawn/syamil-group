package models

type User struct {
	UserID           int
	Username         string
	Email            string
	PasswordHash     string
	RegistrationDate string
}

type Barang struct {
	BarangID   int
	NamaBarang string
	Harga      float64
	Stok       int
}

type Perusahaan struct {
	PerusahaanID   int
	NamaPerusahaan string
	Alamat         string
}

// Define struct models for other tables as well (Barang, Perusahaan, Transaksi, Report)

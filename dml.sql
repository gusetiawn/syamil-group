-- Dummy data untuk tabel User
INSERT INTO "User" (username, email, passwordHash, registrationDate)
VALUES
    ('agus', 'agus@example.com', 'hashed_password_1', '2023-08-01'),
    ('setiawan', 'setiawan@example.com', 'hashed_password_2', '2023-08-05');

-- Dummy data untuk tabel Barang
INSERT INTO Barang (namaBarang, harga, stok)
VALUES
    ('Laptop', 1000.00, 10),
    ('Smartphone', 600.00, 20);

-- Dummy data untuk tabel Perusahaan
INSERT INTO Perusahaan (namaPerusahaan, alamat, telepon)
VALUES
    ('ABC Electronics', 'Rancaekek City', '123-456-7890'),
    ('XYZ Gadgets', 'Letnan Adun', '987-654-3210');

-- Dummy data untuk tabel Transaksi
INSERT INTO Transaksi (userID, barangID, tanggalTransaksi, jumlahBeli, totalHarga)
VALUES
    (1, 1, '2023-08-10', 2, 2000.00),
    (2, 2, '2023-08-12', 3, 1800.00);

-- Dummy data untuk tabel Report
INSERT INTO Report (transaksiID, keterangan, tanggalLaporan)
VALUES
    (1, 'Transaksi ini berhasil', '2023-08-11'),
    (2, 'Transaksi ini memiliki beberapa masalah', '2023-08-14');

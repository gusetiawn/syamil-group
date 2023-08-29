-- Tabel User
CREATE TABLE "User" (
                        userID SERIAL PRIMARY KEY,
                        username VARCHAR(50) NOT NULL,
                        email VARCHAR(100) NOT NULL,
                        passwordHash VARCHAR(128) NOT NULL,
                        registrationDate DATE
);

-- Tabel Barang
CREATE TABLE Barang (
                        barangID SERIAL PRIMARY KEY,
                        namaBarang VARCHAR(100) NOT NULL,
                        harga DECIMAL(10, 2) NOT NULL,
                        stok INT NOT NULL
);

-- Tabel Perusahaan
CREATE TABLE Perusahaan (
                            perusahaanID SERIAL PRIMARY KEY,
                            namaPerusahaan VARCHAR(100) NOT NULL,
                            alamat VARCHAR(200),
                            telepon VARCHAR(20)
);

-- Tabel Transaksi
CREATE TABLE Transaksi (
                           transaksiID SERIAL PRIMARY KEY,
                           userID INT,
                           barangID INT,
                           tanggalTransaksi DATE,
                           jumlahBeli INT,
                           totalHarga DECIMAL(10, 2),
                           FOREIGN KEY (userID) REFERENCES "User"(userID),
                           FOREIGN KEY (barangID) REFERENCES Barang(barangID)
);

-- Tabel Report
CREATE TABLE Report (
                        reportID SERIAL PRIMARY KEY,
                        transaksiID INT,
                        keterangan TEXT,
                        tanggalLaporan DATE,
                        FOREIGN KEY (transaksiID) REFERENCES Transaksi(transaksiID)
);

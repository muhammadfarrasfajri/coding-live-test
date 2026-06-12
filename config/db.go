package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// 1. Konfigurasi Koneksi
	// PERHATIAN: Ganti "sandi_baru_anda" dengan password yang Anda buat sebelumnya!
	connStr := "host=localhost port=5432 user=postgres password=farras dbname=shope sslmode=disable"

	// 2. Membuka "Gerbang" Koneksi
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal membuka koneksi: ", err)
	}

	// Pastikan koneksi selalu ditutup saat program selesai berjalan

	// 3. Menguji Koneksi (Ping)
	// sql.Open hanya mendaftarkan konfigurasi, Ping() yang benar-benar mencoba terhubung
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database! Cek password atau layanan PostgreSQL Anda. Error: ", err)
	}

	fmt.Println("🎉 Berhasil terhubung ke PostgreSQL!")

	// 4. (Opsional) Percobaan menjalankan query SQL sederhana
	var dbVersion string
	err = db.QueryRow("SELECT version()").Scan(&dbVersion)
	if err != nil {
		log.Fatal("Gagal menjalankan query: ", err)
	}

	fmt.Println("Versi Database yang Anda gunakan:\n", dbVersion)
	return db
}

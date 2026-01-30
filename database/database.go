package database

import (
	"database/sql"
	"log"

	"net/url"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Debug: Cek apakah URL ter-parse dengan benar
	u, err := url.Parse(connectionString)
	if err != nil {
		log.Println("ERROR: Connection string tidak valid:", err)
	} else {
		// Log host untuk memastikan kita tidak salah connect (misal ke bagian password)
		// Log query params untuk memastikan sslmode terbaca
		log.Printf("Mencoba connect ke Host: %s, User: %s, Params: %s", u.Host, u.User.Username(), u.RawQuery)
	}

	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}

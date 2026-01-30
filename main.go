package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

// ubah Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()

	if _, err := os.Stat(".env"); err != nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection: Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandlers := handlers.NewProductHandler(productService)

	// Dependency Injection: Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandlers := handlers.NewCategoryHandler(categoryService)

	// Routes: Product
	http.HandleFunc("/api/produk", productHandlers.HandleProducts)
	http.HandleFunc("/api/produk/", productHandlers.HandleProductByID)

	// Routes: Category
	http.HandleFunc("/categories", categoryHandlers.HandleCategories)
	http.HandleFunc("/categories/", categoryHandlers.HandleCategoryByID)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		docs := map[string]interface{}{
			"_message_": "Dokumentasi API Kasir",
			"endpoints": []map[string]string{
				{"method": "GET", "url": "/categories", "description": "Ambil semua kategori"},
				{"method": "GET", "url": "/categories/{id}", "description": "Ambil detail kategori"},
				{"method": "POST", "url": "/categories", "description": "Tambah kategori baru"},
				{"method": "PUT", "url": "/categories/{id}", "description": "Update kategori"},
				{"method": "DELETE", "url": "/categories/{id}", "description": "Hapus kategori"},
				{"method": "GET", "url": "/api/produk", "description": "Ambil semua produk"},
				{"method": "POST", "url": "/api/produk", "description": "Tambah produk baru"},
			},
		}

		json.NewEncoder(w).Encode(docs)
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server berjalan di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

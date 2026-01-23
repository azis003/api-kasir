package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{}
var nextID = 1

func main() {
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(categories)

		case "POST":
			var newCategory Category
			if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
				http.Error(w, "Data JSON salah", http.StatusBadRequest)
				return
			}
			newCategory.ID = nextID
			nextID++
			categories = append(categories, newCategory)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCategory)

		default:
			http.Error(w, "Method salah", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		index := -1
		for i, cat := range categories {
			if cat.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
			return
		}

		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(categories[index])

		case "PUT":
			var updatedData Category
			if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
				http.Error(w, "Data JSON salah", http.StatusBadRequest)
				return
			}
			categories[index].Name = updatedData.Name
			categories[index].Description = updatedData.Description

			categories[index].ID = id

			json.NewEncoder(w).Encode(categories[index])

		case "DELETE":
			categories = append(categories[:index], categories[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Kategori berhasil dihapus"})

		default:
			http.Error(w, "Method salah", http.StatusMethodNotAllowed)
		}
	})

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
			},
		}

		json.NewEncoder(w).Encode(docs)
	})

	fmt.Println("Server berjalan di http://localhost:8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

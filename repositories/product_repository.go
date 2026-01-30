package repositories

import (
	"database/sql"
	"kasir-api/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) Create(p *model.Product) error {
	err := r.DB.QueryRow(
		"INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id",
		p.Name, p.Price, p.Stock,
	).Scan(&p.ID)
	return err
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	var p model.Product
	err := r.DB.QueryRow("SELECT id, name, price, stock FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(p *model.Product) error {
	_, err := r.DB.Exec(
		"UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4",
		p.Name, p.Price, p.Stock, p.ID,
	)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

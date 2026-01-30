package repositories

import (
	"database/sql"
	"kasir-api/model"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.DB.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id int) (*model.Category, error) {
	var c model.Category
	err := r.DB.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).
		Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Create(c *model.Category) error {
	err := r.DB.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id",
		c.Name, c.Description,
	).Scan(&c.ID)
	return err
}

func (r *CategoryRepository) Update(c *model.Category) error {
	_, err := r.DB.Exec(
		"UPDATE categories SET name = $1, description = $2 WHERE id = $3",
		c.Name, c.Description, c.ID,
	)
	return err
}

func (r *CategoryRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	return err
}

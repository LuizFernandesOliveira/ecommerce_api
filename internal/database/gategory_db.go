package database

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (c *CategoryDB) CreateCategory(category *entity.Category) (*entity.Category, error) {
	_, err := c.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	row := c.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id)
	var category entity.Category
	if err := row.Scan(&category.ID, &category.Name); err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (c *CategoryDB) DeleteCategory(categoryID string) error {
	_, err := c.db.Exec("DELETE FROM categories WHERE id = ?", categoryID)
	if err != nil {
		return err
	}
	return nil
}

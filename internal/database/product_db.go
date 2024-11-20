package database

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := p.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDB) GetProduct(id string) (*entity.Product, error) {
	row := p.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id)
	var product entity.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) GetProductsByCategoryId(categoryID string) ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT id, name, description, price, category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

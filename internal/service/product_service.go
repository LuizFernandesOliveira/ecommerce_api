package service

import (
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/database"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(db database.ProductDB) *ProductService {
	return &ProductService{ProductDB: db}
}

func (p *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)
	_, err := p.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := p.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) GetProductByCategoryId(categoryID string) ([]*entity.Product, error) {
	products, err := p.ProductDB.GetProductsByCategoryId(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := p.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

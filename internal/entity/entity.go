package entity

import (
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func NewProduct(name, description, categoryID, imageURL string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}

type Pagination struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
	Items interface{} `json:"items"`
}

func NewPagination(r *http.Request) *Pagination {
	query := r.URL.Query()
	pageStr := query.Get("page")
	sizeStr := query.Get("size")

	page := 1
	size := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}
	if sizeStr != "" {
		if s, err := strconv.Atoi(sizeStr); err == nil {
			size = s
		}
	}

	if size <= 0 {
		size = 10
	}
	if page <= 0 {
		page = 1
	}

	return &Pagination{
		Page: page,
		Size: size,
	}
}

func (p *Pagination) SetItems(items interface{}) *Pagination {
	p.Items = items
	return p
}

func (p *Pagination) SetTotal(total int) *Pagination {
	p.Total = total
	return p
}

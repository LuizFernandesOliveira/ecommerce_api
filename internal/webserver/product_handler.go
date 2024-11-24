package webserver

import (
	"encoding/json"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/entity"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/service"
	"github.com/go-chi/chi"
	"net/http"
)

type WebProductHandler struct {
	ProductService service.ProductService
}

func NewWebProductHandler(ps service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: ps}
}

func (h *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdProduct, err := h.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

func (h *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := h.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pagination := entity.NewPagination(r)
	query := r.URL.Query()
	categoryID := query.Get("category_id")
	pagination, err := h.ProductService.GetProducts(pagination, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pagination)
}

func (h *WebProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "categoryID is required in the URL", http.StatusBadRequest)
		return
	}
	products, err := h.ProductService.GetProductByCategoryId(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

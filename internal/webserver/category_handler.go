package webserver

import (
	"encoding/json"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/entity"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/service"
	"github.com/go-chi/chi"
	"net/http"
)

type WebCategoryHandler struct {
	CategoryService service.CategoryService
}

func NewWebCategoryHandler(cs service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: cs}
}

func (h *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdCategory, err := h.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategory)
}

func (h *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	category, err := h.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (h *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

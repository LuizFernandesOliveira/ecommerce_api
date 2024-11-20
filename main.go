package main

import (
	"database/sql"
	"fmt"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/database"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/service"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)
	webCategoryHandler := webserver.NewWebCategoryHandler(*categoryService)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)
	webProducHandler := webserver.NewWebProductHandler(*productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/categories/{id}", webCategoryHandler.GetCategory)
	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Post("/categories", webCategoryHandler.CreateCategory)

	c.Get("/products/{id}", webProducHandler.GetProduct)
	c.Get("/products", webProducHandler.GetProducts)
	c.Get("/products/category/{categoryID}", webProducHandler.GetProductsByCategory)
	c.Post("/products", webProducHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
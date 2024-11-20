package main

import (
	"database/sql"
	"fmt"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/database"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/service"
	"github.com/LuizFernandesOliveira/ecommerce_api/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

func main() {
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var databaseName = os.Getenv("DB_NAME")
	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, databaseName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connected")
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
	c.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
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

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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Api-Key"},
		AllowCredentials: true,
	}))
	c.Get("/categories/{id}", webCategoryHandler.GetCategory)
	c.Get("/categories", webCategoryHandler.GetCategories)

	c.Get("/products/{id}", webProducHandler.GetProduct)
	c.Get("/products", webProducHandler.GetProducts)

	c.Group(func(r chi.Router) {
		r.Use(ApiKeyForPostMiddleware)
		r.Post("/categories", webCategoryHandler.CreateCategory)
		r.Delete("/categories/{id}", webCategoryHandler.DeleteCategory)
		r.Post("/products", webProducHandler.CreateProduct)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}

func ApiKeyForPostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Api-Key")

		var apiKeyEnv = os.Getenv("API_KEY")
		if apiKey != apiKeyEnv {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

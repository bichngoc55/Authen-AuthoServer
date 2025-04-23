package main

import (
	"AUTHEN-AUTHOSERVER/internal/cache"
	"AUTHEN-AUTHOSERVER/internal/db"
	handlers "AUTHEN-AUTHOSERVER/internal/handler"
	"AUTHEN-AUTHOSERVER/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	cache.InitRedis()

	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/products").Subrouter()
	protected.Use(middleware.JWTAuth)
	protected.HandleFunc("", handlers.GetAllProductsHandler).Methods("GET")
	protected.HandleFunc("", handlers.CreateProductHandler).Methods("POST")
	protected.HandleFunc("/{id}", handlers.GetProductByIDHandler).Methods("GET")
	protected.HandleFunc("/{id}", handlers.UpdateProductHandler).Methods("PUT")
	protected.HandleFunc("/{id}", handlers.DeleteProductHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
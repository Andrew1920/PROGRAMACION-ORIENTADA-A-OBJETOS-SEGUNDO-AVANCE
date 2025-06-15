package routes

import (
	"net/http"
	"tienda/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, productHandlers *handlers.ProductHandlers, cartHandlers *handlers.CartHandlers) {
	// Rutas de Usuarios y Productos
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/api/products", productHandlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products", productHandlers.CreateProductHandler).Methods("POST")
	r.HandleFunc("/api/products/batch", productHandlers.CreateProductsBatchHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", productHandlers.GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", productHandlers.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", productHandlers.DeleteProductHandler).Methods("DELETE")

	//RUTAS PARA EL CARRITO
	r.HandleFunc("/api/cart", cartHandlers.CreateCartHandler).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}", cartHandlers.ViewCartHandler).Methods("GET")
	r.HandleFunc("/api/cart/{cartId}/add", cartHandlers.AddToCartHandler).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}", cartHandlers.DeleteCartHandler).Methods("DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido a la API de E-Commerce! Tienda de tecnologia"))
	}).Methods("GET")
}

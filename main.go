package main

import (
	"log"
	"net/http"
	"tienda/handlers"
	"tienda/routes"
	"tienda/storage"

	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStore()
	productHandlers := handlers.NewProductHandlers(store)
	// Creamos la instancia de los handlers del carrito, inyectando el almacén de productos.
	cartHandlers := handlers.NewCartHandlers(store)

	r := mux.NewRouter()
	// Pasamos ambos grupos de handlers a las rutas.
	routes.RegisterRoutes(r, productHandlers, cartHandlers)

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}

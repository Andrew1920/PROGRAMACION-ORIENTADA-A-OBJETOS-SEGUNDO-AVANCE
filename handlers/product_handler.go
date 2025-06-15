package handlers

import (
	"encoding/json"
	"net/http"
	"tienda/models"
	"tienda/storage" // Importamos el nuevo paquete

	"github.com/gorilla/mux"
)

// ProductHandlers ahora contiene el almacén (la interfaz) como una dependencia.
type ProductHandlers struct {
	store storage.ProductStorer
}

// NewProductHandlers es un "constructor" que crea una nueva instancia de los handlers.
// Recibe la interfaz del almacén, haciendo que el handler sea independiente de la implementación.
func NewProductHandlers(s storage.ProductStorer) *ProductHandlers {
	return &ProductHandlers{store: s}
}

// GetProductsHandler ahora es un método de ProductHandlers.
func (h *ProductHandlers) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Llama al método de la interfaz, sin saber si es un mapa o una base de datos.
	products, err := h.store.Get()
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CreateProductHandler también es un método y usa la interfaz.
func (h *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.store.Create(product)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

// Y así con todos los demás...
func (h *ProductHandlers) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	product, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandlers) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.store.Update(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

func (h *ProductHandlers) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandlers) CreateProductsBatchHandler(w http.ResponseWriter, r *http.Request) {
	var newProducts []models.Product
	if err := json.NewDecoder(r.Body).Decode(&newProducts); err != nil {
		http.Error(w, "Datos inválidos, se esperaba una lista de productos", http.StatusBadRequest)
		return
	}

	createdProducts, err := h.store.CreateBatch(newProducts)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProducts)
}

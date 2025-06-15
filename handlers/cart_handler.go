package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"tienda/models"
	"tienda/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Usamos un bloque var() para declarar tanto el mapa de carritos
// como el mutex que lo protegerá.
var (
	carts = make(map[string]*models.Cart)
	// Este es el "candado" específico para el mapa de carritos.
	cartMutex sync.Mutex
)

// CartHandlers necesita acceso al almacén de productos para obtener sus detalles (precio, nombre).
type CartHandlers struct {
	productStore storage.ProductStorer
}

// NewCartHandlers es el constructor para los handlers del carrito.
func NewCartHandlers(ps storage.ProductStorer) *CartHandlers {
	return &CartHandlers{productStore: ps}
}

// CreateCartHandler crea un nuevo carrito de compras vacío y devuelve su ID.
func (h *CartHandlers) CreateCartHandler(w http.ResponseWriter, r *http.Request) {
	cartMutex.Lock()
	defer cartMutex.Unlock()

	cartID := uuid.NewString()
	newCart := &models.Cart{
		ID:    cartID,
		Items: []models.CartItem{},
		Total: 0,
	}
	carts[cartID] = newCart

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCart)
}

// AddToCartHandler añade un producto a un carrito específico.
func (h *CartHandlers) AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cartId"]

	var req struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if req.Quantity <= 0 {
		http.Error(w, "La cantidad debe ser positiva", http.StatusBadRequest)
		return
	}

	product, err := h.productStore.GetByID(req.ProductID)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	cartMutex.Lock()
	defer cartMutex.Unlock()

	cart, ok := carts[cartID]
	if !ok {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	found := false
	for i, item := range cart.Items {
		if item.ProductID == req.ProductID {
			cart.Items[i].Quantity += req.Quantity
			found = true
			break
		}
	}
	if !found {
		newItem := models.CartItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Quantity:    req.Quantity,
			Price:       product.Price,
		}
		cart.Items = append(cart.Items, newItem)
	}

	var total float64
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	cart.Total = total

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

// ViewCartHandler muestra el contenido de un carrito específico.
func (h *CartHandlers) ViewCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cartId"]

	cartMutex.Lock()
	defer cartMutex.Unlock()

	cart, ok := carts[cartID]
	if !ok {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// DeleteCartHandler simula una compra o vaciado, eliminando el carrito.
func (h *CartHandlers) DeleteCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cartId"]

	cartMutex.Lock()
	defer cartMutex.Unlock()

	if _, ok := carts[cartID]; !ok {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}
	delete(carts, cartID)
	w.WriteHeader(http.StatusNoContent)
}

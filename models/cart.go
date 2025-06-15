package models

// CartItem representa un único artículo dentro del carrito de compras.
type CartItem struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

// Cart representa el carrito de compras completo, ahora identificado por su propio ID.
type Cart struct {
	ID    string     `json:"id"` // Cambiamos UserID por ID
	Items []CartItem `json:"items"`
	Total float64    `json:"total"`
}

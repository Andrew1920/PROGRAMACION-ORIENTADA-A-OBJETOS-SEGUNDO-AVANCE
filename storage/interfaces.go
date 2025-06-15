package storage

import "tienda/models"

// Aquí defino mi interfaz ProductStorer. Para mí, es un "contrato" que
// establece las operaciones obligatorias para cualquier tipo de almacenamiento
// de productos que yo quiera crear en el futuro.
type ProductStorer interface {
	Create(product models.Product) (models.Product, error)
	CreateBatch(products []models.Product) ([]models.Product, error)
	Get() ([]models.Product, error)
	GetByID(id string) (models.Product, error)
	Update(id string, product models.Product) (models.Product, error)
	Delete(id string) error
}

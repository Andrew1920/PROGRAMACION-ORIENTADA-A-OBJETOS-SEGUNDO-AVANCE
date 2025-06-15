package storage

import (
	"fmt"
	"sync"
	"tienda/models"

	"github.com/google/uuid"
)

// MemoryStore es mi implementación de ProductStorer.
// Contiene los datos y el mutex para manejar la concurrencia de forma segura.
type MemoryStore struct {
	data  map[string]models.Product
	mutex sync.Mutex
}

// NewMemoryStore es mi "constructor" para crear una nueva instancia del almacén en memoria.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]models.Product),
	}
}

// Create implementa la creación de un solo producto.
func (s *MemoryStore) Create(product models.Product) (models.Product, error) {
	s.mutex.Lock()         // Bloqueo el mapa para un acceso seguro.
	defer s.mutex.Unlock() // Me aseguro de que se libere al final de la función.

	product.ID = uuid.NewString() // Asigno un ID único.
	s.data[product.ID] = product
	return product, nil
}

// CreateBatch implementa la creación de múltiples productos a la vez.
func (s *MemoryStore) CreateBatch(products []models.Product) ([]models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	createdProducts := make([]models.Product, 0)
	// Recorro la lista de productos que recibí.
	for _, product := range products {
		product.ID = uuid.NewString() // A cada uno le asigno un ID.
		s.data[product.ID] = product
		createdProducts = append(createdProducts, product)
	}
	return createdProducts, nil
}

// Get implementa la obtención de todos los productos.
func (s *MemoryStore) Get() ([]models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Convierto mi mapa de productos en una lista para devolverla.
	list := make([]models.Product, 0, len(s.data))
	for _, prod := range s.data {
		list = append(list, prod)
	}
	return list, nil
}

// GetByID implementa la búsqueda de un producto por su ID.
func (s *MemoryStore) GetByID(id string) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Uso el "comma ok idiom" para verificar si el producto existe.
	product, ok := s.data[id]
	if !ok {
		// Si no existe, devuelvo un error claro.
		return models.Product{}, fmt.Errorf("producto con ID %s no encontrado", id)
	}
	return product, nil
}

// Update implementa la actualización de un producto existente.
func (s *MemoryStore) Update(id string, product models.Product) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Primero, me aseguro de que el producto que quiero actualizar exista.
	if _, ok := s.data[id]; !ok {
		return models.Product{}, fmt.Errorf("producto con ID %s no encontrado para actualizar", id)
	}

	product.ID = id // Me aseguro de que el producto conserve su ID original.
	s.data[id] = product
	return product, nil
}

// Delete implementa la eliminación de un producto.
func (s *MemoryStore) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Verifico que el producto exista antes de intentar borrarlo.
	if _, ok := s.data[id]; !ok {
		return fmt.Errorf("producto con ID %s no encontrado para eliminar", id)
	}

	// Uso la función delete nativa de Go para quitar un elemento de un mapa.
	delete(s.data, id)
	return nil
}

// Package repository handles database operations.
// It provides an abstraction layer between business logic and persistence.
package repository

import "github.com/mennyaboush/restaurant-inventory-ai/internal/models"

// ProductRepository defines operations for managing products.
type ProductRepository interface {
	AddProduct(p *models.Product) (string, error)
	GetProduct(id string) (*models.Product, error)
	ListProducts() []*models.Product
	SearchProducts(query string) []*models.Product
	UpdateProduct(p *models.Product) error
	DeleteProduct(id string) error
}

// StockRepository defines operations for reading and configuring stock.
// UpdateStock is kept for existing callers; new business flows should prefer
// ApplyStockMovement so the stock change and audit record stay together.
type StockRepository interface {
	GetStock(productID string) (*models.Stock, error)
	UpdateStock(productID string, boxes, units int) error
	SetMinStock(productID string, minStock int) error
	GetLowStockProducts() []*models.Product
}

// MovementRepository owns auditable stock mutations.
type MovementRepository interface {
	// ApplyStockMovement atomically changes stock and persists the movement.
	ApplyStockMovement(m *models.StockMovement) error
	ListStockMovements(productID string, limit int) ([]*models.StockMovement, error)
}

// Repository combines the inventory persistence contracts.
type Repository interface {
	ProductRepository
	StockRepository
	MovementRepository
}

var _ Repository = (*MemoryStore)(nil)

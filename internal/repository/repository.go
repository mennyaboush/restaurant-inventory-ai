// Package repository handles database operations.
// It provides an abstraction layer between the business logic
// and the database, making it easier to test and swap databases.
package repository

import "github.com/mennyaboush/restaurant-inventory-ai/internal/models"

// ============================================
// REPOSITORY INTERFACE
// ============================================
// This interface defines WHAT operations we need.
// Different implementations (Memory, PostgreSQL) define HOW.
//
// Benefits:
// 1. Swap storage without changing business logic
// 2. Easy to mock for testing
// 3. Clear contract for what storage must provide

// ProductRepository defines operations for managing products
type ProductRepository interface {
	// AddProduct creates a new product, returns generated ID
	AddProduct(p *models.Product) (string, error)

	// GetProduct retrieves a product by ID
	GetProduct(id string) (*models.Product, error)

	// ListProducts returns all active products
	ListProducts() []*models.Product

	// SearchProducts finds products matching query (name/brand)
	SearchProducts(query string) []*models.Product

	// UpdateProduct updates an existing product
	UpdateProduct(p *models.Product) error

	// DeleteProduct soft-deletes a product
	DeleteProduct(id string) error
}

// StockRepository defines operations for managing stock
type StockRepository interface {
	// GetStock retrieves stock for a product
	GetStock(productID string) (*models.Stock, error)

	// UpdateStock adds/removes stock (negative for removal)
	UpdateStock(productID string, boxes, units int) error

	// SetMinStock sets the minimum stock alert threshold
	SetMinStock(productID string, minStock int) error

	// GetLowStockProducts returns products below minimum
	GetLowStockProducts() []*models.Product
}

// Repository combines all repository interfaces
// This is what most code will use
type Repository interface {
	ProductRepository
	StockRepository
}

// ============================================
// VERIFICATION: MemoryStore implements Repository
// ============================================
// This line causes a compile error if MemoryStore
// doesn't implement all Repository methods.
// It's a Go pattern to catch mistakes early!

var _ Repository = (*MemoryStore)(nil)

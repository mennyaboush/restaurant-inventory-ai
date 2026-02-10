// Package repository handles data storage and retrieval
package repository

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
)

// ============================================
// ERRORS
// ============================================

var (
	ErrProductNotFound   = fmt.Errorf("product not found")
	ErrProductExists     = fmt.Errorf("product already exists")
	ErrStockNotFound     = fmt.Errorf("stock not found")
	ErrInsufficientStock = fmt.Errorf("insufficient stock")
)

// ============================================
// MEMORY STORE
// ============================================

// MemoryStore is an in-memory implementation of product storage
// Data is lost when the program stops - this is for learning!
// Later we'll swap this for PostgreSQL with the same interface
type MemoryStore struct {
	// Maps for O(1) lookup by ID
	products map[string]*models.Product // productID → Product
	stock    map[string]*models.Stock   // productID → Stock

	// Counter for generating IDs
	nextID int

	// Mutex for thread safety (multiple goroutines accessing store)
	// We'll learn about this more in concurrency lessons
	mu sync.RWMutex
}

// NewMemoryStore creates a new empty store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		products: make(map[string]*models.Product),
		stock:    make(map[string]*models.Stock),
		nextID:   1,
	}
}

// ============================================
// PRODUCT OPERATIONS
// ============================================

// AddProduct adds a new product to the store
// Returns the generated ID or error if validation fails
func (s *MemoryStore) AddProduct(p *models.Product) (string, error) {
	// Validate first
	if err := p.Validate(); err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate ID
	id := fmt.Sprintf("PROD-%03d", s.nextID)
	s.nextID++

	// Set ID and ensure active
	p.ID = id
	p.IsActive = true

	// Store product
	s.products[id] = p

	// Initialize stock at zero
	s.stock[id] = &models.Stock{
		ProductID:     id,
		QuantityBoxes: 0,
		QuantityUnits: 0,
		MinStock:      0,
		LastUpdated:   time.Now(),
	}

	return id, nil
}

// GetProduct retrieves a product by ID
// Returns error if not found
func (s *MemoryStore) GetProduct(id string) (*models.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	product, exists := s.products[id]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrProductNotFound, id)
	}

	return product, nil
}

// ListProducts returns all active products
func (s *MemoryStore) ListProducts() []*models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Create slice with capacity = number of products (efficient!)
	result := make([]*models.Product, 0, len(s.products))

	for _, p := range s.products {
		if p.IsActive {
			result = append(result, p)
		}
	}

	return result
}

// SearchProducts finds products where name contains the query
// Case-insensitive search
// Returns empty slice if no matches (not error!)
func (s *MemoryStore) SearchProducts(query string) []*models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	query = strings.ToLower(query)
	var results []*models.Product

	for _, p := range s.products {
		if !p.IsActive {
			continue
		}

		// Search in Name, Brand, or Category
		nameLower := strings.ToLower(p.Name)
		brandLower := strings.ToLower(p.Brand)

		if strings.Contains(nameLower, query) || strings.Contains(brandLower, query) {
			results = append(results, p)
		}
	}

	return results
}

// UpdateProduct updates an existing product
func (s *MemoryStore) UpdateProduct(p *models.Product) error {
	if err := p.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.products[p.ID]; !exists {
		return fmt.Errorf("%w: %s", ErrProductNotFound, p.ID)
	}

	s.products[p.ID] = p
	return nil
}

// DeleteProduct soft-deletes a product (sets IsActive = false)
// We don't really delete to preserve history
func (s *MemoryStore) DeleteProduct(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, exists := s.products[id]
	if !exists {
		return fmt.Errorf("%w: %s", ErrProductNotFound, id)
	}

	product.IsActive = false
	return nil
}

// ============================================
// STOCK OPERATIONS
// ============================================

// GetStock retrieves stock for a product
func (s *MemoryStore) GetStock(productID string) (*models.Stock, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stock, exists := s.stock[productID]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrStockNotFound, productID)
	}

	return stock, nil
}

// UpdateStock adds or removes stock (use negative for removal)
// Returns error if resulting stock would be negative
func (s *MemoryStore) UpdateStock(productID string, boxes, units int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stock, exists := s.stock[productID]
	if !exists {
		return fmt.Errorf("%w: %s", ErrStockNotFound, productID)
	}

	// Calculate new values
	newBoxes := stock.QuantityBoxes + boxes
	newUnits := stock.QuantityUnits + units

	// Check for negative stock
	if newBoxes < 0 || newUnits < 0 {
		return fmt.Errorf("%w: would result in %d boxes, %d units",
			ErrInsufficientStock, newBoxes, newUnits)
	}

	// Update
	stock.QuantityBoxes = newBoxes
	stock.QuantityUnits = newUnits
	stock.LastUpdated = time.Now()

	return nil
}

// SetMinStock sets the minimum stock alert threshold
func (s *MemoryStore) SetMinStock(productID string, minStock int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stock, exists := s.stock[productID]
	if !exists {
		return fmt.Errorf("%w: %s", ErrStockNotFound, productID)
	}

	stock.MinStock = minStock
	return nil
}

// GetLowStockProducts returns products below their minimum stock level
func (s *MemoryStore) GetLowStockProducts() []*models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var lowStock []*models.Product

	for id, stock := range s.stock {
		product := s.products[id]
		if product == nil || !product.IsActive {
			continue
		}

		// Check if below minimum
		if stock.IsLowStock(product.BoxSize) {
			lowStock = append(lowStock, product)
		}
	}

	return lowStock
}

// ============================================
// UTILITY METHODS
// ============================================

// Count returns the number of active products
func (s *MemoryStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	for _, p := range s.products {
		if p.IsActive {
			count++
		}
	}
	return count
}

// Clear removes all data (useful for testing)
func (s *MemoryStore) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.products = make(map[string]*models.Product)
	s.stock = make(map[string]*models.Stock)
	s.nextID = 1
}

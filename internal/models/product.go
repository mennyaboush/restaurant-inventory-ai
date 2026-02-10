// Package models contains the data structures for our inventory system
package models

import (
	"errors"
	"fmt"
	"time"
)

// ============================================
// ERRORS: Define specific errors for the domain
// This is the Go way - errors are values, not exceptions
// ============================================

var (
	// Product errors
	ErrProductNameRequired    = errors.New("product name is required")
	ErrProductInvalidSize     = errors.New("product size must be positive")
	ErrProductInvalidPrice    = errors.New("product price must be positive")
	ErrProductInvalidCategory = errors.New("invalid product category")

	// Stock errors
	ErrStockNegative        = errors.New("stock cannot be negative")
	ErrStockProductRequired = errors.New("product ID is required for stock")

	// Movement errors
	ErrMovementInvalidType = errors.New("invalid movement type")
	ErrMovementNoQuantity  = errors.New("movement must have boxes or units")
	ErrMovementNoPerformer = errors.New("performed_by is required")
)

// Product represents an item in our inventory
// Identified uniquely by: Brand + Size + ContainerType
type Product struct {
	ID            string  // Unique identifier (e.g., "PROD-001")
	Name          string  // Hebrew name: "קוקה קולה 330 מ״ל פחית"
	Brand         string  // Brand name (usually English): "Coca Cola"
	Size          int     // Size in ml or grams
	ContainerType string  // "can", "bottle", "bag", "piece"
	BoxSize       int     // Units per box (0 if sold individually)
	Price         float64 // Price per unit in NIS
	Category      string  // "drinks", "vegetables", "dairy"
	IsActive      bool    // Is product still sold?
}

// Stock tracks inventory levels for a product
type Stock struct {
	ProductID     string    // Links to Product.ID
	QuantityBoxes int       // Full boxes in stock
	QuantityUnits int       // Loose units (not in boxes)
	MinStock      int       // Alert threshold
	LastUpdated   time.Time // Last modification time
}

// TotalUnits calculates total units from boxes and loose units
func (s *Stock) TotalUnits(boxSize int) int {
	return (s.QuantityBoxes * boxSize) + s.QuantityUnits
}

// IsLowStock checks if stock is below minimum threshold
func (s *Stock) IsLowStock(boxSize int) bool {
	return s.TotalUnits(boxSize) < s.MinStock
}

// StockMovement logs every inventory change
// IMPORTANT: Always track WHO did the action and WHO logged it
type StockMovement struct {
	ID          string    // Unique identifier
	ProductID   string    // Which product
	Type        string    // "IN", "OUT", "WASTE", "ADJUSTMENT"
	Boxes       int       // Number of boxes changed
	Units       int       // Number of loose units changed
	PerformedBy string    // WHO actually did the physical action
	ReportedBy  string    // WHO logged it in the system
	Reason      string    // Why: "delivery", "sold", "expired"
	CreatedAt   time.Time // When this was logged
}

// Movement types as constants
const (
	MovementIn         = "IN"         // Stock received
	MovementOut        = "OUT"        // Stock sold/used
	MovementWaste      = "WASTE"      // Stock thrown away
	MovementAdjustment = "ADJUSTMENT" // Inventory correction
)

// Categories in Hebrew and English
var Categories = map[string]string{
	"drinks":     "משקאות",
	"vegetables": "ירקות",
	"dairy":      "מוצרי חלב",
	"meat":       "בשר",
	"dry_goods":  "מוצרים יבשים",
	"sauces":     "רטבים",
	"canned":     "שימורים",
}

// ============================================
// VALIDATION FUNCTIONS
// Go pattern: return (result, error) - always check the error!
// ============================================

// Validate checks if a Product has all required fields
// Returns nil if valid, or an error describing what's wrong
func (p *Product) Validate() error {
	// Check required string fields
	if p.Name == "" {
		return ErrProductNameRequired
	}

	// Check numeric fields
	if p.Size <= 0 {
		return ErrProductInvalidSize
	}
	if p.Price < 0 {
		return ErrProductInvalidPrice
	}

	// Check category is valid
	if p.Category != "" {
		if _, exists := Categories[p.Category]; !exists {
			// Wrap error with context using fmt.Errorf and %w
			return fmt.Errorf("%w: %s", ErrProductInvalidCategory, p.Category)
		}
	}

	return nil // nil means no error = success!
}

// Validate checks if Stock values are valid
func (s *Stock) Validate() error {
	if s.ProductID == "" {
		return ErrStockProductRequired
	}
	if s.QuantityBoxes < 0 || s.QuantityUnits < 0 {
		return ErrStockNegative
	}
	return nil
}

// ValidMovementTypes lists all allowed movement types
var ValidMovementTypes = map[string]bool{
	MovementIn:         true,
	MovementOut:        true,
	MovementWaste:      true,
	MovementAdjustment: true,
}

// Validate checks if a StockMovement is valid
func (m *StockMovement) Validate() error {
	// Check movement type
	if !ValidMovementTypes[m.Type] {
		return fmt.Errorf("%w: %s", ErrMovementInvalidType, m.Type)
	}

	// Must have some quantity
	if m.Boxes == 0 && m.Units == 0 {
		return ErrMovementNoQuantity
	}

	// Must know who did it
	if m.PerformedBy == "" {
		return ErrMovementNoPerformer
	}

	return nil
}

// ============================================
// HELPER FUNCTIONS
// ============================================

// NewProduct creates a new Product with validation
// This is a "constructor" pattern in Go
func NewProduct(name, brand string, size int, containerType string, boxSize int, price float64, category string) (*Product, error) {
	p := &Product{
		ID:            "", // Will be set by database
		Name:          name,
		Brand:         brand,
		Size:          size,
		ContainerType: containerType,
		BoxSize:       boxSize,
		Price:         price,
		Category:      category,
		IsActive:      true,
	}

	// Validate before returning
	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("invalid product: %w", err)
	}

	return p, nil
}

// NewStockMovement creates a validated stock movement
// If reportedBy is empty, it defaults to performedBy (self-reported)
func NewStockMovement(productID, movementType string, boxes, units int, performedBy, reportedBy, reason string) (*StockMovement, error) {
	// Default: if no reporter specified, person reporting themselves
	if reportedBy == "" {
		reportedBy = performedBy
	}

	m := &StockMovement{
		ID:          "", // Will be set by database
		ProductID:   productID,
		Type:        movementType,
		Boxes:       boxes,
		Units:       units,
		PerformedBy: performedBy,
		ReportedBy:  reportedBy,
		Reason:      reason,
		CreatedAt:   time.Now(),
	}

	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("invalid movement: %w", err)
	}

	return m, nil
}

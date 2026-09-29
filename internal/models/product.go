// Package models contains the data structures for our inventory system
package models

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrProductNameRequired    = errors.New("product name is required")
	ErrProductInvalidSize     = errors.New("product size must be positive")
	ErrProductInvalidPrice    = errors.New("product price must be positive")
	ErrProductInvalidCategory = errors.New("invalid product category")

	ErrStockNegative        = errors.New("stock cannot be negative")
	ErrStockProductRequired = errors.New("product ID is required for stock")

	ErrMovementProductRequired = errors.New("product ID is required for movement")
	ErrMovementInvalidType     = errors.New("invalid movement type")
	ErrMovementNoQuantity      = errors.New("movement must have boxes or units")
	ErrMovementNoPerformer     = errors.New("performed_by is required")
)

type Product struct {
	ID string
	Name string
	Brand string
	Size int
	ContainerType string
	BoxSize int
	Price float64
	Category string
	IsActive bool
}

type Stock struct {
	ProductID string
	QuantityBoxes int
	QuantityUnits int
	MinStock int
	LastUpdated time.Time
}

func (s *Stock) TotalUnits(boxSize int) int { return (s.QuantityBoxes * boxSize) + s.QuantityUnits }
func (s *Stock) IsLowStock(boxSize int) bool { return s.TotalUnits(boxSize) < s.MinStock }

type StockMovement struct {
	ID string
	ProductID string
	Type string
	Boxes int
	Units int
	PerformedBy string
	ReportedBy string
	Reason string
	CreatedAt time.Time
}

const (
	MovementIn = "IN"
	MovementOut = "OUT"
	MovementWaste = "WASTE"
	MovementAdjustment = "ADJUSTMENT"
)

var Categories = map[string]string{
	"drinks":"משקאות","vegetables":"ירקות","dairy":"מוצרי חלב","meat":"בשר",
	"dry_goods":"מוצרים יבשים","sauces":"רטבים","canned":"שימורים",
}

func (p *Product) Validate() error {
	if p.Name == "" { return ErrProductNameRequired }
	if p.Size <= 0 { return ErrProductInvalidSize }
	if p.Price < 0 { return ErrProductInvalidPrice }
	if p.Category != "" {
		if _, exists := Categories[p.Category]; !exists { return fmt.Errorf("%w: %s", ErrProductInvalidCategory, p.Category) }
	}
	return nil
}

func (s *Stock) Validate() error {
	if s.ProductID == "" { return ErrStockProductRequired }
	if s.QuantityBoxes < 0 || s.QuantityUnits < 0 { return ErrStockNegative }
	return nil
}

var ValidMovementTypes = map[string]bool{
	MovementIn:true, MovementOut:true, MovementWaste:true, MovementAdjustment:true,
}

func (m *StockMovement) Validate() error {
	if m.ProductID == "" { return ErrMovementProductRequired }
	if !ValidMovementTypes[m.Type] { return fmt.Errorf("%w: %s", ErrMovementInvalidType, m.Type) }
	if m.Boxes == 0 && m.Units == 0 { return ErrMovementNoQuantity }
	if m.PerformedBy == "" { return ErrMovementNoPerformer }
	return nil
}

func NewProduct(name, brand string, size int, containerType string, boxSize int, price float64, category string) (*Product, error) {
	p := &Product{Name:name, Brand:brand, Size:size, ContainerType:containerType, BoxSize:boxSize, Price:price, Category:category, IsActive:true}
	if err := p.Validate(); err != nil { return nil, fmt.Errorf("invalid product: %w", err) }
	return p, nil
}

func NewStockMovement(productID, movementType string, boxes, units int, performedBy, reportedBy, reason string) (*StockMovement, error) {
	if reportedBy == "" { reportedBy = performedBy }
	m := &StockMovement{ProductID:productID, Type:movementType, Boxes:boxes, Units:units, PerformedBy:performedBy, ReportedBy:reportedBy, Reason:reason, CreatedAt:time.Now()}
	if err := m.Validate(); err != nil { return nil, fmt.Errorf("invalid movement: %w", err) }
	return m, nil
}

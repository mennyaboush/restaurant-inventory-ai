// Package service contains inventory business logic.
package service

import (
	"fmt"

	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/repository"
)

// InventoryService exposes semantic inventory operations.
// HTTP handlers and future AI tools should call this layer rather than
// manipulating stock directly.
type InventoryService struct {
	repo repository.Repository
}

func NewInventoryService(repo repository.Repository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) ReceiveStock(productID string, boxes, units int, performedBy, reportedBy, reason string) (*models.Stock, error) {
	return s.apply(productID, models.MovementIn, boxes, units, performedBy, reportedBy, reason)
}

func (s *InventoryService) ConsumeStock(productID string, boxes, units int, performedBy, reportedBy, reason string) (*models.Stock, error) {
	if boxes < 0 || units < 0 {
		return nil, fmt.Errorf("consume quantities must be non-negative")
	}
	return s.apply(productID, models.MovementOut, -boxes, -units, performedBy, reportedBy, reason)
}

func (s *InventoryService) RecordWaste(productID string, boxes, units int, performedBy, reportedBy, reason string) (*models.Stock, error) {
	if boxes < 0 || units < 0 {
		return nil, fmt.Errorf("waste quantities must be non-negative")
	}
	return s.apply(productID, models.MovementWaste, -boxes, -units, performedBy, reportedBy, reason)
}

// AdjustStock accepts signed deltas because an inventory count may correct
// stock either upward or downward.
func (s *InventoryService) AdjustStock(productID string, boxes, units int, performedBy, reportedBy, reason string) (*models.Stock, error) {
	return s.apply(productID, models.MovementAdjustment, boxes, units, performedBy, reportedBy, reason)
}

func (s *InventoryService) GetMovements(productID string, limit int) ([]*models.StockMovement, error) {
	return s.repo.ListStockMovements(productID, limit)
}

func (s *InventoryService) apply(productID, movementType string, boxes, units int, performedBy, reportedBy, reason string) (*models.Stock, error) {
	if productID == "" {
		return nil, models.ErrStockProductRequired
	}
	if movementType == models.MovementIn && (boxes < 0 || units < 0) {
		return nil, fmt.Errorf("received quantities must be non-negative")
	}

	movement, err := models.NewStockMovement(productID, movementType, boxes, units, performedBy, reportedBy, reason)
	if err != nil {
		return nil, err
	}
	if err := s.repo.ApplyStockMovement(movement); err != nil {
		return nil, err
	}
	return s.repo.GetStock(productID)
}

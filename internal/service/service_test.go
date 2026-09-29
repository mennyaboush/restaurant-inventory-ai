package service

import (
	"errors"
	"testing"

	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/repository"
)

func newTestInventory(t *testing.T) (*InventoryService, *repository.MemoryStore, string) {
	t.Helper()
	store := repository.NewMemoryStore()
	id, err := store.AddProduct(&models.Product{
		Name: "Coke 330", Brand: "Coca Cola", Size: 330,
		ContainerType: "can", BoxSize: 24, Price: 5.5,
		Category: "drinks", IsActive: true,
	})
	if err != nil {
		t.Fatalf("add product: %v", err)
	}
	return NewInventoryService(store), store, id
}

func TestInventoryServiceReceiveAndConsume(t *testing.T) {
	svc, store, productID := newTestInventory(t)

	stock, err := svc.ReceiveStock(productID, 2, 5, "Menny", "", "delivery")
	if err != nil {
		t.Fatalf("receive stock: %v", err)
	}
	if stock.QuantityBoxes != 2 || stock.QuantityUnits != 5 {
		t.Fatalf("unexpected stock after receive: %+v", stock)
	}

	stock, err = svc.ConsumeStock(productID, 1, 2, "Worker", "Menny", "used in shift")
	if err != nil {
		t.Fatalf("consume stock: %v", err)
	}
	if stock.QuantityBoxes != 1 || stock.QuantityUnits != 3 {
		t.Fatalf("unexpected stock after consume: %+v", stock)
	}

	movements, err := store.ListStockMovements(productID, 10)
	if err != nil {
		t.Fatalf("list movements: %v", err)
	}
	if len(movements) != 2 {
		t.Fatalf("expected 2 movements, got %d", len(movements))
	}
	if movements[0].Type != models.MovementOut || movements[0].Boxes != -1 || movements[0].Units != -2 {
		t.Fatalf("unexpected consume movement: %+v", movements[0])
	}
	if movements[1].ReportedBy != "Menny" {
		t.Fatalf("expected self-reported receive movement, got %+v", movements[1])
	}
}

func TestInventoryServiceRejectsInsufficientStockWithoutMovement(t *testing.T) {
	svc, store, productID := newTestInventory(t)

	if _, err := svc.ReceiveStock(productID, 0, 2, "Menny", "", "initial"); err != nil {
		t.Fatalf("receive stock: %v", err)
	}

	_, err := svc.RecordWaste(productID, 0, 3, "Worker", "", "expired")
	if !errors.Is(err, repository.ErrInsufficientStock) {
		t.Fatalf("expected ErrInsufficientStock, got %v", err)
	}

	stock, err := store.GetStock(productID)
	if err != nil {
		t.Fatalf("get stock: %v", err)
	}
	if stock.QuantityUnits != 2 {
		t.Fatalf("stock changed after rejected movement: %+v", stock)
	}

	movements, err := store.ListStockMovements(productID, 10)
	if err != nil {
		t.Fatalf("list movements: %v", err)
	}
	if len(movements) != 1 {
		t.Fatalf("rejected movement should not be persisted; got %d movements", len(movements))
	}
}

func TestInventoryServiceAdjustmentCanIncreaseOrDecrease(t *testing.T) {
	svc, _, productID := newTestInventory(t)

	if _, err := svc.ReceiveStock(productID, 1, 5, "Menny", "", "initial"); err != nil {
		t.Fatalf("receive stock: %v", err)
	}

	stock, err := svc.AdjustStock(productID, 0, -2, "Menny", "", "inventory count")
	if err != nil {
		t.Fatalf("adjust stock: %v", err)
	}
	if stock.QuantityBoxes != 1 || stock.QuantityUnits != 3 {
		t.Fatalf("unexpected adjusted stock: %+v", stock)
	}
}

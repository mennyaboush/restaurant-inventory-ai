package repostest

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
)

// Store defines the subset of repository functionality used by the shared tests.
// Defined locally to avoid import cycles between the test helper package and the
// main repository package.
type Store interface {
	AddProduct(*models.Product) (string, error)
	GetProduct(string) (*models.Product, error)
	SearchProducts(string) []*models.Product
	ListProducts() []*models.Product
	UpdateProduct(*models.Product) error
	DeleteProduct(string) error
	GetStock(string) (*models.Stock, error)
	UpdateStock(string, int, int) error
	SetMinStock(string, int) error
	GetLowStockProducts() []*models.Product
}

// RunStoreIntegrationTests runs the common integration tests against any
// Repository implementation supplied by newStore(db).
func RunStoreIntegrationTests(t *testing.T, newStore func(*sql.DB) Store, db *sql.DB) {
	store := newStore(db)

	// Clean prefix for test products
	tsPrefix := fmt.Sprintf("itest-%d", time.Now().UnixNano())

	// 1) Validation edge case: missing name
	bad := &models.Product{Name: "", Brand: "B", Size: 100, ContainerType: "box", BoxSize: 10, Price: 1.0, Category: "drinks"}
	if _, err := store.AddProduct(bad); err == nil {
		t.Fatalf("expected validation error when adding product without name")
	}

	// 2) Happy path: add product
	p := &models.Product{
		Name:          "ITEST Product",
		Brand:         tsPrefix + "-brand",
		Size:          100,
		ContainerType: "box",
		BoxSize:       10,
		Price:         2.5,
		Category:      "drinks",
		IsActive:      true,
	}
	id, err := store.AddProduct(p)
	if err != nil {
		t.Fatalf("AddProduct failed: %v", err)
	}

	// 3) GetProduct
	got, err := store.GetProduct(id)
	if err != nil {
		t.Fatalf("GetProduct failed: %v", err)
	}
	if got.Brand != p.Brand || got.Name != p.Name {
		t.Fatalf("mismatched product retrieved: got=%+v want=%+v", got, p)
	}

	// 4) SearchProducts
	res := store.SearchProducts("ITEST")
	found := false
	for _, r := range res {
		if r.ID == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("SearchProducts did not find inserted product")
	}

	// 5) ListProducts contains it (active)
	all := store.ListProducts()
	found = false
	for _, r := range all {
		if r.ID == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("ListProducts did not contain inserted product")
	}

	// 6) UpdateProduct
	got.Name = "ITEST Product Updated"
	got.Price = 3.14
	if err := store.UpdateProduct(got); err != nil {
		t.Fatalf("UpdateProduct failed: %v", err)
	}
	got2, err := store.GetProduct(id)
	if err != nil {
		t.Fatalf("GetProduct after update failed: %v", err)
	}
	if got2.Price != 3.14 || got2.Name != "ITEST Product Updated" {
		t.Fatalf("Update did not persist: got=%+v", got2)
	}

	// 7) Stock: initial stock row should exist
	st, err := store.GetStock(id)
	if err != nil {
		t.Fatalf("GetStock failed: %v", err)
	}
	if st.QuantityBoxes != 0 || st.QuantityUnits != 0 {
		t.Fatalf("expected zero initial stock, got %+v", st)
	}

	// 8) UpdateStock (add)
	if err := store.UpdateStock(id, 2, 5); err != nil {
		t.Fatalf("UpdateStock add failed: %v", err)
	}
	st2, err := store.GetStock(id)
	if err != nil {
		t.Fatalf("GetStock after add failed: %v", err)
	}
	if st2.QuantityBoxes != 2 || st2.QuantityUnits != 5 {
		t.Fatalf("unexpected stock after add: %+v", st2)
	}

	// 9) UpdateStock (remove too much) -> expect error and no change
	if err := store.UpdateStock(id, -5, 0); err == nil {
		t.Fatalf("expected error when removing more boxes than available")
	}
	st3, err := store.GetStock(id)
	if err != nil {
		t.Fatalf("GetStock after failed remove failed: %v", err)
	}
	if st3.QuantityBoxes != 2 || st3.QuantityUnits != 5 {
		t.Fatalf("stock changed despite failed remove: %+v", st3)
	}

	// 10) SetMinStock and GetLowStockProducts
	if err := store.SetMinStock(id, 1000); err != nil {
		t.Fatalf("SetMinStock failed: %v", err)
	}
	lows := store.GetLowStockProducts()
	found = false
	for _, r := range lows {
		if r.ID == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("GetLowStockProducts did not include product expected to be low")
	}

	// 11) DeleteProduct (soft-delete) and verify IsActive=false
	if err := store.DeleteProduct(id); err != nil {
		t.Fatalf("DeleteProduct failed: %v", err)
	}
	afterDel, err := store.GetProduct(id)
	if err != nil {
		t.Fatalf("GetProduct after delete failed: %v", err)
	}
	if afterDel.IsActive {
		t.Fatalf("expected product to be inactive after DeleteProduct")
	}

	// 12) Search for non-existent
	empt := store.SearchProducts("no-such-product-xyz")
	if len(empt) != 0 {
		t.Fatalf("expected empty search results for non-matching query, got %d", len(empt))
	}
}

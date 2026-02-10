package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mennyaboush/restaurant-inventory-ai/config"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/api"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/repository"
)

func main() {
	fmt.Println("🍕 Restaurant Inventory AI")
	fmt.Println("==========================")

	// Start HTTP server (Week 2)
	startHTTPServer()
}

// startHTTPServer starts a basic HTTP server with GET /products
func startHTTPServer() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	fmt.Println("🔌 Connecting to database...")
	db, err := cfg.ConnectDB()
	if err != nil {
		fmt.Printf("❌ Failed to connect to database: %v\n", err)
		fmt.Println("Make sure PostgreSQL is running with: docker-compose up -d")
		return
	}
	defer db.Close()
	fmt.Println("✅ Database connected successfully")

	// Create PostgresStore (implements Repository)
	store := repository.NewPostgresStore(db)

	// Add demo products to store (for now)
	demoProducts := []struct {
		name      string
		brand     string
		size      int
		container string
		boxSize   int
		price     float64
		category  string
	}{
		{"קוקה קולה 330 מ״ל פחית", "Coca Cola", 330, "can", 24, 5.50, "drinks"},
		{"פנטה 330 מ״ל פחית", "Fanta", 330, "can", 24, 5.50, "drinks"},
		{"פלפל אדום", "ירקות טריים", 1000, "kg", 0, 15.00, "vegetables"},
		{"פלפל ירוק", "ירקות טריים", 1000, "kg", 0, 12.00, "vegetables"},
		{"חומוס 400 גרם", "עשי", 400, "can", 12, 8.00, "canned"},
	}
	for _, p := range demoProducts {
		product := &models.Product{
			Name:          p.name,
			Brand:         p.brand,
			Size:          p.size,
			ContainerType: p.container,
			BoxSize:       p.boxSize,
			Price:         p.price,
			Category:      p.category,
			IsActive:      true,
		}
		_, _ = store.AddProduct(product)
	}

	// Create API and use chi router
	apiHandler := api.NewAPI(store)
	router := apiHandler.Router()

	fmt.Println("🚀 HTTP server running at http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// learnStructs demonstrates how to use our domain models
func learnStructs() {
	fmt.Println("\n🎓 Lesson 1.2: Structs")
	fmt.Println("======================")

	// ============================================
	// SECTION 1: Creating a Product
	// ============================================
	fmt.Println("\n📖 SECTION 1: Creating a Product")

	// Create a product using named fields (recommended)
	cola := models.Product{
		ID:            "PROD-001",
		Name:          "קוקה קולה 330 מ״ל פחית",
		Brand:         "Coca Cola",
		Size:          330,
		ContainerType: "can",
		BoxSize:       24,
		Price:         5.50,
		Category:      "drinks",
		IsActive:      true,
	}

	fmt.Printf("Product: %s (Brand: %s)\n", cola.Name, cola.Brand)
	fmt.Printf("Size: %dml, Box of %d, Price: %.2f NIS\n",
		cola.Size, cola.BoxSize, cola.Price)

	// ============================================
	// SECTION 2: Working with Stock
	// ============================================
	fmt.Println("\n📖 SECTION 2: Stock Tracking")

	stock := models.Stock{
		ProductID:     cola.ID,
		QuantityBoxes: 5,
		QuantityUnits: 12,
		MinStock:      50,
		LastUpdated:   time.Now(),
	}

	// Use the method we defined on Stock
	total := stock.TotalUnits(cola.BoxSize)
	fmt.Printf("Stock: %d boxes + %d loose = %d total units\n",
		stock.QuantityBoxes, stock.QuantityUnits, total)

	// Check if low stock using the method
	if stock.IsLowStock(cola.BoxSize) {
		fmt.Printf("⚠️ מלאי נמוך! Below %d units\n", stock.MinStock)
	} else {
		fmt.Printf("✅ Stock OK (min: %d)\n", stock.MinStock)
	}

	// ============================================
	// SECTION 3: Logging a Movement
	// ============================================
	fmt.Println("\n📖 SECTION 3: Stock Movement (Audit Trail)")

	// Scenario: אבא delivered 3 boxes, מני logged it
	movement := models.StockMovement{
		ID:          "MOV-001",
		ProductID:   cola.ID,
		Type:        models.MovementIn, // Using our constant!
		Boxes:       3,
		Units:       0,
		PerformedBy: "אבא",
		ReportedBy:  "מני",
		Reason:      "delivery from supplier",
		CreatedAt:   time.Now(),
	}

	fmt.Printf("Movement: %s +%d boxes\n", movement.Type, movement.Boxes)
	fmt.Printf("Done by: %s, Logged by: %s\n",
		movement.PerformedBy, movement.ReportedBy)

	// Update stock
	stock.QuantityBoxes += movement.Boxes
	newTotal := stock.TotalUnits(cola.BoxSize)
	fmt.Printf("Stock updated: %d → %d units\n", total, newTotal)

	// ============================================
	// SECTION 4: Pointers (why they matter)
	// ============================================
	fmt.Println("\n📖 SECTION 4: Pointers")

	// Without pointer - creates a COPY
	colaCopy := cola
	colaCopy.Price = 10.00
	fmt.Printf("Original: %.2f NIS, Copy: %.2f NIS\n", cola.Price, colaCopy.Price)

	// With pointer - modifies the ORIGINAL
	colaPtr := &cola
	colaPtr.Price = 6.00
	fmt.Printf("After pointer: %.2f NIS\n", cola.Price)

	fmt.Println("\n✅ Lesson 1.2 Complete!")
}

// learnErrorHandling demonstrates Go's error handling patterns
func learnErrorHandling() {
	fmt.Println("\n🎓 Lesson 1.5: Functions & Error Handling")
	fmt.Println("==========================================")

	// ============================================
	// SECTION 1: The Go Error Pattern
	// ============================================
	fmt.Println("\n📖 SECTION 1: Creating Products with Validation")

	// ✅ Valid product - using our constructor function
	cola, err := models.NewProduct(
		"קוקה קולה 330 מ״ל פחית", // name (Hebrew)
		"Coca Cola",              // brand
		330,                      // size
		"can",                    // containerType
		24,                       // boxSize
		5.50,                     // price
		"drinks",                 // category
	)

	// ALWAYS check the error!
	if err != nil {
		fmt.Printf("❌ Error creating cola: %v\n", err)
	} else {
		fmt.Printf("✅ Created: %s (Brand: %s)\n", cola.Name, cola.Brand)
	}

	// ❌ Invalid product - missing name
	_, err = models.NewProduct(
		"",      // name - EMPTY!
		"House", // brand
		1,       // size
		"piece", // containerType
		0,       // boxSize
		15.00,   // price
		"",      // category
	)

	if err != nil {
		fmt.Printf("❌ Expected error: %v\n", err)
	}

	// ❌ Invalid product - bad category
	_, err = models.NewProduct(
		"חומוס",            // name
		"House",            // brand
		200,                // size
		"container",        // containerType
		0,                  // boxSize
		12.00,              // price
		"invalid_category", // This category doesn't exist!
	)

	if err != nil {
		fmt.Printf("❌ Expected error: %v\n", err)
	}

	// ============================================
	// SECTION 2: Movement Validation
	// ============================================
	fmt.Println("\n📖 SECTION 2: Stock Movements with Validation")

	// ✅ Valid movement - אבא delivered, מני logged
	movement, err := models.NewStockMovement(
		"PROD-001",
		models.MovementIn,
		3,     // boxes
		0,     // units
		"אבא", // performedBy
		"מני", // reportedBy
		"delivery from supplier",
	)

	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Movement: %s +%d boxes by %s\n",
			movement.Type, movement.Boxes, movement.PerformedBy)
	}

	// ✅ Self-reported movement - reportedBy defaults to performedBy
	movement2, err := models.NewStockMovement(
		"PROD-001",
		models.MovementOut,
		0,
		5,
		"אבא",
		"", // Empty = self-reported
		"sold",
	)

	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Self-reported: %s by %s, logged by %s\n",
			movement2.Type, movement2.PerformedBy, movement2.ReportedBy)
	}

	// ❌ Invalid movement - no performer
	_, err = models.NewStockMovement(
		"PROD-001",
		models.MovementWaste,
		1,
		0,
		"", // WHO did it? Required!
		"",
		"expired",
	)

	if err != nil {
		fmt.Printf("❌ Expected error: %v\n", err)
	}

	// ❌ Invalid movement - bad type
	_, err = models.NewStockMovement(
		"PROD-001",
		"STOLEN", // Not a valid type!
		1,
		0,
		"אבא",
		"",
		"",
	)

	if err != nil {
		fmt.Printf("❌ Expected error: %v\n", err)
	}

	// ============================================
	// SECTION 3: Error Checking Patterns
	// ============================================
	fmt.Println("\n📖 SECTION 3: Error Checking Patterns")

	// Pattern 1: Check specific error type
	product := &models.Product{Name: "", Size: 100}
	err = product.Validate()

	if errors.Is(err, models.ErrProductNameRequired) {
		fmt.Println("✅ Caught: Product name is missing")
	}

	// Pattern 2: Check if error contains another error (wrapped)
	_, err = models.NewProduct("", "", 100, "", 0, 10, "invalid")
	if err != nil {
		fmt.Printf("Full error: %v\n", err)
		// The error is wrapped: "invalid product: product name is required"
	}

	// ============================================
	// KEY TAKEAWAYS
	// ============================================
	fmt.Println("\n" + "==================================================")
	fmt.Println("📝 KEY TAKEAWAYS:")
	fmt.Println("==================================================")
	fmt.Println("1. Go returns errors, not exceptions (no try/catch)")
	fmt.Println("2. ALWAYS check: if err != nil { handle it }")
	fmt.Println("3. Define domain errors: var ErrXxx = errors.New(...)")
	fmt.Println("4. Wrap errors for context: fmt.Errorf(\"context: %w\", err)")
	fmt.Println("5. Constructor functions validate before returning")
	fmt.Println("6. Use errors.Is() to check specific error types")

	fmt.Println("\n✅ Lesson 1.5 Complete!")
}

// learnInterfaces demonstrates the Repository interface pattern
func learnInterfaces() {
	fmt.Println("\n🎓 Lesson 1.11/1.12: Interfaces & Repository")
	fmt.Println("=============================================")

	// ============================================
	// SECTION 1: Create Store (implements Repository interface)
	// ============================================
	fmt.Println("\n📖 SECTION 1: Creating the Store")

	// This creates a MemoryStore, but we'll use it as Repository
	store := repository.NewMemoryStore()
	fmt.Println("✅ Created MemoryStore")

	// ============================================
	// SECTION 2: Add Products
	// ============================================
	fmt.Println("\n📖 SECTION 2: Adding Products")

	// Add some products
	products := []struct {
		name      string
		brand     string
		size      int
		container string
		boxSize   int
		price     float64
		category  string
	}{
		{"קוקה קולה 330 מ״ל פחית", "Coca Cola", 330, "can", 24, 5.50, "drinks"},
		{"פנטה 330 מ״ל פחית", "Fanta", 330, "can", 24, 5.50, "drinks"},
		{"פלפל אדום", "ירקות טריים", 1000, "kg", 0, 15.00, "vegetables"},
		{"פלפל ירוק", "ירקות טריים", 1000, "kg", 0, 12.00, "vegetables"},
		{"חומוס 400 גרם", "עשי", 400, "can", 12, 8.00, "canned"},
	}

	for _, p := range products {
		product := &models.Product{
			Name:          p.name,
			Brand:         p.brand,
			Size:          p.size,
			ContainerType: p.container,
			BoxSize:       p.boxSize,
			Price:         p.price,
			Category:      p.category,
		}

		id, err := store.AddProduct(product)
		if err != nil {
			fmt.Printf("❌ Error adding %s: %v\n", p.name, err)
		} else {
			fmt.Printf("✅ Added: %s (ID: %s)\n", p.name, id)
		}
	}

	// ============================================
	// SECTION 3: Search Products (your dad's use case!)
	// ============================================
	fmt.Println("\n📖 SECTION 3: Searching Products")

	// Search for "פלפל" - should find 2 products
	query := "פלפל"
	results := store.SearchProducts(query)
	fmt.Printf("Search '%s' found %d products:\n", query, len(results))
	for _, p := range results {
		fmt.Printf("  - %s (ID: %s)\n", p.Name, p.ID)
	}

	// Search for "cola" - should find 1 product (brand search)
	query = "cola"
	results = store.SearchProducts(query)
	fmt.Printf("Search '%s' found %d products:\n", query, len(results))
	for _, p := range results {
		fmt.Printf("  - %s (Brand: %s)\n", p.Name, p.Brand)
	}

	// ============================================
	// SECTION 4: Update Stock
	// ============================================
	fmt.Println("\n📖 SECTION 4: Updating Stock")

	// Add stock to first product (cola)
	err := store.UpdateStock("PROD-001", 5, 0) // 5 boxes
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		stock, _ := store.GetStock("PROD-001")
		fmt.Printf("✅ Added 5 boxes to PROD-001. Total: %d boxes\n", stock.QuantityBoxes)
	}

	// Try to remove more than we have
	err = store.UpdateStock("PROD-001", -10, 0) // Remove 10 boxes (we only have 5!)
	if err != nil {
		fmt.Printf("✅ Correctly rejected: %v\n", err)
	}

	// ============================================
	// SECTION 5: The Power of Interfaces
	// ============================================
	fmt.Println("\n📖 SECTION 5: Why Interfaces Matter")

	// This function accepts ANY Repository implementation
	demonstrateInterface(store)

	// ============================================
	// KEY TAKEAWAYS
	// ============================================
	fmt.Println("\n==================================================")
	fmt.Println("📝 KEY TAKEAWAYS:")
	fmt.Println("==================================================")
	fmt.Println("1. Interface = contract (WHAT, not HOW)")
	fmt.Println("2. MemoryStore implements Repository interface")
	fmt.Println("3. Later: PostgresStore will ALSO implement Repository")
	fmt.Println("4. Code using Repository works with BOTH!")
	fmt.Println("5. Search finds products by name/brand for user queries")

	fmt.Println("\n✅ Lesson 1.11/1.12 Complete!")
}

// demonstrateInterface shows how interfaces allow flexible code
// This function works with ANY Repository implementation!
func demonstrateInterface(repo repository.Repository) {
	fmt.Println("\n  → This function accepts repository.Repository interface")
	fmt.Println("  → It doesn't know/care if it's MemoryStore or PostgresStore")

	// Use the interface methods
	products := repo.ListProducts()
	fmt.Printf("  → Found %d products in the repository\n", len(products))
}

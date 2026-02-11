package main

import (
	"fmt"
	"net/http"

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

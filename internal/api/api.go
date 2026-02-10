// Package api contains HTTP handlers for the REST API.
// Handlers receive HTTP requests, validate input, call services,
// and return HTTP responses.
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/repository"
)

// API holds dependencies for HTTP handlers (e.g., the repository)
type API struct {
	Store repository.ProductRepository
}

// NewAPI creates a new API instance with the given repository
func NewAPI(store repository.ProductRepository) *API {
	return &API{Store: store}
}

// LoggingMiddleware logs each HTTP request with method, path, and duration
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("→ %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("← %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}

// ErrorResponse represents a JSON error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// respondJSON writes a JSON response with the given status code
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError writes a JSON error response
func respondError(w http.ResponseWriter, status int, errType, message string) {
	respondJSON(w, status, ErrorResponse{Error: errType, Message: message})
}

// validateProductInput validates product input fields and returns error message if invalid
func validateProductInput(name string, size int, price float64) (string, bool) {
	if name == "" {
		return "name is required", false
	}
	if size <= 0 {
		return "size must be positive", false
	}
	if price < 0 {
		return "price cannot be negative", false
	}
	return "", true
}

// Router returns a chi.Router with all API routes registered
func (api *API) Router() chi.Router {
	r := chi.NewRouter()

	// Add middleware
	r.Use(LoggingMiddleware)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", api.handleListProducts)
		r.Post("/", api.handleCreateProduct)
		r.Get("/{id}", api.handleGetProduct)
		r.Put("/{id}", api.handleUpdateProduct)
		r.Delete("/{id}", api.handleDeleteProduct)
	})
	return r
}

// handleListProducts handles GET /products
func (api *API) handleListProducts(w http.ResponseWriter, r *http.Request) {
	products := api.Store.ListProducts()
	respondJSON(w, http.StatusOK, products)
}

// handleCreateProduct handles POST /products
func (api *API) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name          string  `json:"name"`
		Brand         string  `json:"brand"`
		Size          int     `json:"size"`
		ContainerType string  `json:"containerType"`
		BoxSize       int     `json:"boxSize"`
		Price         float64 `json:"price"`
		Category      string  `json:"category"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "invalid_json", "Request body must be valid JSON")
		return
	}
	// Validate input
	if msg, ok := validateProductInput(input.Name, input.Size, input.Price); !ok {
		respondError(w, http.StatusBadRequest, "validation_error", msg)
		return
	}
	id, err := api.Store.AddProduct(&models.Product{
		Name:          input.Name,
		Brand:         input.Brand,
		Size:          input.Size,
		ContainerType: input.ContainerType,
		BoxSize:       input.BoxSize,
		Price:         input.Price,
		Category:      input.Category,
	})
	if err != nil {
		respondError(w, http.StatusBadRequest, "create_error", err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{"id": id})
}

// handleGetProduct handles GET /products/{id}
func (api *API) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	product, err := api.Store.GetProduct(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "not_found", err.Error())
		return
	}
	respondJSON(w, http.StatusOK, product)
}

// handleUpdateProduct handles PUT /products/{id}
func (api *API) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var input struct {
		Name          string  `json:"name"`
		Brand         string  `json:"brand"`
		Size          int     `json:"size"`
		ContainerType string  `json:"containerType"`
		BoxSize       int     `json:"boxSize"`
		Price         float64 `json:"price"`
		Category      string  `json:"category"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "invalid_json", "Request body must be valid JSON")
		return
	}
	// Validate input
	if msg, ok := validateProductInput(input.Name, input.Size, input.Price); !ok {
		respondError(w, http.StatusBadRequest, "validation_error", msg)
		return
	}
	product := &models.Product{
		ID:            id,
		Name:          input.Name,
		Brand:         input.Brand,
		Size:          input.Size,
		ContainerType: input.ContainerType,
		BoxSize:       input.BoxSize,
		Price:         input.Price,
		Category:      input.Category,
	}
	if err := api.Store.UpdateProduct(product); err != nil {
		respondError(w, http.StatusBadRequest, "update_error", err.Error())
		return
	}
	respondJSON(w, http.StatusOK, product)
}

// handleDeleteProduct handles DELETE /products/{id}
func (api *API) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := api.Store.DeleteProduct(id); err != nil {
		respondError(w, http.StatusNotFound, "not_found", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

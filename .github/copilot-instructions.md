# GitHub Copilot Instructions

> These instructions help GitHub Copilot understand how to assist with this project.

## Project Context

This is a **Restaurant Inventory AI** system - a real production application for a Pizza & Falafel restaurant owned by the developer's father.

**Primary Goals:**
1. **Learning:** Teach Go, Kubernetes, Docker, AI/RAG, Linux
2. **Production:** Build a real, usable system

**Languages:** Hebrew (primary UI), English (code)

---

## Code Style Guidelines

### Go
- Use meaningful variable names in English
- Comments in English
- User-facing strings in Hebrew (with English fallback)
- Follow standard Go project layout (`cmd/`, `internal/`, `pkg/`)
- Use interfaces for dependency injection
- Always handle errors explicitly - never ignore them
- Use context.Context for cancellation and timeouts

### Naming Conventions
```go
// Structs: PascalCase
type Product struct {}
type StockMovement struct {}

// Functions/Methods: PascalCase (exported) or camelCase (private)
func GetProduct() {}      // exported
func calculateTotal() {}  // private

// Variables: camelCase
productName := "Cola"
boxSize := 24

// Constants: PascalCase or ALL_CAPS for true constants
const MaxRetries = 3
const DEFAULT_EXPIRY_DAYS = 365
```

### Error Handling
```go
// Always check errors
result, err := doSomething()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Use wrapped errors for context
if err != nil {
    return fmt.Errorf("processing product %s: %w", productID, err)
}
```

---

## Domain-Specific Rules

### Products
- Products are uniquely identified by: Brand + Size + ContainerType
- Always include Hebrew name (`Name`) and English name (`NameEN`)
- `BoxSize` can be null (for items sold by kg/piece without boxes)

### Stock
- Track `QuantityBoxes` (full boxes) and `QuantityUnits` (loose items) separately
- Total = (QuantityBoxes × Product.BoxSize) + QuantityUnits

### Stock Movements
- Every change MUST be logged with:
  - `PerformedBy`: Who actually did the action
  - `ReportedBy`: Who logged it in the system
- If no name mentioned, PerformedBy = ReportedBy (self-reported)
- Types: "IN", "OUT", "WASTE", "ADJUSTMENT"

### AI Behavior
- NEVER change stock without 100% certainty
- If product is ambiguous → ASK for clarification
- Always confirm before making changes
- Parse user input for: WHO, ACTION, PRODUCT, QUANTITY

---

## Architecture

### Microservices
1. **Inventory Service** - Products, Stock, Movements
2. **Auth Service** - Users, JWT, Permissions
3. **AI Service** - Chat, Intent detection, RAG
4. **Notification Service** - Push, WhatsApp, Email

### Database
- PostgreSQL with pgvector for semantic search
- Redis for caching and notification queue

---

## Teaching Mode

This project is for LEARNING. When writing code:

1. **Explain WHY** - Don't just write code, explain the reasoning
2. **Show alternatives** - Mention other approaches and why we chose this one
3. **Link to concepts** - Reference Go patterns, design patterns, Linux commands
4. **Incremental steps** - Small, testable changes
5. **Real examples** - Use inventory domain (products, stock) in examples

---

## File Structure

```
restaurant-inventory-ai/
├── cmd/                    # Entry points
│   └── server/
├── internal/               # Private packages
│   ├── api/               # HTTP handlers
│   ├── models/            # Data structures
│   ├── repository/        # Database access
│   └── service/           # Business logic
├── pkg/                    # Public packages
├── config/                 # Configuration
├── migrations/             # Database migrations
├── docs/                   # Documentation
└── learn/                  # Learning exercises
```

---

## Hebrew Strings

When adding user-facing text, include both languages:

```go
// Messages
const (
    MsgProductNotFound   = "המוצר לא נמצא"
    MsgProductNotFoundEN = "Product not found"
    
    MsgLowStock   = "מלאי נמוך"
    MsgLowStockEN = "Low stock"
)

// Product categories
var Categories = map[string]string{
    "drinks":     "משקאות",
    "vegetables": "ירקות",
    "dairy":      "מוצרי חלב",
}
```

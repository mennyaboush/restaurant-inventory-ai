# 🤖 AI Assistant Context File

> **IMPORTANT:** Read this file at the start of each session to understand the project context.

---

## Project Overview

**Name:** Restaurant Inventory AI  
**Purpose:** AI-powered inventory management for a Pizza & Falafel restaurant  
**Real User:** The developer's father owns the restaurant  
**Language:** Hebrew primary, English secondary

---

## Architecture Decisions

### Why Microservices?
- Learning Kubernetes (multiple pods, services, scaling)
- Different scaling needs (AI needs more resources)
- Failure isolation (AI down ≠ inventory down)

### Services
1. **Inventory Service** (Go) - Products, Stock, Movements
2. **Auth Service** (Go) - Users, JWT, Permissions
3. **AI Service** (Go/Python) - Chat, Intent detection, RAG
4. **Notification Service** (Go) - Push, WhatsApp, Email

### Database: PostgreSQL
- ACID transactions (stock + movement must be atomic)
- Complex queries for reports
- pgvector for AI embeddings and semantic search

### Cache: Redis
- Session storage
- Caching
- Notification queue

---

## Key Domain Concepts

### Product Complexity
Products are identified by: **Brand + Size + Container Type**

```
"Cola Small" is ambiguous because:
├── Coca Cola 330ml Can
├── Coca Cola 330ml Plastic Bottle
├── Coca Cola 330ml Glass Bottle
└── Pepsi 330ml Can

AI must clarify which one before making changes!
```

### Stock Units
- **quantity_boxes:** Full, unopened boxes
- **quantity_units:** Loose items from opened box
- Total = (boxes × box_size) + units

### Movement Logging
Every stock change is logged with:
- **performed_by:** WHO actually did it (e.g., Yosef)
- **reported_by:** WHO logged it (e.g., Manager)

**DEFAULT RULE:** If no name mentioned → performed_by = reported_by
- "יוסף לקח 2 ארגזים" → performed_by: Yosef, reported_by: Manager
- "לקחתי 2 ארגזים" → performed_by: Manager, reported_by: Manager (self)

This enables accountability and audit trails.

### AI Behavior Rules
1. **NEVER change stock without 100% certainty**
2. If ambiguous → ASK for clarification
3. Always confirm before making changes
4. Parse: who did it, what action, what product, how much

---

## Tech Stack

| Component | Technology |
|-----------|------------|
| Backend | Go |
| Database | PostgreSQL + pgvector |
| Cache | Redis |
| LLM (Dev) | Ollama (llama3.2:1b) |
| LLM (Prod) | OpenAI GPT-4 |
| Container | Docker |
| Orchestration | Kubernetes |
| Frontend | TBD (Next.js likely) |

---

## Developer Environment

- **OS:** macOS
- **CPU:** Intel i5-8279U (4 cores)
- **RAM:** 8GB (tight - be mindful!)
- **IDE:** VS Code with Go extension

### Memory Management
- Stop Ollama when not testing AI: `brew services stop ollama`
- Stop Docker when not needed: `docker-compose down`
- Run services one at a time when possible

---

## Project Structure

```
restaurant-inventory-ai/
├── cmd/
│   └── server/
│       └── main.go          ← Main entry point
├── internal/
│   ├── api/api.go           ← HTTP handlers (empty)
│   ├── models/models.go     ← Data structures (empty)
│   ├── repository/repository.go  ← Database access (empty)
│   └── service/service.go   ← Business logic (empty)
├── learn/                   ← Learning exercises
│   └── 01_basics/main.go    ← Current lesson
├── config/                  ← Configuration
├── migrations/              ← Database migrations
├── docs/                    ← Documentation
├── .ai/                     ← AI assistant instructions
│   ├── instructions.md
│   └── skills.md
├── .github/
│   └── copilot-instructions.md
├── WORK_PLAN.md            ← Daily progress tracking
├── CONTEXT.md              ← This file
├── go.mod                  ← Go module definition
└── README.md
```

### Key Paths
- **Run learning files:** `go run learn/01_basics/main.go`
- **Run main server:** `go run cmd/server/main.go`
- **Working directory:** `/Users/mennyaboush/projects/restaurant-inventory-ai`

---

## Data Models Summary

### Product
```go
type Product struct {
    ID               string    // UUID
    Name             string    // "קוקה קולה 330 פחית"
    NameEN           string    // "Coca Cola 330ml Can"
    Brand            string    // "Coca Cola"
    Size             string    // "330ml"
    ContainerType    string    // "can"
    CategoryID       string    // FK to Category
    UnitType         string    // "bottle", "kg", "piece"
    BoxSize          int       // 24 (items per box)
    DefaultExpiryDays int      // 365
    MinStockLevel    int       // 48
    Keywords         []string  // ["קולה", "cola", "פחית"]
    IsActive         bool
}
```

### Stock
```go
type Stock struct {
    ID            string
    ProductID     string
    QuantityBoxes int     // Full boxes
    QuantityUnits int     // Loose items
    ExpiryDate    *time.Time
    UpdatedAt     time.Time
}

// TotalUnits = (QuantityBoxes * Product.BoxSize) + QuantityUnits
```

### StockMovement
```go
type StockMovement struct {
    ID           string
    ProductID    string
    PerformedBy  string    // WHO did the action
    ReportedBy   string    // WHO logged it
    Type         string    // "IN", "OUT", "WASTE", "ADJUSTMENT"
    BoxesChange  int       // +5 or -2
    UnitsChange  int       // +10 or -3
    Reason       string    // nullable
    Notes        string    // nullable
    CreatedAt    time.Time
}
```

---

## Current Progress

**Phase:** 1 - Go Fundamentals  
**Lesson:** 1.1 - Starting  
**Last Updated:** January 31, 2026

### Completed
- [x] Environment setup (Go, Docker, kubectl, psql, Ollama)
- [x] Project structure created
- [x] Requirements documented
- [x] Architecture designed
- [x] Data models finalized

### In Progress
- [ ] Go basics: variables, types, constants

### Next Up
- Structs and methods
- Build Product model in Go

---

## Teaching Approach

1. **📖 CONCEPT:** Explain the theory
2. **👀 EXAMPLE:** Show working code
3. **✋ EXERCISE:** Small practice task
4. **🔨 BUILD:** Add to the project
5. **🎯 CHECKPOINT:** Verify it works

### Linux Commands
Always explain terminal commands before running them:
- What each part does
- Why we're using this command
- What to expect as output

---

## Reminders for AI Assistant

1. **Check WORK_PLAN.md** for current progress
2. **Update WORK_PLAN.md** after completing tasks
3. **Explain WHY** not just HOW
4. **Small steps** - don't overwhelm
5. **Hebrew support** - names, UI text in Hebrew
6. **Memory constraints** - 8GB RAM, be efficient
7. **Test incrementally** - run code after each change

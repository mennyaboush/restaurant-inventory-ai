# 🎯 Project Work Plan

## Current Status
- **Phase:** 1 - Go Fundamentals + Inventory Service
- **Started:** January 31, 2026
- **Current Lesson:** 2.1 - HTTP APIs (Week 2!)

---

## Phase 1: Go Fundamentals + Inventory Service (Weeks 1-3)

### Week 1: Go Basics + Data Models ✅ COMPLETE

| Day | Lesson | Topic | Status | Notes |
|-----|--------|-------|--------|-------|
| 1 | 1.1 | Go basics: variables, types, constants | ✅ | learn/01_basics/main.go |
| 1 | 1.2 | Exercise: Create basic types for inventory | ✅ | internal/models/product.go |
| 2 | 1.3 | Structs and methods | ✅ | TotalUnits(), IsLowStock() |
| 2 | 1.4 | BUILD: Product, Category, Stock structs | ✅ | Product, Stock, StockMovement |
| 3 | 1.5 | Functions and error handling | ✅ | Validate(), error types |
| 3 | 1.6 | BUILD: Stock calculation functions | ✅ | NewProduct(), NewStockMovement() |
| 4 | 1.7 | Slices and maps | ✅ | make(), append, map[K]V |
| 4 | 1.8 | BUILD: In-memory product store | ✅ | internal/repository/memory_store.go |
| 5 | 1.9 | Packages and project organization | ✅ | models/, repository/ packages |
| 5 | 1.10 | BUILD: Organize into packages | ✅ | Already organized! |
| 6 | 1.11 | Pointers and interfaces | ✅ | Interface concepts explained |
| 6 | 1.12 | BUILD: Repository interface | ✅ | repository.go with compile check |
| 7 | 1.13 | 🎯 CHECKPOINT: Review + Quiz | ✅ | Ready for Week 2! |

### Week 2: HTTP APIs

| Day | Lesson | Topic | Status | Notes |
|-----|--------|-------|--------|-------|
| 8 | 2.1 | HTTP basics and REST principles | 🔄 | **NEXT** |
| 8 | 2.1 | HTTP basics and REST principles | ✅ | |
| 8 | 2.2 | Go net/http package | ⬜ | |
| 8 | 2.2 | Go net/http package | ✅ | |
| 9 | 2.3 | BUILD: Basic HTTP server | ⬜ | |
| 9 | 2.3 | BUILD: Basic HTTP server | ✅ | |
| 9 | 2.4 | JSON encoding/decoding | ⬜ | |
| 9 | 2.4 | JSON encoding/decoding | ✅ | |
| 10 | 2.5 | BUILD: GET /products endpoint | ⬜ | |
| 10 | 2.5 | BUILD: GET /products endpoint | ✅ | |
| 10 | 2.6 | BUILD: POST /products endpoint | ✅ | |
| 11 | 2.7 | Chi router introduction | 🔄 | **NEXT** |
| 11 | 2.7 | Chi router introduction | ✅ | |
| 11 | 2.8 | BUILD: Full CRUD for products | ⬜ | |
| 11 | 2.8 | BUILD: Full CRUD for products | ✅ | |
| 12 | 2.9 | Middleware concepts | ⬜ | |
| 12 | 2.9 | Middleware concepts | ✅ | |
| 12 | 2.10 | BUILD: Logging middleware | ⬜ | |
| 12 | 2.10 | BUILD: Logging middleware | ✅ | |
| 13 | 2.11 | Input validation | ⬜ | |
| 13 | 2.11 | Input validation | ✅ | |
| 13 | 2.12 | BUILD: Validate product input | ⬜ | |
| 13 | 2.12 | BUILD: Validate product input | ✅ | |
| 14 | 2.13 | 🎯 CHECKPOINT: API works with curl | ⬜ | |
| 14 | 2.13 | 🎯 CHECKPOINT: API works with curl | ✅ | |

### Week 3: Database Integration

| Day | Lesson | Topic | Status | Notes |
|-----|--------|-------|--------|-------|
| 15 | 3.1 | PostgreSQL and SQL basics | ⬜ | |
| 15 | 3.1 | PostgreSQL and SQL basics | 🔄 | **NEXT** |
| 16 | 3.3 | Go database/sql package | ⬜ | |
| 16 | 3.4 | BUILD: Database connection | ⬜ | |
| 17 | 3.5 | Database migrations | ⬜ | |
| 17 | 3.6 | BUILD: Create tables migration | ⬜ | |
| 18 | 3.7 | Repository pattern | ⬜ | |
| 18 | 3.8 | BUILD: PostgreSQL repository | ⬜ | |
| 19 | 3.9 | Transactions | ⬜ | |
| 19 | 3.10 | BUILD: Stock movement with transaction | ⬜ | |
| 20 | 3.11 | Testing with database | ⬜ | |
| 21 | 3.12 | 🎯 CHECKPOINT: Full CRUD with DB | ⬜ | |

---

## Phase 2: Auth Service (Week 4)

| Day | Lesson | Topic | Status | Notes |
|-----|--------|-------|--------|-------|
| 22 | 4.1 | Authentication concepts (JWT) | ⬜ | |
| 22 | 4.2 | Password hashing | ⬜ | |
| 23 | 4.3 | BUILD: User model and registration | ⬜ | |
| 23 | 4.4 | BUILD: Login endpoint | ⬜ | |
| 24 | 4.5 | JWT creation and validation | ⬜ | |
| 24 | 4.6 | BUILD: Auth middleware | ⬜ | |
| 25 | 4.7 | Role-based access control | ⬜ | |
| 25 | 4.8 | BUILD: Protect inventory endpoints | ⬜ | |
| 26 | 4.9 | 🎯 CHECKPOINT: Auth works | ⬜ | |

---

## Phase 3: Docker & Local Dev (Week 5)

| Day | Lesson | Topic | Status | Notes |
|-----|--------|-------|--------|-------|
| 27 | 5.1 | Dockerfile basics | ⬜ | |
| 27 | 5.2 | BUILD: Inventory service Dockerfile | ⬜ | |
| 28 | 5.3 | Multi-stage builds | ⬜ | |
| 28 | 5.4 | BUILD: Optimized Dockerfile | ⬜ | |
| 29 | 5.5 | Docker Compose deep dive | ⬜ | |
| 29 | 5.6 | BUILD: Full docker-compose.yaml | ⬜ | |
| 30 | 5.7 | Container networking | ⬜ | |
| 30 | 5.8 | BUILD: Services communicate | ⬜ | |
| 31 | 5.9 | 🎯 CHECKPOINT: Run with docker-compose | ⬜ | |

---

## Phase 4: Kubernetes (Weeks 6-7)

| Day | Topic | Status |
|-----|-------|--------|
| 32-33 | K8s concepts: Pods, Deployments, Services | ⬜ |
| 34-35 | BUILD: Deploy inventory to minikube | ⬜ |
| 36-37 | ConfigMaps and Secrets | ⬜ |
| 38-39 | Ingress and external access | ⬜ |
| 40-41 | Health checks and probes | ⬜ |
| 42 | 🎯 CHECKPOINT: Services running on K8s | ⬜ |

---

## Phase 5: AI Service (Weeks 8-10)

| Day | Topic | Status |
|-----|-------|--------|
| 43-45 | Ollama/OpenAI integration | ⬜ |
| 46-48 | Intent detection and entity extraction | ⬜ |
| 49-51 | RAG with pgvector | ⬜ |
| 52-54 | Hebrew NLP handling | ⬜ |
| 55-57 | Clarification flow | ⬜ |
| 58 | 🎯 CHECKPOINT: Chat works | ⬜ |

---

## Phase 6: Frontend & Polish (Weeks 11-12)

| Day | Topic | Status |
|-----|-------|--------|
| 59-63 | Next.js frontend basics | ⬜ |
| 64-66 | Mobile-friendly UI | ⬜ |
| 67-69 | Hebrew RTL support | ⬜ |
| 70 | 🎯 FINAL: MVP Complete | ⬜ |

---

## Quick Commands

```bash
# Start development
cd /Users/mennyaboush/projects/restaurant-inventory-ai
code .

# Run the Go server
go run cmd/server/main.go

# Start PostgreSQL (when needed)
docker-compose up -d postgres

# Local dev: keep DB secrets out of Git. Create a local `.env` with POSTGRES_* values and add it to `.gitignore`.
# You can source it before running commands: `set -a; source .env; set +a`

# Stop PostgreSQL
docker-compose down

# Start Ollama (when needed)
brew services start ollama

# Stop Ollama
brew services stop ollama

# Run tests
go test ./...

# Check for errors
go vet ./...
```

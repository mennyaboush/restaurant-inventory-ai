# 📚 Learning Journey - Restaurant Inventory AI

**Status:** 🟡 IN PROGRESS  
**Started:** January 31, 2026  
**Goal:** Learn AI Engineering, Go, Kubernetes, Linux through building a real project

---

## 🎓 Learning Philosophy

```
┌─────────────────────────────────────────────────────────────────┐
│                    HOW WE LEARN                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   📖 UNDERSTAND → 👀 SEE → ✋ DO → 🔨 BUILD → 🔄 REPEAT        │
│                                                                 │
│   1. Understand the concept (theory)                            │
│   2. See a working example                                      │
│   3. Do a small exercise                                        │
│   4. Build it into the project                                  │
│   5. Repeat with increasing complexity                          │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Learning Markers

| Marker | Meaning |
|--------|---------|
| 📖 **CONCEPT** | Theory and explanation |
| 👀 **EXAMPLE** | Working code to study |
| ✋ **EXERCISE** | Small practice task |
| 🔨 **BUILD** | Add to the project |
| 🎯 **CHECKPOINT** | Verify everything works |
| 💡 **TIP** | Best practice or hint |
| ⚠️ **PITFALL** | Common mistakes to avoid |
| 🧪 **TEST** | Verify your understanding |

---

## 🗺️ Learning Roadmap

### Overview

```
Phase 0: Environment Setup ............ Week 1 (Days 1-3)
Phase 1: Go Fundamentals .............. Week 1-2 (Days 3-10)
Phase 2: Building REST APIs ........... Week 2-3 (Days 10-17)
Phase 3: Database Integration ......... Week 3-4 (Days 17-24)
Phase 4: Docker & Containers .......... Week 4-5 (Days 24-31)
Phase 5: AI & RAG Integration ......... Week 5-7 (Days 31-45)
Phase 6: Kubernetes Deployment ........ Week 7-9 (Days 45-60)
Phase 7: Production & Polish .......... Week 9-10 (Days 60-70)
```

### Dev update (Feb 8, 2026)

- Moved Postgres credentials out of `docker-compose.yml` into a local `.env` file and added `.gitignore` to avoid committing secrets.
- `docker-compose.yml` now uses `env_file: .env`. Keep real passwords out of Git and use `.env` only for local development.

---

## Phase 0: Environment Setup 🛠️

**Duration:** 1-2 days  
**Goal:** Get all tools installed and working  
**Status:** ✅ Completed

### What You'll Learn
- [ ] Setting up a Go development environment
- [ ] Using VS Code for Go development
- [ ] Basic terminal/shell commands
- [ ] Docker basics
- [ ] Git basics

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 0.1 | Install Go | ⬜ | 15 min | `brew install go` |
| 0.2 | Install VS Code Go extension | ⬜ | 10 min | Extension: `golang.go` |
| 0.3 | Install Docker Desktop | ⬜ | 15 min | From docker.com |
| 0.4 | Install PostgreSQL client | ⬜ | 10 min | `brew install postgresql` |
| 0.5 | Install kubectl | ⬜ | 10 min | `brew install kubectl` |
| 0.6 | Install Ollama | ⬜ | 15 min | For local AI |
| 0.7 | Clone/init project repository | ⬜ | 10 min | Git setup |
| 0.8 | 🎯 CHECKPOINT: Run "Hello World" | ⬜ | 15 min | Verify Go works |

### Learning Resources
- [ ] [Go Installation Guide](https://go.dev/doc/install)
- [ ] [VS Code Go Setup](https://code.visualstudio.com/docs/languages/go)
- [ ] [Docker Getting Started](https://docs.docker.com/get-started/)

---

## Phase 1: Go Fundamentals 🐹

**Duration:** 5-7 days  
**Goal:** Understand Go basics well enough to build APIs  
**Status:** ✅ Completed

### What You'll Learn
- [ ] Go syntax and structure
- [ ] Variables, types, and functions
- [ ] Structs and methods
- [ ] Error handling
- [ ] Packages and modules
- [ ] Slices, maps, and loops
- [ ] Pointers (basics)
- [ ] Interfaces (basics)

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 1.1 | 📖 Go basics: variables, types | ⬜ | 1 hr | |
| 1.2 | ✋ Exercise: Calculator program | ⬜ | 30 min | Practice basics |
| 1.3 | 📖 Functions and error handling | ⬜ | 1 hr | Go's error pattern |
| 1.4 | ✋ Exercise: File reader | ⬜ | 30 min | Practice errors |
| 1.5 | 📖 Structs and methods | ⬜ | 1 hr | OOP in Go |
| 1.6 | ✋ Exercise: Inventory item struct | ⬜ | 30 min | First project code! |
| 1.7 | 📖 Slices and maps | ⬜ | 1 hr | Collections |
| 1.8 | ✋ Exercise: Inventory list | ⬜ | 30 min | |
| 1.9 | 📖 Packages and modules | ⬜ | 1 hr | go.mod, imports |
| 1.10 | 🔨 BUILD: Project structure | ⬜ | 1 hr | Set up folders |
| 1.11 | 📖 Interfaces basics | ⬜ | 1 hr | Abstraction |
| 1.12 | 📖 Pointers basics | ⬜ | 1 hr | Memory |
| 1.13 | 🎯 CHECKPOINT: Go fundamentals | ⬜ | - | Quiz yourself |

### Project Deliverable
By the end of Phase 1, you'll have:
- Basic Go project structure
- `InventoryItem` struct defined
- Simple functions to work with items

---

## Phase 2: Building REST APIs 🌐

**Duration:** 5-7 days  
**Goal:** Build HTTP APIs in Go  
**Status:** ✅ Completed

### What You'll Learn
- [ ] HTTP basics (methods, status codes)
- [ ] REST API design principles
- [ ] Go's `net/http` package
- [ ] Routing and handlers
- [ ] JSON encoding/decoding
- [ ] Middleware concepts
- [ ] Input validation
- [ ] API testing

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 2.1 | 📖 HTTP and REST fundamentals | ✅ | 1 hr | Theory |
| 2.2 | 📖 Go net/http package | ✅ | 1 hr | |
| 2.3 | 🔨 BUILD: Basic HTTP server | ✅ | 30 min | Hello World API |
| 2.4 | 📖 Routing patterns | ✅ | 1 hr | Chi router |
| 2.5 | 🔨 BUILD: Health check endpoint | ✅ | 30 min | /health |
| 2.6 | 📖 JSON in Go | ✅ | 1 hr | Marshal/Unmarshal |
| 2.7 | 🔨 BUILD: GET /inventory | ✅ | 1 hr | List items |
| 2.8 | 🔨 BUILD: POST /inventory | ✅ | 1 hr | Create item |
| 2.9 | 🔨 BUILD: PUT /inventory/:id | ✅ | 1 hr | Update item |
| 2.10 | 🔨 BUILD: DELETE /inventory/:id | ✅ | 1 hr | Delete item |
| 2.11 | 📖 Middleware concepts | ✅ | 1 hr | Logging, auth |
| 2.12 | 🔨 BUILD: Logging middleware | ✅ | 30 min | |
| 2.13 | 📖 Input validation | ✅ | 1 hr | |
| 2.14 | 🔨 BUILD: Validate inventory input | ✅ | 30 min | |
| 2.15 | 🧪 TEST: API with curl/Postman | ✅ | 30 min | |
| 2.16 | 🎯 CHECKPOINT: CRUD API works | ✅ | - | |

### Project Deliverable
By the end of Phase 2, you'll have:
- Working REST API for inventory CRUD
- Proper routing and JSON handling
- Logging middleware
- API documentation

---

## Phase 3: Database Integration 💾

**Duration:** 5-7 days  
**Goal:** Store data persistently in PostgreSQL  
**Status:** 🟡 IN PROGRESS

### What You'll Learn
- [ ] SQL fundamentals (SELECT, INSERT, UPDATE, DELETE)
- [ ] PostgreSQL setup and usage
- [ ] Go database/sql package
- [ ] Connection pooling
- [ ] Migrations
- [ ] Repository pattern
- [ ] Transactions

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 3.1 | 📖 Database fundamentals | ⬜ | 1 hr | Relational concepts |
| 3.2 | 📖 SQL basics | ⬜ | 2 hr | Queries |
| 3.3 | 🔨 Run PostgreSQL in Docker | ⬜ | 30 min | docker-compose |
| 3.4 | ✋ Exercise: SQL queries | ⬜ | 1 hr | Practice SQL |
| 3.5 | 📖 Go database patterns | ⬜ | 1 hr | database/sql |
| 3.6 | 🔨 BUILD: Database connection | ⬜ | 1 hr | |
| 3.7 | 📖 Database migrations | ⬜ | 1 hr | Schema versions |
| 3.8 | 🔨 BUILD: Inventory schema | ⬜ | 1 hr | Tables |
| 3.9 | 🔨 BUILD: Repository layer | ⬜ | 2 hr | Data access |
| 3.10 | 🔨 BUILD: Connect API to DB | ⬜ | 1 hr | Wire it up |
| 3.11 | 📖 Transactions | ⬜ | 1 hr | |
| 3.12 | 🧪 TEST: Data persists | ⬜ | 30 min | |
| 3.13 | 🎯 CHECKPOINT: Full CRUD with DB | ⬜ | - | |

### Project Deliverable
By the end of Phase 3, you'll have:
- PostgreSQL database running
- Tables for inventory, suppliers, etc.
- API connected to real database
- Data persists between restarts

### Next Actions (Feb 9, 2026)

- Short verdict on our current Postgres approach:
  - Development: using `docker-compose` + a local `.env` file is fine for local development and testing. Do not commit `.env` to Git. The `cmd/dev` helper and the integration test are convenient for working locally.
  - Production: treat Postgres as a managed/secure service. Use secrets managers (or environment injection in CI), enable SSL, backups, monitoring, and a proven migration tool.

- Immediate next steps (dev):
  1. Run the integration tests: `go test -tags=integration ./internal/repository -v` (verifies DB + migrations + basic queries).
  2. Implement `postgres_store.go` to replace in-memory repository with `database/sql` calls and connection pooling.
  3. Add a proper migration tool (e.g., `golang-migrate`) and wire it into CI for ordered, idempotent migrations.

- Postgres best-practices summary:
  - NEVER commit secrets. Use `.env` for local dev and a secret manager for CI/prod.
  - Use a migration tool for ordered migrations, locking, and rollbacks.
  - Configure connection pooling and timeouts in the app.
  - Monitor DB health, size, and slow queries; schedule regular backups.
  - Prefer managed DBs in production (AWS RDS, Cloud SQL, etc.) for ops simplicity.

---

## Phase 4: Docker & Containers 🐳

**Duration:** 5-7 days  
**Goal:** Containerize the application  
**Status:** ⬜ Not Started

### What You'll Learn
- [ ] Container concepts (vs VMs)
- [ ] Dockerfile syntax
- [ ] Multi-stage builds
- [ ] Docker Compose
- [ ] Container networking
- [ ] Volumes and persistence
- [ ] Container debugging

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 4.1 | 📖 Container fundamentals | ⬜ | 1 hr | Theory |
| 4.2 | 📖 Dockerfile syntax | ⬜ | 1 hr | |
| 4.3 | 🔨 BUILD: Basic Dockerfile | ⬜ | 1 hr | |
| 4.4 | 📖 Multi-stage builds | ⬜ | 1 hr | Smaller images |
| 4.5 | 🔨 BUILD: Optimized Dockerfile | ⬜ | 1 hr | |
| 4.6 | 📖 Docker Compose | ⬜ | 1 hr | Multi-container |
| 4.7 | 🔨 BUILD: docker-compose.yaml | ⬜ | 1 hr | Full stack |
| 4.8 | 📖 Container networking | ⬜ | 1 hr | |
| 4.9 | 📖 Volumes | ⬜ | 30 min | Data persistence |
| 4.10 | ✋ Exercise: Debug container | ⬜ | 1 hr | Exec, logs |
| 4.11 | 🧪 TEST: Full stack in Docker | ⬜ | 30 min | |
| 4.12 | 🎯 CHECKPOINT: docker-compose up | ⬜ | - | |

### Project Deliverable
By the end of Phase 4, you'll have:
- Production-ready Dockerfile
- docker-compose.yaml for local dev
- All services running in containers

---

## Phase 5: AI & RAG Integration 🤖

**Duration:** 10-14 days  
**Goal:** Add AI capabilities using RAG  
**Status:** ⬜ Not Started

### What You'll Learn
- [ ] AI/ML fundamentals (high level)
- [ ] LLM APIs (OpenAI, Ollama)
- [ ] Prompt engineering
- [ ] Embeddings concepts
- [ ] Vector databases (Qdrant)
- [ ] RAG architecture
- [ ] AI application patterns

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 5.1 | 📖 AI/ML overview for engineers | ⬜ | 2 hr | Concepts |
| 5.2 | 📖 What are LLMs | ⬜ | 1 hr | |
| 5.3 | 🔨 Set up Ollama locally | ⬜ | 1 hr | Local AI |
| 5.4 | ✋ Exercise: Chat with Ollama | ⬜ | 30 min | |
| 5.5 | 📖 LLM APIs | ⬜ | 1 hr | API patterns |
| 5.6 | 🔨 BUILD: LLM client in Go | ⬜ | 2 hr | |
| 5.7 | 📖 Prompt engineering | ⬜ | 2 hr | Best practices |
| 5.8 | ✋ Exercise: Write good prompts | ⬜ | 1 hr | |
| 5.9 | 📖 What are embeddings | ⬜ | 2 hr | Vector concepts |
| 5.10 | 📖 Vector databases | ⬜ | 1 hr | Qdrant |
| 5.11 | 🔨 Set up Qdrant | ⬜ | 1 hr | Docker |
| 5.12 | 📖 RAG architecture | ⬜ | 2 hr | Full pattern |
| 5.13 | 🔨 BUILD: Document ingestion | ⬜ | 2 hr | Load docs |
| 5.14 | 🔨 BUILD: Embedding generation | ⬜ | 2 hr | Vectorize |
| 5.15 | 🔨 BUILD: Vector search | ⬜ | 2 hr | Find similar |
| 5.16 | 🔨 BUILD: RAG pipeline | ⬜ | 3 hr | Full flow |
| 5.17 | 🔨 BUILD: AI chat endpoint | ⬜ | 2 hr | /ask API |
| 5.18 | 🔨 BUILD: Demand prediction | ⬜ | 3 hr | AI feature |
| 5.19 | 🔨 BUILD: Order suggestions | ⬜ | 2 hr | AI feature |
| 5.20 | 🧪 TEST: AI features work | ⬜ | 1 hr | |
| 5.21 | 🎯 CHECKPOINT: RAG working | ⬜ | - | |

### Project Deliverable
By the end of Phase 5, you'll have:
- Local AI with Ollama
- Vector database with documents
- RAG pipeline working
- AI-powered features:
  - Natural language queries
  - Demand prediction
  - Order suggestions

---

## Phase 6: Kubernetes Deployment ☸️

**Duration:** 10-14 days  
**Goal:** Deploy to Kubernetes  
**Status:** ⬜ Not Started

### What You'll Learn
- [ ] Why Kubernetes exists
- [ ] K8s architecture
- [ ] Pods, Deployments, Services
- [ ] ConfigMaps and Secrets
- [ ] Persistent Volumes
- [ ] Ingress
- [ ] kubectl commands
- [ ] Debugging in K8s

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 6.1 | 📖 Why Kubernetes? | ⬜ | 1 hr | The problem it solves |
| 6.2 | 📖 K8s architecture | ⬜ | 2 hr | Control plane, nodes |
| 6.3 | 🔨 Set up minikube | ⬜ | 30 min | Local K8s |
| 6.4 | 📖 Pods | ⬜ | 1 hr | Basic unit |
| 6.5 | 📖 Deployments | ⬜ | 1 hr | Managing pods |
| 6.6 | 📖 Services | ⬜ | 1 hr | Networking |
| 6.7 | 🔨 BUILD: Backend deployment | ⬜ | 1 hr | |
| 6.8 | 🔨 BUILD: Backend service | ⬜ | 30 min | |
| 6.9 | 📖 ConfigMaps & Secrets | ⬜ | 1 hr | Configuration |
| 6.10 | 🔨 BUILD: Config for app | ⬜ | 1 hr | |
| 6.11 | 📖 Persistent Volumes | ⬜ | 1 hr | Storage |
| 6.12 | 🔨 BUILD: Database with PV | ⬜ | 1 hr | |
| 6.13 | 🔨 BUILD: Vector DB deployment | ⬜ | 1 hr | |
| 6.14 | 📖 Ingress | ⬜ | 1 hr | External access |
| 6.15 | 🔨 BUILD: Ingress rules | ⬜ | 1 hr | |
| 6.16 | 📖 K8s debugging | ⬜ | 2 hr | Troubleshooting |
| 6.17 | ✋ Exercise: Fix broken deploy | ⬜ | 1 hr | Practice |
| 6.18 | 🧪 TEST: Full app in K8s | ⬜ | 1 hr | |
| 6.19 | 🎯 CHECKPOINT: App runs in K8s | ⬜ | - | |

### Project Deliverable
By the end of Phase 6, you'll have:
- All K8s manifests written
- App running in minikube
- Database with persistent storage
- External access configured

---

## Phase 7: Production & Polish 🚀

**Duration:** 7-10 days  
**Goal:** Production-ready application  
**Status:** ⬜ Not Started

### What You'll Learn
- [ ] Logging best practices
- [ ] Health checks
- [ ] Monitoring basics
- [ ] Security basics
- [ ] CI/CD concepts
- [ ] Documentation

### Tasks

| # | Task | Status | Time | Notes |
|---|------|--------|------|-------|
| 7.1 | 📖 Structured logging | ⬜ | 1 hr | |
| 7.2 | 🔨 BUILD: Better logging | ⬜ | 1 hr | |
| 7.3 | 📖 Health checks | ⬜ | 1 hr | Liveness, readiness |
| 7.4 | 🔨 BUILD: Health endpoints | ⬜ | 1 hr | |
| 7.5 | 📖 Basic security | ⬜ | 2 hr | |
| 7.6 | 🔨 BUILD: Authentication | ⬜ | 3 hr | |
| 7.7 | 📖 Monitoring overview | ⬜ | 1 hr | |
| 7.8 | 🔨 BUILD: Basic metrics | ⬜ | 2 hr | |
| 7.9 | 📖 CI/CD concepts | ⬜ | 1 hr | |
| 7.10 | 🔨 BUILD: GitHub Actions | ⬜ | 2 hr | |
| 7.11 | 🔨 BUILD: Frontend UI | ⬜ | 4 hr | Simple dashboard |
| 7.12 | 📖 Documentation | ⬜ | 1 hr | |
| 7.13 | 🔨 Write final docs | ⬜ | 2 hr | |
| 7.14 | 🎯 CHECKPOINT: Production ready | ⬜ | - | |

---

## 📊 Progress Dashboard

| Phase | Progress | Status |
|-------|----------|--------|
| Phase 0: Setup | ██████████ 100% | ✅ Completed |
| Phase 1: Go | ██████████ 100% | ✅ Completed |
| Phase 2: APIs | ██████████ 100% | ✅ Completed |
| Phase 3: Database | ░░░░░░░░░░ 0% | 🟡 IN PROGRESS |
| Phase 4: Docker | ░░░░░░░░░░ 0% | ⬜ Not Started |
| Phase 5: AI/RAG | ░░░░░░░░░░ 0% | ⬜ Not Started |
| Phase 6: K8s | ░░░░░░░░░░ 0% | ⬜ Not Started |
| Phase 7: Polish | ░░░░░░░░░░ 0% | ⬜ Not Started |

**Overall:** ███░░░░░░░ 30%

---

## 📝 Learning Notes

*Write your notes, questions, and insights here as you progress:*

### General Notes
```
- 
```

### Questions to Research
```
- 
```

### "Aha!" Moments
```
- 
```

---

## 🏆 Milestones

| Milestone | Target Date | Actual Date | Status |
|-----------|-------------|-------------|--------|
| Environment ready | | February 7, 2026 | ✅ Completed |
| First Go program | | February 7, 2026 | ✅ Completed |
| API returns JSON | | February 7, 2026 | ✅ Completed |
| Database working | | | 🟡 IN PROGRESS |
| Running in Docker | | | ⬜ Not Started |
| First AI response | | | ⬜ Not Started |
| RAG working | | | ⬜ Not Started |
| Running in K8s | | | ⬜ Not Started |
| Production ready | | | ⬜ Not Started |
| Father using it! | | | ⬜ Not Started |

---

*Every expert was once a beginner. Let's go! 🚀*

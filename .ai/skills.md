# 🎓 Learning Skills & Exercises

> Structured learning path with clear skills to acquire

---

## Skill Tree Overview

```
                        🎯 RESTAURANT INVENTORY AI
                                   │
        ┌──────────────────────────┼──────────────────────────┐
        │                          │                          │
   🐹 GO LANG              🐳 INFRASTRUCTURE           🤖 AI/ML
        │                          │                          │
   ├── Basics              ├── Docker                 ├── LLM APIs
   ├── Structs             ├── Compose                ├── Embeddings
   ├── Interfaces          ├── Kubernetes             ├── RAG
   ├── HTTP/REST           ├── Networking             ├── Intent Detection
   ├── Database            └── CI/CD                  └── Hebrew NLP
   └── Testing
```

---

## Phase 1: Go Fundamentals

### Skill 1.1: Variables & Types
**Goal:** Understand Go's type system

**Concepts:**
- [ ] Basic types (int, string, bool, float64)
- [ ] Variable declaration (var, :=, const)
- [ ] Zero values
- [ ] Type conversion

**Exercise:** Create variables for a product (name, price, quantity)

**Assessment:** Can you explain why Go won't compile if a variable is unused?

---

### Skill 1.2: Structs & Methods
**Goal:** Model real-world entities

**Concepts:**
- [ ] Struct declaration
- [ ] Struct fields and tags
- [ ] Methods with receivers
- [ ] Pointer vs value receivers

**Exercise:** Create Product struct with a method to calculate total value

**Assessment:** When would you use pointer receiver vs value receiver?

---

### Skill 1.3: Functions & Error Handling
**Goal:** Write robust functions

**Concepts:**
- [ ] Function declaration
- [ ] Multiple return values
- [ ] Error as a value
- [ ] Error wrapping
- [ ] defer, panic, recover

**Exercise:** Write function to validate product data, returning errors

**Assessment:** Why does Go use explicit error returns instead of exceptions?

---

### Skill 1.4: Collections (Slices & Maps)
**Goal:** Work with data collections

**Concepts:**
- [ ] Arrays vs Slices
- [ ] Slice operations (append, copy, slice)
- [ ] Maps (creation, access, iteration)
- [ ] Make vs literal syntax

**Exercise:** Create a product inventory using a map

**Assessment:** What happens if you access a key that doesn't exist in a map?

---

### Skill 1.5: Packages & Modules
**Goal:** Organize code professionally

**Concepts:**
- [ ] Package declaration
- [ ] Import paths
- [ ] go.mod and go.sum
- [ ] Exported vs unexported
- [ ] Internal packages

**Exercise:** Split code into models, service, and repository packages

**Assessment:** Why can't external packages import from `internal/`?

---

### Skill 1.6: Interfaces
**Goal:** Write flexible, testable code

**Concepts:**
- [ ] Interface declaration
- [ ] Implicit implementation
- [ ] Interface composition
- [ ] Empty interface (any)
- [ ] Type assertions

**Exercise:** Create Repository interface for product storage

**Assessment:** How do interfaces enable easier testing?

---

### Skill 1.7: Pointers
**Goal:** Understand memory and references

**Concepts:**
- [ ] Pointer declaration (&, *)
- [ ] nil pointers
- [ ] Pass by value vs pass by reference
- [ ] When to use pointers

**Exercise:** Modify a product in a slice using pointers

**Assessment:** When should you pass a pointer vs a value?

---

## Phase 2: HTTP & APIs

### Skill 2.1: HTTP Fundamentals
**Concepts:**
- [ ] HTTP methods (GET, POST, PUT, DELETE)
- [ ] Status codes
- [ ] Headers
- [ ] Request/Response body
- [ ] REST principles

### Skill 2.2: Go net/http
**Concepts:**
- [ ] http.HandleFunc
- [ ] http.ListenAndServe
- [ ] Request parsing
- [ ] Response writing
- [ ] Middleware pattern

### Skill 2.3: JSON Handling
**Concepts:**
- [ ] json.Marshal / json.Unmarshal
- [ ] Struct tags
- [ ] Custom marshaling
- [ ] Handling null values

### Skill 2.4: Routing (Chi)
**Concepts:**
- [ ] Router setup
- [ ] URL parameters
- [ ] Middleware
- [ ] Grouping routes

---

## Phase 3: Database

### Skill 3.1: SQL Basics
**Concepts:**
- [ ] SELECT, INSERT, UPDATE, DELETE
- [ ] WHERE, ORDER BY, LIMIT
- [ ] JOINs
- [ ] Transactions

### Skill 3.2: Go database/sql
**Concepts:**
- [ ] Connection pooling
- [ ] Query vs QueryRow vs Exec
- [ ] Prepared statements
- [ ] Scanning results

### Skill 3.3: Migrations
**Concepts:**
- [ ] Schema versioning
- [ ] Up and Down migrations
- [ ] Migration tools

### Skill 3.4: Repository Pattern
**Concepts:**
- [ ] Interface definition
- [ ] PostgreSQL implementation
- [ ] Transaction handling
- [ ] Testing with mocks

---

## Phase 4: Docker & Containers

### Skill 4.1: Docker Basics
**Concepts:**
- [ ] Images vs Containers
- [ ] Dockerfile syntax
- [ ] Build and run
- [ ] Volumes and networking

### Skill 4.2: Multi-stage Builds
**Concepts:**
- [ ] Build stage vs runtime stage
- [ ] Minimizing image size
- [ ] Security best practices

### Skill 4.3: Docker Compose
**Concepts:**
- [ ] Services definition
- [ ] Networking between containers
- [ ] Environment variables
- [ ] Volume persistence

---

## Phase 5: Kubernetes

### Skill 5.1: K8s Concepts
**Concepts:**
- [ ] Pods, Deployments, Services
- [ ] ConfigMaps, Secrets
- [ ] Namespaces
- [ ] Labels and Selectors

### Skill 5.2: Deploying Applications
**Concepts:**
- [ ] Writing manifests
- [ ] kubectl commands
- [ ] Rolling updates
- [ ] Health checks

### Skill 5.3: Ingress & Networking
**Concepts:**
- [ ] Service types (ClusterIP, NodePort, LoadBalancer)
- [ ] Ingress controllers
- [ ] TLS/SSL

---

## Phase 6: AI & RAG

### Skill 6.1: LLM Integration
**Concepts:**
- [ ] API calls to Ollama/OpenAI
- [ ] Prompt engineering
- [ ] Token management
- [ ] Response parsing

### Skill 6.2: Embeddings & Vectors
**Concepts:**
- [ ] What are embeddings
- [ ] Creating embeddings
- [ ] pgvector storage
- [ ] Similarity search

### Skill 6.3: RAG Architecture
**Concepts:**
- [ ] Retrieval step
- [ ] Augmentation step
- [ ] Generation step
- [ ] Context window management

### Skill 6.4: Intent Detection
**Concepts:**
- [ ] Classifying user intent
- [ ] Entity extraction
- [ ] Slot filling
- [ ] Clarification dialogs

---

## Linux/Shell Skills (Ongoing)

### Essential Commands
- [ ] Navigation: cd, ls, pwd
- [ ] Files: cat, less, head, tail, grep
- [ ] Permissions: chmod, chown
- [ ] Processes: ps, top, kill
- [ ] Networking: curl, netstat, ping

### Shell Scripting
- [ ] Variables and expansion
- [ ] Conditionals (if, case)
- [ ] Loops (for, while)
- [ ] Pipes and redirection
- [ ] Functions

---

## Progress Tracking

After completing each skill:
1. Mark checkbox in this file
2. Update WORK_PLAN.md
3. Note any questions for review
4. Build something with the skill

# 🎓 System Design Review - Understanding Every Decision

## Table of Contents
1. [Architecture Pattern: Why Microservices?](#1-architecture-pattern)
2. [Database Choices: Why PostgreSQL?](#2-database-choices)
3. [Data Models: Why This Structure?](#3-data-models)
4. [API Design: Why REST?](#4-api-design)
5. [AI Architecture: Why RAG?](#5-ai-architecture)
6. [Technology Stack: Every Choice Explained](#6-technology-stack)

---

## 1. Architecture Pattern

### The Options

| Pattern | Description | Pros | Cons |
|---------|-------------|------|------|
| **Monolith** | Single application, single deployment | Simple, fast to develop, easy debugging | Hard to scale parts independently, single point of failure |
| **Microservices** | Multiple small services, independent deployment | Scale independently, team ownership, technology flexibility | Complex, network overhead, harder debugging |
| **Serverless** | Functions as a Service (AWS Lambda) | Pay per use, auto-scale | Cold starts, vendor lock-in, complex state management |
| **Modular Monolith** | Single deployment, but well-separated modules | Balance of simplicity and organization | Still single deployment |

### Our Choice: Microservices

**Why?**

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    WHY MICROSERVICES FOR THIS PROJECT                   │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  1. LEARNING KUBERNETES                                                 │
│     └── Microservices = Multiple pods, services, deployments           │
│     └── Real K8s experience (service discovery, scaling, networking)   │
│                                                                         │
│  2. DIFFERENT SCALING NEEDS                                             │
│     └── AI Service: Heavy compute, needs GPU                           │
│     └── Inventory Service: Light, but frequent requests                │
│     └── Notification Service: Bursty (many alerts at once)             │
│                                                                         │
│  3. TECHNOLOGY FLEXIBILITY                                              │
│     └── AI Service might need Python for ML libraries                  │
│     └── Other services in Go                                           │
│                                                                         │
│  4. INDEPENDENT DEPLOYMENT                                              │
│     └── Fix bug in Notifications without touching Inventory            │
│     └── Upgrade AI model without downtime for stock tracking           │
│                                                                         │
│  5. FAILURE ISOLATION                                                   │
│     └── AI service crashes? Inventory still works                      │
│     └── Notification down? Core features still available               │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### When Monolith Would Be Better

- If only 1-2 developers
- If time-to-market is critical
- If you don't need to learn Kubernetes
- If all parts have same scaling needs

---

## 2. Database Choices

### The Options

| Database | Type | Best For | Not Good For |
|----------|------|----------|--------------|
| **PostgreSQL** | Relational (SQL) | Structured data, complex queries, transactions | Huge scale, unstructured data |
| **MySQL** | Relational (SQL) | Similar to PostgreSQL | Complex queries, JSON |
| **MongoDB** | Document (NoSQL) | Flexible schema, rapid prototyping | Complex joins, transactions |
| **Redis** | Key-Value | Caching, sessions, queues | Primary storage, complex queries |
| **Elasticsearch** | Search engine | Full-text search, logs | Primary storage |
| **TimescaleDB** | Time-series | Metrics, analytics over time | General purpose |
| **Vector DBs** (Pinecone, pgvector) | Vector | AI embeddings, similarity search | Regular queries |

### Our Choices

#### PostgreSQL (Main Database)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    WHY POSTGRESQL                                       │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  ✅ ACID Transactions                                                   │
│     "Took 5 items" must update stock AND create movement atomically    │
│     If one fails, both fail (no inconsistent data)                     │
│                                                                         │
│  ✅ Complex Queries                                                     │
│     "Show me products low in stock, grouped by category, with          │
│      last movement date and supplier info"                             │
│     → SQL JOINs handle this naturally                                  │
│                                                                         │
│  ✅ Data Integrity                                                      │
│     Foreign keys: Can't delete category if products exist              │
│     Constraints: Stock can't go negative                               │
│     Types: Quantity must be integer, price must be decimal             │
│                                                                         │
│  ✅ pgvector Extension                                                  │
│     Vector embeddings for AI in same database                          │
│     No need for separate vector database                               │
│                                                                         │
│  ✅ JSON Support                                                        │
│     Can store flexible data when needed (product metadata)             │
│                                                                         │
│  ✅ Industry Standard                                                   │
│     Skills transfer to any job                                         │
│     Great documentation, community support                             │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### Why NOT MongoDB?

```
MongoDB would work, but:

❌ No enforced schema
   → Developer can accidentally save "quantity: 'five'" instead of 5
   → PostgreSQL prevents this at database level

❌ Transactions are newer/complex
   → Stock update + movement must be atomic
   → PostgreSQL has 30+ years of reliable transactions

❌ Joins are expensive
   → "Products with category name and supplier" requires $lookup
   → PostgreSQL JOINs are optimized for this

❌ Learning SQL is more valuable
   → Used in 70%+ of companies
   → MongoDB skills are less transferable
```

#### Redis (Cache & Queue)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    WHY REDIS                                            │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  USE CASE 1: Session Storage                                           │
│  └── Store JWT tokens for fast validation                              │
│  └── 0.1ms vs 5ms database lookup                                      │
│                                                                         │
│  USE CASE 2: Caching                                                   │
│  └── Cache product list (changes rarely)                               │
│  └── Cache stock levels (with short TTL)                               │
│  └── Reduce database load                                              │
│                                                                         │
│  USE CASE 3: Rate Limiting                                             │
│  └── "User can only make 100 requests/minute"                          │
│  └── Redis INCR with TTL is perfect for this                           │
│                                                                         │
│  USE CASE 4: Message Queue                                             │
│  └── "Stock low" → Queue notification                                  │
│  └── Notification service processes queue                              │
│  └── Decouples services                                                │
│                                                                         │
│  WHY NOT RabbitMQ/Kafka for queue?                                     │
│  └── Redis is simpler for our scale                                    │
│  └── One less system to manage                                         │
│  └── Can upgrade to Kafka later if needed                              │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Data Models

### Product Model - Why This Structure?

```sql
PRODUCT
├── id (UUID)                    -- Why UUID? See below
├── name (string)                -- Full display name "קוקה קולה 330 פחית"
├── name_en (string, nullable)   -- For English UI option
├── brand (string)               -- "Coca Cola" - for grouping/filtering
├── size (string)                -- "330ml" - for filtering
├── container_type (string)      -- "can/plastic/glass" - differentiator
├── category_id (FK)             -- Link to category
├── unit_type (string)           -- "bottle/kg/piece" - base unit
├── box_size (int, nullable)     -- Items per box (null if no box)
├── default_expiry_days (int)    -- For calculating expiry
├── min_stock_level (int)        -- When to alert
├── keywords (text[])            -- For search ["קולה", "cola", "פחית"]
└── is_active (bool)             -- Soft delete
```

#### Why UUID instead of Auto-Increment ID?

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    UUID vs AUTO-INCREMENT                               │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  AUTO-INCREMENT (1, 2, 3, 4...)                                        │
│  ├── ✅ Simple, readable                                                │
│  ├── ✅ Smaller storage (4 bytes vs 16 bytes)                          │
│  ├── ❌ Exposes info ("We have ~5000 products")                        │
│  ├── ❌ Problems with database sharding/replication                     │
│  ├── ❌ Must hit database to generate ID                               │
│  └── ❌ Conflicts when merging databases                               │
│                                                                         │
│  UUID (550e8400-e29b-41d4-a716-446655440000)                           │
│  ├── ✅ Generate anywhere (client, server, database)                   │
│  ├── ✅ No conflicts ever (practically unique globally)                │
│  ├── ✅ Works with microservices (no central ID generator)             │
│  ├── ✅ Secure (can't guess next ID)                                   │
│  ├── ❌ Larger (16 bytes)                                              │
│  └── ❌ Not human-readable                                             │
│                                                                         │
│  FOR MICROSERVICES → UUID is better                                    │
│  (Each service can generate IDs independently)                         │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### Why Separate brand, size, container_type?

```
Option A: Single name field
name = "קוקה קולה 330 פחית"

→ How to find all "Coca Cola" products? LIKE '%קולה%' (slow, error-prone)
→ How to find all "330ml" products? LIKE '%330%' (might match "330" in other text)

Option B: Separate fields ✅
brand = "Coca Cola"
size = "330ml"  
container_type = "can"

→ Find all Coca Cola: WHERE brand = 'Coca Cola' (fast, exact)
→ Find all cans: WHERE container_type = 'can' (fast, exact)
→ AI can ask: "Which container type? Can, Plastic, or Glass?"
```

### Stock Model - Why This Structure?

```sql
STOCK
├── id (UUID)
├── product_id (FK)           -- One stock record per product
├── quantity_boxes (int)      -- Full, unopened boxes
├── quantity_units (int)      -- Loose items from opened boxes
├── expiry_date (date, null)  -- Earliest expiry in stock
└── updated_at (timestamp)    -- When last changed
```

#### Why Separate boxes and units?

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    BOXES + UNITS MODEL                                  │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  Real-world scenario:                                                  │
│  "We have 3 boxes of Cola and 5 loose bottles"                         │
│                                                                         │
│  Option A: Store only total units                                      │
│  quantity = 77  (3 × 24 + 5)                                           │
│  ❌ Lost information: How many BOXES do we have?                       │
│  ❌ When ordering, supplier thinks in BOXES                            │
│  ❌ "We have 77 bottles" is less useful than "3 boxes + 5"            │
│                                                                         │
│  Option B: Store boxes and units separately ✅                         │
│  quantity_boxes = 3                                                    │
│  quantity_units = 5                                                    │
│  ✅ Can display: "3 ארגזים + 5 בקבוקים"                                │
│  ✅ Can calculate total: (3 × 24) + 5 = 77 when needed                 │
│  ✅ Natural for ordering: "Need to order 2 boxes"                      │
│  ✅ Reflects how inventory is actually stored                          │
│                                                                         │
│  Helper function:                                                      │
│  TotalUnits() = (quantity_boxes × product.box_size) + quantity_units  │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### Movement Model - Why Track History?

```sql
STOCK_MOVEMENT
├── id (UUID)
├── product_id (FK)
├── user_id (FK)              -- WHO made this change
├── type (enum)               -- IN/OUT/WASTE/ADJUSTMENT
├── boxes_change (int)        -- +5 or -2
├── units_change (int)        -- +10 or -3
├── reason (string, null)     -- "Expired", "Damaged", etc.
├── notes (string, null)      -- Free text
└── created_at (timestamp)    -- WHEN it happened
```

#### Why Log Every Movement?

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    WHY MOVEMENT HISTORY                                 │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  1. AUDIT TRAIL                                                        │
│     "Who took 50 bottles yesterday?" → Check movements                 │
│     "The stock is wrong" → Review history to find error                │
│                                                                         │
│  2. ANALYTICS                                                          │
│     "How much cola do we use per week?" → SUM(OUT) WHERE product=cola  │
│     "What day has most usage?" → GROUP BY day_of_week                  │
│     "What's our waste?" → SUM(WASTE)                                   │
│                                                                         │
│  3. AI PREDICTIONS                                                     │
│     "Based on last 30 days, you'll need 10 boxes next week"           │
│     → Only possible with historical data                               │
│                                                                         │
│  4. ACCOUNTABILITY                                                     │
│     Each movement has user_id                                          │
│     "David recorded 100 exits last week" - is that normal?            │
│                                                                         │
│  5. DEBUGGING                                                          │
│     Stock shows 50, should be 100                                      │
│     → Look at movements to find the discrepancy                        │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 4. API Design

### REST vs GraphQL vs gRPC

| Protocol | Best For | Pros | Cons |
|----------|----------|------|------|
| **REST** | Public APIs, web/mobile clients | Simple, widely understood, cacheable | Over/under-fetching, multiple round trips |
| **GraphQL** | Complex client needs, varying data | Client specifies exact data needed | Complexity, caching harder, learning curve |
| **gRPC** | Service-to-service, high performance | Fast (binary), streaming, type-safe | Not browser-friendly, harder to debug |

### Our Choice: REST (External) + gRPC (Internal, Future)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    API PROTOCOL DECISIONS                               │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  CLIENT → SERVICES: REST                                               │
│  ├── Browser/mobile can call directly                                  │
│  ├── Easy to test with curl, Postman                                   │
│  ├── Easy to debug (readable JSON)                                     │
│  ├── Cacheable (GET requests)                                          │
│  └── Everyone knows REST                                               │
│                                                                         │
│  SERVICE → SERVICE: REST now, gRPC later                               │
│  ├── Start simple (REST everywhere)                                    │
│  ├── If performance matters → upgrade to gRPC                          │
│  └── AI ↔ Inventory might benefit from gRPC                           │
│                                                                         │
│  WHY NOT GraphQL?                                                       │
│  ├── Our queries aren't that complex                                   │
│  ├── Learning curve for team                                           │
│  ├── REST is enough for MVP                                            │
│  └── Can add GraphQL layer later if clients need it                   │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 5. AI Architecture

### RAG (Retrieval Augmented Generation) - Why?

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    LLM ALONE vs RAG                                     │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  LLM ALONE:                                                            │
│  User: "How much cola do we have?"                                     │
│  LLM: "I don't know your inventory." ❌                                │
│                                                                         │
│  WITH RAG:                                                              │
│  1. User: "How much cola do we have?"                                  │
│  2. System: Search database for "cola" products                        │
│  3. System: Get stock levels                                           │
│  4. System: Send to LLM with context:                                  │
│     "Inventory data:                                                   │
│      - Cola Small Can: 3 boxes + 5 units                               │
│      - Cola Large Plastic: 2 boxes                                     │
│      User asks: How much cola do we have?"                             │
│  5. LLM: "You have 3 boxes and 5 bottles of small cola cans,           │
│           plus 2 boxes of large cola plastic bottles." ✅              │
│                                                                         │
│  RAG = Give LLM relevant context before asking                         │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### Why pgvector Instead of Pinecone/Weaviate?

```
Vector databases store "embeddings" - numerical representations of text
that capture meaning. Similar texts have similar vectors.

"קולה קטנה" → [0.2, 0.5, -0.1, 0.8, ...]
"cola small" → [0.21, 0.48, -0.12, 0.79, ...]  (similar vectors!)

┌─────────────────────────────────────────────────────────────────────────┐
│                    VECTOR DB OPTIONS                                    │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  PINECONE (Managed service)                                            │
│  ├── ✅ Easy to use, fully managed                                     │
│  ├── ❌ Costs money                                                    │
│  ├── ❌ Another service to manage                                      │
│  └── ❌ Data leaves your control                                       │
│                                                                         │
│  WEAVIATE / MILVUS (Self-hosted)                                       │
│  ├── ✅ Feature-rich                                                   │
│  ├── ❌ Another database to run and maintain                           │
│  └── ❌ More complexity                                                │
│                                                                         │
│  PGVECTOR (PostgreSQL extension) ✅                                    │
│  ├── ✅ Same database we already use                                   │
│  ├── ✅ No extra cost or service                                       │
│  ├── ✅ SQL queries with vectors                                       │
│  ├── ✅ Data stays in one place                                        │
│  ├── ✅ Transactions work (vector + regular data)                      │
│  └── ❌ Not as fast for huge scale (millions of vectors)              │
│                                                                         │
│  FOR OUR SCALE: pgvector is perfect                                    │
│  (Hundreds of products, not millions)                                   │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### Ollama vs OpenAI

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    LLM PROVIDER STRATEGY                                │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  DEVELOPMENT: Ollama (Local)                                           │
│  ├── Free (no API costs while developing)                              │
│  ├── Fast iteration (no rate limits)                                   │
│  ├── Works offline                                                     │
│  ├── Privacy (data stays local)                                        │
│  └── Model: llama3.2 or mistral                                        │
│                                                                         │
│  PRODUCTION: OpenAI (or keep Ollama)                                   │
│  ├── Better Hebrew support (GPT-4)                                     │
│  ├── More reliable                                                     │
│  ├── Faster response                                                   │
│  └── Cost: ~$0.01-0.03 per conversation                               │
│                                                                         │
│  ABSTRACTION: Interface allows switching                               │
│  ├── type LLMProvider interface { Complete(), Embed() }               │
│  ├── OllamaProvider implements LLMProvider                             │
│  ├── OpenAIProvider implements LLMProvider                             │
│  └── Switch via config, not code changes                               │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 6. Technology Stack Summary

| Component | Choice | Why This | Why Not Alternative |
|-----------|--------|----------|---------------------|
| **Language** | Go | Fast, simple, great for APIs, learning goal | Python (slower), Java (verbose), Node (less type-safe) |
| **Database** | PostgreSQL | ACID, complex queries, pgvector | MongoDB (no strong schema), MySQL (less features) |
| **Cache** | Redis | Fast, versatile, queue support | Memcached (less features) |
| **Vector DB** | pgvector | Same DB, simple | Pinecone (cost), Weaviate (complexity) |
| **LLM Dev** | Ollama | Free, local, private | OpenAI (costs during dev) |
| **LLM Prod** | OpenAI/Ollama | Hebrew quality, reliability | Depends on needs |
| **API** | REST | Simple, universal | GraphQL (overkill), gRPC (not browser-friendly) |
| **Auth** | JWT | Stateless, standard | Sessions (need sticky sessions) |
| **Container** | Docker | Standard, K8s compatible | Podman (less ecosystem) |
| **Orchestration** | Kubernetes | Learning goal, production-ready | Docker Compose (not production-grade) |
| **Frontend** | TBD (Next.js likely) | SSR, React ecosystem | Flutter (learning curve), Vue (smaller ecosystem) |

---

## 7. Service Boundaries - Why Split This Way?

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    SERVICE BOUNDARIES                                   │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  INVENTORY SERVICE                                                      │
│  ├── Products, Categories, Stock, Movements                            │
│  ├── WHY together? They're tightly coupled                             │
│  │   - Product change affects stock                                    │
│  │   - Movement creates stock change                                   │
│  │   - Often queried together                                          │
│  └── Single DB transaction for consistency                             │
│                                                                         │
│  AUTH SERVICE (Separate)                                               │
│  ├── Users, Tokens, Permissions                                        │
│  ├── WHY separate?                                                     │
│  │   - Security isolation                                              │
│  │   - Different scaling needs                                         │
│  │   - Could be shared with other apps                                 │
│  └── Token validation is very frequent, can optimize                   │
│                                                                         │
│  AI SERVICE (Separate)                                                 │
│  ├── Chat, Intent detection, RAG                                       │
│  ├── WHY separate?                                                     │
│  │   - Heavy compute (maybe GPU)                                       │
│  │   - Different technology (might need Python)                        │
│  │   - Different scaling (auto-scale on load)                          │
│  └── Can fail without breaking core inventory                          │
│                                                                         │
│  NOTIFICATION SERVICE (Separate)                                       │
│  ├── Push, WhatsApp, Email                                             │
│  ├── WHY separate?                                                     │
│  │   - Async (doesn't block main flow)                                 │
│  │   - External dependencies (WhatsApp API)                            │
│  │   - Different failure modes                                         │
│  └── Queue-based, processes when ready                                 │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 8. Questions to Verify Understanding

Before we code, can you answer these?

### Database
1. Why can't we use MongoDB for the main data?
2. What happens if we don't use transactions when taking stock?
3. Why store boxes and units separately instead of total?

### Architecture  
4. What's the benefit of AI service being separate?
5. If Notification service is down, what still works?
6. Why do we use Redis in addition to PostgreSQL?

### AI
7. What is RAG and why do we need it?
8. Why use Ollama for development but maybe OpenAI for production?

### Data Model
9. Why do we need `brand`, `size`, `container_type` as separate fields?
10. Why track every stock movement instead of just current quantity?

---

Take your time to review this. Any questions about specific decisions? 
